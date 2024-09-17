package kafkaconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&KafkaConfigurationId{})
}

var _ resourceids.ResourceId = &KafkaConfigurationId{}

// KafkaConfigurationId is a struct representing the Resource ID for a Kafka Configuration
type KafkaConfigurationId struct {
	SubscriptionId         string
	ResourceGroupName      string
	AccountName            string
	KafkaConfigurationName string
}

// NewKafkaConfigurationID returns a new KafkaConfigurationId struct
func NewKafkaConfigurationID(subscriptionId string, resourceGroupName string, accountName string, kafkaConfigurationName string) KafkaConfigurationId {
	return KafkaConfigurationId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		AccountName:            accountName,
		KafkaConfigurationName: kafkaConfigurationName,
	}
}

// ParseKafkaConfigurationID parses 'input' into a KafkaConfigurationId
func ParseKafkaConfigurationID(input string) (*KafkaConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KafkaConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KafkaConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseKafkaConfigurationIDInsensitively parses 'input' case-insensitively into a KafkaConfigurationId
// note: this method should only be used for API response data and not user input
func ParseKafkaConfigurationIDInsensitively(input string) (*KafkaConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KafkaConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KafkaConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *KafkaConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AccountName, ok = input.Parsed["accountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accountName", input)
	}

	if id.KafkaConfigurationName, ok = input.Parsed["kafkaConfigurationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "kafkaConfigurationName", input)
	}

	return nil
}

// ValidateKafkaConfigurationID checks that 'input' can be parsed as a Kafka Configuration ID
func ValidateKafkaConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseKafkaConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Kafka Configuration ID
func (id KafkaConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Purview/accounts/%s/kafkaConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.KafkaConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Kafka Configuration ID
func (id KafkaConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPurview", "Microsoft.Purview", "Microsoft.Purview"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticKafkaConfigurations", "kafkaConfigurations", "kafkaConfigurations"),
		resourceids.UserSpecifiedSegment("kafkaConfigurationName", "kafkaConfigurationValue"),
	}
}

// String returns a human-readable description of this Kafka Configuration ID
func (id KafkaConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Kafka Configuration Name: %q", id.KafkaConfigurationName),
	}
	return fmt.Sprintf("Kafka Configuration (%s)", strings.Join(components, "\n"))
}
