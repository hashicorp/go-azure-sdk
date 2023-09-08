package groupsitelistcontenttypebasetype

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListContentTypeBaseTypeId{}

// GroupSiteListContentTypeBaseTypeId is a struct representing the Resource ID for a Group Site List Content Type Base Type
type GroupSiteListContentTypeBaseTypeId struct {
	GroupId        string
	SiteId         string
	ListId         string
	ContentTypeId  string
	ContentTypeId1 string
}

// NewGroupSiteListContentTypeBaseTypeID returns a new GroupSiteListContentTypeBaseTypeId struct
func NewGroupSiteListContentTypeBaseTypeID(groupId string, siteId string, listId string, contentTypeId string, contentTypeId1 string) GroupSiteListContentTypeBaseTypeId {
	return GroupSiteListContentTypeBaseTypeId{
		GroupId:        groupId,
		SiteId:         siteId,
		ListId:         listId,
		ContentTypeId:  contentTypeId,
		ContentTypeId1: contentTypeId1,
	}
}

// ParseGroupSiteListContentTypeBaseTypeID parses 'input' into a GroupSiteListContentTypeBaseTypeId
func ParseGroupSiteListContentTypeBaseTypeID(input string) (*GroupSiteListContentTypeBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListContentTypeBaseTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListContentTypeBaseTypeId{}

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

	if id.ContentTypeId1, ok = parsed.Parsed["contentTypeId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId1", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListContentTypeBaseTypeIDInsensitively parses 'input' case-insensitively into a GroupSiteListContentTypeBaseTypeId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListContentTypeBaseTypeIDInsensitively(input string) (*GroupSiteListContentTypeBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListContentTypeBaseTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListContentTypeBaseTypeId{}

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

	if id.ContentTypeId1, ok = parsed.Parsed["contentTypeId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListContentTypeBaseTypeID checks that 'input' can be parsed as a Group Site List Content Type Base Type ID
func ValidateGroupSiteListContentTypeBaseTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListContentTypeBaseTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Content Type Base Type ID
func (id GroupSiteListContentTypeBaseTypeId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/contentTypes/%s/baseTypes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ContentTypeId, id.ContentTypeId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Content Type Base Type ID
func (id GroupSiteListContentTypeBaseTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
		resourceids.StaticSegment("staticBaseTypes", "baseTypes", "baseTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId1", "contentTypeId1Value"),
	}
}

// String returns a human-readable description of this Group Site List Content Type Base Type ID
func (id GroupSiteListContentTypeBaseTypeId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Content Type Id 1: %q", id.ContentTypeId1),
	}
	return fmt.Sprintf("Group Site List Content Type Base Type (%s)", strings.Join(components, "\n"))
}
