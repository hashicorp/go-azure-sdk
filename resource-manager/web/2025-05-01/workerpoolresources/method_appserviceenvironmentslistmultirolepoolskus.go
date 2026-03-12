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

type AppServiceEnvironmentsListMultiRolePoolSkusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SkuInfo
}

type AppServiceEnvironmentsListMultiRolePoolSkusCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SkuInfo
}

type AppServiceEnvironmentsListMultiRolePoolSkusCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsListMultiRolePoolSkusCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsListMultiRolePoolSkus ...
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRolePoolSkus(ctx context.Context, id commonids.AppServiceEnvironmentId) (result AppServiceEnvironmentsListMultiRolePoolSkusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AppServiceEnvironmentsListMultiRolePoolSkusCustomPager{},
		Path:       fmt.Sprintf("%s/multiRolePools/default/skus", id.ID()),
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
		Values *[]SkuInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AppServiceEnvironmentsListMultiRolePoolSkusComplete retrieves all the results into a single object
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRolePoolSkusComplete(ctx context.Context, id commonids.AppServiceEnvironmentId) (AppServiceEnvironmentsListMultiRolePoolSkusCompleteResult, error) {
	return c.AppServiceEnvironmentsListMultiRolePoolSkusCompleteMatchingPredicate(ctx, id, SkuInfoOperationPredicate{})
}

// AppServiceEnvironmentsListMultiRolePoolSkusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRolePoolSkusCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, predicate SkuInfoOperationPredicate) (result AppServiceEnvironmentsListMultiRolePoolSkusCompleteResult, err error) {
	items := make([]SkuInfo, 0)

	resp, err := c.AppServiceEnvironmentsListMultiRolePoolSkus(ctx, id)
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

	result = AppServiceEnvironmentsListMultiRolePoolSkusCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
