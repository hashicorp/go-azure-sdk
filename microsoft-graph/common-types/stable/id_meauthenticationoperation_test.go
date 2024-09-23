package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationOperationId{}

func TestNewMeAuthenticationOperationID(t *testing.T) {
	id := NewMeAuthenticationOperationID("longRunningOperationId")

	if id.LongRunningOperationId != "longRunningOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'LongRunningOperationId'", id.LongRunningOperationId, "longRunningOperationId")
	}
}

func TestFormatMeAuthenticationOperationID(t *testing.T) {
	actual := NewMeAuthenticationOperationID("longRunningOperationId").ID()
	expected := "/me/authentication/operations/longRunningOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationOperationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/operations/longRunningOperationId",
			Expected: &MeAuthenticationOperationId{
				LongRunningOperationId: "longRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/operations/longRunningOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LongRunningOperationId != v.Expected.LongRunningOperationId {
			t.Fatalf("Expected %q but got %q for LongRunningOperationId", v.Expected.LongRunningOperationId, actual.LongRunningOperationId)
		}

	}
}

func TestParseMeAuthenticationOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationOperationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/operations/longRunningOperationId",
			Expected: &MeAuthenticationOperationId{
				LongRunningOperationId: "longRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/operations/longRunningOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/oPeRaTiOnS/lOnGrUnNiNgOpErAtIoNiD",
			Expected: &MeAuthenticationOperationId{
				LongRunningOperationId: "lOnGrUnNiNgOpErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/oPeRaTiOnS/lOnGrUnNiNgOpErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LongRunningOperationId != v.Expected.LongRunningOperationId {
			t.Fatalf("Expected %q but got %q for LongRunningOperationId", v.Expected.LongRunningOperationId, actual.LongRunningOperationId)
		}

	}
}

func TestSegmentsForMeAuthenticationOperationId(t *testing.T) {
	segments := MeAuthenticationOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
