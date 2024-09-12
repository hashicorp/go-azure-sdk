package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookMasterCategoryId{}

func TestNewUserIdOutlookMasterCategoryID(t *testing.T) {
	id := NewUserIdOutlookMasterCategoryID("userIdValue", "outlookCategoryIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OutlookCategoryId != "outlookCategoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookCategoryId'", id.OutlookCategoryId, "outlookCategoryIdValue")
	}
}

func TestFormatUserIdOutlookMasterCategoryID(t *testing.T) {
	actual := NewUserIdOutlookMasterCategoryID("userIdValue", "outlookCategoryIdValue").ID()
	expected := "/users/userIdValue/outlook/masterCategories/outlookCategoryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOutlookMasterCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookMasterCategoryId
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
			Input: "/users/userIdValue/outlook",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/masterCategories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/outlook/masterCategories/outlookCategoryIdValue",
			Expected: &UserIdOutlookMasterCategoryId{
				UserId:            "userIdValue",
				OutlookCategoryId: "outlookCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/outlook/masterCategories/outlookCategoryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookMasterCategoryID(v.Input)
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

		if actual.OutlookCategoryId != v.Expected.OutlookCategoryId {
			t.Fatalf("Expected %q but got %q for OutlookCategoryId", v.Expected.OutlookCategoryId, actual.OutlookCategoryId)
		}

	}
}

func TestParseUserIdOutlookMasterCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookMasterCategoryId
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
			Input: "/users/userIdValue/outlook",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/masterCategories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/mAsTeRcAtEgOrIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/outlook/masterCategories/outlookCategoryIdValue",
			Expected: &UserIdOutlookMasterCategoryId{
				UserId:            "userIdValue",
				OutlookCategoryId: "outlookCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/outlook/masterCategories/outlookCategoryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/mAsTeRcAtEgOrIeS/oUtLoOkCaTeGoRyIdVaLuE",
			Expected: &UserIdOutlookMasterCategoryId{
				UserId:            "uSeRiDvAlUe",
				OutlookCategoryId: "oUtLoOkCaTeGoRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/mAsTeRcAtEgOrIeS/oUtLoOkCaTeGoRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookMasterCategoryIDInsensitively(v.Input)
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

		if actual.OutlookCategoryId != v.Expected.OutlookCategoryId {
			t.Fatalf("Expected %q but got %q for OutlookCategoryId", v.Expected.OutlookCategoryId, actual.OutlookCategoryId)
		}

	}
}

func TestSegmentsForUserIdOutlookMasterCategoryId(t *testing.T) {
	segments := UserIdOutlookMasterCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOutlookMasterCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
