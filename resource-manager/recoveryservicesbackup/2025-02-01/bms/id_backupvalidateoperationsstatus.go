package bms

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BackupValidateOperationsStatusId{})
}

var _ resourceids.ResourceId = &BackupValidateOperationsStatusId{}

// BackupValidateOperationsStatusId is a struct representing the Resource ID for a Backup Validate Operations Status
type BackupValidateOperationsStatusId struct {
	SubscriptionId    string
	ResourceGroupName string
	VaultName         string
	OperationId       string
}

// NewBackupValidateOperationsStatusID returns a new BackupValidateOperationsStatusId struct
func NewBackupValidateOperationsStatusID(subscriptionId string, resourceGroupName string, vaultName string, operationId string) BackupValidateOperationsStatusId {
	return BackupValidateOperationsStatusId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		VaultName:         vaultName,
		OperationId:       operationId,
	}
}

// ParseBackupValidateOperationsStatusID parses 'input' into a BackupValidateOperationsStatusId
func ParseBackupValidateOperationsStatusID(input string) (*BackupValidateOperationsStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BackupValidateOperationsStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BackupValidateOperationsStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBackupValidateOperationsStatusIDInsensitively parses 'input' case-insensitively into a BackupValidateOperationsStatusId
// note: this method should only be used for API response data and not user input
func ParseBackupValidateOperationsStatusIDInsensitively(input string) (*BackupValidateOperationsStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BackupValidateOperationsStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BackupValidateOperationsStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BackupValidateOperationsStatusId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.OperationId, ok = input.Parsed["operationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationId", input)
	}

	return nil
}

// ValidateBackupValidateOperationsStatusID checks that 'input' can be parsed as a Backup Validate Operations Status ID
func ValidateBackupValidateOperationsStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBackupValidateOperationsStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Backup Validate Operations Status ID
func (id BackupValidateOperationsStatusId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/backupValidateOperationsStatuses/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Backup Validate Operations Status ID
func (id BackupValidateOperationsStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultName"),
		resourceids.StaticSegment("staticBackupValidateOperationsStatuses", "backupValidateOperationsStatuses", "backupValidateOperationsStatuses"),
		resourceids.UserSpecifiedSegment("operationId", "operationId"),
	}
}

// String returns a human-readable description of this Backup Validate Operations Status ID
func (id BackupValidateOperationsStatusId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Backup Validate Operations Status (%s)", strings.Join(components, "\n"))
}
