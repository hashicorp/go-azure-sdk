package directoryadministrativeunitextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryAdministrativeUnitExtensionId{}

// DirectoryAdministrativeUnitExtensionId is a struct representing the Resource ID for a Directory Administrative Unit Extension
type DirectoryAdministrativeUnitExtensionId struct {
	AdministrativeUnitId string
	ExtensionId          string
}

// NewDirectoryAdministrativeUnitExtensionID returns a new DirectoryAdministrativeUnitExtensionId struct
func NewDirectoryAdministrativeUnitExtensionID(administrativeUnitId string, extensionId string) DirectoryAdministrativeUnitExtensionId {
	return DirectoryAdministrativeUnitExtensionId{
		AdministrativeUnitId: administrativeUnitId,
		ExtensionId:          extensionId,
	}
}

// ParseDirectoryAdministrativeUnitExtensionID parses 'input' into a DirectoryAdministrativeUnitExtensionId
func ParseDirectoryAdministrativeUnitExtensionID(input string) (*DirectoryAdministrativeUnitExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryAdministrativeUnitExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryAdministrativeUnitExtensionId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryAdministrativeUnitExtensionIDInsensitively parses 'input' case-insensitively into a DirectoryAdministrativeUnitExtensionId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAdministrativeUnitExtensionIDInsensitively(input string) (*DirectoryAdministrativeUnitExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryAdministrativeUnitExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryAdministrativeUnitExtensionId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryAdministrativeUnitExtensionID checks that 'input' can be parsed as a Directory Administrative Unit Extension ID
func ValidateDirectoryAdministrativeUnitExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAdministrativeUnitExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Administrative Unit Extension ID
func (id DirectoryAdministrativeUnitExtensionId) ID() string {
	fmtString := "/directory/administrativeUnits/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Administrative Unit Extension ID
func (id DirectoryAdministrativeUnitExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticAdministrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Directory Administrative Unit Extension ID
func (id DirectoryAdministrativeUnitExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Directory Administrative Unit Extension (%s)", strings.Join(components, "\n"))
}
