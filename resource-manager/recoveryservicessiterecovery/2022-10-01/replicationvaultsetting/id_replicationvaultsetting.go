package replicationvaultsetting

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationVaultSettingId{}

// ReplicationVaultSettingId is a struct representing the Resource ID for a Replication Vault Setting
type ReplicationVaultSettingId struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourceName      string
	VaultSettingName  string
}

// NewReplicationVaultSettingID returns a new ReplicationVaultSettingId struct
func NewReplicationVaultSettingID(subscriptionId string, resourceGroupName string, resourceName string, vaultSettingName string) ReplicationVaultSettingId {
	return ReplicationVaultSettingId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ResourceName:      resourceName,
		VaultSettingName:  vaultSettingName,
	}
}

// ParseReplicationVaultSettingID parses 'input' into a ReplicationVaultSettingId
func ParseReplicationVaultSettingID(input string) (*ReplicationVaultSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationVaultSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationVaultSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.VaultSettingName, ok = parsed.Parsed["vaultSettingName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultSettingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationVaultSettingIDInsensitively parses 'input' case-insensitively into a ReplicationVaultSettingId
// note: this method should only be used for API response data and not user input
func ParseReplicationVaultSettingIDInsensitively(input string) (*ReplicationVaultSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationVaultSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationVaultSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.VaultSettingName, ok = parsed.Parsed["vaultSettingName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultSettingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationVaultSettingID checks that 'input' can be parsed as a Replication Vault Setting ID
func ValidateReplicationVaultSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationVaultSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Vault Setting ID
func (id ReplicationVaultSettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationVaultSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName, id.VaultSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Vault Setting ID
func (id ReplicationVaultSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticReplicationVaultSettings", "replicationVaultSettings", "replicationVaultSettings"),
		resourceids.UserSpecifiedSegment("vaultSettingName", "vaultSettingValue"),
	}
}

// String returns a human-readable description of this Replication Vault Setting ID
func (id ReplicationVaultSettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Vault Setting Name: %q", id.VaultSettingName),
	}
	return fmt.Sprintf("Replication Vault Setting (%s)", strings.Join(components, "\n"))
}
