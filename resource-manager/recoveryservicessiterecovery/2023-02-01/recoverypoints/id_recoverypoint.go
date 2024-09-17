package recoverypoints

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RecoveryPointId{})
}

var _ resourceids.ResourceId = &RecoveryPointId{}

// RecoveryPointId is a struct representing the Resource ID for a Recovery Point
type RecoveryPointId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	VaultName                          string
	ReplicationFabricName              string
	ReplicationProtectionContainerName string
	ReplicationProtectedItemName       string
	RecoveryPointName                  string
}

// NewRecoveryPointID returns a new RecoveryPointId struct
func NewRecoveryPointID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationProtectionContainerName string, replicationProtectedItemName string, recoveryPointName string) RecoveryPointId {
	return RecoveryPointId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		VaultName:                          vaultName,
		ReplicationFabricName:              replicationFabricName,
		ReplicationProtectionContainerName: replicationProtectionContainerName,
		ReplicationProtectedItemName:       replicationProtectedItemName,
		RecoveryPointName:                  recoveryPointName,
	}
}

// ParseRecoveryPointID parses 'input' into a RecoveryPointId
func ParseRecoveryPointID(input string) (*RecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RecoveryPointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RecoveryPointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRecoveryPointIDInsensitively parses 'input' case-insensitively into a RecoveryPointId
// note: this method should only be used for API response data and not user input
func ParseRecoveryPointIDInsensitively(input string) (*RecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RecoveryPointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RecoveryPointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RecoveryPointId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ReplicationProtectedItemName, ok = input.Parsed["replicationProtectedItemName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationProtectedItemName", input)
	}

	if id.RecoveryPointName, ok = input.Parsed["recoveryPointName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "recoveryPointName", input)
	}

	return nil
}

// ValidateRecoveryPointID checks that 'input' can be parsed as a Recovery Point ID
func ValidateRecoveryPointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRecoveryPointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Recovery Point ID
func (id RecoveryPointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationProtectionContainers/%s/replicationProtectedItems/%s/recoveryPoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationProtectionContainerName, id.ReplicationProtectedItemName, id.RecoveryPointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Recovery Point ID
func (id RecoveryPointId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReplicationProtectedItems", "replicationProtectedItems", "replicationProtectedItems"),
		resourceids.UserSpecifiedSegment("replicationProtectedItemName", "replicationProtectedItemValue"),
		resourceids.StaticSegment("staticRecoveryPoints", "recoveryPoints", "recoveryPoints"),
		resourceids.UserSpecifiedSegment("recoveryPointName", "recoveryPointValue"),
	}
}

// String returns a human-readable description of this Recovery Point ID
func (id RecoveryPointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replication Protection Container Name: %q", id.ReplicationProtectionContainerName),
		fmt.Sprintf("Replication Protected Item Name: %q", id.ReplicationProtectedItemName),
		fmt.Sprintf("Recovery Point Name: %q", id.RecoveryPointName),
	}
	return fmt.Sprintf("Recovery Point (%s)", strings.Join(components, "\n"))
}
