package kubeenvironments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&KubeEnvironmentId{})
}

var _ resourceids.ResourceId = &KubeEnvironmentId{}

// KubeEnvironmentId is a struct representing the Resource ID for a Kube Environment
type KubeEnvironmentId struct {
	SubscriptionId      string
	ResourceGroupName   string
	KubeEnvironmentName string
}

// NewKubeEnvironmentID returns a new KubeEnvironmentId struct
func NewKubeEnvironmentID(subscriptionId string, resourceGroupName string, kubeEnvironmentName string) KubeEnvironmentId {
	return KubeEnvironmentId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		KubeEnvironmentName: kubeEnvironmentName,
	}
}

// ParseKubeEnvironmentID parses 'input' into a KubeEnvironmentId
func ParseKubeEnvironmentID(input string) (*KubeEnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KubeEnvironmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KubeEnvironmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseKubeEnvironmentIDInsensitively parses 'input' case-insensitively into a KubeEnvironmentId
// note: this method should only be used for API response data and not user input
func ParseKubeEnvironmentIDInsensitively(input string) (*KubeEnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KubeEnvironmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KubeEnvironmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *KubeEnvironmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.KubeEnvironmentName, ok = input.Parsed["kubeEnvironmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "kubeEnvironmentName", input)
	}

	return nil
}

// ValidateKubeEnvironmentID checks that 'input' can be parsed as a Kube Environment ID
func ValidateKubeEnvironmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseKubeEnvironmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Kube Environment ID
func (id KubeEnvironmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/kubeEnvironments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.KubeEnvironmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Kube Environment ID
func (id KubeEnvironmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticKubeEnvironments", "kubeEnvironments", "kubeEnvironments"),
		resourceids.UserSpecifiedSegment("kubeEnvironmentName", "kubeEnvironmentValue"),
	}
}

// String returns a human-readable description of this Kube Environment ID
func (id KubeEnvironmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Kube Environment Name: %q", id.KubeEnvironmentName),
	}
	return fmt.Sprintf("Kube Environment (%s)", strings.Join(components, "\n"))
}
