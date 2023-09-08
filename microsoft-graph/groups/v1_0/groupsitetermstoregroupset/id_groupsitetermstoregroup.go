package groupsitetermstoregroupset

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreGroupId{}

// GroupSiteTermStoreGroupId is a struct representing the Resource ID for a Group Site Term Store Group
type GroupSiteTermStoreGroupId struct {
	GroupId  string
	SiteId   string
	GroupId1 string
}

// NewGroupSiteTermStoreGroupID returns a new GroupSiteTermStoreGroupId struct
func NewGroupSiteTermStoreGroupID(groupId string, siteId string, groupId1 string) GroupSiteTermStoreGroupId {
	return GroupSiteTermStoreGroupId{
		GroupId:  groupId,
		SiteId:   siteId,
		GroupId1: groupId1,
	}
}

// ParseGroupSiteTermStoreGroupID parses 'input' into a GroupSiteTermStoreGroupId
func ParseGroupSiteTermStoreGroupID(input string) (*GroupSiteTermStoreGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.GroupId1, ok = parsed.Parsed["groupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId1", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreGroupIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreGroupIDInsensitively(input string) (*GroupSiteTermStoreGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.GroupId1, ok = parsed.Parsed["groupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreGroupID checks that 'input' can be parsed as a Group Site Term Store Group ID
func ValidateGroupSiteTermStoreGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Group ID
func (id GroupSiteTermStoreGroupId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/groups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.GroupId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Group ID
func (id GroupSiteTermStoreGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStore", "termStore", "termStore"),
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId1", "groupId1Value"),
	}
}

// String returns a human-readable description of this Group Site Term Store Group ID
func (id GroupSiteTermStoreGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Group Id 1: %q", id.GroupId1),
	}
	return fmt.Sprintf("Group Site Term Store Group (%s)", strings.Join(components, "\n"))
}
