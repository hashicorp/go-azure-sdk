package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreIdSetIdRelationId{}

// GroupIdSiteIdTermStoreIdSetIdRelationId is a struct representing the Resource ID for a Group Id Site Id Term Store Id Set Id Relation
type GroupIdSiteIdTermStoreIdSetIdRelationId struct {
	GroupId    string
	SiteId     string
	StoreId    string
	SetId      string
	RelationId string
}

// NewGroupIdSiteIdTermStoreIdSetIdRelationID returns a new GroupIdSiteIdTermStoreIdSetIdRelationId struct
func NewGroupIdSiteIdTermStoreIdSetIdRelationID(groupId string, siteId string, storeId string, setId string, relationId string) GroupIdSiteIdTermStoreIdSetIdRelationId {
	return GroupIdSiteIdTermStoreIdSetIdRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		StoreId:    storeId,
		SetId:      setId,
		RelationId: relationId,
	}
}

// ParseGroupIdSiteIdTermStoreIdSetIdRelationID parses 'input' into a GroupIdSiteIdTermStoreIdSetIdRelationId
func ParseGroupIdSiteIdTermStoreIdSetIdRelationID(input string) (*GroupIdSiteIdTermStoreIdSetIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreIdSetIdRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreIdSetIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreIdSetIdRelationIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreIdSetIdRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreIdSetIdRelationIDInsensitively(input string) (*GroupIdSiteIdTermStoreIdSetIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreIdSetIdRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreIdSetIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreIdSetIdRelationId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RelationId, ok = input.Parsed["relationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "relationId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdTermStoreIdSetIdRelationID checks that 'input' can be parsed as a Group Id Site Id Term Store Id Set Id Relation ID
func ValidateGroupIdSiteIdTermStoreIdSetIdRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreIdSetIdRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Id Set Id Relation ID
func (id GroupIdSiteIdTermStoreIdSetIdRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStores/%s/sets/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.StoreId, id.SetId, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Id Set Id Relation ID
func (id GroupIdSiteIdTermStoreIdSetIdRelationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("termStores", "termStores", "termStores"),
		resourceids.UserSpecifiedSegment("storeId", "storeIdValue"),
		resourceids.StaticSegment("sets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("relations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Id Set Id Relation ID
func (id GroupIdSiteIdTermStoreIdSetIdRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Store: %q", id.StoreId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Id Set Id Relation (%s)", strings.Join(components, "\n"))
}
