package securityinformationprotectionsensitivitylabel

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSecurityInformationProtectionSensitivityLabelId{}

func TestNewMeSecurityInformationProtectionSensitivityLabelID(t *testing.T) {
	id := NewMeSecurityInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

	if id.SensitivityLabelId != "sensitivityLabelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelIdValue")
	}
}

func TestFormatMeSecurityInformationProtectionSensitivityLabelID(t *testing.T) {
	actual := NewMeSecurityInformationProtectionSensitivityLabelID("sensitivityLabelIdValue").ID()
	expected := "/me/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeSecurityInformationProtectionSensitivityLabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeSecurityInformationProtectionSensitivityLabelId
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
			Input: "/me/security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/security/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/security/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue",
			Expected: &MeSecurityInformationProtectionSensitivityLabelId{
				SensitivityLabelId: "sensitivityLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeSecurityInformationProtectionSensitivityLabelID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

	}
}

func TestParseMeSecurityInformationProtectionSensitivityLabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeSecurityInformationProtectionSensitivityLabelId
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
			Input: "/me/security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/security/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/security/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue",
			Expected: &MeSecurityInformationProtectionSensitivityLabelId{
				SensitivityLabelId: "sensitivityLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe",
			Expected: &MeSecurityInformationProtectionSensitivityLabelId{
				SensitivityLabelId: "sEnSiTiViTyLaBeLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeSecurityInformationProtectionSensitivityLabelIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

	}
}

func TestSegmentsForMeSecurityInformationProtectionSensitivityLabelId(t *testing.T) {
	segments := MeSecurityInformationProtectionSensitivityLabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeSecurityInformationProtectionSensitivityLabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
