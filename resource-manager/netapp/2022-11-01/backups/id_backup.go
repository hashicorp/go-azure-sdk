package backups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = BackupId{}

// BackupId is a struct representing the Resource ID for a Backup
type BackupId struct {
	SubscriptionId    string
	ResourceGroupName string
	NetAppAccountName string
	CapacityPoolName  string
	VolumeName        string
	BackupName        string
}

// NewBackupID returns a new BackupId struct
func NewBackupID(subscriptionId string, resourceGroupName string, netAppAccountName string, capacityPoolName string, volumeName string, backupName string) BackupId {
	return BackupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NetAppAccountName: netAppAccountName,
		CapacityPoolName:  capacityPoolName,
		VolumeName:        volumeName,
		BackupName:        backupName,
	}
}

// ParseBackupID parses 'input' into a BackupId
func ParseBackupID(input string) (*BackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.NetAppAccountName, ok = parsed.Parsed["netAppAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "netAppAccountName", *parsed)
	}

	if id.CapacityPoolName, ok = parsed.Parsed["capacityPoolName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "capacityPoolName", *parsed)
	}

	if id.VolumeName, ok = parsed.Parsed["volumeName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "volumeName", *parsed)
	}

	if id.BackupName, ok = parsed.Parsed["backupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupName", *parsed)
	}

	return &id, nil
}

// ParseBackupIDInsensitively parses 'input' case-insensitively into a BackupId
// note: this method should only be used for API response data and not user input
func ParseBackupIDInsensitively(input string) (*BackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.NetAppAccountName, ok = parsed.Parsed["netAppAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "netAppAccountName", *parsed)
	}

	if id.CapacityPoolName, ok = parsed.Parsed["capacityPoolName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "capacityPoolName", *parsed)
	}

	if id.VolumeName, ok = parsed.Parsed["volumeName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "volumeName", *parsed)
	}

	if id.BackupName, ok = parsed.Parsed["backupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupName", *parsed)
	}

	return &id, nil
}

// ValidateBackupID checks that 'input' can be parsed as a Backup ID
func ValidateBackupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBackupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Backup ID
func (id BackupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetApp/netAppAccounts/%s/capacityPools/%s/volumes/%s/backups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetAppAccountName, id.CapacityPoolName, id.VolumeName, id.BackupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Backup ID
func (id BackupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticNetAppAccounts", "netAppAccounts", "netAppAccounts"),
		resourceids.UserSpecifiedSegment("netAppAccountName", "netAppAccountValue"),
		resourceids.StaticSegment("staticCapacityPools", "capacityPools", "capacityPools"),
		resourceids.UserSpecifiedSegment("capacityPoolName", "capacityPoolValue"),
		resourceids.StaticSegment("staticVolumes", "volumes", "volumes"),
		resourceids.UserSpecifiedSegment("volumeName", "volumeValue"),
		resourceids.StaticSegment("staticBackups", "backups", "backups"),
		resourceids.UserSpecifiedSegment("backupName", "backupValue"),
	}
}

// String returns a human-readable description of this Backup ID
func (id BackupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Net App Account Name: %q", id.NetAppAccountName),
		fmt.Sprintf("Capacity Pool Name: %q", id.CapacityPoolName),
		fmt.Sprintf("Volume Name: %q", id.VolumeName),
		fmt.Sprintf("Backup Name: %q", id.BackupName),
	}
	return fmt.Sprintf("Backup (%s)", strings.Join(components, "\n"))
}
