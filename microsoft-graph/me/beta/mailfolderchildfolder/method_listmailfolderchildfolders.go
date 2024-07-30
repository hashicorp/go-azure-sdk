package mailfolderchildfolder

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

type ListMailFolderChildFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MailFolder
}

type ListMailFolderChildFoldersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MailFolder
}

type ListMailFolderChildFoldersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderChildFoldersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderChildFolders ...
func (c MailFolderChildFolderClient) ListMailFolderChildFolders(ctx context.Context, id MeMailFolderId) (result ListMailFolderChildFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMailFolderChildFoldersCustomPager{},
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
		Values *[]beta.MailFolder `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMailFolderChildFoldersComplete retrieves all the results into a single object
func (c MailFolderChildFolderClient) ListMailFolderChildFoldersComplete(ctx context.Context, id MeMailFolderId) (ListMailFolderChildFoldersCompleteResult, error) {
	return c.ListMailFolderChildFoldersCompleteMatchingPredicate(ctx, id, MailFolderOperationPredicate{})
}

// ListMailFolderChildFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderChildFolderClient) ListMailFolderChildFoldersCompleteMatchingPredicate(ctx context.Context, id MeMailFolderId, predicate MailFolderOperationPredicate) (result ListMailFolderChildFoldersCompleteResult, err error) {
	items := make([]beta.MailFolder, 0)

	resp, err := c.ListMailFolderChildFolders(ctx, id)
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

	result = ListMailFolderChildFoldersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
