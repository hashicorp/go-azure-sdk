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

type AppServiceEnvironmentsListMultiRolePoolsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkerPoolResource
}

type AppServiceEnvironmentsListMultiRolePoolsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkerPoolResource
}

type AppServiceEnvironmentsListMultiRolePoolsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsListMultiRolePoolsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsListMultiRolePools ...
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRolePools(ctx context.Context, id commonids.AppServiceEnvironmentId) (result AppServiceEnvironmentsListMultiRolePoolsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AppServiceEnvironmentsListMultiRolePoolsCustomPager{},
		Path:       fmt.Sprintf("%s/multiRolePools", id.ID()),
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
		Values *[]WorkerPoolResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AppServiceEnvironmentsListMultiRolePoolsComplete retrieves all the results into a single object
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRolePoolsComplete(ctx context.Context, id commonids.AppServiceEnvironmentId) (AppServiceEnvironmentsListMultiRolePoolsCompleteResult, error) {
	return c.AppServiceEnvironmentsListMultiRolePoolsCompleteMatchingPredicate(ctx, id, WorkerPoolResourceOperationPredicate{})
}

// AppServiceEnvironmentsListMultiRolePoolsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRolePoolsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, predicate WorkerPoolResourceOperationPredicate) (result AppServiceEnvironmentsListMultiRolePoolsCompleteResult, err error) {
	items := make([]WorkerPoolResource, 0)

	resp, err := c.AppServiceEnvironmentsListMultiRolePools(ctx, id)
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

	result = AppServiceEnvironmentsListMultiRolePoolsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
