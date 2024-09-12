package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{}

// GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId is a struct representing the Resource ID for a Group Id Site Id Term Store Set Id Term Id Child Id Relation
type GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId struct {
	GroupId    string
	SiteId     string
	SetId      string
	TermId     string
	TermId1    string
	RelationId string
}

// NewGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID returns a new GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId struct
func NewGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID(groupId string, siteId string, setId string, termId string, termId1 string, relationId string) GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId {
	return GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		SetId:      setId,
		TermId:     termId,
		TermId1:    termId1,
		RelationId: relationId,
	}
}

// ParseGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID parses 'input' into a GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId
func ParseGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID(input string) (*GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationIDInsensitively(input string) (*GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.SetId, ok = input.Parsed["setId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "setId", input)
	}

	if id.TermId, ok = input.Parsed["termId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termId", input)
	}

	if id.TermId1, ok = input.Parsed["termId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termId1", input)
	}

	if id.RelationId, ok = input.Parsed["relationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "relationId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID checks that 'input' can be parsed as a Group Id Site Id Term Store Set Id Term Id Child Id Relation ID
func ValidateGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Set Id Term Id Child Id Relation ID
func (id GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/terms/%s/children/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.TermId, id.TermId1, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Set Id Term Id Child Id Relation ID
func (id GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("termStore", "termStore", "termStore"),
		resourceids.StaticSegment("sets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("terms", "terms", "terms"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("termId1", "termId1Value"),
		resourceids.StaticSegment("relations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Set Id Term Id Child Id Relation ID
func (id GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Term Id 1: %q", id.TermId1),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Set Id Term Id Child Id Relation (%s)", strings.Join(components, "\n"))
}
