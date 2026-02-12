package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &EnrollmentGroupId{}

func TestNewEnrollmentGroupID(t *testing.T) {
	id := NewEnrollmentGroupID("https://endpoint-url.example.com", "enrollmentGroupId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.EnrollmentGroupId != "enrollmentGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentGroupId'", id.EnrollmentGroupId, "enrollmentGroupId")
	}
}

func TestFormatEnrollmentGroupID(t *testing.T) {
	actual := NewEnrollmentGroupID("https://endpoint-url.example.com", "enrollmentGroupId").ID()
	expected := "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseEnrollmentGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnrollmentGroupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/enrollmentGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId",
			Expected: &EnrollmentGroupId{
				BaseURI:           "https://endpoint-url.example.com",
				EnrollmentGroupId: "enrollmentGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnrollmentGroupID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.EnrollmentGroupId != v.Expected.EnrollmentGroupId {
			t.Fatalf("Expected %q but got %q for EnrollmentGroupId", v.Expected.EnrollmentGroupId, actual.EnrollmentGroupId)
		}

	}
}

func TestParseEnrollmentGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnrollmentGroupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/enrollmentGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/eNrOlLmEnTgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId",
			Expected: &EnrollmentGroupId{
				BaseURI:           "https://endpoint-url.example.com",
				EnrollmentGroupId: "enrollmentGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/eNrOlLmEnTgRoUpS/eNrOlLmEnTgRoUpId",
			Expected: &EnrollmentGroupId{
				BaseURI:           "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				EnrollmentGroupId: "eNrOlLmEnTgRoUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/eNrOlLmEnTgRoUpS/eNrOlLmEnTgRoUpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnrollmentGroupIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.EnrollmentGroupId != v.Expected.EnrollmentGroupId {
			t.Fatalf("Expected %q but got %q for EnrollmentGroupId", v.Expected.EnrollmentGroupId, actual.EnrollmentGroupId)
		}

	}
}

func TestSegmentsForEnrollmentGroupId(t *testing.T) {
	segments := EnrollmentGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("EnrollmentGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
