package usermailfolderchildfolderuserconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderChildFolderUserConfigurationId{}

func TestNewUserMailFolderChildFolderUserConfigurationID(t *testing.T) {
	id := NewUserMailFolderChildFolderUserConfigurationID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "userConfigurationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.MailFolderId != "mailFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderIdValue")
	}

	if id.MailFolderId1 != "mailFolderId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId1'", id.MailFolderId1, "mailFolderId1Value")
	}

	if id.UserConfigurationId != "userConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConfigurationId'", id.UserConfigurationId, "userConfigurationIdValue")
	}
}

func TestFormatUserMailFolderChildFolderUserConfigurationID(t *testing.T) {
	actual := NewUserMailFolderChildFolderUserConfigurationID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "userConfigurationIdValue").ID()
	expected := "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/userConfigurations/userConfigurationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserMailFolderChildFolderUserConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMailFolderChildFolderUserConfigurationId
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
			Input: "/users/userIdValue/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/userConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/userConfigurations/userConfigurationIdValue",
			Expected: &UserMailFolderChildFolderUserConfigurationId{
				UserId:              "userIdValue",
				MailFolderId:        "mailFolderIdValue",
				MailFolderId1:       "mailFolderId1Value",
				UserConfigurationId: "userConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/userConfigurations/userConfigurationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMailFolderChildFolderUserConfigurationID(v.Input)
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

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.UserConfigurationId != v.Expected.UserConfigurationId {
			t.Fatalf("Expected %q but got %q for UserConfigurationId", v.Expected.UserConfigurationId, actual.UserConfigurationId)
		}

	}
}

func TestParseUserMailFolderChildFolderUserConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMailFolderChildFolderUserConfigurationId
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
			Input: "/users/userIdValue/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/userConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/uSeRcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/userConfigurations/userConfigurationIdValue",
			Expected: &UserMailFolderChildFolderUserConfigurationId{
				UserId:              "userIdValue",
				MailFolderId:        "mailFolderIdValue",
				MailFolderId1:       "mailFolderId1Value",
				UserConfigurationId: "userConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/userConfigurations/userConfigurationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/uSeRcOnFiGuRaTiOnS/uSeRcOnFiGuRaTiOnIdVaLuE",
			Expected: &UserMailFolderChildFolderUserConfigurationId{
				UserId:              "uSeRiDvAlUe",
				MailFolderId:        "mAiLfOlDeRiDvAlUe",
				MailFolderId1:       "mAiLfOlDeRiD1VaLuE",
				UserConfigurationId: "uSeRcOnFiGuRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/uSeRcOnFiGuRaTiOnS/uSeRcOnFiGuRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMailFolderChildFolderUserConfigurationIDInsensitively(v.Input)
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

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.UserConfigurationId != v.Expected.UserConfigurationId {
			t.Fatalf("Expected %q but got %q for UserConfigurationId", v.Expected.UserConfigurationId, actual.UserConfigurationId)
		}

	}
}

func TestSegmentsForUserMailFolderChildFolderUserConfigurationId(t *testing.T) {
	segments := UserMailFolderChildFolderUserConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserMailFolderChildFolderUserConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
