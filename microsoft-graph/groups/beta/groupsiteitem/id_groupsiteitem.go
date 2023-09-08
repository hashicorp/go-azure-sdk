package groupsiteitem

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteItemId{}

// GroupSiteItemId is a struct representing the Resource ID for a Group Site Item
type GroupSiteItemId struct {
	GroupId    string
	SiteId     string
	BaseItemId string
}

// NewGroupSiteItemID returns a new GroupSiteItemId struct
func NewGroupSiteItemID(groupId string, siteId string, baseItemId string) GroupSiteItemId {
	return GroupSiteItemId{
		GroupId:    groupId,
		SiteId:     siteId,
		BaseItemId: baseItemId,
	}
}

// ParseGroupSiteItemID parses 'input' into a GroupSiteItemId
func ParseGroupSiteItemID(input string) (*GroupSiteItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteItemId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.BaseItemId, ok = parsed.Parsed["baseItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "baseItemId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteItemIDInsensitively parses 'input' case-insensitively into a GroupSiteItemId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteItemIDInsensitively(input string) (*GroupSiteItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteItemId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.BaseItemId, ok = parsed.Parsed["baseItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "baseItemId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteItemID checks that 'input' can be parsed as a Group Site Item ID
func ValidateGroupSiteItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Item ID
func (id GroupSiteItemId) ID() string {
	fmtString := "/groups/%s/sites/%s/items/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BaseItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Item ID
func (id GroupSiteItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticItems", "items", "items"),
		resourceids.UserSpecifiedSegment("baseItemId", "baseItemIdValue"),
	}
}

// String returns a human-readable description of this Group Site Item ID
func (id GroupSiteItemId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Base Item: %q", id.BaseItemId),
	}
	return fmt.Sprintf("Group Site Item (%s)", strings.Join(components, "\n"))
}
