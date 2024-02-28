package iotsecuritysolutions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecuritySolutionsResourceGroupListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]IoTSecuritySolutionModel
}

type IoTSecuritySolutionsResourceGroupListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []IoTSecuritySolutionModel
}

type IoTSecuritySolutionsResourceGroupListOperationOptions struct {
	Filter *string
}

func DefaultIoTSecuritySolutionsResourceGroupListOperationOptions() IoTSecuritySolutionsResourceGroupListOperationOptions {
	return IoTSecuritySolutionsResourceGroupListOperationOptions{}
}

func (o IoTSecuritySolutionsResourceGroupListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o IoTSecuritySolutionsResourceGroupListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o IoTSecuritySolutionsResourceGroupListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// IoTSecuritySolutionsResourceGroupList ...
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsResourceGroupList(ctx context.Context, id commonids.ResourceGroupId, options IoTSecuritySolutionsResourceGroupListOperationOptions) (result IoTSecuritySolutionsResourceGroupListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Security/iotSecuritySolutions", id.ID()),
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
		Values *[]IoTSecuritySolutionModel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// IoTSecuritySolutionsResourceGroupListComplete retrieves all the results into a single object
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsResourceGroupListComplete(ctx context.Context, id commonids.ResourceGroupId, options IoTSecuritySolutionsResourceGroupListOperationOptions) (IoTSecuritySolutionsResourceGroupListCompleteResult, error) {
	return c.IoTSecuritySolutionsResourceGroupListCompleteMatchingPredicate(ctx, id, options, IoTSecuritySolutionModelOperationPredicate{})
}

// IoTSecuritySolutionsResourceGroupListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsResourceGroupListCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options IoTSecuritySolutionsResourceGroupListOperationOptions, predicate IoTSecuritySolutionModelOperationPredicate) (result IoTSecuritySolutionsResourceGroupListCompleteResult, err error) {
	items := make([]IoTSecuritySolutionModel, 0)

	resp, err := c.IoTSecuritySolutionsResourceGroupList(ctx, id, options)
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

	result = IoTSecuritySolutionsResourceGroupListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
