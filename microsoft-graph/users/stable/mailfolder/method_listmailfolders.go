package mailfolder

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

type ListMailFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.MailFolder
}

type ListMailFoldersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.MailFolder
}

type ListMailFoldersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFoldersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolders ...
func (c MailFolderClient) ListMailFolders(ctx context.Context, id UserId) (result ListMailFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMailFoldersCustomPager{},
		Path:       fmt.Sprintf("%s/mailFolders", id.ID()),
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
		Values *[]stable.MailFolder `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMailFoldersComplete retrieves all the results into a single object
func (c MailFolderClient) ListMailFoldersComplete(ctx context.Context, id UserId) (ListMailFoldersCompleteResult, error) {
	return c.ListMailFoldersCompleteMatchingPredicate(ctx, id, MailFolderOperationPredicate{})
}

// ListMailFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderClient) ListMailFoldersCompleteMatchingPredicate(ctx context.Context, id UserId, predicate MailFolderOperationPredicate) (result ListMailFoldersCompleteResult, err error) {
	items := make([]stable.MailFolder, 0)

	resp, err := c.ListMailFolders(ctx, id)
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

	result = ListMailFoldersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
