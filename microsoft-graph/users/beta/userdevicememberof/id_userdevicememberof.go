package userdevicememberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceMemberOfId{}

// UserDeviceMemberOfId is a struct representing the Resource ID for a User Device Member Of
type UserDeviceMemberOfId struct {
	UserId            string
	DeviceId          string
	DirectoryObjectId string
}

// NewUserDeviceMemberOfID returns a new UserDeviceMemberOfId struct
func NewUserDeviceMemberOfID(userId string, deviceId string, directoryObjectId string) UserDeviceMemberOfId {
	return UserDeviceMemberOfId{
		UserId:            userId,
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserDeviceMemberOfID parses 'input' into a UserDeviceMemberOfId
func ParseUserDeviceMemberOfID(input string) (*UserDeviceMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceMemberOfId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserDeviceMemberOfIDInsensitively parses 'input' case-insensitively into a UserDeviceMemberOfId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceMemberOfIDInsensitively(input string) (*UserDeviceMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceMemberOfId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserDeviceMemberOfID checks that 'input' can be parsed as a User Device Member Of ID
func ValidateUserDeviceMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device Member Of ID
func (id UserDeviceMemberOfId) ID() string {
	fmtString := "/users/%s/devices/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device Member Of ID
func (id UserDeviceMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticMemberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Device Member Of ID
func (id UserDeviceMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Device Member Of (%s)", strings.Join(components, "\n"))
}
