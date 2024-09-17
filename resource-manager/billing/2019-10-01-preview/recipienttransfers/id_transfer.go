package recipienttransfers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TransferId{})
}

var _ resourceids.ResourceId = &TransferId{}

// TransferId is a struct representing the Resource ID for a Transfer
type TransferId struct {
	TransferName string
}

// NewTransferID returns a new TransferId struct
func NewTransferID(transferName string) TransferId {
	return TransferId{
		TransferName: transferName,
	}
}

// ParseTransferID parses 'input' into a TransferId
func ParseTransferID(input string) (*TransferId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TransferId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TransferId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTransferIDInsensitively parses 'input' case-insensitively into a TransferId
// note: this method should only be used for API response data and not user input
func ParseTransferIDInsensitively(input string) (*TransferId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TransferId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TransferId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TransferId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TransferName, ok = input.Parsed["transferName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "transferName", input)
	}

	return nil
}

// ValidateTransferID checks that 'input' can be parsed as a Transfer ID
func ValidateTransferID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTransferID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Transfer ID
func (id TransferId) ID() string {
	fmtString := "/providers/Microsoft.Billing/transfers/%s"
	return fmt.Sprintf(fmtString, id.TransferName)
}

// Segments returns a slice of Resource ID Segments which comprise this Transfer ID
func (id TransferId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticTransfers", "transfers", "transfers"),
		resourceids.UserSpecifiedSegment("transferName", "transferValue"),
	}
}

// String returns a human-readable description of this Transfer ID
func (id TransferId) String() string {
	components := []string{
		fmt.Sprintf("Transfer Name: %q", id.TransferName),
	}
	return fmt.Sprintf("Transfer (%s)", strings.Join(components, "\n"))
}
