package appplatform

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BuildpackBindingId{}

// BuildpackBindingId is a struct representing the Resource ID for a Buildpack Binding
type BuildpackBindingId struct {
	SubscriptionId       string
	ResourceGroupName    string
	ServiceName          string
	BuildServiceName     string
	BuilderName          string
	BuildpackBindingName string
}

// NewBuildpackBindingID returns a new BuildpackBindingId struct
func NewBuildpackBindingID(subscriptionId string, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string) BuildpackBindingId {
	return BuildpackBindingId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		ServiceName:          serviceName,
		BuildServiceName:     buildServiceName,
		BuilderName:          builderName,
		BuildpackBindingName: buildpackBindingName,
	}
}

// ParseBuildpackBindingID parses 'input' into a BuildpackBindingId
func ParseBuildpackBindingID(input string) (*BuildpackBindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(BuildpackBindingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BuildpackBindingId{}

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

	if id.BuilderName, ok = parsed.Parsed["builderName"]; !ok {
		return nil, fmt.Errorf("the segment 'builderName' was not found in the resource id %q", input)
	}

	if id.BuildpackBindingName, ok = parsed.Parsed["buildpackBindingName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildpackBindingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseBuildpackBindingIDInsensitively parses 'input' case-insensitively into a BuildpackBindingId
// note: this method should only be used for API response data and not user input
func ParseBuildpackBindingIDInsensitively(input string) (*BuildpackBindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(BuildpackBindingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BuildpackBindingId{}

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

	if id.BuilderName, ok = parsed.Parsed["builderName"]; !ok {
		return nil, fmt.Errorf("the segment 'builderName' was not found in the resource id %q", input)
	}

	if id.BuildpackBindingName, ok = parsed.Parsed["buildpackBindingName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildpackBindingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateBuildpackBindingID checks that 'input' can be parsed as a Buildpack Binding ID
func ValidateBuildpackBindingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBuildpackBindingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Buildpack Binding ID
func (id BuildpackBindingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AppPlatform/spring/%s/buildServices/%s/builders/%s/buildpackBindings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.BuildServiceName, id.BuilderName, id.BuildpackBindingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Buildpack Binding ID
func (id BuildpackBindingId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticBuilders", "builders", "builders"),
		resourceids.UserSpecifiedSegment("builderName", "builderValue"),
		resourceids.StaticSegment("staticBuildpackBindings", "buildpackBindings", "buildpackBindings"),
		resourceids.UserSpecifiedSegment("buildpackBindingName", "buildpackBindingValue"),
	}
}

// String returns a human-readable description of this Buildpack Binding ID
func (id BuildpackBindingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Build Service Name: %q", id.BuildServiceName),
		fmt.Sprintf("Builder Name: %q", id.BuilderName),
		fmt.Sprintf("Buildpack Binding Name: %q", id.BuildpackBindingName),
	}
	return fmt.Sprintf("Buildpack Binding (%s)", strings.Join(components, "\n"))
}
