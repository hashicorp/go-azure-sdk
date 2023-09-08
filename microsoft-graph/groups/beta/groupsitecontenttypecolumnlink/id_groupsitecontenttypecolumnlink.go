package groupsitecontenttypecolumnlink

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteContentTypeColumnLinkId{}

// GroupSiteContentTypeColumnLinkId is a struct representing the Resource ID for a Group Site Content Type Column Link
type GroupSiteContentTypeColumnLinkId struct {
	GroupId       string
	SiteId        string
	ContentTypeId string
	ColumnLinkId  string
}

// NewGroupSiteContentTypeColumnLinkID returns a new GroupSiteContentTypeColumnLinkId struct
func NewGroupSiteContentTypeColumnLinkID(groupId string, siteId string, contentTypeId string, columnLinkId string) GroupSiteContentTypeColumnLinkId {
	return GroupSiteContentTypeColumnLinkId{
		GroupId:       groupId,
		SiteId:        siteId,
		ContentTypeId: contentTypeId,
		ColumnLinkId:  columnLinkId,
	}
}

// ParseGroupSiteContentTypeColumnLinkID parses 'input' into a GroupSiteContentTypeColumnLinkId
func ParseGroupSiteContentTypeColumnLinkID(input string) (*GroupSiteContentTypeColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteContentTypeColumnLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteContentTypeColumnLinkId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", *parsed)
	}

	if id.ColumnLinkId, ok = parsed.Parsed["columnLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnLinkId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteContentTypeColumnLinkIDInsensitively parses 'input' case-insensitively into a GroupSiteContentTypeColumnLinkId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteContentTypeColumnLinkIDInsensitively(input string) (*GroupSiteContentTypeColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteContentTypeColumnLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteContentTypeColumnLinkId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", *parsed)
	}

	if id.ColumnLinkId, ok = parsed.Parsed["columnLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnLinkId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteContentTypeColumnLinkID checks that 'input' can be parsed as a Group Site Content Type Column Link ID
func ValidateGroupSiteContentTypeColumnLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteContentTypeColumnLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Content Type Column Link ID
func (id GroupSiteContentTypeColumnLinkId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentTypes/%s/columnLinks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentTypeId, id.ColumnLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Content Type Column Link ID
func (id GroupSiteContentTypeColumnLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
		resourceids.StaticSegment("staticColumnLinks", "columnLinks", "columnLinks"),
		resourceids.UserSpecifiedSegment("columnLinkId", "columnLinkIdValue"),
	}
}

// String returns a human-readable description of this Group Site Content Type Column Link ID
func (id GroupSiteContentTypeColumnLinkId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Link: %q", id.ColumnLinkId),
	}
	return fmt.Sprintf("Group Site Content Type Column Link (%s)", strings.Join(components, "\n"))
}
