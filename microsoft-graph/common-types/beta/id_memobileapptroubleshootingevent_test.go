package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMobileAppTroubleshootingEventId{}

func TestNewMeMobileAppTroubleshootingEventID(t *testing.T) {
	id := NewMeMobileAppTroubleshootingEventID("mobileAppTroubleshootingEventId")

	if id.MobileAppTroubleshootingEventId != "mobileAppTroubleshootingEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppTroubleshootingEventId'", id.MobileAppTroubleshootingEventId, "mobileAppTroubleshootingEventId")
	}
}

func TestFormatMeMobileAppTroubleshootingEventID(t *testing.T) {
	actual := NewMeMobileAppTroubleshootingEventID("mobileAppTroubleshootingEventId").ID()
	expected := "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMobileAppTroubleshootingEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMobileAppTroubleshootingEventId
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
			// Valid URI
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId",
			Expected: &MeMobileAppTroubleshootingEventId{
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMobileAppTroubleshootingEventID(v.Input)
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

func TestParseMeMobileAppTroubleshootingEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMobileAppTroubleshootingEventId
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
			// Valid URI
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId",
			Expected: &MeMobileAppTroubleshootingEventId{
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId",
			Expected: &MeMobileAppTroubleshootingEventId{
				MobileAppTroubleshootingEventId: "mObIlEaPpTrOuBlEsHoOtInGeVeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMobileAppTroubleshootingEventIDInsensitively(v.Input)
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

func TestSegmentsForMeMobileAppTroubleshootingEventId(t *testing.T) {
	segments := MeMobileAppTroubleshootingEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMobileAppTroubleshootingEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
