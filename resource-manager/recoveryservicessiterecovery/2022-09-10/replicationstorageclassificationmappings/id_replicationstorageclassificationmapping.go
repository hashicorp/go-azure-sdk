package replicationstorageclassificationmappings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationStorageClassificationMappingId{}

// ReplicationStorageClassificationMappingId is a struct representing the Resource ID for a Replication Storage Classification Mapping
type ReplicationStorageClassificationMappingId struct {
	SubscriptionId                              string
	ResourceGroupName                           string
	VaultName                                   string
	ReplicationFabricName                       string
	ReplicationStorageClassificationName        string
	ReplicationStorageClassificationMappingName string
}

// NewReplicationStorageClassificationMappingID returns a new ReplicationStorageClassificationMappingId struct
func NewReplicationStorageClassificationMappingID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationStorageClassificationName string, replicationStorageClassificationMappingName string) ReplicationStorageClassificationMappingId {
	return ReplicationStorageClassificationMappingId{
		SubscriptionId:                              subscriptionId,
		ResourceGroupName:                           resourceGroupName,
		VaultName:                                   vaultName,
		ReplicationFabricName:                       replicationFabricName,
		ReplicationStorageClassificationName:        replicationStorageClassificationName,
		ReplicationStorageClassificationMappingName: replicationStorageClassificationMappingName,
	}
}

// ParseReplicationStorageClassificationMappingID parses 'input' into a ReplicationStorageClassificationMappingId
func ParseReplicationStorageClassificationMappingID(input string) (*ReplicationStorageClassificationMappingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationStorageClassificationMappingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationStorageClassificationMappingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.ReplicationFabricName, ok = parsed.Parsed["replicationFabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationFabricName' was not found in the resource id %q", input)
	}

	if id.ReplicationStorageClassificationName, ok = parsed.Parsed["replicationStorageClassificationName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationStorageClassificationName' was not found in the resource id %q", input)
	}

	if id.ReplicationStorageClassificationMappingName, ok = parsed.Parsed["replicationStorageClassificationMappingName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationStorageClassificationMappingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationStorageClassificationMappingIDInsensitively parses 'input' case-insensitively into a ReplicationStorageClassificationMappingId
// note: this method should only be used for API response data and not user input
func ParseReplicationStorageClassificationMappingIDInsensitively(input string) (*ReplicationStorageClassificationMappingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationStorageClassificationMappingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationStorageClassificationMappingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.ReplicationFabricName, ok = parsed.Parsed["replicationFabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationFabricName' was not found in the resource id %q", input)
	}

	if id.ReplicationStorageClassificationName, ok = parsed.Parsed["replicationStorageClassificationName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationStorageClassificationName' was not found in the resource id %q", input)
	}

	if id.ReplicationStorageClassificationMappingName, ok = parsed.Parsed["replicationStorageClassificationMappingName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationStorageClassificationMappingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationStorageClassificationMappingID checks that 'input' can be parsed as a Replication Storage Classification Mapping ID
func ValidateReplicationStorageClassificationMappingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationStorageClassificationMappingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Storage Classification Mapping ID
func (id ReplicationStorageClassificationMappingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationStorageClassifications/%s/replicationStorageClassificationMappings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationStorageClassificationName, id.ReplicationStorageClassificationMappingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Storage Classification Mapping ID
func (id ReplicationStorageClassificationMappingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticReplicationFabrics", "replicationFabrics", "replicationFabrics"),
		resourceids.UserSpecifiedSegment("replicationFabricName", "replicationFabricValue"),
		resourceids.StaticSegment("staticReplicationStorageClassifications", "replicationStorageClassifications", "replicationStorageClassifications"),
		resourceids.UserSpecifiedSegment("replicationStorageClassificationName", "replicationStorageClassificationValue"),
		resourceids.StaticSegment("staticReplicationStorageClassificationMappings", "replicationStorageClassificationMappings", "replicationStorageClassificationMappings"),
		resourceids.UserSpecifiedSegment("replicationStorageClassificationMappingName", "replicationStorageClassificationMappingValue"),
	}
}

// String returns a human-readable description of this Replication Storage Classification Mapping ID
func (id ReplicationStorageClassificationMappingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replication Storage Classification Name: %q", id.ReplicationStorageClassificationName),
		fmt.Sprintf("Replication Storage Classification Mapping Name: %q", id.ReplicationStorageClassificationMappingName),
	}
	return fmt.Sprintf("Replication Storage Classification Mapping (%s)", strings.Join(components, "\n"))
}
