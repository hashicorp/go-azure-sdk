package useronenotenotebooksectiongroupsectionpage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteNotebookSectionGroupSectionPageId{}

// UserOnenoteNotebookSectionGroupSectionPageId is a struct representing the Resource ID for a User Onenote Notebook Section Group Section Page
type UserOnenoteNotebookSectionGroupSectionPageId struct {
	UserId           string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewUserOnenoteNotebookSectionGroupSectionPageID returns a new UserOnenoteNotebookSectionGroupSectionPageId struct
func NewUserOnenoteNotebookSectionGroupSectionPageID(userId string, notebookId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) UserOnenoteNotebookSectionGroupSectionPageId {
	return UserOnenoteNotebookSectionGroupSectionPageId{
		UserId:           userId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseUserOnenoteNotebookSectionGroupSectionPageID parses 'input' into a UserOnenoteNotebookSectionGroupSectionPageId
func ParseUserOnenoteNotebookSectionGroupSectionPageID(input string) (*UserOnenoteNotebookSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionGroupSectionPageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteNotebookSectionGroupSectionPageIDInsensitively parses 'input' case-insensitively into a UserOnenoteNotebookSectionGroupSectionPageId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteNotebookSectionGroupSectionPageIDInsensitively(input string) (*UserOnenoteNotebookSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionGroupSectionPageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteNotebookSectionGroupSectionPageID checks that 'input' can be parsed as a User Onenote Notebook Section Group Section Page ID
func ValidateUserOnenoteNotebookSectionGroupSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteNotebookSectionGroupSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Notebook Section Group Section Page ID
func (id UserOnenoteNotebookSectionGroupSectionPageId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Notebook Section Group Section Page ID
func (id UserOnenoteNotebookSectionGroupSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Notebook Section Group Section Page ID
func (id UserOnenoteNotebookSectionGroupSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Onenote Notebook Section Group Section Page (%s)", strings.Join(components, "\n"))
}
