package certificateordersdiagnostics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CertificateOrderId{}

// CertificateOrderId is a struct representing the Resource ID for a Certificate Order
type CertificateOrderId struct {
	SubscriptionId       string
	ResourceGroupName    string
	CertificateOrderName string
}

// NewCertificateOrderID returns a new CertificateOrderId struct
func NewCertificateOrderID(subscriptionId string, resourceGroupName string, certificateOrderName string) CertificateOrderId {
	return CertificateOrderId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		CertificateOrderName: certificateOrderName,
	}
}

// ParseCertificateOrderID parses 'input' into a CertificateOrderId
func ParseCertificateOrderID(input string) (*CertificateOrderId, error) {
	parser := resourceids.NewParserFromResourceIdType(CertificateOrderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CertificateOrderId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.CertificateOrderName, ok = parsed.Parsed["certificateOrderName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "certificateOrderName", *parsed)
	}

	return &id, nil
}

// ParseCertificateOrderIDInsensitively parses 'input' case-insensitively into a CertificateOrderId
// note: this method should only be used for API response data and not user input
func ParseCertificateOrderIDInsensitively(input string) (*CertificateOrderId, error) {
	parser := resourceids.NewParserFromResourceIdType(CertificateOrderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CertificateOrderId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.CertificateOrderName, ok = parsed.Parsed["certificateOrderName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "certificateOrderName", *parsed)
	}

	return &id, nil
}

// ValidateCertificateOrderID checks that 'input' can be parsed as a Certificate Order ID
func ValidateCertificateOrderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCertificateOrderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Certificate Order ID
func (id CertificateOrderId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CertificateRegistration/certificateOrders/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.CertificateOrderName)
}

// Segments returns a slice of Resource ID Segments which comprise this Certificate Order ID
func (id CertificateOrderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCertificateRegistration", "Microsoft.CertificateRegistration", "Microsoft.CertificateRegistration"),
		resourceids.StaticSegment("staticCertificateOrders", "certificateOrders", "certificateOrders"),
		resourceids.UserSpecifiedSegment("certificateOrderName", "certificateOrderValue"),
	}
}

// String returns a human-readable description of this Certificate Order ID
func (id CertificateOrderId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Certificate Order Name: %q", id.CertificateOrderName),
	}
	return fmt.Sprintf("Certificate Order (%s)", strings.Join(components, "\n"))
}
