package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogProvisioningId{}

func TestNewAuditLogProvisioningID(t *testing.T) {
	id := NewAuditLogProvisioningID("provisioningObjectSummaryId")

	if id.ProvisioningObjectSummaryId != "provisioningObjectSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'ProvisioningObjectSummaryId'", id.ProvisioningObjectSummaryId, "provisioningObjectSummaryId")
	}
}

func TestFormatAuditLogProvisioningID(t *testing.T) {
	actual := NewAuditLogProvisioningID("provisioningObjectSummaryId").ID()
	expected := "/auditLogs/provisioning/provisioningObjectSummaryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAuditLogProvisioningID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogProvisioningId
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
			Input: "/auditLogs/provisioning",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/provisioning/provisioningObjectSummaryId",
			Expected: &AuditLogProvisioningId{
				ProvisioningObjectSummaryId: "provisioningObjectSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/provisioning/provisioningObjectSummaryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogProvisioningID(v.Input)
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

func TestParseAuditLogProvisioningIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogProvisioningId
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
			Input: "/auditLogs/provisioning",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/pRoViSiOnInG",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/provisioning/provisioningObjectSummaryId",
			Expected: &AuditLogProvisioningId{
				ProvisioningObjectSummaryId: "provisioningObjectSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/provisioning/provisioningObjectSummaryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/pRoViSiOnInG/pRoViSiOnInGoBjEcTsUmMaRyId",
			Expected: &AuditLogProvisioningId{
				ProvisioningObjectSummaryId: "pRoViSiOnInGoBjEcTsUmMaRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/pRoViSiOnInG/pRoViSiOnInGoBjEcTsUmMaRyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogProvisioningIDInsensitively(v.Input)
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

func TestSegmentsForAuditLogProvisioningId(t *testing.T) {
	segments := AuditLogProvisioningId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AuditLogProvisioningId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
