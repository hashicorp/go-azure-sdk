package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowId{}

func TestNewIdentityAuthenticationEventsFlowID(t *testing.T) {
	id := NewIdentityAuthenticationEventsFlowID("authenticationEventsFlowIdValue")

	if id.AuthenticationEventsFlowId != "authenticationEventsFlowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationEventsFlowId'", id.AuthenticationEventsFlowId, "authenticationEventsFlowIdValue")
	}
}

func TestFormatIdentityAuthenticationEventsFlowID(t *testing.T) {
	actual := NewIdentityAuthenticationEventsFlowID("authenticationEventsFlowIdValue").ID()
	expected := "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityAuthenticationEventsFlowID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowId
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
			Input: "/identity/authenticationEventsFlows",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue",
			Expected: &IdentityAuthenticationEventsFlowId{
				AuthenticationEventsFlowId: "authenticationEventsFlowIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationEventsFlowId != v.Expected.AuthenticationEventsFlowId {
			t.Fatalf("Expected %q but got %q for AuthenticationEventsFlowId", v.Expected.AuthenticationEventsFlowId, actual.AuthenticationEventsFlowId)
		}

	}
}

func TestParseIdentityAuthenticationEventsFlowIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowId
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
			Input: "/identity/authenticationEventsFlows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue",
			Expected: &IdentityAuthenticationEventsFlowId{
				AuthenticationEventsFlowId: "authenticationEventsFlowIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe",
			Expected: &IdentityAuthenticationEventsFlowId{
				AuthenticationEventsFlowId: "aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationEventsFlowId != v.Expected.AuthenticationEventsFlowId {
			t.Fatalf("Expected %q but got %q for AuthenticationEventsFlowId", v.Expected.AuthenticationEventsFlowId, actual.AuthenticationEventsFlowId)
		}

	}
}

func TestSegmentsForIdentityAuthenticationEventsFlowId(t *testing.T) {
	segments := IdentityAuthenticationEventsFlowId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityAuthenticationEventsFlowId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
