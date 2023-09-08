package groupsitelistcolumn

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListColumnId{}

// GroupSiteListColumnId is a struct representing the Resource ID for a Group Site List Column
type GroupSiteListColumnId struct {
	GroupId            string
	SiteId             string
	ListId             string
	ColumnDefinitionId string
}

// NewGroupSiteListColumnID returns a new GroupSiteListColumnId struct
func NewGroupSiteListColumnID(groupId string, siteId string, listId string, columnDefinitionId string) GroupSiteListColumnId {
	return GroupSiteListColumnId{
		GroupId:            groupId,
		SiteId:             siteId,
		ListId:             listId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupSiteListColumnID parses 'input' into a GroupSiteListColumnId
func ParseGroupSiteListColumnID(input string) (*GroupSiteListColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListColumnId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.ColumnDefinitionId, ok = parsed.Parsed["columnDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListColumnIDInsensitively parses 'input' case-insensitively into a GroupSiteListColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListColumnIDInsensitively(input string) (*GroupSiteListColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListColumnId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.ColumnDefinitionId, ok = parsed.Parsed["columnDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListColumnID checks that 'input' can be parsed as a Group Site List Column ID
func ValidateGroupSiteListColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Column ID
func (id GroupSiteListColumnId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Column ID
func (id GroupSiteListColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticColumns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Group Site List Column ID
func (id GroupSiteListColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Site List Column (%s)", strings.Join(components, "\n"))
}
