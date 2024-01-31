package servertrustcertificates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServerTrustCertificateId{}

// ServerTrustCertificateId is a struct representing the Resource ID for a Server Trust Certificate
type ServerTrustCertificateId struct {
	SubscriptionId             string
	ResourceGroupName          string
	ManagedInstanceName        string
	ServerTrustCertificateName string
}

// NewServerTrustCertificateID returns a new ServerTrustCertificateId struct
func NewServerTrustCertificateID(subscriptionId string, resourceGroupName string, managedInstanceName string, serverTrustCertificateName string) ServerTrustCertificateId {
	return ServerTrustCertificateId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		ManagedInstanceName:        managedInstanceName,
		ServerTrustCertificateName: serverTrustCertificateName,
	}
}

// ParseServerTrustCertificateID parses 'input' into a ServerTrustCertificateId
func ParseServerTrustCertificateID(input string) (*ServerTrustCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServerTrustCertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServerTrustCertificateId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServerTrustCertificateIDInsensitively parses 'input' case-insensitively into a ServerTrustCertificateId
// note: this method should only be used for API response data and not user input
func ParseServerTrustCertificateIDInsensitively(input string) (*ServerTrustCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServerTrustCertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServerTrustCertificateId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServerTrustCertificateId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ServerTrustCertificateName, ok = input.Parsed["serverTrustCertificateName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serverTrustCertificateName", input)
	}

	return nil
}

// ValidateServerTrustCertificateID checks that 'input' can be parsed as a Server Trust Certificate ID
func ValidateServerTrustCertificateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServerTrustCertificateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Server Trust Certificate ID
func (id ServerTrustCertificateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/serverTrustCertificates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.ServerTrustCertificateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Server Trust Certificate ID
func (id ServerTrustCertificateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticServerTrustCertificates", "serverTrustCertificates", "serverTrustCertificates"),
		resourceids.UserSpecifiedSegment("serverTrustCertificateName", "serverTrustCertificateValue"),
	}
}

// String returns a human-readable description of this Server Trust Certificate ID
func (id ServerTrustCertificateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Server Trust Certificate Name: %q", id.ServerTrustCertificateName),
	}
	return fmt.Sprintf("Server Trust Certificate (%s)", strings.Join(components, "\n"))
}
