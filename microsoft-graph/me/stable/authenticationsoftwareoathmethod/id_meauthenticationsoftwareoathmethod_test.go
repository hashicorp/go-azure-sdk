package authenticationsoftwareoathmethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationSoftwareOathMethodId{}

func TestNewMeAuthenticationSoftwareOathMethodID(t *testing.T) {
	id := NewMeAuthenticationSoftwareOathMethodID("softwareOathAuthenticationMethodIdValue")

	if id.SoftwareOathAuthenticationMethodId != "softwareOathAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SoftwareOathAuthenticationMethodId'", id.SoftwareOathAuthenticationMethodId, "softwareOathAuthenticationMethodIdValue")
	}
}

func TestFormatMeAuthenticationSoftwareOathMethodID(t *testing.T) {
	actual := NewMeAuthenticationSoftwareOathMethodID("softwareOathAuthenticationMethodIdValue").ID()
	expected := "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationSoftwareOathMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationSoftwareOathMethodId
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
			Input: "/me/authentication/softwareOathMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue",
			Expected: &MeAuthenticationSoftwareOathMethodId{
				SoftwareOathAuthenticationMethodId: "softwareOathAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationSoftwareOathMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SoftwareOathAuthenticationMethodId != v.Expected.SoftwareOathAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for SoftwareOathAuthenticationMethodId", v.Expected.SoftwareOathAuthenticationMethodId, actual.SoftwareOathAuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationSoftwareOathMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationSoftwareOathMethodId
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
			Input: "/me/authentication/softwareOathMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue",
			Expected: &MeAuthenticationSoftwareOathMethodId{
				SoftwareOathAuthenticationMethodId: "softwareOathAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs/sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			Expected: &MeAuthenticationSoftwareOathMethodId{
				SoftwareOathAuthenticationMethodId: "sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs/sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationSoftwareOathMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SoftwareOathAuthenticationMethodId != v.Expected.SoftwareOathAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for SoftwareOathAuthenticationMethodId", v.Expected.SoftwareOathAuthenticationMethodId, actual.SoftwareOathAuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationSoftwareOathMethodId(t *testing.T) {
	segments := MeAuthenticationSoftwareOathMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationSoftwareOathMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
