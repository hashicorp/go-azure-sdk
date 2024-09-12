package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}

func TestNewUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(t *testing.T) {
	id := NewUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID("userIdValue", "mobileAppTroubleshootingEventIdValue", "appLogCollectionRequestIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.MobileAppTroubleshootingEventId != "mobileAppTroubleshootingEventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppTroubleshootingEventId'", id.MobileAppTroubleshootingEventId, "mobileAppTroubleshootingEventIdValue")
	}

	if id.AppLogCollectionRequestId != "appLogCollectionRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppLogCollectionRequestId'", id.AppLogCollectionRequestId, "appLogCollectionRequestIdValue")
	}
}

func TestFormatUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(t *testing.T) {
	actual := NewUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID("userIdValue", "mobileAppTroubleshootingEventIdValue", "appLogCollectionRequestIdValue").ID()
	expected := "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue"
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue",
			Expected: &UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
				UserId:                          "userIdValue",
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventIdValue",
				AppLogCollectionRequestId:       "appLogCollectionRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue/extra",
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
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpTrOuBlEsHoOtInGeVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE/aPpLoGcOlLeCtIoNrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue",
			Expected: &UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
				UserId:                          "userIdValue",
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventIdValue",
				AppLogCollectionRequestId:       "appLogCollectionRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/appLogCollectionRequests/appLogCollectionRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE/aPpLoGcOlLeCtIoNrEqUeStS/aPpLoGcOlLeCtIoNrEqUeStIdVaLuE",
			Expected: &UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
				UserId:                          "uSeRiDvAlUe",
				MobileAppTroubleshootingEventId: "mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE",
				AppLogCollectionRequestId:       "aPpLoGcOlLeCtIoNrEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE/aPpLoGcOlLeCtIoNrEqUeStS/aPpLoGcOlLeCtIoNrEqUeStIdVaLuE/extra",
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
