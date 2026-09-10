package aigateways

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AigatewayId{})
}

var _ resourceids.ResourceId = &AigatewayId{}

// AigatewayId is a struct representing the Resource ID for a Aigateway
type AigatewayId struct {
	SubscriptionId    string
	ResourceGroupName string
	AigatewayName     string
}

// NewAigatewayID returns a new AigatewayId struct
func NewAigatewayID(subscriptionId string, resourceGroupName string, aigatewayName string) AigatewayId {
	return AigatewayId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AigatewayName:     aigatewayName,
	}
}

// ParseAigatewayID parses 'input' into a AigatewayId
func ParseAigatewayID(input string) (*AigatewayId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AigatewayId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AigatewayId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAigatewayIDInsensitively parses 'input' case-insensitively into a AigatewayId
// note: this method should only be used for API response data and not user input
func ParseAigatewayIDInsensitively(input string) (*AigatewayId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AigatewayId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AigatewayId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AigatewayId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AigatewayName, ok = input.Parsed["aigatewayName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "aigatewayName", input)
	}

	return nil
}

// ValidateAigatewayID checks that 'input' can be parsed as a Aigateway ID
func ValidateAigatewayID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAigatewayID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Aigateway ID
func (id AigatewayId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/aigateways/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AigatewayName)
}

// Segments returns a slice of Resource ID Segments which comprise this Aigateway ID
func (id AigatewayId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticAigateways", "aigateways", "aigateways"),
		resourceids.UserSpecifiedSegment("aigatewayName", "aigatewayName"),
	}
}

// String returns a human-readable description of this Aigateway ID
func (id AigatewayId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Aigateway Name: %q", id.AigatewayName),
	}
	return fmt.Sprintf("Aigateway (%s)", strings.Join(components, "\n"))
}
