package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreId{}

// GroupIdSiteIdTermStoreId is a struct representing the Resource ID for a Group Id Site Id Term Store
type GroupIdSiteIdTermStoreId struct {
	GroupId string
	SiteId  string
	StoreId string
}

// NewGroupIdSiteIdTermStoreID returns a new GroupIdSiteIdTermStoreId struct
func NewGroupIdSiteIdTermStoreID(groupId string, siteId string, storeId string) GroupIdSiteIdTermStoreId {
	return GroupIdSiteIdTermStoreId{
		GroupId: groupId,
		SiteId:  siteId,
		StoreId: storeId,
	}
}

// ParseGroupIdSiteIdTermStoreID parses 'input' into a GroupIdSiteIdTermStoreId
func ParseGroupIdSiteIdTermStoreID(input string) (*GroupIdSiteIdTermStoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreIDInsensitively(input string) (*GroupIdSiteIdTermStoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.StoreId, ok = input.Parsed["storeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "storeId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdTermStoreID checks that 'input' can be parsed as a Group Id Site Id Term Store ID
func ValidateGroupIdSiteIdTermStoreID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store ID
func (id GroupIdSiteIdTermStoreId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStores/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.StoreId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store ID
func (id GroupIdSiteIdTermStoreId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("termStores", "termStores", "termStores"),
		resourceids.UserSpecifiedSegment("storeId", "storeIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store ID
func (id GroupIdSiteIdTermStoreId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Store: %q", id.StoreId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store (%s)", strings.Join(components, "\n"))
}
