package exchangeroleassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListExchangeRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignment
}

type ListExchangeRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignment
}

type ListExchangeRoleAssignmentsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListExchangeRoleAssignmentsOperationOptions() ListExchangeRoleAssignmentsOperationOptions {
	return ListExchangeRoleAssignmentsOperationOptions{}
}

func (o ListExchangeRoleAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListExchangeRoleAssignmentsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListExchangeRoleAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListExchangeRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeRoleAssignments - List unifiedRoleAssignments. Get a list of unifiedRoleAssignment objects for the
// provider. The following RBAC providers are currently supported: - directory (Microsoft Entra ID) - entitlement
// management (Microsoft Entra entitlement management) - Exchange Online
func (c ExchangeRoleAssignmentClient) ListExchangeRoleAssignments(ctx context.Context, options ListExchangeRoleAssignmentsOperationOptions) (result ListExchangeRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListExchangeRoleAssignmentsCustomPager{},
		Path:          "/roleManagement/exchange/roleAssignments",
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.UnifiedRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeRoleAssignmentsComplete retrieves all the results into a single object
func (c ExchangeRoleAssignmentClient) ListExchangeRoleAssignmentsComplete(ctx context.Context, options ListExchangeRoleAssignmentsOperationOptions) (ListExchangeRoleAssignmentsCompleteResult, error) {
	return c.ListExchangeRoleAssignmentsCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentOperationPredicate{})
}

// ListExchangeRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeRoleAssignmentClient) ListExchangeRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, options ListExchangeRoleAssignmentsOperationOptions, predicate UnifiedRoleAssignmentOperationPredicate) (result ListExchangeRoleAssignmentsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignment, 0)

	resp, err := c.ListExchangeRoleAssignments(ctx, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListExchangeRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
