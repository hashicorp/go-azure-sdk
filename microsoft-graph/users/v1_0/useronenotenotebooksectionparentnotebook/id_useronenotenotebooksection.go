package useronenotenotebooksectionparentnotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteNotebookSectionId{}

// UserOnenoteNotebookSectionId is a struct representing the Resource ID for a User Onenote Notebook Section
type UserOnenoteNotebookSectionId struct {
	UserId           string
	NotebookId       string
	OnenoteSectionId string
}

// NewUserOnenoteNotebookSectionID returns a new UserOnenoteNotebookSectionId struct
func NewUserOnenoteNotebookSectionID(userId string, notebookId string, onenoteSectionId string) UserOnenoteNotebookSectionId {
	return UserOnenoteNotebookSectionId{
		UserId:           userId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseUserOnenoteNotebookSectionID parses 'input' into a UserOnenoteNotebookSectionId
func ParseUserOnenoteNotebookSectionID(input string) (*UserOnenoteNotebookSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteNotebookSectionIDInsensitively parses 'input' case-insensitively into a UserOnenoteNotebookSectionId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteNotebookSectionIDInsensitively(input string) (*UserOnenoteNotebookSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteNotebookSectionID checks that 'input' can be parsed as a User Onenote Notebook Section ID
func ValidateUserOnenoteNotebookSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteNotebookSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Notebook Section ID
func (id UserOnenoteNotebookSectionId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Notebook Section ID
func (id UserOnenoteNotebookSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Notebook Section ID
func (id UserOnenoteNotebookSectionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("User Onenote Notebook Section (%s)", strings.Join(components, "\n"))
}
