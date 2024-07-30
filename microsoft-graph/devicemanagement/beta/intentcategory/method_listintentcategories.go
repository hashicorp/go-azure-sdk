package intentcategory

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListIntentCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementIntentSettingCategory
}

type ListIntentCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementIntentSettingCategory
}

type ListIntentCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntentCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntentCategories ...
func (c IntentCategoryClient) ListIntentCategories(ctx context.Context, id DeviceManagementIntentId) (result ListIntentCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIntentCategoriesCustomPager{},
		Path:       fmt.Sprintf("%s/categories", id.ID()),
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
		Values *[]beta.DeviceManagementIntentSettingCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIntentCategoriesComplete retrieves all the results into a single object
func (c IntentCategoryClient) ListIntentCategoriesComplete(ctx context.Context, id DeviceManagementIntentId) (ListIntentCategoriesCompleteResult, error) {
	return c.ListIntentCategoriesCompleteMatchingPredicate(ctx, id, DeviceManagementIntentSettingCategoryOperationPredicate{})
}

// ListIntentCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntentCategoryClient) ListIntentCategoriesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementIntentId, predicate DeviceManagementIntentSettingCategoryOperationPredicate) (result ListIntentCategoriesCompleteResult, err error) {
	items := make([]beta.DeviceManagementIntentSettingCategory, 0)

	resp, err := c.ListIntentCategories(ctx, id)
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

	result = ListIntentCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
