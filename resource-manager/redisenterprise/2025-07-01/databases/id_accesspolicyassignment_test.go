package databases

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AccessPolicyAssignmentId{}

func TestNewAccessPolicyAssignmentID(t *testing.T) {
	id := NewAccessPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "redisEnterpriseName", "databaseName", "accessPolicyAssignmentName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.RedisEnterpriseName != "redisEnterpriseName" {
		t.Fatalf("Expected %q but got %q for Segment 'RedisEnterpriseName'", id.RedisEnterpriseName, "redisEnterpriseName")
	}

	if id.DatabaseName != "databaseName" {
		t.Fatalf("Expected %q but got %q for Segment 'DatabaseName'", id.DatabaseName, "databaseName")
	}

	if id.AccessPolicyAssignmentName != "accessPolicyAssignmentName" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPolicyAssignmentName'", id.AccessPolicyAssignmentName, "accessPolicyAssignmentName")
	}
}

func TestFormatAccessPolicyAssignmentID(t *testing.T) {
	actual := NewAccessPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "redisEnterpriseName", "databaseName", "accessPolicyAssignmentName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases/databaseName/accessPolicyAssignments/accessPolicyAssignmentName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAccessPolicyAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AccessPolicyAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases/databaseName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases/databaseName/accessPolicyAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases/databaseName/accessPolicyAssignments/accessPolicyAssignmentName",
			Expected: &AccessPolicyAssignmentId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				RedisEnterpriseName:        "redisEnterpriseName",
				DatabaseName:               "databaseName",
				AccessPolicyAssignmentName: "accessPolicyAssignmentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases/databaseName/accessPolicyAssignments/accessPolicyAssignmentName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAccessPolicyAssignmentID(v.Input)
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

		if actual.RedisEnterpriseName != v.Expected.RedisEnterpriseName {
			t.Fatalf("Expected %q but got %q for RedisEnterpriseName", v.Expected.RedisEnterpriseName, actual.RedisEnterpriseName)
		}

		if actual.DatabaseName != v.Expected.DatabaseName {
			t.Fatalf("Expected %q but got %q for DatabaseName", v.Expected.DatabaseName, actual.DatabaseName)
		}

		if actual.AccessPolicyAssignmentName != v.Expected.AccessPolicyAssignmentName {
			t.Fatalf("Expected %q but got %q for AccessPolicyAssignmentName", v.Expected.AccessPolicyAssignmentName, actual.AccessPolicyAssignmentName)
		}

	}
}

func TestParseAccessPolicyAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AccessPolicyAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cAcHe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cAcHe/rEdIsEnTeRpRiSe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cAcHe/rEdIsEnTeRpRiSe/rEdIsEnTeRpRiSeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cAcHe/rEdIsEnTeRpRiSe/rEdIsEnTeRpRiSeNaMe/dAtAbAsEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases/databaseName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cAcHe/rEdIsEnTeRpRiSe/rEdIsEnTeRpRiSeNaMe/dAtAbAsEs/dAtAbAsEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases/databaseName/accessPolicyAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cAcHe/rEdIsEnTeRpRiSe/rEdIsEnTeRpRiSeNaMe/dAtAbAsEs/dAtAbAsEnAmE/aCcEsSpOlIcYaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases/databaseName/accessPolicyAssignments/accessPolicyAssignmentName",
			Expected: &AccessPolicyAssignmentId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				RedisEnterpriseName:        "redisEnterpriseName",
				DatabaseName:               "databaseName",
				AccessPolicyAssignmentName: "accessPolicyAssignmentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Cache/redisEnterprise/redisEnterpriseName/databases/databaseName/accessPolicyAssignments/accessPolicyAssignmentName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cAcHe/rEdIsEnTeRpRiSe/rEdIsEnTeRpRiSeNaMe/dAtAbAsEs/dAtAbAsEnAmE/aCcEsSpOlIcYaSsIgNmEnTs/aCcEsSpOlIcYaSsIgNmEnTnAmE",
			Expected: &AccessPolicyAssignmentId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "eXaMpLe-rEsOuRcE-GrOuP",
				RedisEnterpriseName:        "rEdIsEnTeRpRiSeNaMe",
				DatabaseName:               "dAtAbAsEnAmE",
				AccessPolicyAssignmentName: "aCcEsSpOlIcYaSsIgNmEnTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cAcHe/rEdIsEnTeRpRiSe/rEdIsEnTeRpRiSeNaMe/dAtAbAsEs/dAtAbAsEnAmE/aCcEsSpOlIcYaSsIgNmEnTs/aCcEsSpOlIcYaSsIgNmEnTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAccessPolicyAssignmentIDInsensitively(v.Input)
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

		if actual.RedisEnterpriseName != v.Expected.RedisEnterpriseName {
			t.Fatalf("Expected %q but got %q for RedisEnterpriseName", v.Expected.RedisEnterpriseName, actual.RedisEnterpriseName)
		}

		if actual.DatabaseName != v.Expected.DatabaseName {
			t.Fatalf("Expected %q but got %q for DatabaseName", v.Expected.DatabaseName, actual.DatabaseName)
		}

		if actual.AccessPolicyAssignmentName != v.Expected.AccessPolicyAssignmentName {
			t.Fatalf("Expected %q but got %q for AccessPolicyAssignmentName", v.Expected.AccessPolicyAssignmentName, actual.AccessPolicyAssignmentName)
		}

	}
}

func TestSegmentsForAccessPolicyAssignmentId(t *testing.T) {
	segments := AccessPolicyAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AccessPolicyAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
