package contactfoldercontact

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

type ListContactFolderContactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Contact
}

type ListContactFolderContactsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Contact
}

type ListContactFolderContactsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListContactFolderContactsOperationOptions() ListContactFolderContactsOperationOptions {
	return ListContactFolderContactsOperationOptions{}
}

func (o ListContactFolderContactsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListContactFolderContactsOperationOptions) ToOData() *odata.Query {
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

func (o ListContactFolderContactsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListContactFolderContacts - List contacts. Get all the contacts in the signed-in user's mailbox (.../me/contacts), or
// from the specified contact folder.
func (c ContactFolderContactClient) ListContactFolderContacts(ctx context.Context, id beta.MeContactFolderId, options ListContactFolderContactsOperationOptions) (result ListContactFolderContactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListContactFolderContactsCustomPager{},
		Path:          fmt.Sprintf("%s/contacts", id.ID()),
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

// ListContactFolderContactsComplete retrieves all the results into a single object
func (c ContactFolderContactClient) ListContactFolderContactsComplete(ctx context.Context, id beta.MeContactFolderId, options ListContactFolderContactsOperationOptions) (ListContactFolderContactsCompleteResult, error) {
	return c.ListContactFolderContactsCompleteMatchingPredicate(ctx, id, options, ContactOperationPredicate{})
}

// ListContactFolderContactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderContactClient) ListContactFolderContactsCompleteMatchingPredicate(ctx context.Context, id beta.MeContactFolderId, options ListContactFolderContactsOperationOptions, predicate ContactOperationPredicate) (result ListContactFolderContactsCompleteResult, err error) {
	items := make([]beta.Contact, 0)

	resp, err := c.ListContactFolderContacts(ctx, id, options)
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
