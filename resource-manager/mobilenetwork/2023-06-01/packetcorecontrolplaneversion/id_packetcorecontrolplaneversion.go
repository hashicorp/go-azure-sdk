package packetcorecontrolplaneversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PacketCoreControlPlaneVersionId{})
}

var _ resourceids.ResourceId = &PacketCoreControlPlaneVersionId{}

// PacketCoreControlPlaneVersionId is a struct representing the Resource ID for a Packet Core Control Plane Version
type PacketCoreControlPlaneVersionId struct {
	PacketCoreControlPlaneVersionName string
}

// NewPacketCoreControlPlaneVersionID returns a new PacketCoreControlPlaneVersionId struct
func NewPacketCoreControlPlaneVersionID(packetCoreControlPlaneVersionName string) PacketCoreControlPlaneVersionId {
	return PacketCoreControlPlaneVersionId{
		PacketCoreControlPlaneVersionName: packetCoreControlPlaneVersionName,
	}
}

// ParsePacketCoreControlPlaneVersionID parses 'input' into a PacketCoreControlPlaneVersionId
func ParsePacketCoreControlPlaneVersionID(input string) (*PacketCoreControlPlaneVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PacketCoreControlPlaneVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PacketCoreControlPlaneVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePacketCoreControlPlaneVersionIDInsensitively parses 'input' case-insensitively into a PacketCoreControlPlaneVersionId
// note: this method should only be used for API response data and not user input
func ParsePacketCoreControlPlaneVersionIDInsensitively(input string) (*PacketCoreControlPlaneVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PacketCoreControlPlaneVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PacketCoreControlPlaneVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PacketCoreControlPlaneVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PacketCoreControlPlaneVersionName, ok = input.Parsed["packetCoreControlPlaneVersionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "packetCoreControlPlaneVersionName", input)
	}

	return nil
}

// ValidatePacketCoreControlPlaneVersionID checks that 'input' can be parsed as a Packet Core Control Plane Version ID
func ValidatePacketCoreControlPlaneVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePacketCoreControlPlaneVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Packet Core Control Plane Version ID
func (id PacketCoreControlPlaneVersionId) ID() string {
	fmtString := "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/%s"
	return fmt.Sprintf(fmtString, id.PacketCoreControlPlaneVersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Packet Core Control Plane Version ID
func (id PacketCoreControlPlaneVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMobileNetwork", "Microsoft.MobileNetwork", "Microsoft.MobileNetwork"),
		resourceids.StaticSegment("staticPacketCoreControlPlaneVersions", "packetCoreControlPlaneVersions", "packetCoreControlPlaneVersions"),
		resourceids.UserSpecifiedSegment("packetCoreControlPlaneVersionName", "packetCoreControlPlaneVersionValue"),
	}
}

// String returns a human-readable description of this Packet Core Control Plane Version ID
func (id PacketCoreControlPlaneVersionId) String() string {
	components := []string{
		fmt.Sprintf("Packet Core Control Plane Version Name: %q", id.PacketCoreControlPlaneVersionName),
	}
	return fmt.Sprintf("Packet Core Control Plane Version (%s)", strings.Join(components, "\n"))
}
