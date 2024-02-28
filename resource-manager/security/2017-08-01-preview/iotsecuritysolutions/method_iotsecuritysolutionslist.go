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

type IoTSecuritySolutionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]IoTSecuritySolutionModel
}

type IoTSecuritySolutionsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []IoTSecuritySolutionModel
}

type IoTSecuritySolutionsListOperationOptions struct {
	Filter *string
}

func DefaultIoTSecuritySolutionsListOperationOptions() IoTSecuritySolutionsListOperationOptions {
	return IoTSecuritySolutionsListOperationOptions{}
}

func (o IoTSecuritySolutionsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o IoTSecuritySolutionsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o IoTSecuritySolutionsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// IoTSecuritySolutionsList ...
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsList(ctx context.Context, id commonids.SubscriptionId, options IoTSecuritySolutionsListOperationOptions) (result IoTSecuritySolutionsListOperationResponse, err error) {
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

// IoTSecuritySolutionsListComplete retrieves all the results into a single object
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsListComplete(ctx context.Context, id commonids.SubscriptionId, options IoTSecuritySolutionsListOperationOptions) (IoTSecuritySolutionsListCompleteResult, error) {
	return c.IoTSecuritySolutionsListCompleteMatchingPredicate(ctx, id, options, IoTSecuritySolutionModelOperationPredicate{})
}

// IoTSecuritySolutionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options IoTSecuritySolutionsListOperationOptions, predicate IoTSecuritySolutionModelOperationPredicate) (result IoTSecuritySolutionsListCompleteResult, err error) {
	items := make([]IoTSecuritySolutionModel, 0)

	resp, err := c.IoTSecuritySolutionsList(ctx, id, options)
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

	result = IoTSecuritySolutionsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
