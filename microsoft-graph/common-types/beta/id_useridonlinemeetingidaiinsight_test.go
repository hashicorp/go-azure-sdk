package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdAiInsightId{}

func TestNewUserIdOnlineMeetingIdAiInsightID(t *testing.T) {
	id := NewUserIdOnlineMeetingIdAiInsightID("userId", "onlineMeetingId", "callAiInsightId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.CallAiInsightId != "callAiInsightId" {
		t.Fatalf("Expected %q but got %q for Segment 'CallAiInsightId'", id.CallAiInsightId, "callAiInsightId")
	}
}

func TestFormatUserIdOnlineMeetingIdAiInsightID(t *testing.T) {
	actual := NewUserIdOnlineMeetingIdAiInsightID("userId", "onlineMeetingId", "callAiInsightId").ID()
	expected := "/users/userId/onlineMeetings/onlineMeetingId/aiInsights/callAiInsightId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnlineMeetingIdAiInsightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdAiInsightId
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
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/aiInsights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/aiInsights/callAiInsightId",
			Expected: &UserIdOnlineMeetingIdAiInsightId{
				UserId:          "userId",
				OnlineMeetingId: "onlineMeetingId",
				CallAiInsightId: "callAiInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/aiInsights/callAiInsightId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdAiInsightID(v.Input)
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

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.CallAiInsightId != v.Expected.CallAiInsightId {
			t.Fatalf("Expected %q but got %q for CallAiInsightId", v.Expected.CallAiInsightId, actual.CallAiInsightId)
		}

	}
}

func TestParseUserIdOnlineMeetingIdAiInsightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdAiInsightId
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
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/aiInsights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aIiNsIgHtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/aiInsights/callAiInsightId",
			Expected: &UserIdOnlineMeetingIdAiInsightId{
				UserId:          "userId",
				OnlineMeetingId: "onlineMeetingId",
				CallAiInsightId: "callAiInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/aiInsights/callAiInsightId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aIiNsIgHtS/cAlLaIiNsIgHtId",
			Expected: &UserIdOnlineMeetingIdAiInsightId{
				UserId:          "uSeRiD",
				OnlineMeetingId: "oNlInEmEeTiNgId",
				CallAiInsightId: "cAlLaIiNsIgHtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aIiNsIgHtS/cAlLaIiNsIgHtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdAiInsightIDInsensitively(v.Input)
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

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.CallAiInsightId != v.Expected.CallAiInsightId {
			t.Fatalf("Expected %q but got %q for CallAiInsightId", v.Expected.CallAiInsightId, actual.CallAiInsightId)
		}

	}
}

func TestSegmentsForUserIdOnlineMeetingIdAiInsightId(t *testing.T) {
	segments := UserIdOnlineMeetingIdAiInsightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnlineMeetingIdAiInsightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
