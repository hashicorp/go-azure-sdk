package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdSecurityInformationProtectionSensitivityLabelId{}

func TestNewUserIdSecurityInformationProtectionSensitivityLabelID(t *testing.T) {
	id := NewUserIdSecurityInformationProtectionSensitivityLabelID("userId", "sensitivityLabelId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.SensitivityLabelId != "sensitivityLabelId" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelId")
	}
}

func TestFormatUserIdSecurityInformationProtectionSensitivityLabelID(t *testing.T) {
	actual := NewUserIdSecurityInformationProtectionSensitivityLabelID("userId", "sensitivityLabelId").ID()
	expected := "/users/userId/security/informationProtection/sensitivityLabels/sensitivityLabelId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdSecurityInformationProtectionSensitivityLabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdSecurityInformationProtectionSensitivityLabelId
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
			Input: "/users/userId/security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/security/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/security/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/security/informationProtection/sensitivityLabels/sensitivityLabelId",
			Expected: &UserIdSecurityInformationProtectionSensitivityLabelId{
				UserId:             "userId",
				SensitivityLabelId: "sensitivityLabelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/security/informationProtection/sensitivityLabels/sensitivityLabelId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdSecurityInformationProtectionSensitivityLabelID(v.Input)
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

	}
}

func TestParseUserIdSecurityInformationProtectionSensitivityLabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdSecurityInformationProtectionSensitivityLabelId
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
			Input: "/users/userId/security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/security/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/security/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/security/informationProtection/sensitivityLabels/sensitivityLabelId",
			Expected: &UserIdSecurityInformationProtectionSensitivityLabelId{
				UserId:             "userId",
				SensitivityLabelId: "sensitivityLabelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/security/informationProtection/sensitivityLabels/sensitivityLabelId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD",
			Expected: &UserIdSecurityInformationProtectionSensitivityLabelId{
				UserId:             "uSeRiD",
				SensitivityLabelId: "sEnSiTiViTyLaBeLiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdSecurityInformationProtectionSensitivityLabelIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserIdSecurityInformationProtectionSensitivityLabelId(t *testing.T) {
	segments := UserIdSecurityInformationProtectionSensitivityLabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdSecurityInformationProtectionSensitivityLabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
