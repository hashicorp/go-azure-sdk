package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogSignInId{}

func TestNewAuditLogSignInID(t *testing.T) {
	id := NewAuditLogSignInID("signInId")

	if id.SignInId != "signInId" {
		t.Fatalf("Expected %q but got %q for Segment 'SignInId'", id.SignInId, "signInId")
	}
}

func TestFormatAuditLogSignInID(t *testing.T) {
	actual := NewAuditLogSignInID("signInId").ID()
	expected := "/auditLogs/signIns/signInId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAuditLogSignInID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogSignInId
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
			Input: "/auditLogs/signIns",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/signIns/signInId",
			Expected: &AuditLogSignInId{
				SignInId: "signInId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/signIns/signInId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogSignInID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SignInId != v.Expected.SignInId {
			t.Fatalf("Expected %q but got %q for SignInId", v.Expected.SignInId, actual.SignInId)
		}

	}
}

func TestParseAuditLogSignInIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AuditLogSignInId
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
			Input: "/auditLogs/signIns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/sIgNiNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/auditLogs/signIns/signInId",
			Expected: &AuditLogSignInId{
				SignInId: "signInId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/auditLogs/signIns/signInId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/sIgNiNs/sIgNiNiD",
			Expected: &AuditLogSignInId{
				SignInId: "sIgNiNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aUdItLoGs/sIgNiNs/sIgNiNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAuditLogSignInIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SignInId != v.Expected.SignInId {
			t.Fatalf("Expected %q but got %q for SignInId", v.Expected.SignInId, actual.SignInId)
		}

	}
}

func TestSegmentsForAuditLogSignInId(t *testing.T) {
	segments := AuditLogSignInId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AuditLogSignInId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
