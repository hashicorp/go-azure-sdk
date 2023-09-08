package useroutlooktaskgrouptaskfolder

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskGroupId{}

// UserOutlookTaskGroupId is a struct representing the Resource ID for a User Outlook Task Group
type UserOutlookTaskGroupId struct {
	UserId             string
	OutlookTaskGroupId string
}

// NewUserOutlookTaskGroupID returns a new UserOutlookTaskGroupId struct
func NewUserOutlookTaskGroupID(userId string, outlookTaskGroupId string) UserOutlookTaskGroupId {
	return UserOutlookTaskGroupId{
		UserId:             userId,
		OutlookTaskGroupId: outlookTaskGroupId,
	}
}

// ParseUserOutlookTaskGroupID parses 'input' into a UserOutlookTaskGroupId
func ParseUserOutlookTaskGroupID(input string) (*UserOutlookTaskGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskGroupId, ok = parsed.Parsed["outlookTaskGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", *parsed)
	}

	return &id, nil
}

// ParseUserOutlookTaskGroupIDInsensitively parses 'input' case-insensitively into a UserOutlookTaskGroupId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookTaskGroupIDInsensitively(input string) (*UserOutlookTaskGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskGroupId, ok = parsed.Parsed["outlookTaskGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", *parsed)
	}

	return &id, nil
}

// ValidateUserOutlookTaskGroupID checks that 'input' can be parsed as a User Outlook Task Group ID
func ValidateUserOutlookTaskGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookTaskGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Task Group ID
func (id UserOutlookTaskGroupId) ID() string {
	fmtString := "/users/%s/outlook/taskGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Task Group ID
func (id UserOutlookTaskGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Task Group ID
func (id UserOutlookTaskGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
	}
	return fmt.Sprintf("User Outlook Task Group (%s)", strings.Join(components, "\n"))
}
