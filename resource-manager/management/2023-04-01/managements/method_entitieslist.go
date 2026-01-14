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

type EntitiesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EntityInfo
}

type EntitiesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EntityInfo
}

type EntitiesListOperationOptions struct {
	CacheControl *string
	Filter       *string
	GroupName    *string
	Search       *EntitySearchType
	Select       *string
	Skip         *int64
	Top          *int64
	View         *EntityViewParameterType
}

func DefaultEntitiesListOperationOptions() EntitiesListOperationOptions {
	return EntitiesListOperationOptions{}
}

func (o EntitiesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.CacheControl != nil {
		out.Append("Cache-Control", fmt.Sprintf("%v", *o.CacheControl))
	}
	return &out
}

func (o EntitiesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o EntitiesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.GroupName != nil {
		out.Append("groupName", fmt.Sprintf("%v", *o.GroupName))
	}
	if o.Search != nil {
		out.Append("$search", fmt.Sprintf("%v", *o.Search))
	}
	if o.Select != nil {
		out.Append("$select", fmt.Sprintf("%v", *o.Select))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	if o.View != nil {
		out.Append("$view", fmt.Sprintf("%v", *o.View))
	}
	return &out
}

type EntitiesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *EntitiesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// EntitiesList ...
func (c ManagementsClient) EntitiesList(ctx context.Context, options EntitiesListOperationOptions) (result EntitiesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &EntitiesListCustomPager{},
		Path:          "/providers/Microsoft.Management/getEntities",
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
		Values *[]EntityInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// EntitiesListComplete retrieves all the results into a single object
func (c ManagementsClient) EntitiesListComplete(ctx context.Context, options EntitiesListOperationOptions) (EntitiesListCompleteResult, error) {
	return c.EntitiesListCompleteMatchingPredicate(ctx, options, EntityInfoOperationPredicate{})
}

// EntitiesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagementsClient) EntitiesListCompleteMatchingPredicate(ctx context.Context, options EntitiesListOperationOptions, predicate EntityInfoOperationPredicate) (result EntitiesListCompleteResult, err error) {
	items := make([]EntityInfo, 0)

	resp, err := c.EntitiesList(ctx, options)
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

	result = EntitiesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
