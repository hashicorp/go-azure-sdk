package activity

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

type ListActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserActivity
}

type ListActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserActivity
}

type ListActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListActivities ...
func (c ActivityClient) ListActivities(ctx context.Context) (result ListActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListActivitiesCustomPager{},
		Path:       "/me/activities",
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
		Values *[]beta.UserActivity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListActivitiesComplete retrieves all the results into a single object
func (c ActivityClient) ListActivitiesComplete(ctx context.Context) (ListActivitiesCompleteResult, error) {
	return c.ListActivitiesCompleteMatchingPredicate(ctx, UserActivityOperationPredicate{})
}

// ListActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ActivityClient) ListActivitiesCompleteMatchingPredicate(ctx context.Context, predicate UserActivityOperationPredicate) (result ListActivitiesCompleteResult, err error) {
	items := make([]beta.UserActivity, 0)

	resp, err := c.ListActivities(ctx)
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

	result = ListActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
