package usercontactfoldercontact

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactFolderId{}

func TestNewUserContactFolderID(t *testing.T) {
	id := NewUserContactFolderID("userIdValue", "contactFolderIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ContactFolderId != "contactFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId'", id.ContactFolderId, "contactFolderIdValue")
	}
}

func TestFormatUserContactFolderID(t *testing.T) {
	actual := NewUserContactFolderID("userIdValue", "contactFolderIdValue").ID()
	expected := "/users/userIdValue/contactFolders/contactFolderIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserContactFolderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserContactFolderId
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
			// Valid URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue",
			Expected: &UserContactFolderId{
				UserId:          "userIdValue",
				ContactFolderId: "contactFolderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserContactFolderID(v.Input)
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

	}
}

func TestParseUserContactFolderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserContactFolderId
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
			// Valid URI
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue",
			Expected: &UserContactFolderId{
				UserId:          "userIdValue",
				ContactFolderId: "contactFolderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/contactFolders/contactFolderIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE",
			Expected: &UserContactFolderId{
				UserId:          "uSeRiDvAlUe",
				ContactFolderId: "cOnTaCtFoLdErIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserContactFolderIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserContactFolderId(t *testing.T) {
	segments := UserContactFolderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserContactFolderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
