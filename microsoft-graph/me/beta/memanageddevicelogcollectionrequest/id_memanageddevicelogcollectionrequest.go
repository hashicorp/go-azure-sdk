package memanageddevicelogcollectionrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceLogCollectionRequestId{}

// MeManagedDeviceLogCollectionRequestId is a struct representing the Resource ID for a Me Managed Device Log Collection Request
type MeManagedDeviceLogCollectionRequestId struct {
	ManagedDeviceId               string
	DeviceLogCollectionResponseId string
}

// NewMeManagedDeviceLogCollectionRequestID returns a new MeManagedDeviceLogCollectionRequestId struct
func NewMeManagedDeviceLogCollectionRequestID(managedDeviceId string, deviceLogCollectionResponseId string) MeManagedDeviceLogCollectionRequestId {
	return MeManagedDeviceLogCollectionRequestId{
		ManagedDeviceId:               managedDeviceId,
		DeviceLogCollectionResponseId: deviceLogCollectionResponseId,
	}
}

// ParseMeManagedDeviceLogCollectionRequestID parses 'input' into a MeManagedDeviceLogCollectionRequestId
func ParseMeManagedDeviceLogCollectionRequestID(input string) (*MeManagedDeviceLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceLogCollectionRequestId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceLogCollectionResponseId, ok = parsed.Parsed["deviceLogCollectionResponseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceLogCollectionResponseId", *parsed)
	}

	return &id, nil
}

// ParseMeManagedDeviceLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceLogCollectionRequestIDInsensitively(input string) (*MeManagedDeviceLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceLogCollectionRequestId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceLogCollectionResponseId, ok = parsed.Parsed["deviceLogCollectionResponseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceLogCollectionResponseId", *parsed)
	}

	return &id, nil
}

// ValidateMeManagedDeviceLogCollectionRequestID checks that 'input' can be parsed as a Me Managed Device Log Collection Request ID
func ValidateMeManagedDeviceLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Log Collection Request ID
func (id MeManagedDeviceLogCollectionRequestId) ID() string {
	fmtString := "/me/managedDevices/%s/logCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceLogCollectionResponseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Log Collection Request ID
func (id MeManagedDeviceLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticLogCollectionRequests", "logCollectionRequests", "logCollectionRequests"),
		resourceids.UserSpecifiedSegment("deviceLogCollectionResponseId", "deviceLogCollectionResponseIdValue"),
	}
}

// String returns a human-readable description of this Me Managed Device Log Collection Request ID
func (id MeManagedDeviceLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Log Collection Response: %q", id.DeviceLogCollectionResponseId),
	}
	return fmt.Sprintf("Me Managed Device Log Collection Request (%s)", strings.Join(components, "\n"))
}
