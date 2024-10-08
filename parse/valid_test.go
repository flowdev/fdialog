package parse_test

import (
	"strings"
	"testing"

	"github.com/flowdev/fdialog/parse"
)

func TestValidate(t *testing.T) {
	specs := []struct {
		name             string
		givenUiDescr     map[string]map[string]any
		givenStrict      bool
		expectedErrCount int
	}{
		{
			name:             "empty",
			givenUiDescr:     map[string]map[string]any{},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "oneMinimalInfo",
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
				"confirm1": {
					"keyword":     "dialog",
					"type":        "confirmation",
					"title":       "Please Confirm",
					"message":     "Do you want to confirm?",
					"dismissText": "Oh, no!",
					"confirmText": "Yes, please.",
					"width":       float64(240),
					"height":      float64(200),
					"children": map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
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
					"children": map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
				"": {},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "typeMissing",
			givenUiDescr: map[string]map[string]any{
				"dialog1": {
					"keyword": "dialog",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongKeyword",
			givenUiDescr: map[string]map[string]any{
				"dialog1": {
					"keyword": "dialogue",
					"type":    "info",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongName",
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
				"dialog1": {
					"keyword": "dialog",
					"type":    "inf",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "infoMessageMissing",
			givenUiDescr: map[string]map[string]any{
				"info6": {
					"keyword": "dialog",
					"type":    "info",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongInfo",
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
				"win1": {
					"keyword": "window",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "windowWithConfirmation",
			givenUiDescr: map[string]map[string]any{
				"main": {
					"keyword": "window",
					"title":   "Confirmation",
					"width":   float64(400),
					"height":  float64(200),
					"children": map[string]map[string]any{
						"confirm3": {
							"keyword":     "dialog",
							"type":        "confirmation",
							"message":     "Do you want to confirm?",
							"dismissText": "Oh, no!",
							"confirmText": "Yes, please.",
							"width":       float64(400),
							"height":      float64(200),
							"children": map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
				"main": {
					"keyword": "window",
					"title":   "Confirmation",
					"width":   float64(400),
					"height":  float64(200),
					"children": map[string]map[string]any{
						"confirm3": {
							"keyword":     "dialog",
							"type":        "confirmation",
							"message":     "Do you want to confirm?",
							"dismissText": "Oh, no!",
							"confirmText": "Yes, please.",
							"width":       float64(400),
							"height":      float64(200),
							"children": map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
				"link1": {
					"keyword":     "link",
					"destination": "info1",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "maximalLink",
			givenUiDescr: map[string]map[string]any{
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
			givenUiDescr: map[string]map[string]any{
				"link3": {
					"keyword":     "link",
					"destination": "main:info1",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "minimalAction",
			givenUiDescr: map[string]map[string]any{
				"act1": {
					"keyword": "action",
					"type":    "exit",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "maximalAction",
			givenUiDescr: map[string]map[string]any{
				"act2": {
					"keyword": "action",
					"type":    "exit",
					"code":    int64(127),
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "wrongAction",
			givenUiDescr: map[string]map[string]any{
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
			err := parse.Validate(spec.givenUiDescr, spec.givenStrict)
			var actualErrCount int
			if err != nil {
				actualErrCount = strings.Count(err.Error(), "\n") + 1
			}
			if actualErrCount != spec.expectedErrCount {
				tt.Errorf("Validate() expectedErrCount %v, actual errors: %v", spec.expectedErrCount, err)
			}
		})
	}
}
