package groupsitecolumnsourcecolumn

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteColumnId{}

// GroupSiteColumnId is a struct representing the Resource ID for a Group Site Column
type GroupSiteColumnId struct {
	GroupId            string
	SiteId             string
	ColumnDefinitionId string
}

// NewGroupSiteColumnID returns a new GroupSiteColumnId struct
func NewGroupSiteColumnID(groupId string, siteId string, columnDefinitionId string) GroupSiteColumnId {
	return GroupSiteColumnId{
		GroupId:            groupId,
		SiteId:             siteId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupSiteColumnID parses 'input' into a GroupSiteColumnId
func ParseGroupSiteColumnID(input string) (*GroupSiteColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteColumnId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ColumnDefinitionId, ok = parsed.Parsed["columnDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteColumnIDInsensitively parses 'input' case-insensitively into a GroupSiteColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteColumnIDInsensitively(input string) (*GroupSiteColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteColumnId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ColumnDefinitionId, ok = parsed.Parsed["columnDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteColumnID checks that 'input' can be parsed as a Group Site Column ID
func ValidateGroupSiteColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Column ID
func (id GroupSiteColumnId) ID() string {
	fmtString := "/groups/%s/sites/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Column ID
func (id GroupSiteColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticColumns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Group Site Column ID
func (id GroupSiteColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Site Column (%s)", strings.Join(components, "\n"))
}
