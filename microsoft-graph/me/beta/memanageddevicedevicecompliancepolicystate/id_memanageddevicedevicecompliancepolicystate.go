package memanageddevicedevicecompliancepolicystate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceDeviceCompliancePolicyStateId{}

// MeManagedDeviceDeviceCompliancePolicyStateId is a struct representing the Resource ID for a Me Managed Device Device Compliance Policy State
type MeManagedDeviceDeviceCompliancePolicyStateId struct {
	ManagedDeviceId               string
	DeviceCompliancePolicyStateId string
}

// NewMeManagedDeviceDeviceCompliancePolicyStateID returns a new MeManagedDeviceDeviceCompliancePolicyStateId struct
func NewMeManagedDeviceDeviceCompliancePolicyStateID(managedDeviceId string, deviceCompliancePolicyStateId string) MeManagedDeviceDeviceCompliancePolicyStateId {
	return MeManagedDeviceDeviceCompliancePolicyStateId{
		ManagedDeviceId:               managedDeviceId,
		DeviceCompliancePolicyStateId: deviceCompliancePolicyStateId,
	}
}

// ParseMeManagedDeviceDeviceCompliancePolicyStateID parses 'input' into a MeManagedDeviceDeviceCompliancePolicyStateId
func ParseMeManagedDeviceDeviceCompliancePolicyStateID(input string) (*MeManagedDeviceDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceDeviceCompliancePolicyStateId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceCompliancePolicyStateId, ok = parsed.Parsed["deviceCompliancePolicyStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyStateId", *parsed)
	}

	return &id, nil
}

// ParseMeManagedDeviceDeviceCompliancePolicyStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceDeviceCompliancePolicyStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceDeviceCompliancePolicyStateIDInsensitively(input string) (*MeManagedDeviceDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceDeviceCompliancePolicyStateId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceCompliancePolicyStateId, ok = parsed.Parsed["deviceCompliancePolicyStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyStateId", *parsed)
	}

	return &id, nil
}

// ValidateMeManagedDeviceDeviceCompliancePolicyStateID checks that 'input' can be parsed as a Me Managed Device Device Compliance Policy State ID
func ValidateMeManagedDeviceDeviceCompliancePolicyStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceDeviceCompliancePolicyStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Device Compliance Policy State ID
func (id MeManagedDeviceDeviceCompliancePolicyStateId) ID() string {
	fmtString := "/me/managedDevices/%s/deviceCompliancePolicyStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceCompliancePolicyStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Device Compliance Policy State ID
func (id MeManagedDeviceDeviceCompliancePolicyStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticDeviceCompliancePolicyStates", "deviceCompliancePolicyStates", "deviceCompliancePolicyStates"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyStateId", "deviceCompliancePolicyStateIdValue"),
	}
}

// String returns a human-readable description of this Me Managed Device Device Compliance Policy State ID
func (id MeManagedDeviceDeviceCompliancePolicyStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Compliance Policy State: %q", id.DeviceCompliancePolicyStateId),
	}
	return fmt.Sprintf("Me Managed Device Device Compliance Policy State (%s)", strings.Join(components, "\n"))
}
