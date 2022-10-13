package appplatform

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SupportedStackId{}

// SupportedStackId is a struct representing the Resource ID for a Supported Stack
type SupportedStackId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	BuildServiceName  string
	StackName         string
}

// NewSupportedStackID returns a new SupportedStackId struct
func NewSupportedStackID(subscriptionId string, resourceGroupName string, serviceName string, buildServiceName string, stackName string) SupportedStackId {
	return SupportedStackId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		BuildServiceName:  buildServiceName,
		StackName:         stackName,
	}
}

// ParseSupportedStackID parses 'input' into a SupportedStackId
func ParseSupportedStackID(input string) (*SupportedStackId, error) {
	parser := resourceids.NewParserFromResourceIdType(SupportedStackId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SupportedStackId{}

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

	if id.StackName, ok = parsed.Parsed["stackName"]; !ok {
		return nil, fmt.Errorf("the segment 'stackName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseSupportedStackIDInsensitively parses 'input' case-insensitively into a SupportedStackId
// note: this method should only be used for API response data and not user input
func ParseSupportedStackIDInsensitively(input string) (*SupportedStackId, error) {
	parser := resourceids.NewParserFromResourceIdType(SupportedStackId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SupportedStackId{}

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

	if id.StackName, ok = parsed.Parsed["stackName"]; !ok {
		return nil, fmt.Errorf("the segment 'stackName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateSupportedStackID checks that 'input' can be parsed as a Supported Stack ID
func ValidateSupportedStackID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSupportedStackID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Supported Stack ID
func (id SupportedStackId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AppPlatform/spring/%s/buildServices/%s/supportedStacks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.BuildServiceName, id.StackName)
}

// Segments returns a slice of Resource ID Segments which comprise this Supported Stack ID
func (id SupportedStackId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticSupportedStacks", "supportedStacks", "supportedStacks"),
		resourceids.UserSpecifiedSegment("stackName", "stackValue"),
	}
}

// String returns a human-readable description of this Supported Stack ID
func (id SupportedStackId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Build Service Name: %q", id.BuildServiceName),
		fmt.Sprintf("Stack Name: %q", id.StackName),
	}
	return fmt.Sprintf("Supported Stack (%s)", strings.Join(components, "\n"))
}
