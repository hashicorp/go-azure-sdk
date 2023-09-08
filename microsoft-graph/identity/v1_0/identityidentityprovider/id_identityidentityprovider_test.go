package identityidentityprovider

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityIdentityProviderId{}

func TestNewIdentityIdentityProviderID(t *testing.T) {
	id := NewIdentityIdentityProviderID("identityProviderBaseIdValue")

	if id.IdentityProviderBaseId != "identityProviderBaseIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityProviderBaseId'", id.IdentityProviderBaseId, "identityProviderBaseIdValue")
	}
}

func TestFormatIdentityIdentityProviderID(t *testing.T) {
	actual := NewIdentityIdentityProviderID("identityProviderBaseIdValue").ID()
	expected := "/identity/identityProviders/identityProviderBaseIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityIdentityProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityIdentityProviderId
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
			Input: "/identity/identityProviders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/identityProviders/identityProviderBaseIdValue",
			Expected: &IdentityIdentityProviderId{
				IdentityProviderBaseId: "identityProviderBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/identityProviders/identityProviderBaseIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityIdentityProviderID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestParseIdentityIdentityProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityIdentityProviderId
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
			Input: "/identity/identityProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/iDeNtItYpRoViDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/identityProviders/identityProviderBaseIdValue",
			Expected: &IdentityIdentityProviderId{
				IdentityProviderBaseId: "identityProviderBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/identityProviders/identityProviderBaseIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/iDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiDvAlUe",
			Expected: &IdentityIdentityProviderId{
				IdentityProviderBaseId: "iDeNtItYpRoViDeRbAsEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/iDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityIdentityProviderIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestSegmentsForIdentityIdentityProviderId(t *testing.T) {
	segments := IdentityIdentityProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityIdentityProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
