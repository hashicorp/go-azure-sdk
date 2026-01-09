package pendingaccessreviewinstance

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

type ListPendingAccessReviewInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewInstance
}

type ListPendingAccessReviewInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewInstance
}

type ListPendingAccessReviewInstancesOperationOptions struct {
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

func DefaultListPendingAccessReviewInstancesOperationOptions() ListPendingAccessReviewInstancesOperationOptions {
	return ListPendingAccessReviewInstancesOperationOptions{}
}

func (o ListPendingAccessReviewInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPendingAccessReviewInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListPendingAccessReviewInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPendingAccessReviewInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstances - accessReviewInstance: pendingAccessReviewInstances (deprecated). Retrieve the
// accessReviewInstance objects pending approval by the calling user. A list of zero or more accessReviewInstance
// objects are returned, of which the calling user is an assigned reviewer.
func (c PendingAccessReviewInstanceClient) ListPendingAccessReviewInstances(ctx context.Context, options ListPendingAccessReviewInstancesOperationOptions) (result ListPendingAccessReviewInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPendingAccessReviewInstancesCustomPager{},
		Path:          "/me/pendingAccessReviewInstances",
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
		Values *[]beta.AccessReviewInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstancesComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceClient) ListPendingAccessReviewInstancesComplete(ctx context.Context, options ListPendingAccessReviewInstancesOperationOptions) (ListPendingAccessReviewInstancesCompleteResult, error) {
	return c.ListPendingAccessReviewInstancesCompleteMatchingPredicate(ctx, options, AccessReviewInstanceOperationPredicate{})
}

// ListPendingAccessReviewInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceClient) ListPendingAccessReviewInstancesCompleteMatchingPredicate(ctx context.Context, options ListPendingAccessReviewInstancesOperationOptions, predicate AccessReviewInstanceOperationPredicate) (result ListPendingAccessReviewInstancesCompleteResult, err error) {
	items := make([]beta.AccessReviewInstance, 0)

	resp, err := c.ListPendingAccessReviewInstances(ctx, options)
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

	result = ListPendingAccessReviewInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
