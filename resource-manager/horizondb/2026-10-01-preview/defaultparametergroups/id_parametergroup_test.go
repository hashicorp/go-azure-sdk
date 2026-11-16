package defaultparametergroups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ParameterGroupId{}

func TestNewParameterGroupID(t *testing.T) {
	id := NewParameterGroupID("12345678-1234-9876-4563-123456789012", "locationName", "parameterGroupName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.ParameterGroupName != "parameterGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'ParameterGroupName'", id.ParameterGroupName, "parameterGroupName")
	}
}

func TestFormatParameterGroupID(t *testing.T) {
	actual := NewParameterGroupID("12345678-1234-9876-4563-123456789012", "locationName", "parameterGroupName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations/locationName/parameterGroups/parameterGroupName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseParameterGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ParameterGroupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations/locationName/parameterGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations/locationName/parameterGroups/parameterGroupName",
			Expected: &ParameterGroupId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				LocationName:       "locationName",
				ParameterGroupName: "parameterGroupName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations/locationName/parameterGroups/parameterGroupName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseParameterGroupID(v.Input)
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

		if actual.ParameterGroupName != v.Expected.ParameterGroupName {
			t.Fatalf("Expected %q but got %q for ParameterGroupName", v.Expected.ParameterGroupName, actual.ParameterGroupName)
		}

	}
}

func TestParseParameterGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ParameterGroupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hOrIzOnDb",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/lOcAtIoNs/lOcAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations/locationName/parameterGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/lOcAtIoNs/lOcAtIoNnAmE/pArAmEtErGrOuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations/locationName/parameterGroups/parameterGroupName",
			Expected: &ParameterGroupId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				LocationName:       "locationName",
				ParameterGroupName: "parameterGroupName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HorizonDb/locations/locationName/parameterGroups/parameterGroupName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/lOcAtIoNs/lOcAtIoNnAmE/pArAmEtErGrOuPs/pArAmEtErGrOuPnAmE",
			Expected: &ParameterGroupId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				LocationName:       "lOcAtIoNnAmE",
				ParameterGroupName: "pArAmEtErGrOuPnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/lOcAtIoNs/lOcAtIoNnAmE/pArAmEtErGrOuPs/pArAmEtErGrOuPnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseParameterGroupIDInsensitively(v.Input)
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

		if actual.ParameterGroupName != v.Expected.ParameterGroupName {
			t.Fatalf("Expected %q but got %q for ParameterGroupName", v.Expected.ParameterGroupName, actual.ParameterGroupName)
		}

	}
}

func TestSegmentsForParameterGroupId(t *testing.T) {
	segments := ParameterGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ParameterGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %d unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
