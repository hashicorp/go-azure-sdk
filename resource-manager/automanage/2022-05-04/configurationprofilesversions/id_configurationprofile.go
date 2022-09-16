package configurationprofilesversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ConfigurationProfileId{}

// ConfigurationProfileId is a struct representing the Resource ID for a Configuration Profile
type ConfigurationProfileId struct {
	SubscriptionId           string
	ResourceGroupName        string
	ConfigurationProfileName string
}

// NewConfigurationProfileID returns a new ConfigurationProfileId struct
func NewConfigurationProfileID(subscriptionId string, resourceGroupName string, configurationProfileName string) ConfigurationProfileId {
	return ConfigurationProfileId{
		SubscriptionId:           subscriptionId,
		ResourceGroupName:        resourceGroupName,
		ConfigurationProfileName: configurationProfileName,
	}
}

// ParseConfigurationProfileID parses 'input' into a ConfigurationProfileId
func ParseConfigurationProfileID(input string) (*ConfigurationProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigurationProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigurationProfileId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ConfigurationProfileName, ok = parsed.Parsed["configurationProfileName"]; !ok {
		return nil, fmt.Errorf("the segment 'configurationProfileName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseConfigurationProfileIDInsensitively parses 'input' case-insensitively into a ConfigurationProfileId
// note: this method should only be used for API response data and not user input
func ParseConfigurationProfileIDInsensitively(input string) (*ConfigurationProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigurationProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigurationProfileId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ConfigurationProfileName, ok = parsed.Parsed["configurationProfileName"]; !ok {
		return nil, fmt.Errorf("the segment 'configurationProfileName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateConfigurationProfileID checks that 'input' can be parsed as a Configuration Profile ID
func ValidateConfigurationProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConfigurationProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Configuration Profile ID
func (id ConfigurationProfileId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AutoManage/configurationProfiles/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ConfigurationProfileName)
}

// Segments returns a slice of Resource ID Segments which comprise this Configuration Profile ID
func (id ConfigurationProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutoManage", "Microsoft.AutoManage", "Microsoft.AutoManage"),
		resourceids.StaticSegment("staticConfigurationProfiles", "configurationProfiles", "configurationProfiles"),
		resourceids.UserSpecifiedSegment("configurationProfileName", "configurationProfileValue"),
	}
}

// String returns a human-readable description of this Configuration Profile ID
func (id ConfigurationProfileId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Configuration Profile Name: %q", id.ConfigurationProfileName),
	}
	return fmt.Sprintf("Configuration Profile (%s)", strings.Join(components, "\n"))
}
