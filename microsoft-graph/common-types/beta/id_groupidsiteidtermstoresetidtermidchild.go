package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreSetIdTermIdChildId{}

// GroupIdSiteIdTermStoreSetIdTermIdChildId is a struct representing the Resource ID for a Group Id Site Id Term Store Set Id Term Id Child
type GroupIdSiteIdTermStoreSetIdTermIdChildId struct {
	GroupId string
	SiteId  string
	SetId   string
	TermId  string
	TermId1 string
}

// NewGroupIdSiteIdTermStoreSetIdTermIdChildID returns a new GroupIdSiteIdTermStoreSetIdTermIdChildId struct
func NewGroupIdSiteIdTermStoreSetIdTermIdChildID(groupId string, siteId string, setId string, termId string, termId1 string) GroupIdSiteIdTermStoreSetIdTermIdChildId {
	return GroupIdSiteIdTermStoreSetIdTermIdChildId{
		GroupId: groupId,
		SiteId:  siteId,
		SetId:   setId,
		TermId:  termId,
		TermId1: termId1,
	}
}

// ParseGroupIdSiteIdTermStoreSetIdTermIdChildID parses 'input' into a GroupIdSiteIdTermStoreSetIdTermIdChildId
func ParseGroupIdSiteIdTermStoreSetIdTermIdChildID(input string) (*GroupIdSiteIdTermStoreSetIdTermIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdTermIdChildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdTermIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreSetIdTermIdChildIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreSetIdTermIdChildId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreSetIdTermIdChildIDInsensitively(input string) (*GroupIdSiteIdTermStoreSetIdTermIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdTermIdChildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdTermIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreSetIdTermIdChildId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdTermStoreSetIdTermIdChildID checks that 'input' can be parsed as a Group Id Site Id Term Store Set Id Term Id Child ID
func ValidateGroupIdSiteIdTermStoreSetIdTermIdChildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreSetIdTermIdChildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Set Id Term Id Child ID
func (id GroupIdSiteIdTermStoreSetIdTermIdChildId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/terms/%s/children/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.TermId, id.TermId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Set Id Term Id Child ID
func (id GroupIdSiteIdTermStoreSetIdTermIdChildId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Set Id Term Id Child ID
func (id GroupIdSiteIdTermStoreSetIdTermIdChildId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
		fmt.Sprintf("Term Id 1: %q", id.TermId1),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Set Id Term Id Child (%s)", strings.Join(components, "\n"))
}
