package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSettingStorageQuotaServiceId{}

func TestNewMeSettingStorageQuotaServiceID(t *testing.T) {
	id := NewMeSettingStorageQuotaServiceID("serviceStorageQuotaBreakdownId")

	if id.ServiceStorageQuotaBreakdownId != "serviceStorageQuotaBreakdownId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServiceStorageQuotaBreakdownId'", id.ServiceStorageQuotaBreakdownId, "serviceStorageQuotaBreakdownId")
	}
}

func TestFormatMeSettingStorageQuotaServiceID(t *testing.T) {
	actual := NewMeSettingStorageQuotaServiceID("serviceStorageQuotaBreakdownId").ID()
	expected := "/me/settings/storage/quota/services/serviceStorageQuotaBreakdownId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeSettingStorageQuotaServiceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeSettingStorageQuotaServiceId
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
			Input: "/me/settings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/storage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/storage/quota",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/storage/quota/services",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/settings/storage/quota/services/serviceStorageQuotaBreakdownId",
			Expected: &MeSettingStorageQuotaServiceId{
				ServiceStorageQuotaBreakdownId: "serviceStorageQuotaBreakdownId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/settings/storage/quota/services/serviceStorageQuotaBreakdownId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeSettingStorageQuotaServiceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServiceStorageQuotaBreakdownId != v.Expected.ServiceStorageQuotaBreakdownId {
			t.Fatalf("Expected %q but got %q for ServiceStorageQuotaBreakdownId", v.Expected.ServiceStorageQuotaBreakdownId, actual.ServiceStorageQuotaBreakdownId)
		}

	}
}

func TestParseMeSettingStorageQuotaServiceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeSettingStorageQuotaServiceId
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
			Input: "/me/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/storage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/sToRaGe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/storage/quota",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/sToRaGe/qUoTa",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/storage/quota/services",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/sToRaGe/qUoTa/sErViCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/settings/storage/quota/services/serviceStorageQuotaBreakdownId",
			Expected: &MeSettingStorageQuotaServiceId{
				ServiceStorageQuotaBreakdownId: "serviceStorageQuotaBreakdownId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/settings/storage/quota/services/serviceStorageQuotaBreakdownId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/sToRaGe/qUoTa/sErViCeS/sErViCeStOrAgEqUoTaBrEaKdOwNiD",
			Expected: &MeSettingStorageQuotaServiceId{
				ServiceStorageQuotaBreakdownId: "sErViCeStOrAgEqUoTaBrEaKdOwNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/sToRaGe/qUoTa/sErViCeS/sErViCeStOrAgEqUoTaBrEaKdOwNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeSettingStorageQuotaServiceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServiceStorageQuotaBreakdownId != v.Expected.ServiceStorageQuotaBreakdownId {
			t.Fatalf("Expected %q but got %q for ServiceStorageQuotaBreakdownId", v.Expected.ServiceStorageQuotaBreakdownId, actual.ServiceStorageQuotaBreakdownId)
		}

	}
}

func TestSegmentsForMeSettingStorageQuotaServiceId(t *testing.T) {
	segments := MeSettingStorageQuotaServiceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeSettingStorageQuotaServiceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
