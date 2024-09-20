package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootListItemActivityId{}

// MeDriveIdRootListItemActivityId is a struct representing the Resource ID for a Me Drive Id Root List Item Activity
type MeDriveIdRootListItemActivityId struct {
	DriveId           string
	ItemActivityOLDId string
}

// NewMeDriveIdRootListItemActivityID returns a new MeDriveIdRootListItemActivityId struct
func NewMeDriveIdRootListItemActivityID(driveId string, itemActivityOLDId string) MeDriveIdRootListItemActivityId {
	return MeDriveIdRootListItemActivityId{
		DriveId:           driveId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseMeDriveIdRootListItemActivityID parses 'input' into a MeDriveIdRootListItemActivityId
func ParseMeDriveIdRootListItemActivityID(input string) (*MeDriveIdRootListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootListItemActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootListItemActivityIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootListItemActivityId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootListItemActivityIDInsensitively(input string) (*MeDriveIdRootListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootListItemActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootListItemActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateMeDriveIdRootListItemActivityID checks that 'input' can be parsed as a Me Drive Id Root List Item Activity ID
func ValidateMeDriveIdRootListItemActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootListItemActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root List Item Activity ID
func (id MeDriveIdRootListItemActivityId) ID() string {
	fmtString := "/me/drives/%s/root/listItem/activities/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root List Item Activity ID
func (id MeDriveIdRootListItemActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root List Item Activity ID
func (id MeDriveIdRootListItemActivityId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity O L D: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Me Drive Id Root List Item Activity (%s)", strings.Join(components, "\n"))
}
