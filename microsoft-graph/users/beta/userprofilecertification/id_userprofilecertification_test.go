package userprofilecertification

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileCertificationId{}

func TestNewUserProfileCertificationID(t *testing.T) {
	id := NewUserProfileCertificationID("userIdValue", "personCertificationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PersonCertificationId != "personCertificationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonCertificationId'", id.PersonCertificationId, "personCertificationIdValue")
	}
}

func TestFormatUserProfileCertificationID(t *testing.T) {
	actual := NewUserProfileCertificationID("userIdValue", "personCertificationIdValue").ID()
	expected := "/users/userIdValue/profile/certifications/personCertificationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileCertificationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileCertificationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile/certifications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/certifications/personCertificationIdValue",
			Expected: &UserProfileCertificationId{
				UserId:                "userIdValue",
				PersonCertificationId: "personCertificationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/certifications/personCertificationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileCertificationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.PersonCertificationId != v.Expected.PersonCertificationId {
			t.Fatalf("Expected %q but got %q for PersonCertificationId", v.Expected.PersonCertificationId, actual.PersonCertificationId)
		}

	}
}

func TestParseUserProfileCertificationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileCertificationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile/certifications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/cErTiFiCaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/certifications/personCertificationIdValue",
			Expected: &UserProfileCertificationId{
				UserId:                "userIdValue",
				PersonCertificationId: "personCertificationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/certifications/personCertificationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/cErTiFiCaTiOnS/pErSoNcErTiFiCaTiOnIdVaLuE",
			Expected: &UserProfileCertificationId{
				UserId:                "uSeRiDvAlUe",
				PersonCertificationId: "pErSoNcErTiFiCaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/cErTiFiCaTiOnS/pErSoNcErTiFiCaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileCertificationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.PersonCertificationId != v.Expected.PersonCertificationId {
			t.Fatalf("Expected %q but got %q for PersonCertificationId", v.Expected.PersonCertificationId, actual.PersonCertificationId)
		}

	}
}

func TestSegmentsForUserProfileCertificationId(t *testing.T) {
	segments := UserProfileCertificationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileCertificationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
