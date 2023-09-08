package groupsitetermstoregroupsettermrelation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreGroupSetTermRelationId{}

// GroupSiteTermStoreGroupSetTermRelationId is a struct representing the Resource ID for a Group Site Term Store Group Set Term Relation
type GroupSiteTermStoreGroupSetTermRelationId struct {
	GroupId    string
	SiteId     string
	StoreId    string
	GroupId1   string
	SetId      string
	TermId     string
	RelationId string
}

// NewGroupSiteTermStoreGroupSetTermRelationID returns a new GroupSiteTermStoreGroupSetTermRelationId struct
func NewGroupSiteTermStoreGroupSetTermRelationID(groupId string, siteId string, storeId string, groupId1 string, setId string, termId string, relationId string) GroupSiteTermStoreGroupSetTermRelationId {
	return GroupSiteTermStoreGroupSetTermRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		StoreId:    storeId,
		GroupId1:   groupId1,
		SetId:      setId,
		TermId:     termId,
		RelationId: relationId,
	}
}

// ParseGroupSiteTermStoreGroupSetTermRelationID parses 'input' into a GroupSiteTermStoreGroupSetTermRelationId
func ParseGroupSiteTermStoreGroupSetTermRelationID(input string) (*GroupSiteTermStoreGroupSetTermRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreGroupSetTermRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreGroupSetTermRelationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.StoreId, ok = parsed.Parsed["storeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storeId", *parsed)
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

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreGroupSetTermRelationIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreGroupSetTermRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreGroupSetTermRelationIDInsensitively(input string) (*GroupSiteTermStoreGroupSetTermRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreGroupSetTermRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreGroupSetTermRelationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.StoreId, ok = parsed.Parsed["storeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storeId", *parsed)
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

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreGroupSetTermRelationID checks that 'input' can be parsed as a Group Site Term Store Group Set Term Relation ID
func ValidateGroupSiteTermStoreGroupSetTermRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreGroupSetTermRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Group Set Term Relation ID
func (id GroupSiteTermStoreGroupSetTermRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStores/%s/groups/%s/sets/%s/terms/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.StoreId, id.GroupId1, id.SetId, id.TermId, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Group Set Term Relation ID
func (id GroupSiteTermStoreGroupSetTermRelationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStores", "termStores", "termStores"),
		resourceids.UserSpecifiedSegment("storeId", "storeIdValue"),
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId1", "groupId1Value"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("staticTerms", "terms", "terms"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
		resourceids.StaticSegment("staticRelations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Group Set Term Relation ID
func (id GroupSiteTermStoreGroupSetTermRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Store: %q", id.StoreId),
		fmt.Sprintf("Group Id 1: %q", id.GroupId1),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Site Term Store Group Set Term Relation (%s)", strings.Join(components, "\n"))
}
