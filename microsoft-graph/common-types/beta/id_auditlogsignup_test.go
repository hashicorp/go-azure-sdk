package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogSignUpId{}

func TestNewAuditLogSignUpID(t *testing.T) {
	id := NewAuditLogSignUpID("selfServiceSignUpId")

	if id.SelfServiceSignUpId != "selfServiceSignUpId" {
		t.Fatalf("Expected %q but got %q for Segment 'SelfServiceSignUpId'", id.SelfServiceSignUpId, "selfServiceSignUpId")
	}
}

func TestFormatAuditLogSignUpID(t *testing.T) {
	actual := NewAuditLogSignUpID("selfServiceSignUpId").ID()
	expected := "/auditLogs/signUps/selfServiceSignUpId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAuditLogSignUpID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogSignUpId
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
			Input: "/auditLogs/signUps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/signUps/selfServiceSignUpId",
			Expected: &AuditLogSignUpId{
				SelfServiceSignUpId: "selfServiceSignUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/signUps/selfServiceSignUpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogSignUpID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SelfServiceSignUpId != v.Expected.SelfServiceSignUpId {
			t.Fatalf("Expected %q but got %q for SelfServiceSignUpId", v.Expected.SelfServiceSignUpId, actual.SelfServiceSignUpId)
		}

	}
}

func TestParseAuditLogSignUpIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogSignUpId
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
			Input: "/auditLogs/signUps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/sIgNuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/signUps/selfServiceSignUpId",
			Expected: &AuditLogSignUpId{
				SelfServiceSignUpId: "selfServiceSignUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/signUps/selfServiceSignUpId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/sIgNuPs/sElFsErViCeSiGnUpId",
			Expected: &AuditLogSignUpId{
				SelfServiceSignUpId: "sElFsErViCeSiGnUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/sIgNuPs/sElFsErViCeSiGnUpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogSignUpIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SelfServiceSignUpId != v.Expected.SelfServiceSignUpId {
			t.Fatalf("Expected %q but got %q for SelfServiceSignUpId", v.Expected.SelfServiceSignUpId, actual.SelfServiceSignUpId)
		}

	}
}

func TestSegmentsForAuditLogSignUpId(t *testing.T) {
	segments := AuditLogSignUpId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AuditLogSignUpId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
