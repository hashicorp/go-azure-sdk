package replicationstorageclassificationmappings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ReplicationStorageClassificationMappingId{})
}

var _ resourceids.ResourceId = &ReplicationStorageClassificationMappingId{}

// ReplicationStorageClassificationMappingId is a struct representing the Resource ID for a Replication Storage Classification Mapping
type ReplicationStorageClassificationMappingId struct {
	SubscriptionId                              string
	ResourceGroupName                           string
	VaultName                                   string
	ReplicationFabricName                       string
	ReplicationStorageClassificationName        string
	ReplicationStorageClassificationMappingName string
}

// NewReplicationStorageClassificationMappingID returns a new ReplicationStorageClassificationMappingId struct
func NewReplicationStorageClassificationMappingID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationStorageClassificationName string, replicationStorageClassificationMappingName string) ReplicationStorageClassificationMappingId {
	return ReplicationStorageClassificationMappingId{
		SubscriptionId:                              subscriptionId,
		ResourceGroupName:                           resourceGroupName,
		VaultName:                                   vaultName,
		ReplicationFabricName:                       replicationFabricName,
		ReplicationStorageClassificationName:        replicationStorageClassificationName,
		ReplicationStorageClassificationMappingName: replicationStorageClassificationMappingName,
	}
}

// ParseReplicationStorageClassificationMappingID parses 'input' into a ReplicationStorageClassificationMappingId
func ParseReplicationStorageClassificationMappingID(input string) (*ReplicationStorageClassificationMappingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationStorageClassificationMappingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationStorageClassificationMappingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReplicationStorageClassificationMappingIDInsensitively parses 'input' case-insensitively into a ReplicationStorageClassificationMappingId
// note: this method should only be used for API response data and not user input
func ParseReplicationStorageClassificationMappingIDInsensitively(input string) (*ReplicationStorageClassificationMappingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationStorageClassificationMappingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationStorageClassificationMappingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicationStorageClassificationMappingId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ReplicationStorageClassificationMappingName, ok = input.Parsed["replicationStorageClassificationMappingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationStorageClassificationMappingName", input)
	}

	return nil
}

// ValidateReplicationStorageClassificationMappingID checks that 'input' can be parsed as a Replication Storage Classification Mapping ID
func ValidateReplicationStorageClassificationMappingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationStorageClassificationMappingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Storage Classification Mapping ID
func (id ReplicationStorageClassificationMappingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationStorageClassifications/%s/replicationStorageClassificationMappings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationStorageClassificationName, id.ReplicationStorageClassificationMappingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Storage Classification Mapping ID
func (id ReplicationStorageClassificationMappingId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReplicationStorageClassificationMappings", "replicationStorageClassificationMappings", "replicationStorageClassificationMappings"),
		resourceids.UserSpecifiedSegment("replicationStorageClassificationMappingName", "replicationStorageClassificationMappingValue"),
	}
}

// String returns a human-readable description of this Replication Storage Classification Mapping ID
func (id ReplicationStorageClassificationMappingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replication Storage Classification Name: %q", id.ReplicationStorageClassificationName),
		fmt.Sprintf("Replication Storage Classification Mapping Name: %q", id.ReplicationStorageClassificationMappingName),
	}
	return fmt.Sprintf("Replication Storage Classification Mapping (%s)", strings.Join(components, "\n"))
}
