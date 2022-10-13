package replicationlogicalnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationLogicalNetworkId{}

// ReplicationLogicalNetworkId is a struct representing the Resource ID for a Replication Logical Network
type ReplicationLogicalNetworkId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ResourceName       string
	FabricName         string
	LogicalNetworkName string
}

// NewReplicationLogicalNetworkID returns a new ReplicationLogicalNetworkId struct
func NewReplicationLogicalNetworkID(subscriptionId string, resourceGroupName string, resourceName string, fabricName string, logicalNetworkName string) ReplicationLogicalNetworkId {
	return ReplicationLogicalNetworkId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ResourceName:       resourceName,
		FabricName:         fabricName,
		LogicalNetworkName: logicalNetworkName,
	}
}

// ParseReplicationLogicalNetworkID parses 'input' into a ReplicationLogicalNetworkId
func ParseReplicationLogicalNetworkID(input string) (*ReplicationLogicalNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationLogicalNetworkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationLogicalNetworkId{}

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

	if id.LogicalNetworkName, ok = parsed.Parsed["logicalNetworkName"]; !ok {
		return nil, fmt.Errorf("the segment 'logicalNetworkName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationLogicalNetworkIDInsensitively parses 'input' case-insensitively into a ReplicationLogicalNetworkId
// note: this method should only be used for API response data and not user input
func ParseReplicationLogicalNetworkIDInsensitively(input string) (*ReplicationLogicalNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationLogicalNetworkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationLogicalNetworkId{}

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

	if id.LogicalNetworkName, ok = parsed.Parsed["logicalNetworkName"]; !ok {
		return nil, fmt.Errorf("the segment 'logicalNetworkName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationLogicalNetworkID checks that 'input' can be parsed as a Replication Logical Network ID
func ValidateReplicationLogicalNetworkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationLogicalNetworkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Logical Network ID
func (id ReplicationLogicalNetworkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationLogicalNetworks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName, id.FabricName, id.LogicalNetworkName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Logical Network ID
func (id ReplicationLogicalNetworkId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReplicationLogicalNetworks", "replicationLogicalNetworks", "replicationLogicalNetworks"),
		resourceids.UserSpecifiedSegment("logicalNetworkName", "logicalNetworkValue"),
	}
}

// String returns a human-readable description of this Replication Logical Network ID
func (id ReplicationLogicalNetworkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Fabric Name: %q", id.FabricName),
		fmt.Sprintf("Logical Network Name: %q", id.LogicalNetworkName),
	}
	return fmt.Sprintf("Replication Logical Network (%s)", strings.Join(components, "\n"))
}
