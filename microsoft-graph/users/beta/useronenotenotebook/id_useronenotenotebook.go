package useronenotenotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteNotebookId{}

// UserOnenoteNotebookId is a struct representing the Resource ID for a User Onenote Notebook
type UserOnenoteNotebookId struct {
	UserId     string
	NotebookId string
}

// NewUserOnenoteNotebookID returns a new UserOnenoteNotebookId struct
func NewUserOnenoteNotebookID(userId string, notebookId string) UserOnenoteNotebookId {
	return UserOnenoteNotebookId{
		UserId:     userId,
		NotebookId: notebookId,
	}
}

// ParseUserOnenoteNotebookID parses 'input' into a UserOnenoteNotebookId
func ParseUserOnenoteNotebookID(input string) (*UserOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteNotebookIDInsensitively parses 'input' case-insensitively into a UserOnenoteNotebookId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteNotebookIDInsensitively(input string) (*UserOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteNotebookId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteNotebookId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteNotebookID checks that 'input' can be parsed as a User Onenote Notebook ID
func ValidateUserOnenoteNotebookID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteNotebookID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Notebook ID
func (id UserOnenoteNotebookId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Notebook ID
func (id UserOnenoteNotebookId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Notebook ID
func (id UserOnenoteNotebookId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
	}
	return fmt.Sprintf("User Onenote Notebook (%s)", strings.Join(components, "\n"))
}
