package groupsiteoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOperationId{}

// GroupSiteOperationId is a struct representing the Resource ID for a Group Site Operation
type GroupSiteOperationId struct {
	GroupId                    string
	SiteId                     string
	RichLongRunningOperationId string
}

// NewGroupSiteOperationID returns a new GroupSiteOperationId struct
func NewGroupSiteOperationID(groupId string, siteId string, richLongRunningOperationId string) GroupSiteOperationId {
	return GroupSiteOperationId{
		GroupId:                    groupId,
		SiteId:                     siteId,
		RichLongRunningOperationId: richLongRunningOperationId,
	}
}

// ParseGroupSiteOperationID parses 'input' into a GroupSiteOperationId
func ParseGroupSiteOperationID(input string) (*GroupSiteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.RichLongRunningOperationId, ok = parsed.Parsed["richLongRunningOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "richLongRunningOperationId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOperationIDInsensitively parses 'input' case-insensitively into a GroupSiteOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOperationIDInsensitively(input string) (*GroupSiteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.RichLongRunningOperationId, ok = parsed.Parsed["richLongRunningOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "richLongRunningOperationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOperationID checks that 'input' can be parsed as a Group Site Operation ID
func ValidateGroupSiteOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Operation ID
func (id GroupSiteOperationId) ID() string {
	fmtString := "/groups/%s/sites/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.RichLongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Operation ID
func (id GroupSiteOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("richLongRunningOperationId", "richLongRunningOperationIdValue"),
	}
}

// String returns a human-readable description of this Group Site Operation ID
func (id GroupSiteOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Rich Long Running Operation: %q", id.RichLongRunningOperationId),
	}
	return fmt.Sprintf("Group Site Operation (%s)", strings.Join(components, "\n"))
}
