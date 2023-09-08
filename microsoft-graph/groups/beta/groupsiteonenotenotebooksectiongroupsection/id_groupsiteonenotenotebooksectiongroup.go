package groupsiteonenotenotebooksectiongroupsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteNotebookSectionGroupId{}

// GroupSiteOnenoteNotebookSectionGroupId is a struct representing the Resource ID for a Group Site Onenote Notebook Section Group
type GroupSiteOnenoteNotebookSectionGroupId struct {
	GroupId        string
	SiteId         string
	NotebookId     string
	SectionGroupId string
}

// NewGroupSiteOnenoteNotebookSectionGroupID returns a new GroupSiteOnenoteNotebookSectionGroupId struct
func NewGroupSiteOnenoteNotebookSectionGroupID(groupId string, siteId string, notebookId string, sectionGroupId string) GroupSiteOnenoteNotebookSectionGroupId {
	return GroupSiteOnenoteNotebookSectionGroupId{
		GroupId:        groupId,
		SiteId:         siteId,
		NotebookId:     notebookId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseGroupSiteOnenoteNotebookSectionGroupID parses 'input' into a GroupSiteOnenoteNotebookSectionGroupId
func ParseGroupSiteOnenoteNotebookSectionGroupID(input string) (*GroupSiteOnenoteNotebookSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteNotebookSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteNotebookSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteNotebookSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteNotebookSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteNotebookSectionGroupIDInsensitively(input string) (*GroupSiteOnenoteNotebookSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteNotebookSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteNotebookSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteNotebookSectionGroupID checks that 'input' can be parsed as a Group Site Onenote Notebook Section Group ID
func ValidateGroupSiteOnenoteNotebookSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteNotebookSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Notebook Section Group ID
func (id GroupSiteOnenoteNotebookSectionGroupId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Notebook Section Group ID
func (id GroupSiteOnenoteNotebookSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Notebook Section Group ID
func (id GroupSiteOnenoteNotebookSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Group Site Onenote Notebook Section Group (%s)", strings.Join(components, "\n"))
}
