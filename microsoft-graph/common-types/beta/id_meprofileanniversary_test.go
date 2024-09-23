package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileAnniversaryId{}

func TestNewMeProfileAnniversaryID(t *testing.T) {
	id := NewMeProfileAnniversaryID("personAnnualEventId")

	if id.PersonAnnualEventId != "personAnnualEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonAnnualEventId'", id.PersonAnnualEventId, "personAnnualEventId")
	}
}

func TestFormatMeProfileAnniversaryID(t *testing.T) {
	actual := NewMeProfileAnniversaryID("personAnnualEventId").ID()
	expected := "/me/profile/anniversaries/personAnnualEventId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileAnniversaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileAnniversaryId
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
			Input: "/me/profile/anniversaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/anniversaries/personAnnualEventId",
			Expected: &MeProfileAnniversaryId{
				PersonAnnualEventId: "personAnnualEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/anniversaries/personAnnualEventId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileAnniversaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonAnnualEventId != v.Expected.PersonAnnualEventId {
			t.Fatalf("Expected %q but got %q for PersonAnnualEventId", v.Expected.PersonAnnualEventId, actual.PersonAnnualEventId)
		}

	}
}

func TestParseMeProfileAnniversaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileAnniversaryId
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
			Input: "/me/profile/anniversaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aNnIvErSaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/anniversaries/personAnnualEventId",
			Expected: &MeProfileAnniversaryId{
				PersonAnnualEventId: "personAnnualEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/anniversaries/personAnnualEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aNnIvErSaRiEs/pErSoNaNnUaLeVeNtId",
			Expected: &MeProfileAnniversaryId{
				PersonAnnualEventId: "pErSoNaNnUaLeVeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aNnIvErSaRiEs/pErSoNaNnUaLeVeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileAnniversaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonAnnualEventId != v.Expected.PersonAnnualEventId {
			t.Fatalf("Expected %q but got %q for PersonAnnualEventId", v.Expected.PersonAnnualEventId, actual.PersonAnnualEventId)
		}

	}
}

func TestSegmentsForMeProfileAnniversaryId(t *testing.T) {
	segments := MeProfileAnniversaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileAnniversaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
