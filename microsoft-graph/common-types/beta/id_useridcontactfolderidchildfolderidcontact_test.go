package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactFolderIdChildFolderIdContactId{}

func TestNewUserIdContactFolderIdChildFolderIdContactID(t *testing.T) {
	id := NewUserIdContactFolderIdChildFolderIdContactID("userIdValue", "contactFolderIdValue", "contactFolderId1Value", "contactIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ContactFolderId != "contactFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId'", id.ContactFolderId, "contactFolderIdValue")
	}

	if id.ContactFolderId1 != "contactFolderId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId1'", id.ContactFolderId1, "contactFolderId1Value")
	}

	if id.ContactId != "contactIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactId'", id.ContactId, "contactIdValue")
	}
}

func TestFormatUserIdContactFolderIdChildFolderIdContactID(t *testing.T) {
	actual := NewUserIdContactFolderIdChildFolderIdContactID("userIdValue", "contactFolderIdValue", "contactFolderId1Value", "contactIdValue").ID()
	expected := "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value/contacts/contactIdValue"
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value/contacts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value/contacts/contactIdValue",
			Expected: &UserIdContactFolderIdChildFolderIdContactId{
				UserId:           "userIdValue",
				ContactFolderId:  "contactFolderIdValue",
				ContactFolderId1: "contactFolderId1Value",
				ContactId:        "contactIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value/contacts/contactIdValue/extra",
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
			Input: "/users/userIdValue/contactFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cHiLdFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cHiLdFoLdErS/cOnTaCtFoLdErId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value/contacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cHiLdFoLdErS/cOnTaCtFoLdErId1vAlUe/cOnTaCtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value/contacts/contactIdValue",
			Expected: &UserIdContactFolderIdChildFolderIdContactId{
				UserId:           "userIdValue",
				ContactFolderId:  "contactFolderIdValue",
				ContactFolderId1: "contactFolderId1Value",
				ContactId:        "contactIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value/contacts/contactIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cHiLdFoLdErS/cOnTaCtFoLdErId1vAlUe/cOnTaCtS/cOnTaCtIdVaLuE",
			Expected: &UserIdContactFolderIdChildFolderIdContactId{
				UserId:           "uSeRiDvAlUe",
				ContactFolderId:  "cOnTaCtFoLdErIdVaLuE",
				ContactFolderId1: "cOnTaCtFoLdErId1vAlUe",
				ContactId:        "cOnTaCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cHiLdFoLdErS/cOnTaCtFoLdErId1vAlUe/cOnTaCtS/cOnTaCtIdVaLuE/extra",
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
