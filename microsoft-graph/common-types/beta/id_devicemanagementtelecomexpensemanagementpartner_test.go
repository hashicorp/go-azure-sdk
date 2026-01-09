package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTelecomExpenseManagementPartnerId{}

func TestNewDeviceManagementTelecomExpenseManagementPartnerID(t *testing.T) {
	id := NewDeviceManagementTelecomExpenseManagementPartnerID("telecomExpenseManagementPartnerId")

	if id.TelecomExpenseManagementPartnerId != "telecomExpenseManagementPartnerId" {
		t.Fatalf("Expected %q but got %q for Segment 'TelecomExpenseManagementPartnerId'", id.TelecomExpenseManagementPartnerId, "telecomExpenseManagementPartnerId")
	}
}

func TestFormatDeviceManagementTelecomExpenseManagementPartnerID(t *testing.T) {
	actual := NewDeviceManagementTelecomExpenseManagementPartnerID("telecomExpenseManagementPartnerId").ID()
	expected := "/deviceManagement/telecomExpenseManagementPartners/telecomExpenseManagementPartnerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTelecomExpenseManagementPartnerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTelecomExpenseManagementPartnerId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/telecomExpenseManagementPartners",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/telecomExpenseManagementPartners/telecomExpenseManagementPartnerId",
			Expected: &DeviceManagementTelecomExpenseManagementPartnerId{
				TelecomExpenseManagementPartnerId: "telecomExpenseManagementPartnerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/telecomExpenseManagementPartners/telecomExpenseManagementPartnerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTelecomExpenseManagementPartnerID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TelecomExpenseManagementPartnerId != v.Expected.TelecomExpenseManagementPartnerId {
			t.Fatalf("Expected %q but got %q for TelecomExpenseManagementPartnerId", v.Expected.TelecomExpenseManagementPartnerId, actual.TelecomExpenseManagementPartnerId)
		}

	}
}

func TestParseDeviceManagementTelecomExpenseManagementPartnerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTelecomExpenseManagementPartnerId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/telecomExpenseManagementPartners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tElEcOmExPeNsEmAnAgEmEnTpArTnErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/telecomExpenseManagementPartners/telecomExpenseManagementPartnerId",
			Expected: &DeviceManagementTelecomExpenseManagementPartnerId{
				TelecomExpenseManagementPartnerId: "telecomExpenseManagementPartnerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/telecomExpenseManagementPartners/telecomExpenseManagementPartnerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tElEcOmExPeNsEmAnAgEmEnTpArTnErS/tElEcOmExPeNsEmAnAgEmEnTpArTnErId",
			Expected: &DeviceManagementTelecomExpenseManagementPartnerId{
				TelecomExpenseManagementPartnerId: "tElEcOmExPeNsEmAnAgEmEnTpArTnErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tElEcOmExPeNsEmAnAgEmEnTpArTnErS/tElEcOmExPeNsEmAnAgEmEnTpArTnErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTelecomExpenseManagementPartnerIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TelecomExpenseManagementPartnerId != v.Expected.TelecomExpenseManagementPartnerId {
			t.Fatalf("Expected %q but got %q for TelecomExpenseManagementPartnerId", v.Expected.TelecomExpenseManagementPartnerId, actual.TelecomExpenseManagementPartnerId)
		}

	}
}

func TestSegmentsForDeviceManagementTelecomExpenseManagementPartnerId(t *testing.T) {
	segments := DeviceManagementTelecomExpenseManagementPartnerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTelecomExpenseManagementPartnerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
