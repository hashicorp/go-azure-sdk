package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdLicenseDetailId{}

func TestNewUserIdLicenseDetailID(t *testing.T) {
	id := NewUserIdLicenseDetailID("userId", "licenseDetailsId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.LicenseDetailsId != "licenseDetailsId" {
		t.Fatalf("Expected %q but got %q for Segment 'LicenseDetailsId'", id.LicenseDetailsId, "licenseDetailsId")
	}
}

func TestFormatUserIdLicenseDetailID(t *testing.T) {
	actual := NewUserIdLicenseDetailID("userId", "licenseDetailsId").ID()
	expected := "/users/userId/licenseDetails/licenseDetailsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdLicenseDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdLicenseDetailId
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
			Input: "/users/userId/licenseDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/licenseDetails/licenseDetailsId",
			Expected: &UserIdLicenseDetailId{
				UserId:           "userId",
				LicenseDetailsId: "licenseDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/licenseDetails/licenseDetailsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdLicenseDetailID(v.Input)
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

		if actual.LicenseDetailsId != v.Expected.LicenseDetailsId {
			t.Fatalf("Expected %q but got %q for LicenseDetailsId", v.Expected.LicenseDetailsId, actual.LicenseDetailsId)
		}

	}
}

func TestParseUserIdLicenseDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdLicenseDetailId
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
			Input: "/users/userId/licenseDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/lIcEnSeDeTaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/licenseDetails/licenseDetailsId",
			Expected: &UserIdLicenseDetailId{
				UserId:           "userId",
				LicenseDetailsId: "licenseDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/licenseDetails/licenseDetailsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/lIcEnSeDeTaIlS/lIcEnSeDeTaIlSiD",
			Expected: &UserIdLicenseDetailId{
				UserId:           "uSeRiD",
				LicenseDetailsId: "lIcEnSeDeTaIlSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/lIcEnSeDeTaIlS/lIcEnSeDeTaIlSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdLicenseDetailIDInsensitively(v.Input)
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

		if actual.LicenseDetailsId != v.Expected.LicenseDetailsId {
			t.Fatalf("Expected %q but got %q for LicenseDetailsId", v.Expected.LicenseDetailsId, actual.LicenseDetailsId)
		}

	}
}

func TestSegmentsForUserIdLicenseDetailId(t *testing.T) {
	segments := UserIdLicenseDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdLicenseDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
