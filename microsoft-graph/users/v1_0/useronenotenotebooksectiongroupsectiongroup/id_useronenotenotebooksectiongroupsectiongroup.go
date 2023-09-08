package useronenotenotebooksectiongroupsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteNotebookSectionGroupSectionGroupId{}

// UserOnenoteNotebookSectionGroupSectionGroupId is a struct representing the Resource ID for a User Onenote Notebook Section Group Section Group
type UserOnenoteNotebookSectionGroupSectionGroupId struct {
	UserId          string
	NotebookId      string
	SectionGroupId  string
	SectionGroupId1 string
}

// NewUserOnenoteNotebookSectionGroupSectionGroupID returns a new UserOnenoteNotebookSectionGroupSectionGroupId struct
func NewUserOnenoteNotebookSectionGroupSectionGroupID(userId string, notebookId string, sectionGroupId string, sectionGroupId1 string) UserOnenoteNotebookSectionGroupSectionGroupId {
	return UserOnenoteNotebookSectionGroupSectionGroupId{
		UserId:          userId,
		NotebookId:      notebookId,
		SectionGroupId:  sectionGroupId,
		SectionGroupId1: sectionGroupId1,
	}
}

// ParseUserOnenoteNotebookSectionGroupSectionGroupID parses 'input' into a UserOnenoteNotebookSectionGroupSectionGroupId
func ParseUserOnenoteNotebookSectionGroupSectionGroupID(input string) (*UserOnenoteNotebookSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionGroupSectionGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteNotebookSectionGroupSectionGroupIDInsensitively parses 'input' case-insensitively into a UserOnenoteNotebookSectionGroupSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteNotebookSectionGroupSectionGroupIDInsensitively(input string) (*UserOnenoteNotebookSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookSectionGroupSectionGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteNotebookSectionGroupSectionGroupID checks that 'input' can be parsed as a User Onenote Notebook Section Group Section Group ID
func ValidateUserOnenoteNotebookSectionGroupSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteNotebookSectionGroupSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Notebook Section Group Section Group ID
func (id UserOnenoteNotebookSectionGroupSectionGroupId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sectionGroups/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.SectionGroupId, id.SectionGroupId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Notebook Section Group Section Group ID
func (id UserOnenoteNotebookSectionGroupSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId1", "sectionGroupId1Value"),
	}
}

// String returns a human-readable description of this User Onenote Notebook Section Group Section Group ID
func (id UserOnenoteNotebookSectionGroupSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Section Group Id 1: %q", id.SectionGroupId1),
	}
	return fmt.Sprintf("User Onenote Notebook Section Group Section Group (%s)", strings.Join(components, "\n"))
}
