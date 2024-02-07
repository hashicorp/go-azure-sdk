package replicationprotectionintents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReplicationProtectionIntentId{}

// ReplicationProtectionIntentId is a struct representing the Resource ID for a Replication Protection Intent
type ReplicationProtectionIntentId struct {
	SubscriptionId                  string
	ResourceGroupName               string
	VaultName                       string
	ReplicationProtectionIntentName string
}

// NewReplicationProtectionIntentID returns a new ReplicationProtectionIntentId struct
func NewReplicationProtectionIntentID(subscriptionId string, resourceGroupName string, vaultName string, replicationProtectionIntentName string) ReplicationProtectionIntentId {
	return ReplicationProtectionIntentId{
		SubscriptionId:                  subscriptionId,
		ResourceGroupName:               resourceGroupName,
		VaultName:                       vaultName,
		ReplicationProtectionIntentName: replicationProtectionIntentName,
	}
}

// ParseReplicationProtectionIntentID parses 'input' into a ReplicationProtectionIntentId
func ParseReplicationProtectionIntentID(input string) (*ReplicationProtectionIntentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationProtectionIntentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationProtectionIntentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReplicationProtectionIntentIDInsensitively parses 'input' case-insensitively into a ReplicationProtectionIntentId
// note: this method should only be used for API response data and not user input
func ParseReplicationProtectionIntentIDInsensitively(input string) (*ReplicationProtectionIntentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationProtectionIntentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationProtectionIntentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicationProtectionIntentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ReplicationProtectionIntentName, ok = input.Parsed["replicationProtectionIntentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationProtectionIntentName", input)
	}

	return nil
}

// ValidateReplicationProtectionIntentID checks that 'input' can be parsed as a Replication Protection Intent ID
func ValidateReplicationProtectionIntentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationProtectionIntentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Protection Intent ID
func (id ReplicationProtectionIntentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationProtectionIntents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationProtectionIntentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Protection Intent ID
func (id ReplicationProtectionIntentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticReplicationProtectionIntents", "replicationProtectionIntents", "replicationProtectionIntents"),
		resourceids.UserSpecifiedSegment("replicationProtectionIntentName", "replicationProtectionIntentValue"),
	}
}

// String returns a human-readable description of this Replication Protection Intent ID
func (id ReplicationProtectionIntentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Protection Intent Name: %q", id.ReplicationProtectionIntentName),
	}
	return fmt.Sprintf("Replication Protection Intent (%s)", strings.Join(components, "\n"))
}
