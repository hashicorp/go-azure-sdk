package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMobileAppTroubleshootingEventId{}

func TestNewDeviceManagementMobileAppTroubleshootingEventID(t *testing.T) {
	id := NewDeviceManagementMobileAppTroubleshootingEventID("mobileAppTroubleshootingEventId")

	if id.MobileAppTroubleshootingEventId != "mobileAppTroubleshootingEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppTroubleshootingEventId'", id.MobileAppTroubleshootingEventId, "mobileAppTroubleshootingEventId")
	}
}

func TestFormatDeviceManagementMobileAppTroubleshootingEventID(t *testing.T) {
	actual := NewDeviceManagementMobileAppTroubleshootingEventID("mobileAppTroubleshootingEventId").ID()
	expected := "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMobileAppTroubleshootingEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMobileAppTroubleshootingEventId
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
			Input: "/deviceManagement/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId",
			Expected: &DeviceManagementMobileAppTroubleshootingEventId{
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMobileAppTroubleshootingEventID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MobileAppTroubleshootingEventId != v.Expected.MobileAppTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for MobileAppTroubleshootingEventId", v.Expected.MobileAppTroubleshootingEventId, actual.MobileAppTroubleshootingEventId)
		}

	}
}

func TestParseDeviceManagementMobileAppTroubleshootingEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMobileAppTroubleshootingEventId
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
			Input: "/deviceManagement/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEaPpTrOuBlEsHoOtInGeVeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId",
			Expected: &DeviceManagementMobileAppTroubleshootingEventId{
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId",
			Expected: &DeviceManagementMobileAppTroubleshootingEventId{
				MobileAppTroubleshootingEventId: "mObIlEaPpTrOuBlEsHoOtInGeVeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMobileAppTroubleshootingEventIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MobileAppTroubleshootingEventId != v.Expected.MobileAppTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for MobileAppTroubleshootingEventId", v.Expected.MobileAppTroubleshootingEventId, actual.MobileAppTroubleshootingEventId)
		}

	}
}

func TestSegmentsForDeviceManagementMobileAppTroubleshootingEventId(t *testing.T) {
	segments := DeviceManagementMobileAppTroubleshootingEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMobileAppTroubleshootingEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
