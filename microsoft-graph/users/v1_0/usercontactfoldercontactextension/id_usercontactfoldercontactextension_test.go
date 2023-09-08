package usercontactfoldercontactextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactFolderContactExtensionId{}

func TestNewUserContactFolderContactExtensionID(t *testing.T) {
	id := NewUserContactFolderContactExtensionID("userIdValue", "contactFolderIdValue", "contactIdValue", "extensionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ContactFolderId != "contactFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId'", id.ContactFolderId, "contactFolderIdValue")
	}

	if id.ContactId != "contactIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactId'", id.ContactId, "contactIdValue")
	}

	if id.ExtensionId != "extensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionIdValue")
	}
}

func TestFormatUserContactFolderContactExtensionID(t *testing.T) {
	actual := NewUserContactFolderContactExtensionID("userIdValue", "contactFolderIdValue", "contactIdValue", "extensionIdValue").ID()
	expected := "/users/userIdValue/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserContactFolderContactExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserContactFolderContactExtensionId
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
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts/contactIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue",
			Expected: &UserContactFolderContactExtensionId{
				UserId:          "userIdValue",
				ContactFolderId: "contactFolderIdValue",
				ContactId:       "contactIdValue",
				ExtensionId:     "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserContactFolderContactExtensionID(v.Input)
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

		if actual.ContactId != v.Expected.ContactId {
			t.Fatalf("Expected %q but got %q for ContactId", v.Expected.ContactId, actual.ContactId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseUserContactFolderContactExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserContactFolderContactExtensionId
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
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts/contactIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS/cOnTaCtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS/cOnTaCtIdVaLuE/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue",
			Expected: &UserContactFolderContactExtensionId{
				UserId:          "userIdValue",
				ContactFolderId: "contactFolderIdValue",
				ContactId:       "contactIdValue",
				ExtensionId:     "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS/cOnTaCtIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &UserContactFolderContactExtensionId{
				UserId:          "uSeRiDvAlUe",
				ContactFolderId: "cOnTaCtFoLdErIdVaLuE",
				ContactId:       "cOnTaCtIdVaLuE",
				ExtensionId:     "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS/cOnTaCtIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserContactFolderContactExtensionIDInsensitively(v.Input)
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

		if actual.ContactId != v.Expected.ContactId {
			t.Fatalf("Expected %q but got %q for ContactId", v.Expected.ContactId, actual.ContactId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForUserContactFolderContactExtensionId(t *testing.T) {
	segments := UserContactFolderContactExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserContactFolderContactExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
