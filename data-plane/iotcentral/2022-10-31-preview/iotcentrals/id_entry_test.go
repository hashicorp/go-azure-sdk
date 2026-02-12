package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &EntryId{}

func TestNewEntryID(t *testing.T) {
	id := NewEntryID("https://endpoint-url.example.com", "enrollmentGroupId", "primary")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.EnrollmentGroupId != "enrollmentGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentGroupId'", id.EnrollmentGroupId, "enrollmentGroupId")
	}

	if id.Entry != "primary" {
		t.Fatalf("Expected %q but got %q for Segment 'Entry'", id.Entry, "primary")
	}
}

func TestFormatEntryID(t *testing.T) {
	actual := NewEntryID("https://endpoint-url.example.com", "enrollmentGroupId", "primary").ID()
	expected := "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId/certificates/primary"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseEntryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EntryId
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
			// Incomplete URI
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId/certificates",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId/certificates/primary",
			Expected: &EntryId{
				BaseURI:           "https://endpoint-url.example.com",
				EnrollmentGroupId: "enrollmentGroupId",
				Entry:             "primary",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId/certificates/primary/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEntryID(v.Input)
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

		if actual.Entry != v.Expected.Entry {
			t.Fatalf("Expected %q but got %q for Entry", v.Expected.Entry, actual.Entry)
		}

	}
}

func TestParseEntryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EntryId
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
			// Incomplete URI
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/eNrOlLmEnTgRoUpS/eNrOlLmEnTgRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId/certificates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/eNrOlLmEnTgRoUpS/eNrOlLmEnTgRoUpId/cErTiFiCaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId/certificates/primary",
			Expected: &EntryId{
				BaseURI:           "https://endpoint-url.example.com",
				EnrollmentGroupId: "enrollmentGroupId",
				Entry:             "primary",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/enrollmentGroups/enrollmentGroupId/certificates/primary/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/eNrOlLmEnTgRoUpS/eNrOlLmEnTgRoUpId/cErTiFiCaTeS/pRiMaRy",
			Expected: &EntryId{
				BaseURI:           "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				EnrollmentGroupId: "eNrOlLmEnTgRoUpId",
				Entry:             "primary",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/eNrOlLmEnTgRoUpS/eNrOlLmEnTgRoUpId/cErTiFiCaTeS/pRiMaRy/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEntryIDInsensitively(v.Input)
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

		if actual.Entry != v.Expected.Entry {
			t.Fatalf("Expected %q but got %q for Entry", v.Expected.Entry, actual.Entry)
		}

	}
}

func TestSegmentsForEntryId(t *testing.T) {
	segments := EntryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("EntryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
