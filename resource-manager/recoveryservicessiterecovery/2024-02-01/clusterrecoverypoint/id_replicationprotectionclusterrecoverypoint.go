package clusterrecoverypoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ReplicationProtectionClusterRecoveryPointId{})
}

var _ resourceids.ResourceId = &ReplicationProtectionClusterRecoveryPointId{}

// ReplicationProtectionClusterRecoveryPointId is a struct representing the Resource ID for a Replication Protection Cluster Recovery Point
type ReplicationProtectionClusterRecoveryPointId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	VaultName                          string
	ReplicationFabricName              string
	ReplicationProtectionContainerName string
	ReplicationProtectionClusterName   string
	RecoveryPointName                  string
}

// NewReplicationProtectionClusterRecoveryPointID returns a new ReplicationProtectionClusterRecoveryPointId struct
func NewReplicationProtectionClusterRecoveryPointID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationProtectionContainerName string, replicationProtectionClusterName string, recoveryPointName string) ReplicationProtectionClusterRecoveryPointId {
	return ReplicationProtectionClusterRecoveryPointId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		VaultName:                          vaultName,
		ReplicationFabricName:              replicationFabricName,
		ReplicationProtectionContainerName: replicationProtectionContainerName,
		ReplicationProtectionClusterName:   replicationProtectionClusterName,
		RecoveryPointName:                  recoveryPointName,
	}
}

// ParseReplicationProtectionClusterRecoveryPointID parses 'input' into a ReplicationProtectionClusterRecoveryPointId
func ParseReplicationProtectionClusterRecoveryPointID(input string) (*ReplicationProtectionClusterRecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationProtectionClusterRecoveryPointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationProtectionClusterRecoveryPointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReplicationProtectionClusterRecoveryPointIDInsensitively parses 'input' case-insensitively into a ReplicationProtectionClusterRecoveryPointId
// note: this method should only be used for API response data and not user input
func ParseReplicationProtectionClusterRecoveryPointIDInsensitively(input string) (*ReplicationProtectionClusterRecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationProtectionClusterRecoveryPointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationProtectionClusterRecoveryPointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicationProtectionClusterRecoveryPointId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ReplicationProtectionContainerName, ok = input.Parsed["replicationProtectionContainerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationProtectionContainerName", input)
	}

	if id.ReplicationProtectionClusterName, ok = input.Parsed["replicationProtectionClusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationProtectionClusterName", input)
	}

	if id.RecoveryPointName, ok = input.Parsed["recoveryPointName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "recoveryPointName", input)
	}

	return nil
}

// ValidateReplicationProtectionClusterRecoveryPointID checks that 'input' can be parsed as a Replication Protection Cluster Recovery Point ID
func ValidateReplicationProtectionClusterRecoveryPointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationProtectionClusterRecoveryPointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Protection Cluster Recovery Point ID
func (id ReplicationProtectionClusterRecoveryPointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationProtectionContainers/%s/replicationProtectionClusters/%s/recoveryPoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationProtectionContainerName, id.ReplicationProtectionClusterName, id.RecoveryPointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Protection Cluster Recovery Point ID
func (id ReplicationProtectionClusterRecoveryPointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultName"),
		resourceids.StaticSegment("staticReplicationFabrics", "replicationFabrics", "replicationFabrics"),
		resourceids.UserSpecifiedSegment("replicationFabricName", "replicationFabricName"),
		resourceids.StaticSegment("staticReplicationProtectionContainers", "replicationProtectionContainers", "replicationProtectionContainers"),
		resourceids.UserSpecifiedSegment("replicationProtectionContainerName", "replicationProtectionContainerName"),
		resourceids.StaticSegment("staticReplicationProtectionClusters", "replicationProtectionClusters", "replicationProtectionClusters"),
		resourceids.UserSpecifiedSegment("replicationProtectionClusterName", "replicationProtectionClusterName"),
		resourceids.StaticSegment("staticRecoveryPoints", "recoveryPoints", "recoveryPoints"),
		resourceids.UserSpecifiedSegment("recoveryPointName", "recoveryPointName"),
	}
}

// String returns a human-readable description of this Replication Protection Cluster Recovery Point ID
func (id ReplicationProtectionClusterRecoveryPointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replication Protection Container Name: %q", id.ReplicationProtectionContainerName),
		fmt.Sprintf("Replication Protection Cluster Name: %q", id.ReplicationProtectionClusterName),
		fmt.Sprintf("Recovery Point Name: %q", id.RecoveryPointName),
	}
	return fmt.Sprintf("Replication Protection Cluster Recovery Point (%s)", strings.Join(components, "\n"))
}
