package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePhotoId{}

func TestNewMePhotoID(t *testing.T) {
	id := NewMePhotoID("profilePhotoId")

	if id.ProfilePhotoId != "profilePhotoId" {
		t.Fatalf("Expected %q but got %q for Segment 'ProfilePhotoId'", id.ProfilePhotoId, "profilePhotoId")
	}
}

func TestFormatMePhotoID(t *testing.T) {
	actual := NewMePhotoID("profilePhotoId").ID()
	expected := "/me/photos/profilePhotoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePhotoID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePhotoId
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
			Input: "/me/photos",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/photos/profilePhotoId",
			Expected: &MePhotoId{
				ProfilePhotoId: "profilePhotoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/photos/profilePhotoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePhotoID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProfilePhotoId != v.Expected.ProfilePhotoId {
			t.Fatalf("Expected %q but got %q for ProfilePhotoId", v.Expected.ProfilePhotoId, actual.ProfilePhotoId)
		}

	}
}

func TestParseMePhotoIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePhotoId
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
			Input: "/me/photos",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pHoToS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/photos/profilePhotoId",
			Expected: &MePhotoId{
				ProfilePhotoId: "profilePhotoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/photos/profilePhotoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pHoToS/pRoFiLePhOtOiD",
			Expected: &MePhotoId{
				ProfilePhotoId: "pRoFiLePhOtOiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pHoToS/pRoFiLePhOtOiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePhotoIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProfilePhotoId != v.Expected.ProfilePhotoId {
			t.Fatalf("Expected %q but got %q for ProfilePhotoId", v.Expected.ProfilePhotoId, actual.ProfilePhotoId)
		}

	}
}

func TestSegmentsForMePhotoId(t *testing.T) {
	segments := MePhotoId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePhotoId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
