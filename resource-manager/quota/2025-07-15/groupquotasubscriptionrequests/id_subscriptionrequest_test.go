package groupquotasubscriptionrequests

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SubscriptionRequestId{}

func TestNewSubscriptionRequestID(t *testing.T) {
	id := NewSubscriptionRequestID("managementGroupId", "groupQuotaName", "requestId")

	if id.ManagementGroupId != "managementGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupId'", id.ManagementGroupId, "managementGroupId")
	}

	if id.GroupQuotaName != "groupQuotaName" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupQuotaName'", id.GroupQuotaName, "groupQuotaName")
	}

	if id.RequestId != "requestId" {
		t.Fatalf("Expected %q but got %q for Segment 'RequestId'", id.RequestId, "requestId")
	}
}

func TestFormatSubscriptionRequestID(t *testing.T) {
	actual := NewSubscriptionRequestID("managementGroupId", "groupQuotaName", "requestId").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/subscriptionRequests/requestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSubscriptionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionRequestId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/subscriptionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/subscriptionRequests/requestId",
			Expected: &SubscriptionRequestId{
				ManagementGroupId: "managementGroupId",
				GroupQuotaName:    "groupQuotaName",
				RequestId:         "requestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/subscriptionRequests/requestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionRequestID(v.Input)
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

		if actual.GroupQuotaName != v.Expected.GroupQuotaName {
			t.Fatalf("Expected %q but got %q for GroupQuotaName", v.Expected.GroupQuotaName, actual.GroupQuotaName)
		}

		if actual.RequestId != v.Expected.RequestId {
			t.Fatalf("Expected %q but got %q for RequestId", v.Expected.RequestId, actual.RequestId)
		}

	}
}

func TestParseSubscriptionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionRequestId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/subscriptionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/sUbScRiPtIoNrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/subscriptionRequests/requestId",
			Expected: &SubscriptionRequestId{
				ManagementGroupId: "managementGroupId",
				GroupQuotaName:    "groupQuotaName",
				RequestId:         "requestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/subscriptionRequests/requestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/sUbScRiPtIoNrEqUeStS/rEqUeStId",
			Expected: &SubscriptionRequestId{
				ManagementGroupId: "mAnAgEmEnTgRoUpId",
				GroupQuotaName:    "gRoUpQuOtAnAmE",
				RequestId:         "rEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/sUbScRiPtIoNrEqUeStS/rEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionRequestIDInsensitively(v.Input)
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

		if actual.GroupQuotaName != v.Expected.GroupQuotaName {
			t.Fatalf("Expected %q but got %q for GroupQuotaName", v.Expected.GroupQuotaName, actual.GroupQuotaName)
		}

		if actual.RequestId != v.Expected.RequestId {
			t.Fatalf("Expected %q but got %q for RequestId", v.Expected.RequestId, actual.RequestId)
		}

	}
}

func TestSegmentsForSubscriptionRequestId(t *testing.T) {
	segments := SubscriptionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SubscriptionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
