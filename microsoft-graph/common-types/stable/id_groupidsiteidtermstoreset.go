package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreSetId{}

// GroupIdSiteIdTermStoreSetId is a struct representing the Resource ID for a Group Id Site Id Term Store Set
type GroupIdSiteIdTermStoreSetId struct {
	GroupId string
	SiteId  string
	SetId   string
}

// NewGroupIdSiteIdTermStoreSetID returns a new GroupIdSiteIdTermStoreSetId struct
func NewGroupIdSiteIdTermStoreSetID(groupId string, siteId string, setId string) GroupIdSiteIdTermStoreSetId {
	return GroupIdSiteIdTermStoreSetId{
		GroupId: groupId,
		SiteId:  siteId,
		SetId:   setId,
	}
}

// ParseGroupIdSiteIdTermStoreSetID parses 'input' into a GroupIdSiteIdTermStoreSetId
func ParseGroupIdSiteIdTermStoreSetID(input string) (*GroupIdSiteIdTermStoreSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdTermStoreSetIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdTermStoreSetId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdTermStoreSetIDInsensitively(input string) (*GroupIdSiteIdTermStoreSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdTermStoreSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdTermStoreSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdTermStoreSetId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdTermStoreSetID checks that 'input' can be parsed as a Group Id Site Id Term Store Set ID
func ValidateGroupIdSiteIdTermStoreSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdTermStoreSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Term Store Set ID
func (id GroupIdSiteIdTermStoreSetId) ID() string {
	fmtString := "/groups/%s/sites/%s/termStore/sets/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Term Store Set ID
func (id GroupIdSiteIdTermStoreSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("termStore", "termStore", "termStore"),
		resourceids.StaticSegment("sets", "sets", "sets"),
		resourceids.UserSpecifiedSegment("setId", "setIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Term Store Set ID
func (id GroupIdSiteIdTermStoreSetId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Set: %q", id.SetId),
	}
	return fmt.Sprintf("Group Id Site Id Term Store Set (%s)", strings.Join(components, "\n"))
}
