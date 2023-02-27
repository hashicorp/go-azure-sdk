// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protectionintent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BackupProtectionIntentId{}

// BackupProtectionIntentId is a struct representing the Resource ID for a Backup Protection Intent
type BackupProtectionIntentId struct {
	SubscriptionId             string
	ResourceGroupName          string
	VaultName                  string
	BackupFabricName           string
	BackupProtectionIntentName string
}

// NewBackupProtectionIntentID returns a new BackupProtectionIntentId struct
func NewBackupProtectionIntentID(subscriptionId string, resourceGroupName string, vaultName string, backupFabricName string, backupProtectionIntentName string) BackupProtectionIntentId {
	return BackupProtectionIntentId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		VaultName:                  vaultName,
		BackupFabricName:           backupFabricName,
		BackupProtectionIntentName: backupProtectionIntentName,
	}
}

// ParseBackupProtectionIntentID parses 'input' into a BackupProtectionIntentId
func ParseBackupProtectionIntentID(input string) (*BackupProtectionIntentId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupProtectionIntentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupProtectionIntentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.BackupFabricName, ok = parsed.Parsed["backupFabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'backupFabricName' was not found in the resource id %q", input)
	}

	if id.BackupProtectionIntentName, ok = parsed.Parsed["backupProtectionIntentName"]; !ok {
		return nil, fmt.Errorf("the segment 'backupProtectionIntentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseBackupProtectionIntentIDInsensitively parses 'input' case-insensitively into a BackupProtectionIntentId
// note: this method should only be used for API response data and not user input
func ParseBackupProtectionIntentIDInsensitively(input string) (*BackupProtectionIntentId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupProtectionIntentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupProtectionIntentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.BackupFabricName, ok = parsed.Parsed["backupFabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'backupFabricName' was not found in the resource id %q", input)
	}

	if id.BackupProtectionIntentName, ok = parsed.Parsed["backupProtectionIntentName"]; !ok {
		return nil, fmt.Errorf("the segment 'backupProtectionIntentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateBackupProtectionIntentID checks that 'input' can be parsed as a Backup Protection Intent ID
func ValidateBackupProtectionIntentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBackupProtectionIntentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Backup Protection Intent ID
func (id BackupProtectionIntentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/backupFabrics/%s/backupProtectionIntent/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.BackupFabricName, id.BackupProtectionIntentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Backup Protection Intent ID
func (id BackupProtectionIntentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticBackupFabrics", "backupFabrics", "backupFabrics"),
		resourceids.UserSpecifiedSegment("backupFabricName", "backupFabricValue"),
		resourceids.StaticSegment("staticBackupProtectionIntent", "backupProtectionIntent", "backupProtectionIntent"),
		resourceids.UserSpecifiedSegment("backupProtectionIntentName", "backupProtectionIntentValue"),
	}
}

// String returns a human-readable description of this Backup Protection Intent ID
func (id BackupProtectionIntentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Backup Fabric Name: %q", id.BackupFabricName),
		fmt.Sprintf("Backup Protection Intent Name: %q", id.BackupProtectionIntentName),
	}
	return fmt.Sprintf("Backup Protection Intent (%s)", strings.Join(components, "\n"))
}
