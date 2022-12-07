package appplatform

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SupportedBuildPackId{}

// SupportedBuildPackId is a struct representing the Resource ID for a Supported Build Pack
type SupportedBuildPackId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	BuildServiceName  string
	BuildpackName     string
}

// NewSupportedBuildPackID returns a new SupportedBuildPackId struct
func NewSupportedBuildPackID(subscriptionId string, resourceGroupName string, serviceName string, buildServiceName string, buildpackName string) SupportedBuildPackId {
	return SupportedBuildPackId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		BuildServiceName:  buildServiceName,
		BuildpackName:     buildpackName,
	}
}

// ParseSupportedBuildPackID parses 'input' into a SupportedBuildPackId
func ParseSupportedBuildPackID(input string) (*SupportedBuildPackId, error) {
	parser := resourceids.NewParserFromResourceIdType(SupportedBuildPackId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SupportedBuildPackId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.BuildServiceName, ok = parsed.Parsed["buildServiceName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildServiceName' was not found in the resource id %q", input)
	}

	if id.BuildpackName, ok = parsed.Parsed["buildpackName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildpackName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseSupportedBuildPackIDInsensitively parses 'input' case-insensitively into a SupportedBuildPackId
// note: this method should only be used for API response data and not user input
func ParseSupportedBuildPackIDInsensitively(input string) (*SupportedBuildPackId, error) {
	parser := resourceids.NewParserFromResourceIdType(SupportedBuildPackId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SupportedBuildPackId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.BuildServiceName, ok = parsed.Parsed["buildServiceName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildServiceName' was not found in the resource id %q", input)
	}

	if id.BuildpackName, ok = parsed.Parsed["buildpackName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildpackName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateSupportedBuildPackID checks that 'input' can be parsed as a Supported Build Pack ID
func ValidateSupportedBuildPackID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSupportedBuildPackID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Supported Build Pack ID
func (id SupportedBuildPackId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AppPlatform/spring/%s/buildServices/%s/supportedBuildPacks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.BuildServiceName, id.BuildpackName)
}

// Segments returns a slice of Resource ID Segments which comprise this Supported Build Pack ID
func (id SupportedBuildPackId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAppPlatform", "Microsoft.AppPlatform", "Microsoft.AppPlatform"),
		resourceids.StaticSegment("staticSpring", "spring", "spring"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticBuildServices", "buildServices", "buildServices"),
		resourceids.UserSpecifiedSegment("buildServiceName", "buildServiceValue"),
		resourceids.StaticSegment("staticSupportedBuildPacks", "supportedBuildPacks", "supportedBuildPacks"),
		resourceids.UserSpecifiedSegment("buildpackName", "buildpackValue"),
	}
}

// String returns a human-readable description of this Supported Build Pack ID
func (id SupportedBuildPackId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Build Service Name: %q", id.BuildServiceName),
		fmt.Sprintf("Buildpack Name: %q", id.BuildpackName),
	}
	return fmt.Sprintf("Supported Build Pack (%s)", strings.Join(components, "\n"))
}
