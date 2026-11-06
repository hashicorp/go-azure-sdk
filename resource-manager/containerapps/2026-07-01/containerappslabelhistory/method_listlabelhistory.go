package containerappslabelhistory

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListLabelHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]LabelHistory
}

type ListLabelHistoryCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []LabelHistory
}

type ListLabelHistoryOperationOptions struct {
	Filter *string
}

func DefaultListLabelHistoryOperationOptions() ListLabelHistoryOperationOptions {
	return ListLabelHistoryOperationOptions{}
}

func (o ListLabelHistoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLabelHistoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListLabelHistoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type ListLabelHistoryCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListLabelHistoryCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLabelHistory ...
func (c ContainerAppsLabelHistoryClient) ListLabelHistory(ctx context.Context, id ContainerAppId, options ListLabelHistoryOperationOptions) (result ListLabelHistoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLabelHistoryCustomPager{},
		Path:          fmt.Sprintf("%s/labelHistories", id.ID()),
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
		Values *[]LabelHistory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLabelHistoryComplete retrieves all the results into a single object
func (c ContainerAppsLabelHistoryClient) ListLabelHistoryComplete(ctx context.Context, id ContainerAppId, options ListLabelHistoryOperationOptions) (ListLabelHistoryCompleteResult, error) {
	return c.ListLabelHistoryCompleteMatchingPredicate(ctx, id, options, LabelHistoryOperationPredicate{})
}

// ListLabelHistoryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContainerAppsLabelHistoryClient) ListLabelHistoryCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, options ListLabelHistoryOperationOptions, predicate LabelHistoryOperationPredicate) (result ListLabelHistoryCompleteResult, err error) {
	items := make([]LabelHistory, 0)

	resp, err := c.ListLabelHistory(ctx, id, options)
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

	result = ListLabelHistoryCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
