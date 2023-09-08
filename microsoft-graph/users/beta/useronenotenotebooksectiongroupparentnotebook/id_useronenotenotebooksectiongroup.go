package useronenotenotebooksectiongroupparentnotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteNotebookSectionGroupId{}

// UserOnenoteNotebookSectionGroupId is a struct representing the Resource ID for a User Onenote Notebook Section Group
type UserOnenoteNotebookSectionGroupId struct {
	UserId         string
	NotebookId     string
	SectionGroupId string
}

// NewUserOnenoteNotebookSectionGroupID returns a new UserOnenoteNotebookSectionGroupId struct
func NewUserOnenoteNotebookSectionGroupID(userId string, notebookId string, sectionGroupId string) UserOnenoteNotebookSectionGroupId {
	return UserOnenoteNotebookSectionGroupId{
		UserId:         userId,
		NotebookId:     notebookId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseUserOnenoteNotebookSectionGroupID parses 'input' into a UserOnenoteNotebookSectionGroupId
func ParseUserOnenoteNotebookSectionGroupID(input string) (*UserOnenoteNotebookSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteNotebookSectionGroupIDInsensitively parses 'input' case-insensitively into a UserOnenoteNotebookSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteNotebookSectionGroupIDInsensitively(input string) (*UserOnenoteNotebookSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteNotebookSectionGroupID checks that 'input' can be parsed as a User Onenote Notebook Section Group ID
func ValidateUserOnenoteNotebookSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteNotebookSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Notebook Section Group ID
func (id UserOnenoteNotebookSectionGroupId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Notebook Section Group ID
func (id UserOnenoteNotebookSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Notebook Section Group ID
func (id UserOnenoteNotebookSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("User Onenote Notebook Section Group (%s)", strings.Join(components, "\n"))
}
