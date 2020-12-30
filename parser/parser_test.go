package parser

import (
	"github.com/whosonfirst/go-edtf/level0"
	"github.com/whosonfirst/go-edtf/level1"
	"github.com/whosonfirst/go-edtf/level2"
	"testing"
)

func TestLevels(t *testing.T) {

	for label, tests := range level0.Tests {

		for _, str := range tests {

			if !level0.IsLevel0(str) {
				t.Fatalf("Invalid level 0 string '%s' (%s)", str, label)
			}

			/*
				if level1.IsLevel1(str) {
					t.Fatalf("Level 0 string '%s' (%s) parses as level 1 date", str, label)
				}

				if level2.IsLevel2(str) {
					t.Fatalf("Level 0 string '%s' (%s) parses as level 2 date", str, label)
				}
			*/
		}
	}

	for label, tests := range level1.Tests {

		for _, str := range tests {

			if !level1.IsLevel1(str) {
				t.Fatalf("Invalid level 1 string '%s' (%s)", str, label)
			}

			/*
				if level0.IsLevel0(str) {
					t.Fatalf("Level 1 string '%s' (%s) parses as level 0 date", str, label)
				}

				if level2.IsLevel2(str) {
					t.Fatalf("Level 1 string '%s' (%s) parses as level 2 date", str, label)
				}
			*/
		}
	}

	for label, tests := range level2.Tests {

		for _, str := range tests {

			if !level2.IsLevel2(str) {
				t.Fatalf("Invalid level 2 string '%s' (%s)", str, label)
			}

			/*
				if level0.IsLevel0(str) {
					t.Fatalf("Level 2 string '%s' (%s) parses as level 0 date", str, label)
				}

				if level1.IsLevel1(str) {
					t.Fatalf("Level 2 string '%s' (%s) parses as level 1 date", str, label)
				}
			*/
		}
	}

}

func TestIsValid(t *testing.T) {

	for label, tests := range level0.Tests {

		for _, str := range tests {

			if !IsValid(str) {
				t.Fatalf("Invalid level 0 string '%s' (%s)", str, label)
			}
		}
	}

	for label, tests := range level1.Tests {

		for _, str := range tests {

			if !IsValid(str) {
				t.Fatalf("Invalid level 1 string '%s' (%s)", str, label)
			}
		}
	}

	for label, tests := range level2.Tests {

		for _, str := range tests {

			if !IsValid(str) {
				t.Fatalf("Invalid level 2 string '%s' (%s)", str, label)
			}
		}
	}

}

func TestParseString(t *testing.T) {

	for label, tests := range level0.Tests {

		for _, str := range tests {

			_, err := ParseString(str)

			if err != nil {
				t.Fatalf("Invalid level 0 string '%s' (%s), %v", str, label, err)
			}
		}
	}

	for label, tests := range level1.Tests {

		for _, str := range tests {

			_, err := ParseString(str)

			if err != nil {
				t.Fatalf("Invalid level 1 string '%s' (%s), %v", str, label, err)
			}
		}
	}

	for label, tests := range level2.Tests {

		for _, str := range tests {

			_, err := ParseString(str)

			if err != nil {
				t.Fatalf("Invalid level 2 string '%s' (%s), %v", str, label, err)
			}
		}
	}

}
