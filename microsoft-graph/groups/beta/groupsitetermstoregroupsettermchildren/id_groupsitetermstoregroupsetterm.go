package groupsitetermstoregroupsettermchildren

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreGroupSetTermId{}

// GroupSiteTermStoreGroupSetTermId is a struct representing the Resource ID for a Group Site Term Store Group Set Term
type GroupSiteTermStoreGroupSetTermId struct {
	GroupId  string
	SiteId   string
	GroupId1 string
	SetId    string
	TermId   string
}

// NewGroupSiteTermStoreGroupSetTermID returns a new GroupSiteTermStoreGroupSetTermId struct
func NewGroupSiteTermStoreGroupSetTermID(groupId string, siteId string, groupId1 string, setId string, termId string) GroupSiteTermStoreGroupSetTermId {
	return GroupSiteTermStoreGroupSetTermId{
		GroupId:  groupId,
		SiteId:   siteId,
		GroupId1: groupId1,
		SetId:    setId,
		TermId:   termId,
	}
}

// ParseGroupSiteTermStoreGroupSetTermID parses 'input' into a GroupSiteTermStoreGroupSetTermId
func ParseGroupSiteTermStoreGroupSetTermID(input string) (*GroupSiteTermStoreGroupSetTermId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreGroupSetTermId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreGroupSetTermId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.GroupId1, ok = parsed.Parsed["groupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId1", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.TermId, ok = parsed.Parsed["termId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreGroupSetTermIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreGroupSetTermId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreGroupSetTermIDInsensitively(input string) (*GroupSiteTermStoreGroupSetTermId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreGroupSetTermId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreGroupSetTermId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.GroupId1, ok = parsed.Parsed["groupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId1", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.TermId, ok = parsed.Parsed["termId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreGroupSetTermID checks that 'input' can be parsed as a Group Site Term Store Group Set Term ID
func ValidateGroupSiteTermStoreGroupSetTermID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreGroupSetTermID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Group Set Term ID
func (id GroupSiteTermStoreGroupSetTermId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/groups/%s/sets/%s/terms/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.GroupId1, id.SetId, id.TermId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Group Set Term ID
func (id GroupSiteTermStoreGroupSetTermId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStore", "termStore", "termStore"),
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId1", "groupId1Value"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("staticTerms", "terms", "terms"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Group Set Term ID
func (id GroupSiteTermStoreGroupSetTermId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Group Id 1: %q", id.GroupId1),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
	}
	return fmt.Sprintf("Group Site Term Store Group Set Term (%s)", strings.Join(components, "\n"))
}
