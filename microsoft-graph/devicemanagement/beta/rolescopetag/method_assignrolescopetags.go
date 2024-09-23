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

type AssignRoleScopeTagsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RoleScopeTagAutoAssignment
}

type AssignRoleScopeTagsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RoleScopeTagAutoAssignment
}

type AssignRoleScopeTagsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAssignRoleScopeTagsOperationOptions() AssignRoleScopeTagsOperationOptions {
	return AssignRoleScopeTagsOperationOptions{}
}

func (o AssignRoleScopeTagsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignRoleScopeTagsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o AssignRoleScopeTagsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignRoleScopeTagsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignRoleScopeTagsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignRoleScopeTags - Invoke action assign
func (c RoleScopeTagClient) AssignRoleScopeTags(ctx context.Context, id beta.DeviceManagementRoleScopeTagId, input AssignRoleScopeTagsRequest, options AssignRoleScopeTagsOperationOptions) (result AssignRoleScopeTagsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignRoleScopeTagsCustomPager{},
		Path:          fmt.Sprintf("%s/assign", id.ID()),
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
		Values *[]beta.RoleScopeTagAutoAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignRoleScopeTagsComplete retrieves all the results into a single object
func (c RoleScopeTagClient) AssignRoleScopeTagsComplete(ctx context.Context, id beta.DeviceManagementRoleScopeTagId, input AssignRoleScopeTagsRequest, options AssignRoleScopeTagsOperationOptions) (AssignRoleScopeTagsCompleteResult, error) {
	return c.AssignRoleScopeTagsCompleteMatchingPredicate(ctx, id, input, options, RoleScopeTagAutoAssignmentOperationPredicate{})
}

// AssignRoleScopeTagsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleScopeTagClient) AssignRoleScopeTagsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementRoleScopeTagId, input AssignRoleScopeTagsRequest, options AssignRoleScopeTagsOperationOptions, predicate RoleScopeTagAutoAssignmentOperationPredicate) (result AssignRoleScopeTagsCompleteResult, err error) {
	items := make([]beta.RoleScopeTagAutoAssignment, 0)

	resp, err := c.AssignRoleScopeTags(ctx, id, input, options)
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

	result = AssignRoleScopeTagsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
