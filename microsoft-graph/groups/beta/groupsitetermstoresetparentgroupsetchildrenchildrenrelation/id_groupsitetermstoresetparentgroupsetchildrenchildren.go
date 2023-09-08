package groupsitetermstoresetparentgroupsetchildrenchildrenrelation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetParentGroupSetChildrenChildrenId{}

// GroupSiteTermStoreSetParentGroupSetChildrenChildrenId is a struct representing the Resource ID for a Group Site Term Store Set Parent Group Set Children Children
type GroupSiteTermStoreSetParentGroupSetChildrenChildrenId struct {
	GroupId string
	SiteId  string
	SetId   string
	SetId1  string
	TermId  string
	TermId1 string
}

// NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenID returns a new GroupSiteTermStoreSetParentGroupSetChildrenChildrenId struct
func NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenID(groupId string, siteId string, setId string, setId1 string, termId string, termId1 string) GroupSiteTermStoreSetParentGroupSetChildrenChildrenId {
	return GroupSiteTermStoreSetParentGroupSetChildrenChildrenId{
		GroupId: groupId,
		SiteId:  siteId,
		SetId:   setId,
		SetId1:  setId1,
		TermId:  termId,
		TermId1: termId1,
	}
}

// ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenID parses 'input' into a GroupSiteTermStoreSetParentGroupSetChildrenChildrenId
func ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenID(input string) (*GroupSiteTermStoreSetParentGroupSetChildrenChildrenId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetChildrenChildrenId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetChildrenChildrenId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.SetId1, ok = parsed.Parsed["setId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId1", *parsed)
	}

	if id.TermId, ok = parsed.Parsed["termId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId", *parsed)
	}

	if id.TermId1, ok = parsed.Parsed["termId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId1", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreSetParentGroupSetChildrenChildrenId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenIDInsensitively(input string) (*GroupSiteTermStoreSetParentGroupSetChildrenChildrenId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetChildrenChildrenId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetChildrenChildrenId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.SetId1, ok = parsed.Parsed["setId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId1", *parsed)
	}

	if id.TermId, ok = parsed.Parsed["termId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId", *parsed)
	}

	if id.TermId1, ok = parsed.Parsed["termId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreSetParentGroupSetChildrenChildrenID checks that 'input' can be parsed as a Group Site Term Store Set Parent Group Set Children Children ID
func ValidateGroupSiteTermStoreSetParentGroupSetChildrenChildrenID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Set Parent Group Set Children Children ID
func (id GroupSiteTermStoreSetParentGroupSetChildrenChildrenId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/parentGroup/sets/%s/children/%s/children/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.SetId1, id.TermId, id.TermId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Set Parent Group Set Children Children ID
func (id GroupSiteTermStoreSetParentGroupSetChildrenChildrenId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStore", "termStore", "termStore"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("staticParentGroup", "parentGroup", "parentGroup"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId1", "setId1Value"),
		resourceids.StaticSegment("staticChildren", "children", "children"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
		resourceids.StaticSegment("staticChildren", "children", "children"),
		resourceids.UserSpecifiedSegment("termId1", "termId1Value"),
	}
}

// String returns a human-readable description of this Group Site Term Store Set Parent Group Set Children Children ID
func (id GroupSiteTermStoreSetParentGroupSetChildrenChildrenId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Set Id 1: %q", id.SetId1),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Term Id 1: %q", id.TermId1),
	}
	return fmt.Sprintf("Group Site Term Store Set Parent Group Set Children Children (%s)", strings.Join(components, "\n"))
}
