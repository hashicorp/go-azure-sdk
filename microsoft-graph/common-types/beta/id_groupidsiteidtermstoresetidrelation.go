package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreSetIdRelationId{}

// GroupIdSiteIdTermStoreSetIdRelationId is a struct representing the Resource ID for a Group Id Site Id Term Store Set Id Relation
type GroupIdSiteIdTermStoreSetIdRelationId struct {
	GroupId    string
	SiteId     string
	SetId      string
	RelationId string
}

// NewGroupIdSiteIdTermStoreSetIdRelationID returns a new GroupIdSiteIdTermStoreSetIdRelationId struct
func NewGroupIdSiteIdTermStoreSetIdRelationID(groupId string, siteId string, setId string, relationId string) GroupIdSiteIdTermStoreSetIdRelationId {
	return GroupIdSiteIdTermStoreSetIdRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		SetId:      setId,
		RelationId: relationId,
	}
}

// ParseGroupIdSiteIdTermStoreSetIdRelationID parses 'input' into a GroupIdSiteIdTermStoreSetIdRelationId
func ParseGroupIdSiteIdTermStoreSetIdRelationID(input string) (*GroupIdSiteIdTermStoreSetIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreSetIdRelationIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreSetIdRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreSetIdRelationIDInsensitively(input string) (*GroupIdSiteIdTermStoreSetIdRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreSetIdRelationId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RelationId, ok = input.Parsed["relationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "relationId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdTermStoreSetIdRelationID checks that 'input' can be parsed as a Group Id Site Id Term Store Set Id Relation ID
func ValidateGroupIdSiteIdTermStoreSetIdRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreSetIdRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Set Id Relation ID
func (id GroupIdSiteIdTermStoreSetIdRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Set Id Relation ID
func (id GroupIdSiteIdTermStoreSetIdRelationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("termStore", "termStore", "termStore"),
		resourceids.StaticSegment("sets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("relations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Set Id Relation ID
func (id GroupIdSiteIdTermStoreSetIdRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Set Id Relation (%s)", strings.Join(components, "\n"))
}
