package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}

func TestNewDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID(t *testing.T) {
	id := NewDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID("mobileAppTroubleshootingEventId", "appLogCollectionRequestId")

	if id.MobileAppTroubleshootingEventId != "mobileAppTroubleshootingEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppTroubleshootingEventId'", id.MobileAppTroubleshootingEventId, "mobileAppTroubleshootingEventId")
	}

	if id.AppLogCollectionRequestId != "appLogCollectionRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppLogCollectionRequestId'", id.AppLogCollectionRequestId, "appLogCollectionRequestId")
	}
}

func TestFormatDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID(t *testing.T) {
	actual := NewDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID("mobileAppTroubleshootingEventId", "appLogCollectionRequestId").ID()
	expected := "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId
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
			// Incomplete URI
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId",
			Expected: &DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventId",
				AppLogCollectionRequestId:       "appLogCollectionRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID(v.Input)
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

		if actual.AppLogCollectionRequestId != v.Expected.AppLogCollectionRequestId {
			t.Fatalf("Expected %q but got %q for AppLogCollectionRequestId", v.Expected.AppLogCollectionRequestId, actual.AppLogCollectionRequestId)
		}

	}
}

func TestParseDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId
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
			// Incomplete URI
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId/aPpLoGcOlLeCtIoNrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId",
			Expected: &DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventId",
				AppLogCollectionRequestId:       "appLogCollectionRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId/aPpLoGcOlLeCtIoNrEqUeStS/aPpLoGcOlLeCtIoNrEqUeStId",
			Expected: &DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
				MobileAppTroubleshootingEventId: "mObIlEaPpTrOuBlEsHoOtInGeVeNtId",
				AppLogCollectionRequestId:       "aPpLoGcOlLeCtIoNrEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId/aPpLoGcOlLeCtIoNrEqUeStS/aPpLoGcOlLeCtIoNrEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively(v.Input)
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

		if actual.AppLogCollectionRequestId != v.Expected.AppLogCollectionRequestId {
			t.Fatalf("Expected %q but got %q for AppLogCollectionRequestId", v.Expected.AppLogCollectionRequestId, actual.AppLogCollectionRequestId)
		}

	}
}

func TestSegmentsForDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId(t *testing.T) {
	segments := DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
