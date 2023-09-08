package useroutlooktask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskId{}

// UserOutlookTaskId is a struct representing the Resource ID for a User Outlook Task
type UserOutlookTaskId struct {
	UserId        string
	OutlookTaskId string
}

// NewUserOutlookTaskID returns a new UserOutlookTaskId struct
func NewUserOutlookTaskID(userId string, outlookTaskId string) UserOutlookTaskId {
	return UserOutlookTaskId{
		UserId:        userId,
		OutlookTaskId: outlookTaskId,
	}
}

// ParseUserOutlookTaskID parses 'input' into a UserOutlookTaskId
func ParseUserOutlookTaskID(input string) (*UserOutlookTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	return &id, nil
}

// ParseUserOutlookTaskIDInsensitively parses 'input' case-insensitively into a UserOutlookTaskId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookTaskIDInsensitively(input string) (*UserOutlookTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	return &id, nil
}

// ValidateUserOutlookTaskID checks that 'input' can be parsed as a User Outlook Task ID
func ValidateUserOutlookTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Task ID
func (id UserOutlookTaskId) ID() string {
	fmtString := "/users/%s/outlook/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Task ID
func (id UserOutlookTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Task ID
func (id UserOutlookTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("User Outlook Task (%s)", strings.Join(components, "\n"))
}
