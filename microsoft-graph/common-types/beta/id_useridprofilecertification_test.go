package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileCertificationId{}

func TestNewUserIdProfileCertificationID(t *testing.T) {
	id := NewUserIdProfileCertificationID("userId", "personCertificationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PersonCertificationId != "personCertificationId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonCertificationId'", id.PersonCertificationId, "personCertificationId")
	}
}

func TestFormatUserIdProfileCertificationID(t *testing.T) {
	actual := NewUserIdProfileCertificationID("userId", "personCertificationId").ID()
	expected := "/users/userId/profile/certifications/personCertificationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileCertificationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileCertificationId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile/certifications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/certifications/personCertificationId",
			Expected: &UserIdProfileCertificationId{
				UserId:                "userId",
				PersonCertificationId: "personCertificationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/certifications/personCertificationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileCertificationID(v.Input)
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

func TestParseUserIdProfileCertificationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileCertificationId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile/certifications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/cErTiFiCaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/certifications/personCertificationId",
			Expected: &UserIdProfileCertificationId{
				UserId:                "userId",
				PersonCertificationId: "personCertificationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/certifications/personCertificationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/cErTiFiCaTiOnS/pErSoNcErTiFiCaTiOnId",
			Expected: &UserIdProfileCertificationId{
				UserId:                "uSeRiD",
				PersonCertificationId: "pErSoNcErTiFiCaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/cErTiFiCaTiOnS/pErSoNcErTiFiCaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileCertificationIDInsensitively(v.Input)
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

func TestSegmentsForUserIdProfileCertificationId(t *testing.T) {
	segments := UserIdProfileCertificationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileCertificationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
