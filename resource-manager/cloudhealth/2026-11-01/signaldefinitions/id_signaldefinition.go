package signaldefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SignaldefinitionId{})
}

var _ resourceids.ResourceId = &SignaldefinitionId{}

// SignaldefinitionId is a struct representing the Resource ID for a Signaldefinition
type SignaldefinitionId struct {
	SubscriptionId       string
	ResourceGroupName    string
	HealthmodelName      string
	SignaldefinitionName string
}

// NewSignaldefinitionID returns a new SignaldefinitionId struct
func NewSignaldefinitionID(subscriptionId string, resourceGroupName string, healthmodelName string, signaldefinitionName string) SignaldefinitionId {
	return SignaldefinitionId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		HealthmodelName:      healthmodelName,
		SignaldefinitionName: signaldefinitionName,
	}
}

// ParseSignaldefinitionID parses 'input' into a SignaldefinitionId
func ParseSignaldefinitionID(input string) (*SignaldefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SignaldefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SignaldefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSignaldefinitionIDInsensitively parses 'input' case-insensitively into a SignaldefinitionId
// note: this method should only be used for API response data and not user input
func ParseSignaldefinitionIDInsensitively(input string) (*SignaldefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SignaldefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SignaldefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SignaldefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.HealthmodelName, ok = input.Parsed["healthmodelName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "healthmodelName", input)
	}

	if id.SignaldefinitionName, ok = input.Parsed["signaldefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "signaldefinitionName", input)
	}

	return nil
}

// ValidateSignaldefinitionID checks that 'input' can be parsed as a Signaldefinition ID
func ValidateSignaldefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSignaldefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Signaldefinition ID
func (id SignaldefinitionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CloudHealth/healthmodels/%s/signaldefinitions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HealthmodelName, id.SignaldefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Signaldefinition ID
func (id SignaldefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCloudHealth", "Microsoft.CloudHealth", "Microsoft.CloudHealth"),
		resourceids.StaticSegment("staticHealthmodels", "healthmodels", "healthmodels"),
		resourceids.UserSpecifiedSegment("healthmodelName", "healthmodelName"),
		resourceids.StaticSegment("staticSignaldefinitions", "signaldefinitions", "signaldefinitions"),
		resourceids.UserSpecifiedSegment("signaldefinitionName", "signaldefinitionName"),
	}
}

// String returns a human-readable description of this Signaldefinition ID
func (id SignaldefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Healthmodel Name: %q", id.HealthmodelName),
		fmt.Sprintf("Signaldefinition Name: %q", id.SignaldefinitionName),
	}
	return fmt.Sprintf("Signaldefinition (%s)", strings.Join(components, "\n"))
}
