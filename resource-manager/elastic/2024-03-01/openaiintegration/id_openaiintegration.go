package openaiintegration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&OpenAIIntegrationId{})
}

var _ resourceids.ResourceId = &OpenAIIntegrationId{}

// OpenAIIntegrationId is a struct representing the Resource ID for a Open A I Integration
type OpenAIIntegrationId struct {
	SubscriptionId        string
	ResourceGroupName     string
	MonitorName           string
	OpenAIIntegrationName string
}

// NewOpenAIIntegrationID returns a new OpenAIIntegrationId struct
func NewOpenAIIntegrationID(subscriptionId string, resourceGroupName string, monitorName string, openAIIntegrationName string) OpenAIIntegrationId {
	return OpenAIIntegrationId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		MonitorName:           monitorName,
		OpenAIIntegrationName: openAIIntegrationName,
	}
}

// ParseOpenAIIntegrationID parses 'input' into a OpenAIIntegrationId
func ParseOpenAIIntegrationID(input string) (*OpenAIIntegrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OpenAIIntegrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OpenAIIntegrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseOpenAIIntegrationIDInsensitively parses 'input' case-insensitively into a OpenAIIntegrationId
// note: this method should only be used for API response data and not user input
func ParseOpenAIIntegrationIDInsensitively(input string) (*OpenAIIntegrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OpenAIIntegrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OpenAIIntegrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *OpenAIIntegrationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.MonitorName, ok = input.Parsed["monitorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "monitorName", input)
	}

	if id.OpenAIIntegrationName, ok = input.Parsed["openAIIntegrationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openAIIntegrationName", input)
	}

	return nil
}

// ValidateOpenAIIntegrationID checks that 'input' can be parsed as a Open A I Integration ID
func ValidateOpenAIIntegrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOpenAIIntegrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Open A I Integration ID
func (id OpenAIIntegrationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Elastic/monitors/%s/openAIIntegrations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MonitorName, id.OpenAIIntegrationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Open A I Integration ID
func (id OpenAIIntegrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftElastic", "Microsoft.Elastic", "Microsoft.Elastic"),
		resourceids.StaticSegment("staticMonitors", "monitors", "monitors"),
		resourceids.UserSpecifiedSegment("monitorName", "monitorValue"),
		resourceids.StaticSegment("staticOpenAIIntegrations", "openAIIntegrations", "openAIIntegrations"),
		resourceids.UserSpecifiedSegment("openAIIntegrationName", "openAIIntegrationValue"),
	}
}

// String returns a human-readable description of this Open A I Integration ID
func (id OpenAIIntegrationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Monitor Name: %q", id.MonitorName),
		fmt.Sprintf("Open A I Integration Name: %q", id.OpenAIIntegrationName),
	}
	return fmt.Sprintf("Open A I Integration (%s)", strings.Join(components, "\n"))
}
