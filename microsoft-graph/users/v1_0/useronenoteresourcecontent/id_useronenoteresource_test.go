package useronenoteresourcecontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteResourceId{}

func TestNewUserOnenoteResourceID(t *testing.T) {
	id := NewUserOnenoteResourceID("userIdValue", "onenoteResourceIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OnenoteResourceId != "onenoteResourceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteResourceId'", id.OnenoteResourceId, "onenoteResourceIdValue")
	}
}

func TestFormatUserOnenoteResourceID(t *testing.T) {
	actual := NewUserOnenoteResourceID("userIdValue", "onenoteResourceIdValue").ID()
	expected := "/users/userIdValue/onenote/resources/onenoteResourceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnenoteResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteResourceId
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
			Input: "/users/userIdValue/onenote/resources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/resources/onenoteResourceIdValue",
			Expected: &UserOnenoteResourceId{
				UserId:            "userIdValue",
				OnenoteResourceId: "onenoteResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/resources/onenoteResourceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteResourceID(v.Input)
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

		if actual.OnenoteResourceId != v.Expected.OnenoteResourceId {
			t.Fatalf("Expected %q but got %q for OnenoteResourceId", v.Expected.OnenoteResourceId, actual.OnenoteResourceId)
		}

	}
}

func TestParseUserOnenoteResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteResourceId
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
			Input: "/users/userIdValue/onenote/resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/rEsOuRcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/resources/onenoteResourceIdValue",
			Expected: &UserOnenoteResourceId{
				UserId:            "userIdValue",
				OnenoteResourceId: "onenoteResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/resources/onenoteResourceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/rEsOuRcEs/oNeNoTeReSoUrCeIdVaLuE",
			Expected: &UserOnenoteResourceId{
				UserId:            "uSeRiDvAlUe",
				OnenoteResourceId: "oNeNoTeReSoUrCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/rEsOuRcEs/oNeNoTeReSoUrCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteResourceIDInsensitively(v.Input)
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

		if actual.OnenoteResourceId != v.Expected.OnenoteResourceId {
			t.Fatalf("Expected %q but got %q for OnenoteResourceId", v.Expected.OnenoteResourceId, actual.OnenoteResourceId)
		}

	}
}

func TestSegmentsForUserOnenoteResourceId(t *testing.T) {
	segments := UserOnenoteResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnenoteResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
