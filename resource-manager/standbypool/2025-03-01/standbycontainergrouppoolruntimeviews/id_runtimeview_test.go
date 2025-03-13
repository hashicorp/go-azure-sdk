package standbycontainergrouppoolruntimeviews

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RuntimeViewId{}

func TestNewRuntimeViewID(t *testing.T) {
	id := NewRuntimeViewID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyContainerGroupPoolName", "runtimeViewName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.StandbyContainerGroupPoolName != "standbyContainerGroupPoolName" {
		t.Fatalf("Expected %q but got %q for Segment 'StandbyContainerGroupPoolName'", id.StandbyContainerGroupPoolName, "standbyContainerGroupPoolName")
	}

	if id.RuntimeViewName != "runtimeViewName" {
		t.Fatalf("Expected %q but got %q for Segment 'RuntimeViewName'", id.RuntimeViewName, "runtimeViewName")
	}
}

func TestFormatRuntimeViewID(t *testing.T) {
	actual := NewRuntimeViewID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyContainerGroupPoolName", "runtimeViewName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools/standbyContainerGroupPoolName/runtimeViews/runtimeViewName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRuntimeViewID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RuntimeViewId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools/standbyContainerGroupPoolName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools/standbyContainerGroupPoolName/runtimeViews",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools/standbyContainerGroupPoolName/runtimeViews/runtimeViewName",
			Expected: &RuntimeViewId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				StandbyContainerGroupPoolName: "standbyContainerGroupPoolName",
				RuntimeViewName:               "runtimeViewName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools/standbyContainerGroupPoolName/runtimeViews/runtimeViewName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRuntimeViewID(v.Input)
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

		if actual.StandbyContainerGroupPoolName != v.Expected.StandbyContainerGroupPoolName {
			t.Fatalf("Expected %q but got %q for StandbyContainerGroupPoolName", v.Expected.StandbyContainerGroupPoolName, actual.StandbyContainerGroupPoolName)
		}

		if actual.RuntimeViewName != v.Expected.RuntimeViewName {
			t.Fatalf("Expected %q but got %q for RuntimeViewName", v.Expected.RuntimeViewName, actual.RuntimeViewName)
		}

	}
}

func TestParseRuntimeViewIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RuntimeViewId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByCoNtAiNeRgRoUpPoOlS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools/standbyContainerGroupPoolName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByCoNtAiNeRgRoUpPoOlS/sTaNdByCoNtAiNeRgRoUpPoOlNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools/standbyContainerGroupPoolName/runtimeViews",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByCoNtAiNeRgRoUpPoOlS/sTaNdByCoNtAiNeRgRoUpPoOlNaMe/rUnTiMeViEwS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools/standbyContainerGroupPoolName/runtimeViews/runtimeViewName",
			Expected: &RuntimeViewId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				StandbyContainerGroupPoolName: "standbyContainerGroupPoolName",
				RuntimeViewName:               "runtimeViewName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyContainerGroupPools/standbyContainerGroupPoolName/runtimeViews/runtimeViewName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByCoNtAiNeRgRoUpPoOlS/sTaNdByCoNtAiNeRgRoUpPoOlNaMe/rUnTiMeViEwS/rUnTiMeViEwNaMe",
			Expected: &RuntimeViewId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "eXaMpLe-rEsOuRcE-GrOuP",
				StandbyContainerGroupPoolName: "sTaNdByCoNtAiNeRgRoUpPoOlNaMe",
				RuntimeViewName:               "rUnTiMeViEwNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByCoNtAiNeRgRoUpPoOlS/sTaNdByCoNtAiNeRgRoUpPoOlNaMe/rUnTiMeViEwS/rUnTiMeViEwNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRuntimeViewIDInsensitively(v.Input)
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

		if actual.StandbyContainerGroupPoolName != v.Expected.StandbyContainerGroupPoolName {
			t.Fatalf("Expected %q but got %q for StandbyContainerGroupPoolName", v.Expected.StandbyContainerGroupPoolName, actual.StandbyContainerGroupPoolName)
		}

		if actual.RuntimeViewName != v.Expected.RuntimeViewName {
			t.Fatalf("Expected %q but got %q for RuntimeViewName", v.Expected.RuntimeViewName, actual.RuntimeViewName)
		}

	}
}

func TestSegmentsForRuntimeViewId(t *testing.T) {
	segments := RuntimeViewId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RuntimeViewId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
