package recommendedactions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RecommendationAction
}

type ListByServerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RecommendationAction
}

type ListByServerOperationOptions struct {
	SessionId *string
}

func DefaultListByServerOperationOptions() ListByServerOperationOptions {
	return ListByServerOperationOptions{}
}

func (o ListByServerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByServerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByServerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.SessionId != nil {
		out.Append("sessionId", fmt.Sprintf("%v", *o.SessionId))
	}
	return &out
}

// ListByServer ...
func (c RecommendedActionsClient) ListByServer(ctx context.Context, id AdvisorId, options ListByServerOperationOptions) (result ListByServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/recommendedActions", id.ID()),
		OptionsObject: options,
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
		Values *[]RecommendationAction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByServerComplete retrieves all the results into a single object
func (c RecommendedActionsClient) ListByServerComplete(ctx context.Context, id AdvisorId, options ListByServerOperationOptions) (ListByServerCompleteResult, error) {
	return c.ListByServerCompleteMatchingPredicate(ctx, id, options, RecommendationActionOperationPredicate{})
}

// ListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RecommendedActionsClient) ListByServerCompleteMatchingPredicate(ctx context.Context, id AdvisorId, options ListByServerOperationOptions, predicate RecommendationActionOperationPredicate) (result ListByServerCompleteResult, err error) {
	items := make([]RecommendationAction, 0)

	resp, err := c.ListByServer(ctx, id, options)
	if err != nil {
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

	result = ListByServerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
