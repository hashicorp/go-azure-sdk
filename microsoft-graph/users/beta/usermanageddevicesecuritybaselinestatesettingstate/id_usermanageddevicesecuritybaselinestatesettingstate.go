package usermanageddevicesecuritybaselinestatesettingstate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceSecurityBaselineStateSettingStateId{}

// UserManagedDeviceSecurityBaselineStateSettingStateId is a struct representing the Resource ID for a User Managed Device Security Baseline State Setting State
type UserManagedDeviceSecurityBaselineStateSettingStateId struct {
	UserId                         string
	ManagedDeviceId                string
	SecurityBaselineStateId        string
	SecurityBaselineSettingStateId string
}

// NewUserManagedDeviceSecurityBaselineStateSettingStateID returns a new UserManagedDeviceSecurityBaselineStateSettingStateId struct
func NewUserManagedDeviceSecurityBaselineStateSettingStateID(userId string, managedDeviceId string, securityBaselineStateId string, securityBaselineSettingStateId string) UserManagedDeviceSecurityBaselineStateSettingStateId {
	return UserManagedDeviceSecurityBaselineStateSettingStateId{
		UserId:                         userId,
		ManagedDeviceId:                managedDeviceId,
		SecurityBaselineStateId:        securityBaselineStateId,
		SecurityBaselineSettingStateId: securityBaselineSettingStateId,
	}
}

// ParseUserManagedDeviceSecurityBaselineStateSettingStateID parses 'input' into a UserManagedDeviceSecurityBaselineStateSettingStateId
func ParseUserManagedDeviceSecurityBaselineStateSettingStateID(input string) (*UserManagedDeviceSecurityBaselineStateSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceSecurityBaselineStateSettingStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceSecurityBaselineStateSettingStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.SecurityBaselineStateId, ok = parsed.Parsed["securityBaselineStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", *parsed)
	}

	if id.SecurityBaselineSettingStateId, ok = parsed.Parsed["securityBaselineSettingStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineSettingStateId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedDeviceSecurityBaselineStateSettingStateIDInsensitively parses 'input' case-insensitively into a UserManagedDeviceSecurityBaselineStateSettingStateId
// note: this method should only be used for API response data and not user input
func ParseUserManagedDeviceSecurityBaselineStateSettingStateIDInsensitively(input string) (*UserManagedDeviceSecurityBaselineStateSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceSecurityBaselineStateSettingStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceSecurityBaselineStateSettingStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.SecurityBaselineStateId, ok = parsed.Parsed["securityBaselineStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", *parsed)
	}

	if id.SecurityBaselineSettingStateId, ok = parsed.Parsed["securityBaselineSettingStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineSettingStateId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedDeviceSecurityBaselineStateSettingStateID checks that 'input' can be parsed as a User Managed Device Security Baseline State Setting State ID
func ValidateUserManagedDeviceSecurityBaselineStateSettingStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedDeviceSecurityBaselineStateSettingStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed Device Security Baseline State Setting State ID
func (id UserManagedDeviceSecurityBaselineStateSettingStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/securityBaselineStates/%s/settingStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.SecurityBaselineStateId, id.SecurityBaselineSettingStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed Device Security Baseline State Setting State ID
func (id UserManagedDeviceSecurityBaselineStateSettingStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticSecurityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateIdValue"),
		resourceids.StaticSegment("staticSettingStates", "settingStates", "settingStates"),
		resourceids.UserSpecifiedSegment("securityBaselineSettingStateId", "securityBaselineSettingStateIdValue"),
	}
}

// String returns a human-readable description of this User Managed Device Security Baseline State Setting State ID
func (id UserManagedDeviceSecurityBaselineStateSettingStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
		fmt.Sprintf("Security Baseline Setting State: %q", id.SecurityBaselineSettingStateId),
	}
	return fmt.Sprintf("User Managed Device Security Baseline State Setting State (%s)", strings.Join(components, "\n"))
}
