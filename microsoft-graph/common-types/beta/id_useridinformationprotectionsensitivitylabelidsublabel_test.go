package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionSensitivityLabelIdSublabelId{}

func TestNewUserIdInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	id := NewUserIdInformationProtectionSensitivityLabelIdSublabelID("userId", "sensitivityLabelId", "sensitivityLabelId1")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.SensitivityLabelId != "sensitivityLabelId" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelId")
	}

	if id.SensitivityLabelId1 != "sensitivityLabelId1" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId1'", id.SensitivityLabelId1, "sensitivityLabelId1")
	}
}

func TestFormatUserIdInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	actual := NewUserIdInformationProtectionSensitivityLabelIdSublabelID("userId", "sensitivityLabelId", "sensitivityLabelId1").ID()
	expected := "/users/userId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1"
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &UserIdInformationProtectionSensitivityLabelIdSublabelId{
				UserId:              "userId",
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &UserIdInformationProtectionSensitivityLabelIdSublabelId{
				UserId:              "userId",
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1",
			Expected: &UserIdInformationProtectionSensitivityLabelIdSublabelId{
				UserId:              "uSeRiD",
				SensitivityLabelId:  "sEnSiTiViTyLaBeLiD",
				SensitivityLabelId1: "sEnSiTiViTyLaBeLiD1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1/extra",
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
