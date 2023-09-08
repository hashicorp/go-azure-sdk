package groupsitelistcontenttype

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListId{}

// GroupSiteListId is a struct representing the Resource ID for a Group Site List
type GroupSiteListId struct {
	GroupId string
	SiteId  string
	ListId  string
}

// NewGroupSiteListID returns a new GroupSiteListId struct
func NewGroupSiteListID(groupId string, siteId string, listId string) GroupSiteListId {
	return GroupSiteListId{
		GroupId: groupId,
		SiteId:  siteId,
		ListId:  listId,
	}
}

// ParseGroupSiteListID parses 'input' into a GroupSiteListId
func ParseGroupSiteListID(input string) (*GroupSiteListId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListIDInsensitively parses 'input' case-insensitively into a GroupSiteListId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListIDInsensitively(input string) (*GroupSiteListId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListID checks that 'input' can be parsed as a Group Site List ID
func ValidateGroupSiteListID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List ID
func (id GroupSiteListId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List ID
func (id GroupSiteListId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
	}
}

// String returns a human-readable description of this Group Site List ID
func (id GroupSiteListId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
	}
	return fmt.Sprintf("Group Site List (%s)", strings.Join(components, "\n"))
}
