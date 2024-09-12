package permissionsanalyticazurefinding

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPermissionsAnalyticAzureFindingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Finding
}

type ListPermissionsAnalyticAzureFindingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Finding
}

type ListPermissionsAnalyticAzureFindingsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListPermissionsAnalyticAzureFindingsOperationOptions() ListPermissionsAnalyticAzureFindingsOperationOptions {
	return ListPermissionsAnalyticAzureFindingsOperationOptions{}
}

func (o ListPermissionsAnalyticAzureFindingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPermissionsAnalyticAzureFindingsOperationOptions) ToOData() *odata.Query {
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

func (o ListPermissionsAnalyticAzureFindingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPermissionsAnalyticAzureFindingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionsAnalyticAzureFindingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionsAnalyticAzureFindings - Get findings from identityGovernance. The output of the permissions usage data
// analysis performed by Permissions Management to assess risk with identities and resources.
func (c PermissionsAnalyticAzureFindingClient) ListPermissionsAnalyticAzureFindings(ctx context.Context, options ListPermissionsAnalyticAzureFindingsOperationOptions) (result ListPermissionsAnalyticAzureFindingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPermissionsAnalyticAzureFindingsCustomPager{},
		Path:          "/identityGovernance/permissionsAnalytics/azure/findings",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.Finding, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalFindingImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.Finding (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListPermissionsAnalyticAzureFindingsComplete retrieves all the results into a single object
func (c PermissionsAnalyticAzureFindingClient) ListPermissionsAnalyticAzureFindingsComplete(ctx context.Context, options ListPermissionsAnalyticAzureFindingsOperationOptions) (ListPermissionsAnalyticAzureFindingsCompleteResult, error) {
	return c.ListPermissionsAnalyticAzureFindingsCompleteMatchingPredicate(ctx, options, FindingOperationPredicate{})
}

// ListPermissionsAnalyticAzureFindingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionsAnalyticAzureFindingClient) ListPermissionsAnalyticAzureFindingsCompleteMatchingPredicate(ctx context.Context, options ListPermissionsAnalyticAzureFindingsOperationOptions, predicate FindingOperationPredicate) (result ListPermissionsAnalyticAzureFindingsCompleteResult, err error) {
	items := make([]beta.Finding, 0)

	resp, err := c.ListPermissionsAnalyticAzureFindings(ctx, options)
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

	result = ListPermissionsAnalyticAzureFindingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
