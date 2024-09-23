package permissiongrant

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

type ListPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ResourceSpecificPermissionGrant
}

type ListPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ResourceSpecificPermissionGrant
}

type ListPermissionGrantsOperationOptions struct {
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

func DefaultListPermissionGrantsOperationOptions() ListPermissionGrantsOperationOptions {
	return ListPermissionGrantsOperationOptions{}
}

func (o ListPermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPermissionGrantsOperationOptions) ToOData() *odata.Query {
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

func (o ListPermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionGrants - Get permissionGrants from me. List all resource-specific permission grants of a user.
func (c PermissionGrantClient) ListPermissionGrants(ctx context.Context, options ListPermissionGrantsOperationOptions) (result ListPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPermissionGrantsCustomPager{},
		Path:          "/me/permissionGrants",
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
		Values *[]beta.ResourceSpecificPermissionGrant `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPermissionGrantsComplete retrieves all the results into a single object
func (c PermissionGrantClient) ListPermissionGrantsComplete(ctx context.Context, options ListPermissionGrantsOperationOptions) (ListPermissionGrantsCompleteResult, error) {
	return c.ListPermissionGrantsCompleteMatchingPredicate(ctx, options, ResourceSpecificPermissionGrantOperationPredicate{})
}

// ListPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionGrantClient) ListPermissionGrantsCompleteMatchingPredicate(ctx context.Context, options ListPermissionGrantsOperationOptions, predicate ResourceSpecificPermissionGrantOperationPredicate) (result ListPermissionGrantsCompleteResult, err error) {
	items := make([]beta.ResourceSpecificPermissionGrant, 0)

	resp, err := c.ListPermissionGrants(ctx, options)
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

	result = ListPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
