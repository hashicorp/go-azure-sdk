package groupsitetermstoresetrelation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetId{}

// GroupSiteTermStoreSetId is a struct representing the Resource ID for a Group Site Term Store Set
type GroupSiteTermStoreSetId struct {
	GroupId string
	SiteId  string
	SetId   string
}

// NewGroupSiteTermStoreSetID returns a new GroupSiteTermStoreSetId struct
func NewGroupSiteTermStoreSetID(groupId string, siteId string, setId string) GroupSiteTermStoreSetId {
	return GroupSiteTermStoreSetId{
		GroupId: groupId,
		SiteId:  siteId,
		SetId:   setId,
	}
}

// ParseGroupSiteTermStoreSetID parses 'input' into a GroupSiteTermStoreSetId
func ParseGroupSiteTermStoreSetID(input string) (*GroupSiteTermStoreSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreSetIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreSetId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreSetIDInsensitively(input string) (*GroupSiteTermStoreSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreSetID checks that 'input' can be parsed as a Group Site Term Store Set ID
func ValidateGroupSiteTermStoreSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Set ID
func (id GroupSiteTermStoreSetId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Set ID
func (id GroupSiteTermStoreSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStore", "termStore", "termStore"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Set ID
func (id GroupSiteTermStoreSetId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
	}
	return fmt.Sprintf("Group Site Term Store Set (%s)", strings.Join(components, "\n"))
}
