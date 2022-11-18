package migrationrecoverypoints

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = MigrationRecoveryPointId{}

// MigrationRecoveryPointId is a struct representing the Resource ID for a Migration Recovery Point
type MigrationRecoveryPointId struct {
	SubscriptionId             string
	ResourceGroupName          string
	ResourceName               string
	FabricName                 string
	ProtectionContainerName    string
	MigrationItemName          string
	MigrationRecoveryPointName string
}

// NewMigrationRecoveryPointID returns a new MigrationRecoveryPointId struct
func NewMigrationRecoveryPointID(subscriptionId string, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, migrationItemName string, migrationRecoveryPointName string) MigrationRecoveryPointId {
	return MigrationRecoveryPointId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		ResourceName:               resourceName,
		FabricName:                 fabricName,
		ProtectionContainerName:    protectionContainerName,
		MigrationItemName:          migrationItemName,
		MigrationRecoveryPointName: migrationRecoveryPointName,
	}
}

// ParseMigrationRecoveryPointID parses 'input' into a MigrationRecoveryPointId
func ParseMigrationRecoveryPointID(input string) (*MigrationRecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(MigrationRecoveryPointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MigrationRecoveryPointId{}

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

	if id.MigrationRecoveryPointName, ok = parsed.Parsed["migrationRecoveryPointName"]; !ok {
		return nil, fmt.Errorf("the segment 'migrationRecoveryPointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseMigrationRecoveryPointIDInsensitively parses 'input' case-insensitively into a MigrationRecoveryPointId
// note: this method should only be used for API response data and not user input
func ParseMigrationRecoveryPointIDInsensitively(input string) (*MigrationRecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(MigrationRecoveryPointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MigrationRecoveryPointId{}

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

	if id.MigrationRecoveryPointName, ok = parsed.Parsed["migrationRecoveryPointName"]; !ok {
		return nil, fmt.Errorf("the segment 'migrationRecoveryPointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateMigrationRecoveryPointID checks that 'input' can be parsed as a Migration Recovery Point ID
func ValidateMigrationRecoveryPointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMigrationRecoveryPointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Migration Recovery Point ID
func (id MigrationRecoveryPointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationProtectionContainers/%s/replicationMigrationItems/%s/migrationRecoveryPoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName, id.FabricName, id.ProtectionContainerName, id.MigrationItemName, id.MigrationRecoveryPointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Migration Recovery Point ID
func (id MigrationRecoveryPointId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticMigrationRecoveryPoints", "migrationRecoveryPoints", "migrationRecoveryPoints"),
		resourceids.UserSpecifiedSegment("migrationRecoveryPointName", "migrationRecoveryPointValue"),
	}
}

// String returns a human-readable description of this Migration Recovery Point ID
func (id MigrationRecoveryPointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Fabric Name: %q", id.FabricName),
		fmt.Sprintf("Protection Container Name: %q", id.ProtectionContainerName),
		fmt.Sprintf("Migration Item Name: %q", id.MigrationItemName),
		fmt.Sprintf("Migration Recovery Point Name: %q", id.MigrationRecoveryPointName),
	}
	return fmt.Sprintf("Migration Recovery Point (%s)", strings.Join(components, "\n"))
}
