package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdSettingStorageQuotaServiceId{}

func TestNewUserIdSettingStorageQuotaServiceID(t *testing.T) {
	id := NewUserIdSettingStorageQuotaServiceID("userIdValue", "serviceStorageQuotaBreakdownIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ServiceStorageQuotaBreakdownId != "serviceStorageQuotaBreakdownIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServiceStorageQuotaBreakdownId'", id.ServiceStorageQuotaBreakdownId, "serviceStorageQuotaBreakdownIdValue")
	}
}

func TestFormatUserIdSettingStorageQuotaServiceID(t *testing.T) {
	actual := NewUserIdSettingStorageQuotaServiceID("userIdValue", "serviceStorageQuotaBreakdownIdValue").ID()
	expected := "/users/userIdValue/settings/storage/quota/services/serviceStorageQuotaBreakdownIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdSettingStorageQuotaServiceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdSettingStorageQuotaServiceId
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
			Input: "/users/userIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/settings/storage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/settings/storage/quota",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/settings/storage/quota/services",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/settings/storage/quota/services/serviceStorageQuotaBreakdownIdValue",
			Expected: &UserIdSettingStorageQuotaServiceId{
				UserId:                         "userIdValue",
				ServiceStorageQuotaBreakdownId: "serviceStorageQuotaBreakdownIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/settings/storage/quota/services/serviceStorageQuotaBreakdownIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdSettingStorageQuotaServiceID(v.Input)
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

		if actual.ServiceStorageQuotaBreakdownId != v.Expected.ServiceStorageQuotaBreakdownId {
			t.Fatalf("Expected %q but got %q for ServiceStorageQuotaBreakdownId", v.Expected.ServiceStorageQuotaBreakdownId, actual.ServiceStorageQuotaBreakdownId)
		}

	}
}

func TestParseUserIdSettingStorageQuotaServiceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdSettingStorageQuotaServiceId
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
			Input: "/users/userIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/settings/storage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEtTiNgS/sToRaGe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/settings/storage/quota",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEtTiNgS/sToRaGe/qUoTa",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/settings/storage/quota/services",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEtTiNgS/sToRaGe/qUoTa/sErViCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/settings/storage/quota/services/serviceStorageQuotaBreakdownIdValue",
			Expected: &UserIdSettingStorageQuotaServiceId{
				UserId:                         "userIdValue",
				ServiceStorageQuotaBreakdownId: "serviceStorageQuotaBreakdownIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/settings/storage/quota/services/serviceStorageQuotaBreakdownIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEtTiNgS/sToRaGe/qUoTa/sErViCeS/sErViCeStOrAgEqUoTaBrEaKdOwNiDvAlUe",
			Expected: &UserIdSettingStorageQuotaServiceId{
				UserId:                         "uSeRiDvAlUe",
				ServiceStorageQuotaBreakdownId: "sErViCeStOrAgEqUoTaBrEaKdOwNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/sEtTiNgS/sToRaGe/qUoTa/sErViCeS/sErViCeStOrAgEqUoTaBrEaKdOwNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdSettingStorageQuotaServiceIDInsensitively(v.Input)
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

		if actual.ServiceStorageQuotaBreakdownId != v.Expected.ServiceStorageQuotaBreakdownId {
			t.Fatalf("Expected %q but got %q for ServiceStorageQuotaBreakdownId", v.Expected.ServiceStorageQuotaBreakdownId, actual.ServiceStorageQuotaBreakdownId)
		}

	}
}

func TestSegmentsForUserIdSettingStorageQuotaServiceId(t *testing.T) {
	segments := UserIdSettingStorageQuotaServiceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdSettingStorageQuotaServiceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
