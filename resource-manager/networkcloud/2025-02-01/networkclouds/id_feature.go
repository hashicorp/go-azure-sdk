package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&FeatureId{})
}

var _ resourceids.ResourceId = &FeatureId{}

// FeatureId is a struct representing the Resource ID for a Feature
type FeatureId struct {
	SubscriptionId        string
	ResourceGroupName     string
	KubernetesClusterName string
	FeatureName           string
}

// NewFeatureID returns a new FeatureId struct
func NewFeatureID(subscriptionId string, resourceGroupName string, kubernetesClusterName string, featureName string) FeatureId {
	return FeatureId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		KubernetesClusterName: kubernetesClusterName,
		FeatureName:           featureName,
	}
}

// ParseFeatureID parses 'input' into a FeatureId
func ParseFeatureID(input string) (*FeatureId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FeatureId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFeatureIDInsensitively parses 'input' case-insensitively into a FeatureId
// note: this method should only be used for API response data and not user input
func ParseFeatureIDInsensitively(input string) (*FeatureId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FeatureId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FeatureId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.KubernetesClusterName, ok = input.Parsed["kubernetesClusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "kubernetesClusterName", input)
	}

	if id.FeatureName, ok = input.Parsed["featureName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "featureName", input)
	}

	return nil
}

// ValidateFeatureID checks that 'input' can be parsed as a Feature ID
func ValidateFeatureID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFeatureID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Feature ID
func (id FeatureId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/kubernetesClusters/%s/features/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.KubernetesClusterName, id.FeatureName)
}

// Segments returns a slice of Resource ID Segments which comprise this Feature ID
func (id FeatureId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticKubernetesClusters", "kubernetesClusters", "kubernetesClusters"),
		resourceids.UserSpecifiedSegment("kubernetesClusterName", "kubernetesClusterName"),
		resourceids.StaticSegment("staticFeatures", "features", "features"),
		resourceids.UserSpecifiedSegment("featureName", "featureName"),
	}
}

// String returns a human-readable description of this Feature ID
func (id FeatureId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Kubernetes Cluster Name: %q", id.KubernetesClusterName),
		fmt.Sprintf("Feature Name: %q", id.FeatureName),
	}
	return fmt.Sprintf("Feature (%s)", strings.Join(components, "\n"))
}
