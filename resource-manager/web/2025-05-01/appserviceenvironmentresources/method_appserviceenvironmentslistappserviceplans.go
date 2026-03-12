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

type AppServiceEnvironmentsListAppServicePlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AppServicePlan
}

type AppServiceEnvironmentsListAppServicePlansCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AppServicePlan
}

type AppServiceEnvironmentsListAppServicePlansCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsListAppServicePlansCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsListAppServicePlans ...
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListAppServicePlans(ctx context.Context, id commonids.AppServiceEnvironmentId) (result AppServiceEnvironmentsListAppServicePlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AppServiceEnvironmentsListAppServicePlansCustomPager{},
		Path:       fmt.Sprintf("%s/serverFarms", id.ID()),
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
		Values *[]AppServicePlan `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AppServiceEnvironmentsListAppServicePlansComplete retrieves all the results into a single object
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListAppServicePlansComplete(ctx context.Context, id commonids.AppServiceEnvironmentId) (AppServiceEnvironmentsListAppServicePlansCompleteResult, error) {
	return c.AppServiceEnvironmentsListAppServicePlansCompleteMatchingPredicate(ctx, id, AppServicePlanOperationPredicate{})
}

// AppServiceEnvironmentsListAppServicePlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListAppServicePlansCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, predicate AppServicePlanOperationPredicate) (result AppServiceEnvironmentsListAppServicePlansCompleteResult, err error) {
	items := make([]AppServicePlan, 0)

	resp, err := c.AppServiceEnvironmentsListAppServicePlans(ctx, id)
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

	result = AppServiceEnvironmentsListAppServicePlansCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
