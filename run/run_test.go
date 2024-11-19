package run

import (
	"encoding/json"
	"github.com/flowdev/fdialog/ui"
	"io"
	"math"
	"os"
	"reflect"
	"testing"
)

func TestWrite(t *testing.T) {
	type uiValue struct {
		value any
		key   string
		id    string
		name  string
		group string
	}
	tests := []struct {
		name        string
		gotIDs      map[string]string // map from ID to fullName
		gotUIValues []uiValue
		gotDescr    ui.AttributesDescr
		wantOutput  map[string]any
	}{
		{
			name:        "empty",
			gotUIValues: []uiValue{},
			gotDescr:    ui.AttributesDescr{},
			wantOutput:  nil,
		}, {
			name:        "simple-value",
			gotUIValues: []uiValue{{value: 1, name: "k"}},
			gotDescr:    ui.AttributesDescr{"outputKey": "l", "fullName": "k"},
			wantOutput:  map[string]any{"l": float64(1)},
		}, {
			name:        "simple-id-value",
			gotIDs:      map[string]string{"k": "l"},
			gotUIValues: []uiValue{{value: 2, key: "l", group: ""}},
			gotDescr:    ui.AttributesDescr{"outputKey": "m", "id": "k"},
			wantOutput:  map[string]any{"m": float64(2)},
		}, {
			name:        "simple-group-value",
			gotUIValues: []uiValue{{value: 3, key: "k", group: "g"}},
			gotDescr:    ui.AttributesDescr{"group": "g"},
			wantOutput:  map[string]any{"k": float64(3)},
		}, {
			name:        "simple-nested-value",
			gotUIValues: []uiValue{{value: 4, name: "l.m.n", group: "g"}},
			gotDescr:    ui.AttributesDescr{"group": "g"},
			wantOutput:  map[string]any{"l": map[string]any{"m": map[string]any{"n": float64(4)}}},
		}, {
			name:        "simple-nested-id-value",
			gotIDs:      map[string]string{"i": "j.k"},
			gotUIValues: []uiValue{{value: 5, name: "j.k"}},
			gotDescr:    ui.AttributesDescr{"outputKey": "m.n.o", "id": "i"},
			wantOutput:  map[string]any{"m": map[string]any{"n": map[string]any{"o": float64(5)}}},
		}, {
			name:        "simple-nested-group-value",
			gotUIValues: []uiValue{{value: 6, key: "i.j.k.l", group: "g"}},
			gotDescr:    ui.AttributesDescr{"group": "g"},
			wantOutput:  map[string]any{"i": map[string]any{"j": map[string]any{"k": map[string]any{"l": float64(6)}}}},
		}, {
			name: "multiple-nested-group-values",
			gotUIValues: []uiValue{
				{value: 7, key: "i.j.k.l", group: "g"},
				{value: 8, key: "i.j.m.n", group: "g"},
			},
			gotDescr: ui.AttributesDescr{"group": "g"},
			wantOutput: map[string]any{"i": map[string]any{"j": map[string]any{"k": map[string]any{"l": float64(7)},
				"m": map[string]any{"n": float64(8)}}}},
		}, {
			name:   "wild-mix",
			gotIDs: map[string]string{"a": "i.n"},
			gotUIValues: []uiValue{
				{value: 9, key: "i.j.k", group: "g"},
				{value: 10, name: "i.l.m", group: "g"},
				{value: 11, id: "a", name: "i.n", group: "g"},
				{value: 12, key: "i.o.k", group: "g"},
				{value: 13, name: "i.p"},
			},
			gotDescr: ui.AttributesDescr{"group": "g"},
			wantOutput: map[string]any{"i": map[string]any{"j": map[string]any{"k": float64(9)},
				"l": map[string]any{"m": float64(10)}, "o": map[string]any{"k": float64(12)}}, "a": float64(11)},
		}, {
			name:        "simple-bool",
			gotUIValues: []uiValue{{value: true, name: "b"}},
			gotDescr:    ui.AttributesDescr{"outputKey": "l", "fullName": "b"},
			wantOutput:  map[string]any{"l": true},
		}, {
			name:        "simple-string",
			gotUIValues: []uiValue{{value: "bla", name: "s"}},
			gotDescr:    ui.AttributesDescr{"outputKey": "l", "fullName": "s"},
			wantOutput:  map[string]any{"l": "bla"},
		}, {
			name:        "simple-float",
			gotUIValues: []uiValue{{value: math.Pi, name: "p"}},
			gotDescr:    ui.AttributesDescr{"outputKey": "l", "fullName": "p"},
			wantOutput:  map[string]any{"l": math.Pi},
		}, {
			name:        "list-string",
			gotUIValues: []uiValue{{value: []any{"abc", "def"}, name: "s"}},
			gotDescr:    ui.AttributesDescr{"outputKey": "l", "fullName": "s"},
			wantOutput:  map[string]any{"l": []any{"abc", "def"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.CreateTemp("", "example")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(f.Name()) // clean up
			os.Stdout = f

			ui.DeleteAllValues()
			for id, name := range tt.gotIDs {
				_ = ui.RegisterID(id, name)
			}
			for _, v := range tt.gotUIValues {
				ui.StoreValue(v.value, v.key, v.id, v.name, v.group)
			}

			Write(tt.gotDescr, tt.name, nil, nil)

			output := fileContent(t, f)
			if !reflect.DeepEqual(output, tt.wantOutput) {
				t.Errorf("Output got %v, want %v", output, tt.wantOutput)
			}
		})
	}
}

func fileContent(t *testing.T, f *os.File) map[string]any {
	if err := f.Sync(); err != nil {
		t.Fatal(err)
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	result := make(map[string]any)
	err := json.NewDecoder(f).Decode(&result)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		t.Fatal(err)
	}
	return result
}
