package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionDataLossPreventionPolicyId{}

func TestNewUserIdInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	id := NewUserIdInformationProtectionDataLossPreventionPolicyID("userId", "dataLossPreventionPolicyId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DataLossPreventionPolicyId != "dataLossPreventionPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DataLossPreventionPolicyId'", id.DataLossPreventionPolicyId, "dataLossPreventionPolicyId")
	}
}

func TestFormatUserIdInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	actual := NewUserIdInformationProtectionDataLossPreventionPolicyID("userId", "dataLossPreventionPolicyId").ID()
	expected := "/users/userId/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionDataLossPreventionPolicyId
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
			Input: "/users/userId/informationProtection/dataLossPreventionPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyId",
			Expected: &UserIdInformationProtectionDataLossPreventionPolicyId{
				UserId:                     "userId",
				DataLossPreventionPolicyId: "dataLossPreventionPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionDataLossPreventionPolicyID(v.Input)
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

func TestParseUserIdInformationProtectionDataLossPreventionPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionDataLossPreventionPolicyId
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
			Input: "/users/userId/informationProtection/dataLossPreventionPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyId",
			Expected: &UserIdInformationProtectionDataLossPreventionPolicyId{
				UserId:                     "userId",
				DataLossPreventionPolicyId: "dataLossPreventionPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS/dAtAlOsSpReVeNtIoNpOlIcYiD",
			Expected: &UserIdInformationProtectionDataLossPreventionPolicyId{
				UserId:                     "uSeRiD",
				DataLossPreventionPolicyId: "dAtAlOsSpReVeNtIoNpOlIcYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS/dAtAlOsSpReVeNtIoNpOlIcYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionDataLossPreventionPolicyIDInsensitively(v.Input)
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

func TestSegmentsForUserIdInformationProtectionDataLossPreventionPolicyId(t *testing.T) {
	segments := UserIdInformationProtectionDataLossPreventionPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdInformationProtectionDataLossPreventionPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
