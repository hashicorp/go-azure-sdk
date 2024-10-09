package autonomousdatabasenationalcharactersets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AutonomousDatabaseNationalCharacterSetId{}

func TestNewAutonomousDatabaseNationalCharacterSetID(t *testing.T) {
	id := NewAutonomousDatabaseNationalCharacterSetID("12345678-1234-9876-4563-123456789012", "locationName", "autonomousDatabaseNationalCharacterSetName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.AutonomousDatabaseNationalCharacterSetName != "autonomousDatabaseNationalCharacterSetName" {
		t.Fatalf("Expected %q but got %q for Segment 'AutonomousDatabaseNationalCharacterSetName'", id.AutonomousDatabaseNationalCharacterSetName, "autonomousDatabaseNationalCharacterSetName")
	}
}

func TestFormatAutonomousDatabaseNationalCharacterSetID(t *testing.T) {
	actual := NewAutonomousDatabaseNationalCharacterSetID("12345678-1234-9876-4563-123456789012", "locationName", "autonomousDatabaseNationalCharacterSetName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseNationalCharacterSets/autonomousDatabaseNationalCharacterSetName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAutonomousDatabaseNationalCharacterSetID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutonomousDatabaseNationalCharacterSetId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseNationalCharacterSets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseNationalCharacterSets/autonomousDatabaseNationalCharacterSetName",
			Expected: &AutonomousDatabaseNationalCharacterSetId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationName",
				AutonomousDatabaseNationalCharacterSetName: "autonomousDatabaseNationalCharacterSetName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseNationalCharacterSets/autonomousDatabaseNationalCharacterSetName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutonomousDatabaseNationalCharacterSetID(v.Input)
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

		if actual.AutonomousDatabaseNationalCharacterSetName != v.Expected.AutonomousDatabaseNationalCharacterSetName {
			t.Fatalf("Expected %q but got %q for AutonomousDatabaseNationalCharacterSetName", v.Expected.AutonomousDatabaseNationalCharacterSetName, actual.AutonomousDatabaseNationalCharacterSetName)
		}

	}
}

func TestParseAutonomousDatabaseNationalCharacterSetIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutonomousDatabaseNationalCharacterSetId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseNationalCharacterSets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoNnAmE/aUtOnOmOuSdAtAbAsEnAtIoNaLcHaRaCtErSeTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseNationalCharacterSets/autonomousDatabaseNationalCharacterSetName",
			Expected: &AutonomousDatabaseNationalCharacterSetId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationName",
				AutonomousDatabaseNationalCharacterSetName: "autonomousDatabaseNationalCharacterSetName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/locationName/autonomousDatabaseNationalCharacterSets/autonomousDatabaseNationalCharacterSetName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoNnAmE/aUtOnOmOuSdAtAbAsEnAtIoNaLcHaRaCtErSeTs/aUtOnOmOuSdAtAbAsEnAtIoNaLcHaRaCtErSeTnAmE",
			Expected: &AutonomousDatabaseNationalCharacterSetId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "lOcAtIoNnAmE",
				AutonomousDatabaseNationalCharacterSetName: "aUtOnOmOuSdAtAbAsEnAtIoNaLcHaRaCtErSeTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoNnAmE/aUtOnOmOuSdAtAbAsEnAtIoNaLcHaRaCtErSeTs/aUtOnOmOuSdAtAbAsEnAtIoNaLcHaRaCtErSeTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutonomousDatabaseNationalCharacterSetIDInsensitively(v.Input)
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

		if actual.AutonomousDatabaseNationalCharacterSetName != v.Expected.AutonomousDatabaseNationalCharacterSetName {
			t.Fatalf("Expected %q but got %q for AutonomousDatabaseNationalCharacterSetName", v.Expected.AutonomousDatabaseNationalCharacterSetName, actual.AutonomousDatabaseNationalCharacterSetName)
		}

	}
}

func TestSegmentsForAutonomousDatabaseNationalCharacterSetId(t *testing.T) {
	segments := AutonomousDatabaseNationalCharacterSetId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AutonomousDatabaseNationalCharacterSetId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
