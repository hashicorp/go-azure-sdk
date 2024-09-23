package b2xuserflowuserflowidentityprovider

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

type ListB2xUserFlowIdentityProviderRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListB2xUserFlowIdentityProviderRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListB2xUserFlowIdentityProviderRefsOperationOptions struct {
	Count     *bool
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Skip      *int64
	Top       *int64
}

func DefaultListB2xUserFlowIdentityProviderRefsOperationOptions() ListB2xUserFlowIdentityProviderRefsOperationOptions {
	return ListB2xUserFlowIdentityProviderRefsOperationOptions{}
}

func (o ListB2xUserFlowIdentityProviderRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListB2xUserFlowIdentityProviderRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListB2xUserFlowIdentityProviderRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListB2xUserFlowIdentityProviderRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowIdentityProviderRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowIdentityProviderRefs - Get ref of userFlowIdentityProviders from identity
func (c B2xUserFlowUserFlowIdentityProviderClient) ListB2xUserFlowIdentityProviderRefs(ctx context.Context, id stable.IdentityB2xUserFlowId, options ListB2xUserFlowIdentityProviderRefsOperationOptions) (result ListB2xUserFlowIdentityProviderRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListB2xUserFlowIdentityProviderRefsCustomPager{},
		Path:          fmt.Sprintf("%s/userFlowIdentityProviders/$ref", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	temp := make([]stable.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListB2xUserFlowIdentityProviderRefsComplete retrieves all the results into a single object
func (c B2xUserFlowUserFlowIdentityProviderClient) ListB2xUserFlowIdentityProviderRefsComplete(ctx context.Context, id stable.IdentityB2xUserFlowId, options ListB2xUserFlowIdentityProviderRefsOperationOptions) (ListB2xUserFlowIdentityProviderRefsCompleteResult, error) {
	return c.ListB2xUserFlowIdentityProviderRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListB2xUserFlowIdentityProviderRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowUserFlowIdentityProviderClient) ListB2xUserFlowIdentityProviderRefsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityB2xUserFlowId, options ListB2xUserFlowIdentityProviderRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListB2xUserFlowIdentityProviderRefsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListB2xUserFlowIdentityProviderRefs(ctx, id, options)
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

	result = ListB2xUserFlowIdentityProviderRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
