package groupsitelistcontenttype

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListContentTypeId{}

// GroupSiteListContentTypeId is a struct representing the Resource ID for a Group Site List Content Type
type GroupSiteListContentTypeId struct {
	GroupId       string
	SiteId        string
	ListId        string
	ContentTypeId string
}

// NewGroupSiteListContentTypeID returns a new GroupSiteListContentTypeId struct
func NewGroupSiteListContentTypeID(groupId string, siteId string, listId string, contentTypeId string) GroupSiteListContentTypeId {
	return GroupSiteListContentTypeId{
		GroupId:       groupId,
		SiteId:        siteId,
		ListId:        listId,
		ContentTypeId: contentTypeId,
	}
}

// ParseGroupSiteListContentTypeID parses 'input' into a GroupSiteListContentTypeId
func ParseGroupSiteListContentTypeID(input string) (*GroupSiteListContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListContentTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListContentTypeId{}

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

	return &id, nil
}

// ParseGroupSiteListContentTypeIDInsensitively parses 'input' case-insensitively into a GroupSiteListContentTypeId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListContentTypeIDInsensitively(input string) (*GroupSiteListContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListContentTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListContentTypeId{}

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

	return &id, nil
}

// ValidateGroupSiteListContentTypeID checks that 'input' can be parsed as a Group Site List Content Type ID
func ValidateGroupSiteListContentTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListContentTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Content Type ID
func (id GroupSiteListContentTypeId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/contentTypes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ContentTypeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Content Type ID
func (id GroupSiteListContentTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
	}
}

// String returns a human-readable description of this Group Site List Content Type ID
func (id GroupSiteListContentTypeId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
	}
	return fmt.Sprintf("Group Site List Content Type (%s)", strings.Join(components, "\n"))
}
