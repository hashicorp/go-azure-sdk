package identityconditionalaccesauthenticationcontextclassreference

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesAuthenticationContextClassReferenceId{}

func TestNewIdentityConditionalAccesAuthenticationContextClassReferenceID(t *testing.T) {
	id := NewIdentityConditionalAccesAuthenticationContextClassReferenceID("authenticationContextClassReferenceIdValue")

	if id.AuthenticationContextClassReferenceId != "authenticationContextClassReferenceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationContextClassReferenceId'", id.AuthenticationContextClassReferenceId, "authenticationContextClassReferenceIdValue")
	}
}

func TestFormatIdentityConditionalAccesAuthenticationContextClassReferenceID(t *testing.T) {
	actual := NewIdentityConditionalAccesAuthenticationContextClassReferenceID("authenticationContextClassReferenceIdValue").ID()
	expected := "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccesAuthenticationContextClassReferenceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesAuthenticationContextClassReferenceId
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
			Input: "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceIdValue",
			Expected: &IdentityConditionalAccesAuthenticationContextClassReferenceId{
				AuthenticationContextClassReferenceId: "authenticationContextClassReferenceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesAuthenticationContextClassReferenceID(v.Input)
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

func TestParseIdentityConditionalAccesAuthenticationContextClassReferenceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesAuthenticationContextClassReferenceId
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
			Input: "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceIdValue",
			Expected: &IdentityConditionalAccesAuthenticationContextClassReferenceId{
				AuthenticationContextClassReferenceId: "authenticationContextClassReferenceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationContextClassReferences/authenticationContextClassReferenceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeS/aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeIdVaLuE",
			Expected: &IdentityConditionalAccesAuthenticationContextClassReferenceId{
				AuthenticationContextClassReferenceId: "aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeS/aUtHeNtIcAtIoNcOnTeXtClAsSrEfErEnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesAuthenticationContextClassReferenceIDInsensitively(v.Input)
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

func TestSegmentsForIdentityConditionalAccesAuthenticationContextClassReferenceId(t *testing.T) {
	segments := IdentityConditionalAccesAuthenticationContextClassReferenceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccesAuthenticationContextClassReferenceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
