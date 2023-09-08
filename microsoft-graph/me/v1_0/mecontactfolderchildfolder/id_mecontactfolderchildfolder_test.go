package mecontactfolderchildfolder

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactFolderChildFolderId{}

func TestNewMeContactFolderChildFolderID(t *testing.T) {
	id := NewMeContactFolderChildFolderID("contactFolderIdValue", "contactFolderId1Value")

	if id.ContactFolderId != "contactFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId'", id.ContactFolderId, "contactFolderIdValue")
	}

	if id.ContactFolderId1 != "contactFolderId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId1'", id.ContactFolderId1, "contactFolderId1Value")
	}
}

func TestFormatMeContactFolderChildFolderID(t *testing.T) {
	actual := NewMeContactFolderChildFolderID("contactFolderIdValue", "contactFolderId1Value").ID()
	expected := "/me/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeContactFolderChildFolderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactFolderChildFolderId
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
			Input: "/me/contactFolders/contactFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value",
			Expected: &MeContactFolderChildFolderId{
				ContactFolderId:  "contactFolderIdValue",
				ContactFolderId1: "contactFolderId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactFolderChildFolderID(v.Input)
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

		if actual.ContactFolderId1 != v.Expected.ContactFolderId1 {
			t.Fatalf("Expected %q but got %q for ContactFolderId1", v.Expected.ContactFolderId1, actual.ContactFolderId1)
		}

	}
}

func TestParseMeContactFolderChildFolderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactFolderChildFolderId
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
			Input: "/me/contactFolders/contactFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cHiLdFoLdErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value",
			Expected: &MeContactFolderChildFolderId{
				ContactFolderId:  "contactFolderIdValue",
				ContactFolderId1: "contactFolderId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contactFolders/contactFolderIdValue/childFolders/contactFolderId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cHiLdFoLdErS/cOnTaCtFoLdErId1vAlUe",
			Expected: &MeContactFolderChildFolderId{
				ContactFolderId:  "cOnTaCtFoLdErIdVaLuE",
				ContactFolderId1: "cOnTaCtFoLdErId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErIdVaLuE/cHiLdFoLdErS/cOnTaCtFoLdErId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactFolderChildFolderIDInsensitively(v.Input)
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

		if actual.ContactFolderId1 != v.Expected.ContactFolderId1 {
			t.Fatalf("Expected %q but got %q for ContactFolderId1", v.Expected.ContactFolderId1, actual.ContactFolderId1)
		}

	}
}

func TestSegmentsForMeContactFolderChildFolderId(t *testing.T) {
	segments := MeContactFolderChildFolderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeContactFolderChildFolderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
