package oucontainer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&OuContainerId{})
}

var _ resourceids.ResourceId = &OuContainerId{}

// OuContainerId is a struct representing the Resource ID for a Ou Container
type OuContainerId struct {
	SubscriptionId    string
	ResourceGroupName string
	DomainServiceName string
	OuContainerName   string
}

// NewOuContainerID returns a new OuContainerId struct
func NewOuContainerID(subscriptionId string, resourceGroupName string, domainServiceName string, ouContainerName string) OuContainerId {
	return OuContainerId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		DomainServiceName: domainServiceName,
		OuContainerName:   ouContainerName,
	}
}

// ParseOuContainerID parses 'input' into a OuContainerId
func ParseOuContainerID(input string) (*OuContainerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OuContainerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OuContainerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseOuContainerIDInsensitively parses 'input' case-insensitively into a OuContainerId
// note: this method should only be used for API response data and not user input
func ParseOuContainerIDInsensitively(input string) (*OuContainerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OuContainerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OuContainerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *OuContainerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DomainServiceName, ok = input.Parsed["domainServiceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "domainServiceName", input)
	}

	if id.OuContainerName, ok = input.Parsed["ouContainerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "ouContainerName", input)
	}

	return nil
}

// ValidateOuContainerID checks that 'input' can be parsed as a Ou Container ID
func ValidateOuContainerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOuContainerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Ou Container ID
func (id OuContainerId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AAD/domainServices/%s/ouContainer/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DomainServiceName, id.OuContainerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Ou Container ID
func (id OuContainerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAAD", "Microsoft.AAD", "Microsoft.AAD"),
		resourceids.StaticSegment("staticDomainServices", "domainServices", "domainServices"),
		resourceids.UserSpecifiedSegment("domainServiceName", "domainServiceValue"),
		resourceids.StaticSegment("staticOuContainer", "ouContainer", "ouContainer"),
		resourceids.UserSpecifiedSegment("ouContainerName", "ouContainerValue"),
	}
}

// String returns a human-readable description of this Ou Container ID
func (id OuContainerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Domain Service Name: %q", id.DomainServiceName),
		fmt.Sprintf("Ou Container Name: %q", id.OuContainerName),
	}
	return fmt.Sprintf("Ou Container (%s)", strings.Join(components, "\n"))
}
