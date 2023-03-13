// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stringfmt

import (
	"fmt"
	"strings"
	"testing"
)

func TestQuoteAndSplitString(t *testing.T) {
	testData := []struct {
		Input    string
		Length   int
		Expected []string
	}{
		{
			Input:  "This is an example.",
			Length: 2,
			Expected: []string{
				"> Th",
				"> is",
				">  i",
				"> s ",
				"> an",
				">  e",
				"> xa",
				"> mp",
				"> le",
				"> .",
			},
		},
		{
			Input:  "Short line.",
			Length: 100,
			Expected: []string{
				"> Short line.",
			},
		},
	}
	for i, v := range testData {
		t.Run(fmt.Sprintf("Iteration %d..", i), func(t *testing.T) {
			actual := QuoteAndSplitString(v.Input, ">", v.Length)
			if len(actual) != len(v.Expected) {
				t.Fatalf("expected %d lines but got %d lines: [%s]", len(v.Expected), len(actual), strings.Join(actual, ", "))
			}
			for line := 0; line < len(actual); line++ {
				t.Logf("line %d", line)
				actualLine := actual[line]
				expectedLine := v.Expected[line]
				if actualLine != expectedLine {
					t.Fatalf("lines differ on line %d - expected %q but got %q", line, expectedLine, actualLine)
				}
			}
		})
	}
}
