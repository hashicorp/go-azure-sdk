package contactfoldercontact

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

type ListContactFolderContactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Contact
}

type ListContactFolderContactsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Contact
}

type ListContactFolderContactsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListContactFolderContactsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListContactFolderContacts ...
func (c ContactFolderContactClient) ListContactFolderContacts(ctx context.Context, id MeContactFolderId) (result ListContactFolderContactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListContactFolderContactsCustomPager{},
		Path:       fmt.Sprintf("%s/contacts", id.ID()),
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
		Values *[]stable.Contact `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListContactFolderContactsComplete retrieves all the results into a single object
func (c ContactFolderContactClient) ListContactFolderContactsComplete(ctx context.Context, id MeContactFolderId) (ListContactFolderContactsCompleteResult, error) {
	return c.ListContactFolderContactsCompleteMatchingPredicate(ctx, id, ContactOperationPredicate{})
}

// ListContactFolderContactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderContactClient) ListContactFolderContactsCompleteMatchingPredicate(ctx context.Context, id MeContactFolderId, predicate ContactOperationPredicate) (result ListContactFolderContactsCompleteResult, err error) {
	items := make([]stable.Contact, 0)

	resp, err := c.ListContactFolderContacts(ctx, id)
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

	result = ListContactFolderContactsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
