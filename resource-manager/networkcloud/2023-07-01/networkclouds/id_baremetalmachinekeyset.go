package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BareMetalMachineKeySetId{})
}

var _ resourceids.ResourceId = &BareMetalMachineKeySetId{}

// BareMetalMachineKeySetId is a struct representing the Resource ID for a Bare Metal Machine Key Set
type BareMetalMachineKeySetId struct {
	SubscriptionId             string
	ResourceGroupName          string
	ClusterName                string
	BareMetalMachineKeySetName string
}

// NewBareMetalMachineKeySetID returns a new BareMetalMachineKeySetId struct
func NewBareMetalMachineKeySetID(subscriptionId string, resourceGroupName string, clusterName string, bareMetalMachineKeySetName string) BareMetalMachineKeySetId {
	return BareMetalMachineKeySetId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		ClusterName:                clusterName,
		BareMetalMachineKeySetName: bareMetalMachineKeySetName,
	}
}

// ParseBareMetalMachineKeySetID parses 'input' into a BareMetalMachineKeySetId
func ParseBareMetalMachineKeySetID(input string) (*BareMetalMachineKeySetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BareMetalMachineKeySetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BareMetalMachineKeySetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBareMetalMachineKeySetIDInsensitively parses 'input' case-insensitively into a BareMetalMachineKeySetId
// note: this method should only be used for API response data and not user input
func ParseBareMetalMachineKeySetIDInsensitively(input string) (*BareMetalMachineKeySetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BareMetalMachineKeySetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BareMetalMachineKeySetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BareMetalMachineKeySetId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ClusterName, ok = input.Parsed["clusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterName", input)
	}

	if id.BareMetalMachineKeySetName, ok = input.Parsed["bareMetalMachineKeySetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bareMetalMachineKeySetName", input)
	}

	return nil
}

// ValidateBareMetalMachineKeySetID checks that 'input' can be parsed as a Bare Metal Machine Key Set ID
func ValidateBareMetalMachineKeySetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBareMetalMachineKeySetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Bare Metal Machine Key Set ID
func (id BareMetalMachineKeySetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/clusters/%s/bareMetalMachineKeySets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.BareMetalMachineKeySetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Bare Metal Machine Key Set ID
func (id BareMetalMachineKeySetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticBareMetalMachineKeySets", "bareMetalMachineKeySets", "bareMetalMachineKeySets"),
		resourceids.UserSpecifiedSegment("bareMetalMachineKeySetName", "bareMetalMachineKeySetValue"),
	}
}

// String returns a human-readable description of this Bare Metal Machine Key Set ID
func (id BareMetalMachineKeySetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Bare Metal Machine Key Set Name: %q", id.BareMetalMachineKeySetName),
	}
	return fmt.Sprintf("Bare Metal Machine Key Set (%s)", strings.Join(components, "\n"))
}
