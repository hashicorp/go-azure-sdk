package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdLanguageIdDefaultPageId{}

func TestNewIdentityB2xUserFlowIdLanguageIdDefaultPageID(t *testing.T) {
	id := NewIdentityB2xUserFlowIdLanguageIdDefaultPageID("b2xIdentityUserFlowId", "userFlowLanguageConfigurationId", "userFlowLanguagePageId")

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

func TestFormatIdentityB2xUserFlowIdLanguageIdDefaultPageID(t *testing.T) {
	actual := NewIdentityB2xUserFlowIdLanguageIdDefaultPageID("b2xIdentityUserFlowId", "userFlowLanguageConfigurationId", "userFlowLanguagePageId").ID()
	expected := "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2xUserFlowIdLanguageIdDefaultPageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdLanguageIdDefaultPageId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId",
			Expected: &IdentityB2xUserFlowIdLanguageIdDefaultPageId{
				B2xIdentityUserFlowId:           "b2xIdentityUserFlowId",
				UserFlowLanguageConfigurationId: "userFlowLanguageConfigurationId",
				UserFlowLanguagePageId:          "userFlowLanguagePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdLanguageIdDefaultPageID(v.Input)
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

func TestParseIdentityB2xUserFlowIdLanguageIdDefaultPageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdLanguageIdDefaultPageId
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
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId/dEfAuLtPaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId",
			Expected: &IdentityB2xUserFlowIdLanguageIdDefaultPageId{
				B2xIdentityUserFlowId:           "b2xIdentityUserFlowId",
				UserFlowLanguageConfigurationId: "userFlowLanguageConfigurationId",
				UserFlowLanguagePageId:          "userFlowLanguagePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId/dEfAuLtPaGeS/uSeRfLoWlAnGuAgEpAgEiD",
			Expected: &IdentityB2xUserFlowIdLanguageIdDefaultPageId{
				B2xIdentityUserFlowId:           "b2xIdEnTiTyUsErFlOwId",
				UserFlowLanguageConfigurationId: "uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId",
				UserFlowLanguagePageId:          "uSeRfLoWlAnGuAgEpAgEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId/dEfAuLtPaGeS/uSeRfLoWlAnGuAgEpAgEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdLanguageIdDefaultPageIDInsensitively(v.Input)
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

func TestSegmentsForIdentityB2xUserFlowIdLanguageIdDefaultPageId(t *testing.T) {
	segments := IdentityB2xUserFlowIdLanguageIdDefaultPageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2xUserFlowIdLanguageIdDefaultPageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
