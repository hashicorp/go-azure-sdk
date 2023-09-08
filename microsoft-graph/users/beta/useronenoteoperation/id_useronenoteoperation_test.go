package useronenoteoperation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteOperationId{}

func TestNewUserOnenoteOperationID(t *testing.T) {
	id := NewUserOnenoteOperationID("userIdValue", "onenoteOperationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OnenoteOperationId != "onenoteOperationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteOperationId'", id.OnenoteOperationId, "onenoteOperationIdValue")
	}
}

func TestFormatUserOnenoteOperationID(t *testing.T) {
	actual := NewUserOnenoteOperationID("userIdValue", "onenoteOperationIdValue").ID()
	expected := "/users/userIdValue/onenote/operations/onenoteOperationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnenoteOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteOperationId
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
			Input: "/users/userIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/operations/onenoteOperationIdValue",
			Expected: &UserOnenoteOperationId{
				UserId:             "userIdValue",
				OnenoteOperationId: "onenoteOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/operations/onenoteOperationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteOperationID(v.Input)
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

		if actual.OnenoteOperationId != v.Expected.OnenoteOperationId {
			t.Fatalf("Expected %q but got %q for OnenoteOperationId", v.Expected.OnenoteOperationId, actual.OnenoteOperationId)
		}

	}
}

func TestParseUserOnenoteOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteOperationId
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
			Input: "/users/userIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/operations/onenoteOperationIdValue",
			Expected: &UserOnenoteOperationId{
				UserId:             "userIdValue",
				OnenoteOperationId: "onenoteOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/operations/onenoteOperationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/oPeRaTiOnS/oNeNoTeOpErAtIoNiDvAlUe",
			Expected: &UserOnenoteOperationId{
				UserId:             "uSeRiDvAlUe",
				OnenoteOperationId: "oNeNoTeOpErAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/oPeRaTiOnS/oNeNoTeOpErAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteOperationIDInsensitively(v.Input)
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

		if actual.OnenoteOperationId != v.Expected.OnenoteOperationId {
			t.Fatalf("Expected %q but got %q for OnenoteOperationId", v.Expected.OnenoteOperationId, actual.OnenoteOperationId)
		}

	}
}

func TestSegmentsForUserOnenoteOperationId(t *testing.T) {
	segments := UserOnenoteOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnenoteOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
