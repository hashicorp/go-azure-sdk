package querytexts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = QueryTextId{}

// QueryTextId is a struct representing the Resource ID for a Query Text
type QueryTextId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	QueryId           string
}

// NewQueryTextID returns a new QueryTextId struct
func NewQueryTextID(subscriptionId string, resourceGroupName string, serverName string, queryId string) QueryTextId {
	return QueryTextId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		QueryId:           queryId,
	}
}

// ParseQueryTextID parses 'input' into a QueryTextId
func ParseQueryTextID(input string) (*QueryTextId, error) {
	parser := resourceids.NewParserFromResourceIdType(QueryTextId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QueryTextId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, fmt.Errorf("the segment 'serverName' was not found in the resource id %q", input)
	}

	if id.QueryId, ok = parsed.Parsed["queryId"]; !ok {
		return nil, fmt.Errorf("the segment 'queryId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseQueryTextIDInsensitively parses 'input' case-insensitively into a QueryTextId
// note: this method should only be used for API response data and not user input
func ParseQueryTextIDInsensitively(input string) (*QueryTextId, error) {
	parser := resourceids.NewParserFromResourceIdType(QueryTextId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QueryTextId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, fmt.Errorf("the segment 'serverName' was not found in the resource id %q", input)
	}

	if id.QueryId, ok = parsed.Parsed["queryId"]; !ok {
		return nil, fmt.Errorf("the segment 'queryId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateQueryTextID checks that 'input' can be parsed as a Query Text ID
func ValidateQueryTextID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseQueryTextID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Query Text ID
func (id QueryTextId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/servers/%s/queryTexts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.QueryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Query Text ID
func (id QueryTextId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticQueryTexts", "queryTexts", "queryTexts"),
		resourceids.UserSpecifiedSegment("queryId", "queryIdValue"),
	}
}

// String returns a human-readable description of this Query Text ID
func (id QueryTextId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Query: %q", id.QueryId),
	}
	return fmt.Sprintf("Query Text (%s)", strings.Join(components, "\n"))
}
