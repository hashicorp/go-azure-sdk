package administrativeunitextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AdministrativeUnitExtensionId{}

// AdministrativeUnitExtensionId is a struct representing the Resource ID for a Administrative Unit Extension
type AdministrativeUnitExtensionId struct {
	AdministrativeUnitId string
	ExtensionId          string
}

// NewAdministrativeUnitExtensionID returns a new AdministrativeUnitExtensionId struct
func NewAdministrativeUnitExtensionID(administrativeUnitId string, extensionId string) AdministrativeUnitExtensionId {
	return AdministrativeUnitExtensionId{
		AdministrativeUnitId: administrativeUnitId,
		ExtensionId:          extensionId,
	}
}

// ParseAdministrativeUnitExtensionID parses 'input' into a AdministrativeUnitExtensionId
func ParseAdministrativeUnitExtensionID(input string) (*AdministrativeUnitExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdministrativeUnitExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AdministrativeUnitExtensionId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseAdministrativeUnitExtensionIDInsensitively parses 'input' case-insensitively into a AdministrativeUnitExtensionId
// note: this method should only be used for API response data and not user input
func ParseAdministrativeUnitExtensionIDInsensitively(input string) (*AdministrativeUnitExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdministrativeUnitExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AdministrativeUnitExtensionId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateAdministrativeUnitExtensionID checks that 'input' can be parsed as a Administrative Unit Extension ID
func ValidateAdministrativeUnitExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdministrativeUnitExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Administrative Unit Extension ID
func (id AdministrativeUnitExtensionId) ID() string {
	fmtString := "/administrativeUnits/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Administrative Unit Extension ID
func (id AdministrativeUnitExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticAdministrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Administrative Unit Extension ID
func (id AdministrativeUnitExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Administrative Unit Extension (%s)", strings.Join(components, "\n"))
}
