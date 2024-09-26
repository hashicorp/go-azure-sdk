package modelcapacities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ModelCapacityListResultValueInlined
}

type ListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ModelCapacityListResultValueInlined
}

type ListOperationOptions struct {
	ModelFormat  *string
	ModelName    *string
	ModelVersion *string
}

func DefaultListOperationOptions() ListOperationOptions {
	return ListOperationOptions{}
}

func (o ListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ModelFormat != nil {
		out.Append("modelFormat", fmt.Sprintf("%v", *o.ModelFormat))
	}
	if o.ModelName != nil {
		out.Append("modelName", fmt.Sprintf("%v", *o.ModelName))
	}
	if o.ModelVersion != nil {
		out.Append("modelVersion", fmt.Sprintf("%v", *o.ModelVersion))
	}
	return &out
}

type ListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// List ...
func (c ModelCapacitiesClient) List(ctx context.Context, id commonids.SubscriptionId, options ListOperationOptions) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.CognitiveServices/modelCapacities", id.ID()),
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
		Values *[]ModelCapacityListResultValueInlined `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplete retrieves all the results into a single object
func (c ModelCapacitiesClient) ListComplete(ctx context.Context, id commonids.SubscriptionId, options ListOperationOptions) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, options, ModelCapacityListResultValueInlinedOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ModelCapacitiesClient) ListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options ListOperationOptions, predicate ModelCapacityListResultValueInlinedOperationPredicate) (result ListCompleteResult, err error) {
	items := make([]ModelCapacityListResultValueInlined, 0)

	resp, err := c.List(ctx, id, options)
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

	result = ListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
