package sqlpools

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SqlPoolId{}

// SqlPoolId is a struct representing the Resource ID for a Sql Pool
type SqlPoolId struct {
	BaseURI     string
	SqlPoolName string
}

// NewSqlPoolID returns a new SqlPoolId struct
func NewSqlPoolID(baseURI string, sqlPoolName string) SqlPoolId {
	return SqlPoolId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		SqlPoolName: sqlPoolName,
	}
}

// ParseSqlPoolID parses 'input' into a SqlPoolId
func ParseSqlPoolID(input string) (*SqlPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SqlPoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SqlPoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSqlPoolIDInsensitively parses 'input' case-insensitively into a SqlPoolId
// note: this method should only be used for API response data and not user input
func ParseSqlPoolIDInsensitively(input string) (*SqlPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SqlPoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SqlPoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SqlPoolId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.SqlPoolName, ok = input.Parsed["sqlPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sqlPoolName", input)
	}

	return nil
}

// ValidateSqlPoolID checks that 'input' can be parsed as a Sql Pool ID
func ValidateSqlPoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSqlPoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sql Pool ID
func (id SqlPoolId) ID() string {
	fmtString := "%s/sqlPools/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.SqlPoolName)
}

// Path returns the formatted Sql Pool ID without the BaseURI
func (id SqlPoolId) Path() string {
	fmtString := "/sqlPools/%s"
	return fmt.Sprintf(fmtString, id.SqlPoolName)
}

// PathElements returns the values of Sql Pool ID Segments without the BaseURI
func (id SqlPoolId) PathElements() []any {
	return []any{id.SqlPoolName}
}

// Segments returns a slice of Resource ID Segments which comprise this Sql Pool ID
func (id SqlPoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticSqlPools", "sqlPools", "sqlPools"),
		resourceids.UserSpecifiedSegment("sqlPoolName", "sqlPoolName"),
	}
}

// String returns a human-readable description of this Sql Pool ID
func (id SqlPoolId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Sql Pool Name: %q", id.SqlPoolName),
	}
	return fmt.Sprintf("Sql Pool (%s)", strings.Join(components, "\n"))
}
