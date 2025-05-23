package sitecertificates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SlotCertificateId{})
}

var _ resourceids.ResourceId = &SlotCertificateId{}

// SlotCertificateId is a struct representing the Resource ID for a Slot Certificate
type SlotCertificateId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	CertificateName   string
}

// NewSlotCertificateID returns a new SlotCertificateId struct
func NewSlotCertificateID(subscriptionId string, resourceGroupName string, siteName string, slotName string, certificateName string) SlotCertificateId {
	return SlotCertificateId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		CertificateName:   certificateName,
	}
}

// ParseSlotCertificateID parses 'input' into a SlotCertificateId
func ParseSlotCertificateID(input string) (*SlotCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SlotCertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SlotCertificateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSlotCertificateIDInsensitively parses 'input' case-insensitively into a SlotCertificateId
// note: this method should only be used for API response data and not user input
func ParseSlotCertificateIDInsensitively(input string) (*SlotCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SlotCertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SlotCertificateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SlotCertificateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SiteName, ok = input.Parsed["siteName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteName", input)
	}

	if id.SlotName, ok = input.Parsed["slotName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "slotName", input)
	}

	if id.CertificateName, ok = input.Parsed["certificateName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateName", input)
	}

	return nil
}

// ValidateSlotCertificateID checks that 'input' can be parsed as a Slot Certificate ID
func ValidateSlotCertificateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotCertificateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Certificate ID
func (id SlotCertificateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/certificates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.CertificateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Certificate ID
func (id SlotCertificateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteName"),
		resourceids.StaticSegment("staticSlots", "slots", "slots"),
		resourceids.UserSpecifiedSegment("slotName", "slotName"),
		resourceids.StaticSegment("staticCertificates", "certificates", "certificates"),
		resourceids.UserSpecifiedSegment("certificateName", "certificateName"),
	}
}

// String returns a human-readable description of this Slot Certificate ID
func (id SlotCertificateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Certificate Name: %q", id.CertificateName),
	}
	return fmt.Sprintf("Slot Certificate (%s)", strings.Join(components, "\n"))
}
