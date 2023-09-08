package meoutlooktask

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

type ListMeOutlookTaskByIdCompletesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTask
}

type ListMeOutlookTaskByIdCompletesCompleteResult struct {
	Items []models.OutlookTask
}

// ListMeOutlookTaskByIdCompletes ...
func (c MeOutlookTaskClient) ListMeOutlookTaskByIdCompletes(ctx context.Context, id MeOutlookTaskId) (result ListMeOutlookTaskByIdCompletesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/complete", id.ID()),
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
		Values *[]models.OutlookTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOutlookTaskByIdCompletesComplete retrieves all the results into a single object
func (c MeOutlookTaskClient) ListMeOutlookTaskByIdCompletesComplete(ctx context.Context, id MeOutlookTaskId) (ListMeOutlookTaskByIdCompletesCompleteResult, error) {
	return c.ListMeOutlookTaskByIdCompletesCompleteMatchingPredicate(ctx, id, models.OutlookTaskOperationPredicate{})
}

// ListMeOutlookTaskByIdCompletesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookTaskClient) ListMeOutlookTaskByIdCompletesCompleteMatchingPredicate(ctx context.Context, id MeOutlookTaskId, predicate models.OutlookTaskOperationPredicate) (result ListMeOutlookTaskByIdCompletesCompleteResult, err error) {
	items := make([]models.OutlookTask, 0)

	resp, err := c.ListMeOutlookTaskByIdCompletes(ctx, id)
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

	result = ListMeOutlookTaskByIdCompletesCompleteResult{
		Items: items,
	}
	return
}
