package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactFolderIdContactIdExtensionId{}

func TestNewMeContactFolderIdContactIdExtensionID(t *testing.T) {
	id := NewMeContactFolderIdContactIdExtensionID("contactFolderId", "contactId", "extensionId")

	if id.ContactFolderId != "contactFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId'", id.ContactFolderId, "contactFolderId")
	}

	if id.ContactId != "contactId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactId'", id.ContactId, "contactId")
	}

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatMeContactFolderIdContactIdExtensionID(t *testing.T) {
	actual := NewMeContactFolderIdContactIdExtensionID("contactFolderId", "contactId", "extensionId").ID()
	expected := "/me/contactFolders/contactFolderId/contacts/contactId/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeContactFolderIdContactIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactFolderIdContactIdExtensionId
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
			Input: "/me/contactFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderId/contacts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderId/contacts/contactId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderId/contacts/contactId/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/contactFolders/contactFolderId/contacts/contactId/extensions/extensionId",
			Expected: &MeContactFolderIdContactIdExtensionId{
				ContactFolderId: "contactFolderId",
				ContactId:       "contactId",
				ExtensionId:     "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contactFolders/contactFolderId/contacts/contactId/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactFolderIdContactIdExtensionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestParseMeContactFolderIdContactIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactFolderIdContactIdExtensionId
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
			Input: "/me/contactFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderId/contacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cOnTaCtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderId/contacts/contactId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cOnTaCtS/cOnTaCtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderId/contacts/contactId/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cOnTaCtS/cOnTaCtId/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/contactFolders/contactFolderId/contacts/contactId/extensions/extensionId",
			Expected: &MeContactFolderIdContactIdExtensionId{
				ContactFolderId: "contactFolderId",
				ContactId:       "contactId",
				ExtensionId:     "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contactFolders/contactFolderId/contacts/contactId/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cOnTaCtS/cOnTaCtId/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &MeContactFolderIdContactIdExtensionId{
				ContactFolderId: "cOnTaCtFoLdErId",
				ContactId:       "cOnTaCtId",
				ExtensionId:     "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErId/cOnTaCtS/cOnTaCtId/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactFolderIdContactIdExtensionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestSegmentsForMeContactFolderIdContactIdExtensionId(t *testing.T) {
	segments := MeContactFolderIdContactIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeContactFolderIdContactIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
