package datacollectionruleassociations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedDataCollectionRuleAssociationId{}

func TestNewScopedDataCollectionRuleAssociationID(t *testing.T) {
	id := NewScopedDataCollectionRuleAssociationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "dataCollectionRuleAssociationName")

	if id.ResourceUri != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceUri'", id.ResourceUri, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.DataCollectionRuleAssociationName != "dataCollectionRuleAssociationName" {
		t.Fatalf("Expected %q but got %q for Segment 'DataCollectionRuleAssociationName'", id.DataCollectionRuleAssociationName, "dataCollectionRuleAssociationName")
	}
}

func TestFormatScopedDataCollectionRuleAssociationID(t *testing.T) {
	actual := NewScopedDataCollectionRuleAssociationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "dataCollectionRuleAssociationName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/dataCollectionRuleAssociations/dataCollectionRuleAssociationName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedDataCollectionRuleAssociationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedDataCollectionRuleAssociationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/dataCollectionRuleAssociations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/dataCollectionRuleAssociations/dataCollectionRuleAssociationName",
			Expected: &ScopedDataCollectionRuleAssociationId{
				ResourceUri:                       "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				DataCollectionRuleAssociationName: "dataCollectionRuleAssociationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/dataCollectionRuleAssociations/dataCollectionRuleAssociationName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedDataCollectionRuleAssociationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceUri != v.Expected.ResourceUri {
			t.Fatalf("Expected %q but got %q for ResourceUri", v.Expected.ResourceUri, actual.ResourceUri)
		}

		if actual.DataCollectionRuleAssociationName != v.Expected.DataCollectionRuleAssociationName {
			t.Fatalf("Expected %q but got %q for DataCollectionRuleAssociationName", v.Expected.DataCollectionRuleAssociationName, actual.DataCollectionRuleAssociationName)
		}

	}
}

func TestParseScopedDataCollectionRuleAssociationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedDataCollectionRuleAssociationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.iNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/dataCollectionRuleAssociations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.iNsIgHtS/dAtAcOlLeCtIoNrUlEaSsOcIaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/dataCollectionRuleAssociations/dataCollectionRuleAssociationName",
			Expected: &ScopedDataCollectionRuleAssociationId{
				ResourceUri:                       "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				DataCollectionRuleAssociationName: "dataCollectionRuleAssociationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/dataCollectionRuleAssociations/dataCollectionRuleAssociationName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.iNsIgHtS/dAtAcOlLeCtIoNrUlEaSsOcIaTiOnS/dAtAcOlLeCtIoNrUlEaSsOcIaTiOnNaMe",
			Expected: &ScopedDataCollectionRuleAssociationId{
				ResourceUri:                       "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				DataCollectionRuleAssociationName: "dAtAcOlLeCtIoNrUlEaSsOcIaTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.iNsIgHtS/dAtAcOlLeCtIoNrUlEaSsOcIaTiOnS/dAtAcOlLeCtIoNrUlEaSsOcIaTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedDataCollectionRuleAssociationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceUri != v.Expected.ResourceUri {
			t.Fatalf("Expected %q but got %q for ResourceUri", v.Expected.ResourceUri, actual.ResourceUri)
		}

		if actual.DataCollectionRuleAssociationName != v.Expected.DataCollectionRuleAssociationName {
			t.Fatalf("Expected %q but got %q for DataCollectionRuleAssociationName", v.Expected.DataCollectionRuleAssociationName, actual.DataCollectionRuleAssociationName)
		}

	}
}

func TestSegmentsForScopedDataCollectionRuleAssociationId(t *testing.T) {
	segments := ScopedDataCollectionRuleAssociationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedDataCollectionRuleAssociationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
