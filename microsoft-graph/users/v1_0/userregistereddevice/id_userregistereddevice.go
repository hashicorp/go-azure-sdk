package userregistereddevice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserRegisteredDeviceId{}

// UserRegisteredDeviceId is a struct representing the Resource ID for a User Registered Device
type UserRegisteredDeviceId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserRegisteredDeviceID returns a new UserRegisteredDeviceId struct
func NewUserRegisteredDeviceID(userId string, directoryObjectId string) UserRegisteredDeviceId {
	return UserRegisteredDeviceId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserRegisteredDeviceID parses 'input' into a UserRegisteredDeviceId
func ParseUserRegisteredDeviceID(input string) (*UserRegisteredDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserRegisteredDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserRegisteredDeviceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserRegisteredDeviceIDInsensitively parses 'input' case-insensitively into a UserRegisteredDeviceId
// note: this method should only be used for API response data and not user input
func ParseUserRegisteredDeviceIDInsensitively(input string) (*UserRegisteredDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserRegisteredDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserRegisteredDeviceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserRegisteredDeviceID checks that 'input' can be parsed as a User Registered Device ID
func ValidateUserRegisteredDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserRegisteredDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Registered Device ID
func (id UserRegisteredDeviceId) ID() string {
	fmtString := "/users/%s/registeredDevices/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Registered Device ID
func (id UserRegisteredDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticRegisteredDevices", "registeredDevices", "registeredDevices"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Registered Device ID
func (id UserRegisteredDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Registered Device (%s)", strings.Join(components, "\n"))
}
