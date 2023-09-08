package useroutlooktaskfolder

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskFolderId{}

// UserOutlookTaskFolderId is a struct representing the Resource ID for a User Outlook Task Folder
type UserOutlookTaskFolderId struct {
	UserId              string
	OutlookTaskFolderId string
}

// NewUserOutlookTaskFolderID returns a new UserOutlookTaskFolderId struct
func NewUserOutlookTaskFolderID(userId string, outlookTaskFolderId string) UserOutlookTaskFolderId {
	return UserOutlookTaskFolderId{
		UserId:              userId,
		OutlookTaskFolderId: outlookTaskFolderId,
	}
}

// ParseUserOutlookTaskFolderID parses 'input' into a UserOutlookTaskFolderId
func ParseUserOutlookTaskFolderID(input string) (*UserOutlookTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	return &id, nil
}

// ParseUserOutlookTaskFolderIDInsensitively parses 'input' case-insensitively into a UserOutlookTaskFolderId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookTaskFolderIDInsensitively(input string) (*UserOutlookTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	return &id, nil
}

// ValidateUserOutlookTaskFolderID checks that 'input' can be parsed as a User Outlook Task Folder ID
func ValidateUserOutlookTaskFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookTaskFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Task Folder ID
func (id UserOutlookTaskFolderId) ID() string {
	fmtString := "/users/%s/outlook/taskFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Task Folder ID
func (id UserOutlookTaskFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Task Folder ID
func (id UserOutlookTaskFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
	}
	return fmt.Sprintf("User Outlook Task Folder (%s)", strings.Join(components, "\n"))
}
