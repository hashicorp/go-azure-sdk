package replicationstorageclassificationmappings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationStorageClassificationMappingId{}

// ReplicationStorageClassificationMappingId is a struct representing the Resource ID for a Replication Storage Classification Mapping
type ReplicationStorageClassificationMappingId struct {
	SubscriptionId                   string
	ResourceGroupName                string
	ResourceName                     string
	FabricName                       string
	StorageClassificationName        string
	StorageClassificationMappingName string
}

// NewReplicationStorageClassificationMappingID returns a new ReplicationStorageClassificationMappingId struct
func NewReplicationStorageClassificationMappingID(subscriptionId string, resourceGroupName string, resourceName string, fabricName string, storageClassificationName string, storageClassificationMappingName string) ReplicationStorageClassificationMappingId {
	return ReplicationStorageClassificationMappingId{
		SubscriptionId:                   subscriptionId,
		ResourceGroupName:                resourceGroupName,
		ResourceName:                     resourceName,
		FabricName:                       fabricName,
		StorageClassificationName:        storageClassificationName,
		StorageClassificationMappingName: storageClassificationMappingName,
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

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.FabricName, ok = parsed.Parsed["fabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'fabricName' was not found in the resource id %q", input)
	}

	if id.StorageClassificationName, ok = parsed.Parsed["storageClassificationName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageClassificationName' was not found in the resource id %q", input)
	}

	if id.StorageClassificationMappingName, ok = parsed.Parsed["storageClassificationMappingName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageClassificationMappingName' was not found in the resource id %q", input)
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

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.FabricName, ok = parsed.Parsed["fabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'fabricName' was not found in the resource id %q", input)
	}

	if id.StorageClassificationName, ok = parsed.Parsed["storageClassificationName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageClassificationName' was not found in the resource id %q", input)
	}

	if id.StorageClassificationMappingName, ok = parsed.Parsed["storageClassificationMappingName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageClassificationMappingName' was not found in the resource id %q", input)
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
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName, id.FabricName, id.StorageClassificationName, id.StorageClassificationMappingName)
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
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticReplicationFabrics", "replicationFabrics", "replicationFabrics"),
		resourceids.UserSpecifiedSegment("fabricName", "fabricValue"),
		resourceids.StaticSegment("staticReplicationStorageClassifications", "replicationStorageClassifications", "replicationStorageClassifications"),
		resourceids.UserSpecifiedSegment("storageClassificationName", "storageClassificationValue"),
		resourceids.StaticSegment("staticReplicationStorageClassificationMappings", "replicationStorageClassificationMappings", "replicationStorageClassificationMappings"),
		resourceids.UserSpecifiedSegment("storageClassificationMappingName", "storageClassificationMappingValue"),
	}
}

// String returns a human-readable description of this Replication Storage Classification Mapping ID
func (id ReplicationStorageClassificationMappingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Fabric Name: %q", id.FabricName),
		fmt.Sprintf("Storage Classification Name: %q", id.StorageClassificationName),
		fmt.Sprintf("Storage Classification Mapping Name: %q", id.StorageClassificationMappingName),
	}
	return fmt.Sprintf("Replication Storage Classification Mapping (%s)", strings.Join(components, "\n"))
}
