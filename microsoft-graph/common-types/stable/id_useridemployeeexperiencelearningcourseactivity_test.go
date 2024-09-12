package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEmployeeExperienceLearningCourseActivityId{}

func TestNewUserIdEmployeeExperienceLearningCourseActivityID(t *testing.T) {
	id := NewUserIdEmployeeExperienceLearningCourseActivityID("userIdValue", "learningCourseActivityIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.LearningCourseActivityId != "learningCourseActivityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LearningCourseActivityId'", id.LearningCourseActivityId, "learningCourseActivityIdValue")
	}
}

func TestFormatUserIdEmployeeExperienceLearningCourseActivityID(t *testing.T) {
	actual := NewUserIdEmployeeExperienceLearningCourseActivityID("userIdValue", "learningCourseActivityIdValue").ID()
	expected := "/users/userIdValue/employeeExperience/learningCourseActivities/learningCourseActivityIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdEmployeeExperienceLearningCourseActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEmployeeExperienceLearningCourseActivityId
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
			Input: "/users/userIdValue/employeeExperience",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/employeeExperience/learningCourseActivities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/employeeExperience/learningCourseActivities/learningCourseActivityIdValue",
			Expected: &UserIdEmployeeExperienceLearningCourseActivityId{
				UserId:                   "userIdValue",
				LearningCourseActivityId: "learningCourseActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/employeeExperience/learningCourseActivities/learningCourseActivityIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEmployeeExperienceLearningCourseActivityID(v.Input)
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

		if actual.LearningCourseActivityId != v.Expected.LearningCourseActivityId {
			t.Fatalf("Expected %q but got %q for LearningCourseActivityId", v.Expected.LearningCourseActivityId, actual.LearningCourseActivityId)
		}

	}
}

func TestParseUserIdEmployeeExperienceLearningCourseActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEmployeeExperienceLearningCourseActivityId
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
			Input: "/users/userIdValue/employeeExperience",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eMpLoYeEeXpErIeNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/employeeExperience/learningCourseActivities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eMpLoYeEeXpErIeNcE/lEaRnInGcOuRsEaCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/employeeExperience/learningCourseActivities/learningCourseActivityIdValue",
			Expected: &UserIdEmployeeExperienceLearningCourseActivityId{
				UserId:                   "userIdValue",
				LearningCourseActivityId: "learningCourseActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/employeeExperience/learningCourseActivities/learningCourseActivityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eMpLoYeEeXpErIeNcE/lEaRnInGcOuRsEaCtIvItIeS/lEaRnInGcOuRsEaCtIvItYiDvAlUe",
			Expected: &UserIdEmployeeExperienceLearningCourseActivityId{
				UserId:                   "uSeRiDvAlUe",
				LearningCourseActivityId: "lEaRnInGcOuRsEaCtIvItYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eMpLoYeEeXpErIeNcE/lEaRnInGcOuRsEaCtIvItIeS/lEaRnInGcOuRsEaCtIvItYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEmployeeExperienceLearningCourseActivityIDInsensitively(v.Input)
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

		if actual.LearningCourseActivityId != v.Expected.LearningCourseActivityId {
			t.Fatalf("Expected %q but got %q for LearningCourseActivityId", v.Expected.LearningCourseActivityId, actual.LearningCourseActivityId)
		}

	}
}

func TestSegmentsForUserIdEmployeeExperienceLearningCourseActivityId(t *testing.T) {
	segments := UserIdEmployeeExperienceLearningCourseActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdEmployeeExperienceLearningCourseActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
