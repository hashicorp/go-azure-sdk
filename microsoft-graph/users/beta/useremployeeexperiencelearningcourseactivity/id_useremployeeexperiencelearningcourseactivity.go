package useremployeeexperiencelearningcourseactivity

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserEmployeeExperienceLearningCourseActivityId{}

// UserEmployeeExperienceLearningCourseActivityId is a struct representing the Resource ID for a User Employee Experience Learning Course Activity
type UserEmployeeExperienceLearningCourseActivityId struct {
	UserId                   string
	LearningCourseActivityId string
}

// NewUserEmployeeExperienceLearningCourseActivityID returns a new UserEmployeeExperienceLearningCourseActivityId struct
func NewUserEmployeeExperienceLearningCourseActivityID(userId string, learningCourseActivityId string) UserEmployeeExperienceLearningCourseActivityId {
	return UserEmployeeExperienceLearningCourseActivityId{
		UserId:                   userId,
		LearningCourseActivityId: learningCourseActivityId,
	}
}

// ParseUserEmployeeExperienceLearningCourseActivityID parses 'input' into a UserEmployeeExperienceLearningCourseActivityId
func ParseUserEmployeeExperienceLearningCourseActivityID(input string) (*UserEmployeeExperienceLearningCourseActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEmployeeExperienceLearningCourseActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEmployeeExperienceLearningCourseActivityId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.LearningCourseActivityId, ok = parsed.Parsed["learningCourseActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "learningCourseActivityId", *parsed)
	}

	return &id, nil
}

// ParseUserEmployeeExperienceLearningCourseActivityIDInsensitively parses 'input' case-insensitively into a UserEmployeeExperienceLearningCourseActivityId
// note: this method should only be used for API response data and not user input
func ParseUserEmployeeExperienceLearningCourseActivityIDInsensitively(input string) (*UserEmployeeExperienceLearningCourseActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEmployeeExperienceLearningCourseActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEmployeeExperienceLearningCourseActivityId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.LearningCourseActivityId, ok = parsed.Parsed["learningCourseActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "learningCourseActivityId", *parsed)
	}

	return &id, nil
}

// ValidateUserEmployeeExperienceLearningCourseActivityID checks that 'input' can be parsed as a User Employee Experience Learning Course Activity ID
func ValidateUserEmployeeExperienceLearningCourseActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserEmployeeExperienceLearningCourseActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Employee Experience Learning Course Activity ID
func (id UserEmployeeExperienceLearningCourseActivityId) ID() string {
	fmtString := "/users/%s/employeeExperience/learningCourseActivities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.LearningCourseActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Employee Experience Learning Course Activity ID
func (id UserEmployeeExperienceLearningCourseActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticEmployeeExperience", "employeeExperience", "employeeExperience"),
		resourceids.StaticSegment("staticLearningCourseActivities", "learningCourseActivities", "learningCourseActivities"),
		resourceids.UserSpecifiedSegment("learningCourseActivityId", "learningCourseActivityIdValue"),
	}
}

// String returns a human-readable description of this User Employee Experience Learning Course Activity ID
func (id UserEmployeeExperienceLearningCourseActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Learning Course Activity: %q", id.LearningCourseActivityId),
	}
	return fmt.Sprintf("User Employee Experience Learning Course Activity (%s)", strings.Join(components, "\n"))
}
