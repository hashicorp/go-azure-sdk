package usermanageddevicemanageddevicemobileappconfigurationstate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{}

// UserManagedDeviceManagedDeviceMobileAppConfigurationStateId is a struct representing the Resource ID for a User Managed Device Managed Device Mobile App Configuration State
type UserManagedDeviceManagedDeviceMobileAppConfigurationStateId struct {
	UserId                                     string
	ManagedDeviceId                            string
	ManagedDeviceMobileAppConfigurationStateId string
}

// NewUserManagedDeviceManagedDeviceMobileAppConfigurationStateID returns a new UserManagedDeviceManagedDeviceMobileAppConfigurationStateId struct
func NewUserManagedDeviceManagedDeviceMobileAppConfigurationStateID(userId string, managedDeviceId string, managedDeviceMobileAppConfigurationStateId string) UserManagedDeviceManagedDeviceMobileAppConfigurationStateId {
	return UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{
		UserId:          userId,
		ManagedDeviceId: managedDeviceId,
		ManagedDeviceMobileAppConfigurationStateId: managedDeviceMobileAppConfigurationStateId,
	}
}

// ParseUserManagedDeviceManagedDeviceMobileAppConfigurationStateID parses 'input' into a UserManagedDeviceManagedDeviceMobileAppConfigurationStateId
func ParseUserManagedDeviceManagedDeviceMobileAppConfigurationStateID(input string) (*UserManagedDeviceManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.ManagedDeviceMobileAppConfigurationStateId, ok = parsed.Parsed["managedDeviceMobileAppConfigurationStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceMobileAppConfigurationStateId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedDeviceManagedDeviceMobileAppConfigurationStateIDInsensitively parses 'input' case-insensitively into a UserManagedDeviceManagedDeviceMobileAppConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseUserManagedDeviceManagedDeviceMobileAppConfigurationStateIDInsensitively(input string) (*UserManagedDeviceManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.ManagedDeviceMobileAppConfigurationStateId, ok = parsed.Parsed["managedDeviceMobileAppConfigurationStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceMobileAppConfigurationStateId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedDeviceManagedDeviceMobileAppConfigurationStateID checks that 'input' can be parsed as a User Managed Device Managed Device Mobile App Configuration State ID
func ValidateUserManagedDeviceManagedDeviceMobileAppConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedDeviceManagedDeviceMobileAppConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed Device Managed Device Mobile App Configuration State ID
func (id UserManagedDeviceManagedDeviceMobileAppConfigurationStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/managedDeviceMobileAppConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.ManagedDeviceMobileAppConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed Device Managed Device Mobile App Configuration State ID
func (id UserManagedDeviceManagedDeviceMobileAppConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticManagedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates"),
		resourceids.UserSpecifiedSegment("managedDeviceMobileAppConfigurationStateId", "managedDeviceMobileAppConfigurationStateIdValue"),
	}
}

// String returns a human-readable description of this User Managed Device Managed Device Mobile App Configuration State ID
func (id UserManagedDeviceManagedDeviceMobileAppConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Managed Device Mobile App Configuration State: %q", id.ManagedDeviceMobileAppConfigurationStateId),
	}
	return fmt.Sprintf("User Managed Device Managed Device Mobile App Configuration State (%s)", strings.Join(components, "\n"))
}
