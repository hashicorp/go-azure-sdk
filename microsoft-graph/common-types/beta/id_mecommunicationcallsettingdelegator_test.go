package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCommunicationCallSettingDelegatorId{}

func TestNewMeCommunicationCallSettingDelegatorID(t *testing.T) {
	id := NewMeCommunicationCallSettingDelegatorID("delegationSettingsId")

	if id.DelegationSettingsId != "delegationSettingsId" {
		t.Fatalf("Expected %q but got %q for Segment 'DelegationSettingsId'", id.DelegationSettingsId, "delegationSettingsId")
	}
}

func TestFormatMeCommunicationCallSettingDelegatorID(t *testing.T) {
	actual := NewMeCommunicationCallSettingDelegatorID("delegationSettingsId").ID()
	expected := "/me/communications/callSettings/delegators/delegationSettingsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCommunicationCallSettingDelegatorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCommunicationCallSettingDelegatorId
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
			Input: "/me/communications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/communications/callSettings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/communications/callSettings/delegators",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/communications/callSettings/delegators/delegationSettingsId",
			Expected: &MeCommunicationCallSettingDelegatorId{
				DelegationSettingsId: "delegationSettingsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/communications/callSettings/delegators/delegationSettingsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCommunicationCallSettingDelegatorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DelegationSettingsId != v.Expected.DelegationSettingsId {
			t.Fatalf("Expected %q but got %q for DelegationSettingsId", v.Expected.DelegationSettingsId, actual.DelegationSettingsId)
		}

	}
}

func TestParseMeCommunicationCallSettingDelegatorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCommunicationCallSettingDelegatorId
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
			Input: "/me/communications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOmMuNiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/communications/callSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOmMuNiCaTiOnS/cAlLsEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/communications/callSettings/delegators",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOmMuNiCaTiOnS/cAlLsEtTiNgS/dElEgAtOrS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/communications/callSettings/delegators/delegationSettingsId",
			Expected: &MeCommunicationCallSettingDelegatorId{
				DelegationSettingsId: "delegationSettingsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/communications/callSettings/delegators/delegationSettingsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOmMuNiCaTiOnS/cAlLsEtTiNgS/dElEgAtOrS/dElEgAtIoNsEtTiNgSiD",
			Expected: &MeCommunicationCallSettingDelegatorId{
				DelegationSettingsId: "dElEgAtIoNsEtTiNgSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cOmMuNiCaTiOnS/cAlLsEtTiNgS/dElEgAtOrS/dElEgAtIoNsEtTiNgSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCommunicationCallSettingDelegatorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DelegationSettingsId != v.Expected.DelegationSettingsId {
			t.Fatalf("Expected %q but got %q for DelegationSettingsId", v.Expected.DelegationSettingsId, actual.DelegationSettingsId)
		}

	}
}

func TestSegmentsForMeCommunicationCallSettingDelegatorId(t *testing.T) {
	segments := MeCommunicationCallSettingDelegatorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCommunicationCallSettingDelegatorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
