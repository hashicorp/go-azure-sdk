package autonomousdatabasecharactersets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AutonomousDatabaseCharacterSetId{}

func TestNewAutonomousDatabaseCharacterSetID(t *testing.T) {
	id := NewAutonomousDatabaseCharacterSetID("12345678-1234-9876-4563-123456789012", "locationName", "autonomousDatabaseCharacterSetName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.AutonomousDatabaseCharacterSetName != "autonomousDatabaseCharacterSetName" {
		t.Fatalf("Expected %q but got %q for Segment 'AutonomousDatabaseCharacterSetName'", id.AutonomousDatabaseCharacterSetName, "autonomousDatabaseCharacterSetName")
	}
}

func TestFormatAutonomousDatabaseCharacterSetID(t *testing.T) {
	actual := NewAutonomousDatabaseCharacterSetID("12345678-1234-9876-4563-123456789012", "locationName", "autonomousDatabaseCharacterSetName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseCharacterSets/autonomousDatabaseCharacterSetName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAutonomousDatabaseCharacterSetID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutonomousDatabaseCharacterSetId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseCharacterSets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseCharacterSets/autonomousDatabaseCharacterSetName",
			Expected: &AutonomousDatabaseCharacterSetId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				LocationName:                       "locationName",
				AutonomousDatabaseCharacterSetName: "autonomousDatabaseCharacterSetName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseCharacterSets/autonomousDatabaseCharacterSetName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutonomousDatabaseCharacterSetID(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.AutonomousDatabaseCharacterSetName != v.Expected.AutonomousDatabaseCharacterSetName {
			t.Fatalf("Expected %q but got %q for AutonomousDatabaseCharacterSetName", v.Expected.AutonomousDatabaseCharacterSetName, actual.AutonomousDatabaseCharacterSetName)
		}

	}
}

func TestParseAutonomousDatabaseCharacterSetIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutonomousDatabaseCharacterSetId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseCharacterSets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoNnAmE/aUtOnOmOuSdAtAbAsEcHaRaCtErSeTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseCharacterSets/autonomousDatabaseCharacterSetName",
			Expected: &AutonomousDatabaseCharacterSetId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				LocationName:                       "locationName",
				AutonomousDatabaseCharacterSetName: "autonomousDatabaseCharacterSetName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseCharacterSets/autonomousDatabaseCharacterSetName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoNnAmE/aUtOnOmOuSdAtAbAsEcHaRaCtErSeTs/aUtOnOmOuSdAtAbAsEcHaRaCtErSeTnAmE",
			Expected: &AutonomousDatabaseCharacterSetId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				LocationName:                       "lOcAtIoNnAmE",
				AutonomousDatabaseCharacterSetName: "aUtOnOmOuSdAtAbAsEcHaRaCtErSeTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoNnAmE/aUtOnOmOuSdAtAbAsEcHaRaCtErSeTs/aUtOnOmOuSdAtAbAsEcHaRaCtErSeTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutonomousDatabaseCharacterSetIDInsensitively(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.AutonomousDatabaseCharacterSetName != v.Expected.AutonomousDatabaseCharacterSetName {
			t.Fatalf("Expected %q but got %q for AutonomousDatabaseCharacterSetName", v.Expected.AutonomousDatabaseCharacterSetName, actual.AutonomousDatabaseCharacterSetName)
		}

	}
}

func TestSegmentsForAutonomousDatabaseCharacterSetId(t *testing.T) {
	segments := AutonomousDatabaseCharacterSetId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AutonomousDatabaseCharacterSetId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
