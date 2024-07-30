package outlooktaskgrouptaskfoldertask

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

type ListOutlookTaskGroupTaskFolderTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutlookTask
}

type ListOutlookTaskGroupTaskFolderTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutlookTask
}

type ListOutlookTaskGroupTaskFolderTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookTaskGroupTaskFolderTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookTaskGroupTaskFolderTasks ...
func (c OutlookTaskGroupTaskFolderTaskClient) ListOutlookTaskGroupTaskFolderTasks(ctx context.Context, id UserIdOutlookTaskGroupIdTaskFolderId) (result ListOutlookTaskGroupTaskFolderTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOutlookTaskGroupTaskFolderTasksCustomPager{},
		Path:       fmt.Sprintf("%s/tasks", id.ID()),
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
		Values *[]beta.OutlookTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutlookTaskGroupTaskFolderTasksComplete retrieves all the results into a single object
func (c OutlookTaskGroupTaskFolderTaskClient) ListOutlookTaskGroupTaskFolderTasksComplete(ctx context.Context, id UserIdOutlookTaskGroupIdTaskFolderId) (ListOutlookTaskGroupTaskFolderTasksCompleteResult, error) {
	return c.ListOutlookTaskGroupTaskFolderTasksCompleteMatchingPredicate(ctx, id, OutlookTaskOperationPredicate{})
}

// ListOutlookTaskGroupTaskFolderTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookTaskGroupTaskFolderTaskClient) ListOutlookTaskGroupTaskFolderTasksCompleteMatchingPredicate(ctx context.Context, id UserIdOutlookTaskGroupIdTaskFolderId, predicate OutlookTaskOperationPredicate) (result ListOutlookTaskGroupTaskFolderTasksCompleteResult, err error) {
	items := make([]beta.OutlookTask, 0)

	resp, err := c.ListOutlookTaskGroupTaskFolderTasks(ctx, id)
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

	result = ListOutlookTaskGroupTaskFolderTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
