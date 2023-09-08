package userdeviceenrollmentconfigurationassignment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceEnrollmentConfigurationAssignmentId{}

// UserDeviceEnrollmentConfigurationAssignmentId is a struct representing the Resource ID for a User Device Enrollment Configuration Assignment
type UserDeviceEnrollmentConfigurationAssignmentId struct {
	UserId                              string
	DeviceEnrollmentConfigurationId     string
	EnrollmentConfigurationAssignmentId string
}

// NewUserDeviceEnrollmentConfigurationAssignmentID returns a new UserDeviceEnrollmentConfigurationAssignmentId struct
func NewUserDeviceEnrollmentConfigurationAssignmentID(userId string, deviceEnrollmentConfigurationId string, enrollmentConfigurationAssignmentId string) UserDeviceEnrollmentConfigurationAssignmentId {
	return UserDeviceEnrollmentConfigurationAssignmentId{
		UserId:                              userId,
		DeviceEnrollmentConfigurationId:     deviceEnrollmentConfigurationId,
		EnrollmentConfigurationAssignmentId: enrollmentConfigurationAssignmentId,
	}
}

// ParseUserDeviceEnrollmentConfigurationAssignmentID parses 'input' into a UserDeviceEnrollmentConfigurationAssignmentId
func ParseUserDeviceEnrollmentConfigurationAssignmentID(input string) (*UserDeviceEnrollmentConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceEnrollmentConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceEnrollmentConfigurationAssignmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceEnrollmentConfigurationId, ok = parsed.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", *parsed)
	}

	if id.EnrollmentConfigurationAssignmentId, ok = parsed.Parsed["enrollmentConfigurationAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "enrollmentConfigurationAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseUserDeviceEnrollmentConfigurationAssignmentIDInsensitively parses 'input' case-insensitively into a UserDeviceEnrollmentConfigurationAssignmentId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceEnrollmentConfigurationAssignmentIDInsensitively(input string) (*UserDeviceEnrollmentConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceEnrollmentConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceEnrollmentConfigurationAssignmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceEnrollmentConfigurationId, ok = parsed.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", *parsed)
	}

	if id.EnrollmentConfigurationAssignmentId, ok = parsed.Parsed["enrollmentConfigurationAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "enrollmentConfigurationAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserDeviceEnrollmentConfigurationAssignmentID checks that 'input' can be parsed as a User Device Enrollment Configuration Assignment ID
func ValidateUserDeviceEnrollmentConfigurationAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceEnrollmentConfigurationAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device Enrollment Configuration Assignment ID
func (id UserDeviceEnrollmentConfigurationAssignmentId) ID() string {
	fmtString := "/users/%s/deviceEnrollmentConfigurations/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceEnrollmentConfigurationId, id.EnrollmentConfigurationAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device Enrollment Configuration Assignment ID
func (id UserDeviceEnrollmentConfigurationAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDeviceEnrollmentConfigurations", "deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations"),
		resourceids.UserSpecifiedSegment("deviceEnrollmentConfigurationId", "deviceEnrollmentConfigurationIdValue"),
		resourceids.StaticSegment("staticAssignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("enrollmentConfigurationAssignmentId", "enrollmentConfigurationAssignmentIdValue"),
	}
}

// String returns a human-readable description of this User Device Enrollment Configuration Assignment ID
func (id UserDeviceEnrollmentConfigurationAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device Enrollment Configuration: %q", id.DeviceEnrollmentConfigurationId),
		fmt.Sprintf("Enrollment Configuration Assignment: %q", id.EnrollmentConfigurationAssignmentId),
	}
	return fmt.Sprintf("User Device Enrollment Configuration Assignment (%s)", strings.Join(components, "\n"))
}
