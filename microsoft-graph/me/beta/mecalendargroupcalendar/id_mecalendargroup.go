package mecalendargroupcalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupId{}

// MeCalendarGroupId is a struct representing the Resource ID for a Me Calendar Group
type MeCalendarGroupId struct {
	CalendarGroupId string
}

// NewMeCalendarGroupID returns a new MeCalendarGroupId struct
func NewMeCalendarGroupID(calendarGroupId string) MeCalendarGroupId {
	return MeCalendarGroupId{
		CalendarGroupId: calendarGroupId,
	}
}

// ParseMeCalendarGroupID parses 'input' into a MeCalendarGroupId
func ParseMeCalendarGroupID(input string) (*MeCalendarGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarGroupIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIDInsensitively(input string) (*MeCalendarGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarGroupID checks that 'input' can be parsed as a Me Calendar Group ID
func ValidateMeCalendarGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group ID
func (id MeCalendarGroupId) ID() string {
	fmtString := "/me/calendarGroups/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group ID
func (id MeCalendarGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group ID
func (id MeCalendarGroupId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
	}
	return fmt.Sprintf("Me Calendar Group (%s)", strings.Join(components, "\n"))
}
