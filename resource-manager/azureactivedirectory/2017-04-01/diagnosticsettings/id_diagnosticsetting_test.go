package diagnosticsettings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DiagnosticSettingId{}

func TestNewDiagnosticSettingID(t *testing.T) {
	id := NewDiagnosticSettingID("diagnosticSettingValue")

	if id.DiagnosticSettingName != "diagnosticSettingValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DiagnosticSettingName'", id.DiagnosticSettingName, "diagnosticSettingValue")
	}
}

func TestFormatDiagnosticSettingID(t *testing.T) {
	actual := NewDiagnosticSettingID("diagnosticSettingValue").ID()
	expected := "/providers/Microsoft.AADIAM/diagnosticSettings/diagnosticSettingValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDiagnosticSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DiagnosticSettingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AADIAM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AADIAM/diagnosticSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.AADIAM/diagnosticSettings/diagnosticSettingValue",
			Expected: &DiagnosticSettingId{
				DiagnosticSettingName: "diagnosticSettingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.AADIAM/diagnosticSettings/diagnosticSettingValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDiagnosticSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DiagnosticSettingName != v.Expected.DiagnosticSettingName {
			t.Fatalf("Expected %q but got %q for DiagnosticSettingName", v.Expected.DiagnosticSettingName, actual.DiagnosticSettingName)
		}

	}
}

func TestParseDiagnosticSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DiagnosticSettingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AADIAM",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aAdIaM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.AADIAM/diagnosticSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aAdIaM/dIaGnOsTiCsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.AADIAM/diagnosticSettings/diagnosticSettingValue",
			Expected: &DiagnosticSettingId{
				DiagnosticSettingName: "diagnosticSettingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.AADIAM/diagnosticSettings/diagnosticSettingValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aAdIaM/dIaGnOsTiCsEtTiNgS/dIaGnOsTiCsEtTiNgVaLuE",
			Expected: &DiagnosticSettingId{
				DiagnosticSettingName: "dIaGnOsTiCsEtTiNgVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aAdIaM/dIaGnOsTiCsEtTiNgS/dIaGnOsTiCsEtTiNgVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDiagnosticSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DiagnosticSettingName != v.Expected.DiagnosticSettingName {
			t.Fatalf("Expected %q but got %q for DiagnosticSettingName", v.Expected.DiagnosticSettingName, actual.DiagnosticSettingName)
		}

	}
}

func TestSegmentsForDiagnosticSettingId(t *testing.T) {
	segments := DiagnosticSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DiagnosticSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
