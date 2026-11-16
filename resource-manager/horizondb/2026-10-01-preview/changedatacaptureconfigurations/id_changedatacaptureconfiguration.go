package changedatacaptureconfigurations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ChangeDataCaptureConfigurationId{})
}

var _ resourceids.ResourceId = &ChangeDataCaptureConfigurationId{}

// ChangeDataCaptureConfigurationId is a struct representing the Resource ID for a Change Data Capture Configuration
type ChangeDataCaptureConfigurationId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	ClusterName                        string
	ChangeDataCaptureConfigurationName string
}

// NewChangeDataCaptureConfigurationID returns a new ChangeDataCaptureConfigurationId struct
func NewChangeDataCaptureConfigurationID(subscriptionId string, resourceGroupName string, clusterName string, changeDataCaptureConfigurationName string) ChangeDataCaptureConfigurationId {
	return ChangeDataCaptureConfigurationId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		ClusterName:                        clusterName,
		ChangeDataCaptureConfigurationName: changeDataCaptureConfigurationName,
	}
}

// ParseChangeDataCaptureConfigurationID parses 'input' into a ChangeDataCaptureConfigurationId
func ParseChangeDataCaptureConfigurationID(input string) (*ChangeDataCaptureConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ChangeDataCaptureConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ChangeDataCaptureConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseChangeDataCaptureConfigurationIDInsensitively parses 'input' case-insensitively into a ChangeDataCaptureConfigurationId
// note: this method should only be used for API response data and not user input
func ParseChangeDataCaptureConfigurationIDInsensitively(input string) (*ChangeDataCaptureConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ChangeDataCaptureConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ChangeDataCaptureConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ChangeDataCaptureConfigurationId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChangeDataCaptureConfigurationName, ok = input.Parsed["changeDataCaptureConfigurationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "changeDataCaptureConfigurationName", input)
	}

	return nil
}

// ValidateChangeDataCaptureConfigurationID checks that 'input' can be parsed as a Change Data Capture Configuration ID
func ValidateChangeDataCaptureConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseChangeDataCaptureConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Change Data Capture Configuration ID
func (id ChangeDataCaptureConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HorizonDb/clusters/%s/changeDataCaptureConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.ChangeDataCaptureConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Change Data Capture Configuration ID
func (id ChangeDataCaptureConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHorizonDb", "Microsoft.HorizonDb", "Microsoft.HorizonDb"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterName"),
		resourceids.StaticSegment("staticChangeDataCaptureConfigurations", "changeDataCaptureConfigurations", "changeDataCaptureConfigurations"),
		resourceids.UserSpecifiedSegment("changeDataCaptureConfigurationName", "changeDataCaptureConfigurationName"),
	}
}

// String returns a human-readable description of this Change Data Capture Configuration ID
func (id ChangeDataCaptureConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Change Data Capture Configuration Name: %q", id.ChangeDataCaptureConfigurationName),
	}
	return fmt.Sprintf("Change Data Capture Configuration (%s)", strings.Join(components, "\n"))
}
