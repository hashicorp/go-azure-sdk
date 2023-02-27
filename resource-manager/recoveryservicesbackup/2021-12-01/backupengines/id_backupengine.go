// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupengines

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BackupEngineId{}

// BackupEngineId is a struct representing the Resource ID for a Backup Engine
type BackupEngineId struct {
	SubscriptionId    string
	ResourceGroupName string
	VaultName         string
	BackupEngineName  string
}

// NewBackupEngineID returns a new BackupEngineId struct
func NewBackupEngineID(subscriptionId string, resourceGroupName string, vaultName string, backupEngineName string) BackupEngineId {
	return BackupEngineId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		VaultName:         vaultName,
		BackupEngineName:  backupEngineName,
	}
}

// ParseBackupEngineID parses 'input' into a BackupEngineId
func ParseBackupEngineID(input string) (*BackupEngineId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupEngineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupEngineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.BackupEngineName, ok = parsed.Parsed["backupEngineName"]; !ok {
		return nil, fmt.Errorf("the segment 'backupEngineName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseBackupEngineIDInsensitively parses 'input' case-insensitively into a BackupEngineId
// note: this method should only be used for API response data and not user input
func ParseBackupEngineIDInsensitively(input string) (*BackupEngineId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupEngineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupEngineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.BackupEngineName, ok = parsed.Parsed["backupEngineName"]; !ok {
		return nil, fmt.Errorf("the segment 'backupEngineName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateBackupEngineID checks that 'input' can be parsed as a Backup Engine ID
func ValidateBackupEngineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBackupEngineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Backup Engine ID
func (id BackupEngineId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/backupEngines/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.BackupEngineName)
}

// Segments returns a slice of Resource ID Segments which comprise this Backup Engine ID
func (id BackupEngineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticBackupEngines", "backupEngines", "backupEngines"),
		resourceids.UserSpecifiedSegment("backupEngineName", "backupEngineValue"),
	}
}

// String returns a human-readable description of this Backup Engine ID
func (id BackupEngineId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Backup Engine Name: %q", id.BackupEngineName),
	}
	return fmt.Sprintf("Backup Engine (%s)", strings.Join(components, "\n"))
}
