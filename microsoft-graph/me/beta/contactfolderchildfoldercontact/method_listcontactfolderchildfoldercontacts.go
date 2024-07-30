package contactfolderchildfoldercontact

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

type ListContactFolderChildFolderContactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Contact
}

type ListContactFolderChildFolderContactsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Contact
}

type ListContactFolderChildFolderContactsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListContactFolderChildFolderContactsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListContactFolderChildFolderContacts ...
func (c ContactFolderChildFolderContactClient) ListContactFolderChildFolderContacts(ctx context.Context, id MeContactFolderIdChildFolderId) (result ListContactFolderChildFolderContactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListContactFolderChildFolderContactsCustomPager{},
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
		Values *[]beta.Contact `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListContactFolderChildFolderContactsComplete retrieves all the results into a single object
func (c ContactFolderChildFolderContactClient) ListContactFolderChildFolderContactsComplete(ctx context.Context, id MeContactFolderIdChildFolderId) (ListContactFolderChildFolderContactsCompleteResult, error) {
	return c.ListContactFolderChildFolderContactsCompleteMatchingPredicate(ctx, id, ContactOperationPredicate{})
}

// ListContactFolderChildFolderContactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderChildFolderContactClient) ListContactFolderChildFolderContactsCompleteMatchingPredicate(ctx context.Context, id MeContactFolderIdChildFolderId, predicate ContactOperationPredicate) (result ListContactFolderChildFolderContactsCompleteResult, err error) {
	items := make([]beta.Contact, 0)

	resp, err := c.ListContactFolderChildFolderContacts(ctx, id)
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

	result = ListContactFolderChildFolderContactsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
