package environmenttypes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DevCenterEnvironmentTypeId{}

func TestNewDevCenterEnvironmentTypeID(t *testing.T) {
	id := NewDevCenterEnvironmentTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "devCenterValue", "environmentTypeValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.DevCenterName != "devCenterValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DevCenterName'", id.DevCenterName, "devCenterValue")
	}

	if id.EnvironmentTypeName != "environmentTypeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EnvironmentTypeName'", id.EnvironmentTypeName, "environmentTypeValue")
	}
}

func TestFormatDevCenterEnvironmentTypeID(t *testing.T) {
	actual := NewDevCenterEnvironmentTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "devCenterValue", "environmentTypeValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters/devCenterValue/environmentTypes/environmentTypeValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDevCenterEnvironmentTypeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DevCenterEnvironmentTypeId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters/devCenterValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters/devCenterValue/environmentTypes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters/devCenterValue/environmentTypes/environmentTypeValue",
			Expected: &DevCenterEnvironmentTypeId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				DevCenterName:       "devCenterValue",
				EnvironmentTypeName: "environmentTypeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters/devCenterValue/environmentTypes/environmentTypeValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDevCenterEnvironmentTypeID(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.DevCenterName != v.Expected.DevCenterName {
			t.Fatalf("Expected %q but got %q for DevCenterName", v.Expected.DevCenterName, actual.DevCenterName)
		}

		if actual.EnvironmentTypeName != v.Expected.EnvironmentTypeName {
			t.Fatalf("Expected %q but got %q for EnvironmentTypeName", v.Expected.EnvironmentTypeName, actual.EnvironmentTypeName)
		}

	}
}

func TestParseDevCenterEnvironmentTypeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DevCenterEnvironmentTypeId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEvCeNtEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEvCeNtEr/dEvCeNtErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters/devCenterValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEvCeNtEr/dEvCeNtErS/dEvCeNtErVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters/devCenterValue/environmentTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEvCeNtEr/dEvCeNtErS/dEvCeNtErVaLuE/eNvIrOnMeNtTyPeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters/devCenterValue/environmentTypes/environmentTypeValue",
			Expected: &DevCenterEnvironmentTypeId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				DevCenterName:       "devCenterValue",
				EnvironmentTypeName: "environmentTypeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DevCenter/devCenters/devCenterValue/environmentTypes/environmentTypeValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEvCeNtEr/dEvCeNtErS/dEvCeNtErVaLuE/eNvIrOnMeNtTyPeS/eNvIrOnMeNtTyPeVaLuE",
			Expected: &DevCenterEnvironmentTypeId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "eXaMpLe-rEsOuRcE-GrOuP",
				DevCenterName:       "dEvCeNtErVaLuE",
				EnvironmentTypeName: "eNvIrOnMeNtTyPeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEvCeNtEr/dEvCeNtErS/dEvCeNtErVaLuE/eNvIrOnMeNtTyPeS/eNvIrOnMeNtTyPeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDevCenterEnvironmentTypeIDInsensitively(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.DevCenterName != v.Expected.DevCenterName {
			t.Fatalf("Expected %q but got %q for DevCenterName", v.Expected.DevCenterName, actual.DevCenterName)
		}

		if actual.EnvironmentTypeName != v.Expected.EnvironmentTypeName {
			t.Fatalf("Expected %q but got %q for EnvironmentTypeName", v.Expected.EnvironmentTypeName, actual.EnvironmentTypeName)
		}

	}
}

func TestSegmentsForDevCenterEnvironmentTypeId(t *testing.T) {
	segments := DevCenterEnvironmentTypeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DevCenterEnvironmentTypeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
