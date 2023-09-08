package groupsitecontenttypecolumnposition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteContentTypeId{}

// GroupSiteContentTypeId is a struct representing the Resource ID for a Group Site Content Type
type GroupSiteContentTypeId struct {
	GroupId       string
	SiteId        string
	ContentTypeId string
}

// NewGroupSiteContentTypeID returns a new GroupSiteContentTypeId struct
func NewGroupSiteContentTypeID(groupId string, siteId string, contentTypeId string) GroupSiteContentTypeId {
	return GroupSiteContentTypeId{
		GroupId:       groupId,
		SiteId:        siteId,
		ContentTypeId: contentTypeId,
	}
}

// ParseGroupSiteContentTypeID parses 'input' into a GroupSiteContentTypeId
func ParseGroupSiteContentTypeID(input string) (*GroupSiteContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteContentTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteContentTypeId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteContentTypeIDInsensitively parses 'input' case-insensitively into a GroupSiteContentTypeId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteContentTypeIDInsensitively(input string) (*GroupSiteContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteContentTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteContentTypeId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteContentTypeID checks that 'input' can be parsed as a Group Site Content Type ID
func ValidateGroupSiteContentTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteContentTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Content Type ID
func (id GroupSiteContentTypeId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentTypes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentTypeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Content Type ID
func (id GroupSiteContentTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
	}
}

// String returns a human-readable description of this Group Site Content Type ID
func (id GroupSiteContentTypeId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
	}
	return fmt.Sprintf("Group Site Content Type (%s)", strings.Join(components, "\n"))
}
