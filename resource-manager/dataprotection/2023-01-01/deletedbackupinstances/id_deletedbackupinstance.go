package deletedbackupinstances

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DeletedBackupInstanceId{}

// DeletedBackupInstanceId is a struct representing the Resource ID for a Deleted Backup Instance
type DeletedBackupInstanceId struct {
	SubscriptionId            string
	ResourceGroupName         string
	BackupVaultName           string
	DeletedBackupInstanceName string
}

// NewDeletedBackupInstanceID returns a new DeletedBackupInstanceId struct
func NewDeletedBackupInstanceID(subscriptionId string, resourceGroupName string, backupVaultName string, deletedBackupInstanceName string) DeletedBackupInstanceId {
	return DeletedBackupInstanceId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		BackupVaultName:           backupVaultName,
		DeletedBackupInstanceName: deletedBackupInstanceName,
	}
}

// ParseDeletedBackupInstanceID parses 'input' into a DeletedBackupInstanceId
func ParseDeletedBackupInstanceID(input string) (*DeletedBackupInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeletedBackupInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeletedBackupInstanceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.BackupVaultName, ok = parsed.Parsed["backupVaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupVaultName", *parsed)
	}

	if id.DeletedBackupInstanceName, ok = parsed.Parsed["deletedBackupInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deletedBackupInstanceName", *parsed)
	}

	return &id, nil
}

// ParseDeletedBackupInstanceIDInsensitively parses 'input' case-insensitively into a DeletedBackupInstanceId
// note: this method should only be used for API response data and not user input
func ParseDeletedBackupInstanceIDInsensitively(input string) (*DeletedBackupInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeletedBackupInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeletedBackupInstanceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.BackupVaultName, ok = parsed.Parsed["backupVaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupVaultName", *parsed)
	}

	if id.DeletedBackupInstanceName, ok = parsed.Parsed["deletedBackupInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deletedBackupInstanceName", *parsed)
	}

	return &id, nil
}

// ValidateDeletedBackupInstanceID checks that 'input' can be parsed as a Deleted Backup Instance ID
func ValidateDeletedBackupInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeletedBackupInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Deleted Backup Instance ID
func (id DeletedBackupInstanceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataProtection/backupVaults/%s/deletedBackupInstances/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.BackupVaultName, id.DeletedBackupInstanceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Deleted Backup Instance ID
func (id DeletedBackupInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataProtection", "Microsoft.DataProtection", "Microsoft.DataProtection"),
		resourceids.StaticSegment("staticBackupVaults", "backupVaults", "backupVaults"),
		resourceids.UserSpecifiedSegment("backupVaultName", "backupVaultValue"),
		resourceids.StaticSegment("staticDeletedBackupInstances", "deletedBackupInstances", "deletedBackupInstances"),
		resourceids.UserSpecifiedSegment("deletedBackupInstanceName", "deletedBackupInstanceValue"),
	}
}

// String returns a human-readable description of this Deleted Backup Instance ID
func (id DeletedBackupInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Backup Vault Name: %q", id.BackupVaultName),
		fmt.Sprintf("Deleted Backup Instance Name: %q", id.DeletedBackupInstanceName),
	}
	return fmt.Sprintf("Deleted Backup Instance (%s)", strings.Join(components, "\n"))
}
