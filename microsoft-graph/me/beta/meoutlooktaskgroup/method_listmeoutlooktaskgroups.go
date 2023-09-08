package meoutlooktaskgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeOutlookTaskGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskGroupCollectionResponse
}

type ListMeOutlookTaskGroupsCompleteResult struct {
	Items []models.OutlookTaskGroupCollectionResponse
}

// ListMeOutlookTaskGroups ...
func (c MeOutlookTaskGroupClient) ListMeOutlookTaskGroups(ctx context.Context) (result ListMeOutlookTaskGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/outlook/taskGroups",
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
		Values *[]models.OutlookTaskGroupCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOutlookTaskGroupsComplete retrieves all the results into a single object
func (c MeOutlookTaskGroupClient) ListMeOutlookTaskGroupsComplete(ctx context.Context) (ListMeOutlookTaskGroupsCompleteResult, error) {
	return c.ListMeOutlookTaskGroupsCompleteMatchingPredicate(ctx, models.OutlookTaskGroupCollectionResponseOperationPredicate{})
}

// ListMeOutlookTaskGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookTaskGroupClient) ListMeOutlookTaskGroupsCompleteMatchingPredicate(ctx context.Context, predicate models.OutlookTaskGroupCollectionResponseOperationPredicate) (result ListMeOutlookTaskGroupsCompleteResult, err error) {
	items := make([]models.OutlookTaskGroupCollectionResponse, 0)

	resp, err := c.ListMeOutlookTaskGroups(ctx)
	if err != nil {
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

	result = ListMeOutlookTaskGroupsCompleteResult{
		Items: items,
	}
	return
}
