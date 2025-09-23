package groupquotasentities

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

type GroupQuotasListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GroupQuotasEntity
}

type GroupQuotasListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GroupQuotasEntity
}

type GroupQuotasListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GroupQuotasListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GroupQuotasList ...
func (c GroupQuotasEntitiesClient) GroupQuotasList(ctx context.Context, id commonids.ManagementGroupId) (result GroupQuotasListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GroupQuotasListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Quota/groupQuotas", id.ID()),
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
		Values *[]GroupQuotasEntity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GroupQuotasListComplete retrieves all the results into a single object
func (c GroupQuotasEntitiesClient) GroupQuotasListComplete(ctx context.Context, id commonids.ManagementGroupId) (GroupQuotasListCompleteResult, error) {
	return c.GroupQuotasListCompleteMatchingPredicate(ctx, id, GroupQuotasEntityOperationPredicate{})
}

// GroupQuotasListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupQuotasEntitiesClient) GroupQuotasListCompleteMatchingPredicate(ctx context.Context, id commonids.ManagementGroupId, predicate GroupQuotasEntityOperationPredicate) (result GroupQuotasListCompleteResult, err error) {
	items := make([]GroupQuotasEntity, 0)

	resp, err := c.GroupQuotasList(ctx, id)
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

	result = GroupQuotasListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
