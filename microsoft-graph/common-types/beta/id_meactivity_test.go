package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeActivityId{}

func TestNewMeActivityID(t *testing.T) {
	id := NewMeActivityID("userActivityId")

	if id.UserActivityId != "userActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserActivityId'", id.UserActivityId, "userActivityId")
	}
}

func TestFormatMeActivityID(t *testing.T) {
	actual := NewMeActivityID("userActivityId").ID()
	expected := "/me/activities/userActivityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeActivityId
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
			Input: "/me/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/activities/userActivityId",
			Expected: &MeActivityId{
				UserActivityId: "userActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/activities/userActivityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeActivityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserActivityId != v.Expected.UserActivityId {
			t.Fatalf("Expected %q but got %q for UserActivityId", v.Expected.UserActivityId, actual.UserActivityId)
		}

	}
}

func TestParseMeActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeActivityId
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
			Input: "/me/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/activities/userActivityId",
			Expected: &MeActivityId{
				UserActivityId: "userActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/activities/userActivityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aCtIvItIeS/uSeRaCtIvItYiD",
			Expected: &MeActivityId{
				UserActivityId: "uSeRaCtIvItYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aCtIvItIeS/uSeRaCtIvItYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeActivityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserActivityId != v.Expected.UserActivityId {
			t.Fatalf("Expected %q but got %q for UserActivityId", v.Expected.UserActivityId, actual.UserActivityId)
		}

	}
}

func TestSegmentsForMeActivityId(t *testing.T) {
	segments := MeActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
