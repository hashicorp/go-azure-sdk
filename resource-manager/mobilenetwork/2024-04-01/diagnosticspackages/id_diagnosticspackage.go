package diagnosticspackages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DiagnosticsPackageId{})
}

var _ resourceids.ResourceId = &DiagnosticsPackageId{}

// DiagnosticsPackageId is a struct representing the Resource ID for a Diagnostics Package
type DiagnosticsPackageId struct {
	SubscriptionId             string
	ResourceGroupName          string
	PacketCoreControlPlaneName string
	DiagnosticsPackageName     string
}

// NewDiagnosticsPackageID returns a new DiagnosticsPackageId struct
func NewDiagnosticsPackageID(subscriptionId string, resourceGroupName string, packetCoreControlPlaneName string, diagnosticsPackageName string) DiagnosticsPackageId {
	return DiagnosticsPackageId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		PacketCoreControlPlaneName: packetCoreControlPlaneName,
		DiagnosticsPackageName:     diagnosticsPackageName,
	}
}

// ParseDiagnosticsPackageID parses 'input' into a DiagnosticsPackageId
func ParseDiagnosticsPackageID(input string) (*DiagnosticsPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiagnosticsPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiagnosticsPackageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDiagnosticsPackageIDInsensitively parses 'input' case-insensitively into a DiagnosticsPackageId
// note: this method should only be used for API response data and not user input
func ParseDiagnosticsPackageIDInsensitively(input string) (*DiagnosticsPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiagnosticsPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiagnosticsPackageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DiagnosticsPackageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DiagnosticsPackageName, ok = input.Parsed["diagnosticsPackageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "diagnosticsPackageName", input)
	}

	return nil
}

// ValidateDiagnosticsPackageID checks that 'input' can be parsed as a Diagnostics Package ID
func ValidateDiagnosticsPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiagnosticsPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Diagnostics Package ID
func (id DiagnosticsPackageId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/%s/diagnosticsPackages/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PacketCoreControlPlaneName, id.DiagnosticsPackageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Diagnostics Package ID
func (id DiagnosticsPackageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMobileNetwork", "Microsoft.MobileNetwork", "Microsoft.MobileNetwork"),
		resourceids.StaticSegment("staticPacketCoreControlPlanes", "packetCoreControlPlanes", "packetCoreControlPlanes"),
		resourceids.UserSpecifiedSegment("packetCoreControlPlaneName", "packetCoreControlPlaneValue"),
		resourceids.StaticSegment("staticDiagnosticsPackages", "diagnosticsPackages", "diagnosticsPackages"),
		resourceids.UserSpecifiedSegment("diagnosticsPackageName", "diagnosticsPackageValue"),
	}
}

// String returns a human-readable description of this Diagnostics Package ID
func (id DiagnosticsPackageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Packet Core Control Plane Name: %q", id.PacketCoreControlPlaneName),
		fmt.Sprintf("Diagnostics Package Name: %q", id.DiagnosticsPackageName),
	}
	return fmt.Sprintf("Diagnostics Package (%s)", strings.Join(components, "\n"))
}
