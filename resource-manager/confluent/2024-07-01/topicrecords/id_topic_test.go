package topicrecords

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TopicId{}

func TestNewTopicID(t *testing.T) {
	id := NewTopicID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationName", "environmentId", "clusterId", "topicName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.OrganizationName != "organizationName" {
		t.Fatalf("Expected %q but got %q for Segment 'OrganizationName'", id.OrganizationName, "organizationName")
	}

	if id.EnvironmentId != "environmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'EnvironmentId'", id.EnvironmentId, "environmentId")
	}

	if id.ClusterId != "clusterId" {
		t.Fatalf("Expected %q but got %q for Segment 'ClusterId'", id.ClusterId, "clusterId")
	}

	if id.TopicName != "topicName" {
		t.Fatalf("Expected %q but got %q for Segment 'TopicName'", id.TopicName, "topicName")
	}
}

func TestFormatTopicID(t *testing.T) {
	actual := NewTopicID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationName", "environmentId", "clusterId", "topicName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters/clusterId/topics/topicName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTopicID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TopicId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters/clusterId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters/clusterId/topics",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters/clusterId/topics/topicName",
			Expected: &TopicId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				OrganizationName:  "organizationName",
				EnvironmentId:     "environmentId",
				ClusterId:         "clusterId",
				TopicName:         "topicName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters/clusterId/topics/topicName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTopicID(v.Input)
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

		if actual.OrganizationName != v.Expected.OrganizationName {
			t.Fatalf("Expected %q but got %q for OrganizationName", v.Expected.OrganizationName, actual.OrganizationName)
		}

		if actual.EnvironmentId != v.Expected.EnvironmentId {
			t.Fatalf("Expected %q but got %q for EnvironmentId", v.Expected.EnvironmentId, actual.EnvironmentId)
		}

		if actual.ClusterId != v.Expected.ClusterId {
			t.Fatalf("Expected %q but got %q for ClusterId", v.Expected.ClusterId, actual.ClusterId)
		}

		if actual.TopicName != v.Expected.TopicName {
			t.Fatalf("Expected %q but got %q for TopicName", v.Expected.TopicName, actual.TopicName)
		}

	}
}

func TestParseTopicIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TopicId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt/oRgAnIzAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt/oRgAnIzAtIoNs/oRgAnIzAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt/oRgAnIzAtIoNs/oRgAnIzAtIoNnAmE/eNvIrOnMeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt/oRgAnIzAtIoNs/oRgAnIzAtIoNnAmE/eNvIrOnMeNtS/eNvIrOnMeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt/oRgAnIzAtIoNs/oRgAnIzAtIoNnAmE/eNvIrOnMeNtS/eNvIrOnMeNtId/cLuStErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters/clusterId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt/oRgAnIzAtIoNs/oRgAnIzAtIoNnAmE/eNvIrOnMeNtS/eNvIrOnMeNtId/cLuStErS/cLuStErId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters/clusterId/topics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt/oRgAnIzAtIoNs/oRgAnIzAtIoNnAmE/eNvIrOnMeNtS/eNvIrOnMeNtId/cLuStErS/cLuStErId/tOpIcS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters/clusterId/topics/topicName",
			Expected: &TopicId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				OrganizationName:  "organizationName",
				EnvironmentId:     "environmentId",
				ClusterId:         "clusterId",
				TopicName:         "topicName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Confluent/organizations/organizationName/environments/environmentId/clusters/clusterId/topics/topicName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt/oRgAnIzAtIoNs/oRgAnIzAtIoNnAmE/eNvIrOnMeNtS/eNvIrOnMeNtId/cLuStErS/cLuStErId/tOpIcS/tOpIcNaMe",
			Expected: &TopicId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "eXaMpLe-rEsOuRcE-GrOuP",
				OrganizationName:  "oRgAnIzAtIoNnAmE",
				EnvironmentId:     "eNvIrOnMeNtId",
				ClusterId:         "cLuStErId",
				TopicName:         "tOpIcNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnFlUeNt/oRgAnIzAtIoNs/oRgAnIzAtIoNnAmE/eNvIrOnMeNtS/eNvIrOnMeNtId/cLuStErS/cLuStErId/tOpIcS/tOpIcNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTopicIDInsensitively(v.Input)
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

		if actual.OrganizationName != v.Expected.OrganizationName {
			t.Fatalf("Expected %q but got %q for OrganizationName", v.Expected.OrganizationName, actual.OrganizationName)
		}

		if actual.EnvironmentId != v.Expected.EnvironmentId {
			t.Fatalf("Expected %q but got %q for EnvironmentId", v.Expected.EnvironmentId, actual.EnvironmentId)
		}

		if actual.ClusterId != v.Expected.ClusterId {
			t.Fatalf("Expected %q but got %q for ClusterId", v.Expected.ClusterId, actual.ClusterId)
		}

		if actual.TopicName != v.Expected.TopicName {
			t.Fatalf("Expected %q but got %q for TopicName", v.Expected.TopicName, actual.TopicName)
		}

	}
}

func TestSegmentsForTopicId(t *testing.T) {
	segments := TopicId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TopicId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
