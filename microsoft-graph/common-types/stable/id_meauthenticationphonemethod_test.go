package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationPhoneMethodId{}

func TestNewMeAuthenticationPhoneMethodID(t *testing.T) {
	id := NewMeAuthenticationPhoneMethodID("phoneAuthenticationMethodId")

	if id.PhoneAuthenticationMethodId != "phoneAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'PhoneAuthenticationMethodId'", id.PhoneAuthenticationMethodId, "phoneAuthenticationMethodId")
	}
}

func TestFormatMeAuthenticationPhoneMethodID(t *testing.T) {
	actual := NewMeAuthenticationPhoneMethodID("phoneAuthenticationMethodId").ID()
	expected := "/me/authentication/phoneMethods/phoneAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationPhoneMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationPhoneMethodId
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
			Input: "/me/authentication/phoneMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/phoneMethods/phoneAuthenticationMethodId",
			Expected: &MeAuthenticationPhoneMethodId{
				PhoneAuthenticationMethodId: "phoneAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/phoneMethods/phoneAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationPhoneMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PhoneAuthenticationMethodId != v.Expected.PhoneAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PhoneAuthenticationMethodId", v.Expected.PhoneAuthenticationMethodId, actual.PhoneAuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationPhoneMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationPhoneMethodId
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
			Input: "/me/authentication/phoneMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pHoNeMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/phoneMethods/phoneAuthenticationMethodId",
			Expected: &MeAuthenticationPhoneMethodId{
				PhoneAuthenticationMethodId: "phoneAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/phoneMethods/phoneAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pHoNeMeThOdS/pHoNeAuThEnTiCaTiOnMeThOdId",
			Expected: &MeAuthenticationPhoneMethodId{
				PhoneAuthenticationMethodId: "pHoNeAuThEnTiCaTiOnMeThOdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pHoNeMeThOdS/pHoNeAuThEnTiCaTiOnMeThOdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationPhoneMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PhoneAuthenticationMethodId != v.Expected.PhoneAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PhoneAuthenticationMethodId", v.Expected.PhoneAuthenticationMethodId, actual.PhoneAuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationPhoneMethodId(t *testing.T) {
	segments := MeAuthenticationPhoneMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationPhoneMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
