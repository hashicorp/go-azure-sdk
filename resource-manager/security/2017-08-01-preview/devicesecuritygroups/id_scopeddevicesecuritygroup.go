package devicesecuritygroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ScopedDeviceSecurityGroupId{}

// ScopedDeviceSecurityGroupId is a struct representing the Resource ID for a Scoped Device Security Group
type ScopedDeviceSecurityGroupId struct {
	ResourceId              string
	DeviceSecurityGroupName string
}

// NewScopedDeviceSecurityGroupID returns a new ScopedDeviceSecurityGroupId struct
func NewScopedDeviceSecurityGroupID(resourceId string, deviceSecurityGroupName string) ScopedDeviceSecurityGroupId {
	return ScopedDeviceSecurityGroupId{
		ResourceId:              resourceId,
		DeviceSecurityGroupName: deviceSecurityGroupName,
	}
}

// ParseScopedDeviceSecurityGroupID parses 'input' into a ScopedDeviceSecurityGroupId
func ParseScopedDeviceSecurityGroupID(input string) (*ScopedDeviceSecurityGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedDeviceSecurityGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedDeviceSecurityGroupId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedDeviceSecurityGroupIDInsensitively parses 'input' case-insensitively into a ScopedDeviceSecurityGroupId
// note: this method should only be used for API response data and not user input
func ParseScopedDeviceSecurityGroupIDInsensitively(input string) (*ScopedDeviceSecurityGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedDeviceSecurityGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedDeviceSecurityGroupId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedDeviceSecurityGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ResourceId, ok = input.Parsed["resourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceId", input)
	}

	if id.DeviceSecurityGroupName, ok = input.Parsed["deviceSecurityGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceSecurityGroupName", input)
	}

	return nil
}

// ValidateScopedDeviceSecurityGroupID checks that 'input' can be parsed as a Scoped Device Security Group ID
func ValidateScopedDeviceSecurityGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedDeviceSecurityGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Device Security Group ID
func (id ScopedDeviceSecurityGroupId) ID() string {
	fmtString := "/%s/providers/Microsoft.Security/deviceSecurityGroups/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.ResourceId, "/"), id.DeviceSecurityGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Device Security Group ID
func (id ScopedDeviceSecurityGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("resourceId", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticDeviceSecurityGroups", "deviceSecurityGroups", "deviceSecurityGroups"),
		resourceids.UserSpecifiedSegment("deviceSecurityGroupName", "deviceSecurityGroupValue"),
	}
}

// String returns a human-readable description of this Scoped Device Security Group ID
func (id ScopedDeviceSecurityGroupId) String() string {
	components := []string{
		fmt.Sprintf("Resource: %q", id.ResourceId),
		fmt.Sprintf("Device Security Group Name: %q", id.DeviceSecurityGroupName),
	}
	return fmt.Sprintf("Scoped Device Security Group (%s)", strings.Join(components, "\n"))
}
