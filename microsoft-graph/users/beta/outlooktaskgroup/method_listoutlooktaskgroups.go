package outlooktaskgroup

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

type ListOutlookTaskGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutlookTaskGroup
}

type ListOutlookTaskGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutlookTaskGroup
}

type ListOutlookTaskGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookTaskGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookTaskGroups ...
func (c OutlookTaskGroupClient) ListOutlookTaskGroups(ctx context.Context, id UserId) (result ListOutlookTaskGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOutlookTaskGroupsCustomPager{},
		Path:       fmt.Sprintf("%s/outlook/taskGroups", id.ID()),
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
		Values *[]beta.OutlookTaskGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutlookTaskGroupsComplete retrieves all the results into a single object
func (c OutlookTaskGroupClient) ListOutlookTaskGroupsComplete(ctx context.Context, id UserId) (ListOutlookTaskGroupsCompleteResult, error) {
	return c.ListOutlookTaskGroupsCompleteMatchingPredicate(ctx, id, OutlookTaskGroupOperationPredicate{})
}

// ListOutlookTaskGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookTaskGroupClient) ListOutlookTaskGroupsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate OutlookTaskGroupOperationPredicate) (result ListOutlookTaskGroupsCompleteResult, err error) {
	items := make([]beta.OutlookTaskGroup, 0)

	resp, err := c.ListOutlookTaskGroups(ctx, id)
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

	result = ListOutlookTaskGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
