package recoverypoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RecoveryPointId{}

// RecoveryPointId is a struct representing the Resource ID for a Recovery Point
type RecoveryPointId struct {
	SubscriptionId     string
	ResourceGroupName  string
	BackupVaultName    string
	BackupInstanceName string
	RecoveryPointId    string
}

// NewRecoveryPointID returns a new RecoveryPointId struct
func NewRecoveryPointID(subscriptionId string, resourceGroupName string, backupVaultName string, backupInstanceName string, recoveryPointId string) RecoveryPointId {
	return RecoveryPointId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		BackupVaultName:    backupVaultName,
		BackupInstanceName: backupInstanceName,
		RecoveryPointId:    recoveryPointId,
	}
}

// ParseRecoveryPointID parses 'input' into a RecoveryPointId
func ParseRecoveryPointID(input string) (*RecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecoveryPointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecoveryPointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.BackupVaultName, ok = parsed.Parsed["backupVaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupVaultName", *parsed)
	}

	if id.BackupInstanceName, ok = parsed.Parsed["backupInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupInstanceName", *parsed)
	}

	if id.RecoveryPointId, ok = parsed.Parsed["recoveryPointId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recoveryPointId", *parsed)
	}

	return &id, nil
}

// ParseRecoveryPointIDInsensitively parses 'input' case-insensitively into a RecoveryPointId
// note: this method should only be used for API response data and not user input
func ParseRecoveryPointIDInsensitively(input string) (*RecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecoveryPointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecoveryPointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.BackupVaultName, ok = parsed.Parsed["backupVaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupVaultName", *parsed)
	}

	if id.BackupInstanceName, ok = parsed.Parsed["backupInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupInstanceName", *parsed)
	}

	if id.RecoveryPointId, ok = parsed.Parsed["recoveryPointId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recoveryPointId", *parsed)
	}

	return &id, nil
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
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataProtection/backupVaults/%s/backupInstances/%s/recoveryPoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.BackupVaultName, id.BackupInstanceName, id.RecoveryPointId)
}

// Segments returns a slice of Resource ID Segments which comprise this Recovery Point ID
func (id RecoveryPointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataProtection", "Microsoft.DataProtection", "Microsoft.DataProtection"),
		resourceids.StaticSegment("staticBackupVaults", "backupVaults", "backupVaults"),
		resourceids.UserSpecifiedSegment("backupVaultName", "backupVaultValue"),
		resourceids.StaticSegment("staticBackupInstances", "backupInstances", "backupInstances"),
		resourceids.UserSpecifiedSegment("backupInstanceName", "backupInstanceValue"),
		resourceids.StaticSegment("staticRecoveryPoints", "recoveryPoints", "recoveryPoints"),
		resourceids.UserSpecifiedSegment("recoveryPointId", "recoveryPointIdValue"),
	}
}

// String returns a human-readable description of this Recovery Point ID
func (id RecoveryPointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Backup Vault Name: %q", id.BackupVaultName),
		fmt.Sprintf("Backup Instance Name: %q", id.BackupInstanceName),
		fmt.Sprintf("Recovery Point: %q", id.RecoveryPointId),
	}
	return fmt.Sprintf("Recovery Point (%s)", strings.Join(components, "\n"))
}
