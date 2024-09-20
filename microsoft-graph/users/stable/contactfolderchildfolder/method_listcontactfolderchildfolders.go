package contactfolderchildfolder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListContactFolderChildFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ContactFolder
}

type ListContactFolderChildFoldersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ContactFolder
}

type ListContactFolderChildFoldersOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListContactFolderChildFoldersOperationOptions() ListContactFolderChildFoldersOperationOptions {
	return ListContactFolderChildFoldersOperationOptions{}
}

func (o ListContactFolderChildFoldersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListContactFolderChildFoldersOperationOptions) ToOData() *odata.Query {
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

func (o ListContactFolderChildFoldersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListContactFolderChildFoldersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListContactFolderChildFoldersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListContactFolderChildFolders - Get childFolders from users. The collection of child folders in the folder.
// Navigation property. Read-only. Nullable.
func (c ContactFolderChildFolderClient) ListContactFolderChildFolders(ctx context.Context, id stable.UserIdContactFolderId, options ListContactFolderChildFoldersOperationOptions) (result ListContactFolderChildFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListContactFolderChildFoldersCustomPager{},
		Path:          fmt.Sprintf("%s/childFolders", id.ID()),
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
		Values *[]stable.ContactFolder `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListContactFolderChildFoldersComplete retrieves all the results into a single object
func (c ContactFolderChildFolderClient) ListContactFolderChildFoldersComplete(ctx context.Context, id stable.UserIdContactFolderId, options ListContactFolderChildFoldersOperationOptions) (ListContactFolderChildFoldersCompleteResult, error) {
	return c.ListContactFolderChildFoldersCompleteMatchingPredicate(ctx, id, options, ContactFolderOperationPredicate{})
}

// ListContactFolderChildFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderChildFolderClient) ListContactFolderChildFoldersCompleteMatchingPredicate(ctx context.Context, id stable.UserIdContactFolderId, options ListContactFolderChildFoldersOperationOptions, predicate ContactFolderOperationPredicate) (result ListContactFolderChildFoldersCompleteResult, err error) {
	items := make([]stable.ContactFolder, 0)

	resp, err := c.ListContactFolderChildFolders(ctx, id, options)
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

	result = ListContactFolderChildFoldersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
