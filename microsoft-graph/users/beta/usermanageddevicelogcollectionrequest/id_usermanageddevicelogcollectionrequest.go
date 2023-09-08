package usermanageddevicelogcollectionrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceLogCollectionRequestId{}

// UserManagedDeviceLogCollectionRequestId is a struct representing the Resource ID for a User Managed Device Log Collection Request
type UserManagedDeviceLogCollectionRequestId struct {
	UserId                        string
	ManagedDeviceId               string
	DeviceLogCollectionResponseId string
}

// NewUserManagedDeviceLogCollectionRequestID returns a new UserManagedDeviceLogCollectionRequestId struct
func NewUserManagedDeviceLogCollectionRequestID(userId string, managedDeviceId string, deviceLogCollectionResponseId string) UserManagedDeviceLogCollectionRequestId {
	return UserManagedDeviceLogCollectionRequestId{
		UserId:                        userId,
		ManagedDeviceId:               managedDeviceId,
		DeviceLogCollectionResponseId: deviceLogCollectionResponseId,
	}
}

// ParseUserManagedDeviceLogCollectionRequestID parses 'input' into a UserManagedDeviceLogCollectionRequestId
func ParseUserManagedDeviceLogCollectionRequestID(input string) (*UserManagedDeviceLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceLogCollectionRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceLogCollectionResponseId, ok = parsed.Parsed["deviceLogCollectionResponseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceLogCollectionResponseId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedDeviceLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a UserManagedDeviceLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseUserManagedDeviceLogCollectionRequestIDInsensitively(input string) (*UserManagedDeviceLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceLogCollectionRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceLogCollectionResponseId, ok = parsed.Parsed["deviceLogCollectionResponseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceLogCollectionResponseId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedDeviceLogCollectionRequestID checks that 'input' can be parsed as a User Managed Device Log Collection Request ID
func ValidateUserManagedDeviceLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedDeviceLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed Device Log Collection Request ID
func (id UserManagedDeviceLogCollectionRequestId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/logCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.DeviceLogCollectionResponseId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed Device Log Collection Request ID
func (id UserManagedDeviceLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticLogCollectionRequests", "logCollectionRequests", "logCollectionRequests"),
		resourceids.UserSpecifiedSegment("deviceLogCollectionResponseId", "deviceLogCollectionResponseIdValue"),
	}
}

// String returns a human-readable description of this User Managed Device Log Collection Request ID
func (id UserManagedDeviceLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Log Collection Response: %q", id.DeviceLogCollectionResponseId),
	}
	return fmt.Sprintf("User Managed Device Log Collection Request (%s)", strings.Join(components, "\n"))
}
