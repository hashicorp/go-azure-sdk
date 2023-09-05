package bestpracticesversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VersionId{}

func TestNewVersionID(t *testing.T) {
	id := NewVersionID("bestPracticeValue", "versionValue")

	if id.BestPracticeName != "bestPracticeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BestPracticeName'", id.BestPracticeName, "bestPracticeValue")
	}

	if id.VersionName != "versionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionValue")
	}
}

func TestFormatVersionID(t *testing.T) {
	actual := NewVersionID("bestPracticeValue", "versionValue").ID()
	expected := "/providers/Microsoft.AutoManage/bestPractices/bestPracticeValue/versions/versionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
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
			// Incomplete URI
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeValue/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeValue/versions/versionValue",
			Expected: &VersionId{
				BestPracticeName: "bestPracticeValue",
				VersionName:      "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeValue/versions/versionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionID(v.Input)
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestParseVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
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
			// Incomplete URI
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/bEsTpRaCtIcEs/bEsTpRaCtIcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeValue/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/bEsTpRaCtIcEs/bEsTpRaCtIcEvAlUe/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeValue/versions/versionValue",
			Expected: &VersionId{
				BestPracticeName: "bestPracticeValue",
				VersionName:      "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.AutoManage/bestPractices/bestPracticeValue/versions/versionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/bEsTpRaCtIcEs/bEsTpRaCtIcEvAlUe/vErSiOnS/vErSiOnVaLuE",
			Expected: &VersionId{
				BestPracticeName: "bEsTpRaCtIcEvAlUe",
				VersionName:      "vErSiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/bEsTpRaCtIcEs/bEsTpRaCtIcEvAlUe/vErSiOnS/vErSiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionIDInsensitively(v.Input)
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestSegmentsForVersionId(t *testing.T) {
	segments := VersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
