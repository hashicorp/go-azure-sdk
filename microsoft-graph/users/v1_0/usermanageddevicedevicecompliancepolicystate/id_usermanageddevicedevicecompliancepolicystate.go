package usermanageddevicedevicecompliancepolicystate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceDeviceCompliancePolicyStateId{}

// UserManagedDeviceDeviceCompliancePolicyStateId is a struct representing the Resource ID for a User Managed Device Device Compliance Policy State
type UserManagedDeviceDeviceCompliancePolicyStateId struct {
	UserId                        string
	ManagedDeviceId               string
	DeviceCompliancePolicyStateId string
}

// NewUserManagedDeviceDeviceCompliancePolicyStateID returns a new UserManagedDeviceDeviceCompliancePolicyStateId struct
func NewUserManagedDeviceDeviceCompliancePolicyStateID(userId string, managedDeviceId string, deviceCompliancePolicyStateId string) UserManagedDeviceDeviceCompliancePolicyStateId {
	return UserManagedDeviceDeviceCompliancePolicyStateId{
		UserId:                        userId,
		ManagedDeviceId:               managedDeviceId,
		DeviceCompliancePolicyStateId: deviceCompliancePolicyStateId,
	}
}

// ParseUserManagedDeviceDeviceCompliancePolicyStateID parses 'input' into a UserManagedDeviceDeviceCompliancePolicyStateId
func ParseUserManagedDeviceDeviceCompliancePolicyStateID(input string) (*UserManagedDeviceDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceDeviceCompliancePolicyStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceCompliancePolicyStateId, ok = parsed.Parsed["deviceCompliancePolicyStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyStateId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedDeviceDeviceCompliancePolicyStateIDInsensitively parses 'input' case-insensitively into a UserManagedDeviceDeviceCompliancePolicyStateId
// note: this method should only be used for API response data and not user input
func ParseUserManagedDeviceDeviceCompliancePolicyStateIDInsensitively(input string) (*UserManagedDeviceDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceDeviceCompliancePolicyStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceCompliancePolicyStateId, ok = parsed.Parsed["deviceCompliancePolicyStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyStateId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedDeviceDeviceCompliancePolicyStateID checks that 'input' can be parsed as a User Managed Device Device Compliance Policy State ID
func ValidateUserManagedDeviceDeviceCompliancePolicyStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedDeviceDeviceCompliancePolicyStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed Device Device Compliance Policy State ID
func (id UserManagedDeviceDeviceCompliancePolicyStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/deviceCompliancePolicyStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.DeviceCompliancePolicyStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed Device Device Compliance Policy State ID
func (id UserManagedDeviceDeviceCompliancePolicyStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticDeviceCompliancePolicyStates", "deviceCompliancePolicyStates", "deviceCompliancePolicyStates"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyStateId", "deviceCompliancePolicyStateIdValue"),
	}
}

// String returns a human-readable description of this User Managed Device Device Compliance Policy State ID
func (id UserManagedDeviceDeviceCompliancePolicyStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Compliance Policy State: %q", id.DeviceCompliancePolicyStateId),
	}
	return fmt.Sprintf("User Managed Device Device Compliance Policy State (%s)", strings.Join(components, "\n"))
}
