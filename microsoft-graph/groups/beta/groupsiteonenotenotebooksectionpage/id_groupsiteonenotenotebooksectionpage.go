package groupsiteonenotenotebooksectionpage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteNotebookSectionPageId{}

// GroupSiteOnenoteNotebookSectionPageId is a struct representing the Resource ID for a Group Site Onenote Notebook Section Page
type GroupSiteOnenoteNotebookSectionPageId struct {
	GroupId          string
	SiteId           string
	NotebookId       string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupSiteOnenoteNotebookSectionPageID returns a new GroupSiteOnenoteNotebookSectionPageId struct
func NewGroupSiteOnenoteNotebookSectionPageID(groupId string, siteId string, notebookId string, onenoteSectionId string, onenotePageId string) GroupSiteOnenoteNotebookSectionPageId {
	return GroupSiteOnenoteNotebookSectionPageId{
		GroupId:          groupId,
		SiteId:           siteId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupSiteOnenoteNotebookSectionPageID parses 'input' into a GroupSiteOnenoteNotebookSectionPageId
func ParseGroupSiteOnenoteNotebookSectionPageID(input string) (*GroupSiteOnenoteNotebookSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteNotebookSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteNotebookSectionPageId{}

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

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteNotebookSectionPageIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteNotebookSectionPageId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteNotebookSectionPageIDInsensitively(input string) (*GroupSiteOnenoteNotebookSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteNotebookSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteNotebookSectionPageId{}

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

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteNotebookSectionPageID checks that 'input' can be parsed as a Group Site Onenote Notebook Section Page ID
func ValidateGroupSiteOnenoteNotebookSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteNotebookSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Notebook Section Page ID
func (id GroupSiteOnenoteNotebookSectionPageId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Notebook Section Page ID
func (id GroupSiteOnenoteNotebookSectionPageId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Notebook Section Page ID
func (id GroupSiteOnenoteNotebookSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Site Onenote Notebook Section Page (%s)", strings.Join(components, "\n"))
}
