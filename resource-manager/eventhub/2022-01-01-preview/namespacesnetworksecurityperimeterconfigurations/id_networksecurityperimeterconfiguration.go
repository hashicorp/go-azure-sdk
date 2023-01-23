package namespacesnetworksecurityperimeterconfigurations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = NetworkSecurityPerimeterConfigurationId{}

// NetworkSecurityPerimeterConfigurationId is a struct representing the Resource ID for a Network Security Perimeter Configuration
type NetworkSecurityPerimeterConfigurationId struct {
	SubscriptionId                            string
	ResourceGroupName                         string
	NamespaceName                             string
	NetworkSecurityPerimeterConfigurationName string
}

// NewNetworkSecurityPerimeterConfigurationID returns a new NetworkSecurityPerimeterConfigurationId struct
func NewNetworkSecurityPerimeterConfigurationID(subscriptionId string, resourceGroupName string, namespaceName string, networkSecurityPerimeterConfigurationName string) NetworkSecurityPerimeterConfigurationId {
	return NetworkSecurityPerimeterConfigurationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NamespaceName:     namespaceName,
		NetworkSecurityPerimeterConfigurationName: networkSecurityPerimeterConfigurationName,
	}
}

// ParseNetworkSecurityPerimeterConfigurationID parses 'input' into a NetworkSecurityPerimeterConfigurationId
func ParseNetworkSecurityPerimeterConfigurationID(input string) (*NetworkSecurityPerimeterConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(NetworkSecurityPerimeterConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NetworkSecurityPerimeterConfigurationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.NamespaceName, ok = parsed.Parsed["namespaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'namespaceName' was not found in the resource id %q", input)
	}

	if id.NetworkSecurityPerimeterConfigurationName, ok = parsed.Parsed["networkSecurityPerimeterConfigurationName"]; !ok {
		return nil, fmt.Errorf("the segment 'networkSecurityPerimeterConfigurationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseNetworkSecurityPerimeterConfigurationIDInsensitively parses 'input' case-insensitively into a NetworkSecurityPerimeterConfigurationId
// note: this method should only be used for API response data and not user input
func ParseNetworkSecurityPerimeterConfigurationIDInsensitively(input string) (*NetworkSecurityPerimeterConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(NetworkSecurityPerimeterConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NetworkSecurityPerimeterConfigurationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.NamespaceName, ok = parsed.Parsed["namespaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'namespaceName' was not found in the resource id %q", input)
	}

	if id.NetworkSecurityPerimeterConfigurationName, ok = parsed.Parsed["networkSecurityPerimeterConfigurationName"]; !ok {
		return nil, fmt.Errorf("the segment 'networkSecurityPerimeterConfigurationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateNetworkSecurityPerimeterConfigurationID checks that 'input' can be parsed as a Network Security Perimeter Configuration ID
func ValidateNetworkSecurityPerimeterConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNetworkSecurityPerimeterConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Network Security Perimeter Configuration ID
func (id NetworkSecurityPerimeterConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.EventHub/namespaces/%s/networkSecurityPerimeterConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NamespaceName, id.NetworkSecurityPerimeterConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Network Security Perimeter Configuration ID
func (id NetworkSecurityPerimeterConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftEventHub", "Microsoft.EventHub", "Microsoft.EventHub"),
		resourceids.StaticSegment("staticNamespaces", "namespaces", "namespaces"),
		resourceids.UserSpecifiedSegment("namespaceName", "namespaceValue"),
		resourceids.StaticSegment("staticNetworkSecurityPerimeterConfigurations", "networkSecurityPerimeterConfigurations", "networkSecurityPerimeterConfigurations"),
		resourceids.UserSpecifiedSegment("networkSecurityPerimeterConfigurationName", "networkSecurityPerimeterConfigurationValue"),
	}
}

// String returns a human-readable description of this Network Security Perimeter Configuration ID
func (id NetworkSecurityPerimeterConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Namespace Name: %q", id.NamespaceName),
		fmt.Sprintf("Network Security Perimeter Configuration Name: %q", id.NetworkSecurityPerimeterConfigurationName),
	}
	return fmt.Sprintf("Network Security Perimeter Configuration (%s)", strings.Join(components, "\n"))
}
