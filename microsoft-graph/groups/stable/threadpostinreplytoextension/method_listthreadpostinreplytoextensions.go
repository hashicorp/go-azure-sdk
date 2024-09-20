package threadpostinreplytoextension

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

type ListThreadPostInReplyToExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Extension
}

type ListThreadPostInReplyToExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Extension
}

type ListThreadPostInReplyToExtensionsOperationOptions struct {
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

func DefaultListThreadPostInReplyToExtensionsOperationOptions() ListThreadPostInReplyToExtensionsOperationOptions {
	return ListThreadPostInReplyToExtensionsOperationOptions{}
}

func (o ListThreadPostInReplyToExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListThreadPostInReplyToExtensionsOperationOptions) ToOData() *odata.Query {
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

func (o ListThreadPostInReplyToExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListThreadPostInReplyToExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListThreadPostInReplyToExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListThreadPostInReplyToExtensions - Get extensions from groups. The collection of open extensions defined for the
// post. Read-only. Nullable. Supports $expand.
func (c ThreadPostInReplyToExtensionClient) ListThreadPostInReplyToExtensions(ctx context.Context, id stable.GroupIdThreadIdPostId, options ListThreadPostInReplyToExtensionsOperationOptions) (result ListThreadPostInReplyToExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListThreadPostInReplyToExtensionsCustomPager{},
		Path:          fmt.Sprintf("%s/inReplyTo/extensions", id.ID()),
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

// ListThreadPostInReplyToExtensionsComplete retrieves all the results into a single object
func (c ThreadPostInReplyToExtensionClient) ListThreadPostInReplyToExtensionsComplete(ctx context.Context, id stable.GroupIdThreadIdPostId, options ListThreadPostInReplyToExtensionsOperationOptions) (ListThreadPostInReplyToExtensionsCompleteResult, error) {
	return c.ListThreadPostInReplyToExtensionsCompleteMatchingPredicate(ctx, id, options, ExtensionOperationPredicate{})
}

// ListThreadPostInReplyToExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ThreadPostInReplyToExtensionClient) ListThreadPostInReplyToExtensionsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdThreadIdPostId, options ListThreadPostInReplyToExtensionsOperationOptions, predicate ExtensionOperationPredicate) (result ListThreadPostInReplyToExtensionsCompleteResult, err error) {
	items := make([]stable.Extension, 0)

	resp, err := c.ListThreadPostInReplyToExtensions(ctx, id, options)
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

	result = ListThreadPostInReplyToExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
