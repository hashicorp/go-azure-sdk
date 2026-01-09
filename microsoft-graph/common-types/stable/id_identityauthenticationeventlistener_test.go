package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventListenerId{}

func TestNewIdentityAuthenticationEventListenerID(t *testing.T) {
	id := NewIdentityAuthenticationEventListenerID("authenticationEventListenerId")

	if id.AuthenticationEventListenerId != "authenticationEventListenerId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationEventListenerId'", id.AuthenticationEventListenerId, "authenticationEventListenerId")
	}
}

func TestFormatIdentityAuthenticationEventListenerID(t *testing.T) {
	actual := NewIdentityAuthenticationEventListenerID("authenticationEventListenerId").ID()
	expected := "/identity/authenticationEventListeners/authenticationEventListenerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityAuthenticationEventListenerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventListenerId
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
			Input: "/identity/authenticationEventListeners",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventListeners/authenticationEventListenerId",
			Expected: &IdentityAuthenticationEventListenerId{
				AuthenticationEventListenerId: "authenticationEventListenerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventListeners/authenticationEventListenerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventListenerID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationEventListenerId != v.Expected.AuthenticationEventListenerId {
			t.Fatalf("Expected %q but got %q for AuthenticationEventListenerId", v.Expected.AuthenticationEventListenerId, actual.AuthenticationEventListenerId)
		}

	}
}

func TestParseIdentityAuthenticationEventListenerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventListenerId
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
			Input: "/identity/authenticationEventListeners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtLiStEnErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventListeners/authenticationEventListenerId",
			Expected: &IdentityAuthenticationEventListenerId{
				AuthenticationEventListenerId: "authenticationEventListenerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventListeners/authenticationEventListenerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtLiStEnErS/aUtHeNtIcAtIoNeVeNtLiStEnErId",
			Expected: &IdentityAuthenticationEventListenerId{
				AuthenticationEventListenerId: "aUtHeNtIcAtIoNeVeNtLiStEnErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtLiStEnErS/aUtHeNtIcAtIoNeVeNtLiStEnErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventListenerIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationEventListenerId != v.Expected.AuthenticationEventListenerId {
			t.Fatalf("Expected %q but got %q for AuthenticationEventListenerId", v.Expected.AuthenticationEventListenerId, actual.AuthenticationEventListenerId)
		}

	}
}

func TestSegmentsForIdentityAuthenticationEventListenerId(t *testing.T) {
	segments := IdentityAuthenticationEventListenerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityAuthenticationEventListenerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
