package groupsitelistitemactivitydriveitem

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListItemActivityId{}

// GroupSiteListItemActivityId is a struct representing the Resource ID for a Group Site List Item Activity
type GroupSiteListItemActivityId struct {
	GroupId           string
	SiteId            string
	ListId            string
	ListItemId        string
	ItemActivityOLDId string
}

// NewGroupSiteListItemActivityID returns a new GroupSiteListItemActivityId struct
func NewGroupSiteListItemActivityID(groupId string, siteId string, listId string, listItemId string, itemActivityOLDId string) GroupSiteListItemActivityId {
	return GroupSiteListItemActivityId{
		GroupId:           groupId,
		SiteId:            siteId,
		ListId:            listId,
		ListItemId:        listItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseGroupSiteListItemActivityID parses 'input' into a GroupSiteListItemActivityId
func ParseGroupSiteListItemActivityID(input string) (*GroupSiteListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListItemActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListItemActivityId{}

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

	if id.ItemActivityOLDId, ok = parsed.Parsed["itemActivityOLDId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListItemActivityIDInsensitively parses 'input' case-insensitively into a GroupSiteListItemActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListItemActivityIDInsensitively(input string) (*GroupSiteListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListItemActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListItemActivityId{}

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

	if id.ItemActivityOLDId, ok = parsed.Parsed["itemActivityOLDId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListItemActivityID checks that 'input' can be parsed as a Group Site List Item Activity ID
func ValidateGroupSiteListItemActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListItemActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Item Activity ID
func (id GroupSiteListItemActivityId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/items/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ListItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Item Activity ID
func (id GroupSiteListItemActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticItems", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemIdValue"),
		resourceids.StaticSegment("staticActivities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDIdValue"),
	}
}

// String returns a human-readable description of this Group Site List Item Activity ID
func (id GroupSiteListItemActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Item Activity O L D: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Group Site List Item Activity (%s)", strings.Join(components, "\n"))
}
