package appserviceenvironmentresources

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

type AppServiceEnvironmentsListCapacitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StampCapacity
}

type AppServiceEnvironmentsListCapacitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StampCapacity
}

type AppServiceEnvironmentsListCapacitiesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsListCapacitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsListCapacities ...
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListCapacities(ctx context.Context, id commonids.AppServiceEnvironmentId) (result AppServiceEnvironmentsListCapacitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AppServiceEnvironmentsListCapacitiesCustomPager{},
		Path:       fmt.Sprintf("%s/capacities/compute", id.ID()),
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
		Values *[]StampCapacity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AppServiceEnvironmentsListCapacitiesComplete retrieves all the results into a single object
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListCapacitiesComplete(ctx context.Context, id commonids.AppServiceEnvironmentId) (AppServiceEnvironmentsListCapacitiesCompleteResult, error) {
	return c.AppServiceEnvironmentsListCapacitiesCompleteMatchingPredicate(ctx, id, StampCapacityOperationPredicate{})
}

// AppServiceEnvironmentsListCapacitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListCapacitiesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, predicate StampCapacityOperationPredicate) (result AppServiceEnvironmentsListCapacitiesCompleteResult, err error) {
	items := make([]StampCapacity, 0)

	resp, err := c.AppServiceEnvironmentsListCapacities(ctx, id)
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

	result = AppServiceEnvironmentsListCapacitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
