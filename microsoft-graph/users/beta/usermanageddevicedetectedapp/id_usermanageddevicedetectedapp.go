package usermanageddevicedetectedapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceDetectedAppId{}

// UserManagedDeviceDetectedAppId is a struct representing the Resource ID for a User Managed Device Detected App
type UserManagedDeviceDetectedAppId struct {
	UserId          string
	ManagedDeviceId string
	DetectedAppId   string
}

// NewUserManagedDeviceDetectedAppID returns a new UserManagedDeviceDetectedAppId struct
func NewUserManagedDeviceDetectedAppID(userId string, managedDeviceId string, detectedAppId string) UserManagedDeviceDetectedAppId {
	return UserManagedDeviceDetectedAppId{
		UserId:          userId,
		ManagedDeviceId: managedDeviceId,
		DetectedAppId:   detectedAppId,
	}
}

// ParseUserManagedDeviceDetectedAppID parses 'input' into a UserManagedDeviceDetectedAppId
func ParseUserManagedDeviceDetectedAppID(input string) (*UserManagedDeviceDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceDetectedAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceDetectedAppId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DetectedAppId, ok = parsed.Parsed["detectedAppId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedDeviceDetectedAppIDInsensitively parses 'input' case-insensitively into a UserManagedDeviceDetectedAppId
// note: this method should only be used for API response data and not user input
func ParseUserManagedDeviceDetectedAppIDInsensitively(input string) (*UserManagedDeviceDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceDetectedAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceDetectedAppId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DetectedAppId, ok = parsed.Parsed["detectedAppId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedDeviceDetectedAppID checks that 'input' can be parsed as a User Managed Device Detected App ID
func ValidateUserManagedDeviceDetectedAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedDeviceDetectedAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed Device Detected App ID
func (id UserManagedDeviceDetectedAppId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/detectedApps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.DetectedAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed Device Detected App ID
func (id UserManagedDeviceDetectedAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticDetectedApps", "detectedApps", "detectedApps"),
		resourceids.UserSpecifiedSegment("detectedAppId", "detectedAppIdValue"),
	}
}

// String returns a human-readable description of this User Managed Device Detected App ID
func (id UserManagedDeviceDetectedAppId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Detected App: %q", id.DetectedAppId),
	}
	return fmt.Sprintf("User Managed Device Detected App (%s)", strings.Join(components, "\n"))
}
