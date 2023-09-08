package memanageddevicemanageddevicemobileappconfigurationstate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceManagedDeviceMobileAppConfigurationStateId{}

// MeManagedDeviceManagedDeviceMobileAppConfigurationStateId is a struct representing the Resource ID for a Me Managed Device Managed Device Mobile App Configuration State
type MeManagedDeviceManagedDeviceMobileAppConfigurationStateId struct {
	ManagedDeviceId                            string
	ManagedDeviceMobileAppConfigurationStateId string
}

// NewMeManagedDeviceManagedDeviceMobileAppConfigurationStateID returns a new MeManagedDeviceManagedDeviceMobileAppConfigurationStateId struct
func NewMeManagedDeviceManagedDeviceMobileAppConfigurationStateID(managedDeviceId string, managedDeviceMobileAppConfigurationStateId string) MeManagedDeviceManagedDeviceMobileAppConfigurationStateId {
	return MeManagedDeviceManagedDeviceMobileAppConfigurationStateId{
		ManagedDeviceId: managedDeviceId,
		ManagedDeviceMobileAppConfigurationStateId: managedDeviceMobileAppConfigurationStateId,
	}
}

// ParseMeManagedDeviceManagedDeviceMobileAppConfigurationStateID parses 'input' into a MeManagedDeviceManagedDeviceMobileAppConfigurationStateId
func ParseMeManagedDeviceManagedDeviceMobileAppConfigurationStateID(input string) (*MeManagedDeviceManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceManagedDeviceMobileAppConfigurationStateId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.ManagedDeviceMobileAppConfigurationStateId, ok = parsed.Parsed["managedDeviceMobileAppConfigurationStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceMobileAppConfigurationStateId", *parsed)
	}

	return &id, nil
}

// ParseMeManagedDeviceManagedDeviceMobileAppConfigurationStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceManagedDeviceMobileAppConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceManagedDeviceMobileAppConfigurationStateIDInsensitively(input string) (*MeManagedDeviceManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceManagedDeviceMobileAppConfigurationStateId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.ManagedDeviceMobileAppConfigurationStateId, ok = parsed.Parsed["managedDeviceMobileAppConfigurationStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceMobileAppConfigurationStateId", *parsed)
	}

	return &id, nil
}

// ValidateMeManagedDeviceManagedDeviceMobileAppConfigurationStateID checks that 'input' can be parsed as a Me Managed Device Managed Device Mobile App Configuration State ID
func ValidateMeManagedDeviceManagedDeviceMobileAppConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceManagedDeviceMobileAppConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Managed Device Mobile App Configuration State ID
func (id MeManagedDeviceManagedDeviceMobileAppConfigurationStateId) ID() string {
	fmtString := "/me/managedDevices/%s/managedDeviceMobileAppConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.ManagedDeviceMobileAppConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Managed Device Mobile App Configuration State ID
func (id MeManagedDeviceManagedDeviceMobileAppConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticManagedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates"),
		resourceids.UserSpecifiedSegment("managedDeviceMobileAppConfigurationStateId", "managedDeviceMobileAppConfigurationStateIdValue"),
	}
}

// String returns a human-readable description of this Me Managed Device Managed Device Mobile App Configuration State ID
func (id MeManagedDeviceManagedDeviceMobileAppConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Managed Device Mobile App Configuration State: %q", id.ManagedDeviceMobileAppConfigurationStateId),
	}
	return fmt.Sprintf("Me Managed Device Managed Device Mobile App Configuration State (%s)", strings.Join(components, "\n"))
}
