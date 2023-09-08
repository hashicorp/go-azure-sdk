package groupsitelistcontenttypecolumnposition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListContentTypeColumnPositionId{}

// GroupSiteListContentTypeColumnPositionId is a struct representing the Resource ID for a Group Site List Content Type Column Position
type GroupSiteListContentTypeColumnPositionId struct {
	GroupId            string
	SiteId             string
	ListId             string
	ContentTypeId      string
	ColumnDefinitionId string
}

// NewGroupSiteListContentTypeColumnPositionID returns a new GroupSiteListContentTypeColumnPositionId struct
func NewGroupSiteListContentTypeColumnPositionID(groupId string, siteId string, listId string, contentTypeId string, columnDefinitionId string) GroupSiteListContentTypeColumnPositionId {
	return GroupSiteListContentTypeColumnPositionId{
		GroupId:            groupId,
		SiteId:             siteId,
		ListId:             listId,
		ContentTypeId:      contentTypeId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupSiteListContentTypeColumnPositionID parses 'input' into a GroupSiteListContentTypeColumnPositionId
func ParseGroupSiteListContentTypeColumnPositionID(input string) (*GroupSiteListContentTypeColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListContentTypeColumnPositionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListContentTypeColumnPositionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", *parsed)
	}

	if id.ColumnDefinitionId, ok = parsed.Parsed["columnDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListContentTypeColumnPositionIDInsensitively parses 'input' case-insensitively into a GroupSiteListContentTypeColumnPositionId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListContentTypeColumnPositionIDInsensitively(input string) (*GroupSiteListContentTypeColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListContentTypeColumnPositionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListContentTypeColumnPositionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", *parsed)
	}

	if id.ColumnDefinitionId, ok = parsed.Parsed["columnDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListContentTypeColumnPositionID checks that 'input' can be parsed as a Group Site List Content Type Column Position ID
func ValidateGroupSiteListContentTypeColumnPositionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListContentTypeColumnPositionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Content Type Column Position ID
func (id GroupSiteListContentTypeColumnPositionId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/contentTypes/%s/columnPositions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ContentTypeId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Content Type Column Position ID
func (id GroupSiteListContentTypeColumnPositionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
		resourceids.StaticSegment("staticColumnPositions", "columnPositions", "columnPositions"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Group Site List Content Type Column Position ID
func (id GroupSiteListContentTypeColumnPositionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Site List Content Type Column Position (%s)", strings.Join(components, "\n"))
}
