package grouponenotepageparentsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenotePageId{}

// GroupOnenotePageId is a struct representing the Resource ID for a Group Onenote Page
type GroupOnenotePageId struct {
	GroupId       string
	OnenotePageId string
}

// NewGroupOnenotePageID returns a new GroupOnenotePageId struct
func NewGroupOnenotePageID(groupId string, onenotePageId string) GroupOnenotePageId {
	return GroupOnenotePageId{
		GroupId:       groupId,
		OnenotePageId: onenotePageId,
	}
}

// ParseGroupOnenotePageID parses 'input' into a GroupOnenotePageId
func ParseGroupOnenotePageID(input string) (*GroupOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenotePageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenotePageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenotePageIDInsensitively parses 'input' case-insensitively into a GroupOnenotePageId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenotePageIDInsensitively(input string) (*GroupOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenotePageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenotePageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenotePageID checks that 'input' can be parsed as a Group Onenote Page ID
func ValidateGroupOnenotePageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenotePageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Page ID
func (id GroupOnenotePageId) ID() string {
	fmtString := "/groups/%s/onenote/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Page ID
func (id GroupOnenotePageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Page ID
func (id GroupOnenotePageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Onenote Page (%s)", strings.Join(components, "\n"))
}
