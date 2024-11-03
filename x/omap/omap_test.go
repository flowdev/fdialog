package omap_test

import (
	"github.com/flowdev/fdialog/x/omap"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	tests := []struct {
		name        string
		givenSize   int    // for New
		givenKeys   []int  // for Add/Build
		givenValues []bool // for Add/Build
		givenExists []int  // for Exists
		givenNoKey  int    // for Get
		wantOKs     []bool // for Add
		wantLen     int    // for Len
		wantExists  []bool // for Exists
		wantKeys    []int  // for All/Get
		wantValues  []bool // for All/Get
	}{
		{
			name:        "empty",
			givenSize:   0,
			givenKeys:   nil,
			givenValues: nil,
			givenExists: nil,
			givenNoKey:  -1,
			wantOKs:     nil,
			wantLen:     0,
			wantExists:  nil,
			wantKeys:    nil,
			wantValues:  nil,
		}, {
			name:        "one-entry",
			givenSize:   1,
			givenKeys:   []int{1},
			givenValues: []bool{true},
			givenExists: []int{1, 2},
			givenNoKey:  -1,
			wantOKs:     []bool{true},
			wantLen:     1,
			wantExists:  []bool{true, false},
			wantKeys:    []int{1},
			wantValues:  []bool{true},
		}, {
			name:        "double-entry",
			givenSize:   2,
			givenKeys:   []int{1, 1},
			givenValues: []bool{true, false},
			givenExists: []int{2, 1},
			givenNoKey:  -1,
			wantOKs:     []bool{true, false},
			wantLen:     1,
			wantExists:  []bool{false, true},
			wantKeys:    []int{1},
			wantValues:  []bool{true},
		}, {
			name:        "size-too-small",
			givenSize:   0,
			givenKeys:   []int{1, 2, 3},
			givenValues: []bool{true, true, true},
			givenExists: []int{1, 2, 3},
			givenNoKey:  -1,
			wantOKs:     []bool{true, true, true},
			wantLen:     3,
			wantExists:  []bool{true, true, true},
			wantKeys:    []int{1, 2, 3},
			wantValues:  []bool{true, true, true},
		}, {
			name:        "size-too-big",
			givenSize:   32,
			givenKeys:   []int{1, 2, 3},
			givenValues: []bool{true, true, true},
			givenExists: []int{1, 2, 3},
			givenNoKey:  -1,
			wantOKs:     []bool{true, true, true},
			wantLen:     3,
			wantExists:  []bool{true, true, true},
			wantKeys:    []int{1, 2, 3},
			wantValues:  []bool{true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			om := omap.New[int, bool](tt.givenSize)
			for i := 0; i < len(tt.givenKeys); i++ {
				ok := om.Add(tt.givenKeys[i], tt.givenValues[i])
				if ok != tt.wantOKs[i] {
					t.Errorf("index %d: Add() = %t, want %t", i, ok, tt.wantOKs[i])
				}
			}
			if om.Len() != tt.wantLen {
				t.Errorf("Len() = %d, want %d", om.Len(), tt.wantLen)
			}
			for i := 0; i < len(tt.givenExists); i++ {
				if actExists := om.Exists(tt.givenExists[i]); actExists != tt.wantExists[i] {
					t.Errorf("index %d: Exists() = %t, want %t", i, actExists, tt.wantExists[i])
				}
			}
			i := 0
			for actKey, actValue := range om.All() {
				if actKey != tt.wantKeys[i] {
					t.Errorf("index %d: All(key) = %v, want %v", i, actKey, tt.wantKeys[i])
				}
				if actValue != tt.wantValues[i] {
					t.Errorf("index %d: All(value) = %v, want %v", i, actValue, tt.wantValues[i])
				}
				i++
			}
			if actValue, actOK := om.Get(tt.givenNoKey); actOK || actValue {
				t.Errorf("Get(no key) = (%v, %t), want (false, false)", actValue, actOK)
			}
			for i := 0; i < len(tt.wantKeys); i++ {
				if actValue, actOK := om.Get(tt.wantKeys[i]); !actOK || actValue != tt.wantValues[i] {
					t.Errorf("index %d: Get() = (%v, %t), want (%v, true)", i, actValue, actOK, tt.wantValues[i])
				}
			}

			// Test Build() with new map:
			om2 := omap.New[int, bool](tt.givenSize)
			for i := 0; i < len(tt.givenKeys); i++ {
				om2 = om2.Build(tt.givenKeys[i], tt.givenValues[i])
			}
			if om2.Len() != tt.wantLen {
				t.Errorf("After Build(): Len() = %d, want %d", om2.Len(), tt.wantLen)
			}
			i = 0
			for actKey, actValue := range om2.All() {
				if actKey != tt.wantKeys[i] {
					t.Errorf("After Build(): index %d: All(key) = %v, want %v", i, actKey, tt.wantKeys[i])
				}
				if actValue != tt.wantValues[i] {
					t.Errorf("After Build(): index %d: All(value) = %v, want %v", i, actValue, tt.wantValues[i])
				}
				i++
			}
		})
	}
}
