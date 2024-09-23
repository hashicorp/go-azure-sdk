package customauthenticationextension

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

type ListCustomAuthenticationExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CustomAuthenticationExtension
}

type ListCustomAuthenticationExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CustomAuthenticationExtension
}

type ListCustomAuthenticationExtensionsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListCustomAuthenticationExtensionsOperationOptions() ListCustomAuthenticationExtensionsOperationOptions {
	return ListCustomAuthenticationExtensionsOperationOptions{}
}

func (o ListCustomAuthenticationExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCustomAuthenticationExtensionsOperationOptions) ToOData() *odata.Query {
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

func (o ListCustomAuthenticationExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCustomAuthenticationExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCustomAuthenticationExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCustomAuthenticationExtensions - List customAuthenticationExtensions. Get a list of the
// customAuthenticationExtension objects and their properties. The following derived types are supported.
func (c CustomAuthenticationExtensionClient) ListCustomAuthenticationExtensions(ctx context.Context, options ListCustomAuthenticationExtensionsOperationOptions) (result ListCustomAuthenticationExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCustomAuthenticationExtensionsCustomPager{},
		Path:          "/identity/customAuthenticationExtensions",
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

	temp := make([]stable.CustomAuthenticationExtension, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalCustomAuthenticationExtensionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.CustomAuthenticationExtension (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListCustomAuthenticationExtensionsComplete retrieves all the results into a single object
func (c CustomAuthenticationExtensionClient) ListCustomAuthenticationExtensionsComplete(ctx context.Context, options ListCustomAuthenticationExtensionsOperationOptions) (ListCustomAuthenticationExtensionsCompleteResult, error) {
	return c.ListCustomAuthenticationExtensionsCompleteMatchingPredicate(ctx, options, CustomAuthenticationExtensionOperationPredicate{})
}

// ListCustomAuthenticationExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CustomAuthenticationExtensionClient) ListCustomAuthenticationExtensionsCompleteMatchingPredicate(ctx context.Context, options ListCustomAuthenticationExtensionsOperationOptions, predicate CustomAuthenticationExtensionOperationPredicate) (result ListCustomAuthenticationExtensionsCompleteResult, err error) {
	items := make([]stable.CustomAuthenticationExtension, 0)

	resp, err := c.ListCustomAuthenticationExtensions(ctx, options)
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

	result = ListCustomAuthenticationExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
