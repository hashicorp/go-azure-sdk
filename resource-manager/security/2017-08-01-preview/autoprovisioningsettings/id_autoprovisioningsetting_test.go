package autoprovisioningsettings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AutoProvisioningSettingId{}

func TestNewAutoProvisioningSettingID(t *testing.T) {
	id := NewAutoProvisioningSettingID("12345678-1234-9876-4563-123456789012", "autoProvisioningSettingValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.AutoProvisioningSettingName != "autoProvisioningSettingValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AutoProvisioningSettingName'", id.AutoProvisioningSettingName, "autoProvisioningSettingValue")
	}
}

func TestFormatAutoProvisioningSettingID(t *testing.T) {
	actual := NewAutoProvisioningSettingID("12345678-1234-9876-4563-123456789012", "autoProvisioningSettingValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/autoProvisioningSettings/autoProvisioningSettingValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAutoProvisioningSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutoProvisioningSettingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/autoProvisioningSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/autoProvisioningSettings/autoProvisioningSettingValue",
			Expected: &AutoProvisioningSettingId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				AutoProvisioningSettingName: "autoProvisioningSettingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/autoProvisioningSettings/autoProvisioningSettingValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutoProvisioningSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.AutoProvisioningSettingName != v.Expected.AutoProvisioningSettingName {
			t.Fatalf("Expected %q but got %q for AutoProvisioningSettingName", v.Expected.AutoProvisioningSettingName, actual.AutoProvisioningSettingName)
		}

	}
}

func TestParseAutoProvisioningSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutoProvisioningSettingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/autoProvisioningSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/aUtOpRoViSiOnInGsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/autoProvisioningSettings/autoProvisioningSettingValue",
			Expected: &AutoProvisioningSettingId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				AutoProvisioningSettingName: "autoProvisioningSettingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/autoProvisioningSettings/autoProvisioningSettingValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/aUtOpRoViSiOnInGsEtTiNgS/aUtOpRoViSiOnInGsEtTiNgVaLuE",
			Expected: &AutoProvisioningSettingId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				AutoProvisioningSettingName: "aUtOpRoViSiOnInGsEtTiNgVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/aUtOpRoViSiOnInGsEtTiNgS/aUtOpRoViSiOnInGsEtTiNgVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutoProvisioningSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.AutoProvisioningSettingName != v.Expected.AutoProvisioningSettingName {
			t.Fatalf("Expected %q but got %q for AutoProvisioningSettingName", v.Expected.AutoProvisioningSettingName, actual.AutoProvisioningSettingName)
		}

	}
}

func TestSegmentsForAutoProvisioningSettingId(t *testing.T) {
	segments := AutoProvisioningSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AutoProvisioningSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
