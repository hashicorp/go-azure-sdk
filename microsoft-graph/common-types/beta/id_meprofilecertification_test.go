package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileCertificationId{}

func TestNewMeProfileCertificationID(t *testing.T) {
	id := NewMeProfileCertificationID("personCertificationId")

	if id.PersonCertificationId != "personCertificationId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonCertificationId'", id.PersonCertificationId, "personCertificationId")
	}
}

func TestFormatMeProfileCertificationID(t *testing.T) {
	actual := NewMeProfileCertificationID("personCertificationId").ID()
	expected := "/me/profile/certifications/personCertificationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileCertificationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileCertificationId
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
			Input: "/me/profile/certifications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/certifications/personCertificationId",
			Expected: &MeProfileCertificationId{
				PersonCertificationId: "personCertificationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/certifications/personCertificationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileCertificationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonCertificationId != v.Expected.PersonCertificationId {
			t.Fatalf("Expected %q but got %q for PersonCertificationId", v.Expected.PersonCertificationId, actual.PersonCertificationId)
		}

	}
}

func TestParseMeProfileCertificationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileCertificationId
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
			Input: "/me/profile/certifications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/cErTiFiCaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/certifications/personCertificationId",
			Expected: &MeProfileCertificationId{
				PersonCertificationId: "personCertificationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/certifications/personCertificationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/cErTiFiCaTiOnS/pErSoNcErTiFiCaTiOnId",
			Expected: &MeProfileCertificationId{
				PersonCertificationId: "pErSoNcErTiFiCaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/cErTiFiCaTiOnS/pErSoNcErTiFiCaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileCertificationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonCertificationId != v.Expected.PersonCertificationId {
			t.Fatalf("Expected %q but got %q for PersonCertificationId", v.Expected.PersonCertificationId, actual.PersonCertificationId)
		}

	}
}

func TestSegmentsForMeProfileCertificationId(t *testing.T) {
	segments := MeProfileCertificationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileCertificationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
