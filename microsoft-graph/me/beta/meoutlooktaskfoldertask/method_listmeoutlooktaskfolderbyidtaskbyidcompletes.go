package meoutlooktaskfoldertask

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

type ListMeOutlookTaskFolderByIdTaskByIdCompletesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTask
}

type ListMeOutlookTaskFolderByIdTaskByIdCompletesCompleteResult struct {
	Items []models.OutlookTask
}

// ListMeOutlookTaskFolderByIdTaskByIdCompletes ...
func (c MeOutlookTaskFolderTaskClient) ListMeOutlookTaskFolderByIdTaskByIdCompletes(ctx context.Context, id MeOutlookTaskFolderTaskId) (result ListMeOutlookTaskFolderByIdTaskByIdCompletesOperationResponse, err error) {
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

// ListMeOutlookTaskFolderByIdTaskByIdCompletesComplete retrieves all the results into a single object
func (c MeOutlookTaskFolderTaskClient) ListMeOutlookTaskFolderByIdTaskByIdCompletesComplete(ctx context.Context, id MeOutlookTaskFolderTaskId) (ListMeOutlookTaskFolderByIdTaskByIdCompletesCompleteResult, error) {
	return c.ListMeOutlookTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate(ctx, id, models.OutlookTaskOperationPredicate{})
}

// ListMeOutlookTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookTaskFolderTaskClient) ListMeOutlookTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate(ctx context.Context, id MeOutlookTaskFolderTaskId, predicate models.OutlookTaskOperationPredicate) (result ListMeOutlookTaskFolderByIdTaskByIdCompletesCompleteResult, err error) {
	items := make([]models.OutlookTask, 0)

	resp, err := c.ListMeOutlookTaskFolderByIdTaskByIdCompletes(ctx, id)
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

	result = ListMeOutlookTaskFolderByIdTaskByIdCompletesCompleteResult{
		Items: items,
	}
	return
}
