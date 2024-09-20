package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionSensitivityLabelIdSublabelId{}

func TestNewMeInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	id := NewMeInformationProtectionSensitivityLabelIdSublabelID("sensitivityLabelId", "sensitivityLabelId1")

	if id.SensitivityLabelId != "sensitivityLabelId" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelId")
	}

	if id.SensitivityLabelId1 != "sensitivityLabelId1" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId1'", id.SensitivityLabelId1, "sensitivityLabelId1")
	}
}

func TestFormatMeInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	actual := NewMeInformationProtectionSensitivityLabelIdSublabelID("sensitivityLabelId", "sensitivityLabelId1").ID()
	expected := "/me/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionSensitivityLabelIdSublabelId
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
			Input: "/me/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &MeInformationProtectionSensitivityLabelIdSublabelId{
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionSensitivityLabelIdSublabelID(v.Input)
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

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestParseMeInformationProtectionSensitivityLabelIdSublabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionSensitivityLabelIdSublabelId
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
			Input: "/me/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &MeInformationProtectionSensitivityLabelIdSublabelId{
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1",
			Expected: &MeInformationProtectionSensitivityLabelIdSublabelId{
				SensitivityLabelId:  "sEnSiTiViTyLaBeLiD",
				SensitivityLabelId1: "sEnSiTiViTyLaBeLiD1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionSensitivityLabelIdSublabelIDInsensitively(v.Input)
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

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestSegmentsForMeInformationProtectionSensitivityLabelIdSublabelId(t *testing.T) {
	segments := MeInformationProtectionSensitivityLabelIdSublabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInformationProtectionSensitivityLabelIdSublabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
