package hypervmachines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllMachinesInSiteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HyperVMachine
}

type GetAllMachinesInSiteCompleteResult struct {
	Items []HyperVMachine
}

type GetAllMachinesInSiteOperationOptions struct {
	ContinuationToken *string
	Filter            *string
	Top               *int64
	TotalRecordCount  *int64
}

func DefaultGetAllMachinesInSiteOperationOptions() GetAllMachinesInSiteOperationOptions {
	return GetAllMachinesInSiteOperationOptions{}
}

func (o GetAllMachinesInSiteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAllMachinesInSiteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetAllMachinesInSiteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ContinuationToken != nil {
		out.Append("continuationToken", fmt.Sprintf("%v", *o.ContinuationToken))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	if o.TotalRecordCount != nil {
		out.Append("totalRecordCount", fmt.Sprintf("%v", *o.TotalRecordCount))
	}
	return &out
}

// GetAllMachinesInSite ...
func (c HyperVMachinesClient) GetAllMachinesInSite(ctx context.Context, id HyperVSiteId, options GetAllMachinesInSiteOperationOptions) (result GetAllMachinesInSiteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/machines", id.ID()),
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
		Values *[]HyperVMachine `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAllMachinesInSiteComplete retrieves all the results into a single object
func (c HyperVMachinesClient) GetAllMachinesInSiteComplete(ctx context.Context, id HyperVSiteId, options GetAllMachinesInSiteOperationOptions) (GetAllMachinesInSiteCompleteResult, error) {
	return c.GetAllMachinesInSiteCompleteMatchingPredicate(ctx, id, options, HyperVMachineOperationPredicate{})
}

// GetAllMachinesInSiteCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HyperVMachinesClient) GetAllMachinesInSiteCompleteMatchingPredicate(ctx context.Context, id HyperVSiteId, options GetAllMachinesInSiteOperationOptions, predicate HyperVMachineOperationPredicate) (result GetAllMachinesInSiteCompleteResult, err error) {
	items := make([]HyperVMachine, 0)

	resp, err := c.GetAllMachinesInSite(ctx, id, options)
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

	result = GetAllMachinesInSiteCompleteResult{
		Items: items,
	}
	return
}
