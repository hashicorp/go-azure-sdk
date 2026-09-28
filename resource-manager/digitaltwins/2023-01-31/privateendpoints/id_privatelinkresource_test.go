package privateendpoints

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PrivateLinkResourceId{}

func TestNewPrivateLinkResourceID(t *testing.T) {
	id := NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "digitalTwinsInstanceName", "resourceId")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.DigitalTwinsInstanceName != "digitalTwinsInstanceName" {
		t.Fatalf("Expected %q but got %q for Segment 'DigitalTwinsInstanceName'", id.DigitalTwinsInstanceName, "digitalTwinsInstanceName")
	}

	if id.ResourceId != "resourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceId'", id.ResourceId, "resourceId")
	}
}

func TestFormatPrivateLinkResourceID(t *testing.T) {
	actual := NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "digitalTwinsInstanceName", "resourceId").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances/digitalTwinsInstanceName/privateLinkResources/resourceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePrivateLinkResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PrivateLinkResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances/digitalTwinsInstanceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances/digitalTwinsInstanceName/privateLinkResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances/digitalTwinsInstanceName/privateLinkResources/resourceId",
			Expected: &PrivateLinkResourceId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "example-resource-group",
				DigitalTwinsInstanceName: "digitalTwinsInstanceName",
				ResourceId:               "resourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances/digitalTwinsInstanceName/privateLinkResources/resourceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePrivateLinkResourceID(v.Input)
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

		if actual.DigitalTwinsInstanceName != v.Expected.DigitalTwinsInstanceName {
			t.Fatalf("Expected %q but got %q for DigitalTwinsInstanceName", v.Expected.DigitalTwinsInstanceName, actual.DigitalTwinsInstanceName)
		}

		if actual.ResourceId != v.Expected.ResourceId {
			t.Fatalf("Expected %q but got %q for ResourceId", v.Expected.ResourceId, actual.ResourceId)
		}

	}
}

func TestParsePrivateLinkResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PrivateLinkResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIgItAlTwInS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIgItAlTwInS/dIgItAlTwInSiNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances/digitalTwinsInstanceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIgItAlTwInS/dIgItAlTwInSiNsTaNcEs/dIgItAlTwInSiNsTaNcEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances/digitalTwinsInstanceName/privateLinkResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIgItAlTwInS/dIgItAlTwInSiNsTaNcEs/dIgItAlTwInSiNsTaNcEnAmE/pRiVaTeLiNkReSoUrCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances/digitalTwinsInstanceName/privateLinkResources/resourceId",
			Expected: &PrivateLinkResourceId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "example-resource-group",
				DigitalTwinsInstanceName: "digitalTwinsInstanceName",
				ResourceId:               "resourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DigitalTwins/digitalTwinsInstances/digitalTwinsInstanceName/privateLinkResources/resourceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIgItAlTwInS/dIgItAlTwInSiNsTaNcEs/dIgItAlTwInSiNsTaNcEnAmE/pRiVaTeLiNkReSoUrCeS/rEsOuRcEiD",
			Expected: &PrivateLinkResourceId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "eXaMpLe-rEsOuRcE-GrOuP",
				DigitalTwinsInstanceName: "dIgItAlTwInSiNsTaNcEnAmE",
				ResourceId:               "rEsOuRcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIgItAlTwInS/dIgItAlTwInSiNsTaNcEs/dIgItAlTwInSiNsTaNcEnAmE/pRiVaTeLiNkReSoUrCeS/rEsOuRcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePrivateLinkResourceIDInsensitively(v.Input)
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

		if actual.DigitalTwinsInstanceName != v.Expected.DigitalTwinsInstanceName {
			t.Fatalf("Expected %q but got %q for DigitalTwinsInstanceName", v.Expected.DigitalTwinsInstanceName, actual.DigitalTwinsInstanceName)
		}

		if actual.ResourceId != v.Expected.ResourceId {
			t.Fatalf("Expected %q but got %q for ResourceId", v.Expected.ResourceId, actual.ResourceId)
		}

	}
}

func TestSegmentsForPrivateLinkResourceId(t *testing.T) {
	segments := PrivateLinkResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PrivateLinkResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %d unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
