package meowneddevice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOwnedDeviceId{}

// MeOwnedDeviceId is a struct representing the Resource ID for a Me Owned Device
type MeOwnedDeviceId struct {
	DirectoryObjectId string
}

// NewMeOwnedDeviceID returns a new MeOwnedDeviceId struct
func NewMeOwnedDeviceID(directoryObjectId string) MeOwnedDeviceId {
	return MeOwnedDeviceId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeOwnedDeviceID parses 'input' into a MeOwnedDeviceId
func ParseMeOwnedDeviceID(input string) (*MeOwnedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOwnedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOwnedDeviceId{}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseMeOwnedDeviceIDInsensitively parses 'input' case-insensitively into a MeOwnedDeviceId
// note: this method should only be used for API response data and not user input
func ParseMeOwnedDeviceIDInsensitively(input string) (*MeOwnedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOwnedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOwnedDeviceId{}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateMeOwnedDeviceID checks that 'input' can be parsed as a Me Owned Device ID
func ValidateMeOwnedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOwnedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Owned Device ID
func (id MeOwnedDeviceId) ID() string {
	fmtString := "/me/ownedDevices/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Owned Device ID
func (id MeOwnedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOwnedDevices", "ownedDevices", "ownedDevices"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Me Owned Device ID
func (id MeOwnedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Owned Device (%s)", strings.Join(components, "\n"))
}
