package groupeventinstancecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupEventInstanceId{}

// GroupEventInstanceId is a struct representing the Resource ID for a Group Event Instance
type GroupEventInstanceId struct {
	GroupId  string
	EventId  string
	EventId1 string
}

// NewGroupEventInstanceID returns a new GroupEventInstanceId struct
func NewGroupEventInstanceID(groupId string, eventId string, eventId1 string) GroupEventInstanceId {
	return GroupEventInstanceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseGroupEventInstanceID parses 'input' into a GroupEventInstanceId
func ParseGroupEventInstanceID(input string) (*GroupEventInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventInstanceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseGroupEventInstanceIDInsensitively parses 'input' case-insensitively into a GroupEventInstanceId
// note: this method should only be used for API response data and not user input
func ParseGroupEventInstanceIDInsensitively(input string) (*GroupEventInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventInstanceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupEventInstanceID checks that 'input' can be parsed as a Group Event Instance ID
func ValidateGroupEventInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupEventInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Event Instance ID
func (id GroupEventInstanceId) ID() string {
	fmtString := "/groups/%s/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Event Instance ID
func (id GroupEventInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Group Event Instance ID
func (id GroupEventInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Group Event Instance (%s)", strings.Join(components, "\n"))
}
