package environmentversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = EnvironmentVersionId{}

// EnvironmentVersionId is a struct representing the Resource ID for a Environment Version
type EnvironmentVersionId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	Name              string
	Version           string
}

// NewEnvironmentVersionID returns a new EnvironmentVersionId struct
func NewEnvironmentVersionID(subscriptionId string, resourceGroupName string, workspaceName string, name string, version string) EnvironmentVersionId {
	return EnvironmentVersionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		Name:              name,
		Version:           version,
	}
}

// ParseEnvironmentVersionID parses 'input' into a EnvironmentVersionId
func ParseEnvironmentVersionID(input string) (*EnvironmentVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(EnvironmentVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EnvironmentVersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.Name, ok = parsed.Parsed["name"]; !ok {
		return nil, fmt.Errorf("the segment 'name' was not found in the resource id %q", input)
	}

	if id.Version, ok = parsed.Parsed["version"]; !ok {
		return nil, fmt.Errorf("the segment 'version' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseEnvironmentVersionIDInsensitively parses 'input' case-insensitively into a EnvironmentVersionId
// note: this method should only be used for API response data and not user input
func ParseEnvironmentVersionIDInsensitively(input string) (*EnvironmentVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(EnvironmentVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EnvironmentVersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.Name, ok = parsed.Parsed["name"]; !ok {
		return nil, fmt.Errorf("the segment 'name' was not found in the resource id %q", input)
	}

	if id.Version, ok = parsed.Parsed["version"]; !ok {
		return nil, fmt.Errorf("the segment 'version' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateEnvironmentVersionID checks that 'input' can be parsed as a Environment Version ID
func ValidateEnvironmentVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEnvironmentVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Environment Version ID
func (id EnvironmentVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/environments/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.Name, id.Version)
}

// Segments returns a slice of Resource ID Segments which comprise this Environment Version ID
func (id EnvironmentVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticEnvironments", "environments", "environments"),
		resourceids.UserSpecifiedSegment("name", "nameValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("version", "versionValue"),
	}
}

// String returns a human-readable description of this Environment Version ID
func (id EnvironmentVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Name: %q", id.Name),
		fmt.Sprintf("Version: %q", id.Version),
	}
	return fmt.Sprintf("Environment Version (%s)", strings.Join(components, "\n"))
}
