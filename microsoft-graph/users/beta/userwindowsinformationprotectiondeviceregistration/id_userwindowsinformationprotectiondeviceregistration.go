package userwindowsinformationprotectiondeviceregistration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserWindowsInformationProtectionDeviceRegistrationId{}

// UserWindowsInformationProtectionDeviceRegistrationId is a struct representing the Resource ID for a User Windows Information Protection Device Registration
type UserWindowsInformationProtectionDeviceRegistrationId struct {
	UserId                                           string
	WindowsInformationProtectionDeviceRegistrationId string
}

// NewUserWindowsInformationProtectionDeviceRegistrationID returns a new UserWindowsInformationProtectionDeviceRegistrationId struct
func NewUserWindowsInformationProtectionDeviceRegistrationID(userId string, windowsInformationProtectionDeviceRegistrationId string) UserWindowsInformationProtectionDeviceRegistrationId {
	return UserWindowsInformationProtectionDeviceRegistrationId{
		UserId: userId,
		WindowsInformationProtectionDeviceRegistrationId: windowsInformationProtectionDeviceRegistrationId,
	}
}

// ParseUserWindowsInformationProtectionDeviceRegistrationID parses 'input' into a UserWindowsInformationProtectionDeviceRegistrationId
func ParseUserWindowsInformationProtectionDeviceRegistrationID(input string) (*UserWindowsInformationProtectionDeviceRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserWindowsInformationProtectionDeviceRegistrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserWindowsInformationProtectionDeviceRegistrationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.WindowsInformationProtectionDeviceRegistrationId, ok = parsed.Parsed["windowsInformationProtectionDeviceRegistrationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "windowsInformationProtectionDeviceRegistrationId", *parsed)
	}

	return &id, nil
}

// ParseUserWindowsInformationProtectionDeviceRegistrationIDInsensitively parses 'input' case-insensitively into a UserWindowsInformationProtectionDeviceRegistrationId
// note: this method should only be used for API response data and not user input
func ParseUserWindowsInformationProtectionDeviceRegistrationIDInsensitively(input string) (*UserWindowsInformationProtectionDeviceRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserWindowsInformationProtectionDeviceRegistrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserWindowsInformationProtectionDeviceRegistrationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.WindowsInformationProtectionDeviceRegistrationId, ok = parsed.Parsed["windowsInformationProtectionDeviceRegistrationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "windowsInformationProtectionDeviceRegistrationId", *parsed)
	}

	return &id, nil
}

// ValidateUserWindowsInformationProtectionDeviceRegistrationID checks that 'input' can be parsed as a User Windows Information Protection Device Registration ID
func ValidateUserWindowsInformationProtectionDeviceRegistrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserWindowsInformationProtectionDeviceRegistrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Windows Information Protection Device Registration ID
func (id UserWindowsInformationProtectionDeviceRegistrationId) ID() string {
	fmtString := "/users/%s/windowsInformationProtectionDeviceRegistrations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WindowsInformationProtectionDeviceRegistrationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Windows Information Protection Device Registration ID
func (id UserWindowsInformationProtectionDeviceRegistrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticWindowsInformationProtectionDeviceRegistrations", "windowsInformationProtectionDeviceRegistrations", "windowsInformationProtectionDeviceRegistrations"),
		resourceids.UserSpecifiedSegment("windowsInformationProtectionDeviceRegistrationId", "windowsInformationProtectionDeviceRegistrationIdValue"),
	}
}

// String returns a human-readable description of this User Windows Information Protection Device Registration ID
func (id UserWindowsInformationProtectionDeviceRegistrationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Windows Information Protection Device Registration: %q", id.WindowsInformationProtectionDeviceRegistrationId),
	}
	return fmt.Sprintf("User Windows Information Protection Device Registration (%s)", strings.Join(components, "\n"))
}
