package certificateordersdiagnostics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CertificateOrderDetectorId{})
}

var _ resourceids.ResourceId = &CertificateOrderDetectorId{}

// CertificateOrderDetectorId is a struct representing the Resource ID for a Certificate Order Detector
type CertificateOrderDetectorId struct {
	SubscriptionId       string
	ResourceGroupName    string
	CertificateOrderName string
	DetectorName         string
}

// NewCertificateOrderDetectorID returns a new CertificateOrderDetectorId struct
func NewCertificateOrderDetectorID(subscriptionId string, resourceGroupName string, certificateOrderName string, detectorName string) CertificateOrderDetectorId {
	return CertificateOrderDetectorId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		CertificateOrderName: certificateOrderName,
		DetectorName:         detectorName,
	}
}

// ParseCertificateOrderDetectorID parses 'input' into a CertificateOrderDetectorId
func ParseCertificateOrderDetectorID(input string) (*CertificateOrderDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CertificateOrderDetectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CertificateOrderDetectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCertificateOrderDetectorIDInsensitively parses 'input' case-insensitively into a CertificateOrderDetectorId
// note: this method should only be used for API response data and not user input
func ParseCertificateOrderDetectorIDInsensitively(input string) (*CertificateOrderDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CertificateOrderDetectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CertificateOrderDetectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CertificateOrderDetectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.CertificateOrderName, ok = input.Parsed["certificateOrderName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateOrderName", input)
	}

	if id.DetectorName, ok = input.Parsed["detectorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectorName", input)
	}

	return nil
}

// ValidateCertificateOrderDetectorID checks that 'input' can be parsed as a Certificate Order Detector ID
func ValidateCertificateOrderDetectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCertificateOrderDetectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Certificate Order Detector ID
func (id CertificateOrderDetectorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CertificateRegistration/certificateOrders/%s/detectors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.CertificateOrderName, id.DetectorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Certificate Order Detector ID
func (id CertificateOrderDetectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCertificateRegistration", "Microsoft.CertificateRegistration", "Microsoft.CertificateRegistration"),
		resourceids.StaticSegment("staticCertificateOrders", "certificateOrders", "certificateOrders"),
		resourceids.UserSpecifiedSegment("certificateOrderName", "certificateOrderName"),
		resourceids.StaticSegment("staticDetectors", "detectors", "detectors"),
		resourceids.UserSpecifiedSegment("detectorName", "detectorName"),
	}
}

// String returns a human-readable description of this Certificate Order Detector ID
func (id CertificateOrderDetectorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Certificate Order Name: %q", id.CertificateOrderName),
		fmt.Sprintf("Detector Name: %q", id.DetectorName),
	}
	return fmt.Sprintf("Certificate Order Detector (%s)", strings.Join(components, "\n"))
}
