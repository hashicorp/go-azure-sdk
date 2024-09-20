package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComplianceCategoryId{}

// DeviceManagementComplianceCategoryId is a struct representing the Resource ID for a Device Management Compliance Category
type DeviceManagementComplianceCategoryId struct {
	DeviceManagementConfigurationCategoryId string
}

// NewDeviceManagementComplianceCategoryID returns a new DeviceManagementComplianceCategoryId struct
func NewDeviceManagementComplianceCategoryID(deviceManagementConfigurationCategoryId string) DeviceManagementComplianceCategoryId {
	return DeviceManagementComplianceCategoryId{
		DeviceManagementConfigurationCategoryId: deviceManagementConfigurationCategoryId,
	}
}

// ParseDeviceManagementComplianceCategoryID parses 'input' into a DeviceManagementComplianceCategoryId
func ParseDeviceManagementComplianceCategoryID(input string) (*DeviceManagementComplianceCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComplianceCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComplianceCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComplianceCategoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementComplianceCategoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComplianceCategoryIDInsensitively(input string) (*DeviceManagementComplianceCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComplianceCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComplianceCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComplianceCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationCategoryId, ok = input.Parsed["deviceManagementConfigurationCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationCategoryId", input)
	}

	return nil
}

// ValidateDeviceManagementComplianceCategoryID checks that 'input' can be parsed as a Device Management Compliance Category ID
func ValidateDeviceManagementComplianceCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComplianceCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Compliance Category ID
func (id DeviceManagementComplianceCategoryId) ID() string {
	fmtString := "/deviceManagement/complianceCategories/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Compliance Category ID
func (id DeviceManagementComplianceCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("complianceCategories", "complianceCategories", "complianceCategories"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationCategoryId", "deviceManagementConfigurationCategoryId"),
	}
}

// String returns a human-readable description of this Device Management Compliance Category ID
func (id DeviceManagementComplianceCategoryId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Category: %q", id.DeviceManagementConfigurationCategoryId),
	}
	return fmt.Sprintf("Device Management Compliance Category (%s)", strings.Join(components, "\n"))
}
