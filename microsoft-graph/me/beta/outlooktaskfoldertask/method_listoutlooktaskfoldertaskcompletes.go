package outlooktaskfoldertask

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

type ListOutlookTaskFolderTaskCompletesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutlookTask
}

type ListOutlookTaskFolderTaskCompletesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutlookTask
}

type ListOutlookTaskFolderTaskCompletesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookTaskFolderTaskCompletesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookTaskFolderTaskCompletes ...
func (c OutlookTaskFolderTaskClient) ListOutlookTaskFolderTaskCompletes(ctx context.Context, id MeOutlookTaskFolderIdTaskId) (result ListOutlookTaskFolderTaskCompletesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListOutlookTaskFolderTaskCompletesCustomPager{},
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
		Values *[]beta.OutlookTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutlookTaskFolderTaskCompletesComplete retrieves all the results into a single object
func (c OutlookTaskFolderTaskClient) ListOutlookTaskFolderTaskCompletesComplete(ctx context.Context, id MeOutlookTaskFolderIdTaskId) (ListOutlookTaskFolderTaskCompletesCompleteResult, error) {
	return c.ListOutlookTaskFolderTaskCompletesCompleteMatchingPredicate(ctx, id, OutlookTaskOperationPredicate{})
}

// ListOutlookTaskFolderTaskCompletesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookTaskFolderTaskClient) ListOutlookTaskFolderTaskCompletesCompleteMatchingPredicate(ctx context.Context, id MeOutlookTaskFolderIdTaskId, predicate OutlookTaskOperationPredicate) (result ListOutlookTaskFolderTaskCompletesCompleteResult, err error) {
	items := make([]beta.OutlookTask, 0)

	resp, err := c.ListOutlookTaskFolderTaskCompletes(ctx, id)
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

	result = ListOutlookTaskFolderTaskCompletesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
