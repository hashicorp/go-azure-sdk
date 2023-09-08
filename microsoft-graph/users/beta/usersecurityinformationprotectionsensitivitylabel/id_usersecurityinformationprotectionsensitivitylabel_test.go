package usersecurityinformationprotectionsensitivitylabel

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserSecurityInformationProtectionSensitivityLabelId{}

func TestNewUserSecurityInformationProtectionSensitivityLabelID(t *testing.T) {
	id := NewUserSecurityInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.SensitivityLabelId != "sensitivityLabelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelIdValue")
	}
}

func TestFormatUserSecurityInformationProtectionSensitivityLabelID(t *testing.T) {
	actual := NewUserSecurityInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue").ID()
	expected := "/users/userIdValue/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserSecurityInformationProtectionSensitivityLabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserSecurityInformationProtectionSensitivityLabelId
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
			Input: "/users/userIdValue/security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/security/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/security/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue",
			Expected: &UserSecurityInformationProtectionSensitivityLabelId{
				UserId:             "userIdValue",
				SensitivityLabelId: "sensitivityLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserSecurityInformationProtectionSensitivityLabelID(v.Input)
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

func TestParseUserSecurityInformationProtectionSensitivityLabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserSecurityInformationProtectionSensitivityLabelId
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
			Input: "/users/userIdValue/security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/security/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/security/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue",
			Expected: &UserSecurityInformationProtectionSensitivityLabelId{
				UserId:             "userIdValue",
				SensitivityLabelId: "sensitivityLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/security/informationProtection/sensitivityLabels/sensitivityLabelIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe",
			Expected: &UserSecurityInformationProtectionSensitivityLabelId{
				UserId:             "uSeRiDvAlUe",
				SensitivityLabelId: "sEnSiTiViTyLaBeLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEcUrItY/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserSecurityInformationProtectionSensitivityLabelIDInsensitively(v.Input)
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

func TestSegmentsForUserSecurityInformationProtectionSensitivityLabelId(t *testing.T) {
	segments := UserSecurityInformationProtectionSensitivityLabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserSecurityInformationProtectionSensitivityLabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
