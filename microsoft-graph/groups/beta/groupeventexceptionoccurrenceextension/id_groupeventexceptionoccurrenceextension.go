package groupeventexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupEventExceptionOccurrenceExtensionId{}

// GroupEventExceptionOccurrenceExtensionId is a struct representing the Resource ID for a Group Event Exception Occurrence Extension
type GroupEventExceptionOccurrenceExtensionId struct {
	GroupId     string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewGroupEventExceptionOccurrenceExtensionID returns a new GroupEventExceptionOccurrenceExtensionId struct
func NewGroupEventExceptionOccurrenceExtensionID(groupId string, eventId string, eventId1 string, extensionId string) GroupEventExceptionOccurrenceExtensionId {
	return GroupEventExceptionOccurrenceExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseGroupEventExceptionOccurrenceExtensionID parses 'input' into a GroupEventExceptionOccurrenceExtensionId
func ParseGroupEventExceptionOccurrenceExtensionID(input string) (*GroupEventExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventExceptionOccurrenceExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseGroupEventExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a GroupEventExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupEventExceptionOccurrenceExtensionIDInsensitively(input string) (*GroupEventExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventExceptionOccurrenceExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupEventExceptionOccurrenceExtensionID checks that 'input' can be parsed as a Group Event Exception Occurrence Extension ID
func ValidateGroupEventExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupEventExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Event Exception Occurrence Extension ID
func (id GroupEventExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/groups/%s/events/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Event Exception Occurrence Extension ID
func (id GroupEventExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Event Exception Occurrence Extension ID
func (id GroupEventExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Event Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
