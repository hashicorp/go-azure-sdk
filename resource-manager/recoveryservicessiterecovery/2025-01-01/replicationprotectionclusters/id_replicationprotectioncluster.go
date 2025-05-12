package replicationprotectionclusters

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ReplicationProtectionClusterId{})
}

var _ resourceids.ResourceId = &ReplicationProtectionClusterId{}

// ReplicationProtectionClusterId is a struct representing the Resource ID for a Replication Protection Cluster
type ReplicationProtectionClusterId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	VaultName                          string
	ReplicationFabricName              string
	ReplicationProtectionContainerName string
	ReplicationProtectionClusterName   string
}

// NewReplicationProtectionClusterID returns a new ReplicationProtectionClusterId struct
func NewReplicationProtectionClusterID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationProtectionContainerName string, replicationProtectionClusterName string) ReplicationProtectionClusterId {
	return ReplicationProtectionClusterId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		VaultName:                          vaultName,
		ReplicationFabricName:              replicationFabricName,
		ReplicationProtectionContainerName: replicationProtectionContainerName,
		ReplicationProtectionClusterName:   replicationProtectionClusterName,
	}
}

// ParseReplicationProtectionClusterID parses 'input' into a ReplicationProtectionClusterId
func ParseReplicationProtectionClusterID(input string) (*ReplicationProtectionClusterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationProtectionClusterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationProtectionClusterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReplicationProtectionClusterIDInsensitively parses 'input' case-insensitively into a ReplicationProtectionClusterId
// note: this method should only be used for API response data and not user input
func ParseReplicationProtectionClusterIDInsensitively(input string) (*ReplicationProtectionClusterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationProtectionClusterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationProtectionClusterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicationProtectionClusterId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateReplicationProtectionClusterID checks that 'input' can be parsed as a Replication Protection Cluster ID
func ValidateReplicationProtectionClusterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationProtectionClusterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Protection Cluster ID
func (id ReplicationProtectionClusterId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationProtectionContainers/%s/replicationProtectionClusters/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationProtectionContainerName, id.ReplicationProtectionClusterName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Protection Cluster ID
func (id ReplicationProtectionClusterId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Replication Protection Cluster ID
func (id ReplicationProtectionClusterId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replication Protection Container Name: %q", id.ReplicationProtectionContainerName),
		fmt.Sprintf("Replication Protection Cluster Name: %q", id.ReplicationProtectionClusterName),
	}
	return fmt.Sprintf("Replication Protection Cluster (%s)", strings.Join(components, "\n"))
}
