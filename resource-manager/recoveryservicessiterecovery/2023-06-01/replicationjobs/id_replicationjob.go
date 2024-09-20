package replicationjobs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ReplicationJobId{})
}

var _ resourceids.ResourceId = &ReplicationJobId{}

// ReplicationJobId is a struct representing the Resource ID for a Replication Job
type ReplicationJobId struct {
	SubscriptionId     string
	ResourceGroupName  string
	VaultName          string
	ReplicationJobName string
}

// NewReplicationJobID returns a new ReplicationJobId struct
func NewReplicationJobID(subscriptionId string, resourceGroupName string, vaultName string, replicationJobName string) ReplicationJobId {
	return ReplicationJobId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		VaultName:          vaultName,
		ReplicationJobName: replicationJobName,
	}
}

// ParseReplicationJobID parses 'input' into a ReplicationJobId
func ParseReplicationJobID(input string) (*ReplicationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReplicationJobIDInsensitively parses 'input' case-insensitively into a ReplicationJobId
// note: this method should only be used for API response data and not user input
func ParseReplicationJobIDInsensitively(input string) (*ReplicationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReplicationJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicationJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicationJobId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ReplicationJobName, ok = input.Parsed["replicationJobName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationJobName", input)
	}

	return nil
}

// ValidateReplicationJobID checks that 'input' can be parsed as a Replication Job ID
func ValidateReplicationJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Job ID
func (id ReplicationJobId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationJobs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationJobName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Job ID
func (id ReplicationJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "resourceName"),
		resourceids.StaticSegment("staticReplicationJobs", "replicationJobs", "replicationJobs"),
		resourceids.UserSpecifiedSegment("replicationJobName", "jobName"),
	}
}

// String returns a human-readable description of this Replication Job ID
func (id ReplicationJobId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Job Name: %q", id.ReplicationJobName),
	}
	return fmt.Sprintf("Replication Job (%s)", strings.Join(components, "\n"))
}
