package diagnosticsettingscategories

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedDiagnosticSettingsCategoryId{}

func TestNewScopedDiagnosticSettingsCategoryID(t *testing.T) {
	id := NewScopedDiagnosticSettingsCategoryID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "name")

	if id.ResourceUri != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceUri'", id.ResourceUri, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.DiagnosticSettingsCategoryName != "name" {
		t.Fatalf("Expected %q but got %q for Segment 'DiagnosticSettingsCategoryName'", id.DiagnosticSettingsCategoryName, "name")
	}
}

func TestFormatScopedDiagnosticSettingsCategoryID(t *testing.T) {
	actual := NewScopedDiagnosticSettingsCategoryID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "name").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/diagnosticSettingsCategories/name"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedDiagnosticSettingsCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedDiagnosticSettingsCategoryId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/diagnosticSettingsCategories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/diagnosticSettingsCategories/name",
			Expected: &ScopedDiagnosticSettingsCategoryId{
				ResourceUri:                    "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				DiagnosticSettingsCategoryName: "name",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/diagnosticSettingsCategories/name/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedDiagnosticSettingsCategoryID(v.Input)
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

		if actual.DiagnosticSettingsCategoryName != v.Expected.DiagnosticSettingsCategoryName {
			t.Fatalf("Expected %q but got %q for DiagnosticSettingsCategoryName", v.Expected.DiagnosticSettingsCategoryName, actual.DiagnosticSettingsCategoryName)
		}

	}
}

func TestParseScopedDiagnosticSettingsCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedDiagnosticSettingsCategoryId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/diagnosticSettingsCategories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.iNsIgHtS/dIaGnOsTiCsEtTiNgScAtEgOrIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/diagnosticSettingsCategories/name",
			Expected: &ScopedDiagnosticSettingsCategoryId{
				ResourceUri:                    "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				DiagnosticSettingsCategoryName: "name",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Insights/diagnosticSettingsCategories/name/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.iNsIgHtS/dIaGnOsTiCsEtTiNgScAtEgOrIeS/nAmE",
			Expected: &ScopedDiagnosticSettingsCategoryId{
				ResourceUri:                    "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				DiagnosticSettingsCategoryName: "nAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.iNsIgHtS/dIaGnOsTiCsEtTiNgScAtEgOrIeS/nAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedDiagnosticSettingsCategoryIDInsensitively(v.Input)
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

		if actual.DiagnosticSettingsCategoryName != v.Expected.DiagnosticSettingsCategoryName {
			t.Fatalf("Expected %q but got %q for DiagnosticSettingsCategoryName", v.Expected.DiagnosticSettingsCategoryName, actual.DiagnosticSettingsCategoryName)
		}

	}
}

func TestSegmentsForScopedDiagnosticSettingsCategoryId(t *testing.T) {
	segments := ScopedDiagnosticSettingsCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedDiagnosticSettingsCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
