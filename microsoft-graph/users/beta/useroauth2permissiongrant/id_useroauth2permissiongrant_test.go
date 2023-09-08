package useroauth2permissiongrant

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOauth2PermissionGrantId{}

func TestNewUserOauth2PermissionGrantID(t *testing.T) {
	id := NewUserOauth2PermissionGrantID("userIdValue", "oAuth2PermissionGrantIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OAuth2PermissionGrantId != "oAuth2PermissionGrantIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OAuth2PermissionGrantId'", id.OAuth2PermissionGrantId, "oAuth2PermissionGrantIdValue")
	}
}

func TestFormatUserOauth2PermissionGrantID(t *testing.T) {
	actual := NewUserOauth2PermissionGrantID("userIdValue", "oAuth2PermissionGrantIdValue").ID()
	expected := "/users/userIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOauth2PermissionGrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOauth2PermissionGrantId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/oauth2PermissionGrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue",
			Expected: &UserOauth2PermissionGrantId{
				UserId:                  "userIdValue",
				OAuth2PermissionGrantId: "oAuth2PermissionGrantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOauth2PermissionGrantID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.OAuth2PermissionGrantId != v.Expected.OAuth2PermissionGrantId {
			t.Fatalf("Expected %q but got %q for OAuth2PermissionGrantId", v.Expected.OAuth2PermissionGrantId, actual.OAuth2PermissionGrantId)
		}

	}
}

func TestParseUserOauth2PermissionGrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOauth2PermissionGrantId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/oauth2PermissionGrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oAuTh2pErMiSsIoNgRaNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue",
			Expected: &UserOauth2PermissionGrantId{
				UserId:                  "userIdValue",
				OAuth2PermissionGrantId: "oAuth2PermissionGrantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oAuTh2pErMiSsIoNgRaNtS/oAuTh2pErMiSsIoNgRaNtIdVaLuE",
			Expected: &UserOauth2PermissionGrantId{
				UserId:                  "uSeRiDvAlUe",
				OAuth2PermissionGrantId: "oAuTh2pErMiSsIoNgRaNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oAuTh2pErMiSsIoNgRaNtS/oAuTh2pErMiSsIoNgRaNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOauth2PermissionGrantIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.OAuth2PermissionGrantId != v.Expected.OAuth2PermissionGrantId {
			t.Fatalf("Expected %q but got %q for OAuth2PermissionGrantId", v.Expected.OAuth2PermissionGrantId, actual.OAuth2PermissionGrantId)
		}

	}
}

func TestSegmentsForUserOauth2PermissionGrantId(t *testing.T) {
	segments := UserOauth2PermissionGrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOauth2PermissionGrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
