package managements

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroupsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagementGroupInfo
}

type ManagementGroupsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagementGroupInfo
}

type ManagementGroupsListOperationOptions struct {
	CacheControl *string
}

func DefaultManagementGroupsListOperationOptions() ManagementGroupsListOperationOptions {
	return ManagementGroupsListOperationOptions{}
}

func (o ManagementGroupsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.CacheControl != nil {
		out.Append("Cache-Control", fmt.Sprintf("%v", *o.CacheControl))
	}
	return &out
}

func (o ManagementGroupsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ManagementGroupsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ManagementGroupsListCustomPager struct {
	NextLink *odata.Link `json:"@nextLink"`
}

func (p *ManagementGroupsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ManagementGroupsList ...
func (c ManagementsClient) ManagementGroupsList(ctx context.Context, options ManagementGroupsListOperationOptions) (result ManagementGroupsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ManagementGroupsListCustomPager{},
		Path:          "/providers/Microsoft.Management/managementGroups",
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
		Values *[]ManagementGroupInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ManagementGroupsListComplete retrieves all the results into a single object
func (c ManagementsClient) ManagementGroupsListComplete(ctx context.Context, options ManagementGroupsListOperationOptions) (ManagementGroupsListCompleteResult, error) {
	return c.ManagementGroupsListCompleteMatchingPredicate(ctx, options, ManagementGroupInfoOperationPredicate{})
}

// ManagementGroupsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagementsClient) ManagementGroupsListCompleteMatchingPredicate(ctx context.Context, options ManagementGroupsListOperationOptions, predicate ManagementGroupInfoOperationPredicate) (result ManagementGroupsListCompleteResult, err error) {
	items := make([]ManagementGroupInfo, 0)

	resp, err := c.ManagementGroupsList(ctx, options)
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

	result = ManagementGroupsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
