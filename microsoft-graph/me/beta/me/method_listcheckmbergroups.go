package me

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCheckmberGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type ListCheckmberGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type ListCheckmberGroupsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListCheckmberGroupsOperationOptions() ListCheckmberGroupsOperationOptions {
	return ListCheckmberGroupsOperationOptions{}
}

func (o ListCheckmberGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCheckmberGroupsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListCheckmberGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCheckmberGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCheckmberGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCheckmberGroups - Invoke action checkMemberGroups. Check for membership in a specified list of group IDs, and
// return from that list those groups (identified by IDs) of which the specified user, group, service principal,
// organizational contact, device, or directory object is a member. This function is transitive. You can check up to a
// maximum of 20 groups per request. This function supports all groups provisioned in Microsoft Entra ID. Because
// Microsoft 365 groups cannot contain other groups, membership in a Microsoft 365 group is always direct.
func (c MeClient) ListCheckmberGroups(ctx context.Context, input ListCheckmberGroupsRequest, options ListCheckmberGroupsOperationOptions) (result ListCheckmberGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListCheckmberGroupsCustomPager{},
		Path:          "/me/checkMemberGroups",
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
		Values *[]string `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCheckmberGroupsComplete retrieves all the results into a single object
func (c MeClient) ListCheckmberGroupsComplete(ctx context.Context, input ListCheckmberGroupsRequest, options ListCheckmberGroupsOperationOptions) (result ListCheckmberGroupsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.ListCheckmberGroups(ctx, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = ListCheckmberGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
