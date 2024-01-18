package longtermretentionbackup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LtrBackupOperationId{}

func TestNewLtrBackupOperationID(t *testing.T) {
	id := NewLtrBackupOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerValue", "ltrBackupOperationValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.FlexibleServerName != "flexibleServerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FlexibleServerName'", id.FlexibleServerName, "flexibleServerValue")
	}

	if id.LtrBackupOperationName != "ltrBackupOperationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LtrBackupOperationName'", id.LtrBackupOperationName, "ltrBackupOperationValue")
	}
}

func TestFormatLtrBackupOperationID(t *testing.T) {
	actual := NewLtrBackupOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerValue", "ltrBackupOperationValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServerValue/ltrBackupOperations/ltrBackupOperationValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLtrBackupOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LtrBackupOperationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServerValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServerValue/ltrBackupOperations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServerValue/ltrBackupOperations/ltrBackupOperationValue",
			Expected: &LtrBackupOperationId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				FlexibleServerName:     "flexibleServerValue",
				LtrBackupOperationName: "ltrBackupOperationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServerValue/ltrBackupOperations/ltrBackupOperationValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLtrBackupOperationID(v.Input)
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

		if actual.FlexibleServerName != v.Expected.FlexibleServerName {
			t.Fatalf("Expected %q but got %q for FlexibleServerName", v.Expected.FlexibleServerName, actual.FlexibleServerName)
		}

		if actual.LtrBackupOperationName != v.Expected.LtrBackupOperationName {
			t.Fatalf("Expected %q but got %q for LtrBackupOperationName", v.Expected.LtrBackupOperationName, actual.LtrBackupOperationName)
		}

	}
}

func TestParseLtrBackupOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LtrBackupOperationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/fLeXiBlEsErVeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServerValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/fLeXiBlEsErVeRs/fLeXiBlEsErVeRvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServerValue/ltrBackupOperations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/fLeXiBlEsErVeRs/fLeXiBlEsErVeRvAlUe/lTrBaCkUpOpErAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServerValue/ltrBackupOperations/ltrBackupOperationValue",
			Expected: &LtrBackupOperationId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				FlexibleServerName:     "flexibleServerValue",
				LtrBackupOperationName: "ltrBackupOperationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServerValue/ltrBackupOperations/ltrBackupOperationValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/fLeXiBlEsErVeRs/fLeXiBlEsErVeRvAlUe/lTrBaCkUpOpErAtIoNs/lTrBaCkUpOpErAtIoNvAlUe",
			Expected: &LtrBackupOperationId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "eXaMpLe-rEsOuRcE-GrOuP",
				FlexibleServerName:     "fLeXiBlEsErVeRvAlUe",
				LtrBackupOperationName: "lTrBaCkUpOpErAtIoNvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/fLeXiBlEsErVeRs/fLeXiBlEsErVeRvAlUe/lTrBaCkUpOpErAtIoNs/lTrBaCkUpOpErAtIoNvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLtrBackupOperationIDInsensitively(v.Input)
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

		if actual.FlexibleServerName != v.Expected.FlexibleServerName {
			t.Fatalf("Expected %q but got %q for FlexibleServerName", v.Expected.FlexibleServerName, actual.FlexibleServerName)
		}

		if actual.LtrBackupOperationName != v.Expected.LtrBackupOperationName {
			t.Fatalf("Expected %q but got %q for LtrBackupOperationName", v.Expected.LtrBackupOperationName, actual.LtrBackupOperationName)
		}

	}
}

func TestSegmentsForLtrBackupOperationId(t *testing.T) {
	segments := LtrBackupOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LtrBackupOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
