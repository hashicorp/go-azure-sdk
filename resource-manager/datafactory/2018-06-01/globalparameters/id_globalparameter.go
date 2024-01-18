package globalparameters

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GlobalParameterId{}

// GlobalParameterId is a struct representing the Resource ID for a Global Parameter
type GlobalParameterId struct {
	SubscriptionId      string
	ResourceGroupName   string
	FactoryName         string
	GlobalParameterName string
}

// NewGlobalParameterID returns a new GlobalParameterId struct
func NewGlobalParameterID(subscriptionId string, resourceGroupName string, factoryName string, globalParameterName string) GlobalParameterId {
	return GlobalParameterId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		FactoryName:         factoryName,
		GlobalParameterName: globalParameterName,
	}
}

// ParseGlobalParameterID parses 'input' into a GlobalParameterId
func ParseGlobalParameterID(input string) (*GlobalParameterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GlobalParameterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GlobalParameterId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGlobalParameterIDInsensitively parses 'input' case-insensitively into a GlobalParameterId
// note: this method should only be used for API response data and not user input
func ParseGlobalParameterIDInsensitively(input string) (*GlobalParameterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GlobalParameterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GlobalParameterId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GlobalParameterId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.FactoryName, ok = input.Parsed["factoryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "factoryName", input)
	}

	if id.GlobalParameterName, ok = input.Parsed["globalParameterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "globalParameterName", input)
	}

	return nil
}

// ValidateGlobalParameterID checks that 'input' can be parsed as a Global Parameter ID
func ValidateGlobalParameterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGlobalParameterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Global Parameter ID
func (id GlobalParameterId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataFactory/factories/%s/globalParameters/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FactoryName, id.GlobalParameterName)
}

// Segments returns a slice of Resource ID Segments which comprise this Global Parameter ID
func (id GlobalParameterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataFactory", "Microsoft.DataFactory", "Microsoft.DataFactory"),
		resourceids.StaticSegment("staticFactories", "factories", "factories"),
		resourceids.UserSpecifiedSegment("factoryName", "factoryValue"),
		resourceids.StaticSegment("staticGlobalParameters", "globalParameters", "globalParameters"),
		resourceids.UserSpecifiedSegment("globalParameterName", "globalParameterValue"),
	}
}

// String returns a human-readable description of this Global Parameter ID
func (id GlobalParameterId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Factory Name: %q", id.FactoryName),
		fmt.Sprintf("Global Parameter Name: %q", id.GlobalParameterName),
	}
	return fmt.Sprintf("Global Parameter (%s)", strings.Join(components, "\n"))
}
