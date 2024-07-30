package b2xuserflowlanguagedefaultpage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdLanguageId{}

func TestNewIdentityB2xUserFlowIdLanguageID(t *testing.T) {
	id := NewIdentityB2xUserFlowIdLanguageID("b2xIdentityUserFlowIdValue", "userFlowLanguageConfigurationIdValue")

	if id.B2xIdentityUserFlowId != "b2xIdentityUserFlowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'B2xIdentityUserFlowId'", id.B2xIdentityUserFlowId, "b2xIdentityUserFlowIdValue")
	}

	if id.UserFlowLanguageConfigurationId != "userFlowLanguageConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserFlowLanguageConfigurationId'", id.UserFlowLanguageConfigurationId, "userFlowLanguageConfigurationIdValue")
	}
}

func TestFormatIdentityB2xUserFlowIdLanguageID(t *testing.T) {
	actual := NewIdentityB2xUserFlowIdLanguageID("b2xIdentityUserFlowIdValue", "userFlowLanguageConfigurationIdValue").ID()
	expected := "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/languages/userFlowLanguageConfigurationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2xUserFlowIdLanguageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdLanguageId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/languages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/languages/userFlowLanguageConfigurationIdValue",
			Expected: &IdentityB2xUserFlowIdLanguageId{
				B2xIdentityUserFlowId:           "b2xIdentityUserFlowIdValue",
				UserFlowLanguageConfigurationId: "userFlowLanguageConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/languages/userFlowLanguageConfigurationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdLanguageID(v.Input)
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

		if actual.UserFlowLanguageConfigurationId != v.Expected.UserFlowLanguageConfigurationId {
			t.Fatalf("Expected %q but got %q for UserFlowLanguageConfigurationId", v.Expected.UserFlowLanguageConfigurationId, actual.UserFlowLanguageConfigurationId)
		}

	}
}

func TestParseIdentityB2xUserFlowIdLanguageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdLanguageId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/languages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/lAnGuAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/languages/userFlowLanguageConfigurationIdValue",
			Expected: &IdentityB2xUserFlowIdLanguageId{
				B2xIdentityUserFlowId:           "b2xIdentityUserFlowIdValue",
				UserFlowLanguageConfigurationId: "userFlowLanguageConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/languages/userFlowLanguageConfigurationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnIdVaLuE",
			Expected: &IdentityB2xUserFlowIdLanguageId{
				B2xIdentityUserFlowId:           "b2xIdEnTiTyUsErFlOwIdVaLuE",
				UserFlowLanguageConfigurationId: "uSeRfLoWlAnGuAgEcOnFiGuRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdLanguageIDInsensitively(v.Input)
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

		if actual.UserFlowLanguageConfigurationId != v.Expected.UserFlowLanguageConfigurationId {
			t.Fatalf("Expected %q but got %q for UserFlowLanguageConfigurationId", v.Expected.UserFlowLanguageConfigurationId, actual.UserFlowLanguageConfigurationId)
		}

	}
}

func TestSegmentsForIdentityB2xUserFlowIdLanguageId(t *testing.T) {
	segments := IdentityB2xUserFlowIdLanguageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2xUserFlowIdLanguageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
