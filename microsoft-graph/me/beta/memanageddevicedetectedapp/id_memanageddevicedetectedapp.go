package memanageddevicedetectedapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceDetectedAppId{}

// MeManagedDeviceDetectedAppId is a struct representing the Resource ID for a Me Managed Device Detected App
type MeManagedDeviceDetectedAppId struct {
	ManagedDeviceId string
	DetectedAppId   string
}

// NewMeManagedDeviceDetectedAppID returns a new MeManagedDeviceDetectedAppId struct
func NewMeManagedDeviceDetectedAppID(managedDeviceId string, detectedAppId string) MeManagedDeviceDetectedAppId {
	return MeManagedDeviceDetectedAppId{
		ManagedDeviceId: managedDeviceId,
		DetectedAppId:   detectedAppId,
	}
}

// ParseMeManagedDeviceDetectedAppID parses 'input' into a MeManagedDeviceDetectedAppId
func ParseMeManagedDeviceDetectedAppID(input string) (*MeManagedDeviceDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceDetectedAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceDetectedAppId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DetectedAppId, ok = parsed.Parsed["detectedAppId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", *parsed)
	}

	return &id, nil
}

// ParseMeManagedDeviceDetectedAppIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceDetectedAppId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceDetectedAppIDInsensitively(input string) (*MeManagedDeviceDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceDetectedAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceDetectedAppId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DetectedAppId, ok = parsed.Parsed["detectedAppId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", *parsed)
	}

	return &id, nil
}

// ValidateMeManagedDeviceDetectedAppID checks that 'input' can be parsed as a Me Managed Device Detected App ID
func ValidateMeManagedDeviceDetectedAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceDetectedAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Detected App ID
func (id MeManagedDeviceDetectedAppId) ID() string {
	fmtString := "/me/managedDevices/%s/detectedApps/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DetectedAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Detected App ID
func (id MeManagedDeviceDetectedAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticDetectedApps", "detectedApps", "detectedApps"),
		resourceids.UserSpecifiedSegment("detectedAppId", "detectedAppIdValue"),
	}
}

// String returns a human-readable description of this Me Managed Device Detected App ID
func (id MeManagedDeviceDetectedAppId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Detected App: %q", id.DetectedAppId),
	}
	return fmt.Sprintf("Me Managed Device Detected App (%s)", strings.Join(components, "\n"))
}
