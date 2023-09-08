package userinsightsharedresource

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInsightSharedId{}

func TestNewUserInsightSharedID(t *testing.T) {
	id := NewUserInsightSharedID("userIdValue", "sharedInsightIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.SharedInsightId != "sharedInsightIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedInsightId'", id.SharedInsightId, "sharedInsightIdValue")
	}
}

func TestFormatUserInsightSharedID(t *testing.T) {
	actual := NewUserInsightSharedID("userIdValue", "sharedInsightIdValue").ID()
	expected := "/users/userIdValue/insights/shared/sharedInsightIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserInsightSharedID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserInsightSharedId
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
			Input: "/users/userIdValue/insights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/insights/shared",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/insights/shared/sharedInsightIdValue",
			Expected: &UserInsightSharedId{
				UserId:          "userIdValue",
				SharedInsightId: "sharedInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/insights/shared/sharedInsightIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserInsightSharedID(v.Input)
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

		if actual.SharedInsightId != v.Expected.SharedInsightId {
			t.Fatalf("Expected %q but got %q for SharedInsightId", v.Expected.SharedInsightId, actual.SharedInsightId)
		}

	}
}

func TestParseUserInsightSharedIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserInsightSharedId
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
			Input: "/users/userIdValue/insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/insights/shared",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS/sHaReD",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/insights/shared/sharedInsightIdValue",
			Expected: &UserInsightSharedId{
				UserId:          "userIdValue",
				SharedInsightId: "sharedInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/insights/shared/sharedInsightIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS/sHaReD/sHaReDiNsIgHtIdVaLuE",
			Expected: &UserInsightSharedId{
				UserId:          "uSeRiDvAlUe",
				SharedInsightId: "sHaReDiNsIgHtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS/sHaReD/sHaReDiNsIgHtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserInsightSharedIDInsensitively(v.Input)
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

		if actual.SharedInsightId != v.Expected.SharedInsightId {
			t.Fatalf("Expected %q but got %q for SharedInsightId", v.Expected.SharedInsightId, actual.SharedInsightId)
		}

	}
}

func TestSegmentsForUserInsightSharedId(t *testing.T) {
	segments := UserInsightSharedId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserInsightSharedId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
