package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactFolderIdChildFolderIdContactId{}

func TestNewUserIdContactFolderIdChildFolderIdContactID(t *testing.T) {
	id := NewUserIdContactFolderIdChildFolderIdContactID("userId", "contactFolderId", "contactFolderId1", "contactId")

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
}

func TestFormatUserIdContactFolderIdChildFolderIdContactID(t *testing.T) {
	actual := NewUserIdContactFolderIdChildFolderIdContactID("userId", "contactFolderId", "contactFolderId1", "contactId").ID()
	expected := "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdContactFolderIdChildFolderIdContactID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdContactFolderIdChildFolderIdContactId
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
			// Valid URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId",
			Expected: &UserIdContactFolderIdChildFolderIdContactId{
				UserId:           "userId",
				ContactFolderId:  "contactFolderId",
				ContactFolderId1: "contactFolderId1",
				ContactId:        "contactId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdContactFolderIdChildFolderIdContactID(v.Input)
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

	}
}

func TestParseUserIdContactFolderIdChildFolderIdContactIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdContactFolderIdChildFolderIdContactId
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
			// Valid URI
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId",
			Expected: &UserIdContactFolderIdChildFolderIdContactId{
				UserId:           "userId",
				ContactFolderId:  "contactFolderId",
				ContactFolderId1: "contactFolderId1",
				ContactId:        "contactId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/contactFolders/contactFolderId/childFolders/contactFolderId1/contacts/contactId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cHiLdFoLdErS/cOnTaCtFoLdErId1/cOnTaCtS/cOnTaCtId",
			Expected: &UserIdContactFolderIdChildFolderIdContactId{
				UserId:           "uSeRiD",
				ContactFolderId:  "cOnTaCtFoLdErId",
				ContactFolderId1: "cOnTaCtFoLdErId1",
				ContactId:        "cOnTaCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cHiLdFoLdErS/cOnTaCtFoLdErId1/cOnTaCtS/cOnTaCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdContactFolderIdChildFolderIdContactIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserIdContactFolderIdChildFolderIdContactId(t *testing.T) {
	segments := UserIdContactFolderIdChildFolderIdContactId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdContactFolderIdChildFolderIdContactId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
