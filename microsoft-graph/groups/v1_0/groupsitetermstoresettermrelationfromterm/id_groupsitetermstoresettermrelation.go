package groupsitetermstoresettermrelationfromterm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetTermRelationId{}

// GroupSiteTermStoreSetTermRelationId is a struct representing the Resource ID for a Group Site Term Store Set Term Relation
type GroupSiteTermStoreSetTermRelationId struct {
	GroupId    string
	SiteId     string
	SetId      string
	TermId     string
	RelationId string
}

// NewGroupSiteTermStoreSetTermRelationID returns a new GroupSiteTermStoreSetTermRelationId struct
func NewGroupSiteTermStoreSetTermRelationID(groupId string, siteId string, setId string, termId string, relationId string) GroupSiteTermStoreSetTermRelationId {
	return GroupSiteTermStoreSetTermRelationId{
		GroupId:    groupId,
		SiteId:     siteId,
		SetId:      setId,
		TermId:     termId,
		RelationId: relationId,
	}
}

// ParseGroupSiteTermStoreSetTermRelationID parses 'input' into a GroupSiteTermStoreSetTermRelationId
func ParseGroupSiteTermStoreSetTermRelationID(input string) (*GroupSiteTermStoreSetTermRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetTermRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetTermRelationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
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

// ParseGroupSiteTermStoreSetTermRelationIDInsensitively parses 'input' case-insensitively into a GroupSiteTermStoreSetTermRelationId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteTermStoreSetTermRelationIDInsensitively(input string) (*GroupSiteTermStoreSetTermRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteTermStoreSetTermRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteTermStoreSetTermRelationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
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

// ValidateGroupSiteTermStoreSetTermRelationID checks that 'input' can be parsed as a Group Site Term Store Set Term Relation ID
func ValidateGroupSiteTermStoreSetTermRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteTermStoreSetTermRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Term Store Set Term Relation ID
func (id GroupSiteTermStoreSetTermRelationId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/terms/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.TermId, id.RelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Term Store Set Term Relation ID
func (id GroupSiteTermStoreSetTermRelationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticTermStore", "termStore", "termStore"),
		resourceids.StaticSegment("staticSets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("staticTerms", "terms", "terms"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
		resourceids.StaticSegment("staticRelations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationId", "relationIdValue"),
	}
}

// String returns a human-readable description of this Group Site Term Store Set Term Relation ID
func (id GroupSiteTermStoreSetTermRelationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Relation: %q", id.RelationId),
	}
	return fmt.Sprintf("Group Site Term Store Set Term Relation (%s)", strings.Join(components, "\n"))
}
