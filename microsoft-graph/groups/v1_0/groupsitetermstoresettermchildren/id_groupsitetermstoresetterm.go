package groupsitetermstoresettermchildren

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetTermId{}

// GroupSiteTermStoreSetTermId is a struct representing the Resource ID for a Group Site Term Store Set Term
type GroupSiteTermStoreSetTermId struct {
	GroupId string
	SiteId  string
	SetId   string
	TermId  string
}

// NewGroupSiteTermStoreSetTermID returns a new GroupSiteTermStoreSetTermId struct
func NewGroupSiteTermStoreSetTermID(groupId string, siteId string, setId string, termId string) GroupSiteTermStoreSetTermId {
	return GroupSiteTermStoreSetTermId{
		GroupId: groupId,
		SiteId:  siteId,
		SetId:   setId,
		TermId:  termId,
	}
}

// ParseGroupSiteTermStoreSetTermID parses 'input' into a GroupSiteTermStoreSetTermId
func ParseGroupSiteTermStoreSetTermID(input string) (*GroupSiteTermStoreSetTermId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetTermId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetTermId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.TermId, ok = parsed.Parsed["termId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreSetTermIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreSetTermId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreSetTermIDInsensitively(input string) (*GroupSiteTermStoreSetTermId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetTermId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetTermId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.TermId, ok = parsed.Parsed["termId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreSetTermID checks that 'input' can be parsed as a Group Site Term Store Set Term ID
func ValidateGroupSiteTermStoreSetTermID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreSetTermID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Set Term ID
func (id GroupSiteTermStoreSetTermId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/terms/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.TermId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Set Term ID
func (id GroupSiteTermStoreSetTermId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStore", "termStore", "termStore"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("staticTerms", "terms", "terms"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Set Term ID
func (id GroupSiteTermStoreSetTermId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
	}
	return fmt.Sprintf("Group Site Term Store Set Term (%s)", strings.Join(components, "\n"))
}
