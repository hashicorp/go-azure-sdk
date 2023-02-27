// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrationrecoverypoints

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = MigrationRecoveryPointId{}

// MigrationRecoveryPointId is a struct representing the Resource ID for a Migration Recovery Point
type MigrationRecoveryPointId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	VaultName                          string
	ReplicationFabricName              string
	ReplicationProtectionContainerName string
	ReplicationMigrationItemName       string
	MigrationRecoveryPointName         string
}

// NewMigrationRecoveryPointID returns a new MigrationRecoveryPointId struct
func NewMigrationRecoveryPointID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationProtectionContainerName string, replicationMigrationItemName string, migrationRecoveryPointName string) MigrationRecoveryPointId {
	return MigrationRecoveryPointId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		VaultName:                          vaultName,
		ReplicationFabricName:              replicationFabricName,
		ReplicationProtectionContainerName: replicationProtectionContainerName,
		ReplicationMigrationItemName:       replicationMigrationItemName,
		MigrationRecoveryPointName:         migrationRecoveryPointName,
	}
}

// ParseMigrationRecoveryPointID parses 'input' into a MigrationRecoveryPointId
func ParseMigrationRecoveryPointID(input string) (*MigrationRecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(MigrationRecoveryPointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MigrationRecoveryPointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.ReplicationFabricName, ok = parsed.Parsed["replicationFabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationFabricName' was not found in the resource id %q", input)
	}

	if id.ReplicationProtectionContainerName, ok = parsed.Parsed["replicationProtectionContainerName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationProtectionContainerName' was not found in the resource id %q", input)
	}

	if id.ReplicationMigrationItemName, ok = parsed.Parsed["replicationMigrationItemName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationMigrationItemName' was not found in the resource id %q", input)
	}

	if id.MigrationRecoveryPointName, ok = parsed.Parsed["migrationRecoveryPointName"]; !ok {
		return nil, fmt.Errorf("the segment 'migrationRecoveryPointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseMigrationRecoveryPointIDInsensitively parses 'input' case-insensitively into a MigrationRecoveryPointId
// note: this method should only be used for API response data and not user input
func ParseMigrationRecoveryPointIDInsensitively(input string) (*MigrationRecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(MigrationRecoveryPointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MigrationRecoveryPointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.ReplicationFabricName, ok = parsed.Parsed["replicationFabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationFabricName' was not found in the resource id %q", input)
	}

	if id.ReplicationProtectionContainerName, ok = parsed.Parsed["replicationProtectionContainerName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationProtectionContainerName' was not found in the resource id %q", input)
	}

	if id.ReplicationMigrationItemName, ok = parsed.Parsed["replicationMigrationItemName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationMigrationItemName' was not found in the resource id %q", input)
	}

	if id.MigrationRecoveryPointName, ok = parsed.Parsed["migrationRecoveryPointName"]; !ok {
		return nil, fmt.Errorf("the segment 'migrationRecoveryPointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateMigrationRecoveryPointID checks that 'input' can be parsed as a Migration Recovery Point ID
func ValidateMigrationRecoveryPointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMigrationRecoveryPointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Migration Recovery Point ID
func (id MigrationRecoveryPointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationProtectionContainers/%s/replicationMigrationItems/%s/migrationRecoveryPoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationProtectionContainerName, id.ReplicationMigrationItemName, id.MigrationRecoveryPointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Migration Recovery Point ID
func (id MigrationRecoveryPointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticReplicationFabrics", "replicationFabrics", "replicationFabrics"),
		resourceids.UserSpecifiedSegment("replicationFabricName", "replicationFabricValue"),
		resourceids.StaticSegment("staticReplicationProtectionContainers", "replicationProtectionContainers", "replicationProtectionContainers"),
		resourceids.UserSpecifiedSegment("replicationProtectionContainerName", "replicationProtectionContainerValue"),
		resourceids.StaticSegment("staticReplicationMigrationItems", "replicationMigrationItems", "replicationMigrationItems"),
		resourceids.UserSpecifiedSegment("replicationMigrationItemName", "replicationMigrationItemValue"),
		resourceids.StaticSegment("staticMigrationRecoveryPoints", "migrationRecoveryPoints", "migrationRecoveryPoints"),
		resourceids.UserSpecifiedSegment("migrationRecoveryPointName", "migrationRecoveryPointValue"),
	}
}

// String returns a human-readable description of this Migration Recovery Point ID
func (id MigrationRecoveryPointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replication Protection Container Name: %q", id.ReplicationProtectionContainerName),
		fmt.Sprintf("Replication Migration Item Name: %q", id.ReplicationMigrationItemName),
		fmt.Sprintf("Migration Recovery Point Name: %q", id.MigrationRecoveryPointName),
	}
	return fmt.Sprintf("Migration Recovery Point (%s)", strings.Join(components, "\n"))
}
