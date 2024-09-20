package managedapplogcollectionrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListManagedAppLogCollectionRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedAppLogCollectionRequest
}

type ListManagedAppLogCollectionRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedAppLogCollectionRequest
}

type ListManagedAppLogCollectionRequestsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListManagedAppLogCollectionRequestsOperationOptions() ListManagedAppLogCollectionRequestsOperationOptions {
	return ListManagedAppLogCollectionRequestsOperationOptions{}
}

func (o ListManagedAppLogCollectionRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListManagedAppLogCollectionRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListManagedAppLogCollectionRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListManagedAppLogCollectionRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedAppLogCollectionRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedAppLogCollectionRequests - Get managedAppLogCollectionRequests from me. Zero or more log collection
// requests triggered for the user.
func (c ManagedAppLogCollectionRequestClient) ListManagedAppLogCollectionRequests(ctx context.Context, options ListManagedAppLogCollectionRequestsOperationOptions) (result ListManagedAppLogCollectionRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListManagedAppLogCollectionRequestsCustomPager{},
		Path:          "/me/managedAppLogCollectionRequests",
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
		Values *[]beta.ManagedAppLogCollectionRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedAppLogCollectionRequestsComplete retrieves all the results into a single object
func (c ManagedAppLogCollectionRequestClient) ListManagedAppLogCollectionRequestsComplete(ctx context.Context, options ListManagedAppLogCollectionRequestsOperationOptions) (ListManagedAppLogCollectionRequestsCompleteResult, error) {
	return c.ListManagedAppLogCollectionRequestsCompleteMatchingPredicate(ctx, options, ManagedAppLogCollectionRequestOperationPredicate{})
}

// ListManagedAppLogCollectionRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedAppLogCollectionRequestClient) ListManagedAppLogCollectionRequestsCompleteMatchingPredicate(ctx context.Context, options ListManagedAppLogCollectionRequestsOperationOptions, predicate ManagedAppLogCollectionRequestOperationPredicate) (result ListManagedAppLogCollectionRequestsCompleteResult, err error) {
	items := make([]beta.ManagedAppLogCollectionRequest, 0)

	resp, err := c.ListManagedAppLogCollectionRequests(ctx, options)
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

	result = ListManagedAppLogCollectionRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
