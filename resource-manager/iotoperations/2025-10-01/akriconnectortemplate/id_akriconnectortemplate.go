package akriconnectortemplate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AkriConnectorTemplateId{})
}

var _ resourceids.ResourceId = &AkriConnectorTemplateId{}

// AkriConnectorTemplateId is a struct representing the Resource ID for a Akri Connector Template
type AkriConnectorTemplateId struct {
	SubscriptionId            string
	ResourceGroupName         string
	InstanceName              string
	AkriConnectorTemplateName string
}

// NewAkriConnectorTemplateID returns a new AkriConnectorTemplateId struct
func NewAkriConnectorTemplateID(subscriptionId string, resourceGroupName string, instanceName string, akriConnectorTemplateName string) AkriConnectorTemplateId {
	return AkriConnectorTemplateId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		InstanceName:              instanceName,
		AkriConnectorTemplateName: akriConnectorTemplateName,
	}
}

// ParseAkriConnectorTemplateID parses 'input' into a AkriConnectorTemplateId
func ParseAkriConnectorTemplateID(input string) (*AkriConnectorTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AkriConnectorTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AkriConnectorTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAkriConnectorTemplateIDInsensitively parses 'input' case-insensitively into a AkriConnectorTemplateId
// note: this method should only be used for API response data and not user input
func ParseAkriConnectorTemplateIDInsensitively(input string) (*AkriConnectorTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AkriConnectorTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AkriConnectorTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AkriConnectorTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.InstanceName, ok = input.Parsed["instanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "instanceName", input)
	}

	if id.AkriConnectorTemplateName, ok = input.Parsed["akriConnectorTemplateName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "akriConnectorTemplateName", input)
	}

	return nil
}

// ValidateAkriConnectorTemplateID checks that 'input' can be parsed as a Akri Connector Template ID
func ValidateAkriConnectorTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAkriConnectorTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Akri Connector Template ID
func (id AkriConnectorTemplateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.IoTOperations/instances/%s/akriConnectorTemplates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.InstanceName, id.AkriConnectorTemplateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Akri Connector Template ID
func (id AkriConnectorTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftIoTOperations", "Microsoft.IoTOperations", "Microsoft.IoTOperations"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("instanceName", "instanceName"),
		resourceids.StaticSegment("staticAkriConnectorTemplates", "akriConnectorTemplates", "akriConnectorTemplates"),
		resourceids.UserSpecifiedSegment("akriConnectorTemplateName", "akriConnectorTemplateName"),
	}
}

// String returns a human-readable description of this Akri Connector Template ID
func (id AkriConnectorTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Instance Name: %q", id.InstanceName),
		fmt.Sprintf("Akri Connector Template Name: %q", id.AkriConnectorTemplateName),
	}
	return fmt.Sprintf("Akri Connector Template (%s)", strings.Join(components, "\n"))
}
