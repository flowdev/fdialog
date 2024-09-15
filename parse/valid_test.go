package parse

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	specs := []struct {
		name             string
		givenUiDescr     []map[string]any
		givenStrict      bool
		expectedErrCount int
	}{
		{
			name:             "empty",
			givenUiDescr:     []map[string]any{},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "oneMinimalInfo",
			givenUiDescr: []map[string]any{
				{
					"keyword": "window",
					"name":    "info1",
					"type":    "info",
					"message": "Message for you.",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "oneMaximalInfo",
			givenUiDescr: []map[string]any{
				{
					"keyword":    "window",
					"name":       "info2",
					"type":       "info",
					"message":    "Message for you.",
					"title":      "My Info",
					"buttonText": "Okay...",
					"width":      int64(240),
					"height":     int64(200),
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "oneMaximalError",
			givenUiDescr: []map[string]any{
				{
					"keyword":    "window",
					"name":       "error1",
					"type":       "error",
					"message":    "Error for you.",
					"buttonText": "Oh, shit...",
					"width":      int64(240),
					"height":     int64(200),
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "oneMaximalConfirmation",
			givenUiDescr: []map[string]any{
				{
					"keyword":     "window",
					"name":        "confirm1",
					"type":        "confirmation",
					"title":       "Please Confirm",
					"message":     "Do you want to confirm?",
					"dismissText": "Oh, no!",
					"confirmText": "Yes, please.",
					"width":       int64(240),
					"height":      int64(200),
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "twoInfos",
			givenUiDescr: []map[string]any{
				{
					"keyword": "window",
					"name":    "info3",
					"type":    "info",
					"message": "Info no. one",
				}, {
					"keyword": "window",
					"name":    "info4",
					"type":    "info",
					"message": "Info no. two",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "infoErrorConfirmation",
			givenUiDescr: []map[string]any{
				{
					"keyword": "window",
					"name":    "info5",
					"type":    "info",
					"message": "Info no. three (or five?)",
				}, {
					"keyword": "window",
					"name":    "error2",
					"type":    "error",
					"message": "Error no. two",
				}, {
					"keyword": "window",
					"name":    "confirm2",
					"type":    "confirmation",
					"message": "Please confirm (no. two)",
				},
			},
			givenStrict:      true,
			expectedErrCount: 0,
		}, {
			name: "allMissing",
			givenUiDescr: []map[string]any{
				{},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "nameAndTypeMissing",
			givenUiDescr: []map[string]any{
				{
					"keyword": "window",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "typeMissing",
			givenUiDescr: []map[string]any{
				{
					"keyword": "window",
					"name":    "win1",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongKeyword",
			givenUiDescr: []map[string]any{
				{
					"keyword": "windoof",
					"name":    "win1",
					"type":    "info",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongType",
			givenUiDescr: []map[string]any{
				{
					"keyword": "window",
					"name":    "win1",
					"type":    "inf",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "infoMessageMissing",
			givenUiDescr: []map[string]any{
				{
					"keyword": "window",
					"name":    "info6",
					"type":    "info",
				},
			},
			givenStrict:      true,
			expectedErrCount: 1,
		}, {
			name: "wrongInfo",
			givenUiDescr: []map[string]any{
				{
					"keyword":    "window",
					"name":       "",
					"type":       "info",
					"message":    "",
					"title":      "",
					"buttonText": "",
					"width":      int64(49),
					"height":     int64(79),
				},
			},
			givenStrict:      true,
			expectedErrCount: 6,
		}, {
			name: "wrongError",
			givenUiDescr: []map[string]any{
				{
					"keyword":    "window",
					"name":       "",
					"type":       "error",
					"message":    "",
					"buttonText": "",
					"width":      int64(49),
					"height":     int64(79),
				},
			},
			givenStrict:      true,
			expectedErrCount: 5,
		}, {
			name: "wrongConfirmation",
			givenUiDescr: []map[string]any{
				{
					"keyword":     "window",
					"name":        "",
					"type":        "confirmation",
					"title":       "",
					"message":     "",
					"dismissText": "",
					"confirmText": "",
					"width":       int64(49),
					"height":      int64(79),
				},
			},
			givenStrict:      true,
			expectedErrCount: 7,
		},
	}

	for _, spec := range specs {
		t.Run(spec.name, func(tt *testing.T) {
			err := Validate(spec.givenUiDescr, spec.givenStrict)
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
