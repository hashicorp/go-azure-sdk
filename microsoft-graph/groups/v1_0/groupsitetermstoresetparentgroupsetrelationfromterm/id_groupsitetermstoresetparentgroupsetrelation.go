package groupsitetermstoresetparentgroupsetrelationfromterm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetParentGroupSetRelationId{}

// GroupSiteTermStoreSetParentGroupSetRelationId is a struct representing the Resource ID for a Group Site Term Store Set Parent Group Set Relation
type GroupSiteTermStoreSetParentGroupSetRelationId struct {
	GroupId    string
	SiteId     string
	StoreId    string
	SetId      string
	SetId1     string
	RelationId string
}

// NewGroupSiteTermStoreSetParentGroupSetRelationID returns a new GroupSiteTermStoreSetParentGroupSetRelationId struct
func NewGroupSiteTermStoreSetParentGroupSetRelationID(groupId string, siteId string, storeId string, setId string, setId1 string, relationId string) GroupSiteTermStoreSetParentGroupSetRelationId {
	return GroupSiteTermStoreSetParentGroupSetRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		StoreId:    storeId,
		SetId:      setId,
		SetId1:     setId1,
		RelationId: relationId,
	}
}

// ParseGroupSiteTermStoreSetParentGroupSetRelationID parses 'input' into a GroupSiteTermStoreSetParentGroupSetRelationId
func ParseGroupSiteTermStoreSetParentGroupSetRelationID(input string) (*GroupSiteTermStoreSetParentGroupSetRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetRelationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.StoreId, ok = parsed.Parsed["storeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storeId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.SetId1, ok = parsed.Parsed["setId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId1", *parsed)
	}

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteTermStoreSetParentGroupSetRelationIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreSetParentGroupSetRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreSetParentGroupSetRelationIDInsensitively(input string) (*GroupSiteTermStoreSetParentGroupSetRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetParentGroupSetRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetParentGroupSetRelationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.StoreId, ok = parsed.Parsed["storeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storeId", *parsed)
	}

	if id.SetId, ok = parsed.Parsed["setId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId", *parsed)
	}

	if id.SetId1, ok = parsed.Parsed["setId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "setId1", *parsed)
	}

	if id.RelationId, ok = parsed.Parsed["relationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteTermStoreSetParentGroupSetRelationID checks that 'input' can be parsed as a Group Site Term Store Set Parent Group Set Relation ID
func ValidateGroupSiteTermStoreSetParentGroupSetRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreSetParentGroupSetRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Set Parent Group Set Relation ID
func (id GroupSiteTermStoreSetParentGroupSetRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStores/%s/sets/%s/parentGroup/sets/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.StoreId, id.SetId, id.SetId1, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Set Parent Group Set Relation ID
func (id GroupSiteTermStoreSetParentGroupSetRelationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStores", "termStores", "termStores"),
		resourceids.UserSpecifiedSegment("storeId", "storeIdValue"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("staticParentGroup", "parentGroup", "parentGroup"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId1", "setId1Value"),
		resourceids.StaticSegment("staticRelations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Set Parent Group Set Relation ID
func (id GroupSiteTermStoreSetParentGroupSetRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Store: %q", id.StoreId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Set Id 1: %q", id.SetId1),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Site Term Store Set Parent Group Set Relation (%s)", strings.Join(components, "\n"))
}
