package appplatform

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ResultId{}

// ResultId is a struct representing the Resource ID for a Result
type ResultId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	BuildServiceName  string
	BuildName         string
	BuildResultName   string
}

// NewResultID returns a new ResultId struct
func NewResultID(subscriptionId string, resourceGroupName string, serviceName string, buildServiceName string, buildName string, buildResultName string) ResultId {
	return ResultId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		BuildServiceName:  buildServiceName,
		BuildName:         buildName,
		BuildResultName:   buildResultName,
	}
}

// ParseResultID parses 'input' into a ResultId
func ParseResultID(input string) (*ResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ResultId{}

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

	if id.BuildName, ok = parsed.Parsed["buildName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildName' was not found in the resource id %q", input)
	}

	if id.BuildResultName, ok = parsed.Parsed["buildResultName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildResultName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseResultIDInsensitively parses 'input' case-insensitively into a ResultId
// note: this method should only be used for API response data and not user input
func ParseResultIDInsensitively(input string) (*ResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ResultId{}

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

	if id.BuildName, ok = parsed.Parsed["buildName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildName' was not found in the resource id %q", input)
	}

	if id.BuildResultName, ok = parsed.Parsed["buildResultName"]; !ok {
		return nil, fmt.Errorf("the segment 'buildResultName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateResultID checks that 'input' can be parsed as a Result ID
func ValidateResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Result ID
func (id ResultId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AppPlatform/spring/%s/buildServices/%s/builds/%s/results/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.BuildServiceName, id.BuildName, id.BuildResultName)
}

// Segments returns a slice of Resource ID Segments which comprise this Result ID
func (id ResultId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticBuilds", "builds", "builds"),
		resourceids.UserSpecifiedSegment("buildName", "buildValue"),
		resourceids.StaticSegment("staticResults", "results", "results"),
		resourceids.UserSpecifiedSegment("buildResultName", "buildResultValue"),
	}
}

// String returns a human-readable description of this Result ID
func (id ResultId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Build Service Name: %q", id.BuildServiceName),
		fmt.Sprintf("Build Name: %q", id.BuildName),
		fmt.Sprintf("Build Result Name: %q", id.BuildResultName),
	}
	return fmt.Sprintf("Result (%s)", strings.Join(components, "\n"))
}
