package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreIdSetIdTermIdRelationId{}

// GroupIdSiteIdTermStoreIdSetIdTermIdRelationId is a struct representing the Resource ID for a Group Id Site Id Term Store Id Set Id Term Id Relation
type GroupIdSiteIdTermStoreIdSetIdTermIdRelationId struct {
	GroupId    string
	SiteId     string
	StoreId    string
	SetId      string
	TermId     string
	RelationId string
}

// NewGroupIdSiteIdTermStoreIdSetIdTermIdRelationID returns a new GroupIdSiteIdTermStoreIdSetIdTermIdRelationId struct
func NewGroupIdSiteIdTermStoreIdSetIdTermIdRelationID(groupId string, siteId string, storeId string, setId string, termId string, relationId string) GroupIdSiteIdTermStoreIdSetIdTermIdRelationId {
	return GroupIdSiteIdTermStoreIdSetIdTermIdRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		StoreId:    storeId,
		SetId:      setId,
		TermId:     termId,
		RelationId: relationId,
	}
}

// ParseGroupIdSiteIdTermStoreIdSetIdTermIdRelationID parses 'input' into a GroupIdSiteIdTermStoreIdSetIdTermIdRelationId
func ParseGroupIdSiteIdTermStoreIdSetIdTermIdRelationID(input string) (*GroupIdSiteIdTermStoreIdSetIdTermIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreIdSetIdTermIdRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreIdSetIdTermIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreIdSetIdTermIdRelationIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreIdSetIdTermIdRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreIdSetIdTermIdRelationIDInsensitively(input string) (*GroupIdSiteIdTermStoreIdSetIdTermIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreIdSetIdTermIdRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreIdSetIdTermIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreIdSetIdTermIdRelationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.StoreId, ok = input.Parsed["storeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "storeId", input)
	}

	if id.SetId, ok = input.Parsed["setId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "setId", input)
	}

	if id.TermId, ok = input.Parsed["termId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termId", input)
	}

	if id.RelationId, ok = input.Parsed["relationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "relationId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdTermStoreIdSetIdTermIdRelationID checks that 'input' can be parsed as a Group Id Site Id Term Store Id Set Id Term Id Relation ID
func ValidateGroupIdSiteIdTermStoreIdSetIdTermIdRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreIdSetIdTermIdRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Id Set Id Term Id Relation ID
func (id GroupIdSiteIdTermStoreIdSetIdTermIdRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStores/%s/sets/%s/terms/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.StoreId, id.SetId, id.TermId, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Id Set Id Term Id Relation ID
func (id GroupIdSiteIdTermStoreIdSetIdTermIdRelationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("termStores", "termStores", "termStores"),
		resourceids.UserSpecifiedSegment("storeId", "storeIdValue"),
		resourceids.StaticSegment("sets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("terms", "terms", "terms"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
		resourceids.StaticSegment("relations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Id Set Id Term Id Relation ID
func (id GroupIdSiteIdTermStoreIdSetIdTermIdRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Store: %q", id.StoreId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Id Set Id Term Id Relation (%s)", strings.Join(components, "\n"))
}
