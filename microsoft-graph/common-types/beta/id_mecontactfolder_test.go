package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactFolderId{}

func TestNewMeContactFolderID(t *testing.T) {
	id := NewMeContactFolderID("contactFolderId")

	if id.ContactFolderId != "contactFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactFolderId'", id.ContactFolderId, "contactFolderId")
	}
}

func TestFormatMeContactFolderID(t *testing.T) {
	actual := NewMeContactFolderID("contactFolderId").ID()
	expected := "/me/contactFolders/contactFolderId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeContactFolderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactFolderId
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
			// Valid URI
			Input: "/me/contactFolders/contactFolderId",
			Expected: &MeContactFolderId{
				ContactFolderId: "contactFolderId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contactFolders/contactFolderId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactFolderID(v.Input)
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

	}
}

func TestParseMeContactFolderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactFolderId
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
			// Valid URI
			Input: "/me/contactFolders/contactFolderId",
			Expected: &MeContactFolderId{
				ContactFolderId: "contactFolderId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contactFolders/contactFolderId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErId",
			Expected: &MeContactFolderId{
				ContactFolderId: "cOnTaCtFoLdErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtFoLdErS/cOnTaCtFoLdErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactFolderIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeContactFolderId(t *testing.T) {
	segments := MeContactFolderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeContactFolderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
