package userowneddevice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOwnedDeviceId{}

// UserOwnedDeviceId is a struct representing the Resource ID for a User Owned Device
type UserOwnedDeviceId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserOwnedDeviceID returns a new UserOwnedDeviceId struct
func NewUserOwnedDeviceID(userId string, directoryObjectId string) UserOwnedDeviceId {
	return UserOwnedDeviceId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserOwnedDeviceID parses 'input' into a UserOwnedDeviceId
func ParseUserOwnedDeviceID(input string) (*UserOwnedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOwnedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOwnedDeviceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserOwnedDeviceIDInsensitively parses 'input' case-insensitively into a UserOwnedDeviceId
// note: this method should only be used for API response data and not user input
func ParseUserOwnedDeviceIDInsensitively(input string) (*UserOwnedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOwnedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOwnedDeviceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserOwnedDeviceID checks that 'input' can be parsed as a User Owned Device ID
func ValidateUserOwnedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOwnedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Owned Device ID
func (id UserOwnedDeviceId) ID() string {
	fmtString := "/users/%s/ownedDevices/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Owned Device ID
func (id UserOwnedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOwnedDevices", "ownedDevices", "ownedDevices"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Owned Device ID
func (id UserOwnedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Owned Device (%s)", strings.Join(components, "\n"))
}
