package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PublicCertificateId{}

// PublicCertificateId is a struct representing the Resource ID for a Public Certificate
type PublicCertificateId struct {
	SubscriptionId        string
	ResourceGroupName     string
	SiteName              string
	PublicCertificateName string
}

// NewPublicCertificateID returns a new PublicCertificateId struct
func NewPublicCertificateID(subscriptionId string, resourceGroupName string, siteName string, publicCertificateName string) PublicCertificateId {
	return PublicCertificateId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		SiteName:              siteName,
		PublicCertificateName: publicCertificateName,
	}
}

// ParsePublicCertificateID parses 'input' into a PublicCertificateId
func ParsePublicCertificateID(input string) (*PublicCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(PublicCertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PublicCertificateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.PublicCertificateName, ok = parsed.Parsed["publicCertificateName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "publicCertificateName", *parsed)
	}

	return &id, nil
}

// ParsePublicCertificateIDInsensitively parses 'input' case-insensitively into a PublicCertificateId
// note: this method should only be used for API response data and not user input
func ParsePublicCertificateIDInsensitively(input string) (*PublicCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(PublicCertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PublicCertificateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.PublicCertificateName, ok = parsed.Parsed["publicCertificateName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "publicCertificateName", *parsed)
	}

	return &id, nil
}

// ValidatePublicCertificateID checks that 'input' can be parsed as a Public Certificate ID
func ValidatePublicCertificateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePublicCertificateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Public Certificate ID
func (id PublicCertificateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/publicCertificates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.PublicCertificateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Public Certificate ID
func (id PublicCertificateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticPublicCertificates", "publicCertificates", "publicCertificates"),
		resourceids.UserSpecifiedSegment("publicCertificateName", "publicCertificateValue"),
	}
}

// String returns a human-readable description of this Public Certificate ID
func (id PublicCertificateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Public Certificate Name: %q", id.PublicCertificateName),
	}
	return fmt.Sprintf("Public Certificate (%s)", strings.Join(components, "\n"))
}
