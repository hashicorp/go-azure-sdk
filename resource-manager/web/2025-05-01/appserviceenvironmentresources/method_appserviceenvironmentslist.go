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

type AppServiceEnvironmentsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AppServiceEnvironmentResource
}

type AppServiceEnvironmentsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AppServiceEnvironmentResource
}

type AppServiceEnvironmentsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsList ...
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsList(ctx context.Context, id commonids.SubscriptionId) (result AppServiceEnvironmentsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AppServiceEnvironmentsListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Web/hostingEnvironments", id.ID()),
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
		Values *[]AppServiceEnvironmentResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AppServiceEnvironmentsListComplete retrieves all the results into a single object
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListComplete(ctx context.Context, id commonids.SubscriptionId) (AppServiceEnvironmentsListCompleteResult, error) {
	return c.AppServiceEnvironmentsListCompleteMatchingPredicate(ctx, id, AppServiceEnvironmentResourceOperationPredicate{})
}

// AppServiceEnvironmentsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate AppServiceEnvironmentResourceOperationPredicate) (result AppServiceEnvironmentsListCompleteResult, err error) {
	items := make([]AppServiceEnvironmentResource, 0)

	resp, err := c.AppServiceEnvironmentsList(ctx, id)
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

	result = AppServiceEnvironmentsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
