package contactfolderchildfolder

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

type ListContactFolderChildFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ContactFolder
}

type ListContactFolderChildFoldersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ContactFolder
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

// ListContactFolderChildFolders ...
func (c ContactFolderChildFolderClient) ListContactFolderChildFolders(ctx context.Context, id UserIdContactFolderId) (result ListContactFolderChildFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListContactFolderChildFoldersCustomPager{},
		Path:       fmt.Sprintf("%s/childFolders", id.ID()),
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
		Values *[]beta.ContactFolder `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListContactFolderChildFoldersComplete retrieves all the results into a single object
func (c ContactFolderChildFolderClient) ListContactFolderChildFoldersComplete(ctx context.Context, id UserIdContactFolderId) (ListContactFolderChildFoldersCompleteResult, error) {
	return c.ListContactFolderChildFoldersCompleteMatchingPredicate(ctx, id, ContactFolderOperationPredicate{})
}

// ListContactFolderChildFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderChildFolderClient) ListContactFolderChildFoldersCompleteMatchingPredicate(ctx context.Context, id UserIdContactFolderId, predicate ContactFolderOperationPredicate) (result ListContactFolderChildFoldersCompleteResult, err error) {
	items := make([]beta.ContactFolder, 0)

	resp, err := c.ListContactFolderChildFolders(ctx, id)
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
