package groupsitetermstoregroupsetchildrenchildrenrelation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreGroupSetChildrenChildrenRelationId{}

// GroupSiteTermStoreGroupSetChildrenChildrenRelationId is a struct representing the Resource ID for a Group Site Term Store Group Set Children Children Relation
type GroupSiteTermStoreGroupSetChildrenChildrenRelationId struct {
	GroupId    string
	SiteId     string
	GroupId1   string
	SetId      string
	TermId     string
	TermId1    string
	RelationId string
}

// NewGroupSiteTermStoreGroupSetChildrenChildrenRelationID returns a new GroupSiteTermStoreGroupSetChildrenChildrenRelationId struct
func NewGroupSiteTermStoreGroupSetChildrenChildrenRelationID(groupId string, siteId string, groupId1 string, setId string, termId string, termId1 string, relationId string) GroupSiteTermStoreGroupSetChildrenChildrenRelationId {
	return GroupSiteTermStoreGroupSetChildrenChildrenRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		GroupId1:   groupId1,
		SetId:      setId,
		TermId:     termId,
		TermId1:    termId1,
		RelationId: relationId,
	}
}

// ParseGroupSiteTermStoreGroupSetChildrenChildrenRelationID parses 'input' into a GroupSiteTermStoreGroupSetChildrenChildrenRelationId
func ParseGroupSiteTermStoreGroupSetChildrenChildrenRelationID(input string) (*GroupSiteTermStoreGroupSetChildrenChildrenRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreGroupSetChildrenChildrenRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreGroupSetChildrenChildrenRelationId{}

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

	if id.TermId1, ok = parsed.Parsed["termId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId1", *parsed)
	}

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreGroupSetChildrenChildrenRelationIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreGroupSetChildrenChildrenRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreGroupSetChildrenChildrenRelationIDInsensitively(input string) (*GroupSiteTermStoreGroupSetChildrenChildrenRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreGroupSetChildrenChildrenRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreGroupSetChildrenChildrenRelationId{}

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

	if id.TermId1, ok = parsed.Parsed["termId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "termId1", *parsed)
	}

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreGroupSetChildrenChildrenRelationID checks that 'input' can be parsed as a Group Site Term Store Group Set Children Children Relation ID
func ValidateGroupSiteTermStoreGroupSetChildrenChildrenRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreGroupSetChildrenChildrenRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Group Set Children Children Relation ID
func (id GroupSiteTermStoreGroupSetChildrenChildrenRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/groups/%s/sets/%s/children/%s/children/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.GroupId1, id.SetId, id.TermId, id.TermId1, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Group Set Children Children Relation ID
func (id GroupSiteTermStoreGroupSetChildrenChildrenRelationId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticChildren", "children", "children"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
		resourceids.StaticSegment("staticChildren", "children", "children"),
		resourceids.UserSpecifiedSegment("termId1", "termId1Value"),
		resourceids.StaticSegment("staticRelations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Group Set Children Children Relation ID
func (id GroupSiteTermStoreGroupSetChildrenChildrenRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Group Id 1: %q", id.GroupId1),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Term Id 1: %q", id.TermId1),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Site Term Store Group Set Children Children Relation (%s)", strings.Join(components, "\n"))
}
