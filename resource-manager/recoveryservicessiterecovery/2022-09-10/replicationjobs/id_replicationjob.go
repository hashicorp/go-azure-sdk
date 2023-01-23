package replicationjobs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationJobId{}

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
	parser := resourceids.NewParserFromResourceIdType(ReplicationJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationJobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.ReplicationJobName, ok = parsed.Parsed["replicationJobName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationJobName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationJobIDInsensitively parses 'input' case-insensitively into a ReplicationJobId
// note: this method should only be used for API response data and not user input
func ParseReplicationJobIDInsensitively(input string) (*ReplicationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationJobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.ReplicationJobName, ok = parsed.Parsed["replicationJobName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationJobName' was not found in the resource id %q", input)
	}

	return &id, nil
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
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticReplicationJobs", "replicationJobs", "replicationJobs"),
		resourceids.UserSpecifiedSegment("replicationJobName", "replicationJobValue"),
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
