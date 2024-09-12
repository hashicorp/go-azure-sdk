package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogCustomSecurityAttributeAuditId{}

func TestNewAuditLogCustomSecurityAttributeAuditID(t *testing.T) {
	id := NewAuditLogCustomSecurityAttributeAuditID("customSecurityAttributeAuditIdValue")

	if id.CustomSecurityAttributeAuditId != "customSecurityAttributeAuditIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomSecurityAttributeAuditId'", id.CustomSecurityAttributeAuditId, "customSecurityAttributeAuditIdValue")
	}
}

func TestFormatAuditLogCustomSecurityAttributeAuditID(t *testing.T) {
	actual := NewAuditLogCustomSecurityAttributeAuditID("customSecurityAttributeAuditIdValue").ID()
	expected := "/auditLogs/customSecurityAttributeAudits/customSecurityAttributeAuditIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAuditLogCustomSecurityAttributeAuditID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogCustomSecurityAttributeAuditId
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
			Input: "/auditLogs/customSecurityAttributeAudits",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/customSecurityAttributeAudits/customSecurityAttributeAuditIdValue",
			Expected: &AuditLogCustomSecurityAttributeAuditId{
				CustomSecurityAttributeAuditId: "customSecurityAttributeAuditIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/customSecurityAttributeAudits/customSecurityAttributeAuditIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogCustomSecurityAttributeAuditID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomSecurityAttributeAuditId != v.Expected.CustomSecurityAttributeAuditId {
			t.Fatalf("Expected %q but got %q for CustomSecurityAttributeAuditId", v.Expected.CustomSecurityAttributeAuditId, actual.CustomSecurityAttributeAuditId)
		}

	}
}

func TestParseAuditLogCustomSecurityAttributeAuditIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogCustomSecurityAttributeAuditId
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
			Input: "/auditLogs/customSecurityAttributeAudits",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/cUsToMsEcUrItYaTtRiBuTeAuDiTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/customSecurityAttributeAudits/customSecurityAttributeAuditIdValue",
			Expected: &AuditLogCustomSecurityAttributeAuditId{
				CustomSecurityAttributeAuditId: "customSecurityAttributeAuditIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/customSecurityAttributeAudits/customSecurityAttributeAuditIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/cUsToMsEcUrItYaTtRiBuTeAuDiTs/cUsToMsEcUrItYaTtRiBuTeAuDiTiDvAlUe",
			Expected: &AuditLogCustomSecurityAttributeAuditId{
				CustomSecurityAttributeAuditId: "cUsToMsEcUrItYaTtRiBuTeAuDiTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/cUsToMsEcUrItYaTtRiBuTeAuDiTs/cUsToMsEcUrItYaTtRiBuTeAuDiTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogCustomSecurityAttributeAuditIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomSecurityAttributeAuditId != v.Expected.CustomSecurityAttributeAuditId {
			t.Fatalf("Expected %q but got %q for CustomSecurityAttributeAuditId", v.Expected.CustomSecurityAttributeAuditId, actual.CustomSecurityAttributeAuditId)
		}

	}
}

func TestSegmentsForAuditLogCustomSecurityAttributeAuditId(t *testing.T) {
	segments := AuditLogCustomSecurityAttributeAuditId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AuditLogCustomSecurityAttributeAuditId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
