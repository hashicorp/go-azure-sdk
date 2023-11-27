package replicationstorageclassifications

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ReplicationStorageClassificationId{}

// ReplicationStorageClassificationId is a struct representing the Resource ID for a Replication Storage Classification
type ReplicationStorageClassificationId struct {
	SubscriptionId                       string
	ResourceGroupName                    string
	VaultName                            string
	ReplicationFabricName                string
	ReplicationStorageClassificationName string
}

// NewReplicationStorageClassificationID returns a new ReplicationStorageClassificationId struct
func NewReplicationStorageClassificationID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationStorageClassificationName string) ReplicationStorageClassificationId {
	return ReplicationStorageClassificationId{
		SubscriptionId:                       subscriptionId,
		ResourceGroupName:                    resourceGroupName,
		VaultName:                            vaultName,
		ReplicationFabricName:                replicationFabricName,
		ReplicationStorageClassificationName: replicationStorageClassificationName,
	}
}

// ParseReplicationStorageClassificationID parses 'input' into a ReplicationStorageClassificationId
func ParseReplicationStorageClassificationID(input string) (*ReplicationStorageClassificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationStorageClassificationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationStorageClassificationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReplicationStorageClassificationIDInsensitively parses 'input' case-insensitively into a ReplicationStorageClassificationId
// note: this method should only be used for API response data and not user input
func ParseReplicationStorageClassificationIDInsensitively(input string) (*ReplicationStorageClassificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationStorageClassificationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationStorageClassificationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicationStorageClassificationId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ReplicationStorageClassificationName, ok = input.Parsed["replicationStorageClassificationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationStorageClassificationName", input)
	}

	return nil
}

// ValidateReplicationStorageClassificationID checks that 'input' can be parsed as a Replication Storage Classification ID
func ValidateReplicationStorageClassificationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationStorageClassificationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Storage Classification ID
func (id ReplicationStorageClassificationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationStorageClassifications/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationStorageClassificationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Storage Classification ID
func (id ReplicationStorageClassificationId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReplicationStorageClassifications", "replicationStorageClassifications", "replicationStorageClassifications"),
		resourceids.UserSpecifiedSegment("replicationStorageClassificationName", "replicationStorageClassificationValue"),
	}
}

// String returns a human-readable description of this Replication Storage Classification ID
func (id ReplicationStorageClassificationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replication Storage Classification Name: %q", id.ReplicationStorageClassificationName),
	}
	return fmt.Sprintf("Replication Storage Classification (%s)", strings.Join(components, "\n"))
}
