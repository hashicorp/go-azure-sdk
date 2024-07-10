package endpointcertificates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&EndpointCertificateId{})
}

var _ resourceids.ResourceId = &EndpointCertificateId{}

// EndpointCertificateId is a struct representing the Resource ID for a Endpoint Certificate
type EndpointCertificateId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ManagedInstanceName     string
	EndpointCertificateName string
}

// NewEndpointCertificateID returns a new EndpointCertificateId struct
func NewEndpointCertificateID(subscriptionId string, resourceGroupName string, managedInstanceName string, endpointCertificateName string) EndpointCertificateId {
	return EndpointCertificateId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ManagedInstanceName:     managedInstanceName,
		EndpointCertificateName: endpointCertificateName,
	}
}

// ParseEndpointCertificateID parses 'input' into a EndpointCertificateId
func ParseEndpointCertificateID(input string) (*EndpointCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EndpointCertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EndpointCertificateId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEndpointCertificateIDInsensitively parses 'input' case-insensitively into a EndpointCertificateId
// note: this method should only be used for API response data and not user input
func ParseEndpointCertificateIDInsensitively(input string) (*EndpointCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EndpointCertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EndpointCertificateId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EndpointCertificateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedInstanceName, ok = input.Parsed["managedInstanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", input)
	}

	if id.EndpointCertificateName, ok = input.Parsed["endpointCertificateName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "endpointCertificateName", input)
	}

	return nil
}

// ValidateEndpointCertificateID checks that 'input' can be parsed as a Endpoint Certificate ID
func ValidateEndpointCertificateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEndpointCertificateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Endpoint Certificate ID
func (id EndpointCertificateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/endpointCertificates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.EndpointCertificateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Endpoint Certificate ID
func (id EndpointCertificateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticEndpointCertificates", "endpointCertificates", "endpointCertificates"),
		resourceids.UserSpecifiedSegment("endpointCertificateName", "endpointCertificateValue"),
	}
}

// String returns a human-readable description of this Endpoint Certificate ID
func (id EndpointCertificateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Endpoint Certificate Name: %q", id.EndpointCertificateName),
	}
	return fmt.Sprintf("Endpoint Certificate (%s)", strings.Join(components, "\n"))
}
