package workerpoolresources

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

type AppServiceEnvironmentsListMultiRoleUsagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Usage
}

type AppServiceEnvironmentsListMultiRoleUsagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Usage
}

type AppServiceEnvironmentsListMultiRoleUsagesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsListMultiRoleUsagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsListMultiRoleUsages ...
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRoleUsages(ctx context.Context, id commonids.AppServiceEnvironmentId) (result AppServiceEnvironmentsListMultiRoleUsagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AppServiceEnvironmentsListMultiRoleUsagesCustomPager{},
		Path:       fmt.Sprintf("%s/multiRolePools/default/usages", id.ID()),
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
		Values *[]Usage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AppServiceEnvironmentsListMultiRoleUsagesComplete retrieves all the results into a single object
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRoleUsagesComplete(ctx context.Context, id commonids.AppServiceEnvironmentId) (AppServiceEnvironmentsListMultiRoleUsagesCompleteResult, error) {
	return c.AppServiceEnvironmentsListMultiRoleUsagesCompleteMatchingPredicate(ctx, id, UsageOperationPredicate{})
}

// AppServiceEnvironmentsListMultiRoleUsagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRoleUsagesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, predicate UsageOperationPredicate) (result AppServiceEnvironmentsListMultiRoleUsagesCompleteResult, err error) {
	items := make([]Usage, 0)

	resp, err := c.AppServiceEnvironmentsListMultiRoleUsages(ctx, id)
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

	result = AppServiceEnvironmentsListMultiRoleUsagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
