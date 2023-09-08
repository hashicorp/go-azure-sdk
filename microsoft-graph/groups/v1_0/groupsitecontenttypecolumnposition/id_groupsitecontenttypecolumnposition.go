package groupsitecontenttypecolumnposition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteContentTypeColumnPositionId{}

// GroupSiteContentTypeColumnPositionId is a struct representing the Resource ID for a Group Site Content Type Column Position
type GroupSiteContentTypeColumnPositionId struct {
	GroupId            string
	SiteId             string
	ContentTypeId      string
	ColumnDefinitionId string
}

// NewGroupSiteContentTypeColumnPositionID returns a new GroupSiteContentTypeColumnPositionId struct
func NewGroupSiteContentTypeColumnPositionID(groupId string, siteId string, contentTypeId string, columnDefinitionId string) GroupSiteContentTypeColumnPositionId {
	return GroupSiteContentTypeColumnPositionId{
		GroupId:            groupId,
		SiteId:             siteId,
		ContentTypeId:      contentTypeId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupSiteContentTypeColumnPositionID parses 'input' into a GroupSiteContentTypeColumnPositionId
func ParseGroupSiteContentTypeColumnPositionID(input string) (*GroupSiteContentTypeColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteContentTypeColumnPositionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteContentTypeColumnPositionId{}

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

// ParseGroupSiteContentTypeColumnPositionIDInsensitively parses 'input' case-insensitively into a GroupSiteContentTypeColumnPositionId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteContentTypeColumnPositionIDInsensitively(input string) (*GroupSiteContentTypeColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteContentTypeColumnPositionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteContentTypeColumnPositionId{}

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

// ValidateGroupSiteContentTypeColumnPositionID checks that 'input' can be parsed as a Group Site Content Type Column Position ID
func ValidateGroupSiteContentTypeColumnPositionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteContentTypeColumnPositionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Content Type Column Position ID
func (id GroupSiteContentTypeColumnPositionId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentTypes/%s/columnPositions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentTypeId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Content Type Column Position ID
func (id GroupSiteContentTypeColumnPositionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
		resourceids.StaticSegment("staticColumnPositions", "columnPositions", "columnPositions"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Group Site Content Type Column Position ID
func (id GroupSiteContentTypeColumnPositionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Site Content Type Column Position (%s)", strings.Join(components, "\n"))
}
