package uidl

import (
	"fmt"
	"github.com/flowdev/fdialog/parse"
	"github.com/flowdev/gparselib"
	"io"
	"strconv"
)

const UIDLVersion = 1

func ParseUIDL(input io.Reader) (map[string]map[string]any, error) {
	inputData, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	_ = inputData
	return nil, nil
}

func parseUIDL(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{
		parseVersion,
		parseCommands,
		gparselib.NewParseEOFPlugin(nil),
	}, func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
		versionResult := pd.SubResults[0]
		version := versionResult.Value.(int64)
		if version != 1 {
			pd.AddError(
				versionResult.Pos,
				fmt.Sprintf("expected version %d, got: %d", UIDLVersion, version),
				nil,
			)
		}
		pd.Result.Value = pd.SubResults[1].Value
		return pd, ctx
	})
}

func parseVersion(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	pMustSpace := gparselib.NewParseSpacePlugin(nil, true)
	pVersionLiteral := gparselib.NewParseLiteralPlugin(nil, "version")
	pVersionNumber, err := gparselib.NewParseNaturalPlugin(nil, 10)
	if err != nil {
		return nil, err
	}

	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{
		parseSpaceComment,
		pVersionLiteral,
		pMustSpace,
		pVersionNumber,
		pMustSpace,
		parseSpaceComment,
	}, func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
		pd.Result.Value = pd.SubResults[3].Value
		return pd, ctx
	})
}

func parseCommand(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{
		gparselib.NewParseIdentPlugin(textSemantic, "", ""),
		parseSpaceComment,
		gparselib.NewParseIdentPlugin(textSemantic, "_", "_"),
		parseSpaceComment,
		parseAttributes,
		gparselib.NewParseOptionalPlugin(parseCommandBody, nil),
	}, func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
		pd.Result.Value = commandValue{
			keyword:      pd.SubResults[0].Value.(string),
			name:         pd.SubResults[2].Value.(string),
			attributeMap: pd.SubResults[4].Value.(map[string]any),
			// TODO: add pd.SubResults[5].Value.(map[string]map[string]any) as key "children" to map
		}
		return pd, ctx
	})
}

func parseCommands(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseMulti0(pd, ctx, parseCommand,
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			// convert all commands to a map
			commandMap := make(map[string]map[string]any, len(pd.SubResults))

			for _, subResult := range pd.SubResults {
				command := subResult.Value.(commandValue)
				if _, ok := commandMap[command.name]; ok {
					pd.AddError(subResult.Pos,
						fmt.Sprintf("duplicate command name: %q", command.name),
						nil,
					)
					continue
				}
				commandMap[command.name] = command.attributeMap
			}

			pd.Result.Value = commandMap
			return pd, ctx
		})
}

// CommandBody        <- '{' Spacing Command '}' Spacing
func parseCommandBody(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{
		gparselib.NewParseLiteralPlugin(nil, "{"),
		parseSpaceComment,
		parseCommands,
		gparselib.NewParseLiteralPlugin(nil, "}"),
		parseSpaceComment,
	}, func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
		pd.Result.Value = pd.SubResults[2].Value
		return pd, ctx
	})
}

func parseAttributes(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	pParenOpen := gparselib.NewParseLiteralPlugin(nil, "(")
	pParenClose := gparselib.NewParseLiteralPlugin(nil, ")")
	pComma := gparselib.NewParseLiteralPlugin(nil, ",")

	pNoAttributes := gparselib.NewParseAllPlugin([]gparselib.SubparserOp{
		pParenOpen,
		parseSpaceComment,
		pParenClose,
		parseSpaceComment,
	}, nil)

	pCommaAttribute := gparselib.NewParseAllPlugin([]gparselib.SubparserOp{
		pComma,
		parseSpaceComment,
		parseAttribute,
	}, func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
		pd.Result.Value = pd.SubResults[2].Value
		return pd, ctx
	})
	pManyAttributes := gparselib.NewParseMulti0Plugin(pCommaAttribute,
		func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
			allAttrs := make([]attributeValue, len(pd.SubResults))
			for i, subResult := range pd.SubResults { // copy over for the type
				allAttrs[i] = subResult.Value.(attributeValue)
			}
			pd.Result.Value = allAttrs
			return pd, ctx
		})
	pAllAttributes := gparselib.NewParseAllPlugin([]gparselib.SubparserOp{
		pParenOpen,
		parseSpaceComment,
		parseAttribute,
		pManyAttributes,
		pParenClose,
		parseSpaceComment,
	}, func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
		// convert all attributes to a map
		manyAttrs := pd.SubResults[3].Value.([]attributeValue)
		attrMap := make(map[string]any, len(manyAttrs))
		attr := pd.SubResults[2].Value.(attributeValue)
		if checkAttributeKey(attr.key, pd) {
			attrMap[attr.key] = attr.value
		}

		for _, attr = range manyAttrs {
			if !checkAttributeKey(attr.key, pd) {
				continue
			}
			if _, ok := attrMap[attr.key]; ok {
				pd.AddError(pd.SubResults[3].Pos,
					fmt.Sprintf("duplicate attribute key: %q", attr.key),
					nil,
				)
				continue
			}
			attrMap[attr.key] = attr.value
		}

		pd.Result.Value = attrMap
		return pd, ctx
	})

	return gparselib.ParseAny(pd, ctx, []gparselib.SubparserOp{
		pNoAttributes,
		pAllAttributes,
	}, nil)
}

func checkAttributeKey(attributeKey string, pd *gparselib.ParseData) bool {
	reservedAttributeKeys := map[string]bool{
		parse.KeyKeyword:  true,
		parse.KeyChildren: true,
		parse.KeyType:     true,
		parse.KeyName:     true,
	}
	if reservedAttributeKeys[attributeKey] {
		pd.AddError(
			pd.Result.Pos,
			fmt.Sprintf("reserved key %q can't be used as an attribute name", attributeKey),
			nil,
		)
		return false
	}
	return true
}

func parseAttribute(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{
		gparselib.NewParseIdentPlugin(textSemantic, "_", "_"),
		parseSpaceComment,
		gparselib.NewParseLiteralPlugin(nil, "="),
		parseSpaceComment,
		parseValue,
		parseSpaceComment,
	}, func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
		pd.Result.Value = attributeValue{
			key:   pd.SubResults[0].Value.(string),
			value: pd.SubResults[4].Value,
		}
		return pd, ctx
	})
}

func parseValue(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	return gparselib.ParseAny(pd, ctx, []gparselib.SubparserOp{
		parseNormalString, parseBacktickString, parseBoolValue, parseFloatValue, parseIntValue,
	}, nil)
}

func parseNormalString(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	pDQuote := gparselib.NewParseLiteralPlugin(nil, `"`)

	afterBackslash := false
	pInnerString := gparselib.NewParseGoodRunesPlugin(textSemantic,
		func(r rune) bool {
			if afterBackslash {
				afterBackslash = false
				return true
			}
			if r == '\\' {
				afterBackslash = true
				return true
			}
			return r != '"'
		})

	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{pDQuote, pInnerString, pDQuote},
		func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
			pd.Result.Value = pd.SubResults[1].Value
			return pd, ctx
		})
}
func parseBacktickString(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	pBacktick := gparselib.NewParseLiteralPlugin(nil, "`")
	pNoBacktick := gparselib.NewParseGoodRunesPlugin(textSemantic,
		func(r rune) bool {
			return r != '`'
		})

	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{pBacktick, pNoBacktick, pBacktick},
		func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
			pd.Result.Value = pd.SubResults[1].Value
			return pd, ctx
		})
}

func parseFloatValue(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	pPlus := gparselib.NewParseLiteralPlugin(nil, "+")
	pMinus := gparselib.NewParseLiteralPlugin(nil, "-")
	pSign := gparselib.NewParseAnyPlugin([]gparselib.SubparserOp{pPlus, pMinus}, nil)
	pOptSign := gparselib.NewParseOptionalPlugin(pSign, nil)
	pNum, err := gparselib.NewParseNaturalPlugin(nil, 10)
	if err != nil {
		panic(err) // can only happen if cfgRadix is invalid
	}
	pDot := gparselib.NewParseLiteralPlugin(nil, ".")
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{pOptSign, pNum, pDot, pNum},
		func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
			pd.Result.Value, err = strconv.ParseFloat(pd.Result.Text, 64)
			if err != nil {
				pd.AddError(pd.Result.Pos, err.Error(), nil)
			}
			return pd, ctx
		})
}

func parseIntValue(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	pPlus := gparselib.NewParseLiteralPlugin(nil, "+")
	pMinus := gparselib.NewParseLiteralPlugin(nil, "-")
	pSign := gparselib.NewParseAnyPlugin([]gparselib.SubparserOp{pPlus, pMinus}, nil)
	pOptSign := gparselib.NewParseOptionalPlugin(pSign, nil)
	pNum, err := gparselib.NewParseNaturalPlugin(nil, 10)
	if err != nil {
		panic(err) // can only happen if cfgRadix is invalid
	}
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{pOptSign, pNum},
		func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
			pd.Result.Value, err = strconv.ParseInt(pd.Result.Text, 10, 64) // TODO: handle overflow?
			if err != nil {
				pd.AddError(pd.Result.Pos, err.Error(), nil)
			}
			return pd, ctx
		})
}

func parseBoolValue(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	pTrue := gparselib.NewParseLiteralPlugin(
		func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
			pd.Result.Value = true
			return pd, ctx
		}, "true")
	pFalse := gparselib.NewParseLiteralPlugin(
		func(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
			pd.Result.Value = false
			return pd, ctx
		}, "false")

	return gparselib.ParseAny(pd, ctx, []gparselib.SubparserOp{pTrue, pFalse}, nil)
}

// parseSpaceComment parses any amount of space (including newline) and line
// (`#` ... <NL>) comments.
//   - Semantic result: The parsed text plus a signal whether a newline was
//     parsed (spaceCommentSemValue).
func parseSpaceComment(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	pSpc := gparselib.NewParseSpacePlugin(textSemantic, true)
	pLnCmnt, err := gparselib.NewParseLineCommentPlugin(textSemantic, `#`)
	if err != nil {
		panic(err) // can only be a programming error!
	}
	pSpaceOrComment := gparselib.NewParseAnyPlugin(
		[]gparselib.SubparserOp{pSpc, pLnCmnt},
		textSemantic,
	)
	return gparselib.ParseMulti0(pd, ctx, pSpaceOrComment, nil)
}

//
// ----------------------------------------------------------------------------
// Semantics:
//

type commandValue struct {
	keyword      string
	name         string
	attributeMap map[string]any
}

type attributeValue struct {
	key   string
	value any
}

// textSemantic stores the successfully parsed text as semantic value.
func textSemantic(pd *gparselib.ParseData, ctx interface{}) (*gparselib.ParseData, interface{}) {
	pd.Result.Value = pd.Result.Text
	return pd, ctx
}
