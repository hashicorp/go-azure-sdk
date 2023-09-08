package groupsiteonenoteresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteResourceId{}

// GroupSiteOnenoteResourceId is a struct representing the Resource ID for a Group Site Onenote Resource
type GroupSiteOnenoteResourceId struct {
	GroupId           string
	SiteId            string
	OnenoteResourceId string
}

// NewGroupSiteOnenoteResourceID returns a new GroupSiteOnenoteResourceId struct
func NewGroupSiteOnenoteResourceID(groupId string, siteId string, onenoteResourceId string) GroupSiteOnenoteResourceId {
	return GroupSiteOnenoteResourceId{
		GroupId:           groupId,
		SiteId:            siteId,
		OnenoteResourceId: onenoteResourceId,
	}
}

// ParseGroupSiteOnenoteResourceID parses 'input' into a GroupSiteOnenoteResourceId
func ParseGroupSiteOnenoteResourceID(input string) (*GroupSiteOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteResourceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.OnenoteResourceId, ok = parsed.Parsed["onenoteResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteResourceIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteResourceId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteResourceIDInsensitively(input string) (*GroupSiteOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteResourceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.OnenoteResourceId, ok = parsed.Parsed["onenoteResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteResourceID checks that 'input' can be parsed as a Group Site Onenote Resource ID
func ValidateGroupSiteOnenoteResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Resource ID
func (id GroupSiteOnenoteResourceId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/resources/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.OnenoteResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Resource ID
func (id GroupSiteOnenoteResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticResources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("onenoteResourceId", "onenoteResourceIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Resource ID
func (id GroupSiteOnenoteResourceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Onenote Resource: %q", id.OnenoteResourceId),
	}
	return fmt.Sprintf("Group Site Onenote Resource (%s)", strings.Join(components, "\n"))
}
