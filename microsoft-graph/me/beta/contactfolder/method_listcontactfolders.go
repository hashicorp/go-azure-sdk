package contactfolder

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

type ListContactFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ContactFolder
}

type ListContactFoldersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ContactFolder
}

type ListContactFoldersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListContactFoldersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListContactFolders ...
func (c ContactFolderClient) ListContactFolders(ctx context.Context) (result ListContactFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListContactFoldersCustomPager{},
		Path:       "/me/contactFolders",
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

// ListContactFoldersComplete retrieves all the results into a single object
func (c ContactFolderClient) ListContactFoldersComplete(ctx context.Context) (ListContactFoldersCompleteResult, error) {
	return c.ListContactFoldersCompleteMatchingPredicate(ctx, ContactFolderOperationPredicate{})
}

// ListContactFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderClient) ListContactFoldersCompleteMatchingPredicate(ctx context.Context, predicate ContactFolderOperationPredicate) (result ListContactFoldersCompleteResult, err error) {
	items := make([]beta.ContactFolder, 0)

	resp, err := c.ListContactFolders(ctx)
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

	result = ListContactFoldersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
