package networkclouds

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

type StorageAppliancesListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StorageAppliance
}

type StorageAppliancesListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StorageAppliance
}

type StorageAppliancesListByResourceGroupOperationOptions struct {
	Top *int64
}

func DefaultStorageAppliancesListByResourceGroupOperationOptions() StorageAppliancesListByResourceGroupOperationOptions {
	return StorageAppliancesListByResourceGroupOperationOptions{}
}

func (o StorageAppliancesListByResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StorageAppliancesListByResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o StorageAppliancesListByResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type StorageAppliancesListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StorageAppliancesListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StorageAppliancesListByResourceGroup ...
func (c NetworkcloudsClient) StorageAppliancesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options StorageAppliancesListByResourceGroupOperationOptions) (result StorageAppliancesListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &StorageAppliancesListByResourceGroupCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/storageAppliances", id.ID()),
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
		Values *[]StorageAppliance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StorageAppliancesListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) StorageAppliancesListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId, options StorageAppliancesListByResourceGroupOperationOptions) (StorageAppliancesListByResourceGroupCompleteResult, error) {
	return c.StorageAppliancesListByResourceGroupCompleteMatchingPredicate(ctx, id, options, StorageApplianceOperationPredicate{})
}

// StorageAppliancesListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) StorageAppliancesListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options StorageAppliancesListByResourceGroupOperationOptions, predicate StorageApplianceOperationPredicate) (result StorageAppliancesListByResourceGroupCompleteResult, err error) {
	items := make([]StorageAppliance, 0)

	resp, err := c.StorageAppliancesListByResourceGroup(ctx, id, options)
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

	result = StorageAppliancesListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
