package meoutlooktaskgrouptaskfolder

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOutlookTaskGroupTaskFolderId{}

// MeOutlookTaskGroupTaskFolderId is a struct representing the Resource ID for a Me Outlook Task Group Task Folder
type MeOutlookTaskGroupTaskFolderId struct {
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
}

// NewMeOutlookTaskGroupTaskFolderID returns a new MeOutlookTaskGroupTaskFolderId struct
func NewMeOutlookTaskGroupTaskFolderID(outlookTaskGroupId string, outlookTaskFolderId string) MeOutlookTaskGroupTaskFolderId {
	return MeOutlookTaskGroupTaskFolderId{
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
	}
}

// ParseMeOutlookTaskGroupTaskFolderID parses 'input' into a MeOutlookTaskGroupTaskFolderId
func ParseMeOutlookTaskGroupTaskFolderID(input string) (*MeOutlookTaskGroupTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskGroupTaskFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskGroupTaskFolderId{}

	if id.OutlookTaskGroupId, ok = parsed.Parsed["outlookTaskGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	return &id, nil
}

// ParseMeOutlookTaskGroupTaskFolderIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskGroupTaskFolderId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskGroupTaskFolderIDInsensitively(input string) (*MeOutlookTaskGroupTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskGroupTaskFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskGroupTaskFolderId{}

	if id.OutlookTaskGroupId, ok = parsed.Parsed["outlookTaskGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	return &id, nil
}

// ValidateMeOutlookTaskGroupTaskFolderID checks that 'input' can be parsed as a Me Outlook Task Group Task Folder ID
func ValidateMeOutlookTaskGroupTaskFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskGroupTaskFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Group Task Folder ID
func (id MeOutlookTaskGroupTaskFolderId) ID() string {
	fmtString := "/me/outlook/taskGroups/%s/taskFolders/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskGroupId, id.OutlookTaskFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Group Task Folder ID
func (id MeOutlookTaskGroupTaskFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupIdValue"),
		resourceids.StaticSegment("staticTaskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderIdValue"),
	}
}

// String returns a human-readable description of this Me Outlook Task Group Task Folder ID
func (id MeOutlookTaskGroupTaskFolderId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
	}
	return fmt.Sprintf("Me Outlook Task Group Task Folder (%s)", strings.Join(components, "\n"))
}
