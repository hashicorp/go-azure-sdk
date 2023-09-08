package groupsiteonenotenotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteNotebookId{}

// GroupSiteOnenoteNotebookId is a struct representing the Resource ID for a Group Site Onenote Notebook
type GroupSiteOnenoteNotebookId struct {
	GroupId    string
	SiteId     string
	NotebookId string
}

// NewGroupSiteOnenoteNotebookID returns a new GroupSiteOnenoteNotebookId struct
func NewGroupSiteOnenoteNotebookID(groupId string, siteId string, notebookId string) GroupSiteOnenoteNotebookId {
	return GroupSiteOnenoteNotebookId{
		GroupId:    groupId,
		SiteId:     siteId,
		NotebookId: notebookId,
	}
}

// ParseGroupSiteOnenoteNotebookID parses 'input' into a GroupSiteOnenoteNotebookId
func ParseGroupSiteOnenoteNotebookID(input string) (*GroupSiteOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteNotebookId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteNotebookId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteNotebookIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteNotebookId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteNotebookIDInsensitively(input string) (*GroupSiteOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteNotebookId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteNotebookId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteNotebookID checks that 'input' can be parsed as a Group Site Onenote Notebook ID
func ValidateGroupSiteOnenoteNotebookID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteNotebookID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Notebook ID
func (id GroupSiteOnenoteNotebookId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Notebook ID
func (id GroupSiteOnenoteNotebookId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Notebook ID
func (id GroupSiteOnenoteNotebookId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
	}
	return fmt.Sprintf("Group Site Onenote Notebook (%s)", strings.Join(components, "\n"))
}
