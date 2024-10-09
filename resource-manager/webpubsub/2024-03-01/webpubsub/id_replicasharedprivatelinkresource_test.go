package webpubsub

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReplicaSharedPrivateLinkResourceId{}

func TestNewReplicaSharedPrivateLinkResourceID(t *testing.T) {
	id := NewReplicaSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "webPubSubName", "replicaName", "sharedPrivateLinkResourceName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.WebPubSubName != "webPubSubName" {
		t.Fatalf("Expected %q but got %q for Segment 'WebPubSubName'", id.WebPubSubName, "webPubSubName")
	}

	if id.ReplicaName != "replicaName" {
		t.Fatalf("Expected %q but got %q for Segment 'ReplicaName'", id.ReplicaName, "replicaName")
	}

	if id.SharedPrivateLinkResourceName != "sharedPrivateLinkResourceName" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedPrivateLinkResourceName'", id.SharedPrivateLinkResourceName, "sharedPrivateLinkResourceName")
	}
}

func TestFormatReplicaSharedPrivateLinkResourceID(t *testing.T) {
	actual := NewReplicaSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "webPubSubName", "replicaName", "sharedPrivateLinkResourceName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas/replicaName/sharedPrivateLinkResources/sharedPrivateLinkResourceName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReplicaSharedPrivateLinkResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReplicaSharedPrivateLinkResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas/replicaName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas/replicaName/sharedPrivateLinkResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas/replicaName/sharedPrivateLinkResources/sharedPrivateLinkResourceName",
			Expected: &ReplicaSharedPrivateLinkResourceId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				WebPubSubName:                 "webPubSubName",
				ReplicaName:                   "replicaName",
				SharedPrivateLinkResourceName: "sharedPrivateLinkResourceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas/replicaName/sharedPrivateLinkResources/sharedPrivateLinkResourceName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReplicaSharedPrivateLinkResourceID(v.Input)
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

		if actual.WebPubSubName != v.Expected.WebPubSubName {
			t.Fatalf("Expected %q but got %q for WebPubSubName", v.Expected.WebPubSubName, actual.WebPubSubName)
		}

		if actual.ReplicaName != v.Expected.ReplicaName {
			t.Fatalf("Expected %q but got %q for ReplicaName", v.Expected.ReplicaName, actual.ReplicaName)
		}

		if actual.SharedPrivateLinkResourceName != v.Expected.SharedPrivateLinkResourceName {
			t.Fatalf("Expected %q but got %q for SharedPrivateLinkResourceName", v.Expected.SharedPrivateLinkResourceName, actual.SharedPrivateLinkResourceName)
		}

	}
}

func TestParseReplicaSharedPrivateLinkResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReplicaSharedPrivateLinkResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sIgNaLrSeRvIcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sIgNaLrSeRvIcE/wEbPuBsUb",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sIgNaLrSeRvIcE/wEbPuBsUb/wEbPuBsUbNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sIgNaLrSeRvIcE/wEbPuBsUb/wEbPuBsUbNaMe/rEpLiCaS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas/replicaName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sIgNaLrSeRvIcE/wEbPuBsUb/wEbPuBsUbNaMe/rEpLiCaS/rEpLiCaNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas/replicaName/sharedPrivateLinkResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sIgNaLrSeRvIcE/wEbPuBsUb/wEbPuBsUbNaMe/rEpLiCaS/rEpLiCaNaMe/sHaReDpRiVaTeLiNkReSoUrCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas/replicaName/sharedPrivateLinkResources/sharedPrivateLinkResourceName",
			Expected: &ReplicaSharedPrivateLinkResourceId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				WebPubSubName:                 "webPubSubName",
				ReplicaName:                   "replicaName",
				SharedPrivateLinkResourceName: "sharedPrivateLinkResourceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SignalRService/webPubSub/webPubSubName/replicas/replicaName/sharedPrivateLinkResources/sharedPrivateLinkResourceName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sIgNaLrSeRvIcE/wEbPuBsUb/wEbPuBsUbNaMe/rEpLiCaS/rEpLiCaNaMe/sHaReDpRiVaTeLiNkReSoUrCeS/sHaReDpRiVaTeLiNkReSoUrCeNaMe",
			Expected: &ReplicaSharedPrivateLinkResourceId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "eXaMpLe-rEsOuRcE-GrOuP",
				WebPubSubName:                 "wEbPuBsUbNaMe",
				ReplicaName:                   "rEpLiCaNaMe",
				SharedPrivateLinkResourceName: "sHaReDpRiVaTeLiNkReSoUrCeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sIgNaLrSeRvIcE/wEbPuBsUb/wEbPuBsUbNaMe/rEpLiCaS/rEpLiCaNaMe/sHaReDpRiVaTeLiNkReSoUrCeS/sHaReDpRiVaTeLiNkReSoUrCeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReplicaSharedPrivateLinkResourceIDInsensitively(v.Input)
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

		if actual.WebPubSubName != v.Expected.WebPubSubName {
			t.Fatalf("Expected %q but got %q for WebPubSubName", v.Expected.WebPubSubName, actual.WebPubSubName)
		}

		if actual.ReplicaName != v.Expected.ReplicaName {
			t.Fatalf("Expected %q but got %q for ReplicaName", v.Expected.ReplicaName, actual.ReplicaName)
		}

		if actual.SharedPrivateLinkResourceName != v.Expected.SharedPrivateLinkResourceName {
			t.Fatalf("Expected %q but got %q for SharedPrivateLinkResourceName", v.Expected.SharedPrivateLinkResourceName, actual.SharedPrivateLinkResourceName)
		}

	}
}

func TestSegmentsForReplicaSharedPrivateLinkResourceId(t *testing.T) {
	segments := ReplicaSharedPrivateLinkResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReplicaSharedPrivateLinkResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
