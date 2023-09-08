package grouponenotesection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteSectionId{}

// GroupOnenoteSectionId is a struct representing the Resource ID for a Group Onenote Section
type GroupOnenoteSectionId struct {
	GroupId          string
	OnenoteSectionId string
}

// NewGroupOnenoteSectionID returns a new GroupOnenoteSectionId struct
func NewGroupOnenoteSectionID(groupId string, onenoteSectionId string) GroupOnenoteSectionId {
	return GroupOnenoteSectionId{
		GroupId:          groupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupOnenoteSectionID parses 'input' into a GroupOnenoteSectionId
func ParseGroupOnenoteSectionID(input string) (*GroupOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteSectionIDInsensitively parses 'input' case-insensitively into a GroupOnenoteSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteSectionIDInsensitively(input string) (*GroupOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteSectionID checks that 'input' can be parsed as a Group Onenote Section ID
func ValidateGroupOnenoteSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Section ID
func (id GroupOnenoteSectionId) ID() string {
	fmtString := "/groups/%s/onenote/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Section ID
func (id GroupOnenoteSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Section ID
func (id GroupOnenoteSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Onenote Section (%s)", strings.Join(components, "\n"))
}
