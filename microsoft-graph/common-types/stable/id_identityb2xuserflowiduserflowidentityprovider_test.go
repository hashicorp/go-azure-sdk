package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdUserFlowIdentityProviderId{}

func TestNewIdentityB2xUserFlowIdUserFlowIdentityProviderID(t *testing.T) {
	id := NewIdentityB2xUserFlowIdUserFlowIdentityProviderID("b2xIdentityUserFlowId", "identityProviderBaseId")

	if id.B2xIdentityUserFlowId != "b2xIdentityUserFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'B2xIdentityUserFlowId'", id.B2xIdentityUserFlowId, "b2xIdentityUserFlowId")
	}

	if id.IdentityProviderBaseId != "identityProviderBaseId" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityProviderBaseId'", id.IdentityProviderBaseId, "identityProviderBaseId")
	}
}

func TestFormatIdentityB2xUserFlowIdUserFlowIdentityProviderID(t *testing.T) {
	actual := NewIdentityB2xUserFlowIdUserFlowIdentityProviderID("b2xIdentityUserFlowId", "identityProviderBaseId").ID()
	expected := "/identity/b2xUserFlows/b2xIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2xUserFlowIdUserFlowIdentityProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdUserFlowIdentityProviderId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userFlowIdentityProviders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId",
			Expected: &IdentityB2xUserFlowIdUserFlowIdentityProviderId{
				B2xIdentityUserFlowId:  "b2xIdentityUserFlowId",
				IdentityProviderBaseId: "identityProviderBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdUserFlowIdentityProviderID(v.Input)
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

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestParseIdentityB2xUserFlowIdUserFlowIdentityProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdUserFlowIdentityProviderId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userFlowIdentityProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/uSeRfLoWiDeNtItYpRoViDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId",
			Expected: &IdentityB2xUserFlowIdUserFlowIdentityProviderId{
				B2xIdentityUserFlowId:  "b2xIdentityUserFlowId",
				IdentityProviderBaseId: "identityProviderBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/uSeRfLoWiDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiD",
			Expected: &IdentityB2xUserFlowIdUserFlowIdentityProviderId{
				B2xIdentityUserFlowId:  "b2xIdEnTiTyUsErFlOwId",
				IdentityProviderBaseId: "iDeNtItYpRoViDeRbAsEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/uSeRfLoWiDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdUserFlowIdentityProviderIDInsensitively(v.Input)
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

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestSegmentsForIdentityB2xUserFlowIdUserFlowIdentityProviderId(t *testing.T) {
	segments := IdentityB2xUserFlowIdUserFlowIdentityProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2xUserFlowIdUserFlowIdentityProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
