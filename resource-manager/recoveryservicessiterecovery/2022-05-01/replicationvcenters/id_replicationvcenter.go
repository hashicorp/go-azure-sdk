package replicationvcenters

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationvCenterId{}

// ReplicationvCenterId is a struct representing the Resource ID for a Replicationv Center
type ReplicationvCenterId struct {
	SubscriptionId         string
	ResourceGroupName      string
	VaultName              string
	ReplicationFabricName  string
	ReplicationvCenterName string
}

// NewReplicationvCenterID returns a new ReplicationvCenterId struct
func NewReplicationvCenterID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationvCenterName string) ReplicationvCenterId {
	return ReplicationvCenterId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		VaultName:              vaultName,
		ReplicationFabricName:  replicationFabricName,
		ReplicationvCenterName: replicationvCenterName,
	}
}

// ParseReplicationvCenterID parses 'input' into a ReplicationvCenterId
func ParseReplicationvCenterID(input string) (*ReplicationvCenterId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationvCenterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationvCenterId{}

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

	if id.ReplicationvCenterName, ok = parsed.Parsed["replicationvCenterName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationvCenterName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationvCenterIDInsensitively parses 'input' case-insensitively into a ReplicationvCenterId
// note: this method should only be used for API response data and not user input
func ParseReplicationvCenterIDInsensitively(input string) (*ReplicationvCenterId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationvCenterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationvCenterId{}

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

	if id.ReplicationvCenterName, ok = parsed.Parsed["replicationvCenterName"]; !ok {
		return nil, fmt.Errorf("the segment 'replicationvCenterName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationvCenterID checks that 'input' can be parsed as a Replicationv Center ID
func ValidateReplicationvCenterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationvCenterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replicationv Center ID
func (id ReplicationvCenterId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationvCenters/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationvCenterName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replicationv Center ID
func (id ReplicationvCenterId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReplicationvCenters", "replicationvCenters", "replicationvCenters"),
		resourceids.UserSpecifiedSegment("replicationvCenterName", "replicationvCenterValue"),
	}
}

// String returns a human-readable description of this Replicationv Center ID
func (id ReplicationvCenterId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replicationv Center Name: %q", id.ReplicationvCenterName),
	}
	return fmt.Sprintf("Replicationv Center (%s)", strings.Join(components, "\n"))
}
