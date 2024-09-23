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

type ListOutlookTaskFoldersOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListOutlookTaskFoldersOperationOptions() ListOutlookTaskFoldersOperationOptions {
	return ListOutlookTaskFoldersOperationOptions{}
}

func (o ListOutlookTaskFoldersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOutlookTaskFoldersOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListOutlookTaskFoldersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListOutlookTaskFolders - Get taskFolders from users. The user's Outlook task folders. Read-only. Nullable.
func (c OutlookTaskFolderClient) ListOutlookTaskFolders(ctx context.Context, id beta.UserId, options ListOutlookTaskFoldersOperationOptions) (result ListOutlookTaskFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOutlookTaskFoldersCustomPager{},
		Path:          fmt.Sprintf("%s/outlook/taskFolders", id.ID()),
		RetryFunc:     options.RetryFunc,
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
func (c OutlookTaskFolderClient) ListOutlookTaskFoldersComplete(ctx context.Context, id beta.UserId, options ListOutlookTaskFoldersOperationOptions) (ListOutlookTaskFoldersCompleteResult, error) {
	return c.ListOutlookTaskFoldersCompleteMatchingPredicate(ctx, id, options, OutlookTaskFolderOperationPredicate{})
}

// ListOutlookTaskFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookTaskFolderClient) ListOutlookTaskFoldersCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListOutlookTaskFoldersOperationOptions, predicate OutlookTaskFolderOperationPredicate) (result ListOutlookTaskFoldersCompleteResult, err error) {
	items := make([]beta.OutlookTaskFolder, 0)

	resp, err := c.ListOutlookTaskFolders(ctx, id, options)
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
