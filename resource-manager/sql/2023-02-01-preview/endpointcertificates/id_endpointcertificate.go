package endpointcertificates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = EndpointCertificateId{}

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
	parser := resourceids.NewParserFromResourceIdType(EndpointCertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EndpointCertificateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.EndpointCertificateName, ok = parsed.Parsed["endpointCertificateName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "endpointCertificateName", *parsed)
	}

	return &id, nil
}

// ParseEndpointCertificateIDInsensitively parses 'input' case-insensitively into a EndpointCertificateId
// note: this method should only be used for API response data and not user input
func ParseEndpointCertificateIDInsensitively(input string) (*EndpointCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(EndpointCertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EndpointCertificateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.EndpointCertificateName, ok = parsed.Parsed["endpointCertificateName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "endpointCertificateName", *parsed)
	}

	return &id, nil
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
