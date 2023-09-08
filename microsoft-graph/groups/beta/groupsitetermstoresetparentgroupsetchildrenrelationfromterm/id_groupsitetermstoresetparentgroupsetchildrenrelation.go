package groupsitetermstoresetparentgroupsetchildrenrelationfromterm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetParentGroupSetChildrenRelationId{}

// GroupSiteTermStoreSetParentGroupSetChildrenRelationId is a struct representing the Resource ID for a Group Site Term Store Set Parent Group Set Children Relation
type GroupSiteTermStoreSetParentGroupSetChildrenRelationId struct {
	GroupId    string
	SiteId     string
	SetId      string
	SetId1     string
	TermId     string
	RelationId string
}

// NewGroupSiteTermStoreSetParentGroupSetChildrenRelationID returns a new GroupSiteTermStoreSetParentGroupSetChildrenRelationId struct
func NewGroupSiteTermStoreSetParentGroupSetChildrenRelationID(groupId string, siteId string, setId string, setId1 string, termId string, relationId string) GroupSiteTermStoreSetParentGroupSetChildrenRelationId {
	return GroupSiteTermStoreSetParentGroupSetChildrenRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		SetId:      setId,
		SetId1:     setId1,
		TermId:     termId,
		RelationId: relationId,
	}
}

// ParseGroupSiteTermStoreSetParentGroupSetChildrenRelationID parses 'input' into a GroupSiteTermStoreSetParentGroupSetChildrenRelationId
func ParseGroupSiteTermStoreSetParentGroupSetChildrenRelationID(input string) (*GroupSiteTermStoreSetParentGroupSetChildrenRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetChildrenRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetChildrenRelationId{}

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

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreSetParentGroupSetChildrenRelationIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreSetParentGroupSetChildrenRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreSetParentGroupSetChildrenRelationIDInsensitively(input string) (*GroupSiteTermStoreSetParentGroupSetChildrenRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetChildrenRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetChildrenRelationId{}

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

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreSetParentGroupSetChildrenRelationID checks that 'input' can be parsed as a Group Site Term Store Set Parent Group Set Children Relation ID
func ValidateGroupSiteTermStoreSetParentGroupSetChildrenRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreSetParentGroupSetChildrenRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Set Parent Group Set Children Relation ID
func (id GroupSiteTermStoreSetParentGroupSetChildrenRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/parentGroup/sets/%s/children/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.SetId1, id.TermId, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Set Parent Group Set Children Relation ID
func (id GroupSiteTermStoreSetParentGroupSetChildrenRelationId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticRelations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Set Parent Group Set Children Relation ID
func (id GroupSiteTermStoreSetParentGroupSetChildrenRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Set Id 1: %q", id.SetId1),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Site Term Store Set Parent Group Set Children Relation (%s)", strings.Join(components, "\n"))
}
