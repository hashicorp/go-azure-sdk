package mecontactfoldercontactextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactFolderContactExtensionId{}

func TestNewMeContactFolderContactExtensionID(t *testing.T) {
	id := NewMeContactFolderContactExtensionID("contactFolderIdValue", "contactIdValue", "extensionIdValue")

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

func TestFormatMeContactFolderContactExtensionID(t *testing.T) {
	actual := NewMeContactFolderContactExtensionID("contactFolderIdValue", "contactIdValue", "extensionIdValue").ID()
	expected := "/me/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeContactFolderContactExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactFolderContactExtensionId
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
			Input: "/me/contactFolders/contactFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderIdValue/contacts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderIdValue/contacts/contactIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue",
			Expected: &MeContactFolderContactExtensionId{
				ContactFolderId: "contactFolderIdValue",
				ContactId:       "contactIdValue",
				ExtensionId:     "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactFolderContactExtensionID(v.Input)
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

func TestParseMeContactFolderContactExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactFolderContactExtensionId
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
			Input: "/me/contactFolders/contactFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderIdValue/contacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderIdValue/contacts/contactIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS/cOnTaCtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS/cOnTaCtIdVaLuE/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue",
			Expected: &MeContactFolderContactExtensionId{
				ContactFolderId: "contactFolderIdValue",
				ContactId:       "contactIdValue",
				ExtensionId:     "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contactFolders/contactFolderIdValue/contacts/contactIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS/cOnTaCtIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &MeContactFolderContactExtensionId{
				ContactFolderId: "cOnTaCtFoLdErIdVaLuE",
				ContactId:       "cOnTaCtIdVaLuE",
				ExtensionId:     "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cOnTaCtS/cOnTaCtIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactFolderContactExtensionIDInsensitively(v.Input)
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

func TestSegmentsForMeContactFolderContactExtensionId(t *testing.T) {
	segments := MeContactFolderContactExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeContactFolderContactExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
