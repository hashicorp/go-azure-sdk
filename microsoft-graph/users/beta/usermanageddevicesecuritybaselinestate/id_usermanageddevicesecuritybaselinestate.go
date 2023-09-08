package usermanageddevicesecuritybaselinestate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceSecurityBaselineStateId{}

// UserManagedDeviceSecurityBaselineStateId is a struct representing the Resource ID for a User Managed Device Security Baseline State
type UserManagedDeviceSecurityBaselineStateId struct {
	UserId                  string
	ManagedDeviceId         string
	SecurityBaselineStateId string
}

// NewUserManagedDeviceSecurityBaselineStateID returns a new UserManagedDeviceSecurityBaselineStateId struct
func NewUserManagedDeviceSecurityBaselineStateID(userId string, managedDeviceId string, securityBaselineStateId string) UserManagedDeviceSecurityBaselineStateId {
	return UserManagedDeviceSecurityBaselineStateId{
		UserId:                  userId,
		ManagedDeviceId:         managedDeviceId,
		SecurityBaselineStateId: securityBaselineStateId,
	}
}

// ParseUserManagedDeviceSecurityBaselineStateID parses 'input' into a UserManagedDeviceSecurityBaselineStateId
func ParseUserManagedDeviceSecurityBaselineStateID(input string) (*UserManagedDeviceSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceSecurityBaselineStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.SecurityBaselineStateId, ok = parsed.Parsed["securityBaselineStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedDeviceSecurityBaselineStateIDInsensitively parses 'input' case-insensitively into a UserManagedDeviceSecurityBaselineStateId
// note: this method should only be used for API response data and not user input
func ParseUserManagedDeviceSecurityBaselineStateIDInsensitively(input string) (*UserManagedDeviceSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceSecurityBaselineStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.SecurityBaselineStateId, ok = parsed.Parsed["securityBaselineStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedDeviceSecurityBaselineStateID checks that 'input' can be parsed as a User Managed Device Security Baseline State ID
func ValidateUserManagedDeviceSecurityBaselineStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedDeviceSecurityBaselineStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed Device Security Baseline State ID
func (id UserManagedDeviceSecurityBaselineStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/securityBaselineStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.SecurityBaselineStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed Device Security Baseline State ID
func (id UserManagedDeviceSecurityBaselineStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticSecurityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateIdValue"),
	}
}

// String returns a human-readable description of this User Managed Device Security Baseline State ID
func (id UserManagedDeviceSecurityBaselineStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
	}
	return fmt.Sprintf("User Managed Device Security Baseline State (%s)", strings.Join(components, "\n"))
}
