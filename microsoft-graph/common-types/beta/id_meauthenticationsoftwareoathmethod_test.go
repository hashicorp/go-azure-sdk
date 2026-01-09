package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationSoftwareOathMethodId{}

func TestNewMeAuthenticationSoftwareOathMethodID(t *testing.T) {
	id := NewMeAuthenticationSoftwareOathMethodID("softwareOathAuthenticationMethodId")

	if id.SoftwareOathAuthenticationMethodId != "softwareOathAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'SoftwareOathAuthenticationMethodId'", id.SoftwareOathAuthenticationMethodId, "softwareOathAuthenticationMethodId")
	}
}

func TestFormatMeAuthenticationSoftwareOathMethodID(t *testing.T) {
	actual := NewMeAuthenticationSoftwareOathMethodID("softwareOathAuthenticationMethodId").ID()
	expected := "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodId"
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
			Input: "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodId",
			Expected: &MeAuthenticationSoftwareOathMethodId{
				SoftwareOathAuthenticationMethodId: "softwareOathAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodId/extra",
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
			Input: "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodId",
			Expected: &MeAuthenticationSoftwareOathMethodId{
				SoftwareOathAuthenticationMethodId: "softwareOathAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/softwareOathMethods/softwareOathAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs/sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			Expected: &MeAuthenticationSoftwareOathMethodId{
				SoftwareOathAuthenticationMethodId: "sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs/sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/extra",
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
