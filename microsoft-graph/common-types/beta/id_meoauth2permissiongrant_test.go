package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOAuth2PermissionGrantId{}

func TestNewMeOAuth2PermissionGrantID(t *testing.T) {
	id := NewMeOAuth2PermissionGrantID("oAuth2PermissionGrantId")

	if id.OAuth2PermissionGrantId != "oAuth2PermissionGrantId" {
		t.Fatalf("Expected %q but got %q for Segment 'OAuth2PermissionGrantId'", id.OAuth2PermissionGrantId, "oAuth2PermissionGrantId")
	}
}

func TestFormatMeOAuth2PermissionGrantID(t *testing.T) {
	actual := NewMeOAuth2PermissionGrantID("oAuth2PermissionGrantId").ID()
	expected := "/me/oauth2PermissionGrants/oAuth2PermissionGrantId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOAuth2PermissionGrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOAuth2PermissionGrantId
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
			Input: "/me/oauth2PermissionGrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/oauth2PermissionGrants/oAuth2PermissionGrantId",
			Expected: &MeOAuth2PermissionGrantId{
				OAuth2PermissionGrantId: "oAuth2PermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/oauth2PermissionGrants/oAuth2PermissionGrantId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOAuth2PermissionGrantID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OAuth2PermissionGrantId != v.Expected.OAuth2PermissionGrantId {
			t.Fatalf("Expected %q but got %q for OAuth2PermissionGrantId", v.Expected.OAuth2PermissionGrantId, actual.OAuth2PermissionGrantId)
		}

	}
}

func TestParseMeOAuth2PermissionGrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOAuth2PermissionGrantId
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
			Input: "/me/oauth2PermissionGrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oAuTh2pErMiSsIoNgRaNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/oauth2PermissionGrants/oAuth2PermissionGrantId",
			Expected: &MeOAuth2PermissionGrantId{
				OAuth2PermissionGrantId: "oAuth2PermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/oauth2PermissionGrants/oAuth2PermissionGrantId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oAuTh2pErMiSsIoNgRaNtS/oAuTh2pErMiSsIoNgRaNtId",
			Expected: &MeOAuth2PermissionGrantId{
				OAuth2PermissionGrantId: "oAuTh2pErMiSsIoNgRaNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oAuTh2pErMiSsIoNgRaNtS/oAuTh2pErMiSsIoNgRaNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOAuth2PermissionGrantIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OAuth2PermissionGrantId != v.Expected.OAuth2PermissionGrantId {
			t.Fatalf("Expected %q but got %q for OAuth2PermissionGrantId", v.Expected.OAuth2PermissionGrantId, actual.OAuth2PermissionGrantId)
		}

	}
}

func TestSegmentsForMeOAuth2PermissionGrantId(t *testing.T) {
	segments := MeOAuth2PermissionGrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOAuth2PermissionGrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
