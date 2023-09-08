package groupsiteonenotepageparentnotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenotePageId{}

// GroupSiteOnenotePageId is a struct representing the Resource ID for a Group Site Onenote Page
type GroupSiteOnenotePageId struct {
	GroupId       string
	SiteId        string
	OnenotePageId string
}

// NewGroupSiteOnenotePageID returns a new GroupSiteOnenotePageId struct
func NewGroupSiteOnenotePageID(groupId string, siteId string, onenotePageId string) GroupSiteOnenotePageId {
	return GroupSiteOnenotePageId{
		GroupId:       groupId,
		SiteId:        siteId,
		OnenotePageId: onenotePageId,
	}
}

// ParseGroupSiteOnenotePageID parses 'input' into a GroupSiteOnenotePageId
func ParseGroupSiteOnenotePageID(input string) (*GroupSiteOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenotePageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenotePageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenotePageIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenotePageId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenotePageIDInsensitively(input string) (*GroupSiteOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenotePageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenotePageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenotePageID checks that 'input' can be parsed as a Group Site Onenote Page ID
func ValidateGroupSiteOnenotePageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenotePageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Page ID
func (id GroupSiteOnenotePageId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Page ID
func (id GroupSiteOnenotePageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Page ID
func (id GroupSiteOnenotePageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Site Onenote Page (%s)", strings.Join(components, "\n"))
}
