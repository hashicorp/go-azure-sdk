package insightused

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInsightUsedId{}

func TestNewUserIdInsightUsedID(t *testing.T) {
	id := NewUserIdInsightUsedID("userIdValue", "usedInsightIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.UsedInsightId != "usedInsightIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UsedInsightId'", id.UsedInsightId, "usedInsightIdValue")
	}
}

func TestFormatUserIdInsightUsedID(t *testing.T) {
	actual := NewUserIdInsightUsedID("userIdValue", "usedInsightIdValue").ID()
	expected := "/users/userIdValue/insights/used/usedInsightIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdInsightUsedID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInsightUsedId
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
			Input: "/users/userIdValue/insights/used",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/insights/used/usedInsightIdValue",
			Expected: &UserIdInsightUsedId{
				UserId:        "userIdValue",
				UsedInsightId: "usedInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/insights/used/usedInsightIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInsightUsedID(v.Input)
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

		if actual.UsedInsightId != v.Expected.UsedInsightId {
			t.Fatalf("Expected %q but got %q for UsedInsightId", v.Expected.UsedInsightId, actual.UsedInsightId)
		}

	}
}

func TestParseUserIdInsightUsedIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInsightUsedId
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
			Input: "/users/userIdValue/insights/used",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS/uSeD",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/insights/used/usedInsightIdValue",
			Expected: &UserIdInsightUsedId{
				UserId:        "userIdValue",
				UsedInsightId: "usedInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/insights/used/usedInsightIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS/uSeD/uSeDiNsIgHtIdVaLuE",
			Expected: &UserIdInsightUsedId{
				UserId:        "uSeRiDvAlUe",
				UsedInsightId: "uSeDiNsIgHtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS/uSeD/uSeDiNsIgHtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInsightUsedIDInsensitively(v.Input)
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

		if actual.UsedInsightId != v.Expected.UsedInsightId {
			t.Fatalf("Expected %q but got %q for UsedInsightId", v.Expected.UsedInsightId, actual.UsedInsightId)
		}

	}
}

func TestSegmentsForUserIdInsightUsedId(t *testing.T) {
	segments := UserIdInsightUsedId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdInsightUsedId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
