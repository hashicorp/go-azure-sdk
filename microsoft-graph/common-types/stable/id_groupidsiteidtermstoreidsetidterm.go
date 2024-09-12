package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreIdSetIdTermId{}

// GroupIdSiteIdTermStoreIdSetIdTermId is a struct representing the Resource ID for a Group Id Site Id Term Store Id Set Id Term
type GroupIdSiteIdTermStoreIdSetIdTermId struct {
	GroupId string
	SiteId  string
	StoreId string
	SetId   string
	TermId  string
}

// NewGroupIdSiteIdTermStoreIdSetIdTermID returns a new GroupIdSiteIdTermStoreIdSetIdTermId struct
func NewGroupIdSiteIdTermStoreIdSetIdTermID(groupId string, siteId string, storeId string, setId string, termId string) GroupIdSiteIdTermStoreIdSetIdTermId {
	return GroupIdSiteIdTermStoreIdSetIdTermId{
		GroupId: groupId,
		SiteId:  siteId,
		StoreId: storeId,
		SetId:   setId,
		TermId:  termId,
	}
}

// ParseGroupIdSiteIdTermStoreIdSetIdTermID parses 'input' into a GroupIdSiteIdTermStoreIdSetIdTermId
func ParseGroupIdSiteIdTermStoreIdSetIdTermID(input string) (*GroupIdSiteIdTermStoreIdSetIdTermId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreIdSetIdTermId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreIdSetIdTermId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreIdSetIdTermIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreIdSetIdTermId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreIdSetIdTermIDInsensitively(input string) (*GroupIdSiteIdTermStoreIdSetIdTermId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreIdSetIdTermId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreIdSetIdTermId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreIdSetIdTermId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.TermId, ok = input.Parsed["termId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdTermStoreIdSetIdTermID checks that 'input' can be parsed as a Group Id Site Id Term Store Id Set Id Term ID
func ValidateGroupIdSiteIdTermStoreIdSetIdTermID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreIdSetIdTermID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Id Set Id Term ID
func (id GroupIdSiteIdTermStoreIdSetIdTermId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStores/%s/sets/%s/terms/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.StoreId, id.SetId, id.TermId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Id Set Id Term ID
func (id GroupIdSiteIdTermStoreIdSetIdTermId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("termStores", "termStores", "termStores"),
		resourceids.UserSpecifiedSegment("storeId", "storeIdValue"),
		resourceids.StaticSegment("sets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("terms", "terms", "terms"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Id Set Id Term ID
func (id GroupIdSiteIdTermStoreIdSetIdTermId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Store: %q", id.StoreId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Id Set Id Term (%s)", strings.Join(components, "\n"))
}
