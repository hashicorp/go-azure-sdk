package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdUserConfigurationId{}

func TestNewUserIdMailFolderIdUserConfigurationID(t *testing.T) {
	id := NewUserIdMailFolderIdUserConfigurationID("userId", "mailFolderId", "userConfigurationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.MailFolderId != "mailFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderId")
	}

	if id.UserConfigurationId != "userConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConfigurationId'", id.UserConfigurationId, "userConfigurationId")
	}
}

func TestFormatUserIdMailFolderIdUserConfigurationID(t *testing.T) {
	actual := NewUserIdMailFolderIdUserConfigurationID("userId", "mailFolderId", "userConfigurationId").ID()
	expected := "/users/userId/mailFolders/mailFolderId/userConfigurations/userConfigurationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdMailFolderIdUserConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdUserConfigurationId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/userConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/userConfigurations/userConfigurationId",
			Expected: &UserIdMailFolderIdUserConfigurationId{
				UserId:              "userId",
				MailFolderId:        "mailFolderId",
				UserConfigurationId: "userConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/userConfigurations/userConfigurationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdUserConfigurationID(v.Input)
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

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.UserConfigurationId != v.Expected.UserConfigurationId {
			t.Fatalf("Expected %q but got %q for UserConfigurationId", v.Expected.UserConfigurationId, actual.UserConfigurationId)
		}

	}
}

func TestParseUserIdMailFolderIdUserConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdUserConfigurationId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/userConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/uSeRcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/userConfigurations/userConfigurationId",
			Expected: &UserIdMailFolderIdUserConfigurationId{
				UserId:              "userId",
				MailFolderId:        "mailFolderId",
				UserConfigurationId: "userConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/userConfigurations/userConfigurationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/uSeRcOnFiGuRaTiOnS/uSeRcOnFiGuRaTiOnId",
			Expected: &UserIdMailFolderIdUserConfigurationId{
				UserId:              "uSeRiD",
				MailFolderId:        "mAiLfOlDeRiD",
				UserConfigurationId: "uSeRcOnFiGuRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/uSeRcOnFiGuRaTiOnS/uSeRcOnFiGuRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdUserConfigurationIDInsensitively(v.Input)
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

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.UserConfigurationId != v.Expected.UserConfigurationId {
			t.Fatalf("Expected %q but got %q for UserConfigurationId", v.Expected.UserConfigurationId, actual.UserConfigurationId)
		}

	}
}

func TestSegmentsForUserIdMailFolderIdUserConfigurationId(t *testing.T) {
	segments := UserIdMailFolderIdUserConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdMailFolderIdUserConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
