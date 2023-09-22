package policysetdefinitionversions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PolicySetDefinitionVersion
}

type ListByManagementGroupCompleteResult struct {
	Items []PolicySetDefinitionVersion
}

type ListByManagementGroupOperationOptions struct {
	Top *int64
}

func DefaultListByManagementGroupOperationOptions() ListByManagementGroupOperationOptions {
	return ListByManagementGroupOperationOptions{}
}

func (o ListByManagementGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByManagementGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByManagementGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListByManagementGroup ...
func (c PolicySetDefinitionVersionsClient) ListByManagementGroup(ctx context.Context, id Providers2PolicySetDefinitionId, options ListByManagementGroupOperationOptions) (result ListByManagementGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]PolicySetDefinitionVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByManagementGroupComplete retrieves all the results into a single object
func (c PolicySetDefinitionVersionsClient) ListByManagementGroupComplete(ctx context.Context, id Providers2PolicySetDefinitionId, options ListByManagementGroupOperationOptions) (ListByManagementGroupCompleteResult, error) {
	return c.ListByManagementGroupCompleteMatchingPredicate(ctx, id, options, PolicySetDefinitionVersionOperationPredicate{})
}

// ListByManagementGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicySetDefinitionVersionsClient) ListByManagementGroupCompleteMatchingPredicate(ctx context.Context, id Providers2PolicySetDefinitionId, options ListByManagementGroupOperationOptions, predicate PolicySetDefinitionVersionOperationPredicate) (result ListByManagementGroupCompleteResult, err error) {
	items := make([]PolicySetDefinitionVersion, 0)

	resp, err := c.ListByManagementGroup(ctx, id, options)
	if err != nil {
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

	result = ListByManagementGroupCompleteResult{
		Items: items,
	}
	return
}
