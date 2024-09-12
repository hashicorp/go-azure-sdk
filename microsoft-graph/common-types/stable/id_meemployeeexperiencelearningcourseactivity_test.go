package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEmployeeExperienceLearningCourseActivityId{}

func TestNewMeEmployeeExperienceLearningCourseActivityID(t *testing.T) {
	id := NewMeEmployeeExperienceLearningCourseActivityID("learningCourseActivityIdValue")

	if id.LearningCourseActivityId != "learningCourseActivityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LearningCourseActivityId'", id.LearningCourseActivityId, "learningCourseActivityIdValue")
	}
}

func TestFormatMeEmployeeExperienceLearningCourseActivityID(t *testing.T) {
	actual := NewMeEmployeeExperienceLearningCourseActivityID("learningCourseActivityIdValue").ID()
	expected := "/me/employeeExperience/learningCourseActivities/learningCourseActivityIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeEmployeeExperienceLearningCourseActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEmployeeExperienceLearningCourseActivityId
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
			Input: "/me/employeeExperience",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/employeeExperience/learningCourseActivities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/employeeExperience/learningCourseActivities/learningCourseActivityIdValue",
			Expected: &MeEmployeeExperienceLearningCourseActivityId{
				LearningCourseActivityId: "learningCourseActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/employeeExperience/learningCourseActivities/learningCourseActivityIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEmployeeExperienceLearningCourseActivityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LearningCourseActivityId != v.Expected.LearningCourseActivityId {
			t.Fatalf("Expected %q but got %q for LearningCourseActivityId", v.Expected.LearningCourseActivityId, actual.LearningCourseActivityId)
		}

	}
}

func TestParseMeEmployeeExperienceLearningCourseActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEmployeeExperienceLearningCourseActivityId
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
			Input: "/me/employeeExperience",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/employeeExperience/learningCourseActivities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/lEaRnInGcOuRsEaCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/employeeExperience/learningCourseActivities/learningCourseActivityIdValue",
			Expected: &MeEmployeeExperienceLearningCourseActivityId{
				LearningCourseActivityId: "learningCourseActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/employeeExperience/learningCourseActivities/learningCourseActivityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/lEaRnInGcOuRsEaCtIvItIeS/lEaRnInGcOuRsEaCtIvItYiDvAlUe",
			Expected: &MeEmployeeExperienceLearningCourseActivityId{
				LearningCourseActivityId: "lEaRnInGcOuRsEaCtIvItYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/lEaRnInGcOuRsEaCtIvItIeS/lEaRnInGcOuRsEaCtIvItYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEmployeeExperienceLearningCourseActivityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LearningCourseActivityId != v.Expected.LearningCourseActivityId {
			t.Fatalf("Expected %q but got %q for LearningCourseActivityId", v.Expected.LearningCourseActivityId, actual.LearningCourseActivityId)
		}

	}
}

func TestSegmentsForMeEmployeeExperienceLearningCourseActivityId(t *testing.T) {
	segments := MeEmployeeExperienceLearningCourseActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeEmployeeExperienceLearningCourseActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
