package serverlessruntimes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ServerlessRuntimeId{})
}

var _ resourceids.ResourceId = &ServerlessRuntimeId{}

// ServerlessRuntimeId is a struct representing the Resource ID for a Serverless Runtime
type ServerlessRuntimeId struct {
	SubscriptionId        string
	ResourceGroupName     string
	OrganizationName      string
	ServerlessRuntimeName string
}

// NewServerlessRuntimeID returns a new ServerlessRuntimeId struct
func NewServerlessRuntimeID(subscriptionId string, resourceGroupName string, organizationName string, serverlessRuntimeName string) ServerlessRuntimeId {
	return ServerlessRuntimeId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		OrganizationName:      organizationName,
		ServerlessRuntimeName: serverlessRuntimeName,
	}
}

// ParseServerlessRuntimeID parses 'input' into a ServerlessRuntimeId
func ParseServerlessRuntimeID(input string) (*ServerlessRuntimeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServerlessRuntimeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServerlessRuntimeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServerlessRuntimeIDInsensitively parses 'input' case-insensitively into a ServerlessRuntimeId
// note: this method should only be used for API response data and not user input
func ParseServerlessRuntimeIDInsensitively(input string) (*ServerlessRuntimeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServerlessRuntimeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServerlessRuntimeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServerlessRuntimeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.OrganizationName, ok = input.Parsed["organizationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "organizationName", input)
	}

	if id.ServerlessRuntimeName, ok = input.Parsed["serverlessRuntimeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serverlessRuntimeName", input)
	}

	return nil
}

// ValidateServerlessRuntimeID checks that 'input' can be parsed as a Serverless Runtime ID
func ValidateServerlessRuntimeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServerlessRuntimeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Serverless Runtime ID
func (id ServerlessRuntimeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Informatica.DataManagement/organizations/%s/serverlessRuntimes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.OrganizationName, id.ServerlessRuntimeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Serverless Runtime ID
func (id ServerlessRuntimeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticInformaticaDataManagement", "Informatica.DataManagement", "Informatica.DataManagement"),
		resourceids.StaticSegment("staticOrganizations", "organizations", "organizations"),
		resourceids.UserSpecifiedSegment("organizationName", "organizationName"),
		resourceids.StaticSegment("staticServerlessRuntimes", "serverlessRuntimes", "serverlessRuntimes"),
		resourceids.UserSpecifiedSegment("serverlessRuntimeName", "serverlessRuntimeName"),
	}
}

// String returns a human-readable description of this Serverless Runtime ID
func (id ServerlessRuntimeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Organization Name: %q", id.OrganizationName),
		fmt.Sprintf("Serverless Runtime Name: %q", id.ServerlessRuntimeName),
	}
	return fmt.Sprintf("Serverless Runtime (%s)", strings.Join(components, "\n"))
}
