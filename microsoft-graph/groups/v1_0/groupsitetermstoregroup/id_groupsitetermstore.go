package groupsitetermstoregroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreId{}

// GroupSiteTermStoreId is a struct representing the Resource ID for a Group Site Term Store
type GroupSiteTermStoreId struct {
	GroupId string
	SiteId  string
	StoreId string
}

// NewGroupSiteTermStoreID returns a new GroupSiteTermStoreId struct
func NewGroupSiteTermStoreID(groupId string, siteId string, storeId string) GroupSiteTermStoreId {
	return GroupSiteTermStoreId{
		GroupId: groupId,
		SiteId:  siteId,
		StoreId: storeId,
	}
}

// ParseGroupSiteTermStoreID parses 'input' into a GroupSiteTermStoreId
func ParseGroupSiteTermStoreID(input string) (*GroupSiteTermStoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.StoreId, ok = parsed.Parsed["storeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storeId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreIDInsensitively(input string) (*GroupSiteTermStoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.StoreId, ok = parsed.Parsed["storeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storeId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreID checks that 'input' can be parsed as a Group Site Term Store ID
func ValidateGroupSiteTermStoreID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store ID
func (id GroupSiteTermStoreId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStores/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.StoreId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store ID
func (id GroupSiteTermStoreId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStores", "termStores", "termStores"),
		resourceids.UserSpecifiedSegment("storeId", "storeIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store ID
func (id GroupSiteTermStoreId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Store: %q", id.StoreId),
	}
	return fmt.Sprintf("Group Site Term Store (%s)", strings.Join(components, "\n"))
}
