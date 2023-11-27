package packetcorecontrolplaneversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProviderPacketCoreControlPlaneVersionId{}

// ProviderPacketCoreControlPlaneVersionId is a struct representing the Resource ID for a Provider Packet Core Control Plane Version
type ProviderPacketCoreControlPlaneVersionId struct {
	SubscriptionId                    string
	PacketCoreControlPlaneVersionName string
}

// NewProviderPacketCoreControlPlaneVersionID returns a new ProviderPacketCoreControlPlaneVersionId struct
func NewProviderPacketCoreControlPlaneVersionID(subscriptionId string, packetCoreControlPlaneVersionName string) ProviderPacketCoreControlPlaneVersionId {
	return ProviderPacketCoreControlPlaneVersionId{
		SubscriptionId:                    subscriptionId,
		PacketCoreControlPlaneVersionName: packetCoreControlPlaneVersionName,
	}
}

// ParseProviderPacketCoreControlPlaneVersionID parses 'input' into a ProviderPacketCoreControlPlaneVersionId
func ParseProviderPacketCoreControlPlaneVersionID(input string) (*ProviderPacketCoreControlPlaneVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderPacketCoreControlPlaneVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderPacketCoreControlPlaneVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviderPacketCoreControlPlaneVersionIDInsensitively parses 'input' case-insensitively into a ProviderPacketCoreControlPlaneVersionId
// note: this method should only be used for API response data and not user input
func ParseProviderPacketCoreControlPlaneVersionIDInsensitively(input string) (*ProviderPacketCoreControlPlaneVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderPacketCoreControlPlaneVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderPacketCoreControlPlaneVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProviderPacketCoreControlPlaneVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.PacketCoreControlPlaneVersionName, ok = input.Parsed["packetCoreControlPlaneVersionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "packetCoreControlPlaneVersionName", input)
	}

	return nil
}

// ValidateProviderPacketCoreControlPlaneVersionID checks that 'input' can be parsed as a Provider Packet Core Control Plane Version ID
func ValidateProviderPacketCoreControlPlaneVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderPacketCoreControlPlaneVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Packet Core Control Plane Version ID
func (id ProviderPacketCoreControlPlaneVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.PacketCoreControlPlaneVersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Packet Core Control Plane Version ID
func (id ProviderPacketCoreControlPlaneVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMobileNetwork", "Microsoft.MobileNetwork", "Microsoft.MobileNetwork"),
		resourceids.StaticSegment("staticPacketCoreControlPlaneVersions", "packetCoreControlPlaneVersions", "packetCoreControlPlaneVersions"),
		resourceids.UserSpecifiedSegment("packetCoreControlPlaneVersionName", "packetCoreControlPlaneVersionValue"),
	}
}

// String returns a human-readable description of this Provider Packet Core Control Plane Version ID
func (id ProviderPacketCoreControlPlaneVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Packet Core Control Plane Version Name: %q", id.PacketCoreControlPlaneVersionName),
	}
	return fmt.Sprintf("Provider Packet Core Control Plane Version (%s)", strings.Join(components, "\n"))
}
