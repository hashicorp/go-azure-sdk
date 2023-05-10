package sourcecontrolconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SourceControlConfigurationId{}

// SourceControlConfigurationId is a struct representing the Resource ID for a Source Control Configuration
type SourceControlConfigurationId struct {
	SubscriptionId                 string
	ResourceGroupName              string
	ProviderName                   string
	ClusterResourceName            string
	ClusterName                    string
	SourceControlConfigurationName string
}

// NewSourceControlConfigurationID returns a new SourceControlConfigurationId struct
func NewSourceControlConfigurationID(subscriptionId string, resourceGroupName string, providerName string, clusterResourceName string, clusterName string, sourceControlConfigurationName string) SourceControlConfigurationId {
	return SourceControlConfigurationId{
		SubscriptionId:                 subscriptionId,
		ResourceGroupName:              resourceGroupName,
		ProviderName:                   providerName,
		ClusterResourceName:            clusterResourceName,
		ClusterName:                    clusterName,
		SourceControlConfigurationName: sourceControlConfigurationName,
	}
}

// ParseSourceControlConfigurationID parses 'input' into a SourceControlConfigurationId
func ParseSourceControlConfigurationID(input string) (*SourceControlConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(SourceControlConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SourceControlConfigurationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ClusterResourceName, ok = parsed.Parsed["clusterResourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterResourceName", *parsed)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterName", *parsed)
	}

	if id.SourceControlConfigurationName, ok = parsed.Parsed["sourceControlConfigurationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sourceControlConfigurationName", *parsed)
	}

	return &id, nil
}

// ParseSourceControlConfigurationIDInsensitively parses 'input' case-insensitively into a SourceControlConfigurationId
// note: this method should only be used for API response data and not user input
func ParseSourceControlConfigurationIDInsensitively(input string) (*SourceControlConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(SourceControlConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SourceControlConfigurationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ClusterResourceName, ok = parsed.Parsed["clusterResourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterResourceName", *parsed)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterName", *parsed)
	}

	if id.SourceControlConfigurationName, ok = parsed.Parsed["sourceControlConfigurationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sourceControlConfigurationName", *parsed)
	}

	return &id, nil
}

// ValidateSourceControlConfigurationID checks that 'input' can be parsed as a Source Control Configuration ID
func ValidateSourceControlConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSourceControlConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Source Control Configuration ID
func (id SourceControlConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ClusterResourceName, id.ClusterName, id.SourceControlConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Source Control Configuration ID
func (id SourceControlConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("clusterResourceName", "clusterResourceValue"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftKubernetesConfiguration", "Microsoft.KubernetesConfiguration", "Microsoft.KubernetesConfiguration"),
		resourceids.StaticSegment("staticSourceControlConfigurations", "sourceControlConfigurations", "sourceControlConfigurations"),
		resourceids.UserSpecifiedSegment("sourceControlConfigurationName", "sourceControlConfigurationValue"),
	}
}

// String returns a human-readable description of this Source Control Configuration ID
func (id SourceControlConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Cluster Resource Name: %q", id.ClusterResourceName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Source Control Configuration Name: %q", id.SourceControlConfigurationName),
	}
	return fmt.Sprintf("Source Control Configuration (%s)", strings.Join(components, "\n"))
}
