package adaptivenetworkhardenings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedAdaptiveNetworkHardeningId{}

func TestNewScopedAdaptiveNetworkHardeningID(t *testing.T) {
	id := NewScopedAdaptiveNetworkHardeningID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "adaptiveNetworkHardeningValue")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.AdaptiveNetworkHardeningName != "adaptiveNetworkHardeningValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AdaptiveNetworkHardeningName'", id.AdaptiveNetworkHardeningName, "adaptiveNetworkHardeningValue")
	}
}

func TestFormatScopedAdaptiveNetworkHardeningID(t *testing.T) {
	actual := NewScopedAdaptiveNetworkHardeningID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "adaptiveNetworkHardeningValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedAdaptiveNetworkHardeningID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedAdaptiveNetworkHardeningId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/adaptiveNetworkHardenings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue",
			Expected: &ScopedAdaptiveNetworkHardeningId{
				Scope:                        "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				AdaptiveNetworkHardeningName: "adaptiveNetworkHardeningValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedAdaptiveNetworkHardeningID(v.Input)
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

		if actual.AdaptiveNetworkHardeningName != v.Expected.AdaptiveNetworkHardeningName {
			t.Fatalf("Expected %q but got %q for AdaptiveNetworkHardeningName", v.Expected.AdaptiveNetworkHardeningName, actual.AdaptiveNetworkHardeningName)
		}

	}
}

func TestParseScopedAdaptiveNetworkHardeningIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedAdaptiveNetworkHardeningId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/adaptiveNetworkHardenings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aDaPtIvEnEtWoRkHaRdEnInGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue",
			Expected: &ScopedAdaptiveNetworkHardeningId{
				Scope:                        "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				AdaptiveNetworkHardeningName: "adaptiveNetworkHardeningValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aDaPtIvEnEtWoRkHaRdEnInGs/aDaPtIvEnEtWoRkHaRdEnInGvAlUe",
			Expected: &ScopedAdaptiveNetworkHardeningId{
				Scope:                        "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				AdaptiveNetworkHardeningName: "aDaPtIvEnEtWoRkHaRdEnInGvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aDaPtIvEnEtWoRkHaRdEnInGs/aDaPtIvEnEtWoRkHaRdEnInGvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedAdaptiveNetworkHardeningIDInsensitively(v.Input)
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

		if actual.AdaptiveNetworkHardeningName != v.Expected.AdaptiveNetworkHardeningName {
			t.Fatalf("Expected %q but got %q for AdaptiveNetworkHardeningName", v.Expected.AdaptiveNetworkHardeningName, actual.AdaptiveNetworkHardeningName)
		}

	}
}

func TestSegmentsForScopedAdaptiveNetworkHardeningId(t *testing.T) {
	segments := ScopedAdaptiveNetworkHardeningId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedAdaptiveNetworkHardeningId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
