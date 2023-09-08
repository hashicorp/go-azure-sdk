package memanageddevicesecuritybaselinestatesettingstate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceSecurityBaselineStateSettingStateId{}

// MeManagedDeviceSecurityBaselineStateSettingStateId is a struct representing the Resource ID for a Me Managed Device Security Baseline State Setting State
type MeManagedDeviceSecurityBaselineStateSettingStateId struct {
	ManagedDeviceId                string
	SecurityBaselineStateId        string
	SecurityBaselineSettingStateId string
}

// NewMeManagedDeviceSecurityBaselineStateSettingStateID returns a new MeManagedDeviceSecurityBaselineStateSettingStateId struct
func NewMeManagedDeviceSecurityBaselineStateSettingStateID(managedDeviceId string, securityBaselineStateId string, securityBaselineSettingStateId string) MeManagedDeviceSecurityBaselineStateSettingStateId {
	return MeManagedDeviceSecurityBaselineStateSettingStateId{
		ManagedDeviceId:                managedDeviceId,
		SecurityBaselineStateId:        securityBaselineStateId,
		SecurityBaselineSettingStateId: securityBaselineSettingStateId,
	}
}

// ParseMeManagedDeviceSecurityBaselineStateSettingStateID parses 'input' into a MeManagedDeviceSecurityBaselineStateSettingStateId
func ParseMeManagedDeviceSecurityBaselineStateSettingStateID(input string) (*MeManagedDeviceSecurityBaselineStateSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceSecurityBaselineStateSettingStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceSecurityBaselineStateSettingStateId{}

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

// ParseMeManagedDeviceSecurityBaselineStateSettingStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceSecurityBaselineStateSettingStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceSecurityBaselineStateSettingStateIDInsensitively(input string) (*MeManagedDeviceSecurityBaselineStateSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceSecurityBaselineStateSettingStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceSecurityBaselineStateSettingStateId{}

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

// ValidateMeManagedDeviceSecurityBaselineStateSettingStateID checks that 'input' can be parsed as a Me Managed Device Security Baseline State Setting State ID
func ValidateMeManagedDeviceSecurityBaselineStateSettingStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceSecurityBaselineStateSettingStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Security Baseline State Setting State ID
func (id MeManagedDeviceSecurityBaselineStateSettingStateId) ID() string {
	fmtString := "/me/managedDevices/%s/securityBaselineStates/%s/settingStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.SecurityBaselineStateId, id.SecurityBaselineSettingStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Security Baseline State Setting State ID
func (id MeManagedDeviceSecurityBaselineStateSettingStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticSecurityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateIdValue"),
		resourceids.StaticSegment("staticSettingStates", "settingStates", "settingStates"),
		resourceids.UserSpecifiedSegment("securityBaselineSettingStateId", "securityBaselineSettingStateIdValue"),
	}
}

// String returns a human-readable description of this Me Managed Device Security Baseline State Setting State ID
func (id MeManagedDeviceSecurityBaselineStateSettingStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
		fmt.Sprintf("Security Baseline Setting State: %q", id.SecurityBaselineSettingStateId),
	}
	return fmt.Sprintf("Me Managed Device Security Baseline State Setting State (%s)", strings.Join(components, "\n"))
}
