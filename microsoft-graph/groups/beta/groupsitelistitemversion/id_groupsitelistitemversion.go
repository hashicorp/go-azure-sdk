package groupsitelistitemversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListItemVersionId{}

// GroupSiteListItemVersionId is a struct representing the Resource ID for a Group Site List Item Version
type GroupSiteListItemVersionId struct {
	GroupId           string
	SiteId            string
	ListId            string
	ListItemId        string
	ListItemVersionId string
}

// NewGroupSiteListItemVersionID returns a new GroupSiteListItemVersionId struct
func NewGroupSiteListItemVersionID(groupId string, siteId string, listId string, listItemId string, listItemVersionId string) GroupSiteListItemVersionId {
	return GroupSiteListItemVersionId{
		GroupId:           groupId,
		SiteId:            siteId,
		ListId:            listId,
		ListItemId:        listItemId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseGroupSiteListItemVersionID parses 'input' into a GroupSiteListItemVersionId
func ParseGroupSiteListItemVersionID(input string) (*GroupSiteListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListItemVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListItemVersionId{}

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

	if id.ListItemVersionId, ok = parsed.Parsed["listItemVersionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListItemVersionIDInsensitively parses 'input' case-insensitively into a GroupSiteListItemVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListItemVersionIDInsensitively(input string) (*GroupSiteListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListItemVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListItemVersionId{}

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

	if id.ListItemVersionId, ok = parsed.Parsed["listItemVersionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListItemVersionID checks that 'input' can be parsed as a Group Site List Item Version ID
func ValidateGroupSiteListItemVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListItemVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Item Version ID
func (id GroupSiteListItemVersionId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/items/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ListItemId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Item Version ID
func (id GroupSiteListItemVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticItems", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemIdValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionIdValue"),
	}
}

// String returns a human-readable description of this Group Site List Item Version ID
func (id GroupSiteListItemVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("Group Site List Item Version (%s)", strings.Join(components, "\n"))
}
