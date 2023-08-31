package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = KubernetesClusterId{}

// KubernetesClusterId is a struct representing the Resource ID for a Kubernetes Cluster
type KubernetesClusterId struct {
	SubscriptionId        string
	ResourceGroupName     string
	KubernetesClusterName string
}

// NewKubernetesClusterID returns a new KubernetesClusterId struct
func NewKubernetesClusterID(subscriptionId string, resourceGroupName string, kubernetesClusterName string) KubernetesClusterId {
	return KubernetesClusterId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		KubernetesClusterName: kubernetesClusterName,
	}
}

// ParseKubernetesClusterID parses 'input' into a KubernetesClusterId
func ParseKubernetesClusterID(input string) (*KubernetesClusterId, error) {
	parser := resourceids.NewParserFromResourceIdType(KubernetesClusterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := KubernetesClusterId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.KubernetesClusterName, ok = parsed.Parsed["kubernetesClusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "kubernetesClusterName", *parsed)
	}

	return &id, nil
}

// ParseKubernetesClusterIDInsensitively parses 'input' case-insensitively into a KubernetesClusterId
// note: this method should only be used for API response data and not user input
func ParseKubernetesClusterIDInsensitively(input string) (*KubernetesClusterId, error) {
	parser := resourceids.NewParserFromResourceIdType(KubernetesClusterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := KubernetesClusterId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.KubernetesClusterName, ok = parsed.Parsed["kubernetesClusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "kubernetesClusterName", *parsed)
	}

	return &id, nil
}

// ValidateKubernetesClusterID checks that 'input' can be parsed as a Kubernetes Cluster ID
func ValidateKubernetesClusterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseKubernetesClusterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Kubernetes Cluster ID
func (id KubernetesClusterId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/kubernetesClusters/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.KubernetesClusterName)
}

// Segments returns a slice of Resource ID Segments which comprise this Kubernetes Cluster ID
func (id KubernetesClusterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticKubernetesClusters", "kubernetesClusters", "kubernetesClusters"),
		resourceids.UserSpecifiedSegment("kubernetesClusterName", "kubernetesClusterValue"),
	}
}

// String returns a human-readable description of this Kubernetes Cluster ID
func (id KubernetesClusterId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Kubernetes Cluster Name: %q", id.KubernetesClusterName),
	}
	return fmt.Sprintf("Kubernetes Cluster (%s)", strings.Join(components, "\n"))
}
