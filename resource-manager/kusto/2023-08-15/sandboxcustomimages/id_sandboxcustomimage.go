package sandboxcustomimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SandboxCustomImageId{})
}

var _ resourceids.ResourceId = &SandboxCustomImageId{}

// SandboxCustomImageId is a struct representing the Resource ID for a Sandbox Custom Image
type SandboxCustomImageId struct {
	SubscriptionId         string
	ResourceGroupName      string
	ClusterName            string
	SandboxCustomImageName string
}

// NewSandboxCustomImageID returns a new SandboxCustomImageId struct
func NewSandboxCustomImageID(subscriptionId string, resourceGroupName string, clusterName string, sandboxCustomImageName string) SandboxCustomImageId {
	return SandboxCustomImageId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		ClusterName:            clusterName,
		SandboxCustomImageName: sandboxCustomImageName,
	}
}

// ParseSandboxCustomImageID parses 'input' into a SandboxCustomImageId
func ParseSandboxCustomImageID(input string) (*SandboxCustomImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SandboxCustomImageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SandboxCustomImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSandboxCustomImageIDInsensitively parses 'input' case-insensitively into a SandboxCustomImageId
// note: this method should only be used for API response data and not user input
func ParseSandboxCustomImageIDInsensitively(input string) (*SandboxCustomImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SandboxCustomImageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SandboxCustomImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SandboxCustomImageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ClusterName, ok = input.Parsed["clusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterName", input)
	}

	if id.SandboxCustomImageName, ok = input.Parsed["sandboxCustomImageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sandboxCustomImageName", input)
	}

	return nil
}

// ValidateSandboxCustomImageID checks that 'input' can be parsed as a Sandbox Custom Image ID
func ValidateSandboxCustomImageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSandboxCustomImageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sandbox Custom Image ID
func (id SandboxCustomImageId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Kusto/clusters/%s/sandboxCustomImages/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.SandboxCustomImageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sandbox Custom Image ID
func (id SandboxCustomImageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftKusto", "Microsoft.Kusto", "Microsoft.Kusto"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticSandboxCustomImages", "sandboxCustomImages", "sandboxCustomImages"),
		resourceids.UserSpecifiedSegment("sandboxCustomImageName", "sandboxCustomImageValue"),
	}
}

// String returns a human-readable description of this Sandbox Custom Image ID
func (id SandboxCustomImageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Sandbox Custom Image Name: %q", id.SandboxCustomImageName),
	}
	return fmt.Sprintf("Sandbox Custom Image (%s)", strings.Join(components, "\n"))
}
