package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityCustomAuthenticationExtensionId{}

func TestNewIdentityCustomAuthenticationExtensionID(t *testing.T) {
	id := NewIdentityCustomAuthenticationExtensionID("customAuthenticationExtensionId")

	if id.CustomAuthenticationExtensionId != "customAuthenticationExtensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomAuthenticationExtensionId'", id.CustomAuthenticationExtensionId, "customAuthenticationExtensionId")
	}
}

func TestFormatIdentityCustomAuthenticationExtensionID(t *testing.T) {
	actual := NewIdentityCustomAuthenticationExtensionID("customAuthenticationExtensionId").ID()
	expected := "/identity/customAuthenticationExtensions/customAuthenticationExtensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityCustomAuthenticationExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityCustomAuthenticationExtensionId
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
			Input: "/identity/customAuthenticationExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/customAuthenticationExtensions/customAuthenticationExtensionId",
			Expected: &IdentityCustomAuthenticationExtensionId{
				CustomAuthenticationExtensionId: "customAuthenticationExtensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/customAuthenticationExtensions/customAuthenticationExtensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityCustomAuthenticationExtensionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomAuthenticationExtensionId != v.Expected.CustomAuthenticationExtensionId {
			t.Fatalf("Expected %q but got %q for CustomAuthenticationExtensionId", v.Expected.CustomAuthenticationExtensionId, actual.CustomAuthenticationExtensionId)
		}

	}
}

func TestParseIdentityCustomAuthenticationExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityCustomAuthenticationExtensionId
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
			Input: "/identity/customAuthenticationExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cUsToMaUtHeNtIcAtIoNeXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/customAuthenticationExtensions/customAuthenticationExtensionId",
			Expected: &IdentityCustomAuthenticationExtensionId{
				CustomAuthenticationExtensionId: "customAuthenticationExtensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/customAuthenticationExtensions/customAuthenticationExtensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cUsToMaUtHeNtIcAtIoNeXtEnSiOnS/cUsToMaUtHeNtIcAtIoNeXtEnSiOnId",
			Expected: &IdentityCustomAuthenticationExtensionId{
				CustomAuthenticationExtensionId: "cUsToMaUtHeNtIcAtIoNeXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cUsToMaUtHeNtIcAtIoNeXtEnSiOnS/cUsToMaUtHeNtIcAtIoNeXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityCustomAuthenticationExtensionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomAuthenticationExtensionId != v.Expected.CustomAuthenticationExtensionId {
			t.Fatalf("Expected %q but got %q for CustomAuthenticationExtensionId", v.Expected.CustomAuthenticationExtensionId, actual.CustomAuthenticationExtensionId)
		}

	}
}

func TestSegmentsForIdentityCustomAuthenticationExtensionId(t *testing.T) {
	segments := IdentityCustomAuthenticationExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityCustomAuthenticationExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
