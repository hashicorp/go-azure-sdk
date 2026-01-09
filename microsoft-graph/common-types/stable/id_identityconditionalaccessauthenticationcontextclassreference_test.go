package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessAuthenticationContextClassReferenceId{}

func TestNewIdentityConditionalAccessAuthenticationContextClassReferenceID(t *testing.T) {
	id := NewIdentityConditionalAccessAuthenticationContextClassReferenceID("authenticationContextClassReferenceId")

	if id.AuthenticationContextClassReferenceId != "authenticationContextClassReferenceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationContextClassReferenceId'", id.AuthenticationContextClassReferenceId, "authenticationContextClassReferenceId")
	}
}

func TestFormatIdentityConditionalAccessAuthenticationContextClassReferenceID(t *testing.T) {
	actual := NewIdentityConditionalAccessAuthenticationContextClassReferenceID("authenticationContextClassReferenceId").ID()
	expected := "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccessAuthenticationContextClassReferenceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessAuthenticationContextClassReferenceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationContextClassReferences",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceId",
			Expected: &IdentityConditionalAccessAuthenticationContextClassReferenceId{
				AuthenticationContextClassReferenceId: "authenticationContextClassReferenceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessAuthenticationContextClassReferenceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationContextClassReferenceId != v.Expected.AuthenticationContextClassReferenceId {
			t.Fatalf("Expected %q but got %q for AuthenticationContextClassReferenceId", v.Expected.AuthenticationContextClassReferenceId, actual.AuthenticationContextClassReferenceId)
		}

	}
}

func TestParseIdentityConditionalAccessAuthenticationContextClassReferenceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessAuthenticationContextClassReferenceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationContextClassReferences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceId",
			Expected: &IdentityConditionalAccessAuthenticationContextClassReferenceId{
				AuthenticationContextClassReferenceId: "authenticationContextClassReferenceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeS/aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeId",
			Expected: &IdentityConditionalAccessAuthenticationContextClassReferenceId{
				AuthenticationContextClassReferenceId: "aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeS/aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessAuthenticationContextClassReferenceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationContextClassReferenceId != v.Expected.AuthenticationContextClassReferenceId {
			t.Fatalf("Expected %q but got %q for AuthenticationContextClassReferenceId", v.Expected.AuthenticationContextClassReferenceId, actual.AuthenticationContextClassReferenceId)
		}

	}
}

func TestSegmentsForIdentityConditionalAccessAuthenticationContextClassReferenceId(t *testing.T) {
	segments := IdentityConditionalAccessAuthenticationContextClassReferenceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccessAuthenticationContextClassReferenceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
