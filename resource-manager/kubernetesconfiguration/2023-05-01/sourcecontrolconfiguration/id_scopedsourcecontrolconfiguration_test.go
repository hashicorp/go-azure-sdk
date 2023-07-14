package sourcecontrolconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ScopedSourceControlConfigurationId{}

func TestNewScopedSourceControlConfigurationID(t *testing.T) {
	id := NewScopedSourceControlConfigurationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "sourceControlConfigurationValue")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.SourceControlConfigurationName != "sourceControlConfigurationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SourceControlConfigurationName'", id.SourceControlConfigurationName, "sourceControlConfigurationValue")
	}
}

func TestFormatScopedSourceControlConfigurationID(t *testing.T) {
	actual := NewScopedSourceControlConfigurationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "sourceControlConfigurationValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/sourceControlConfigurationValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedSourceControlConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedSourceControlConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.KubernetesConfiguration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/sourceControlConfigurationValue",
			Expected: &ScopedSourceControlConfigurationId{
				Scope:                          "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				SourceControlConfigurationName: "sourceControlConfigurationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/sourceControlConfigurationValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedSourceControlConfigurationID(v.Input)
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

		if actual.SourceControlConfigurationName != v.Expected.SourceControlConfigurationName {
			t.Fatalf("Expected %q but got %q for SourceControlConfigurationName", v.Expected.SourceControlConfigurationName, actual.SourceControlConfigurationName)
		}

	}
}

func TestParseScopedSourceControlConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedSourceControlConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.KubernetesConfiguration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.kUbErNeTeScOnFiGuRaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.kUbErNeTeScOnFiGuRaTiOn/sOuRcEcOnTrOlCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/sourceControlConfigurationValue",
			Expected: &ScopedSourceControlConfigurationId{
				Scope:                          "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				SourceControlConfigurationName: "sourceControlConfigurationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/sourceControlConfigurationValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.kUbErNeTeScOnFiGuRaTiOn/sOuRcEcOnTrOlCoNfIgUrAtIoNs/sOuRcEcOnTrOlCoNfIgUrAtIoNvAlUe",
			Expected: &ScopedSourceControlConfigurationId{
				Scope:                          "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				SourceControlConfigurationName: "sOuRcEcOnTrOlCoNfIgUrAtIoNvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.kUbErNeTeScOnFiGuRaTiOn/sOuRcEcOnTrOlCoNfIgUrAtIoNs/sOuRcEcOnTrOlCoNfIgUrAtIoNvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedSourceControlConfigurationIDInsensitively(v.Input)
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

		if actual.SourceControlConfigurationName != v.Expected.SourceControlConfigurationName {
			t.Fatalf("Expected %q but got %q for SourceControlConfigurationName", v.Expected.SourceControlConfigurationName, actual.SourceControlConfigurationName)
		}

	}
}

func TestSegmentsForScopedSourceControlConfigurationId(t *testing.T) {
	segments := ScopedSourceControlConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedSourceControlConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
