package groupsitetermstoresetparentgroupsetchildrenchildrenrelationtoterm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId{}

// GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId is a struct representing the Resource ID for a Group Site Term Store Set Parent Group Set Children Children Relation
type GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId struct {
	GroupId    string
	SiteId     string
	SetId      string
	SetId1     string
	TermId     string
	TermId1    string
	RelationId string
}

// NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationID returns a new GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId struct
func NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationID(groupId string, siteId string, setId string, setId1 string, termId string, termId1 string, relationId string) GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId {
	return GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		SetId:      setId,
		SetId1:     setId1,
		TermId:     termId,
		TermId1:    termId1,
		RelationId: relationId,
	}
}

// ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationID parses 'input' into a GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId
func ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationID(input string) (*GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId{}

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

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationIDInsensitively(input string) (*GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId{}

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

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationID checks that 'input' can be parsed as a Group Site Term Store Set Parent Group Set Children Children Relation ID
func ValidateGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Set Parent Group Set Children Children Relation ID
func (id GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/parentGroup/sets/%s/children/%s/children/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.SetId1, id.TermId, id.TermId1, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Set Parent Group Set Children Children Relation ID
func (id GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticRelations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Set Parent Group Set Children Children Relation ID
func (id GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Set Id 1: %q", id.SetId1),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Term Id 1: %q", id.TermId1),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Site Term Store Set Parent Group Set Children Children Relation (%s)", strings.Join(components, "\n"))
}
