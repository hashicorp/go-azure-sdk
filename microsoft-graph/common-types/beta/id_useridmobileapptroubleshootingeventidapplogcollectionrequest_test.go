package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}

func TestNewUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(t *testing.T) {
	id := NewUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID("userId", "mobileAppTroubleshootingEventId", "appLogCollectionRequestId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.MobileAppTroubleshootingEventId != "mobileAppTroubleshootingEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppTroubleshootingEventId'", id.MobileAppTroubleshootingEventId, "mobileAppTroubleshootingEventId")
	}

	if id.AppLogCollectionRequestId != "appLogCollectionRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppLogCollectionRequestId'", id.AppLogCollectionRequestId, "appLogCollectionRequestId")
	}
}

func TestFormatUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(t *testing.T) {
	actual := NewUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID("userId", "mobileAppTroubleshootingEventId", "appLogCollectionRequestId").ID()
	expected := "/users/userId/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId
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
			Input: "/users/userId/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId",
			Expected: &UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
				UserId:                          "userId",
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventId",
				AppLogCollectionRequestId:       "appLogCollectionRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(v.Input)
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

		if actual.MobileAppTroubleshootingEventId != v.Expected.MobileAppTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for MobileAppTroubleshootingEventId", v.Expected.MobileAppTroubleshootingEventId, actual.MobileAppTroubleshootingEventId)
		}

		if actual.AppLogCollectionRequestId != v.Expected.AppLogCollectionRequestId {
			t.Fatalf("Expected %q but got %q for AppLogCollectionRequestId", v.Expected.AppLogCollectionRequestId, actual.AppLogCollectionRequestId)
		}

	}
}

func TestParseUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId
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
			Input: "/users/userId/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mObIlEaPpTrOuBlEsHoOtInGeVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId/aPpLoGcOlLeCtIoNrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId",
			Expected: &UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
				UserId:                          "userId",
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventId",
				AppLogCollectionRequestId:       "appLogCollectionRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/appLogCollectionRequests/appLogCollectionRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId/aPpLoGcOlLeCtIoNrEqUeStS/aPpLoGcOlLeCtIoNrEqUeStId",
			Expected: &UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
				UserId:                          "uSeRiD",
				MobileAppTroubleshootingEventId: "mObIlEaPpTrOuBlEsHoOtInGeVeNtId",
				AppLogCollectionRequestId:       "aPpLoGcOlLeCtIoNrEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId/aPpLoGcOlLeCtIoNrEqUeStS/aPpLoGcOlLeCtIoNrEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively(v.Input)
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

		if actual.MobileAppTroubleshootingEventId != v.Expected.MobileAppTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for MobileAppTroubleshootingEventId", v.Expected.MobileAppTroubleshootingEventId, actual.MobileAppTroubleshootingEventId)
		}

		if actual.AppLogCollectionRequestId != v.Expected.AppLogCollectionRequestId {
			t.Fatalf("Expected %q but got %q for AppLogCollectionRequestId", v.Expected.AppLogCollectionRequestId, actual.AppLogCollectionRequestId)
		}

	}
}

func TestSegmentsForUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId(t *testing.T) {
	segments := UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
