package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactFolderIdChildFolderIdContactIdExtensionId{}

func TestNewUserIdContactFolderIdChildFolderIdContactIdExtensionID(t *testing.T) {
	id := NewUserIdContactFolderIdChildFolderIdContactIdExtensionID("userId", "contactFolderId", "contactFolderId1", "contactId", "extensionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ContactFolderId != "contactFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId'", id.ContactFolderId, "contactFolderId")
	}

	if id.ContactFolderId1 != "contactFolderId1" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId1'", id.ContactFolderId1, "contactFolderId1")
	}

	if id.ContactId != "contactId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactId'", id.ContactId, "contactId")
	}

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatUserIdContactFolderIdChildFolderIdContactIdExtensionID(t *testing.T) {
	actual := NewUserIdContactFolderIdChildFolderIdContactIdExtensionID("userId", "contactFolderId", "contactFolderId1", "contactId", "extensionId").ID()
	expected := "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdContactFolderIdChildFolderIdContactIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdContactFolderIdChildFolderIdContactIdExtensionId
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
			Input: "/users/userId/contactFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId/extensions/extensionId",
			Expected: &UserIdContactFolderIdChildFolderIdContactIdExtensionId{
				UserId:           "userId",
				ContactFolderId:  "contactFolderId",
				ContactFolderId1: "contactFolderId1",
				ContactId:        "contactId",
				ExtensionId:      "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdContactFolderIdChildFolderIdContactIdExtensionID(v.Input)
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

		if actual.ContactFolderId != v.Expected.ContactFolderId {
			t.Fatalf("Expected %q but got %q for ContactFolderId", v.Expected.ContactFolderId, actual.ContactFolderId)
		}

		if actual.ContactFolderId1 != v.Expected.ContactFolderId1 {
			t.Fatalf("Expected %q but got %q for ContactFolderId1", v.Expected.ContactFolderId1, actual.ContactFolderId1)
		}

		if actual.ContactId != v.Expected.ContactId {
			t.Fatalf("Expected %q but got %q for ContactId", v.Expected.ContactId, actual.ContactId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseUserIdContactFolderIdChildFolderIdContactIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdContactFolderIdChildFolderIdContactIdExtensionId
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
			Input: "/users/userId/contactFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cHiLdFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cHiLdFoLdErS/cOnTaCtFoLdErId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cHiLdFoLdErS/cOnTaCtFoLdErId1/cOnTaCtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cHiLdFoLdErS/cOnTaCtFoLdErId1/cOnTaCtS/cOnTaCtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cHiLdFoLdErS/cOnTaCtFoLdErId1/cOnTaCtS/cOnTaCtId/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId/extensions/extensionId",
			Expected: &UserIdContactFolderIdChildFolderIdContactIdExtensionId{
				UserId:           "userId",
				ContactFolderId:  "contactFolderId",
				ContactFolderId1: "contactFolderId1",
				ContactId:        "contactId",
				ExtensionId:      "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cHiLdFoLdErS/cOnTaCtFoLdErId1/cOnTaCtS/cOnTaCtId/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &UserIdContactFolderIdChildFolderIdContactIdExtensionId{
				UserId:           "uSeRiD",
				ContactFolderId:  "cOnTaCtFoLdErId",
				ContactFolderId1: "cOnTaCtFoLdErId1",
				ContactId:        "cOnTaCtId",
				ExtensionId:      "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cHiLdFoLdErS/cOnTaCtFoLdErId1/cOnTaCtS/cOnTaCtId/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdContactFolderIdChildFolderIdContactIdExtensionIDInsensitively(v.Input)
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

		if actual.ContactFolderId != v.Expected.ContactFolderId {
			t.Fatalf("Expected %q but got %q for ContactFolderId", v.Expected.ContactFolderId, actual.ContactFolderId)
		}

		if actual.ContactFolderId1 != v.Expected.ContactFolderId1 {
			t.Fatalf("Expected %q but got %q for ContactFolderId1", v.Expected.ContactFolderId1, actual.ContactFolderId1)
		}

		if actual.ContactId != v.Expected.ContactId {
			t.Fatalf("Expected %q but got %q for ContactId", v.Expected.ContactId, actual.ContactId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForUserIdContactFolderIdChildFolderIdContactIdExtensionId(t *testing.T) {
	segments := UserIdContactFolderIdChildFolderIdContactIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdContactFolderIdChildFolderIdContactIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
