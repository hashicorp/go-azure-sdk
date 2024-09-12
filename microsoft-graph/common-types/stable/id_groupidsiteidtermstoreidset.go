package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreIdSetId{}

// GroupIdSiteIdTermStoreIdSetId is a struct representing the Resource ID for a Group Id Site Id Term Store Id Set
type GroupIdSiteIdTermStoreIdSetId struct {
	GroupId string
	SiteId  string
	StoreId string
	SetId   string
}

// NewGroupIdSiteIdTermStoreIdSetID returns a new GroupIdSiteIdTermStoreIdSetId struct
func NewGroupIdSiteIdTermStoreIdSetID(groupId string, siteId string, storeId string, setId string) GroupIdSiteIdTermStoreIdSetId {
	return GroupIdSiteIdTermStoreIdSetId{
		GroupId: groupId,
		SiteId:  siteId,
		StoreId: storeId,
		SetId:   setId,
	}
}

// ParseGroupIdSiteIdTermStoreIdSetID parses 'input' into a GroupIdSiteIdTermStoreIdSetId
func ParseGroupIdSiteIdTermStoreIdSetID(input string) (*GroupIdSiteIdTermStoreIdSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreIdSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreIdSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreIdSetIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreIdSetId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreIdSetIDInsensitively(input string) (*GroupIdSiteIdTermStoreIdSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreIdSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreIdSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreIdSetId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.SetId, ok = input.Parsed["setId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "setId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdTermStoreIdSetID checks that 'input' can be parsed as a Group Id Site Id Term Store Id Set ID
func ValidateGroupIdSiteIdTermStoreIdSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreIdSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Id Set ID
func (id GroupIdSiteIdTermStoreIdSetId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStores/%s/sets/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.StoreId, id.SetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Id Set ID
func (id GroupIdSiteIdTermStoreIdSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("termStores", "termStores", "termStores"),
		resourceids.UserSpecifiedSegment("storeId", "storeIdValue"),
		resourceids.StaticSegment("sets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Id Set ID
func (id GroupIdSiteIdTermStoreIdSetId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Store: %q", id.StoreId),
		fmt.Sprintf("Set: %q", id.SetId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Id Set (%s)", strings.Join(components, "\n"))
}
