package usermanageddevicedetectedapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceId{}

// UserManagedDeviceId is a struct representing the Resource ID for a User Managed Device
type UserManagedDeviceId struct {
	UserId          string
	ManagedDeviceId string
}

// NewUserManagedDeviceID returns a new UserManagedDeviceId struct
func NewUserManagedDeviceID(userId string, managedDeviceId string) UserManagedDeviceId {
	return UserManagedDeviceId{
		UserId:          userId,
		ManagedDeviceId: managedDeviceId,
	}
}

// ParseUserManagedDeviceID parses 'input' into a UserManagedDeviceId
func ParseUserManagedDeviceID(input string) (*UserManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedDeviceIDInsensitively parses 'input' case-insensitively into a UserManagedDeviceId
// note: this method should only be used for API response data and not user input
func ParseUserManagedDeviceIDInsensitively(input string) (*UserManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedDeviceID checks that 'input' can be parsed as a User Managed Device ID
func ValidateUserManagedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed Device ID
func (id UserManagedDeviceId) ID() string {
	fmtString := "/users/%s/managedDevices/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed Device ID
func (id UserManagedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
	}
}

// String returns a human-readable description of this User Managed Device ID
func (id UserManagedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
	}
	return fmt.Sprintf("User Managed Device (%s)", strings.Join(components, "\n"))
}
