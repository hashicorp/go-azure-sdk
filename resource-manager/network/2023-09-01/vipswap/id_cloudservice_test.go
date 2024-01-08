package vipswap

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CloudServiceId{}

func TestNewCloudServiceID(t *testing.T) {
	id := NewCloudServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "cloudServiceValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "resourceGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "resourceGroupValue")
	}

	if id.CloudServiceName != "cloudServiceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudServiceName'", id.CloudServiceName, "cloudServiceValue")
	}
}

func TestFormatCloudServiceID(t *testing.T) {
	actual := NewCloudServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "cloudServiceValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.Compute/cloudServices/cloudServiceValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCloudServiceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CloudServiceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.Compute/cloudServices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.Compute/cloudServices/cloudServiceValue",
			Expected: &CloudServiceId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "resourceGroupValue",
				CloudServiceName:  "cloudServiceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.Compute/cloudServices/cloudServiceValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCloudServiceID(v.Input)
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

		if actual.CloudServiceName != v.Expected.CloudServiceName {
			t.Fatalf("Expected %q but got %q for CloudServiceName", v.Expected.CloudServiceName, actual.CloudServiceName)
		}

	}
}

func TestParseCloudServiceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CloudServiceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.cOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.Compute/cloudServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.Compute/cloudServices/cloudServiceValue",
			Expected: &CloudServiceId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "resourceGroupValue",
				CloudServiceName:  "cloudServiceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.Compute/cloudServices/cloudServiceValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe",
			Expected: &CloudServiceId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "rEsOuRcEgRoUpVaLuE",
				CloudServiceName:  "cLoUdSeRvIcEvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCloudServiceIDInsensitively(v.Input)
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

		if actual.CloudServiceName != v.Expected.CloudServiceName {
			t.Fatalf("Expected %q but got %q for CloudServiceName", v.Expected.CloudServiceName, actual.CloudServiceName)
		}

	}
}

func TestSegmentsForCloudServiceId(t *testing.T) {
	segments := CloudServiceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CloudServiceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
