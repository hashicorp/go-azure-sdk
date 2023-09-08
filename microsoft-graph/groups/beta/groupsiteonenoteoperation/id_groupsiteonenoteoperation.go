package groupsiteonenoteoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteOperationId{}

// GroupSiteOnenoteOperationId is a struct representing the Resource ID for a Group Site Onenote Operation
type GroupSiteOnenoteOperationId struct {
	GroupId            string
	SiteId             string
	OnenoteOperationId string
}

// NewGroupSiteOnenoteOperationID returns a new GroupSiteOnenoteOperationId struct
func NewGroupSiteOnenoteOperationID(groupId string, siteId string, onenoteOperationId string) GroupSiteOnenoteOperationId {
	return GroupSiteOnenoteOperationId{
		GroupId:            groupId,
		SiteId:             siteId,
		OnenoteOperationId: onenoteOperationId,
	}
}

// ParseGroupSiteOnenoteOperationID parses 'input' into a GroupSiteOnenoteOperationId
func ParseGroupSiteOnenoteOperationID(input string) (*GroupSiteOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.OnenoteOperationId, ok = parsed.Parsed["onenoteOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteOperationIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteOperationIDInsensitively(input string) (*GroupSiteOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.OnenoteOperationId, ok = parsed.Parsed["onenoteOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteOperationID checks that 'input' can be parsed as a Group Site Onenote Operation ID
func ValidateGroupSiteOnenoteOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Operation ID
func (id GroupSiteOnenoteOperationId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.OnenoteOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Operation ID
func (id GroupSiteOnenoteOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("onenoteOperationId", "onenoteOperationIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Operation ID
func (id GroupSiteOnenoteOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Onenote Operation: %q", id.OnenoteOperationId),
	}
	return fmt.Sprintf("Group Site Onenote Operation (%s)", strings.Join(components, "\n"))
}
