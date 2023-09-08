package memanageddevicesecuritybaselinestate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceSecurityBaselineStateId{}

// MeManagedDeviceSecurityBaselineStateId is a struct representing the Resource ID for a Me Managed Device Security Baseline State
type MeManagedDeviceSecurityBaselineStateId struct {
	ManagedDeviceId         string
	SecurityBaselineStateId string
}

// NewMeManagedDeviceSecurityBaselineStateID returns a new MeManagedDeviceSecurityBaselineStateId struct
func NewMeManagedDeviceSecurityBaselineStateID(managedDeviceId string, securityBaselineStateId string) MeManagedDeviceSecurityBaselineStateId {
	return MeManagedDeviceSecurityBaselineStateId{
		ManagedDeviceId:         managedDeviceId,
		SecurityBaselineStateId: securityBaselineStateId,
	}
}

// ParseMeManagedDeviceSecurityBaselineStateID parses 'input' into a MeManagedDeviceSecurityBaselineStateId
func ParseMeManagedDeviceSecurityBaselineStateID(input string) (*MeManagedDeviceSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceSecurityBaselineStateId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.SecurityBaselineStateId, ok = parsed.Parsed["securityBaselineStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", *parsed)
	}

	return &id, nil
}

// ParseMeManagedDeviceSecurityBaselineStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceSecurityBaselineStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceSecurityBaselineStateIDInsensitively(input string) (*MeManagedDeviceSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceSecurityBaselineStateId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.SecurityBaselineStateId, ok = parsed.Parsed["securityBaselineStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", *parsed)
	}

	return &id, nil
}

// ValidateMeManagedDeviceSecurityBaselineStateID checks that 'input' can be parsed as a Me Managed Device Security Baseline State ID
func ValidateMeManagedDeviceSecurityBaselineStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceSecurityBaselineStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Security Baseline State ID
func (id MeManagedDeviceSecurityBaselineStateId) ID() string {
	fmtString := "/me/managedDevices/%s/securityBaselineStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.SecurityBaselineStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Security Baseline State ID
func (id MeManagedDeviceSecurityBaselineStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticSecurityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateIdValue"),
	}
}

// String returns a human-readable description of this Me Managed Device Security Baseline State ID
func (id MeManagedDeviceSecurityBaselineStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
	}
	return fmt.Sprintf("Me Managed Device Security Baseline State (%s)", strings.Join(components, "\n"))
}
