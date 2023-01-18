package integrationserviceenvironmentmanagedapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ManagedApiId{}

// ManagedApiId is a struct representing the Resource ID for a Managed Api
type ManagedApiId struct {
	SubscriptionId                    string
	ResourceGroup                     string
	IntegrationServiceEnvironmentName string
	ApiName                           string
}

// NewManagedApiID returns a new ManagedApiId struct
func NewManagedApiID(subscriptionId string, resourceGroup string, integrationServiceEnvironmentName string, apiName string) ManagedApiId {
	return ManagedApiId{
		SubscriptionId:                    subscriptionId,
		ResourceGroup:                     resourceGroup,
		IntegrationServiceEnvironmentName: integrationServiceEnvironmentName,
		ApiName:                           apiName,
	}
}

// ParseManagedApiID parses 'input' into a ManagedApiId
func ParseManagedApiID(input string) (*ManagedApiId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedApiId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedApiId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroup, ok = parsed.Parsed["resourceGroup"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroup' was not found in the resource id %q", input)
	}

	if id.IntegrationServiceEnvironmentName, ok = parsed.Parsed["integrationServiceEnvironmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'integrationServiceEnvironmentName' was not found in the resource id %q", input)
	}

	if id.ApiName, ok = parsed.Parsed["apiName"]; !ok {
		return nil, fmt.Errorf("the segment 'apiName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseManagedApiIDInsensitively parses 'input' case-insensitively into a ManagedApiId
// note: this method should only be used for API response data and not user input
func ParseManagedApiIDInsensitively(input string) (*ManagedApiId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedApiId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedApiId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroup, ok = parsed.Parsed["resourceGroup"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroup' was not found in the resource id %q", input)
	}

	if id.IntegrationServiceEnvironmentName, ok = parsed.Parsed["integrationServiceEnvironmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'integrationServiceEnvironmentName' was not found in the resource id %q", input)
	}

	if id.ApiName, ok = parsed.Parsed["apiName"]; !ok {
		return nil, fmt.Errorf("the segment 'apiName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateManagedApiID checks that 'input' can be parsed as a Managed Api ID
func ValidateManagedApiID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedApiID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Api ID
func (id ManagedApiId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Logic/integrationServiceEnvironments/%s/managedApis/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.IntegrationServiceEnvironmentName, id.ApiName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Api ID
func (id ManagedApiId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroup", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftLogic", "Microsoft.Logic", "Microsoft.Logic"),
		resourceids.StaticSegment("staticIntegrationServiceEnvironments", "integrationServiceEnvironments", "integrationServiceEnvironments"),
		resourceids.UserSpecifiedSegment("integrationServiceEnvironmentName", "integrationServiceEnvironmentValue"),
		resourceids.StaticSegment("staticManagedApis", "managedApis", "managedApis"),
		resourceids.UserSpecifiedSegment("apiName", "apiValue"),
	}
}

// String returns a human-readable description of this Managed Api ID
func (id ManagedApiId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group: %q", id.ResourceGroup),
		fmt.Sprintf("Integration Service Environment Name: %q", id.IntegrationServiceEnvironmentName),
		fmt.Sprintf("Api Name: %q", id.ApiName),
	}
	return fmt.Sprintf("Managed Api (%s)", strings.Join(components, "\n"))
}
