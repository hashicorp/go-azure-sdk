package dbservers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DbServerId{}

func TestNewDbServerID(t *testing.T) {
	id := NewDbServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudexadatainfrastructurename", "dbserverocid")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.CloudExadataInfrastructureName != "cloudexadatainfrastructurename" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudExadataInfrastructureName'", id.CloudExadataInfrastructureName, "cloudexadatainfrastructurename")
	}

	if id.DbServerName != "dbserverocid" {
		t.Fatalf("Expected %q but got %q for Segment 'DbServerName'", id.DbServerName, "dbserverocid")
	}
}

func TestFormatDbServerID(t *testing.T) {
	actual := NewDbServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudexadatainfrastructurename", "dbserverocid").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures/cloudexadatainfrastructurename/dbServers/dbserverocid"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDbServerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DbServerId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures/cloudexadatainfrastructurename",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures/cloudexadatainfrastructurename/dbServers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures/cloudexadatainfrastructurename/dbServers/dbserverocid",
			Expected: &DbServerId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "example-resource-group",
				CloudExadataInfrastructureName: "cloudexadatainfrastructurename",
				DbServerName:                   "dbserverocid",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures/cloudexadatainfrastructurename/dbServers/dbserverocid/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDbServerID(v.Input)
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

		if actual.CloudExadataInfrastructureName != v.Expected.CloudExadataInfrastructureName {
			t.Fatalf("Expected %q but got %q for CloudExadataInfrastructureName", v.Expected.CloudExadataInfrastructureName, actual.CloudExadataInfrastructureName)
		}

		if actual.DbServerName != v.Expected.DbServerName {
			t.Fatalf("Expected %q but got %q for DbServerName", v.Expected.DbServerName, actual.DbServerName)
		}

	}
}

func TestParseDbServerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DbServerId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/cLoUdExAdAtAiNfRaStRuCtUrEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures/cloudexadatainfrastructurename",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/cLoUdExAdAtAiNfRaStRuCtUrEs/cLoUdExAdAtAiNfRaStRuCtUrEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures/cloudexadatainfrastructurename/dbServers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/cLoUdExAdAtAiNfRaStRuCtUrEs/cLoUdExAdAtAiNfRaStRuCtUrEnAmE/dBsErVeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures/cloudexadatainfrastructurename/dbServers/dbserverocid",
			Expected: &DbServerId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "example-resource-group",
				CloudExadataInfrastructureName: "cloudexadatainfrastructurename",
				DbServerName:                   "dbserverocid",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/cloudExadataInfrastructures/cloudexadatainfrastructurename/dbServers/dbserverocid/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/cLoUdExAdAtAiNfRaStRuCtUrEs/cLoUdExAdAtAiNfRaStRuCtUrEnAmE/dBsErVeRs/dBsErVeRoCiD",
			Expected: &DbServerId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "eXaMpLe-rEsOuRcE-GrOuP",
				CloudExadataInfrastructureName: "cLoUdExAdAtAiNfRaStRuCtUrEnAmE",
				DbServerName:                   "dBsErVeRoCiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/cLoUdExAdAtAiNfRaStRuCtUrEs/cLoUdExAdAtAiNfRaStRuCtUrEnAmE/dBsErVeRs/dBsErVeRoCiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDbServerIDInsensitively(v.Input)
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

		if actual.CloudExadataInfrastructureName != v.Expected.CloudExadataInfrastructureName {
			t.Fatalf("Expected %q but got %q for CloudExadataInfrastructureName", v.Expected.CloudExadataInfrastructureName, actual.CloudExadataInfrastructureName)
		}

		if actual.DbServerName != v.Expected.DbServerName {
			t.Fatalf("Expected %q but got %q for DbServerName", v.Expected.DbServerName, actual.DbServerName)
		}

	}
}

func TestSegmentsForDbServerId(t *testing.T) {
	segments := DbServerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DbServerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
