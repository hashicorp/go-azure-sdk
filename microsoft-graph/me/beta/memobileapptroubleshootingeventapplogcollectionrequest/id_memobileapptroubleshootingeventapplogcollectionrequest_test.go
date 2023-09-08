package memobileapptroubleshootingeventapplogcollectionrequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMobileAppTroubleshootingEventAppLogCollectionRequestId{}

func TestNewMeMobileAppTroubleshootingEventAppLogCollectionRequestID(t *testing.T) {
	id := NewMeMobileAppTroubleshootingEventAppLogCollectionRequestID("mobileAppTroubleshootingEventIdValue", "appLogCollectionRequestIdValue")

	if id.MobileAppTroubleshootingEventId != "mobileAppTroubleshootingEventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppTroubleshootingEventId'", id.MobileAppTroubleshootingEventId, "mobileAppTroubleshootingEventIdValue")
	}

	if id.AppLogCollectionRequestId != "appLogCollectionRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppLogCollectionRequestId'", id.AppLogCollectionRequestId, "appLogCollectionRequestIdValue")
	}
}

func TestFormatMeMobileAppTroubleshootingEventAppLogCollectionRequestID(t *testing.T) {
	actual := NewMeMobileAppTroubleshootingEventAppLogCollectionRequestID("mobileAppTroubleshootingEventIdValue", "appLogCollectionRequestIdValue").ID()
	expected := "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMobileAppTroubleshootingEventAppLogCollectionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMobileAppTroubleshootingEventAppLogCollectionRequestId
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
			Input: "/me/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue",
			Expected: &MeMobileAppTroubleshootingEventAppLogCollectionRequestId{
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventIdValue",
				AppLogCollectionRequestId:       "appLogCollectionRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMobileAppTroubleshootingEventAppLogCollectionRequestID(v.Input)
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

func TestParseMeMobileAppTroubleshootingEventAppLogCollectionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMobileAppTroubleshootingEventAppLogCollectionRequestId
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
			Input: "/me/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpTrOuBlEsHoOtInGeVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE/aPpLoGcOlLeCtIoNrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue",
			Expected: &MeMobileAppTroubleshootingEventAppLogCollectionRequestId{
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventIdValue",
				AppLogCollectionRequestId:       "appLogCollectionRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE/aPpLoGcOlLeCtIoNrEqUeStS/aPpLoGcOlLeCtIoNrEqUeStIdVaLuE",
			Expected: &MeMobileAppTroubleshootingEventAppLogCollectionRequestId{
				MobileAppTroubleshootingEventId: "mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE",
				AppLogCollectionRequestId:       "aPpLoGcOlLeCtIoNrEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE/aPpLoGcOlLeCtIoNrEqUeStS/aPpLoGcOlLeCtIoNrEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMobileAppTroubleshootingEventAppLogCollectionRequestIDInsensitively(v.Input)
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

func TestSegmentsForMeMobileAppTroubleshootingEventAppLogCollectionRequestId(t *testing.T) {
	segments := MeMobileAppTroubleshootingEventAppLogCollectionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMobileAppTroubleshootingEventAppLogCollectionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
