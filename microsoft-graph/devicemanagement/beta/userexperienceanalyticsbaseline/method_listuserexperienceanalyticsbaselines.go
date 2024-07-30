package userexperienceanalyticsbaseline

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

type ListUserExperienceAnalyticsBaselinesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsBaseline
}

type ListUserExperienceAnalyticsBaselinesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsBaseline
}

type ListUserExperienceAnalyticsBaselinesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsBaselinesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsBaselines ...
func (c UserExperienceAnalyticsBaselineClient) ListUserExperienceAnalyticsBaselines(ctx context.Context) (result ListUserExperienceAnalyticsBaselinesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsBaselinesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsBaselines",
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
		Values *[]beta.UserExperienceAnalyticsBaseline `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsBaselinesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsBaselineClient) ListUserExperienceAnalyticsBaselinesComplete(ctx context.Context) (ListUserExperienceAnalyticsBaselinesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsBaselinesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsBaselineOperationPredicate{})
}

// ListUserExperienceAnalyticsBaselinesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsBaselineClient) ListUserExperienceAnalyticsBaselinesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsBaselineOperationPredicate) (result ListUserExperienceAnalyticsBaselinesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsBaseline, 0)

	resp, err := c.ListUserExperienceAnalyticsBaselines(ctx)
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

	result = ListUserExperienceAnalyticsBaselinesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
