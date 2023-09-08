package grouponenoteoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteOperationId{}

// GroupOnenoteOperationId is a struct representing the Resource ID for a Group Onenote Operation
type GroupOnenoteOperationId struct {
	GroupId            string
	OnenoteOperationId string
}

// NewGroupOnenoteOperationID returns a new GroupOnenoteOperationId struct
func NewGroupOnenoteOperationID(groupId string, onenoteOperationId string) GroupOnenoteOperationId {
	return GroupOnenoteOperationId{
		GroupId:            groupId,
		OnenoteOperationId: onenoteOperationId,
	}
}

// ParseGroupOnenoteOperationID parses 'input' into a GroupOnenoteOperationId
func ParseGroupOnenoteOperationID(input string) (*GroupOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenoteOperationId, ok = parsed.Parsed["onenoteOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteOperationIDInsensitively parses 'input' case-insensitively into a GroupOnenoteOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteOperationIDInsensitively(input string) (*GroupOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenoteOperationId, ok = parsed.Parsed["onenoteOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteOperationID checks that 'input' can be parsed as a Group Onenote Operation ID
func ValidateGroupOnenoteOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Operation ID
func (id GroupOnenoteOperationId) ID() string {
	fmtString := "/groups/%s/onenote/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenoteOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Operation ID
func (id GroupOnenoteOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("onenoteOperationId", "onenoteOperationIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Operation ID
func (id GroupOnenoteOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Operation: %q", id.OnenoteOperationId),
	}
	return fmt.Sprintf("Group Onenote Operation (%s)", strings.Join(components, "\n"))
}
