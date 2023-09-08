package medeviceenrollmentconfigurationassignment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceEnrollmentConfigurationAssignmentId{}

// MeDeviceEnrollmentConfigurationAssignmentId is a struct representing the Resource ID for a Me Device Enrollment Configuration Assignment
type MeDeviceEnrollmentConfigurationAssignmentId struct {
	DeviceEnrollmentConfigurationId     string
	EnrollmentConfigurationAssignmentId string
}

// NewMeDeviceEnrollmentConfigurationAssignmentID returns a new MeDeviceEnrollmentConfigurationAssignmentId struct
func NewMeDeviceEnrollmentConfigurationAssignmentID(deviceEnrollmentConfigurationId string, enrollmentConfigurationAssignmentId string) MeDeviceEnrollmentConfigurationAssignmentId {
	return MeDeviceEnrollmentConfigurationAssignmentId{
		DeviceEnrollmentConfigurationId:     deviceEnrollmentConfigurationId,
		EnrollmentConfigurationAssignmentId: enrollmentConfigurationAssignmentId,
	}
}

// ParseMeDeviceEnrollmentConfigurationAssignmentID parses 'input' into a MeDeviceEnrollmentConfigurationAssignmentId
func ParseMeDeviceEnrollmentConfigurationAssignmentID(input string) (*MeDeviceEnrollmentConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceEnrollmentConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceEnrollmentConfigurationAssignmentId{}

	if id.DeviceEnrollmentConfigurationId, ok = parsed.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", *parsed)
	}

	if id.EnrollmentConfigurationAssignmentId, ok = parsed.Parsed["enrollmentConfigurationAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "enrollmentConfigurationAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseMeDeviceEnrollmentConfigurationAssignmentIDInsensitively parses 'input' case-insensitively into a MeDeviceEnrollmentConfigurationAssignmentId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceEnrollmentConfigurationAssignmentIDInsensitively(input string) (*MeDeviceEnrollmentConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceEnrollmentConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceEnrollmentConfigurationAssignmentId{}

	if id.DeviceEnrollmentConfigurationId, ok = parsed.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", *parsed)
	}

	if id.EnrollmentConfigurationAssignmentId, ok = parsed.Parsed["enrollmentConfigurationAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "enrollmentConfigurationAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeDeviceEnrollmentConfigurationAssignmentID checks that 'input' can be parsed as a Me Device Enrollment Configuration Assignment ID
func ValidateMeDeviceEnrollmentConfigurationAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceEnrollmentConfigurationAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Enrollment Configuration Assignment ID
func (id MeDeviceEnrollmentConfigurationAssignmentId) ID() string {
	fmtString := "/me/deviceEnrollmentConfigurations/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceEnrollmentConfigurationId, id.EnrollmentConfigurationAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Enrollment Configuration Assignment ID
func (id MeDeviceEnrollmentConfigurationAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticDeviceEnrollmentConfigurations", "deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations"),
		resourceids.UserSpecifiedSegment("deviceEnrollmentConfigurationId", "deviceEnrollmentConfigurationIdValue"),
		resourceids.StaticSegment("staticAssignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("enrollmentConfigurationAssignmentId", "enrollmentConfigurationAssignmentIdValue"),
	}
}

// String returns a human-readable description of this Me Device Enrollment Configuration Assignment ID
func (id MeDeviceEnrollmentConfigurationAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Enrollment Configuration: %q", id.DeviceEnrollmentConfigurationId),
		fmt.Sprintf("Enrollment Configuration Assignment: %q", id.EnrollmentConfigurationAssignmentId),
	}
	return fmt.Sprintf("Me Device Enrollment Configuration Assignment (%s)", strings.Join(components, "\n"))
}
