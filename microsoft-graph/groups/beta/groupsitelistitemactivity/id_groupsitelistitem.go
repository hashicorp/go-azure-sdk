package groupsitelistitemactivity

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListItemId{}

// GroupSiteListItemId is a struct representing the Resource ID for a Group Site List Item
type GroupSiteListItemId struct {
	GroupId    string
	SiteId     string
	ListId     string
	ListItemId string
}

// NewGroupSiteListItemID returns a new GroupSiteListItemId struct
func NewGroupSiteListItemID(groupId string, siteId string, listId string, listItemId string) GroupSiteListItemId {
	return GroupSiteListItemId{
		GroupId:    groupId,
		SiteId:     siteId,
		ListId:     listId,
		ListItemId: listItemId,
	}
}

// ParseGroupSiteListItemID parses 'input' into a GroupSiteListItemId
func ParseGroupSiteListItemID(input string) (*GroupSiteListItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListItemId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.ListItemId, ok = parsed.Parsed["listItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listItemId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListItemIDInsensitively parses 'input' case-insensitively into a GroupSiteListItemId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListItemIDInsensitively(input string) (*GroupSiteListItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListItemId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.ListItemId, ok = parsed.Parsed["listItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listItemId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListItemID checks that 'input' can be parsed as a Group Site List Item ID
func ValidateGroupSiteListItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Item ID
func (id GroupSiteListItemId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/items/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ListItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Item ID
func (id GroupSiteListItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticItems", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemIdValue"),
	}
}

// String returns a human-readable description of this Group Site List Item ID
func (id GroupSiteListItemId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
	}
	return fmt.Sprintf("Group Site List Item (%s)", strings.Join(components, "\n"))
}
