package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdOperationId{}

func TestNewMeMailFolderIdOperationID(t *testing.T) {
	id := NewMeMailFolderIdOperationID("mailFolderId", "mailFolderOperationId")

	if id.MailFolderId != "mailFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderId")
	}

	if id.MailFolderOperationId != "mailFolderOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderOperationId'", id.MailFolderOperationId, "mailFolderOperationId")
	}
}

func TestFormatMeMailFolderIdOperationID(t *testing.T) {
	actual := NewMeMailFolderIdOperationID("mailFolderId", "mailFolderOperationId").ID()
	expected := "/me/mailFolders/mailFolderId/operations/mailFolderOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMailFolderIdOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderIdOperationId
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
			Input: "/me/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderId/operations/mailFolderOperationId",
			Expected: &MeMailFolderIdOperationId{
				MailFolderId:          "mailFolderId",
				MailFolderOperationId: "mailFolderOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderId/operations/mailFolderOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderIdOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.MailFolderOperationId != v.Expected.MailFolderOperationId {
			t.Fatalf("Expected %q but got %q for MailFolderOperationId", v.Expected.MailFolderOperationId, actual.MailFolderOperationId)
		}

	}
}

func TestParseMeMailFolderIdOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderIdOperationId
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
			Input: "/me/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderId/operations/mailFolderOperationId",
			Expected: &MeMailFolderIdOperationId{
				MailFolderId:          "mailFolderId",
				MailFolderOperationId: "mailFolderOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderId/operations/mailFolderOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/oPeRaTiOnS/mAiLfOlDeRoPeRaTiOnId",
			Expected: &MeMailFolderIdOperationId{
				MailFolderId:          "mAiLfOlDeRiD",
				MailFolderOperationId: "mAiLfOlDeRoPeRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/oPeRaTiOnS/mAiLfOlDeRoPeRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderIdOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.MailFolderOperationId != v.Expected.MailFolderOperationId {
			t.Fatalf("Expected %q but got %q for MailFolderOperationId", v.Expected.MailFolderOperationId, actual.MailFolderOperationId)
		}

	}
}

func TestSegmentsForMeMailFolderIdOperationId(t *testing.T) {
	segments := MeMailFolderIdOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMailFolderIdOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
