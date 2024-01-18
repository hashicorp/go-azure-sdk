package deployments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ProviderDeploymentId{}

func TestNewProviderDeploymentID(t *testing.T) {
	id := NewProviderDeploymentID("12345678-1234-9876-4563-123456789012", "deploymentValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.DeploymentName != "deploymentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeploymentName'", id.DeploymentName, "deploymentValue")
	}
}

func TestFormatProviderDeploymentID(t *testing.T) {
	actual := NewProviderDeploymentID("12345678-1234-9876-4563-123456789012", "deploymentValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Resources/deployments/deploymentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviderDeploymentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderDeploymentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Resources/deployments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Resources/deployments/deploymentValue",
			Expected: &ProviderDeploymentId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				DeploymentName: "deploymentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Resources/deployments/deploymentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderDeploymentID(v.Input)
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

		if actual.DeploymentName != v.Expected.DeploymentName {
			t.Fatalf("Expected %q but got %q for DeploymentName", v.Expected.DeploymentName, actual.DeploymentName)
		}

	}
}

func TestParseProviderDeploymentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderDeploymentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Resources/deployments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Resources/deployments/deploymentValue",
			Expected: &ProviderDeploymentId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				DeploymentName: "deploymentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Resources/deployments/deploymentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs/dEpLoYmEnTvAlUe",
			Expected: &ProviderDeploymentId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				DeploymentName: "dEpLoYmEnTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs/dEpLoYmEnTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderDeploymentIDInsensitively(v.Input)
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

		if actual.DeploymentName != v.Expected.DeploymentName {
			t.Fatalf("Expected %q but got %q for DeploymentName", v.Expected.DeploymentName, actual.DeploymentName)
		}

	}
}

func TestSegmentsForProviderDeploymentId(t *testing.T) {
	segments := ProviderDeploymentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProviderDeploymentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
