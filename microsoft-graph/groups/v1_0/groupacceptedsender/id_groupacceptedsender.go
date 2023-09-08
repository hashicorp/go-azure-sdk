package groupacceptedsender

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupAcceptedSenderId{}

// GroupAcceptedSenderId is a struct representing the Resource ID for a Group Accepted Sender
type GroupAcceptedSenderId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupAcceptedSenderID returns a new GroupAcceptedSenderId struct
func NewGroupAcceptedSenderID(groupId string, directoryObjectId string) GroupAcceptedSenderId {
	return GroupAcceptedSenderId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupAcceptedSenderID parses 'input' into a GroupAcceptedSenderId
func ParseGroupAcceptedSenderID(input string) (*GroupAcceptedSenderId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupAcceptedSenderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupAcceptedSenderId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseGroupAcceptedSenderIDInsensitively parses 'input' case-insensitively into a GroupAcceptedSenderId
// note: this method should only be used for API response data and not user input
func ParseGroupAcceptedSenderIDInsensitively(input string) (*GroupAcceptedSenderId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupAcceptedSenderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupAcceptedSenderId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateGroupAcceptedSenderID checks that 'input' can be parsed as a Group Accepted Sender ID
func ValidateGroupAcceptedSenderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupAcceptedSenderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Accepted Sender ID
func (id GroupAcceptedSenderId) ID() string {
	fmtString := "/groups/%s/acceptedSenders/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Accepted Sender ID
func (id GroupAcceptedSenderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticAcceptedSenders", "acceptedSenders", "acceptedSenders"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Group Accepted Sender ID
func (id GroupAcceptedSenderId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Accepted Sender (%s)", strings.Join(components, "\n"))
}
