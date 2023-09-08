package userinformationprotectiondatalosspreventionpolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInformationProtectionDataLossPreventionPolicyId{}

func TestNewUserInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	id := NewUserInformationProtectionDataLossPreventionPolicyID("userIdValue", "dataLossPreventionPolicyIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.DataLossPreventionPolicyId != "dataLossPreventionPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DataLossPreventionPolicyId'", id.DataLossPreventionPolicyId, "dataLossPreventionPolicyIdValue")
	}
}

func TestFormatUserInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	actual := NewUserInformationProtectionDataLossPreventionPolicyID("userIdValue", "dataLossPreventionPolicyIdValue").ID()
	expected := "/users/userIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserInformationProtectionDataLossPreventionPolicyId
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
			Input: "/users/userIdValue/informationProtection/dataLossPreventionPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue",
			Expected: &UserInformationProtectionDataLossPreventionPolicyId{
				UserId:                     "userIdValue",
				DataLossPreventionPolicyId: "dataLossPreventionPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserInformationProtectionDataLossPreventionPolicyID(v.Input)
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

		if actual.DataLossPreventionPolicyId != v.Expected.DataLossPreventionPolicyId {
			t.Fatalf("Expected %q but got %q for DataLossPreventionPolicyId", v.Expected.DataLossPreventionPolicyId, actual.DataLossPreventionPolicyId)
		}

	}
}

func TestParseUserInformationProtectionDataLossPreventionPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserInformationProtectionDataLossPreventionPolicyId
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
			Input: "/users/userIdValue/informationProtection/dataLossPreventionPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue",
			Expected: &UserInformationProtectionDataLossPreventionPolicyId{
				UserId:                     "userIdValue",
				DataLossPreventionPolicyId: "dataLossPreventionPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS/dAtAlOsSpReVeNtIoNpOlIcYiDvAlUe",
			Expected: &UserInformationProtectionDataLossPreventionPolicyId{
				UserId:                     "uSeRiDvAlUe",
				DataLossPreventionPolicyId: "dAtAlOsSpReVeNtIoNpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS/dAtAlOsSpReVeNtIoNpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserInformationProtectionDataLossPreventionPolicyIDInsensitively(v.Input)
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

		if actual.DataLossPreventionPolicyId != v.Expected.DataLossPreventionPolicyId {
			t.Fatalf("Expected %q but got %q for DataLossPreventionPolicyId", v.Expected.DataLossPreventionPolicyId, actual.DataLossPreventionPolicyId)
		}

	}
}

func TestSegmentsForUserInformationProtectionDataLossPreventionPolicyId(t *testing.T) {
	segments := UserInformationProtectionDataLossPreventionPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserInformationProtectionDataLossPreventionPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
