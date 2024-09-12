package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreSetIdChildId{}

// GroupIdSiteIdTermStoreSetIdChildId is a struct representing the Resource ID for a Group Id Site Id Term Store Set Id Child
type GroupIdSiteIdTermStoreSetIdChildId struct {
	GroupId string
	SiteId  string
	SetId   string
	TermId  string
}

// NewGroupIdSiteIdTermStoreSetIdChildID returns a new GroupIdSiteIdTermStoreSetIdChildId struct
func NewGroupIdSiteIdTermStoreSetIdChildID(groupId string, siteId string, setId string, termId string) GroupIdSiteIdTermStoreSetIdChildId {
	return GroupIdSiteIdTermStoreSetIdChildId{
		GroupId: groupId,
		SiteId:  siteId,
		SetId:   setId,
		TermId:  termId,
	}
}

// ParseGroupIdSiteIdTermStoreSetIdChildID parses 'input' into a GroupIdSiteIdTermStoreSetIdChildId
func ParseGroupIdSiteIdTermStoreSetIdChildID(input string) (*GroupIdSiteIdTermStoreSetIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdChildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreSetIdChildIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreSetIdChildId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreSetIdChildIDInsensitively(input string) (*GroupIdSiteIdTermStoreSetIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdChildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreSetIdChildId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdTermStoreSetIdChildID checks that 'input' can be parsed as a Group Id Site Id Term Store Set Id Child ID
func ValidateGroupIdSiteIdTermStoreSetIdChildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreSetIdChildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Set Id Child ID
func (id GroupIdSiteIdTermStoreSetIdChildId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/children/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.TermId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Set Id Child ID
func (id GroupIdSiteIdTermStoreSetIdChildId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("termStore", "termStore", "termStore"),
		resourceids.StaticSegment("sets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("termId", "termIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Set Id Child ID
func (id GroupIdSiteIdTermStoreSetIdChildId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Set Id Child (%s)", strings.Join(components, "\n"))
}
