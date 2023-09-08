package userdeviceenrollmentconfigurationassignment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceEnrollmentConfigurationId{}

// UserDeviceEnrollmentConfigurationId is a struct representing the Resource ID for a User Device Enrollment Configuration
type UserDeviceEnrollmentConfigurationId struct {
	UserId                          string
	DeviceEnrollmentConfigurationId string
}

// NewUserDeviceEnrollmentConfigurationID returns a new UserDeviceEnrollmentConfigurationId struct
func NewUserDeviceEnrollmentConfigurationID(userId string, deviceEnrollmentConfigurationId string) UserDeviceEnrollmentConfigurationId {
	return UserDeviceEnrollmentConfigurationId{
		UserId:                          userId,
		DeviceEnrollmentConfigurationId: deviceEnrollmentConfigurationId,
	}
}

// ParseUserDeviceEnrollmentConfigurationID parses 'input' into a UserDeviceEnrollmentConfigurationId
func ParseUserDeviceEnrollmentConfigurationID(input string) (*UserDeviceEnrollmentConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceEnrollmentConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceEnrollmentConfigurationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceEnrollmentConfigurationId, ok = parsed.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", *parsed)
	}

	return &id, nil
}

// ParseUserDeviceEnrollmentConfigurationIDInsensitively parses 'input' case-insensitively into a UserDeviceEnrollmentConfigurationId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceEnrollmentConfigurationIDInsensitively(input string) (*UserDeviceEnrollmentConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceEnrollmentConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceEnrollmentConfigurationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceEnrollmentConfigurationId, ok = parsed.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", *parsed)
	}

	return &id, nil
}

// ValidateUserDeviceEnrollmentConfigurationID checks that 'input' can be parsed as a User Device Enrollment Configuration ID
func ValidateUserDeviceEnrollmentConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceEnrollmentConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device Enrollment Configuration ID
func (id UserDeviceEnrollmentConfigurationId) ID() string {
	fmtString := "/users/%s/deviceEnrollmentConfigurations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceEnrollmentConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device Enrollment Configuration ID
func (id UserDeviceEnrollmentConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDeviceEnrollmentConfigurations", "deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations"),
		resourceids.UserSpecifiedSegment("deviceEnrollmentConfigurationId", "deviceEnrollmentConfigurationIdValue"),
	}
}

// String returns a human-readable description of this User Device Enrollment Configuration ID
func (id UserDeviceEnrollmentConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device Enrollment Configuration: %q", id.DeviceEnrollmentConfigurationId),
	}
	return fmt.Sprintf("User Device Enrollment Configuration (%s)", strings.Join(components, "\n"))
}
