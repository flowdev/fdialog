package valid_test

import (
	"github.com/flowdev/fdialog/x/omap"
	"testing"

	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/uimain"
	"github.com/flowdev/fdialog/valid"
)

func TestValidate(t *testing.T) {
	_ = uimain.RegisterEverything()

	specs := []struct {
		name         string
		givenUiDescr ui.CommandsDescr
		givenStrict  bool
		expectedOK   bool
	}{
		{
			name:         "empty",
			givenUiDescr: omap.New[string, ui.AttributesDescr](0),
			givenStrict:  true,
			expectedOK:   true,
		}, {
			name: "oneMinimalInfo",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"info1", map[string]any{
					":keyword": "dialog",
					"type":     "info",
					"message":  "Message for you.",
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "oneMaximalInfo",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"info2", map[string]any{
					":keyword":   "dialog",
					"type":       "info",
					"message":    "Message for you.",
					"title":      "My Info",
					"buttonText": "Okay...",
					"width":      float64(240),
					"height":     float64(200),
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "oneMaximalError",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"error1", map[string]any{
					":keyword":   "dialog",
					"type":       "error",
					"message":    "Error for you.",
					"buttonText": "Oh, shit...",
					"width":      float64(240),
					"height":     float64(200),
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "oneMaximalConfirmation",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"confirm1", map[string]any{
					":keyword":    "dialog",
					"type":        "confirmation",
					"title":       "Please Confirm",
					"message":     "Do you want to confirm?",
					"dismissText": "Oh, no!",
					"confirmText": "Yes, please.",
					"width":       float64(240),
					"height":      float64(200),
					":children": omap.New[string, ui.AttributesDescr](2).Build(
						"confirm", map[string]any{
							":keyword": "action",
							"type":    "exit",
							"code":    int64(0),
						}).Build(
						"dismiss", map[string]any{
							":keyword": "action",
							"type":    "exit",
							"code":    int64(1),
						}),
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "twoInfos",
			givenUiDescr: omap.New[string, ui.AttributesDescr](2).Build(
				"info3", map[string]any{
					":keyword": "dialog",
					"type":     "info",
					"message":  "Info no. one",
				}).Build(
				"info4", map[string]any{
					":keyword": "dialog",
					"type":     "info",
					"message":  "Info no. two",
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "infoErrorConfirmation",
			givenUiDescr: omap.New[string, ui.AttributesDescr](3).Build(
				"info5", map[string]any{
					":keyword": "dialog",
					"type":     "info",
					"message":  "Info no. three (or five?)",
				}).Build(
				"error2", map[string]any{
					":keyword": "dialog",
					"type":     "error",
					"message":  "Error no. two",
				}).Build(
				"confirm2", map[string]any{
					":keyword": "dialog",
					"type":     "confirmation",
					"message":  "Please confirm (no. two)",
					":children": omap.New[string, ui.AttributesDescr](2).Build(
						"confirm", map[string]any{
							":keyword": "action",
							"type":     "exit",
							"code":     int64(0),
						}).Build(
						"dismiss", map[string]any{
							":keyword": "action",
							"type":     "exit",
							"code":     int64(1),
						}),
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name:         "allMissing",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build("", map[string]any{}),
			givenStrict:  true,
			expectedOK:   false,
		}, {
			name: "typeMissing",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"dialog1", map[string]any{
					":keyword": "dialog",
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "wrongKeyword",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"dialog1", map[string]any{
					":keyword": "dialogue",
					"type":     "info",
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "wrongName",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"dia.log1", map[string]any{
					":keyword": "dialog",
					"type":     "info",
					"message":  "Your info",
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "wrongType",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"dialog1", map[string]any{
					":keyword": "dialog",
					"type":     "inf",
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "infoMessageMissing",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"info6", map[string]any{
					":keyword": "dialog",
					"type":     "info",
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "wrongInfo",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"", map[string]any{
					":keyword":   "dialog",
					"type":       "info",
					"message":    "",
					"title":      "",
					"buttonText": "",
					"width":      float64(49),
					"height":     float64(79),
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "wrongError",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"", map[string]any{
					":keyword":   "dialog",
					"type":       "error",
					"message":    "",
					"buttonText": "",
					"width":      float64(49),
					"height":     float64(79),
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "wrongConfirmation",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"", map[string]any{
					":keyword":    "dialog",
					"type":        "confirmation",
					"title":       "",
					"message":     "",
					"dismissText": "",
					"confirmText": "",
					"width":       float64(49),
					"height":      float64(79),
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "errorWithTitle",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"error3", map[string]any{
					":keyword":   "dialog",
					"type":       "error",
					"title":      "Error?",
					"message":    "Error for you.",
					"buttonText": "Oh, shit...",
					"width":      float64(240),
					"height":     float64(200),
				}),
			givenStrict: false,
			expectedOK:  true,
		}, {
			name: "infoWithExtraAttrs",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"error4", map[string]any{
					":keyword":          "dialog",
					"type":              "info",
					"titli":             "Info?",
					"message":           "Info for you.",
					"buttonFext":        "Fine",
					"with":              float64(240),
					"heiht":             float64(200),
					"myMadeUpAttribute": "bla",
				}),
			givenStrict: false,
			expectedOK:  true,
		}, {
			name: "minimalWindow",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"win1", map[string]any{
					":keyword": "window",
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "windowWithConfirmation",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"main", map[string]any{
					":keyword": "window",
					"title":    "Confirmation",
					"width":    float64(400),
					"height":   float64(200),
					":children": omap.New[string, ui.AttributesDescr](1).Build(
						"confirm3", map[string]any{
							":keyword":    "dialog",
							"type":        "confirmation",
							"message":     "Do you want to confirm?",
							"dismissText": "Oh, no!",
							"confirmText": "Yes, please.",
							"width":       float64(400),
							"height":      float64(200),
							":children": omap.New[string, ui.AttributesDescr](2).Build(
								"confirm", map[string]any{
									":keyword": "action",
									"type":     "exit",
									"code":     int64(0),
								}).Build(
								"dismiss", map[string]any{
									":keyword": "action",
									"type":     "exit",
									"code":     int64(1),
								}),
						}),
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "windowWithNestedConfirmationError",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"main", map[string]any{
					":keyword": "window",
					"title":    "Confirmation",
					"width":    float64(400),
					"height":   float64(200),
					":children": omap.New[string, ui.AttributesDescr](1).Build(
						"confirm3", map[string]any{
							":keyword":    "dialog",
							"type":        "confirmation",
							"message":     "Do you want to confirm?",
							"dismissText": "Oh, no!",
							"confirmText": "Yes, please.",
							"width":       float64(400),
							"height":      float64(200),
							":children": omap.New[string, ui.AttributesDescr](2).Build(
								"confirm", map[string]any{
									":keyword": "action",
									"type":     "ext",
									"code":     int64(0),
								}).Build(
								"dismiss", map[string]any{
									":keyword": "action",
									"type":     "exit",
									"code":     int64(128),
								}),
						}),
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "minimalLink",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"link1", map[string]any{
					":keyword":    "link",
					"destination": "info1",
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "maximalLink",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"link2", map[string]any{
					":keyword":    "link",
					"type":        "local",
					"destination": "main.info1",
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "wrongLink",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"link3", map[string]any{
					":keyword":    "link",
					"destination": "main:info1",
				}),
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "minimalAction",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"act1", map[string]any{
					":keyword": "action",
					"type":     "exit",
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "maximalAction",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"act2", map[string]any{
					":keyword": "action",
					"type":     "exit",
					"code":     int64(125),
				}),
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "wrongAction",
			givenUiDescr: omap.New[string, ui.AttributesDescr](1).Build(
				"act2", map[string]any{
					":keyword": "action",
					"type":     "exit",
					"code":     int64(-1),
				}),
			givenStrict: true,
			expectedOK:  false,
		},
	}

	for _, spec := range specs {
		t.Run(spec.name, func(tt *testing.T) {
			ok := valid.UIDescription(spec.givenUiDescr, spec.givenStrict)
			if ok != spec.expectedOK {
				tt.Errorf("UIDescription() expectedOK %v, actual OK: %v", spec.expectedOK, ok)
			}
		})
	}
}
