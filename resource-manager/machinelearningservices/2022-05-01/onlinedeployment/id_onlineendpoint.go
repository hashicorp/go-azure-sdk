package onlinedeployment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = OnlineEndpointId{}

// OnlineEndpointId is a struct representing the Resource ID for a Online Endpoint
type OnlineEndpointId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	EndpointName      string
}

// NewOnlineEndpointID returns a new OnlineEndpointId struct
func NewOnlineEndpointID(subscriptionId string, resourceGroupName string, workspaceName string, endpointName string) OnlineEndpointId {
	return OnlineEndpointId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		EndpointName:      endpointName,
	}
}

// ParseOnlineEndpointID parses 'input' into a OnlineEndpointId
func ParseOnlineEndpointID(input string) (*OnlineEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(OnlineEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OnlineEndpointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EndpointName, ok = parsed.Parsed["endpointName"]; !ok {
		return nil, fmt.Errorf("the segment 'endpointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseOnlineEndpointIDInsensitively parses 'input' case-insensitively into a OnlineEndpointId
// note: this method should only be used for API response data and not user input
func ParseOnlineEndpointIDInsensitively(input string) (*OnlineEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(OnlineEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OnlineEndpointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EndpointName, ok = parsed.Parsed["endpointName"]; !ok {
		return nil, fmt.Errorf("the segment 'endpointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateOnlineEndpointID checks that 'input' can be parsed as a Online Endpoint ID
func ValidateOnlineEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOnlineEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Online Endpoint ID
func (id OnlineEndpointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/onlineEndpoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.EndpointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Online Endpoint ID
func (id OnlineEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticOnlineEndpoints", "onlineEndpoints", "onlineEndpoints"),
		resourceids.UserSpecifiedSegment("endpointName", "endpointValue"),
	}
}

// String returns a human-readable description of this Online Endpoint ID
func (id OnlineEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Endpoint Name: %q", id.EndpointName),
	}
	return fmt.Sprintf("Online Endpoint (%s)", strings.Join(components, "\n"))
}
