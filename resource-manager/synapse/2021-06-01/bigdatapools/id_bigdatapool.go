package bigdatapools

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BigDataPoolId{}

// BigDataPoolId is a struct representing the Resource ID for a Big Data Pool
type BigDataPoolId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	BigDataPoolName   string
}

// NewBigDataPoolID returns a new BigDataPoolId struct
func NewBigDataPoolID(subscriptionId string, resourceGroupName string, workspaceName string, bigDataPoolName string) BigDataPoolId {
	return BigDataPoolId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		BigDataPoolName:   bigDataPoolName,
	}
}

// ParseBigDataPoolID parses 'input' into a BigDataPoolId
func ParseBigDataPoolID(input string) (*BigDataPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(BigDataPoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BigDataPoolId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.BigDataPoolName, ok = parsed.Parsed["bigDataPoolName"]; !ok {
		return nil, fmt.Errorf("the segment 'bigDataPoolName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseBigDataPoolIDInsensitively parses 'input' case-insensitively into a BigDataPoolId
// note: this method should only be used for API response data and not user input
func ParseBigDataPoolIDInsensitively(input string) (*BigDataPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(BigDataPoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BigDataPoolId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.BigDataPoolName, ok = parsed.Parsed["bigDataPoolName"]; !ok {
		return nil, fmt.Errorf("the segment 'bigDataPoolName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateBigDataPoolID checks that 'input' can be parsed as a Big Data Pool ID
func ValidateBigDataPoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBigDataPoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Big Data Pool ID
func (id BigDataPoolId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/bigDataPools/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.BigDataPoolName)
}

// Segments returns a slice of Resource ID Segments which comprise this Big Data Pool ID
func (id BigDataPoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticBigDataPools", "bigDataPools", "bigDataPools"),
		resourceids.UserSpecifiedSegment("bigDataPoolName", "bigDataPoolValue"),
	}
}

// String returns a human-readable description of this Big Data Pool ID
func (id BigDataPoolId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Big Data Pool Name: %q", id.BigDataPoolName),
	}
	return fmt.Sprintf("Big Data Pool (%s)", strings.Join(components, "\n"))
}
