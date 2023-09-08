package groupevent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupEventId{}

// GroupEventId is a struct representing the Resource ID for a Group Event
type GroupEventId struct {
	GroupId string
	EventId string
}

// NewGroupEventID returns a new GroupEventId struct
func NewGroupEventID(groupId string, eventId string) GroupEventId {
	return GroupEventId{
		GroupId: groupId,
		EventId: eventId,
	}
}

// ParseGroupEventID parses 'input' into a GroupEventId
func ParseGroupEventID(input string) (*GroupEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseGroupEventIDInsensitively parses 'input' case-insensitively into a GroupEventId
// note: this method should only be used for API response data and not user input
func ParseGroupEventIDInsensitively(input string) (*GroupEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateGroupEventID checks that 'input' can be parsed as a Group Event ID
func ValidateGroupEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Event ID
func (id GroupEventId) ID() string {
	fmtString := "/groups/%s/events/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Event ID
func (id GroupEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Group Event ID
func (id GroupEventId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Group Event (%s)", strings.Join(components, "\n"))
}
