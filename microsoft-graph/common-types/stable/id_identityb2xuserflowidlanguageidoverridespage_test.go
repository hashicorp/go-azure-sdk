package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdLanguageIdOverridesPageId{}

func TestNewIdentityB2xUserFlowIdLanguageIdOverridesPageID(t *testing.T) {
	id := NewIdentityB2xUserFlowIdLanguageIdOverridesPageID("b2xIdentityUserFlowId", "userFlowLanguageConfigurationId", "userFlowLanguagePageId")

	if id.B2xIdentityUserFlowId != "b2xIdentityUserFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'B2xIdentityUserFlowId'", id.B2xIdentityUserFlowId, "b2xIdentityUserFlowId")
	}

	if id.UserFlowLanguageConfigurationId != "userFlowLanguageConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserFlowLanguageConfigurationId'", id.UserFlowLanguageConfigurationId, "userFlowLanguageConfigurationId")
	}

	if id.UserFlowLanguagePageId != "userFlowLanguagePageId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserFlowLanguagePageId'", id.UserFlowLanguagePageId, "userFlowLanguagePageId")
	}
}

func TestFormatIdentityB2xUserFlowIdLanguageIdOverridesPageID(t *testing.T) {
	actual := NewIdentityB2xUserFlowIdLanguageIdOverridesPageID("b2xIdentityUserFlowId", "userFlowLanguageConfigurationId", "userFlowLanguagePageId").ID()
	expected := "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/overridesPages/userFlowLanguagePageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2xUserFlowIdLanguageIdOverridesPageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdLanguageIdOverridesPageId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/overridesPages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/overridesPages/userFlowLanguagePageId",
			Expected: &IdentityB2xUserFlowIdLanguageIdOverridesPageId{
				B2xIdentityUserFlowId:           "b2xIdentityUserFlowId",
				UserFlowLanguageConfigurationId: "userFlowLanguageConfigurationId",
				UserFlowLanguagePageId:          "userFlowLanguagePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/overridesPages/userFlowLanguagePageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdLanguageIdOverridesPageID(v.Input)
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

		if actual.UserFlowLanguagePageId != v.Expected.UserFlowLanguagePageId {
			t.Fatalf("Expected %q but got %q for UserFlowLanguagePageId", v.Expected.UserFlowLanguagePageId, actual.UserFlowLanguagePageId)
		}

	}
}

func TestParseIdentityB2xUserFlowIdLanguageIdOverridesPageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdLanguageIdOverridesPageId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/lAnGuAgEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/overridesPages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId/oVeRrIdEsPaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/overridesPages/userFlowLanguagePageId",
			Expected: &IdentityB2xUserFlowIdLanguageIdOverridesPageId{
				B2xIdentityUserFlowId:           "b2xIdentityUserFlowId",
				UserFlowLanguageConfigurationId: "userFlowLanguageConfigurationId",
				UserFlowLanguagePageId:          "userFlowLanguagePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/overridesPages/userFlowLanguagePageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId/oVeRrIdEsPaGeS/uSeRfLoWlAnGuAgEpAgEiD",
			Expected: &IdentityB2xUserFlowIdLanguageIdOverridesPageId{
				B2xIdentityUserFlowId:           "b2xIdEnTiTyUsErFlOwId",
				UserFlowLanguageConfigurationId: "uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId",
				UserFlowLanguagePageId:          "uSeRfLoWlAnGuAgEpAgEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId/oVeRrIdEsPaGeS/uSeRfLoWlAnGuAgEpAgEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdLanguageIdOverridesPageIDInsensitively(v.Input)
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

		if actual.UserFlowLanguagePageId != v.Expected.UserFlowLanguagePageId {
			t.Fatalf("Expected %q but got %q for UserFlowLanguagePageId", v.Expected.UserFlowLanguagePageId, actual.UserFlowLanguagePageId)
		}

	}
}

func TestSegmentsForIdentityB2xUserFlowIdLanguageIdOverridesPageId(t *testing.T) {
	segments := IdentityB2xUserFlowIdLanguageIdOverridesPageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2xUserFlowIdLanguageIdOverridesPageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
