package protectionintent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BackupProtectionIntentId{})
}

var _ resourceids.ResourceId = &BackupProtectionIntentId{}

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
	parser := resourceids.NewParserFromResourceIdType(&BackupProtectionIntentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BackupProtectionIntentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBackupProtectionIntentIDInsensitively parses 'input' case-insensitively into a BackupProtectionIntentId
// note: this method should only be used for API response data and not user input
func ParseBackupProtectionIntentIDInsensitively(input string) (*BackupProtectionIntentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BackupProtectionIntentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BackupProtectionIntentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BackupProtectionIntentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.BackupProtectionIntentName, ok = input.Parsed["backupProtectionIntentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "backupProtectionIntentName", input)
	}

	return nil
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
