package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyCrossTenantAccessPolicyPartnerId{}

func TestNewPolicyCrossTenantAccessPolicyPartnerID(t *testing.T) {
	id := NewPolicyCrossTenantAccessPolicyPartnerID("crossTenantAccessPolicyConfigurationPartnerTenantId")

	if id.CrossTenantAccessPolicyConfigurationPartnerTenantId != "crossTenantAccessPolicyConfigurationPartnerTenantId" {
		t.Fatalf("Expected %q but got %q for Segment 'CrossTenantAccessPolicyConfigurationPartnerTenantId'", id.CrossTenantAccessPolicyConfigurationPartnerTenantId, "crossTenantAccessPolicyConfigurationPartnerTenantId")
	}
}

func TestFormatPolicyCrossTenantAccessPolicyPartnerID(t *testing.T) {
	actual := NewPolicyCrossTenantAccessPolicyPartnerID("crossTenantAccessPolicyConfigurationPartnerTenantId").ID()
	expected := "/policies/crossTenantAccessPolicy/partners/crossTenantAccessPolicyConfigurationPartnerTenantId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyCrossTenantAccessPolicyPartnerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyCrossTenantAccessPolicyPartnerId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/crossTenantAccessPolicy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/crossTenantAccessPolicy/partners",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/crossTenantAccessPolicy/partners/crossTenantAccessPolicyConfigurationPartnerTenantId",
			Expected: &PolicyCrossTenantAccessPolicyPartnerId{
				CrossTenantAccessPolicyConfigurationPartnerTenantId: "crossTenantAccessPolicyConfigurationPartnerTenantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/crossTenantAccessPolicy/partners/crossTenantAccessPolicyConfigurationPartnerTenantId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyCrossTenantAccessPolicyPartnerID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CrossTenantAccessPolicyConfigurationPartnerTenantId != v.Expected.CrossTenantAccessPolicyConfigurationPartnerTenantId {
			t.Fatalf("Expected %q but got %q for CrossTenantAccessPolicyConfigurationPartnerTenantId", v.Expected.CrossTenantAccessPolicyConfigurationPartnerTenantId, actual.CrossTenantAccessPolicyConfigurationPartnerTenantId)
		}

	}
}

func TestParsePolicyCrossTenantAccessPolicyPartnerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyCrossTenantAccessPolicyPartnerId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/crossTenantAccessPolicy",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cRoSsTeNaNtAcCeSsPoLiCy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/crossTenantAccessPolicy/partners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cRoSsTeNaNtAcCeSsPoLiCy/pArTnErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/crossTenantAccessPolicy/partners/crossTenantAccessPolicyConfigurationPartnerTenantId",
			Expected: &PolicyCrossTenantAccessPolicyPartnerId{
				CrossTenantAccessPolicyConfigurationPartnerTenantId: "crossTenantAccessPolicyConfigurationPartnerTenantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/crossTenantAccessPolicy/partners/crossTenantAccessPolicyConfigurationPartnerTenantId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cRoSsTeNaNtAcCeSsPoLiCy/pArTnErS/cRoSsTeNaNtAcCeSsPoLiCyCoNfIgUrAtIoNpArTnErTeNaNtId",
			Expected: &PolicyCrossTenantAccessPolicyPartnerId{
				CrossTenantAccessPolicyConfigurationPartnerTenantId: "cRoSsTeNaNtAcCeSsPoLiCyCoNfIgUrAtIoNpArTnErTeNaNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cRoSsTeNaNtAcCeSsPoLiCy/pArTnErS/cRoSsTeNaNtAcCeSsPoLiCyCoNfIgUrAtIoNpArTnErTeNaNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyCrossTenantAccessPolicyPartnerIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CrossTenantAccessPolicyConfigurationPartnerTenantId != v.Expected.CrossTenantAccessPolicyConfigurationPartnerTenantId {
			t.Fatalf("Expected %q but got %q for CrossTenantAccessPolicyConfigurationPartnerTenantId", v.Expected.CrossTenantAccessPolicyConfigurationPartnerTenantId, actual.CrossTenantAccessPolicyConfigurationPartnerTenantId)
		}

	}
}

func TestSegmentsForPolicyCrossTenantAccessPolicyPartnerId(t *testing.T) {
	segments := PolicyCrossTenantAccessPolicyPartnerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyCrossTenantAccessPolicyPartnerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
