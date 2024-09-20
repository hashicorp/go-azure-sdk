package replicationlogicalnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ReplicationLogicalNetworkId{})
}

var _ resourceids.ResourceId = &ReplicationLogicalNetworkId{}

// ReplicationLogicalNetworkId is a struct representing the Resource ID for a Replication Logical Network
type ReplicationLogicalNetworkId struct {
	SubscriptionId                string
	ResourceGroupName             string
	VaultName                     string
	ReplicationFabricName         string
	ReplicationLogicalNetworkName string
}

// NewReplicationLogicalNetworkID returns a new ReplicationLogicalNetworkId struct
func NewReplicationLogicalNetworkID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationLogicalNetworkName string) ReplicationLogicalNetworkId {
	return ReplicationLogicalNetworkId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		VaultName:                     vaultName,
		ReplicationFabricName:         replicationFabricName,
		ReplicationLogicalNetworkName: replicationLogicalNetworkName,
	}
}

// ParseReplicationLogicalNetworkID parses 'input' into a ReplicationLogicalNetworkId
func ParseReplicationLogicalNetworkID(input string) (*ReplicationLogicalNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationLogicalNetworkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationLogicalNetworkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReplicationLogicalNetworkIDInsensitively parses 'input' case-insensitively into a ReplicationLogicalNetworkId
// note: this method should only be used for API response data and not user input
func ParseReplicationLogicalNetworkIDInsensitively(input string) (*ReplicationLogicalNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationLogicalNetworkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationLogicalNetworkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicationLogicalNetworkId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.VaultName, ok = input.Parsed["vaultName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "vaultName", input)
	}

	if id.ReplicationFabricName, ok = input.Parsed["replicationFabricName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationFabricName", input)
	}

	if id.ReplicationLogicalNetworkName, ok = input.Parsed["replicationLogicalNetworkName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationLogicalNetworkName", input)
	}

	return nil
}

// ValidateReplicationLogicalNetworkID checks that 'input' can be parsed as a Replication Logical Network ID
func ValidateReplicationLogicalNetworkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationLogicalNetworkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Logical Network ID
func (id ReplicationLogicalNetworkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationLogicalNetworks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationLogicalNetworkName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Logical Network ID
func (id ReplicationLogicalNetworkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "resourceName"),
		resourceids.StaticSegment("staticReplicationFabrics", "replicationFabrics", "replicationFabrics"),
		resourceids.UserSpecifiedSegment("replicationFabricName", "fabricName"),
		resourceids.StaticSegment("staticReplicationLogicalNetworks", "replicationLogicalNetworks", "replicationLogicalNetworks"),
		resourceids.UserSpecifiedSegment("replicationLogicalNetworkName", "logicalNetworkName"),
	}
}

// String returns a human-readable description of this Replication Logical Network ID
func (id ReplicationLogicalNetworkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replication Logical Network Name: %q", id.ReplicationLogicalNetworkName),
	}
	return fmt.Sprintf("Replication Logical Network (%s)", strings.Join(components, "\n"))
}
