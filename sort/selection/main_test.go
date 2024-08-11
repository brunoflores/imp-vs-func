package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
)

func TestMyStringComparable(t *testing.T) {
	tests := []struct {
		left, right MyString
		want        Comp
	}{
		{
			left:  MyString{"A"},
			right: MyString{"A"},
			want:  Eq,
		},
		{
			left:  MyString{"A"},
			right: MyString{"B"},
			want:  Less,
		},
		{
			left:  MyString{"B"},
			right: MyString{"A"},
			want:  More,
		},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("test case #%d", i), func(t *testing.T) {
			if got := tc.left.CompareTo(&tc.right); got != tc.want {
				t.Fatalf("expected: %v, got %v", tc.want, got)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	var testData = "./testdata"
	entries, err := os.ReadDir(testData)
	if err != nil {
		panic(err)
	}
	var fix = func(s []byte) string {
		return strings.TrimSpace(string(s))
	}
	for _, dir := range entries {
		if !dir.IsDir() {
			continue
		}
		fNameExpected := path.Join(testData, dir.Name(), "expected.txt")
		fNameInput := path.Join(testData, dir.Name(), "input.txt")
		t.Run(dir.Name(), func(t *testing.T) {
			input, err := os.ReadFile(fNameInput)
			if err != nil {
				panic(err)
			}
			expected, err := os.ReadFile(fNameExpected)
			if err != nil {
				panic(err)
			}
			want := fix(expected)
			myStrings := Map(strings.Split(fix(input), " "), NewMyString)
			selectionSort := SelectionSort[*MyString]{Items: myStrings}
			selectionSort.Sort()
			if got := selectionSort.Show(); got != want {
				if os.Getenv("PROMOTE") == "y" {
					f, err := os.OpenFile(fNameExpected, os.O_WRONLY, 0600)
					if err != nil {
						panic(err)
					}
					defer f.Close()
					if _, err = f.WriteString(got); err != nil {
						panic(err)
					}
					return
				}
				t.Fatalf(fmt.Sprintf("\ngot:\t%s\nwant:\t%s", got, want))
			}
		})
	}
}
