package profileeducationalactivity

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileEducationalActivityId{}

func TestNewMeProfileEducationalActivityID(t *testing.T) {
	id := NewMeProfileEducationalActivityID("educationalActivityIdValue")

	if id.EducationalActivityId != "educationalActivityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EducationalActivityId'", id.EducationalActivityId, "educationalActivityIdValue")
	}
}

func TestFormatMeProfileEducationalActivityID(t *testing.T) {
	actual := NewMeProfileEducationalActivityID("educationalActivityIdValue").ID()
	expected := "/me/profile/educationalActivities/educationalActivityIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileEducationalActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileEducationalActivityId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/educationalActivities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/educationalActivities/educationalActivityIdValue",
			Expected: &MeProfileEducationalActivityId{
				EducationalActivityId: "educationalActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/educationalActivities/educationalActivityIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileEducationalActivityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EducationalActivityId != v.Expected.EducationalActivityId {
			t.Fatalf("Expected %q but got %q for EducationalActivityId", v.Expected.EducationalActivityId, actual.EducationalActivityId)
		}

	}
}

func TestParseMeProfileEducationalActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileEducationalActivityId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/educationalActivities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/eDuCaTiOnAlAcTiViTiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/educationalActivities/educationalActivityIdValue",
			Expected: &MeProfileEducationalActivityId{
				EducationalActivityId: "educationalActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/educationalActivities/educationalActivityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/eDuCaTiOnAlAcTiViTiEs/eDuCaTiOnAlAcTiViTyIdVaLuE",
			Expected: &MeProfileEducationalActivityId{
				EducationalActivityId: "eDuCaTiOnAlAcTiViTyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/eDuCaTiOnAlAcTiViTiEs/eDuCaTiOnAlAcTiViTyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileEducationalActivityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EducationalActivityId != v.Expected.EducationalActivityId {
			t.Fatalf("Expected %q but got %q for EducationalActivityId", v.Expected.EducationalActivityId, actual.EducationalActivityId)
		}

	}
}

func TestSegmentsForMeProfileEducationalActivityId(t *testing.T) {
	segments := MeProfileEducationalActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileEducationalActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
