package extendedueinformation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&UeId{})
}

var _ resourceids.ResourceId = &UeId{}

// UeId is a struct representing the Resource ID for a Ue
type UeId struct {
	SubscriptionId             string
	ResourceGroupName          string
	PacketCoreControlPlaneName string
	UeId                       string
}

// NewUeID returns a new UeId struct
func NewUeID(subscriptionId string, resourceGroupName string, packetCoreControlPlaneName string, ueId string) UeId {
	return UeId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		PacketCoreControlPlaneName: packetCoreControlPlaneName,
		UeId:                       ueId,
	}
}

// ParseUeID parses 'input' into a UeId
func ParseUeID(input string) (*UeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUeIDInsensitively parses 'input' case-insensitively into a UeId
// note: this method should only be used for API response data and not user input
func ParseUeIDInsensitively(input string) (*UeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PacketCoreControlPlaneName, ok = input.Parsed["packetCoreControlPlaneName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "packetCoreControlPlaneName", input)
	}

	if id.UeId, ok = input.Parsed["ueId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "ueId", input)
	}

	return nil
}

// ValidateUeID checks that 'input' can be parsed as a Ue ID
func ValidateUeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Ue ID
func (id UeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/%s/ues/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PacketCoreControlPlaneName, id.UeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Ue ID
func (id UeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMobileNetwork", "Microsoft.MobileNetwork", "Microsoft.MobileNetwork"),
		resourceids.StaticSegment("staticPacketCoreControlPlanes", "packetCoreControlPlanes", "packetCoreControlPlanes"),
		resourceids.UserSpecifiedSegment("packetCoreControlPlaneName", "packetCoreControlPlaneValue"),
		resourceids.StaticSegment("staticUes", "ues", "ues"),
		resourceids.UserSpecifiedSegment("ueId", "ueIdValue"),
	}
}

// String returns a human-readable description of this Ue ID
func (id UeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Packet Core Control Plane Name: %q", id.PacketCoreControlPlaneName),
		fmt.Sprintf("Ue: %q", id.UeId),
	}
	return fmt.Sprintf("Ue (%s)", strings.Join(components, "\n"))
}
