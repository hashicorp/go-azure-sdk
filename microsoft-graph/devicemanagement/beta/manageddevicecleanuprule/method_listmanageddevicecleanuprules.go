package manageddevicecleanuprule

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

type ListManagedDeviceCleanupRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedDeviceCleanupRule
}

type ListManagedDeviceCleanupRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedDeviceCleanupRule
}

type ListManagedDeviceCleanupRulesOperationOptions struct {
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

func DefaultListManagedDeviceCleanupRulesOperationOptions() ListManagedDeviceCleanupRulesOperationOptions {
	return ListManagedDeviceCleanupRulesOperationOptions{}
}

func (o ListManagedDeviceCleanupRulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListManagedDeviceCleanupRulesOperationOptions) ToOData() *odata.Query {
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

func (o ListManagedDeviceCleanupRulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListManagedDeviceCleanupRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceCleanupRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceCleanupRules - Get managedDeviceCleanupRules from deviceManagement. Device cleanup rule V2
func (c ManagedDeviceCleanupRuleClient) ListManagedDeviceCleanupRules(ctx context.Context, options ListManagedDeviceCleanupRulesOperationOptions) (result ListManagedDeviceCleanupRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListManagedDeviceCleanupRulesCustomPager{},
		Path:          "/deviceManagement/managedDeviceCleanupRules",
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
		Values *[]beta.ManagedDeviceCleanupRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceCleanupRulesComplete retrieves all the results into a single object
func (c ManagedDeviceCleanupRuleClient) ListManagedDeviceCleanupRulesComplete(ctx context.Context, options ListManagedDeviceCleanupRulesOperationOptions) (ListManagedDeviceCleanupRulesCompleteResult, error) {
	return c.ListManagedDeviceCleanupRulesCompleteMatchingPredicate(ctx, options, ManagedDeviceCleanupRuleOperationPredicate{})
}

// ListManagedDeviceCleanupRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceCleanupRuleClient) ListManagedDeviceCleanupRulesCompleteMatchingPredicate(ctx context.Context, options ListManagedDeviceCleanupRulesOperationOptions, predicate ManagedDeviceCleanupRuleOperationPredicate) (result ListManagedDeviceCleanupRulesCompleteResult, err error) {
	items := make([]beta.ManagedDeviceCleanupRule, 0)

	resp, err := c.ListManagedDeviceCleanupRules(ctx, options)
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

	result = ListManagedDeviceCleanupRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
