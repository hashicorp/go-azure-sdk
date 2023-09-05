package configurationprofilesversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ConfigurationProfileVersionId{}

// ConfigurationProfileVersionId is a struct representing the Resource ID for a Configuration Profile Version
type ConfigurationProfileVersionId struct {
	SubscriptionId           string
	ResourceGroupName        string
	ConfigurationProfileName string
	VersionName              string
}

// NewConfigurationProfileVersionID returns a new ConfigurationProfileVersionId struct
func NewConfigurationProfileVersionID(subscriptionId string, resourceGroupName string, configurationProfileName string, versionName string) ConfigurationProfileVersionId {
	return ConfigurationProfileVersionId{
		SubscriptionId:           subscriptionId,
		ResourceGroupName:        resourceGroupName,
		ConfigurationProfileName: configurationProfileName,
		VersionName:              versionName,
	}
}

// ParseConfigurationProfileVersionID parses 'input' into a ConfigurationProfileVersionId
func ParseConfigurationProfileVersionID(input string) (*ConfigurationProfileVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigurationProfileVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigurationProfileVersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ConfigurationProfileName, ok = parsed.Parsed["configurationProfileName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileName", *parsed)
	}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionName", *parsed)
	}

	return &id, nil
}

// ParseConfigurationProfileVersionIDInsensitively parses 'input' case-insensitively into a ConfigurationProfileVersionId
// note: this method should only be used for API response data and not user input
func ParseConfigurationProfileVersionIDInsensitively(input string) (*ConfigurationProfileVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigurationProfileVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigurationProfileVersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ConfigurationProfileName, ok = parsed.Parsed["configurationProfileName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileName", *parsed)
	}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionName", *parsed)
	}

	return &id, nil
}

// ValidateConfigurationProfileVersionID checks that 'input' can be parsed as a Configuration Profile Version ID
func ValidateConfigurationProfileVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConfigurationProfileVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Configuration Profile Version ID
func (id ConfigurationProfileVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AutoManage/configurationProfiles/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ConfigurationProfileName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Configuration Profile Version ID
func (id ConfigurationProfileVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutoManage", "Microsoft.AutoManage", "Microsoft.AutoManage"),
		resourceids.StaticSegment("staticConfigurationProfiles", "configurationProfiles", "configurationProfiles"),
		resourceids.UserSpecifiedSegment("configurationProfileName", "configurationProfileValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionValue"),
	}
}

// String returns a human-readable description of this Configuration Profile Version ID
func (id ConfigurationProfileVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Configuration Profile Name: %q", id.ConfigurationProfileName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Configuration Profile Version (%s)", strings.Join(components, "\n"))
}
