package authenticationplatformcredentialmethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationPlatformCredentialMethodId{}

func TestNewMeAuthenticationPlatformCredentialMethodID(t *testing.T) {
	id := NewMeAuthenticationPlatformCredentialMethodID("platformCredentialAuthenticationMethodIdValue")

	if id.PlatformCredentialAuthenticationMethodId != "platformCredentialAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlatformCredentialAuthenticationMethodId'", id.PlatformCredentialAuthenticationMethodId, "platformCredentialAuthenticationMethodIdValue")
	}
}

func TestFormatMeAuthenticationPlatformCredentialMethodID(t *testing.T) {
	actual := NewMeAuthenticationPlatformCredentialMethodID("platformCredentialAuthenticationMethodIdValue").ID()
	expected := "/me/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationPlatformCredentialMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationPlatformCredentialMethodId
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
			Input: "/me/authentication/platformCredentialMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue",
			Expected: &MeAuthenticationPlatformCredentialMethodId{
				PlatformCredentialAuthenticationMethodId: "platformCredentialAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationPlatformCredentialMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PlatformCredentialAuthenticationMethodId != v.Expected.PlatformCredentialAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PlatformCredentialAuthenticationMethodId", v.Expected.PlatformCredentialAuthenticationMethodId, actual.PlatformCredentialAuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationPlatformCredentialMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationPlatformCredentialMethodId
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
			Input: "/me/authentication/platformCredentialMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pLaTfOrMcReDeNtIaLmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue",
			Expected: &MeAuthenticationPlatformCredentialMethodId{
				PlatformCredentialAuthenticationMethodId: "platformCredentialAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pLaTfOrMcReDeNtIaLmEtHoDs/pLaTfOrMcReDeNtIaLaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			Expected: &MeAuthenticationPlatformCredentialMethodId{
				PlatformCredentialAuthenticationMethodId: "pLaTfOrMcReDeNtIaLaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pLaTfOrMcReDeNtIaLmEtHoDs/pLaTfOrMcReDeNtIaLaUtHeNtIcAtIoNmEtHoDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationPlatformCredentialMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PlatformCredentialAuthenticationMethodId != v.Expected.PlatformCredentialAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PlatformCredentialAuthenticationMethodId", v.Expected.PlatformCredentialAuthenticationMethodId, actual.PlatformCredentialAuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationPlatformCredentialMethodId(t *testing.T) {
	segments := MeAuthenticationPlatformCredentialMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationPlatformCredentialMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
