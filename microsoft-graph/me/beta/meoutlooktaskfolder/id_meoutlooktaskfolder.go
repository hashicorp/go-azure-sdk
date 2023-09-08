package meoutlooktaskfolder

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOutlookTaskFolderId{}

// MeOutlookTaskFolderId is a struct representing the Resource ID for a Me Outlook Task Folder
type MeOutlookTaskFolderId struct {
	OutlookTaskFolderId string
}

// NewMeOutlookTaskFolderID returns a new MeOutlookTaskFolderId struct
func NewMeOutlookTaskFolderID(outlookTaskFolderId string) MeOutlookTaskFolderId {
	return MeOutlookTaskFolderId{
		OutlookTaskFolderId: outlookTaskFolderId,
	}
}

// ParseMeOutlookTaskFolderID parses 'input' into a MeOutlookTaskFolderId
func ParseMeOutlookTaskFolderID(input string) (*MeOutlookTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskFolderId{}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	return &id, nil
}

// ParseMeOutlookTaskFolderIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskFolderId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskFolderIDInsensitively(input string) (*MeOutlookTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskFolderId{}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	return &id, nil
}

// ValidateMeOutlookTaskFolderID checks that 'input' can be parsed as a Me Outlook Task Folder ID
func ValidateMeOutlookTaskFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Folder ID
func (id MeOutlookTaskFolderId) ID() string {
	fmtString := "/me/outlook/taskFolders/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Folder ID
func (id MeOutlookTaskFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderIdValue"),
	}
}

// String returns a human-readable description of this Me Outlook Task Folder ID
func (id MeOutlookTaskFolderId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
	}
	return fmt.Sprintf("Me Outlook Task Folder (%s)", strings.Join(components, "\n"))
}
