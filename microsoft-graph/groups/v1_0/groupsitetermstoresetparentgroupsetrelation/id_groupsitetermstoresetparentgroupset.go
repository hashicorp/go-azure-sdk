package groupsitetermstoresetparentgroupsetrelation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetParentGroupSetId{}

// GroupSiteTermStoreSetParentGroupSetId is a struct representing the Resource ID for a Group Site Term Store Set Parent Group Set
type GroupSiteTermStoreSetParentGroupSetId struct {
	GroupId string
	SiteId  string
	StoreId string
	SetId   string
	SetId1  string
}

// NewGroupSiteTermStoreSetParentGroupSetID returns a new GroupSiteTermStoreSetParentGroupSetId struct
func NewGroupSiteTermStoreSetParentGroupSetID(groupId string, siteId string, storeId string, setId string, setId1 string) GroupSiteTermStoreSetParentGroupSetId {
	return GroupSiteTermStoreSetParentGroupSetId{
		GroupId: groupId,
		SiteId:  siteId,
		StoreId: storeId,
		SetId:   setId,
		SetId1:  setId1,
	}
}

// ParseGroupSiteTermStoreSetParentGroupSetID parses 'input' into a GroupSiteTermStoreSetParentGroupSetId
func ParseGroupSiteTermStoreSetParentGroupSetID(input string) (*GroupSiteTermStoreSetParentGroupSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.StoreId, ok = parsed.Parsed["storeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storeId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.SetId1, ok = parsed.Parsed["setId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId1", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreSetParentGroupSetIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreSetParentGroupSetId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreSetParentGroupSetIDInsensitively(input string) (*GroupSiteTermStoreSetParentGroupSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.StoreId, ok = parsed.Parsed["storeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storeId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.SetId1, ok = parsed.Parsed["setId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreSetParentGroupSetID checks that 'input' can be parsed as a Group Site Term Store Set Parent Group Set ID
func ValidateGroupSiteTermStoreSetParentGroupSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreSetParentGroupSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Set Parent Group Set ID
func (id GroupSiteTermStoreSetParentGroupSetId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStores/%s/sets/%s/parentGroup/sets/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.StoreId, id.SetId, id.SetId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Set Parent Group Set ID
func (id GroupSiteTermStoreSetParentGroupSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStores", "termStores", "termStores"),
		resourceids.UserSpecifiedSegment("storeId", "storeIdValue"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("staticParentGroup", "parentGroup", "parentGroup"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId1", "setId1Value"),
	}
}

// String returns a human-readable description of this Group Site Term Store Set Parent Group Set ID
func (id GroupSiteTermStoreSetParentGroupSetId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Store: %q", id.StoreId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Set Id 1: %q", id.SetId1),
	}
	return fmt.Sprintf("Group Site Term Store Set Parent Group Set (%s)", strings.Join(components, "\n"))
}
