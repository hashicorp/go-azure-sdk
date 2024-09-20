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

type ListContactFolderChildFolderContactsOperationOptions struct {
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

func DefaultListContactFolderChildFolderContactsOperationOptions() ListContactFolderChildFolderContactsOperationOptions {
	return ListContactFolderChildFolderContactsOperationOptions{}
}

func (o ListContactFolderChildFolderContactsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListContactFolderChildFolderContactsOperationOptions) ToOData() *odata.Query {
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

func (o ListContactFolderChildFolderContactsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListContactFolderChildFolderContacts - Get contacts from users. The contacts in the folder. Navigation property.
// Read-only. Nullable.
func (c ContactFolderChildFolderContactClient) ListContactFolderChildFolderContacts(ctx context.Context, id beta.UserIdContactFolderIdChildFolderId, options ListContactFolderChildFolderContactsOperationOptions) (result ListContactFolderChildFolderContactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListContactFolderChildFolderContactsCustomPager{},
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

// ListContactFolderChildFolderContactsComplete retrieves all the results into a single object
func (c ContactFolderChildFolderContactClient) ListContactFolderChildFolderContactsComplete(ctx context.Context, id beta.UserIdContactFolderIdChildFolderId, options ListContactFolderChildFolderContactsOperationOptions) (ListContactFolderChildFolderContactsCompleteResult, error) {
	return c.ListContactFolderChildFolderContactsCompleteMatchingPredicate(ctx, id, options, ContactOperationPredicate{})
}

// ListContactFolderChildFolderContactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderChildFolderContactClient) ListContactFolderChildFolderContactsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdContactFolderIdChildFolderId, options ListContactFolderChildFolderContactsOperationOptions, predicate ContactOperationPredicate) (result ListContactFolderChildFolderContactsCompleteResult, err error) {
	items := make([]beta.Contact, 0)

	resp, err := c.ListContactFolderChildFolderContacts(ctx, id, options)
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
