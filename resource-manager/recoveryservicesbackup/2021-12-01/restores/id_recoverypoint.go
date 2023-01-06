package restores

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = RecoveryPointId{}

// RecoveryPointId is a struct representing the Resource ID for a Recovery Point
type RecoveryPointId struct {
	SubscriptionId    string
	ResourceGroupName string
	VaultName         string
	FabricName        string
	ContainerName     string
	ProtectedItemName string
	RecoveryPointId   string
}

// NewRecoveryPointID returns a new RecoveryPointId struct
func NewRecoveryPointID(subscriptionId string, resourceGroupName string, vaultName string, fabricName string, containerName string, protectedItemName string, recoveryPointId string) RecoveryPointId {
	return RecoveryPointId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		VaultName:         vaultName,
		FabricName:        fabricName,
		ContainerName:     containerName,
		ProtectedItemName: protectedItemName,
		RecoveryPointId:   recoveryPointId,
	}
}

// ParseRecoveryPointID parses 'input' into a RecoveryPointId
func ParseRecoveryPointID(input string) (*RecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecoveryPointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecoveryPointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.FabricName, ok = parsed.Parsed["fabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'fabricName' was not found in the resource id %q", input)
	}

	if id.ContainerName, ok = parsed.Parsed["containerName"]; !ok {
		return nil, fmt.Errorf("the segment 'containerName' was not found in the resource id %q", input)
	}

	if id.ProtectedItemName, ok = parsed.Parsed["protectedItemName"]; !ok {
		return nil, fmt.Errorf("the segment 'protectedItemName' was not found in the resource id %q", input)
	}

	if id.RecoveryPointId, ok = parsed.Parsed["recoveryPointId"]; !ok {
		return nil, fmt.Errorf("the segment 'recoveryPointId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRecoveryPointIDInsensitively parses 'input' case-insensitively into a RecoveryPointId
// note: this method should only be used for API response data and not user input
func ParseRecoveryPointIDInsensitively(input string) (*RecoveryPointId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecoveryPointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecoveryPointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.FabricName, ok = parsed.Parsed["fabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'fabricName' was not found in the resource id %q", input)
	}

	if id.ContainerName, ok = parsed.Parsed["containerName"]; !ok {
		return nil, fmt.Errorf("the segment 'containerName' was not found in the resource id %q", input)
	}

	if id.ProtectedItemName, ok = parsed.Parsed["protectedItemName"]; !ok {
		return nil, fmt.Errorf("the segment 'protectedItemName' was not found in the resource id %q", input)
	}

	if id.RecoveryPointId, ok = parsed.Parsed["recoveryPointId"]; !ok {
		return nil, fmt.Errorf("the segment 'recoveryPointId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRecoveryPointID checks that 'input' can be parsed as a Recovery Point ID
func ValidateRecoveryPointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRecoveryPointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Recovery Point ID
func (id RecoveryPointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/backupFabrics/%s/protectionContainers/%s/protectedItems/%s/recoveryPoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.FabricName, id.ContainerName, id.ProtectedItemName, id.RecoveryPointId)
}

// Segments returns a slice of Resource ID Segments which comprise this Recovery Point ID
func (id RecoveryPointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticBackupFabrics", "backupFabrics", "backupFabrics"),
		resourceids.UserSpecifiedSegment("fabricName", "fabricValue"),
		resourceids.StaticSegment("staticProtectionContainers", "protectionContainers", "protectionContainers"),
		resourceids.UserSpecifiedSegment("containerName", "containerValue"),
		resourceids.StaticSegment("staticProtectedItems", "protectedItems", "protectedItems"),
		resourceids.UserSpecifiedSegment("protectedItemName", "protectedItemValue"),
		resourceids.StaticSegment("staticRecoveryPoints", "recoveryPoints", "recoveryPoints"),
		resourceids.UserSpecifiedSegment("recoveryPointId", "recoveryPointIdValue"),
	}
}

// String returns a human-readable description of this Recovery Point ID
func (id RecoveryPointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Fabric Name: %q", id.FabricName),
		fmt.Sprintf("Container Name: %q", id.ContainerName),
		fmt.Sprintf("Protected Item Name: %q", id.ProtectedItemName),
		fmt.Sprintf("Recovery Point: %q", id.RecoveryPointId),
	}
	return fmt.Sprintf("Recovery Point (%s)", strings.Join(components, "\n"))
}
