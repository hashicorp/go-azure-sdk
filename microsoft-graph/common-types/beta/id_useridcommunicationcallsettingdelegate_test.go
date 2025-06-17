package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCommunicationCallSettingDelegateId{}

func TestNewUserIdCommunicationCallSettingDelegateID(t *testing.T) {
	id := NewUserIdCommunicationCallSettingDelegateID("userId", "delegationSettingsId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DelegationSettingsId != "delegationSettingsId" {
		t.Fatalf("Expected %q but got %q for Segment 'DelegationSettingsId'", id.DelegationSettingsId, "delegationSettingsId")
	}
}

func TestFormatUserIdCommunicationCallSettingDelegateID(t *testing.T) {
	actual := NewUserIdCommunicationCallSettingDelegateID("userId", "delegationSettingsId").ID()
	expected := "/users/userId/communications/callSettings/delegates/delegationSettingsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdCommunicationCallSettingDelegateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCommunicationCallSettingDelegateId
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
			Input: "/users/userId/communications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/communications/callSettings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/communications/callSettings/delegates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/communications/callSettings/delegates/delegationSettingsId",
			Expected: &UserIdCommunicationCallSettingDelegateId{
				UserId:               "userId",
				DelegationSettingsId: "delegationSettingsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/communications/callSettings/delegates/delegationSettingsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCommunicationCallSettingDelegateID(v.Input)
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

		if actual.DelegationSettingsId != v.Expected.DelegationSettingsId {
			t.Fatalf("Expected %q but got %q for DelegationSettingsId", v.Expected.DelegationSettingsId, actual.DelegationSettingsId)
		}

	}
}

func TestParseUserIdCommunicationCallSettingDelegateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCommunicationCallSettingDelegateId
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
			Input: "/users/userId/communications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOmMuNiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/communications/callSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOmMuNiCaTiOnS/cAlLsEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/communications/callSettings/delegates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOmMuNiCaTiOnS/cAlLsEtTiNgS/dElEgAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/communications/callSettings/delegates/delegationSettingsId",
			Expected: &UserIdCommunicationCallSettingDelegateId{
				UserId:               "userId",
				DelegationSettingsId: "delegationSettingsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/communications/callSettings/delegates/delegationSettingsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOmMuNiCaTiOnS/cAlLsEtTiNgS/dElEgAtEs/dElEgAtIoNsEtTiNgSiD",
			Expected: &UserIdCommunicationCallSettingDelegateId{
				UserId:               "uSeRiD",
				DelegationSettingsId: "dElEgAtIoNsEtTiNgSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOmMuNiCaTiOnS/cAlLsEtTiNgS/dElEgAtEs/dElEgAtIoNsEtTiNgSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCommunicationCallSettingDelegateIDInsensitively(v.Input)
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

		if actual.DelegationSettingsId != v.Expected.DelegationSettingsId {
			t.Fatalf("Expected %q but got %q for DelegationSettingsId", v.Expected.DelegationSettingsId, actual.DelegationSettingsId)
		}

	}
}

func TestSegmentsForUserIdCommunicationCallSettingDelegateId(t *testing.T) {
	segments := UserIdCommunicationCallSettingDelegateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdCommunicationCallSettingDelegateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
