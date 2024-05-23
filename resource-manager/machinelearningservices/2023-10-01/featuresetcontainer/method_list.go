package featuresetcontainer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FeaturesetContainerResource
}

type ListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []FeaturesetContainerResource
}

type ListOperationOptions struct {
	CreatedBy    *string
	Description  *string
	ListViewType *ListViewType
	Name         *string
	PageSize     *int64
	Skip         *string
	Tags         *string
}

func DefaultListOperationOptions() ListOperationOptions {
	return ListOperationOptions{}
}

func (o ListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.CreatedBy != nil {
		out.Append("createdBy", fmt.Sprintf("%v", *o.CreatedBy))
	}
	if o.Description != nil {
		out.Append("description", fmt.Sprintf("%v", *o.Description))
	}
	if o.ListViewType != nil {
		out.Append("listViewType", fmt.Sprintf("%v", *o.ListViewType))
	}
	if o.Name != nil {
		out.Append("name", fmt.Sprintf("%v", *o.Name))
	}
	if o.PageSize != nil {
		out.Append("pageSize", fmt.Sprintf("%v", *o.PageSize))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Tags != nil {
		out.Append("tags", fmt.Sprintf("%v", *o.Tags))
	}
	return &out
}

// List ...
func (c FeaturesetContainerClient) List(ctx context.Context, id WorkspaceId, options ListOperationOptions) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/featureSets", id.ID()),
		OptionsObject: options,
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
		Values *[]FeaturesetContainerResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplete retrieves all the results into a single object
func (c FeaturesetContainerClient) ListComplete(ctx context.Context, id WorkspaceId, options ListOperationOptions) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, options, FeaturesetContainerResourceOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FeaturesetContainerClient) ListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options ListOperationOptions, predicate FeaturesetContainerResourceOperationPredicate) (result ListCompleteResult, err error) {
	items := make([]FeaturesetContainerResource, 0)

	resp, err := c.List(ctx, id, options)
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

	result = ListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
