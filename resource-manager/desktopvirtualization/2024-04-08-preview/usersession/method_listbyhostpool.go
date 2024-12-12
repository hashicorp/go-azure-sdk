package usersession

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]UserSession
}

type ListByHostPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []UserSession
}

type ListByHostPoolOperationOptions struct {
	Filter       *string
	InitialSkip  *int64
	IsDescending *bool
	PageSize     *int64
}

func DefaultListByHostPoolOperationOptions() ListByHostPoolOperationOptions {
	return ListByHostPoolOperationOptions{}
}

func (o ListByHostPoolOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByHostPoolOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByHostPoolOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.InitialSkip != nil {
		out.Append("initialSkip", fmt.Sprintf("%v", *o.InitialSkip))
	}
	if o.IsDescending != nil {
		out.Append("isDescending", fmt.Sprintf("%v", *o.IsDescending))
	}
	if o.PageSize != nil {
		out.Append("pageSize", fmt.Sprintf("%v", *o.PageSize))
	}
	return &out
}

type ListByHostPoolCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByHostPoolCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByHostPool ...
func (c UserSessionClient) ListByHostPool(ctx context.Context, id HostPoolId, options ListByHostPoolOperationOptions) (result ListByHostPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByHostPoolCustomPager{},
		Path:          fmt.Sprintf("%s/userSessions", id.ID()),
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
		Values *[]UserSession `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByHostPoolComplete retrieves all the results into a single object
func (c UserSessionClient) ListByHostPoolComplete(ctx context.Context, id HostPoolId, options ListByHostPoolOperationOptions) (ListByHostPoolCompleteResult, error) {
	return c.ListByHostPoolCompleteMatchingPredicate(ctx, id, options, UserSessionOperationPredicate{})
}

// ListByHostPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserSessionClient) ListByHostPoolCompleteMatchingPredicate(ctx context.Context, id HostPoolId, options ListByHostPoolOperationOptions, predicate UserSessionOperationPredicate) (result ListByHostPoolCompleteResult, err error) {
	items := make([]UserSession, 0)

	resp, err := c.ListByHostPool(ctx, id, options)
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

	result = ListByHostPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
