package replicationevents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ReplicationEventId{}

// ReplicationEventId is a struct representing the Resource ID for a Replication Event
type ReplicationEventId struct {
	SubscriptionId       string
	ResourceGroupName    string
	VaultName            string
	ReplicationEventName string
}

// NewReplicationEventID returns a new ReplicationEventId struct
func NewReplicationEventID(subscriptionId string, resourceGroupName string, vaultName string, replicationEventName string) ReplicationEventId {
	return ReplicationEventId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		VaultName:            vaultName,
		ReplicationEventName: replicationEventName,
	}
}

// ParseReplicationEventID parses 'input' into a ReplicationEventId
func ParseReplicationEventID(input string) (*ReplicationEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationEventId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.ReplicationEventName, ok = parsed.Parsed["replicationEventName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationEventName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationEventIDInsensitively parses 'input' case-insensitively into a ReplicationEventId
// note: this method should only be used for API response data and not user input
func ParseReplicationEventIDInsensitively(input string) (*ReplicationEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationEventId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.ReplicationEventName, ok = parsed.Parsed["replicationEventName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationEventName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationEventID checks that 'input' can be parsed as a Replication Event ID
func ValidateReplicationEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Event ID
func (id ReplicationEventId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationEvents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationEventName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Event ID
func (id ReplicationEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticReplicationEvents", "replicationEvents", "replicationEvents"),
		resourceids.UserSpecifiedSegment("replicationEventName", "replicationEventValue"),
	}
}

// String returns a human-readable description of this Replication Event ID
func (id ReplicationEventId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Event Name: %q", id.ReplicationEventName),
	}
	return fmt.Sprintf("Replication Event (%s)", strings.Join(components, "\n"))
}
