package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionPolicyLabelId{}

func TestNewUserIdInformationProtectionPolicyLabelID(t *testing.T) {
	id := NewUserIdInformationProtectionPolicyLabelID("userId", "informationProtectionLabelId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.InformationProtectionLabelId != "informationProtectionLabelId" {
		t.Fatalf("Expected %q but got %q for Segment 'InformationProtectionLabelId'", id.InformationProtectionLabelId, "informationProtectionLabelId")
	}
}

func TestFormatUserIdInformationProtectionPolicyLabelID(t *testing.T) {
	actual := NewUserIdInformationProtectionPolicyLabelID("userId", "informationProtectionLabelId").ID()
	expected := "/users/userId/informationProtection/policy/labels/informationProtectionLabelId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdInformationProtectionPolicyLabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionPolicyLabelId
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
			Input: "/users/userId/informationProtection/policy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/policy/labels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/informationProtection/policy/labels/informationProtectionLabelId",
			Expected: &UserIdInformationProtectionPolicyLabelId{
				UserId:                       "userId",
				InformationProtectionLabelId: "informationProtectionLabelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/informationProtection/policy/labels/informationProtectionLabelId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionPolicyLabelID(v.Input)
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

		if actual.InformationProtectionLabelId != v.Expected.InformationProtectionLabelId {
			t.Fatalf("Expected %q but got %q for InformationProtectionLabelId", v.Expected.InformationProtectionLabelId, actual.InformationProtectionLabelId)
		}

	}
}

func TestParseUserIdInformationProtectionPolicyLabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionPolicyLabelId
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
			Input: "/users/userId/informationProtection/policy",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/pOlIcY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/policy/labels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/pOlIcY/lAbElS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/informationProtection/policy/labels/informationProtectionLabelId",
			Expected: &UserIdInformationProtectionPolicyLabelId{
				UserId:                       "userId",
				InformationProtectionLabelId: "informationProtectionLabelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/informationProtection/policy/labels/informationProtectionLabelId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/pOlIcY/lAbElS/iNfOrMaTiOnPrOtEcTiOnLaBeLiD",
			Expected: &UserIdInformationProtectionPolicyLabelId{
				UserId:                       "uSeRiD",
				InformationProtectionLabelId: "iNfOrMaTiOnPrOtEcTiOnLaBeLiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/pOlIcY/lAbElS/iNfOrMaTiOnPrOtEcTiOnLaBeLiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionPolicyLabelIDInsensitively(v.Input)
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

		if actual.InformationProtectionLabelId != v.Expected.InformationProtectionLabelId {
			t.Fatalf("Expected %q but got %q for InformationProtectionLabelId", v.Expected.InformationProtectionLabelId, actual.InformationProtectionLabelId)
		}

	}
}

func TestSegmentsForUserIdInformationProtectionPolicyLabelId(t *testing.T) {
	segments := UserIdInformationProtectionPolicyLabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdInformationProtectionPolicyLabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
