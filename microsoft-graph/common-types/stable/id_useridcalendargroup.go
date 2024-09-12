package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarGroupId{}

// UserIdCalendarGroupId is a struct representing the Resource ID for a User Id Calendar Group
type UserIdCalendarGroupId struct {
	UserId          string
	CalendarGroupId string
}

// NewUserIdCalendarGroupID returns a new UserIdCalendarGroupId struct
func NewUserIdCalendarGroupID(userId string, calendarGroupId string) UserIdCalendarGroupId {
	return UserIdCalendarGroupId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
	}
}

// ParseUserIdCalendarGroupID parses 'input' into a UserIdCalendarGroupId
func ParseUserIdCalendarGroupID(input string) (*UserIdCalendarGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarGroupIDInsensitively parses 'input' case-insensitively into a UserIdCalendarGroupId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarGroupIDInsensitively(input string) (*UserIdCalendarGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CalendarGroupId, ok = input.Parsed["calendarGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", input)
	}

	return nil
}

// ValidateUserIdCalendarGroupID checks that 'input' can be parsed as a User Id Calendar Group ID
func ValidateUserIdCalendarGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Group ID
func (id UserIdCalendarGroupId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Group ID
func (id UserIdCalendarGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Group ID
func (id UserIdCalendarGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
	}
	return fmt.Sprintf("User Id Calendar Group (%s)", strings.Join(components, "\n"))
}
