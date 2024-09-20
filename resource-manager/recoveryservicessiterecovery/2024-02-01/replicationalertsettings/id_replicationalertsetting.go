package replicationalertsettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ReplicationAlertSettingId{})
}

var _ resourceids.ResourceId = &ReplicationAlertSettingId{}

// ReplicationAlertSettingId is a struct representing the Resource ID for a Replication Alert Setting
type ReplicationAlertSettingId struct {
	SubscriptionId              string
	ResourceGroupName           string
	VaultName                   string
	ReplicationAlertSettingName string
}

// NewReplicationAlertSettingID returns a new ReplicationAlertSettingId struct
func NewReplicationAlertSettingID(subscriptionId string, resourceGroupName string, vaultName string, replicationAlertSettingName string) ReplicationAlertSettingId {
	return ReplicationAlertSettingId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		VaultName:                   vaultName,
		ReplicationAlertSettingName: replicationAlertSettingName,
	}
}

// ParseReplicationAlertSettingID parses 'input' into a ReplicationAlertSettingId
func ParseReplicationAlertSettingID(input string) (*ReplicationAlertSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationAlertSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationAlertSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReplicationAlertSettingIDInsensitively parses 'input' case-insensitively into a ReplicationAlertSettingId
// note: this method should only be used for API response data and not user input
func ParseReplicationAlertSettingIDInsensitively(input string) (*ReplicationAlertSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationAlertSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationAlertSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicationAlertSettingId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ReplicationAlertSettingName, ok = input.Parsed["replicationAlertSettingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationAlertSettingName", input)
	}

	return nil
}

// ValidateReplicationAlertSettingID checks that 'input' can be parsed as a Replication Alert Setting ID
func ValidateReplicationAlertSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationAlertSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Alert Setting ID
func (id ReplicationAlertSettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationAlertSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationAlertSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Alert Setting ID
func (id ReplicationAlertSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "resourceName"),
		resourceids.StaticSegment("staticReplicationAlertSettings", "replicationAlertSettings", "replicationAlertSettings"),
		resourceids.UserSpecifiedSegment("replicationAlertSettingName", "alertSettingName"),
	}
}

// String returns a human-readable description of this Replication Alert Setting ID
func (id ReplicationAlertSettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Alert Setting Name: %q", id.ReplicationAlertSettingName),
	}
	return fmt.Sprintf("Replication Alert Setting (%s)", strings.Join(components, "\n"))
}
