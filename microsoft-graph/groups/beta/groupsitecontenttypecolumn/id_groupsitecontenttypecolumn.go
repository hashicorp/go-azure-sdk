package groupsitecontenttypecolumn

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteContentTypeColumnId{}

// GroupSiteContentTypeColumnId is a struct representing the Resource ID for a Group Site Content Type Column
type GroupSiteContentTypeColumnId struct {
	GroupId            string
	SiteId             string
	ContentTypeId      string
	ColumnDefinitionId string
}

// NewGroupSiteContentTypeColumnID returns a new GroupSiteContentTypeColumnId struct
func NewGroupSiteContentTypeColumnID(groupId string, siteId string, contentTypeId string, columnDefinitionId string) GroupSiteContentTypeColumnId {
	return GroupSiteContentTypeColumnId{
		GroupId:            groupId,
		SiteId:             siteId,
		ContentTypeId:      contentTypeId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupSiteContentTypeColumnID parses 'input' into a GroupSiteContentTypeColumnId
func ParseGroupSiteContentTypeColumnID(input string) (*GroupSiteContentTypeColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteContentTypeColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteContentTypeColumnId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", *parsed)
	}

	if id.ColumnDefinitionId, ok = parsed.Parsed["columnDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteContentTypeColumnIDInsensitively parses 'input' case-insensitively into a GroupSiteContentTypeColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteContentTypeColumnIDInsensitively(input string) (*GroupSiteContentTypeColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteContentTypeColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteContentTypeColumnId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", *parsed)
	}

	if id.ColumnDefinitionId, ok = parsed.Parsed["columnDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteContentTypeColumnID checks that 'input' can be parsed as a Group Site Content Type Column ID
func ValidateGroupSiteContentTypeColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteContentTypeColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Content Type Column ID
func (id GroupSiteContentTypeColumnId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentTypes/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentTypeId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Content Type Column ID
func (id GroupSiteContentTypeColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
		resourceids.StaticSegment("staticColumns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Group Site Content Type Column ID
func (id GroupSiteContentTypeColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Site Content Type Column (%s)", strings.Join(components, "\n"))
}
