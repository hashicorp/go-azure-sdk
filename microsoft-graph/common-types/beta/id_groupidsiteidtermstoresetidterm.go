package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreSetIdTermId{}

// GroupIdSiteIdTermStoreSetIdTermId is a struct representing the Resource ID for a Group Id Site Id Term Store Set Id Term
type GroupIdSiteIdTermStoreSetIdTermId struct {
	GroupId string
	SiteId  string
	SetId   string
	TermId  string
}

// NewGroupIdSiteIdTermStoreSetIdTermID returns a new GroupIdSiteIdTermStoreSetIdTermId struct
func NewGroupIdSiteIdTermStoreSetIdTermID(groupId string, siteId string, setId string, termId string) GroupIdSiteIdTermStoreSetIdTermId {
	return GroupIdSiteIdTermStoreSetIdTermId{
		GroupId: groupId,
		SiteId:  siteId,
		SetId:   setId,
		TermId:  termId,
	}
}

// ParseGroupIdSiteIdTermStoreSetIdTermID parses 'input' into a GroupIdSiteIdTermStoreSetIdTermId
func ParseGroupIdSiteIdTermStoreSetIdTermID(input string) (*GroupIdSiteIdTermStoreSetIdTermId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdTermId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdTermId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreSetIdTermIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreSetIdTermId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreSetIdTermIDInsensitively(input string) (*GroupIdSiteIdTermStoreSetIdTermId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetIdTermId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetIdTermId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreSetIdTermId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdSiteIdTermStoreSetIdTermID checks that 'input' can be parsed as a Group Id Site Id Term Store Set Id Term ID
func ValidateGroupIdSiteIdTermStoreSetIdTermID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreSetIdTermID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Set Id Term ID
func (id GroupIdSiteIdTermStoreSetIdTermId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s/terms/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId, id.TermId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Set Id Term ID
func (id GroupIdSiteIdTermStoreSetIdTermId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Set Id Term ID
func (id GroupIdSiteIdTermStoreSetIdTermId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
		fmt.Sprintf("Term: %q", id.TermId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Set Id Term (%s)", strings.Join(components, "\n"))
}
