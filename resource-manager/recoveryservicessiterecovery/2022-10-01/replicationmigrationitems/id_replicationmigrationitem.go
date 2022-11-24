package replicationmigrationitems

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationMigrationItemId{}

// ReplicationMigrationItemId is a struct representing the Resource ID for a Replication Migration Item
type ReplicationMigrationItemId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ResourceName            string
	FabricName              string
	ProtectionContainerName string
	MigrationItemName       string
}

// NewReplicationMigrationItemID returns a new ReplicationMigrationItemId struct
func NewReplicationMigrationItemID(subscriptionId string, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, migrationItemName string) ReplicationMigrationItemId {
	return ReplicationMigrationItemId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ResourceName:            resourceName,
		FabricName:              fabricName,
		ProtectionContainerName: protectionContainerName,
		MigrationItemName:       migrationItemName,
	}
}

// ParseReplicationMigrationItemID parses 'input' into a ReplicationMigrationItemId
func ParseReplicationMigrationItemID(input string) (*ReplicationMigrationItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationMigrationItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationMigrationItemId{}

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

	if id.ProtectionContainerName, ok = parsed.Parsed["protectionContainerName"]; !ok {
		return nil, fmt.Errorf("the segment 'protectionContainerName' was not found in the resource id %q", input)
	}

	if id.MigrationItemName, ok = parsed.Parsed["migrationItemName"]; !ok {
		return nil, fmt.Errorf("the segment 'migrationItemName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationMigrationItemIDInsensitively parses 'input' case-insensitively into a ReplicationMigrationItemId
// note: this method should only be used for API response data and not user input
func ParseReplicationMigrationItemIDInsensitively(input string) (*ReplicationMigrationItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationMigrationItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationMigrationItemId{}

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

	if id.ProtectionContainerName, ok = parsed.Parsed["protectionContainerName"]; !ok {
		return nil, fmt.Errorf("the segment 'protectionContainerName' was not found in the resource id %q", input)
	}

	if id.MigrationItemName, ok = parsed.Parsed["migrationItemName"]; !ok {
		return nil, fmt.Errorf("the segment 'migrationItemName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationMigrationItemID checks that 'input' can be parsed as a Replication Migration Item ID
func ValidateReplicationMigrationItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationMigrationItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Migration Item ID
func (id ReplicationMigrationItemId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationProtectionContainers/%s/replicationMigrationItems/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName, id.FabricName, id.ProtectionContainerName, id.MigrationItemName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Migration Item ID
func (id ReplicationMigrationItemId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReplicationProtectionContainers", "replicationProtectionContainers", "replicationProtectionContainers"),
		resourceids.UserSpecifiedSegment("protectionContainerName", "protectionContainerValue"),
		resourceids.StaticSegment("staticReplicationMigrationItems", "replicationMigrationItems", "replicationMigrationItems"),
		resourceids.UserSpecifiedSegment("migrationItemName", "migrationItemValue"),
	}
}

// String returns a human-readable description of this Replication Migration Item ID
func (id ReplicationMigrationItemId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Fabric Name: %q", id.FabricName),
		fmt.Sprintf("Protection Container Name: %q", id.ProtectionContainerName),
		fmt.Sprintf("Migration Item Name: %q", id.MigrationItemName),
	}
	return fmt.Sprintf("Replication Migration Item (%s)", strings.Join(components, "\n"))
}
