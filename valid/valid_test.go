package valid_test

import (
	"strings"
	"testing"

	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/uimain"
	"github.com/flowdev/fdialog/valid"
)

func TestValidate(t *testing.T) {
	_ = uimain.RegisterEverything()

	specs := []struct {
		name             string
		givenUiDescr     ui.CommandsDescr
		givenStrict      bool
		expectedErrCount int
	}{
		{
			name:             "empty",
			givenUiDescr:     ui.CommandsDescr{},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "oneMinimalInfo",
			givenUiDescr: ui.CommandsDescr{
				"info1": {
					"keyword": "dialog",
					"type":    "info",
					"message": "Message for you.",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "oneMaximalInfo",
			givenUiDescr: ui.CommandsDescr{
				"info2": {
					"keyword":    "dialog",
					"type":       "info",
					"message":    "Message for you.",
					"title":      "My Info",
					"buttonText": "Okay...",
					"width":      float64(240),
					"height":     float64(200),
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "oneMaximalError",
			givenUiDescr: ui.CommandsDescr{
				"error1": {
					"keyword":    "dialog",
					"type":       "error",
					"message":    "Error for you.",
					"buttonText": "Oh, shit...",
					"width":      float64(240),
					"height":     float64(200),
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "oneMaximalConfirmation",
			givenUiDescr: ui.CommandsDescr{
				"confirm1": {
					"keyword":     "dialog",
					"type":        "confirmation",
					"title":       "Please Confirm",
					"message":     "Do you want to confirm?",
					"dismissText": "Oh, no!",
					"confirmText": "Yes, please.",
					"width":       float64(240),
					"height":      float64(200),
					"children": ui.CommandsDescr{
						"confirm": {
							"keyword": "action",
							"type":    "exit",
							"code":    int64(0),
						},
						"dismiss": {
							"keyword": "action",
							"type":    "exit",
							"code":    int64(1),
						},
					},
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "twoInfos",
			givenUiDescr: ui.CommandsDescr{
				"info3": {
					"keyword": "dialog",
					"type":    "info",
					"message": "Info no. one",
				},
				"info4": {
					"keyword": "dialog",
					"type":    "info",
					"message": "Info no. two",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "infoErrorConfirmation",
			givenUiDescr: ui.CommandsDescr{
				"info5": {
					"keyword": "dialog",
					"type":    "info",
					"message": "Info no. three (or five?)",
				},
				"error2": {
					"keyword": "dialog",
					"type":    "error",
					"message": "Error no. two",
				},
				"confirm2": {
					"keyword": "dialog",
					"type":    "confirmation",
					"message": "Please confirm (no. two)",
					"children": ui.CommandsDescr{
						"confirm": {
							"keyword": "action",
							"type":    "exit",
							"code":    int64(0),
						},
						"dismiss": {
							"keyword": "action",
							"type":    "exit",
							"code":    int64(1),
						},
					},
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "allMissing",
			givenUiDescr: ui.CommandsDescr{
				"": {},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "typeMissing",
			givenUiDescr: ui.CommandsDescr{
				"dialog1": {
					"keyword": "dialog",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongKeyword",
			givenUiDescr: ui.CommandsDescr{
				"dialog1": {
					"keyword": "dialogue",
					"type":    "info",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongName",
			givenUiDescr: ui.CommandsDescr{
				"dia.log1": {
					"keyword": "dialog",
					"type":    "info",
					"message": "Your info",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongType",
			givenUiDescr: ui.CommandsDescr{
				"dialog1": {
					"keyword": "dialog",
					"type":    "inf",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "infoMessageMissing",
			givenUiDescr: ui.CommandsDescr{
				"info6": {
					"keyword": "dialog",
					"type":    "info",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongInfo",
			givenUiDescr: ui.CommandsDescr{
				"": {
					"keyword":    "dialog",
					"type":       "info",
					"message":    "",
					"title":      "",
					"buttonText": "",
					"width":      float64(49),
					"height":     float64(79),
				},
			},
			givenStrict:      true,
			expectedErrCount: 6,
		}, {
			name: "wrongError",
			givenUiDescr: ui.CommandsDescr{
				"": {
					"keyword":    "dialog",
					"type":       "error",
					"message":    "",
					"buttonText": "",
					"width":      float64(49),
					"height":     float64(79),
				},
			},
			givenStrict:      true,
			expectedErrCount: 5,
		}, {
			name: "wrongConfirmation",
			givenUiDescr: ui.CommandsDescr{
				"": {
					"keyword":     "dialog",
					"type":        "confirmation",
					"title":       "",
					"message":     "",
					"dismissText": "",
					"confirmText": "",
					"width":       float64(49),
					"height":      float64(79),
				},
			},
			givenStrict:      true,
			expectedErrCount: 8,
		}, {
			name: "errorWithTitle",
			givenUiDescr: ui.CommandsDescr{
				"error3": {
					"keyword":    "dialog",
					"type":       "error",
					"title":      "Error?",
					"message":    "Error for you.",
					"buttonText": "Oh, shit...",
					"width":      float64(240),
					"height":     float64(200),
				},
			},
			givenStrict:      false,
			expectedErrCount: 0,
		}, {
			name: "infoWithExtraAttrs",
			givenUiDescr: ui.CommandsDescr{
				"error4": {
					"keyword":           "dialog",
					"type":              "info",
					"titli":             "Info?",
					"message":           "Info for you.",
					"buttonFext":        "Fine",
					"with":              float64(240),
					"heiht":             float64(200),
					"myMadeUpAttribute": "bla",
				},
			},
			givenStrict:      false,
			expectedErrCount: 0,
		}, {
			name: "minimalWindow",
			givenUiDescr: ui.CommandsDescr{
				"win1": {
					"keyword": "window",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "windowWithConfirmation",
			givenUiDescr: ui.CommandsDescr{
				"main": {
					"keyword": "window",
					"title":   "Confirmation",
					"width":   float64(400),
					"height":  float64(200),
					"children": ui.CommandsDescr{
						"confirm3": {
							"keyword":     "dialog",
							"type":        "confirmation",
							"message":     "Do you want to confirm?",
							"dismissText": "Oh, no!",
							"confirmText": "Yes, please.",
							"width":       float64(400),
							"height":      float64(200),
							"children": ui.CommandsDescr{
								"confirm": {
									"keyword": "action",
									"type":    "exit",
									"code":    int64(0),
								},
								"dismiss": {
									"keyword": "action",
									"type":    "exit",
									"code":    int64(1),
								},
							},
						},
					},
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "windowWithNestedConfirmationError",
			givenUiDescr: ui.CommandsDescr{
				"main": {
					"keyword": "window",
					"title":   "Confirmation",
					"width":   float64(400),
					"height":  float64(200),
					"children": ui.CommandsDescr{
						"confirm3": {
							"keyword":     "dialog",
							"type":        "confirmation",
							"message":     "Do you want to confirm?",
							"dismissText": "Oh, no!",
							"confirmText": "Yes, please.",
							"width":       float64(400),
							"height":      float64(200),
							"children": ui.CommandsDescr{
								"confirm": {
									"keyword": "action",
									"type":    "ext",
									"code":    int64(0),
								},
								"dismiss": {
									"keyword": "action",
									"type":    "exit",
									"code":    int64(128),
								},
							},
						},
					},
				},
			},
			givenStrict:      true,
			expectedErrCount: 2,
		}, {
			name: "minimalLink",
			givenUiDescr: ui.CommandsDescr{
				"link1": {
					"keyword":     "link",
					"destination": "info1",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "maximalLink",
			givenUiDescr: ui.CommandsDescr{
				"link2": {
					"keyword":     "link",
					"type":        "local",
					"destination": "main.info1",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "wrongLink",
			givenUiDescr: ui.CommandsDescr{
				"link3": {
					"keyword":     "link",
					"destination": "main:info1",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "minimalAction",
			givenUiDescr: ui.CommandsDescr{
				"act1": {
					"keyword": "action",
					"type":    "exit",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "maximalAction",
			givenUiDescr: ui.CommandsDescr{
				"act2": {
					"keyword": "action",
					"type":    "exit",
					"code":    int64(125),
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "wrongAction",
			givenUiDescr: ui.CommandsDescr{
				"act2": {
					"keyword": "action",
					"type":    "exit",
					"code":    int64(-1),
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		},
	}

	for _, spec := range specs {
		t.Run(spec.name, func(tt *testing.T) {
			err := valid.UIDescription(spec.givenUiDescr, spec.givenStrict)
			var actualErrCount int
			if err != nil {
				actualErrCount = strings.Count(err.Error(), "\n") + 1
			}
			if actualErrCount != spec.expectedErrCount {
				tt.Errorf("UIDescription() expectedErrCount %v, actual errors: %v", spec.expectedErrCount, err)
			}
		})
	}
}
