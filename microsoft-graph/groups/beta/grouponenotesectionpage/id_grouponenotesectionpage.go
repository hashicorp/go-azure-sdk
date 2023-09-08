package grouponenotesectionpage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteSectionPageId{}

// GroupOnenoteSectionPageId is a struct representing the Resource ID for a Group Onenote Section Page
type GroupOnenoteSectionPageId struct {
	GroupId          string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupOnenoteSectionPageID returns a new GroupOnenoteSectionPageId struct
func NewGroupOnenoteSectionPageID(groupId string, onenoteSectionId string, onenotePageId string) GroupOnenoteSectionPageId {
	return GroupOnenoteSectionPageId{
		GroupId:          groupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupOnenoteSectionPageID parses 'input' into a GroupOnenoteSectionPageId
func ParseGroupOnenoteSectionPageID(input string) (*GroupOnenoteSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionPageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteSectionPageIDInsensitively parses 'input' case-insensitively into a GroupOnenoteSectionPageId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteSectionPageIDInsensitively(input string) (*GroupOnenoteSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionPageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteSectionPageID checks that 'input' can be parsed as a Group Onenote Section Page ID
func ValidateGroupOnenoteSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Section Page ID
func (id GroupOnenoteSectionPageId) ID() string {
	fmtString := "/groups/%s/onenote/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Section Page ID
func (id GroupOnenoteSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Section Page ID
func (id GroupOnenoteSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Onenote Section Page (%s)", strings.Join(components, "\n"))
}
