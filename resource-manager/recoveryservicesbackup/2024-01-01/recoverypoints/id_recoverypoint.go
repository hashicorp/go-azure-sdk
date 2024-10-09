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
	SubscriptionId          string
	ResourceGroupName       string
	VaultName               string
	BackupFabricName        string
	ProtectionContainerName string
	ProtectedItemName       string
	RecoveryPointId         string
}

// NewRecoveryPointID returns a new RecoveryPointId struct
func NewRecoveryPointID(subscriptionId string, resourceGroupName string, vaultName string, backupFabricName string, protectionContainerName string, protectedItemName string, recoveryPointId string) RecoveryPointId {
	return RecoveryPointId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		VaultName:               vaultName,
		BackupFabricName:        backupFabricName,
		ProtectionContainerName: protectionContainerName,
		ProtectedItemName:       protectedItemName,
		RecoveryPointId:         recoveryPointId,
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

	if id.BackupFabricName, ok = input.Parsed["backupFabricName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "backupFabricName", input)
	}

	if id.ProtectionContainerName, ok = input.Parsed["protectionContainerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "protectionContainerName", input)
	}

	if id.ProtectedItemName, ok = input.Parsed["protectedItemName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "protectedItemName", input)
	}

	if id.RecoveryPointId, ok = input.Parsed["recoveryPointId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "recoveryPointId", input)
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
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/backupFabrics/%s/protectionContainers/%s/protectedItems/%s/recoveryPoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.BackupFabricName, id.ProtectionContainerName, id.ProtectedItemName, id.RecoveryPointId)
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
		resourceids.UserSpecifiedSegment("vaultName", "vaultName"),
		resourceids.StaticSegment("staticBackupFabrics", "backupFabrics", "backupFabrics"),
		resourceids.UserSpecifiedSegment("backupFabricName", "backupFabricName"),
		resourceids.StaticSegment("staticProtectionContainers", "protectionContainers", "protectionContainers"),
		resourceids.UserSpecifiedSegment("protectionContainerName", "protectionContainerName"),
		resourceids.StaticSegment("staticProtectedItems", "protectedItems", "protectedItems"),
		resourceids.UserSpecifiedSegment("protectedItemName", "protectedItemName"),
		resourceids.StaticSegment("staticRecoveryPoints", "recoveryPoints", "recoveryPoints"),
		resourceids.UserSpecifiedSegment("recoveryPointId", "recoveryPointId"),
	}
}

// String returns a human-readable description of this Recovery Point ID
func (id RecoveryPointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Backup Fabric Name: %q", id.BackupFabricName),
		fmt.Sprintf("Protection Container Name: %q", id.ProtectionContainerName),
		fmt.Sprintf("Protected Item Name: %q", id.ProtectedItemName),
		fmt.Sprintf("Recovery Point: %q", id.RecoveryPointId),
	}
	return fmt.Sprintf("Recovery Point (%s)", strings.Join(components, "\n"))
}
