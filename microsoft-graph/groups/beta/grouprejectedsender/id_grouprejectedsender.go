package grouprejectedsender

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupRejectedSenderId{}

// GroupRejectedSenderId is a struct representing the Resource ID for a Group Rejected Sender
type GroupRejectedSenderId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupRejectedSenderID returns a new GroupRejectedSenderId struct
func NewGroupRejectedSenderID(groupId string, directoryObjectId string) GroupRejectedSenderId {
	return GroupRejectedSenderId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupRejectedSenderID parses 'input' into a GroupRejectedSenderId
func ParseGroupRejectedSenderID(input string) (*GroupRejectedSenderId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupRejectedSenderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupRejectedSenderId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseGroupRejectedSenderIDInsensitively parses 'input' case-insensitively into a GroupRejectedSenderId
// note: this method should only be used for API response data and not user input
func ParseGroupRejectedSenderIDInsensitively(input string) (*GroupRejectedSenderId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupRejectedSenderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupRejectedSenderId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateGroupRejectedSenderID checks that 'input' can be parsed as a Group Rejected Sender ID
func ValidateGroupRejectedSenderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupRejectedSenderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Rejected Sender ID
func (id GroupRejectedSenderId) ID() string {
	fmtString := "/groups/%s/rejectedSenders/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Rejected Sender ID
func (id GroupRejectedSenderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticRejectedSenders", "rejectedSenders", "rejectedSenders"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Group Rejected Sender ID
func (id GroupRejectedSenderId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Rejected Sender (%s)", strings.Join(components, "\n"))
}
