package calendarcalendarpermission

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarPermissionId{}

// MeCalendarCalendarPermissionId is a struct representing the Resource ID for a Me Calendar Calendar Permission
type MeCalendarCalendarPermissionId struct {
	CalendarPermissionId string
}

// NewMeCalendarCalendarPermissionID returns a new MeCalendarCalendarPermissionId struct
func NewMeCalendarCalendarPermissionID(calendarPermissionId string) MeCalendarCalendarPermissionId {
	return MeCalendarCalendarPermissionId{
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseMeCalendarCalendarPermissionID parses 'input' into a MeCalendarCalendarPermissionId
func ParseMeCalendarCalendarPermissionID(input string) (*MeCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarPermissionIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarPermissionIDInsensitively(input string) (*MeCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarPermissionId, ok = input.Parsed["calendarPermissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", input)
	}

	return nil
}

// ValidateMeCalendarCalendarPermissionID checks that 'input' can be parsed as a Me Calendar Calendar Permission ID
func ValidateMeCalendarCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar Permission ID
func (id MeCalendarCalendarPermissionId) ID() string {
	fmtString := "/me/calendar/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar Permission ID
func (id MeCalendarCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar Permission ID
func (id MeCalendarCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("Me Calendar Calendar Permission (%s)", strings.Join(components, "\n"))
}
