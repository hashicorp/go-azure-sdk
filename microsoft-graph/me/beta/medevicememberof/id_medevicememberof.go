package medevicememberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceMemberOfId{}

// MeDeviceMemberOfId is a struct representing the Resource ID for a Me Device Member Of
type MeDeviceMemberOfId struct {
	DeviceId          string
	DirectoryObjectId string
}

// NewMeDeviceMemberOfID returns a new MeDeviceMemberOfId struct
func NewMeDeviceMemberOfID(deviceId string, directoryObjectId string) MeDeviceMemberOfId {
	return MeDeviceMemberOfId{
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeDeviceMemberOfID parses 'input' into a MeDeviceMemberOfId
func ParseMeDeviceMemberOfID(input string) (*MeDeviceMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceMemberOfId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseMeDeviceMemberOfIDInsensitively parses 'input' case-insensitively into a MeDeviceMemberOfId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceMemberOfIDInsensitively(input string) (*MeDeviceMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceMemberOfId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateMeDeviceMemberOfID checks that 'input' can be parsed as a Me Device Member Of ID
func ValidateMeDeviceMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Member Of ID
func (id MeDeviceMemberOfId) ID() string {
	fmtString := "/me/devices/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Member Of ID
func (id MeDeviceMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticMemberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Me Device Member Of ID
func (id MeDeviceMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Device Member Of (%s)", strings.Join(components, "\n"))
}
