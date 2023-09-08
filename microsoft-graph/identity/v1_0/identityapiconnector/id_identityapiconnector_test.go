package identityapiconnector

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityApiConnectorId{}

func TestNewIdentityApiConnectorID(t *testing.T) {
	id := NewIdentityApiConnectorID("identityApiConnectorIdValue")

	if id.IdentityApiConnectorId != "identityApiConnectorIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityApiConnectorId'", id.IdentityApiConnectorId, "identityApiConnectorIdValue")
	}
}

func TestFormatIdentityApiConnectorID(t *testing.T) {
	actual := NewIdentityApiConnectorID("identityApiConnectorIdValue").ID()
	expected := "/identity/apiConnectors/identityApiConnectorIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityApiConnectorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityApiConnectorId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/apiConnectors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/apiConnectors/identityApiConnectorIdValue",
			Expected: &IdentityApiConnectorId{
				IdentityApiConnectorId: "identityApiConnectorIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/apiConnectors/identityApiConnectorIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityApiConnectorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityApiConnectorId != v.Expected.IdentityApiConnectorId {
			t.Fatalf("Expected %q but got %q for IdentityApiConnectorId", v.Expected.IdentityApiConnectorId, actual.IdentityApiConnectorId)
		}

	}
}

func TestParseIdentityApiConnectorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityApiConnectorId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/apiConnectors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aPiCoNnEcToRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/apiConnectors/identityApiConnectorIdValue",
			Expected: &IdentityApiConnectorId{
				IdentityApiConnectorId: "identityApiConnectorIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/apiConnectors/identityApiConnectorIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aPiCoNnEcToRs/iDeNtItYaPiCoNnEcToRiDvAlUe",
			Expected: &IdentityApiConnectorId{
				IdentityApiConnectorId: "iDeNtItYaPiCoNnEcToRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aPiCoNnEcToRs/iDeNtItYaPiCoNnEcToRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityApiConnectorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityApiConnectorId != v.Expected.IdentityApiConnectorId {
			t.Fatalf("Expected %q but got %q for IdentityApiConnectorId", v.Expected.IdentityApiConnectorId, actual.IdentityApiConnectorId)
		}

	}
}

func TestSegmentsForIdentityApiConnectorId(t *testing.T) {
	segments := IdentityApiConnectorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityApiConnectorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
