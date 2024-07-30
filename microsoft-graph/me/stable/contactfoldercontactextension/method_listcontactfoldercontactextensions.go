package contactfoldercontactextension

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

type ListContactFolderContactExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Extension
}

type ListContactFolderContactExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Extension
}

type ListContactFolderContactExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListContactFolderContactExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListContactFolderContactExtensions ...
func (c ContactFolderContactExtensionClient) ListContactFolderContactExtensions(ctx context.Context, id MeContactFolderIdContactId) (result ListContactFolderContactExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListContactFolderContactExtensionsCustomPager{},
		Path:       fmt.Sprintf("%s/extensions", id.ID()),
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
		Values *[]stable.Extension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListContactFolderContactExtensionsComplete retrieves all the results into a single object
func (c ContactFolderContactExtensionClient) ListContactFolderContactExtensionsComplete(ctx context.Context, id MeContactFolderIdContactId) (ListContactFolderContactExtensionsCompleteResult, error) {
	return c.ListContactFolderContactExtensionsCompleteMatchingPredicate(ctx, id, ExtensionOperationPredicate{})
}

// ListContactFolderContactExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderContactExtensionClient) ListContactFolderContactExtensionsCompleteMatchingPredicate(ctx context.Context, id MeContactFolderIdContactId, predicate ExtensionOperationPredicate) (result ListContactFolderContactExtensionsCompleteResult, err error) {
	items := make([]stable.Extension, 0)

	resp, err := c.ListContactFolderContactExtensions(ctx, id)
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

	result = ListContactFolderContactExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
