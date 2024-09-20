package contactfoldercontactextension

import (
	"context"
	"encoding/json"
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

type ListContactFolderContactExtensionsOperationOptions struct {
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

func DefaultListContactFolderContactExtensionsOperationOptions() ListContactFolderContactExtensionsOperationOptions {
	return ListContactFolderContactExtensionsOperationOptions{}
}

func (o ListContactFolderContactExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListContactFolderContactExtensionsOperationOptions) ToOData() *odata.Query {
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

func (o ListContactFolderContactExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListContactFolderContactExtensions - Get extensions from me. The collection of open extensions defined for the
// contact. Read-only. Nullable.
func (c ContactFolderContactExtensionClient) ListContactFolderContactExtensions(ctx context.Context, id stable.MeContactFolderIdContactId, options ListContactFolderContactExtensionsOperationOptions) (result ListContactFolderContactExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListContactFolderContactExtensionsCustomPager{},
		Path:          fmt.Sprintf("%s/extensions", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.Extension, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalExtensionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.Extension (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListContactFolderContactExtensionsComplete retrieves all the results into a single object
func (c ContactFolderContactExtensionClient) ListContactFolderContactExtensionsComplete(ctx context.Context, id stable.MeContactFolderIdContactId, options ListContactFolderContactExtensionsOperationOptions) (ListContactFolderContactExtensionsCompleteResult, error) {
	return c.ListContactFolderContactExtensionsCompleteMatchingPredicate(ctx, id, options, ExtensionOperationPredicate{})
}

// ListContactFolderContactExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderContactExtensionClient) ListContactFolderContactExtensionsCompleteMatchingPredicate(ctx context.Context, id stable.MeContactFolderIdContactId, options ListContactFolderContactExtensionsOperationOptions, predicate ExtensionOperationPredicate) (result ListContactFolderContactExtensionsCompleteResult, err error) {
	items := make([]stable.Extension, 0)

	resp, err := c.ListContactFolderContactExtensions(ctx, id, options)
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
