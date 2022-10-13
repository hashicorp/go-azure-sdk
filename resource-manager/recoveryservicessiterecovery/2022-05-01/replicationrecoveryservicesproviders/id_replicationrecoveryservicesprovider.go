package replicationrecoveryservicesproviders

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationRecoveryServicesProviderId{}

// ReplicationRecoveryServicesProviderId is a struct representing the Resource ID for a Replication Recovery Services Provider
type ReplicationRecoveryServicesProviderId struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourceName      string
	FabricName        string
	ProviderName      string
}

// NewReplicationRecoveryServicesProviderID returns a new ReplicationRecoveryServicesProviderId struct
func NewReplicationRecoveryServicesProviderID(subscriptionId string, resourceGroupName string, resourceName string, fabricName string, providerName string) ReplicationRecoveryServicesProviderId {
	return ReplicationRecoveryServicesProviderId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ResourceName:      resourceName,
		FabricName:        fabricName,
		ProviderName:      providerName,
	}
}

// ParseReplicationRecoveryServicesProviderID parses 'input' into a ReplicationRecoveryServicesProviderId
func ParseReplicationRecoveryServicesProviderID(input string) (*ReplicationRecoveryServicesProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationRecoveryServicesProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationRecoveryServicesProviderId{}

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

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationRecoveryServicesProviderIDInsensitively parses 'input' case-insensitively into a ReplicationRecoveryServicesProviderId
// note: this method should only be used for API response data and not user input
func ParseReplicationRecoveryServicesProviderIDInsensitively(input string) (*ReplicationRecoveryServicesProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationRecoveryServicesProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationRecoveryServicesProviderId{}

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

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationRecoveryServicesProviderID checks that 'input' can be parsed as a Replication Recovery Services Provider ID
func ValidateReplicationRecoveryServicesProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationRecoveryServicesProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Recovery Services Provider ID
func (id ReplicationRecoveryServicesProviderId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationRecoveryServicesProviders/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName, id.FabricName, id.ProviderName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Recovery Services Provider ID
func (id ReplicationRecoveryServicesProviderId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReplicationRecoveryServicesProviders", "replicationRecoveryServicesProviders", "replicationRecoveryServicesProviders"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
	}
}

// String returns a human-readable description of this Replication Recovery Services Provider ID
func (id ReplicationRecoveryServicesProviderId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Fabric Name: %q", id.FabricName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
	}
	return fmt.Sprintf("Replication Recovery Services Provider (%s)", strings.Join(components, "\n"))
}
