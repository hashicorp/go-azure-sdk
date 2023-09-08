package groupsiteonenotenotebooksectionparentsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteNotebookSectionId{}

// GroupSiteOnenoteNotebookSectionId is a struct representing the Resource ID for a Group Site Onenote Notebook Section
type GroupSiteOnenoteNotebookSectionId struct {
	GroupId          string
	SiteId           string
	NotebookId       string
	OnenoteSectionId string
}

// NewGroupSiteOnenoteNotebookSectionID returns a new GroupSiteOnenoteNotebookSectionId struct
func NewGroupSiteOnenoteNotebookSectionID(groupId string, siteId string, notebookId string, onenoteSectionId string) GroupSiteOnenoteNotebookSectionId {
	return GroupSiteOnenoteNotebookSectionId{
		GroupId:          groupId,
		SiteId:           siteId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupSiteOnenoteNotebookSectionID parses 'input' into a GroupSiteOnenoteNotebookSectionId
func ParseGroupSiteOnenoteNotebookSectionID(input string) (*GroupSiteOnenoteNotebookSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteNotebookSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteNotebookSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteNotebookSectionIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteNotebookSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteNotebookSectionIDInsensitively(input string) (*GroupSiteOnenoteNotebookSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteNotebookSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteNotebookSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteNotebookSectionID checks that 'input' can be parsed as a Group Site Onenote Notebook Section ID
func ValidateGroupSiteOnenoteNotebookSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteNotebookSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Notebook Section ID
func (id GroupSiteOnenoteNotebookSectionId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Notebook Section ID
func (id GroupSiteOnenoteNotebookSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Notebook Section ID
func (id GroupSiteOnenoteNotebookSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Site Onenote Notebook Section (%s)", strings.Join(components, "\n"))
}
