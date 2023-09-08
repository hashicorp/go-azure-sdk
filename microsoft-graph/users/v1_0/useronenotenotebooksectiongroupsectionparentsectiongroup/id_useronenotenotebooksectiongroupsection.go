package useronenotenotebooksectiongroupsectionparentsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteNotebookSectionGroupSectionId{}

// UserOnenoteNotebookSectionGroupSectionId is a struct representing the Resource ID for a User Onenote Notebook Section Group Section
type UserOnenoteNotebookSectionGroupSectionId struct {
	UserId           string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewUserOnenoteNotebookSectionGroupSectionID returns a new UserOnenoteNotebookSectionGroupSectionId struct
func NewUserOnenoteNotebookSectionGroupSectionID(userId string, notebookId string, sectionGroupId string, onenoteSectionId string) UserOnenoteNotebookSectionGroupSectionId {
	return UserOnenoteNotebookSectionGroupSectionId{
		UserId:           userId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseUserOnenoteNotebookSectionGroupSectionID parses 'input' into a UserOnenoteNotebookSectionGroupSectionId
func ParseUserOnenoteNotebookSectionGroupSectionID(input string) (*UserOnenoteNotebookSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionGroupSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionGroupSectionId{}

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

	return &id, nil
}

// ParseUserOnenoteNotebookSectionGroupSectionIDInsensitively parses 'input' case-insensitively into a UserOnenoteNotebookSectionGroupSectionId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteNotebookSectionGroupSectionIDInsensitively(input string) (*UserOnenoteNotebookSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionGroupSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionGroupSectionId{}

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

	return &id, nil
}

// ValidateUserOnenoteNotebookSectionGroupSectionID checks that 'input' can be parsed as a User Onenote Notebook Section Group Section ID
func ValidateUserOnenoteNotebookSectionGroupSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteNotebookSectionGroupSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Notebook Section Group Section ID
func (id UserOnenoteNotebookSectionGroupSectionId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Notebook Section Group Section ID
func (id UserOnenoteNotebookSectionGroupSectionId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this User Onenote Notebook Section Group Section ID
func (id UserOnenoteNotebookSectionGroupSectionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("User Onenote Notebook Section Group Section (%s)", strings.Join(components, "\n"))
}
