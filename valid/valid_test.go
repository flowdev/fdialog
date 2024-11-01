package valid_test

import (
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
			givenUiDescr: ui.CommandsDescr{},
			givenStrict:  true,
			expectedOK:   true,
		}, {
			name: "oneMinimalInfo",
			givenUiDescr: ui.CommandsDescr{
				"info1": {
					"keyword": "dialog",
					"type":    "info",
					"message": "Message for you.",
				},
			},
			givenStrict: true,
			expectedOK:  true,
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
			givenStrict: true,
			expectedOK:  true,
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
			givenStrict: true,
			expectedOK:  true,
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
			givenStrict: true,
			expectedOK:  true,
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
			givenStrict: true,
			expectedOK:  true,
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
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "allMissing",
			givenUiDescr: ui.CommandsDescr{
				"": {},
			},
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "typeMissing",
			givenUiDescr: ui.CommandsDescr{
				"dialog1": {
					"keyword": "dialog",
				},
			},
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "wrongKeyword",
			givenUiDescr: ui.CommandsDescr{
				"dialog1": {
					"keyword": "dialogue",
					"type":    "info",
				},
			},
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "wrongName",
			givenUiDescr: ui.CommandsDescr{
				"dia.log1": {
					"keyword": "dialog",
					"type":    "info",
					"message": "Your info",
				},
			},
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "wrongType",
			givenUiDescr: ui.CommandsDescr{
				"dialog1": {
					"keyword": "dialog",
					"type":    "inf",
				},
			},
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "infoMessageMissing",
			givenUiDescr: ui.CommandsDescr{
				"info6": {
					"keyword": "dialog",
					"type":    "info",
				},
			},
			givenStrict: true,
			expectedOK:  false,
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
			givenStrict: true,
			expectedOK:  false,
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
			givenStrict: true,
			expectedOK:  false,
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
			givenStrict: true,
			expectedOK:  false,
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
			givenStrict: false,
			expectedOK:  true,
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
			givenStrict: false,
			expectedOK:  true,
		}, {
			name: "minimalWindow",
			givenUiDescr: ui.CommandsDescr{
				"win1": {
					"keyword": "window",
				},
			},
			givenStrict: true,
			expectedOK:  true,
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
			givenStrict: true,
			expectedOK:  true,
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
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "minimalLink",
			givenUiDescr: ui.CommandsDescr{
				"link1": {
					"keyword":     "link",
					"destination": "info1",
				},
			},
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "maximalLink",
			givenUiDescr: ui.CommandsDescr{
				"link2": {
					"keyword":     "link",
					"type":        "local",
					"destination": "main.info1",
				},
			},
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "wrongLink",
			givenUiDescr: ui.CommandsDescr{
				"link3": {
					"keyword":     "link",
					"destination": "main:info1",
				},
			},
			givenStrict: true,
			expectedOK:  false,
		}, {
			name: "minimalAction",
			givenUiDescr: ui.CommandsDescr{
				"act1": {
					"keyword": "action",
					"type":    "exit",
				},
			},
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "maximalAction",
			givenUiDescr: ui.CommandsDescr{
				"act2": {
					"keyword": "action",
					"type":    "exit",
					"code":    int64(125),
				},
			},
			givenStrict: true,
			expectedOK:  true,
		}, {
			name: "wrongAction",
			givenUiDescr: ui.CommandsDescr{
				"act2": {
					"keyword": "action",
					"type":    "exit",
					"code":    int64(-1),
				},
			},
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
