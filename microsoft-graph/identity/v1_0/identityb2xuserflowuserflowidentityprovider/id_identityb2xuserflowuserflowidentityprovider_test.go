package identityb2xuserflowuserflowidentityprovider

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityB2xUserFlowUserFlowIdentityProviderId{}

func TestNewIdentityB2xUserFlowUserFlowIdentityProviderID(t *testing.T) {
	id := NewIdentityB2xUserFlowUserFlowIdentityProviderID("b2xIdentityUserFlowIdValue", "identityProviderBaseIdValue")

	if id.B2xIdentityUserFlowId != "b2xIdentityUserFlowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'B2xIdentityUserFlowId'", id.B2xIdentityUserFlowId, "b2xIdentityUserFlowIdValue")
	}

	if id.IdentityProviderBaseId != "identityProviderBaseIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityProviderBaseId'", id.IdentityProviderBaseId, "identityProviderBaseIdValue")
	}
}

func TestFormatIdentityB2xUserFlowUserFlowIdentityProviderID(t *testing.T) {
	actual := NewIdentityB2xUserFlowUserFlowIdentityProviderID("b2xIdentityUserFlowIdValue", "identityProviderBaseIdValue").ID()
	expected := "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/userFlowIdentityProviders/identityProviderBaseIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2xUserFlowUserFlowIdentityProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowUserFlowIdentityProviderId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/userFlowIdentityProviders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/userFlowIdentityProviders/identityProviderBaseIdValue",
			Expected: &IdentityB2xUserFlowUserFlowIdentityProviderId{
				B2xIdentityUserFlowId:  "b2xIdentityUserFlowIdValue",
				IdentityProviderBaseId: "identityProviderBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/userFlowIdentityProviders/identityProviderBaseIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowUserFlowIdentityProviderID(v.Input)
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

func TestParseIdentityB2xUserFlowUserFlowIdentityProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowUserFlowIdentityProviderId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/userFlowIdentityProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/uSeRfLoWiDeNtItYpRoViDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/userFlowIdentityProviders/identityProviderBaseIdValue",
			Expected: &IdentityB2xUserFlowUserFlowIdentityProviderId{
				B2xIdentityUserFlowId:  "b2xIdentityUserFlowIdValue",
				IdentityProviderBaseId: "identityProviderBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/userFlowIdentityProviders/identityProviderBaseIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/uSeRfLoWiDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiDvAlUe",
			Expected: &IdentityB2xUserFlowUserFlowIdentityProviderId{
				B2xIdentityUserFlowId:  "b2xIdEnTiTyUsErFlOwIdVaLuE",
				IdentityProviderBaseId: "iDeNtItYpRoViDeRbAsEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/uSeRfLoWiDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowUserFlowIdentityProviderIDInsensitively(v.Input)
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

func TestSegmentsForIdentityB2xUserFlowUserFlowIdentityProviderId(t *testing.T) {
	segments := IdentityB2xUserFlowUserFlowIdentityProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2xUserFlowUserFlowIdentityProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
