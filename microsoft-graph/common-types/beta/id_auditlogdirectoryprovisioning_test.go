package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogDirectoryProvisioningId{}

func TestNewAuditLogDirectoryProvisioningID(t *testing.T) {
	id := NewAuditLogDirectoryProvisioningID("provisioningObjectSummaryId")

	if id.ProvisioningObjectSummaryId != "provisioningObjectSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'ProvisioningObjectSummaryId'", id.ProvisioningObjectSummaryId, "provisioningObjectSummaryId")
	}
}

func TestFormatAuditLogDirectoryProvisioningID(t *testing.T) {
	actual := NewAuditLogDirectoryProvisioningID("provisioningObjectSummaryId").ID()
	expected := "/auditLogs/directoryProvisioning/provisioningObjectSummaryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAuditLogDirectoryProvisioningID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogDirectoryProvisioningId
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
			Input: "/auditLogs/directoryProvisioning",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/directoryProvisioning/provisioningObjectSummaryId",
			Expected: &AuditLogDirectoryProvisioningId{
				ProvisioningObjectSummaryId: "provisioningObjectSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/directoryProvisioning/provisioningObjectSummaryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogDirectoryProvisioningID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProvisioningObjectSummaryId != v.Expected.ProvisioningObjectSummaryId {
			t.Fatalf("Expected %q but got %q for ProvisioningObjectSummaryId", v.Expected.ProvisioningObjectSummaryId, actual.ProvisioningObjectSummaryId)
		}

	}
}

func TestParseAuditLogDirectoryProvisioningIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogDirectoryProvisioningId
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
			Input: "/auditLogs/directoryProvisioning",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/dIrEcToRyPrOvIsIoNiNg",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/directoryProvisioning/provisioningObjectSummaryId",
			Expected: &AuditLogDirectoryProvisioningId{
				ProvisioningObjectSummaryId: "provisioningObjectSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/directoryProvisioning/provisioningObjectSummaryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/dIrEcToRyPrOvIsIoNiNg/pRoViSiOnInGoBjEcTsUmMaRyId",
			Expected: &AuditLogDirectoryProvisioningId{
				ProvisioningObjectSummaryId: "pRoViSiOnInGoBjEcTsUmMaRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/dIrEcToRyPrOvIsIoNiNg/pRoViSiOnInGoBjEcTsUmMaRyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogDirectoryProvisioningIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProvisioningObjectSummaryId != v.Expected.ProvisioningObjectSummaryId {
			t.Fatalf("Expected %q but got %q for ProvisioningObjectSummaryId", v.Expected.ProvisioningObjectSummaryId, actual.ProvisioningObjectSummaryId)
		}

	}
}

func TestSegmentsForAuditLogDirectoryProvisioningId(t *testing.T) {
	segments := AuditLogDirectoryProvisioningId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AuditLogDirectoryProvisioningId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
