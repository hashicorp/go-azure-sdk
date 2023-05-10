package azurebackupjob

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = BackupJobId{}

// BackupJobId is a struct representing the Resource ID for a Backup Job
type BackupJobId struct {
	SubscriptionId    string
	ResourceGroupName string
	BackupVaultName   string
	JobId             string
}

// NewBackupJobID returns a new BackupJobId struct
func NewBackupJobID(subscriptionId string, resourceGroupName string, backupVaultName string, jobId string) BackupJobId {
	return BackupJobId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		BackupVaultName:   backupVaultName,
		JobId:             jobId,
	}
}

// ParseBackupJobID parses 'input' into a BackupJobId
func ParseBackupJobID(input string) (*BackupJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupJobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.BackupVaultName, ok = parsed.Parsed["backupVaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupVaultName", *parsed)
	}

	if id.JobId, ok = parsed.Parsed["jobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobId", *parsed)
	}

	return &id, nil
}

// ParseBackupJobIDInsensitively parses 'input' case-insensitively into a BackupJobId
// note: this method should only be used for API response data and not user input
func ParseBackupJobIDInsensitively(input string) (*BackupJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupJobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.BackupVaultName, ok = parsed.Parsed["backupVaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "backupVaultName", *parsed)
	}

	if id.JobId, ok = parsed.Parsed["jobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobId", *parsed)
	}

	return &id, nil
}

// ValidateBackupJobID checks that 'input' can be parsed as a Backup Job ID
func ValidateBackupJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBackupJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Backup Job ID
func (id BackupJobId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataProtection/backupVaults/%s/backupJobs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.BackupVaultName, id.JobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Backup Job ID
func (id BackupJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataProtection", "Microsoft.DataProtection", "Microsoft.DataProtection"),
		resourceids.StaticSegment("staticBackupVaults", "backupVaults", "backupVaults"),
		resourceids.UserSpecifiedSegment("backupVaultName", "backupVaultValue"),
		resourceids.StaticSegment("staticBackupJobs", "backupJobs", "backupJobs"),
		resourceids.UserSpecifiedSegment("jobId", "jobIdValue"),
	}
}

// String returns a human-readable description of this Backup Job ID
func (id BackupJobId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Backup Vault Name: %q", id.BackupVaultName),
		fmt.Sprintf("Job: %q", id.JobId),
	}
	return fmt.Sprintf("Backup Job (%s)", strings.Join(components, "\n"))
}
