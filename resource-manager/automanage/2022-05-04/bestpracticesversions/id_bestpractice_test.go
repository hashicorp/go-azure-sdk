package bestpracticesversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BestPracticeId{}

func TestNewBestPracticeID(t *testing.T) {
	id := NewBestPracticeID("bestPracticeName")

	if id.BestPracticeName != "bestPracticeName" {
		t.Fatalf("Expected %q but got %q for Segment 'BestPracticeName'", id.BestPracticeName, "bestPracticeName")
	}
}

func TestFormatBestPracticeID(t *testing.T) {
	actual := NewBestPracticeID("bestPracticeName").ID()
	expected := "/providers/Microsoft.AutoManage/bestPractices/bestPracticeName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBestPracticeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BestPracticeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AutoManage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AutoManage/bestPractices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeName",
			Expected: &BestPracticeId{
				BestPracticeName: "bestPracticeName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBestPracticeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BestPracticeName != v.Expected.BestPracticeName {
			t.Fatalf("Expected %q but got %q for BestPracticeName", v.Expected.BestPracticeName, actual.BestPracticeName)
		}

	}
}

func TestParseBestPracticeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BestPracticeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AutoManage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AutoManage/bestPractices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/bEsTpRaCtIcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeName",
			Expected: &BestPracticeId{
				BestPracticeName: "bestPracticeName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/bEsTpRaCtIcEs/bEsTpRaCtIcEnAmE",
			Expected: &BestPracticeId{
				BestPracticeName: "bEsTpRaCtIcEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/bEsTpRaCtIcEs/bEsTpRaCtIcEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBestPracticeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BestPracticeName != v.Expected.BestPracticeName {
			t.Fatalf("Expected %q but got %q for BestPracticeName", v.Expected.BestPracticeName, actual.BestPracticeName)
		}

	}
}

func TestSegmentsForBestPracticeId(t *testing.T) {
	segments := BestPracticeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BestPracticeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
