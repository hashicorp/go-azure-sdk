package informationprotectionsensitivitylabelsublabel

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionSensitivityLabelIdSublabelId{}

func TestNewUserIdInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	id := NewUserIdInformationProtectionSensitivityLabelIdSublabelID("userIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.SensitivityLabelId != "sensitivityLabelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelIdValue")
	}

	if id.SensitivityLabelId1 != "sensitivityLabelId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId1'", id.SensitivityLabelId1, "sensitivityLabelId1Value")
	}
}

func TestFormatUserIdInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	actual := NewUserIdInformationProtectionSensitivityLabelIdSublabelID("userIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value").ID()
	expected := "/users/userIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue/sublabels/sensitivityLabelId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionSensitivityLabelIdSublabelId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue/sublabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue/sublabels/sensitivityLabelId1Value",
			Expected: &UserIdInformationProtectionSensitivityLabelIdSublabelId{
				UserId:              "userIdValue",
				SensitivityLabelId:  "sensitivityLabelIdValue",
				SensitivityLabelId1: "sensitivityLabelId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue/sublabels/sensitivityLabelId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionSensitivityLabelIdSublabelID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestParseUserIdInformationProtectionSensitivityLabelIdSublabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionSensitivityLabelIdSublabelId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue/sublabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe/sUbLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue/sublabels/sensitivityLabelId1Value",
			Expected: &UserIdInformationProtectionSensitivityLabelIdSublabelId{
				UserId:              "userIdValue",
				SensitivityLabelId:  "sensitivityLabelIdValue",
				SensitivityLabelId1: "sensitivityLabelId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue/sublabels/sensitivityLabelId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe/sUbLaBeLs/sEnSiTiViTyLaBeLiD1VaLuE",
			Expected: &UserIdInformationProtectionSensitivityLabelIdSublabelId{
				UserId:              "uSeRiDvAlUe",
				SensitivityLabelId:  "sEnSiTiViTyLaBeLiDvAlUe",
				SensitivityLabelId1: "sEnSiTiViTyLaBeLiD1VaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe/sUbLaBeLs/sEnSiTiViTyLaBeLiD1VaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionSensitivityLabelIdSublabelIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestSegmentsForUserIdInformationProtectionSensitivityLabelIdSublabelId(t *testing.T) {
	segments := UserIdInformationProtectionSensitivityLabelIdSublabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdInformationProtectionSensitivityLabelIdSublabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
