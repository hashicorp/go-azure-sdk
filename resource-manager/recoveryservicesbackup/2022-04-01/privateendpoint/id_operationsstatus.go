package privateendpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = OperationsStatusId{}

// OperationsStatusId is a struct representing the Resource ID for a Operations Status
type OperationsStatusId struct {
	SubscriptionId                string
	ResourceGroupName             string
	VaultName                     string
	PrivateEndpointConnectionName string
	OperationId                   string
}

// NewOperationsStatusID returns a new OperationsStatusId struct
func NewOperationsStatusID(subscriptionId string, resourceGroupName string, vaultName string, privateEndpointConnectionName string, operationId string) OperationsStatusId {
	return OperationsStatusId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		VaultName:                     vaultName,
		PrivateEndpointConnectionName: privateEndpointConnectionName,
		OperationId:                   operationId,
	}
}

// ParseOperationsStatusID parses 'input' into a OperationsStatusId
func ParseOperationsStatusID(input string) (*OperationsStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationsStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationsStatusId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.PrivateEndpointConnectionName, ok = parsed.Parsed["privateEndpointConnectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateEndpointConnectionName' was not found in the resource id %q", input)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, fmt.Errorf("the segment 'operationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseOperationsStatusIDInsensitively parses 'input' case-insensitively into a OperationsStatusId
// note: this method should only be used for API response data and not user input
func ParseOperationsStatusIDInsensitively(input string) (*OperationsStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationsStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationsStatusId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.PrivateEndpointConnectionName, ok = parsed.Parsed["privateEndpointConnectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateEndpointConnectionName' was not found in the resource id %q", input)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, fmt.Errorf("the segment 'operationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateOperationsStatusID checks that 'input' can be parsed as a Operations Status ID
func ValidateOperationsStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOperationsStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Operations Status ID
func (id OperationsStatusId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/privateEndpointConnections/%s/operationsStatus/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.PrivateEndpointConnectionName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Operations Status ID
func (id OperationsStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticPrivateEndpointConnections", "privateEndpointConnections", "privateEndpointConnections"),
		resourceids.UserSpecifiedSegment("privateEndpointConnectionName", "privateEndpointConnectionValue"),
		resourceids.StaticSegment("staticOperationsStatus", "operationsStatus", "operationsStatus"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Operations Status ID
func (id OperationsStatusId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Private Endpoint Connection Name: %q", id.PrivateEndpointConnectionName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Operations Status (%s)", strings.Join(components, "\n"))
}
