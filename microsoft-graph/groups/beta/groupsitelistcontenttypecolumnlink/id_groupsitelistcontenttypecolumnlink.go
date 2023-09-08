package groupsitelistcontenttypecolumnlink

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListContentTypeColumnLinkId{}

// GroupSiteListContentTypeColumnLinkId is a struct representing the Resource ID for a Group Site List Content Type Column Link
type GroupSiteListContentTypeColumnLinkId struct {
	GroupId       string
	SiteId        string
	ListId        string
	ContentTypeId string
	ColumnLinkId  string
}

// NewGroupSiteListContentTypeColumnLinkID returns a new GroupSiteListContentTypeColumnLinkId struct
func NewGroupSiteListContentTypeColumnLinkID(groupId string, siteId string, listId string, contentTypeId string, columnLinkId string) GroupSiteListContentTypeColumnLinkId {
	return GroupSiteListContentTypeColumnLinkId{
		GroupId:       groupId,
		SiteId:        siteId,
		ListId:        listId,
		ContentTypeId: contentTypeId,
		ColumnLinkId:  columnLinkId,
	}
}

// ParseGroupSiteListContentTypeColumnLinkID parses 'input' into a GroupSiteListContentTypeColumnLinkId
func ParseGroupSiteListContentTypeColumnLinkID(input string) (*GroupSiteListContentTypeColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListContentTypeColumnLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListContentTypeColumnLinkId{}

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

	if id.ColumnLinkId, ok = parsed.Parsed["columnLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnLinkId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListContentTypeColumnLinkIDInsensitively parses 'input' case-insensitively into a GroupSiteListContentTypeColumnLinkId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListContentTypeColumnLinkIDInsensitively(input string) (*GroupSiteListContentTypeColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListContentTypeColumnLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListContentTypeColumnLinkId{}

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

	if id.ColumnLinkId, ok = parsed.Parsed["columnLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnLinkId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListContentTypeColumnLinkID checks that 'input' can be parsed as a Group Site List Content Type Column Link ID
func ValidateGroupSiteListContentTypeColumnLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListContentTypeColumnLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Content Type Column Link ID
func (id GroupSiteListContentTypeColumnLinkId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/contentTypes/%s/columnLinks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ContentTypeId, id.ColumnLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Content Type Column Link ID
func (id GroupSiteListContentTypeColumnLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
		resourceids.StaticSegment("staticColumnLinks", "columnLinks", "columnLinks"),
		resourceids.UserSpecifiedSegment("columnLinkId", "columnLinkIdValue"),
	}
}

// String returns a human-readable description of this Group Site List Content Type Column Link ID
func (id GroupSiteListContentTypeColumnLinkId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Link: %q", id.ColumnLinkId),
	}
	return fmt.Sprintf("Group Site List Content Type Column Link (%s)", strings.Join(components, "\n"))
}
