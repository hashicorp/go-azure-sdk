package modelcapacities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationBasedModelCapacitiesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ModelCapacityListResultValueInlined
}

type LocationBasedModelCapacitiesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ModelCapacityListResultValueInlined
}

type LocationBasedModelCapacitiesListOperationOptions struct {
	ModelFormat  *string
	ModelName    *string
	ModelVersion *string
}

func DefaultLocationBasedModelCapacitiesListOperationOptions() LocationBasedModelCapacitiesListOperationOptions {
	return LocationBasedModelCapacitiesListOperationOptions{}
}

func (o LocationBasedModelCapacitiesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o LocationBasedModelCapacitiesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o LocationBasedModelCapacitiesListOperationOptions) ToQuery() *client.QueryParams {
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

type LocationBasedModelCapacitiesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *LocationBasedModelCapacitiesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// LocationBasedModelCapacitiesList ...
func (c ModelCapacitiesClient) LocationBasedModelCapacitiesList(ctx context.Context, id LocationId, options LocationBasedModelCapacitiesListOperationOptions) (result LocationBasedModelCapacitiesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &LocationBasedModelCapacitiesListCustomPager{},
		Path:          fmt.Sprintf("%s/modelCapacities", id.ID()),
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

// LocationBasedModelCapacitiesListComplete retrieves all the results into a single object
func (c ModelCapacitiesClient) LocationBasedModelCapacitiesListComplete(ctx context.Context, id LocationId, options LocationBasedModelCapacitiesListOperationOptions) (LocationBasedModelCapacitiesListCompleteResult, error) {
	return c.LocationBasedModelCapacitiesListCompleteMatchingPredicate(ctx, id, options, ModelCapacityListResultValueInlinedOperationPredicate{})
}

// LocationBasedModelCapacitiesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ModelCapacitiesClient) LocationBasedModelCapacitiesListCompleteMatchingPredicate(ctx context.Context, id LocationId, options LocationBasedModelCapacitiesListOperationOptions, predicate ModelCapacityListResultValueInlinedOperationPredicate) (result LocationBasedModelCapacitiesListCompleteResult, err error) {
	items := make([]ModelCapacityListResultValueInlined, 0)

	resp, err := c.LocationBasedModelCapacitiesList(ctx, id, options)
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

	result = LocationBasedModelCapacitiesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
