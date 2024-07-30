package userexperienceanalyticscategory

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

type ListUserExperienceAnalyticsCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsCategory
}

type ListUserExperienceAnalyticsCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsCategory
}

type ListUserExperienceAnalyticsCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsCategories ...
func (c UserExperienceAnalyticsCategoryClient) ListUserExperienceAnalyticsCategories(ctx context.Context) (result ListUserExperienceAnalyticsCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsCategoriesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsCategories",
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
		Values *[]beta.UserExperienceAnalyticsCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsCategoriesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsCategoryClient) ListUserExperienceAnalyticsCategoriesComplete(ctx context.Context) (ListUserExperienceAnalyticsCategoriesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsCategoriesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsCategoryOperationPredicate{})
}

// ListUserExperienceAnalyticsCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsCategoryClient) ListUserExperienceAnalyticsCategoriesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsCategoryOperationPredicate) (result ListUserExperienceAnalyticsCategoriesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsCategory, 0)

	resp, err := c.ListUserExperienceAnalyticsCategories(ctx)
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

	result = ListUserExperienceAnalyticsCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
