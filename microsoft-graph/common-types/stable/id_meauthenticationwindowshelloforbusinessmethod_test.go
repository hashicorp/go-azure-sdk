package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationWindowsHelloForBusinessMethodId{}

func TestNewMeAuthenticationWindowsHelloForBusinessMethodID(t *testing.T) {
	id := NewMeAuthenticationWindowsHelloForBusinessMethodID("windowsHelloForBusinessAuthenticationMethodId")

	if id.WindowsHelloForBusinessAuthenticationMethodId != "windowsHelloForBusinessAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsHelloForBusinessAuthenticationMethodId'", id.WindowsHelloForBusinessAuthenticationMethodId, "windowsHelloForBusinessAuthenticationMethodId")
	}
}

func TestFormatMeAuthenticationWindowsHelloForBusinessMethodID(t *testing.T) {
	actual := NewMeAuthenticationWindowsHelloForBusinessMethodID("windowsHelloForBusinessAuthenticationMethodId").ID()
	expected := "/me/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationWindowsHelloForBusinessMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationWindowsHelloForBusinessMethodId
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
			Input: "/me/authentication/windowsHelloForBusinessMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodId",
			Expected: &MeAuthenticationWindowsHelloForBusinessMethodId{
				WindowsHelloForBusinessAuthenticationMethodId: "windowsHelloForBusinessAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationWindowsHelloForBusinessMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsHelloForBusinessAuthenticationMethodId != v.Expected.WindowsHelloForBusinessAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for WindowsHelloForBusinessAuthenticationMethodId", v.Expected.WindowsHelloForBusinessAuthenticationMethodId, actual.WindowsHelloForBusinessAuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationWindowsHelloForBusinessMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationWindowsHelloForBusinessMethodId
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
			Input: "/me/authentication/windowsHelloForBusinessMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/wInDoWsHeLlOfOrBuSiNeSsMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodId",
			Expected: &MeAuthenticationWindowsHelloForBusinessMethodId{
				WindowsHelloForBusinessAuthenticationMethodId: "windowsHelloForBusinessAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/wInDoWsHeLlOfOrBuSiNeSsMeThOdS/wInDoWsHeLlOfOrBuSiNeSsAuThEnTiCaTiOnMeThOdId",
			Expected: &MeAuthenticationWindowsHelloForBusinessMethodId{
				WindowsHelloForBusinessAuthenticationMethodId: "wInDoWsHeLlOfOrBuSiNeSsAuThEnTiCaTiOnMeThOdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/wInDoWsHeLlOfOrBuSiNeSsMeThOdS/wInDoWsHeLlOfOrBuSiNeSsAuThEnTiCaTiOnMeThOdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationWindowsHelloForBusinessMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsHelloForBusinessAuthenticationMethodId != v.Expected.WindowsHelloForBusinessAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for WindowsHelloForBusinessAuthenticationMethodId", v.Expected.WindowsHelloForBusinessAuthenticationMethodId, actual.WindowsHelloForBusinessAuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationWindowsHelloForBusinessMethodId(t *testing.T) {
	segments := MeAuthenticationWindowsHelloForBusinessMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationWindowsHelloForBusinessMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
