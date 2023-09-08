package meoutlooktaskgrouptaskfoldertask

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

type ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTask
}

type ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteResult struct {
	Items []models.OutlookTask
}

// ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletes ...
func (c MeOutlookTaskGroupTaskFolderTaskClient) ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletes(ctx context.Context, id MeOutlookTaskGroupTaskFolderTaskId) (result ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesOperationResponse, err error) {
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

// ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesComplete retrieves all the results into a single object
func (c MeOutlookTaskGroupTaskFolderTaskClient) ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesComplete(ctx context.Context, id MeOutlookTaskGroupTaskFolderTaskId) (ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteResult, error) {
	return c.ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate(ctx, id, models.OutlookTaskOperationPredicate{})
}

// ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookTaskGroupTaskFolderTaskClient) ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate(ctx context.Context, id MeOutlookTaskGroupTaskFolderTaskId, predicate models.OutlookTaskOperationPredicate) (result ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteResult, err error) {
	items := make([]models.OutlookTask, 0)

	resp, err := c.ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletes(ctx, id)
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

	result = ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteResult{
		Items: items,
	}
	return
}
