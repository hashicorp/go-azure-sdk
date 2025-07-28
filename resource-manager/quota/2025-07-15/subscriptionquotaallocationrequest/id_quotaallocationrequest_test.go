package subscriptionquotaallocationrequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &QuotaAllocationRequestId{}

func TestNewQuotaAllocationRequestID(t *testing.T) {
	id := NewQuotaAllocationRequestID("managementGroupId", "12345678-1234-9876-4563-123456789012", "groupQuotaName", "resourceProviderName", "allocationId")

	if id.ManagementGroupId != "managementGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupId'", id.ManagementGroupId, "managementGroupId")
	}

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.GroupQuotaName != "groupQuotaName" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupQuotaName'", id.GroupQuotaName, "groupQuotaName")
	}

	if id.ResourceProviderName != "resourceProviderName" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceProviderName'", id.ResourceProviderName, "resourceProviderName")
	}

	if id.AllocationId != "allocationId" {
		t.Fatalf("Expected %q but got %q for Segment 'AllocationId'", id.AllocationId, "allocationId")
	}
}

func TestFormatQuotaAllocationRequestID(t *testing.T) {
	actual := NewQuotaAllocationRequestID("managementGroupId", "12345678-1234-9876-4563-123456789012", "groupQuotaName", "resourceProviderName", "allocationId").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/quotaAllocationRequests/allocationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseQuotaAllocationRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *QuotaAllocationRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/quotaAllocationRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/quotaAllocationRequests/allocationId",
			Expected: &QuotaAllocationRequestId{
				ManagementGroupId:    "managementGroupId",
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				GroupQuotaName:       "groupQuotaName",
				ResourceProviderName: "resourceProviderName",
				AllocationId:         "allocationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/quotaAllocationRequests/allocationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseQuotaAllocationRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.GroupQuotaName != v.Expected.GroupQuotaName {
			t.Fatalf("Expected %q but got %q for GroupQuotaName", v.Expected.GroupQuotaName, actual.GroupQuotaName)
		}

		if actual.ResourceProviderName != v.Expected.ResourceProviderName {
			t.Fatalf("Expected %q but got %q for ResourceProviderName", v.Expected.ResourceProviderName, actual.ResourceProviderName)
		}

		if actual.AllocationId != v.Expected.AllocationId {
			t.Fatalf("Expected %q but got %q for AllocationId", v.Expected.AllocationId, actual.AllocationId)
		}

	}
}

func TestParseQuotaAllocationRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *QuotaAllocationRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.qUoTa",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs/rEsOuRcEpRoViDeRnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/quotaAllocationRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs/rEsOuRcEpRoViDeRnAmE/qUoTaAlLoCaTiOnReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/quotaAllocationRequests/allocationId",
			Expected: &QuotaAllocationRequestId{
				ManagementGroupId:    "managementGroupId",
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				GroupQuotaName:       "groupQuotaName",
				ResourceProviderName: "resourceProviderName",
				AllocationId:         "allocationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/quotaAllocationRequests/allocationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs/rEsOuRcEpRoViDeRnAmE/qUoTaAlLoCaTiOnReQuEsTs/aLlOcAtIoNiD",
			Expected: &QuotaAllocationRequestId{
				ManagementGroupId:    "mAnAgEmEnTgRoUpId",
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				GroupQuotaName:       "gRoUpQuOtAnAmE",
				ResourceProviderName: "rEsOuRcEpRoViDeRnAmE",
				AllocationId:         "aLlOcAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs/rEsOuRcEpRoViDeRnAmE/qUoTaAlLoCaTiOnReQuEsTs/aLlOcAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseQuotaAllocationRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.GroupQuotaName != v.Expected.GroupQuotaName {
			t.Fatalf("Expected %q but got %q for GroupQuotaName", v.Expected.GroupQuotaName, actual.GroupQuotaName)
		}

		if actual.ResourceProviderName != v.Expected.ResourceProviderName {
			t.Fatalf("Expected %q but got %q for ResourceProviderName", v.Expected.ResourceProviderName, actual.ResourceProviderName)
		}

		if actual.AllocationId != v.Expected.AllocationId {
			t.Fatalf("Expected %q but got %q for AllocationId", v.Expected.AllocationId, actual.AllocationId)
		}

	}
}

func TestSegmentsForQuotaAllocationRequestId(t *testing.T) {
	segments := QuotaAllocationRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("QuotaAllocationRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
