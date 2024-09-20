package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogDirectoryAuditId{}

func TestNewAuditLogDirectoryAuditID(t *testing.T) {
	id := NewAuditLogDirectoryAuditID("directoryAuditId")

	if id.DirectoryAuditId != "directoryAuditId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryAuditId'", id.DirectoryAuditId, "directoryAuditId")
	}
}

func TestFormatAuditLogDirectoryAuditID(t *testing.T) {
	actual := NewAuditLogDirectoryAuditID("directoryAuditId").ID()
	expected := "/auditLogs/directoryAudits/directoryAuditId"
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
			Input: "/auditLogs/directoryAudits/directoryAuditId",
			Expected: &AuditLogDirectoryAuditId{
				DirectoryAuditId: "directoryAuditId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/directoryAudits/directoryAuditId/extra",
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
			Input: "/auditLogs/directoryAudits/directoryAuditId",
			Expected: &AuditLogDirectoryAuditId{
				DirectoryAuditId: "directoryAuditId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/directoryAudits/directoryAuditId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/dIrEcToRyAuDiTs/dIrEcToRyAuDiTiD",
			Expected: &AuditLogDirectoryAuditId{
				DirectoryAuditId: "dIrEcToRyAuDiTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/dIrEcToRyAuDiTs/dIrEcToRyAuDiTiD/extra",
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
