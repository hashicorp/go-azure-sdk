package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInsightSharedId{}

func TestNewUserIdInsightSharedID(t *testing.T) {
	id := NewUserIdInsightSharedID("userId", "sharedInsightId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.SharedInsightId != "sharedInsightId" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedInsightId'", id.SharedInsightId, "sharedInsightId")
	}
}

func TestFormatUserIdInsightSharedID(t *testing.T) {
	actual := NewUserIdInsightSharedID("userId", "sharedInsightId").ID()
	expected := "/users/userId/insights/shared/sharedInsightId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdInsightSharedID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInsightSharedId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/insights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/insights/shared",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/insights/shared/sharedInsightId",
			Expected: &UserIdInsightSharedId{
				UserId:          "userId",
				SharedInsightId: "sharedInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/insights/shared/sharedInsightId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInsightSharedID(v.Input)
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

func TestParseUserIdInsightSharedIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInsightSharedId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/insights/shared",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNsIgHtS/sHaReD",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/insights/shared/sharedInsightId",
			Expected: &UserIdInsightSharedId{
				UserId:          "userId",
				SharedInsightId: "sharedInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/insights/shared/sharedInsightId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNsIgHtS/sHaReD/sHaReDiNsIgHtId",
			Expected: &UserIdInsightSharedId{
				UserId:          "uSeRiD",
				SharedInsightId: "sHaReDiNsIgHtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNsIgHtS/sHaReD/sHaReDiNsIgHtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInsightSharedIDInsensitively(v.Input)
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

func TestSegmentsForUserIdInsightSharedId(t *testing.T) {
	segments := UserIdInsightSharedId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdInsightSharedId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
