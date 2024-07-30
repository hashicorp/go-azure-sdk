package informationprotectiondatalosspreventionpolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionDataLossPreventionPolicyId{}

func TestNewMeInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	id := NewMeInformationProtectionDataLossPreventionPolicyID("dataLossPreventionPolicyIdValue")

	if id.DataLossPreventionPolicyId != "dataLossPreventionPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DataLossPreventionPolicyId'", id.DataLossPreventionPolicyId, "dataLossPreventionPolicyIdValue")
	}
}

func TestFormatMeInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	actual := NewMeInformationProtectionDataLossPreventionPolicyID("dataLossPreventionPolicyIdValue").ID()
	expected := "/me/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionDataLossPreventionPolicyId
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
			Input: "/me/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/dataLossPreventionPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue",
			Expected: &MeInformationProtectionDataLossPreventionPolicyId{
				DataLossPreventionPolicyId: "dataLossPreventionPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionDataLossPreventionPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DataLossPreventionPolicyId != v.Expected.DataLossPreventionPolicyId {
			t.Fatalf("Expected %q but got %q for DataLossPreventionPolicyId", v.Expected.DataLossPreventionPolicyId, actual.DataLossPreventionPolicyId)
		}

	}
}

func TestParseMeInformationProtectionDataLossPreventionPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionDataLossPreventionPolicyId
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
			Input: "/me/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/dataLossPreventionPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue",
			Expected: &MeInformationProtectionDataLossPreventionPolicyId{
				DataLossPreventionPolicyId: "dataLossPreventionPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS/dAtAlOsSpReVeNtIoNpOlIcYiDvAlUe",
			Expected: &MeInformationProtectionDataLossPreventionPolicyId{
				DataLossPreventionPolicyId: "dAtAlOsSpReVeNtIoNpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS/dAtAlOsSpReVeNtIoNpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionDataLossPreventionPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DataLossPreventionPolicyId != v.Expected.DataLossPreventionPolicyId {
			t.Fatalf("Expected %q but got %q for DataLossPreventionPolicyId", v.Expected.DataLossPreventionPolicyId, actual.DataLossPreventionPolicyId)
		}

	}
}

func TestSegmentsForMeInformationProtectionDataLossPreventionPolicyId(t *testing.T) {
	segments := MeInformationProtectionDataLossPreventionPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInformationProtectionDataLossPreventionPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
