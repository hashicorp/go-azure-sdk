package groupsitetermstoresetchildrenchildren

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetChildrenId{}

// GroupSiteTermStoreSetChildrenId is a struct representing the Resource ID for a Group Site Term Store Set Children
type GroupSiteTermStoreSetChildrenId struct {
	GroupId string
	SiteId  string
	SetId   string
	TermId  string
}

// NewGroupSiteTermStoreSetChildrenID returns a new GroupSiteTermStoreSetChildrenId struct
func NewGroupSiteTermStoreSetChildrenID(groupId string, siteId string, setId string, termId string) GroupSiteTermStoreSetChildrenId {
	return GroupSiteTermStoreSetChildrenId{
		GroupId: groupId,
		SiteId:  siteId,
		SetId:   setId,
		TermId:  termId,
	}
}

// ParseGroupSiteTermStoreSetChildrenID parses 'input' into a GroupSiteTermStoreSetChildrenId
func ParseGroupSiteTermStoreSetChildrenID(input string) (*GroupSiteTermStoreSetChildrenId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetChildrenId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetChildrenId{}

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

// ParseGroupSiteTermStoreSetChildrenIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreSetChildrenId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreSetChildrenIDInsensitively(input string) (*GroupSiteTermStoreSetChildrenId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetChildrenId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetChildrenId{}

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

// ValidateGroupSiteTermStoreSetChildrenID checks that 'input' can be parsed as a Group Site Term Store Set Children ID
func ValidateGroupSiteTermStoreSetChildrenID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreSetChildrenID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Set Children ID
func (id GroupSiteTermStoreSetChildrenId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/children/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.TermId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Set Children ID
func (id GroupSiteTermStoreSetChildrenId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStore", "termStore", "termStore"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("staticChildren", "children", "children"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Set Children ID
func (id GroupSiteTermStoreSetChildrenId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
	}
	return fmt.Sprintf("Group Site Term Store Set Children (%s)", strings.Join(components, "\n"))
}
