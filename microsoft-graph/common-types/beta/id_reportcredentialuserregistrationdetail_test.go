package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportCredentialUserRegistrationDetailId{}

func TestNewReportCredentialUserRegistrationDetailID(t *testing.T) {
	id := NewReportCredentialUserRegistrationDetailID("credentialUserRegistrationDetailsId")

	if id.CredentialUserRegistrationDetailsId != "credentialUserRegistrationDetailsId" {
		t.Fatalf("Expected %q but got %q for Segment 'CredentialUserRegistrationDetailsId'", id.CredentialUserRegistrationDetailsId, "credentialUserRegistrationDetailsId")
	}
}

func TestFormatReportCredentialUserRegistrationDetailID(t *testing.T) {
	actual := NewReportCredentialUserRegistrationDetailID("credentialUserRegistrationDetailsId").ID()
	expected := "/reports/credentialUserRegistrationDetails/credentialUserRegistrationDetailsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportCredentialUserRegistrationDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportCredentialUserRegistrationDetailId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/credentialUserRegistrationDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/credentialUserRegistrationDetails/credentialUserRegistrationDetailsId",
			Expected: &ReportCredentialUserRegistrationDetailId{
				CredentialUserRegistrationDetailsId: "credentialUserRegistrationDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/credentialUserRegistrationDetails/credentialUserRegistrationDetailsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportCredentialUserRegistrationDetailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CredentialUserRegistrationDetailsId != v.Expected.CredentialUserRegistrationDetailsId {
			t.Fatalf("Expected %q but got %q for CredentialUserRegistrationDetailsId", v.Expected.CredentialUserRegistrationDetailsId, actual.CredentialUserRegistrationDetailsId)
		}

	}
}

func TestParseReportCredentialUserRegistrationDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportCredentialUserRegistrationDetailId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/credentialUserRegistrationDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/cReDeNtIaLuSeRrEgIsTrAtIoNdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/credentialUserRegistrationDetails/credentialUserRegistrationDetailsId",
			Expected: &ReportCredentialUserRegistrationDetailId{
				CredentialUserRegistrationDetailsId: "credentialUserRegistrationDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/credentialUserRegistrationDetails/credentialUserRegistrationDetailsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/cReDeNtIaLuSeRrEgIsTrAtIoNdEtAiLs/cReDeNtIaLuSeRrEgIsTrAtIoNdEtAiLsId",
			Expected: &ReportCredentialUserRegistrationDetailId{
				CredentialUserRegistrationDetailsId: "cReDeNtIaLuSeRrEgIsTrAtIoNdEtAiLsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/cReDeNtIaLuSeRrEgIsTrAtIoNdEtAiLs/cReDeNtIaLuSeRrEgIsTrAtIoNdEtAiLsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportCredentialUserRegistrationDetailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CredentialUserRegistrationDetailsId != v.Expected.CredentialUserRegistrationDetailsId {
			t.Fatalf("Expected %q but got %q for CredentialUserRegistrationDetailsId", v.Expected.CredentialUserRegistrationDetailsId, actual.CredentialUserRegistrationDetailsId)
		}

	}
}

func TestSegmentsForReportCredentialUserRegistrationDetailId(t *testing.T) {
	segments := ReportCredentialUserRegistrationDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportCredentialUserRegistrationDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
