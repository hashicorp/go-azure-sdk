package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdChildFolderIdOperationId{}

func TestNewMeMailFolderIdChildFolderIdOperationID(t *testing.T) {
	id := NewMeMailFolderIdChildFolderIdOperationID("mailFolderId", "mailFolderId1", "mailFolderOperationId")

	if id.MailFolderId != "mailFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderId")
	}

	if id.MailFolderId1 != "mailFolderId1" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId1'", id.MailFolderId1, "mailFolderId1")
	}

	if id.MailFolderOperationId != "mailFolderOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderOperationId'", id.MailFolderOperationId, "mailFolderOperationId")
	}
}

func TestFormatMeMailFolderIdChildFolderIdOperationID(t *testing.T) {
	actual := NewMeMailFolderIdChildFolderIdOperationID("mailFolderId", "mailFolderId1", "mailFolderOperationId").ID()
	expected := "/me/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMailFolderIdChildFolderIdOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderIdChildFolderIdOperationId
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
			Input: "/me/mailFolders/mailFolderId/childFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/childFolders/mailFolderId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/childFolders/mailFolderId1/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId",
			Expected: &MeMailFolderIdChildFolderIdOperationId{
				MailFolderId:          "mailFolderId",
				MailFolderId1:         "mailFolderId1",
				MailFolderOperationId: "mailFolderOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderIdChildFolderIdOperationID(v.Input)
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

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.MailFolderOperationId != v.Expected.MailFolderOperationId {
			t.Fatalf("Expected %q but got %q for MailFolderOperationId", v.Expected.MailFolderOperationId, actual.MailFolderOperationId)
		}

	}
}

func TestParseMeMailFolderIdChildFolderIdOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderIdChildFolderIdOperationId
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
			Input: "/me/mailFolders/mailFolderId/childFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/childFolders/mailFolderId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/childFolders/mailFolderId1/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId",
			Expected: &MeMailFolderIdChildFolderIdOperationId{
				MailFolderId:          "mailFolderId",
				MailFolderId1:         "mailFolderId1",
				MailFolderOperationId: "mailFolderOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1/oPeRaTiOnS/mAiLfOlDeRoPeRaTiOnId",
			Expected: &MeMailFolderIdChildFolderIdOperationId{
				MailFolderId:          "mAiLfOlDeRiD",
				MailFolderId1:         "mAiLfOlDeRiD1",
				MailFolderOperationId: "mAiLfOlDeRoPeRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1/oPeRaTiOnS/mAiLfOlDeRoPeRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderIdChildFolderIdOperationIDInsensitively(v.Input)
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

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.MailFolderOperationId != v.Expected.MailFolderOperationId {
			t.Fatalf("Expected %q but got %q for MailFolderOperationId", v.Expected.MailFolderOperationId, actual.MailFolderOperationId)
		}

	}
}

func TestSegmentsForMeMailFolderIdChildFolderIdOperationId(t *testing.T) {
	segments := MeMailFolderIdChildFolderIdOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMailFolderIdChildFolderIdOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
