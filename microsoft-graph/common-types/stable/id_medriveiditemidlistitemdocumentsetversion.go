package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdListItemDocumentSetVersionId{}

// MeDriveIdItemIdListItemDocumentSetVersionId is a struct representing the Resource ID for a Me Drive Id Item Id List Item Document Set Version
type MeDriveIdItemIdListItemDocumentSetVersionId struct {
	DriveId              string
	DriveItemId          string
	DocumentSetVersionId string
}

// NewMeDriveIdItemIdListItemDocumentSetVersionID returns a new MeDriveIdItemIdListItemDocumentSetVersionId struct
func NewMeDriveIdItemIdListItemDocumentSetVersionID(driveId string, driveItemId string, documentSetVersionId string) MeDriveIdItemIdListItemDocumentSetVersionId {
	return MeDriveIdItemIdListItemDocumentSetVersionId{
		DriveId:              driveId,
		DriveItemId:          driveItemId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseMeDriveIdItemIdListItemDocumentSetVersionID parses 'input' into a MeDriveIdItemIdListItemDocumentSetVersionId
func ParseMeDriveIdItemIdListItemDocumentSetVersionID(input string) (*MeDriveIdItemIdListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdListItemDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdListItemDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdListItemDocumentSetVersionIDInsensitively(input string) (*MeDriveIdItemIdListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdListItemDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.DocumentSetVersionId, ok = input.Parsed["documentSetVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentSetVersionId", input)
	}

	return nil
}

// ValidateMeDriveIdItemIdListItemDocumentSetVersionID checks that 'input' can be parsed as a Me Drive Id Item Id List Item Document Set Version ID
func ValidateMeDriveIdItemIdListItemDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdListItemDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id List Item Document Set Version ID
func (id MeDriveIdItemIdListItemDocumentSetVersionId) ID() string {
	fmtString := "/me/drives/%s/items/%s/listItem/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id List Item Document Set Version ID
func (id MeDriveIdItemIdListItemDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id List Item Document Set Version ID
func (id MeDriveIdItemIdListItemDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("Me Drive Id Item Id List Item Document Set Version (%s)", strings.Join(components, "\n"))
}
