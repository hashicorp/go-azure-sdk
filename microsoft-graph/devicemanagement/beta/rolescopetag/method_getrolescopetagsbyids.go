package rolescopetag

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

type GetRoleScopeTagsByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RoleScopeTag
}

type GetRoleScopeTagsByIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RoleScopeTag
}

type GetRoleScopeTagsByIdsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultGetRoleScopeTagsByIdsOperationOptions() GetRoleScopeTagsByIdsOperationOptions {
	return GetRoleScopeTagsByIdsOperationOptions{}
}

func (o GetRoleScopeTagsByIdsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetRoleScopeTagsByIdsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetRoleScopeTagsByIdsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetRoleScopeTagsByIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetRoleScopeTagsByIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetRoleScopeTagsByIds - Invoke action getRoleScopeTagsById
func (c RoleScopeTagClient) GetRoleScopeTagsByIds(ctx context.Context, input GetRoleScopeTagsByIdsRequest, options GetRoleScopeTagsByIdsOperationOptions) (result GetRoleScopeTagsByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetRoleScopeTagsByIdsCustomPager{},
		Path:          "/deviceManagement/roleScopeTags/getRoleScopeTagsById",
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
		Values *[]beta.RoleScopeTag `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetRoleScopeTagsByIdsComplete retrieves all the results into a single object
func (c RoleScopeTagClient) GetRoleScopeTagsByIdsComplete(ctx context.Context, input GetRoleScopeTagsByIdsRequest, options GetRoleScopeTagsByIdsOperationOptions) (GetRoleScopeTagsByIdsCompleteResult, error) {
	return c.GetRoleScopeTagsByIdsCompleteMatchingPredicate(ctx, input, options, RoleScopeTagOperationPredicate{})
}

// GetRoleScopeTagsByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleScopeTagClient) GetRoleScopeTagsByIdsCompleteMatchingPredicate(ctx context.Context, input GetRoleScopeTagsByIdsRequest, options GetRoleScopeTagsByIdsOperationOptions, predicate RoleScopeTagOperationPredicate) (result GetRoleScopeTagsByIdsCompleteResult, err error) {
	items := make([]beta.RoleScopeTag, 0)

	resp, err := c.GetRoleScopeTagsByIds(ctx, input, options)
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

	result = GetRoleScopeTagsByIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
