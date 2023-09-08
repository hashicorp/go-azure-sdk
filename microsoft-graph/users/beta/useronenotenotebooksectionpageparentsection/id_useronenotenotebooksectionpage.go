package useronenotenotebooksectionpageparentsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteNotebookSectionPageId{}

// UserOnenoteNotebookSectionPageId is a struct representing the Resource ID for a User Onenote Notebook Section Page
type UserOnenoteNotebookSectionPageId struct {
	UserId           string
	NotebookId       string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewUserOnenoteNotebookSectionPageID returns a new UserOnenoteNotebookSectionPageId struct
func NewUserOnenoteNotebookSectionPageID(userId string, notebookId string, onenoteSectionId string, onenotePageId string) UserOnenoteNotebookSectionPageId {
	return UserOnenoteNotebookSectionPageId{
		UserId:           userId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseUserOnenoteNotebookSectionPageID parses 'input' into a UserOnenoteNotebookSectionPageId
func ParseUserOnenoteNotebookSectionPageID(input string) (*UserOnenoteNotebookSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionPageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
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

// ParseUserOnenoteNotebookSectionPageIDInsensitively parses 'input' case-insensitively into a UserOnenoteNotebookSectionPageId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteNotebookSectionPageIDInsensitively(input string) (*UserOnenoteNotebookSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionPageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
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

// ValidateUserOnenoteNotebookSectionPageID checks that 'input' can be parsed as a User Onenote Notebook Section Page ID
func ValidateUserOnenoteNotebookSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteNotebookSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Notebook Section Page ID
func (id UserOnenoteNotebookSectionPageId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Notebook Section Page ID
func (id UserOnenoteNotebookSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Notebook Section Page ID
func (id UserOnenoteNotebookSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Onenote Notebook Section Page (%s)", strings.Join(components, "\n"))
}
