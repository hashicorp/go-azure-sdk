package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdIdentityProviderId{}

func TestNewIdentityB2xUserFlowIdIdentityProviderID(t *testing.T) {
	id := NewIdentityB2xUserFlowIdIdentityProviderID("b2xIdentityUserFlowIdValue", "identityProviderIdValue")

	if id.B2xIdentityUserFlowId != "b2xIdentityUserFlowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'B2xIdentityUserFlowId'", id.B2xIdentityUserFlowId, "b2xIdentityUserFlowIdValue")
	}

	if id.IdentityProviderId != "identityProviderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityProviderId'", id.IdentityProviderId, "identityProviderIdValue")
	}
}

func TestFormatIdentityB2xUserFlowIdIdentityProviderID(t *testing.T) {
	actual := NewIdentityB2xUserFlowIdIdentityProviderID("b2xIdentityUserFlowIdValue", "identityProviderIdValue").ID()
	expected := "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/identityProviders/identityProviderIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2xUserFlowIdIdentityProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdIdentityProviderId
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
			Input: "/identity/b2xUserFlows",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/identityProviders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/identityProviders/identityProviderIdValue",
			Expected: &IdentityB2xUserFlowIdIdentityProviderId{
				B2xIdentityUserFlowId: "b2xIdentityUserFlowIdValue",
				IdentityProviderId:    "identityProviderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/identityProviders/identityProviderIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdIdentityProviderID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.B2xIdentityUserFlowId != v.Expected.B2xIdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for B2xIdentityUserFlowId", v.Expected.B2xIdentityUserFlowId, actual.B2xIdentityUserFlowId)
		}

		if actual.IdentityProviderId != v.Expected.IdentityProviderId {
			t.Fatalf("Expected %q but got %q for IdentityProviderId", v.Expected.IdentityProviderId, actual.IdentityProviderId)
		}

	}
}

func TestParseIdentityB2xUserFlowIdIdentityProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdIdentityProviderId
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
			Input: "/identity/b2xUserFlows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/identityProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/iDeNtItYpRoViDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/identityProviders/identityProviderIdValue",
			Expected: &IdentityB2xUserFlowIdIdentityProviderId{
				B2xIdentityUserFlowId: "b2xIdentityUserFlowIdValue",
				IdentityProviderId:    "identityProviderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/identityProviders/identityProviderIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/iDeNtItYpRoViDeRs/iDeNtItYpRoViDeRiDvAlUe",
			Expected: &IdentityB2xUserFlowIdIdentityProviderId{
				B2xIdentityUserFlowId: "b2xIdEnTiTyUsErFlOwIdVaLuE",
				IdentityProviderId:    "iDeNtItYpRoViDeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/iDeNtItYpRoViDeRs/iDeNtItYpRoViDeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdIdentityProviderIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.B2xIdentityUserFlowId != v.Expected.B2xIdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for B2xIdentityUserFlowId", v.Expected.B2xIdentityUserFlowId, actual.B2xIdentityUserFlowId)
		}

		if actual.IdentityProviderId != v.Expected.IdentityProviderId {
			t.Fatalf("Expected %q but got %q for IdentityProviderId", v.Expected.IdentityProviderId, actual.IdentityProviderId)
		}

	}
}

func TestSegmentsForIdentityB2xUserFlowIdIdentityProviderId(t *testing.T) {
	segments := IdentityB2xUserFlowIdIdentityProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2xUserFlowIdIdentityProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
