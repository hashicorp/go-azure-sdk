package environmentversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RegistryEnvironmentVersionId{})
}

var _ resourceids.ResourceId = &RegistryEnvironmentVersionId{}

// RegistryEnvironmentVersionId is a struct representing the Resource ID for a Registry Environment Version
type RegistryEnvironmentVersionId struct {
	SubscriptionId    string
	ResourceGroupName string
	RegistryName      string
	EnvironmentName   string
	VersionName       string
}

// NewRegistryEnvironmentVersionID returns a new RegistryEnvironmentVersionId struct
func NewRegistryEnvironmentVersionID(subscriptionId string, resourceGroupName string, registryName string, environmentName string, versionName string) RegistryEnvironmentVersionId {
	return RegistryEnvironmentVersionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		RegistryName:      registryName,
		EnvironmentName:   environmentName,
		VersionName:       versionName,
	}
}

// ParseRegistryEnvironmentVersionID parses 'input' into a RegistryEnvironmentVersionId
func ParseRegistryEnvironmentVersionID(input string) (*RegistryEnvironmentVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RegistryEnvironmentVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegistryEnvironmentVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRegistryEnvironmentVersionIDInsensitively parses 'input' case-insensitively into a RegistryEnvironmentVersionId
// note: this method should only be used for API response data and not user input
func ParseRegistryEnvironmentVersionIDInsensitively(input string) (*RegistryEnvironmentVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RegistryEnvironmentVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegistryEnvironmentVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RegistryEnvironmentVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.RegistryName, ok = input.Parsed["registryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "registryName", input)
	}

	if id.EnvironmentName, ok = input.Parsed["environmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "environmentName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateRegistryEnvironmentVersionID checks that 'input' can be parsed as a Registry Environment Version ID
func ValidateRegistryEnvironmentVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRegistryEnvironmentVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Registry Environment Version ID
func (id RegistryEnvironmentVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/registries/%s/environments/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RegistryName, id.EnvironmentName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Registry Environment Version ID
func (id RegistryEnvironmentVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticRegistries", "registries", "registries"),
		resourceids.UserSpecifiedSegment("registryName", "registryName"),
		resourceids.StaticSegment("staticEnvironments", "environments", "environments"),
		resourceids.UserSpecifiedSegment("environmentName", "environmentName"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionName"),
	}
}

// String returns a human-readable description of this Registry Environment Version ID
func (id RegistryEnvironmentVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Registry Name: %q", id.RegistryName),
		fmt.Sprintf("Environment Name: %q", id.EnvironmentName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Registry Environment Version (%s)", strings.Join(components, "\n"))
}
