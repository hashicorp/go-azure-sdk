package outlooktaskfolder

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

type ListOutlookTaskFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutlookTaskFolder
}

type ListOutlookTaskFoldersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutlookTaskFolder
}

type ListOutlookTaskFoldersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookTaskFoldersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookTaskFolders ...
func (c OutlookTaskFolderClient) ListOutlookTaskFolders(ctx context.Context) (result ListOutlookTaskFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOutlookTaskFoldersCustomPager{},
		Path:       "/me/outlook/taskFolders",
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
		Values *[]beta.OutlookTaskFolder `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutlookTaskFoldersComplete retrieves all the results into a single object
func (c OutlookTaskFolderClient) ListOutlookTaskFoldersComplete(ctx context.Context) (ListOutlookTaskFoldersCompleteResult, error) {
	return c.ListOutlookTaskFoldersCompleteMatchingPredicate(ctx, OutlookTaskFolderOperationPredicate{})
}

// ListOutlookTaskFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookTaskFolderClient) ListOutlookTaskFoldersCompleteMatchingPredicate(ctx context.Context, predicate OutlookTaskFolderOperationPredicate) (result ListOutlookTaskFoldersCompleteResult, err error) {
	items := make([]beta.OutlookTaskFolder, 0)

	resp, err := c.ListOutlookTaskFolders(ctx)
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

	result = ListOutlookTaskFoldersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
