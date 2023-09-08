package grouponenoteresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteResourceId{}

// GroupOnenoteResourceId is a struct representing the Resource ID for a Group Onenote Resource
type GroupOnenoteResourceId struct {
	GroupId           string
	OnenoteResourceId string
}

// NewGroupOnenoteResourceID returns a new GroupOnenoteResourceId struct
func NewGroupOnenoteResourceID(groupId string, onenoteResourceId string) GroupOnenoteResourceId {
	return GroupOnenoteResourceId{
		GroupId:           groupId,
		OnenoteResourceId: onenoteResourceId,
	}
}

// ParseGroupOnenoteResourceID parses 'input' into a GroupOnenoteResourceId
func ParseGroupOnenoteResourceID(input string) (*GroupOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteResourceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenoteResourceId, ok = parsed.Parsed["onenoteResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteResourceIDInsensitively parses 'input' case-insensitively into a GroupOnenoteResourceId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteResourceIDInsensitively(input string) (*GroupOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteResourceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenoteResourceId, ok = parsed.Parsed["onenoteResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteResourceID checks that 'input' can be parsed as a Group Onenote Resource ID
func ValidateGroupOnenoteResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Resource ID
func (id GroupOnenoteResourceId) ID() string {
	fmtString := "/groups/%s/onenote/resources/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenoteResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Resource ID
func (id GroupOnenoteResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticResources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("onenoteResourceId", "onenoteResourceIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Resource ID
func (id GroupOnenoteResourceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Resource: %q", id.OnenoteResourceId),
	}
	return fmt.Sprintf("Group Onenote Resource (%s)", strings.Join(components, "\n"))
}
