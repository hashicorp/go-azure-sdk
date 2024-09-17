package workloadnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PortMirroringProfileId{})
}

var _ resourceids.ResourceId = &PortMirroringProfileId{}

// PortMirroringProfileId is a struct representing the Resource ID for a Port Mirroring Profile
type PortMirroringProfileId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	PortMirroringId   string
}

// NewPortMirroringProfileID returns a new PortMirroringProfileId struct
func NewPortMirroringProfileID(subscriptionId string, resourceGroupName string, privateCloudName string, portMirroringId string) PortMirroringProfileId {
	return PortMirroringProfileId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		PortMirroringId:   portMirroringId,
	}
}

// ParsePortMirroringProfileID parses 'input' into a PortMirroringProfileId
func ParsePortMirroringProfileID(input string) (*PortMirroringProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PortMirroringProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PortMirroringProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePortMirroringProfileIDInsensitively parses 'input' case-insensitively into a PortMirroringProfileId
// note: this method should only be used for API response data and not user input
func ParsePortMirroringProfileIDInsensitively(input string) (*PortMirroringProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PortMirroringProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PortMirroringProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PortMirroringProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PrivateCloudName, ok = input.Parsed["privateCloudName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateCloudName", input)
	}

	if id.PortMirroringId, ok = input.Parsed["portMirroringId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "portMirroringId", input)
	}

	return nil
}

// ValidatePortMirroringProfileID checks that 'input' can be parsed as a Port Mirroring Profile ID
func ValidatePortMirroringProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePortMirroringProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Port Mirroring Profile ID
func (id PortMirroringProfileId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/portMirroringProfiles/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.PortMirroringId)
}

// Segments returns a slice of Resource ID Segments which comprise this Port Mirroring Profile ID
func (id PortMirroringProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticWorkloadNetworks", "workloadNetworks", "workloadNetworks"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticPortMirroringProfiles", "portMirroringProfiles", "portMirroringProfiles"),
		resourceids.UserSpecifiedSegment("portMirroringId", "portMirroringIdValue"),
	}
}

// String returns a human-readable description of this Port Mirroring Profile ID
func (id PortMirroringProfileId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Port Mirroring: %q", id.PortMirroringId),
	}
	return fmt.Sprintf("Port Mirroring Profile (%s)", strings.Join(components, "\n"))
}
