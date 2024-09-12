package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreSetIdTermIdRelationId{}

// GroupIdSiteIdTermStoreSetIdTermIdRelationId is a struct representing the Resource ID for a Group Id Site Id Term Store Set Id Term Id Relation
type GroupIdSiteIdTermStoreSetIdTermIdRelationId struct {
	GroupId    string
	SiteId     string
	SetId      string
	TermId     string
	RelationId string
}

// NewGroupIdSiteIdTermStoreSetIdTermIdRelationID returns a new GroupIdSiteIdTermStoreSetIdTermIdRelationId struct
func NewGroupIdSiteIdTermStoreSetIdTermIdRelationID(groupId string, siteId string, setId string, termId string, relationId string) GroupIdSiteIdTermStoreSetIdTermIdRelationId {
	return GroupIdSiteIdTermStoreSetIdTermIdRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		SetId:      setId,
		TermId:     termId,
		RelationId: relationId,
	}
}

// ParseGroupIdSiteIdTermStoreSetIdTermIdRelationID parses 'input' into a GroupIdSiteIdTermStoreSetIdTermIdRelationId
func ParseGroupIdSiteIdTermStoreSetIdTermIdRelationID(input string) (*GroupIdSiteIdTermStoreSetIdTermIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdTermIdRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdTermIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreSetIdTermIdRelationIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreSetIdTermIdRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreSetIdTermIdRelationIDInsensitively(input string) (*GroupIdSiteIdTermStoreSetIdTermIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdTermIdRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdTermIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreSetIdTermIdRelationId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RelationId, ok = input.Parsed["relationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "relationId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdTermStoreSetIdTermIdRelationID checks that 'input' can be parsed as a Group Id Site Id Term Store Set Id Term Id Relation ID
func ValidateGroupIdSiteIdTermStoreSetIdTermIdRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreSetIdTermIdRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Set Id Term Id Relation ID
func (id GroupIdSiteIdTermStoreSetIdTermIdRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/terms/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.TermId, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Set Id Term Id Relation ID
func (id GroupIdSiteIdTermStoreSetIdTermIdRelationId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("relations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Set Id Term Id Relation ID
func (id GroupIdSiteIdTermStoreSetIdTermIdRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Set Id Term Id Relation (%s)", strings.Join(components, "\n"))
}
