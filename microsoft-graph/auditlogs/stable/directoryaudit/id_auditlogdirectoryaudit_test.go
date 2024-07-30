package directoryaudit

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogDirectoryAuditId{}

func TestNewAuditLogDirectoryAuditID(t *testing.T) {
	id := NewAuditLogDirectoryAuditID("directoryAuditIdValue")

	if id.DirectoryAuditId != "directoryAuditIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryAuditId'", id.DirectoryAuditId, "directoryAuditIdValue")
	}
}

func TestFormatAuditLogDirectoryAuditID(t *testing.T) {
	actual := NewAuditLogDirectoryAuditID("directoryAuditIdValue").ID()
	expected := "/auditLogs/directoryAudits/directoryAuditIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAuditLogDirectoryAuditID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogDirectoryAuditId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/auditLogs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/auditLogs/directoryAudits",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/directoryAudits/directoryAuditIdValue",
			Expected: &AuditLogDirectoryAuditId{
				DirectoryAuditId: "directoryAuditIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/directoryAudits/directoryAuditIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogDirectoryAuditID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryAuditId != v.Expected.DirectoryAuditId {
			t.Fatalf("Expected %q but got %q for DirectoryAuditId", v.Expected.DirectoryAuditId, actual.DirectoryAuditId)
		}

	}
}

func TestParseAuditLogDirectoryAuditIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogDirectoryAuditId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/auditLogs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/auditLogs/directoryAudits",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/dIrEcToRyAuDiTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/directoryAudits/directoryAuditIdValue",
			Expected: &AuditLogDirectoryAuditId{
				DirectoryAuditId: "directoryAuditIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/directoryAudits/directoryAuditIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/dIrEcToRyAuDiTs/dIrEcToRyAuDiTiDvAlUe",
			Expected: &AuditLogDirectoryAuditId{
				DirectoryAuditId: "dIrEcToRyAuDiTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/dIrEcToRyAuDiTs/dIrEcToRyAuDiTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogDirectoryAuditIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryAuditId != v.Expected.DirectoryAuditId {
			t.Fatalf("Expected %q but got %q for DirectoryAuditId", v.Expected.DirectoryAuditId, actual.DirectoryAuditId)
		}

	}
}

func TestSegmentsForAuditLogDirectoryAuditId(t *testing.T) {
	segments := AuditLogDirectoryAuditId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AuditLogDirectoryAuditId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
