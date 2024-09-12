package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionPolicyLabelId{}

func TestNewMeInformationProtectionPolicyLabelID(t *testing.T) {
	id := NewMeInformationProtectionPolicyLabelID("informationProtectionLabelIdValue")

	if id.InformationProtectionLabelId != "informationProtectionLabelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InformationProtectionLabelId'", id.InformationProtectionLabelId, "informationProtectionLabelIdValue")
	}
}

func TestFormatMeInformationProtectionPolicyLabelID(t *testing.T) {
	actual := NewMeInformationProtectionPolicyLabelID("informationProtectionLabelIdValue").ID()
	expected := "/me/informationProtection/policy/labels/informationProtectionLabelIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInformationProtectionPolicyLabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionPolicyLabelId
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
			Input: "/me/informationProtection/policy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/policy/labels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/policy/labels/informationProtectionLabelIdValue",
			Expected: &MeInformationProtectionPolicyLabelId{
				InformationProtectionLabelId: "informationProtectionLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/policy/labels/informationProtectionLabelIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionPolicyLabelID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InformationProtectionLabelId != v.Expected.InformationProtectionLabelId {
			t.Fatalf("Expected %q but got %q for InformationProtectionLabelId", v.Expected.InformationProtectionLabelId, actual.InformationProtectionLabelId)
		}

	}
}

func TestParseMeInformationProtectionPolicyLabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionPolicyLabelId
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
			Input: "/me/informationProtection/policy",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/pOlIcY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/policy/labels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/pOlIcY/lAbElS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/policy/labels/informationProtectionLabelIdValue",
			Expected: &MeInformationProtectionPolicyLabelId{
				InformationProtectionLabelId: "informationProtectionLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/policy/labels/informationProtectionLabelIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/pOlIcY/lAbElS/iNfOrMaTiOnPrOtEcTiOnLaBeLiDvAlUe",
			Expected: &MeInformationProtectionPolicyLabelId{
				InformationProtectionLabelId: "iNfOrMaTiOnPrOtEcTiOnLaBeLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/pOlIcY/lAbElS/iNfOrMaTiOnPrOtEcTiOnLaBeLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionPolicyLabelIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InformationProtectionLabelId != v.Expected.InformationProtectionLabelId {
			t.Fatalf("Expected %q but got %q for InformationProtectionLabelId", v.Expected.InformationProtectionLabelId, actual.InformationProtectionLabelId)
		}

	}
}

func TestSegmentsForMeInformationProtectionPolicyLabelId(t *testing.T) {
	segments := MeInformationProtectionPolicyLabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInformationProtectionPolicyLabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
