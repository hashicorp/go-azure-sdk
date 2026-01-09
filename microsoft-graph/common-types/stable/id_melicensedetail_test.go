package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeLicenseDetailId{}

func TestNewMeLicenseDetailID(t *testing.T) {
	id := NewMeLicenseDetailID("licenseDetailsId")

	if id.LicenseDetailsId != "licenseDetailsId" {
		t.Fatalf("Expected %q but got %q for Segment 'LicenseDetailsId'", id.LicenseDetailsId, "licenseDetailsId")
	}
}

func TestFormatMeLicenseDetailID(t *testing.T) {
	actual := NewMeLicenseDetailID("licenseDetailsId").ID()
	expected := "/me/licenseDetails/licenseDetailsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeLicenseDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeLicenseDetailId
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
			Input: "/me/licenseDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/licenseDetails/licenseDetailsId",
			Expected: &MeLicenseDetailId{
				LicenseDetailsId: "licenseDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/licenseDetails/licenseDetailsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeLicenseDetailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LicenseDetailsId != v.Expected.LicenseDetailsId {
			t.Fatalf("Expected %q but got %q for LicenseDetailsId", v.Expected.LicenseDetailsId, actual.LicenseDetailsId)
		}

	}
}

func TestParseMeLicenseDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeLicenseDetailId
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
			Input: "/me/licenseDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/lIcEnSeDeTaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/licenseDetails/licenseDetailsId",
			Expected: &MeLicenseDetailId{
				LicenseDetailsId: "licenseDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/licenseDetails/licenseDetailsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/lIcEnSeDeTaIlS/lIcEnSeDeTaIlSiD",
			Expected: &MeLicenseDetailId{
				LicenseDetailsId: "lIcEnSeDeTaIlSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/lIcEnSeDeTaIlS/lIcEnSeDeTaIlSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeLicenseDetailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LicenseDetailsId != v.Expected.LicenseDetailsId {
			t.Fatalf("Expected %q but got %q for LicenseDetailsId", v.Expected.LicenseDetailsId, actual.LicenseDetailsId)
		}

	}
}

func TestSegmentsForMeLicenseDetailId(t *testing.T) {
	segments := MeLicenseDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeLicenseDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
