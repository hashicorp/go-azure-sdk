package webapps

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TriggeredWebJobHistoryId{}

func TestNewTriggeredWebJobHistoryID(t *testing.T) {
	id := NewTriggeredWebJobHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "triggeredWebJobValue", "historyValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.SiteName != "siteValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteName'", id.SiteName, "siteValue")
	}

	if id.SlotName != "slotValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SlotName'", id.SlotName, "slotValue")
	}

	if id.TriggeredWebJobName != "triggeredWebJobValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TriggeredWebJobName'", id.TriggeredWebJobName, "triggeredWebJobValue")
	}

	if id.HistoryName != "historyValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HistoryName'", id.HistoryName, "historyValue")
	}
}

func TestFormatTriggeredWebJobHistoryID(t *testing.T) {
	actual := NewTriggeredWebJobHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "triggeredWebJobValue", "historyValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs/triggeredWebJobValue/history/historyValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTriggeredWebJobHistoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TriggeredWebJobHistoryId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs/triggeredWebJobValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs/triggeredWebJobValue/history",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs/triggeredWebJobValue/history/historyValue",
			Expected: &TriggeredWebJobHistoryId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				SiteName:            "siteValue",
				SlotName:            "slotValue",
				TriggeredWebJobName: "triggeredWebJobValue",
				HistoryName:         "historyValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs/triggeredWebJobValue/history/historyValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTriggeredWebJobHistoryID(v.Input)
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

		if actual.SiteName != v.Expected.SiteName {
			t.Fatalf("Expected %q but got %q for SiteName", v.Expected.SiteName, actual.SiteName)
		}

		if actual.SlotName != v.Expected.SlotName {
			t.Fatalf("Expected %q but got %q for SlotName", v.Expected.SlotName, actual.SlotName)
		}

		if actual.TriggeredWebJobName != v.Expected.TriggeredWebJobName {
			t.Fatalf("Expected %q but got %q for TriggeredWebJobName", v.Expected.TriggeredWebJobName, actual.TriggeredWebJobName)
		}

		if actual.HistoryName != v.Expected.HistoryName {
			t.Fatalf("Expected %q but got %q for HistoryName", v.Expected.HistoryName, actual.HistoryName)
		}

	}
}

func TestParseTriggeredWebJobHistoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TriggeredWebJobHistoryId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sItEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sItEs/sItEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sItEs/sItEvAlUe/sLoTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sItEs/sItEvAlUe/sLoTs/sLoTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sItEs/sItEvAlUe/sLoTs/sLoTvAlUe/tRiGgErEdWeBjObS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs/triggeredWebJobValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sItEs/sItEvAlUe/sLoTs/sLoTvAlUe/tRiGgErEdWeBjObS/tRiGgErEdWeBjObVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs/triggeredWebJobValue/history",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sItEs/sItEvAlUe/sLoTs/sLoTvAlUe/tRiGgErEdWeBjObS/tRiGgErEdWeBjObVaLuE/hIsToRy",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs/triggeredWebJobValue/history/historyValue",
			Expected: &TriggeredWebJobHistoryId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				SiteName:            "siteValue",
				SlotName:            "slotValue",
				TriggeredWebJobName: "triggeredWebJobValue",
				HistoryName:         "historyValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue/slots/slotValue/triggeredWebJobs/triggeredWebJobValue/history/historyValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sItEs/sItEvAlUe/sLoTs/sLoTvAlUe/tRiGgErEdWeBjObS/tRiGgErEdWeBjObVaLuE/hIsToRy/hIsToRyVaLuE",
			Expected: &TriggeredWebJobHistoryId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "eXaMpLe-rEsOuRcE-GrOuP",
				SiteName:            "sItEvAlUe",
				SlotName:            "sLoTvAlUe",
				TriggeredWebJobName: "tRiGgErEdWeBjObVaLuE",
				HistoryName:         "hIsToRyVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sItEs/sItEvAlUe/sLoTs/sLoTvAlUe/tRiGgErEdWeBjObS/tRiGgErEdWeBjObVaLuE/hIsToRy/hIsToRyVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTriggeredWebJobHistoryIDInsensitively(v.Input)
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

		if actual.SiteName != v.Expected.SiteName {
			t.Fatalf("Expected %q but got %q for SiteName", v.Expected.SiteName, actual.SiteName)
		}

		if actual.SlotName != v.Expected.SlotName {
			t.Fatalf("Expected %q but got %q for SlotName", v.Expected.SlotName, actual.SlotName)
		}

		if actual.TriggeredWebJobName != v.Expected.TriggeredWebJobName {
			t.Fatalf("Expected %q but got %q for TriggeredWebJobName", v.Expected.TriggeredWebJobName, actual.TriggeredWebJobName)
		}

		if actual.HistoryName != v.Expected.HistoryName {
			t.Fatalf("Expected %q but got %q for HistoryName", v.Expected.HistoryName, actual.HistoryName)
		}

	}
}

func TestSegmentsForTriggeredWebJobHistoryId(t *testing.T) {
	segments := TriggeredWebJobHistoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TriggeredWebJobHistoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
