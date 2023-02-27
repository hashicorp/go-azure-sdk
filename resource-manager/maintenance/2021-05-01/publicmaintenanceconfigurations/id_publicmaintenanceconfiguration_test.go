// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package publicmaintenanceconfigurations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PublicMaintenanceConfigurationId{}

func TestNewPublicMaintenanceConfigurationID(t *testing.T) {
	id := NewPublicMaintenanceConfigurationID("12345678-1234-9876-4563-123456789012", "publicMaintenanceConfigurationValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.PublicMaintenanceConfigurationName != "publicMaintenanceConfigurationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PublicMaintenanceConfigurationName'", id.PublicMaintenanceConfigurationName, "publicMaintenanceConfigurationValue")
	}
}

func TestFormatPublicMaintenanceConfigurationID(t *testing.T) {
	actual := NewPublicMaintenanceConfigurationID("12345678-1234-9876-4563-123456789012", "publicMaintenanceConfigurationValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/publicMaintenanceConfigurationValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePublicMaintenanceConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PublicMaintenanceConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Maintenance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Maintenance/publicMaintenanceConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/publicMaintenanceConfigurationValue",
			Expected: &PublicMaintenanceConfigurationId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				PublicMaintenanceConfigurationName: "publicMaintenanceConfigurationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/publicMaintenanceConfigurationValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePublicMaintenanceConfigurationID(v.Input)
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

		if actual.PublicMaintenanceConfigurationName != v.Expected.PublicMaintenanceConfigurationName {
			t.Fatalf("Expected %q but got %q for PublicMaintenanceConfigurationName", v.Expected.PublicMaintenanceConfigurationName, actual.PublicMaintenanceConfigurationName)
		}

	}
}

func TestParsePublicMaintenanceConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PublicMaintenanceConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Maintenance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Maintenance/publicMaintenanceConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe/pUbLiCmAiNtEnAnCeCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/publicMaintenanceConfigurationValue",
			Expected: &PublicMaintenanceConfigurationId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				PublicMaintenanceConfigurationName: "publicMaintenanceConfigurationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/publicMaintenanceConfigurationValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe/pUbLiCmAiNtEnAnCeCoNfIgUrAtIoNs/pUbLiCmAiNtEnAnCeCoNfIgUrAtIoNvAlUe",
			Expected: &PublicMaintenanceConfigurationId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				PublicMaintenanceConfigurationName: "pUbLiCmAiNtEnAnCeCoNfIgUrAtIoNvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe/pUbLiCmAiNtEnAnCeCoNfIgUrAtIoNs/pUbLiCmAiNtEnAnCeCoNfIgUrAtIoNvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePublicMaintenanceConfigurationIDInsensitively(v.Input)
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

		if actual.PublicMaintenanceConfigurationName != v.Expected.PublicMaintenanceConfigurationName {
			t.Fatalf("Expected %q but got %q for PublicMaintenanceConfigurationName", v.Expected.PublicMaintenanceConfigurationName, actual.PublicMaintenanceConfigurationName)
		}

	}
}

func TestSegmentsForPublicMaintenanceConfigurationId(t *testing.T) {
	segments := PublicMaintenanceConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PublicMaintenanceConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
