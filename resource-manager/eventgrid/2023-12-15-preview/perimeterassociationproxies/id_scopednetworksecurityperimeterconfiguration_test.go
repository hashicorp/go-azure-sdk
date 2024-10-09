package perimeterassociationproxies

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedNetworkSecurityPerimeterConfigurationId{}

func TestNewScopedNetworkSecurityPerimeterConfigurationID(t *testing.T) {
	id := NewScopedNetworkSecurityPerimeterConfigurationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "networkSecurityPerimeterConfigurationName")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.NetworkSecurityPerimeterConfigurationName != "networkSecurityPerimeterConfigurationName" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkSecurityPerimeterConfigurationName'", id.NetworkSecurityPerimeterConfigurationName, "networkSecurityPerimeterConfigurationName")
	}
}

func TestFormatScopedNetworkSecurityPerimeterConfigurationID(t *testing.T) {
	actual := NewScopedNetworkSecurityPerimeterConfigurationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "networkSecurityPerimeterConfigurationName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedNetworkSecurityPerimeterConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedNetworkSecurityPerimeterConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/networkSecurityPerimeterConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName",
			Expected: &ScopedNetworkSecurityPerimeterConfigurationId{
				Scope: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				NetworkSecurityPerimeterConfigurationName: "networkSecurityPerimeterConfigurationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedNetworkSecurityPerimeterConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.NetworkSecurityPerimeterConfigurationName != v.Expected.NetworkSecurityPerimeterConfigurationName {
			t.Fatalf("Expected %q but got %q for NetworkSecurityPerimeterConfigurationName", v.Expected.NetworkSecurityPerimeterConfigurationName, actual.NetworkSecurityPerimeterConfigurationName)
		}

	}
}

func TestParseScopedNetworkSecurityPerimeterConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedNetworkSecurityPerimeterConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/networkSecurityPerimeterConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName",
			Expected: &ScopedNetworkSecurityPerimeterConfigurationId{
				Scope: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				NetworkSecurityPerimeterConfigurationName: "networkSecurityPerimeterConfigurationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnS/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnNaMe",
			Expected: &ScopedNetworkSecurityPerimeterConfigurationId{
				Scope: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				NetworkSecurityPerimeterConfigurationName: "nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnS/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedNetworkSecurityPerimeterConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.NetworkSecurityPerimeterConfigurationName != v.Expected.NetworkSecurityPerimeterConfigurationName {
			t.Fatalf("Expected %q but got %q for NetworkSecurityPerimeterConfigurationName", v.Expected.NetworkSecurityPerimeterConfigurationName, actual.NetworkSecurityPerimeterConfigurationName)
		}

	}
}

func TestSegmentsForScopedNetworkSecurityPerimeterConfigurationId(t *testing.T) {
	segments := ScopedNetworkSecurityPerimeterConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedNetworkSecurityPerimeterConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
