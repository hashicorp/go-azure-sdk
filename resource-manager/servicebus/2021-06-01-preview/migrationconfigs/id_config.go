package migrationconfigs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ConfigId{}

// ConfigId is a struct representing the Resource ID for a Config
type ConfigId struct {
	SubscriptionId    string
	ResourceGroupName string
	NamespaceName     string
	ConfigName        MigrationConfigurationName
}

// NewConfigID returns a new ConfigId struct
func NewConfigID(subscriptionId string, resourceGroupName string, namespaceName string, configName MigrationConfigurationName) ConfigId {
	return ConfigId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NamespaceName:     namespaceName,
		ConfigName:        configName,
	}
}

// ParseConfigID parses 'input' into a ConfigId
func ParseConfigID(input string) (*ConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.NamespaceName, ok = parsed.Parsed["namespaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'namespaceName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["configName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'configName' was not found in the resource id %q", input)
		}

		configName, err := parseMigrationConfigurationName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ConfigName = *configName
	}

	return &id, nil
}

// ParseConfigIDInsensitively parses 'input' case-insensitively into a ConfigId
// note: this method should only be used for API response data and not user input
func ParseConfigIDInsensitively(input string) (*ConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.NamespaceName, ok = parsed.Parsed["namespaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'namespaceName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["configName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'configName' was not found in the resource id %q", input)
		}

		configName, err := parseMigrationConfigurationName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ConfigName = *configName
	}

	return &id, nil
}

// ValidateConfigID checks that 'input' can be parsed as a Config ID
func ValidateConfigID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConfigID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Config ID
func (id ConfigId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ServiceBus/namespaces/%s/migrationConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NamespaceName, string(id.ConfigName))
}

// Segments returns a slice of Resource ID Segments which comprise this Config ID
func (id ConfigId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceBus", "Microsoft.ServiceBus", "Microsoft.ServiceBus"),
		resourceids.StaticSegment("staticNamespaces", "namespaces", "namespaces"),
		resourceids.UserSpecifiedSegment("namespaceName", "namespaceValue"),
		resourceids.StaticSegment("staticMigrationConfigurations", "migrationConfigurations", "migrationConfigurations"),
		resourceids.ConstantSegment("configName", PossibleValuesForMigrationConfigurationName(), "$default"),
	}
}

// String returns a human-readable description of this Config ID
func (id ConfigId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Namespace Name: %q", id.NamespaceName),
		fmt.Sprintf("Config Name: %q", string(id.ConfigName)),
	}
	return fmt.Sprintf("Config (%s)", strings.Join(components, "\n"))
}
