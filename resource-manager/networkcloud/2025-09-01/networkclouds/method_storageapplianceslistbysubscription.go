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

type StorageAppliancesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StorageAppliance
}

type StorageAppliancesListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StorageAppliance
}

type StorageAppliancesListBySubscriptionOperationOptions struct {
	Top *int64
}

func DefaultStorageAppliancesListBySubscriptionOperationOptions() StorageAppliancesListBySubscriptionOperationOptions {
	return StorageAppliancesListBySubscriptionOperationOptions{}
}

func (o StorageAppliancesListBySubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StorageAppliancesListBySubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o StorageAppliancesListBySubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type StorageAppliancesListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StorageAppliancesListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StorageAppliancesListBySubscription ...
func (c NetworkcloudsClient) StorageAppliancesListBySubscription(ctx context.Context, id commonids.SubscriptionId, options StorageAppliancesListBySubscriptionOperationOptions) (result StorageAppliancesListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &StorageAppliancesListBySubscriptionCustomPager{},
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

// StorageAppliancesListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) StorageAppliancesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options StorageAppliancesListBySubscriptionOperationOptions) (StorageAppliancesListBySubscriptionCompleteResult, error) {
	return c.StorageAppliancesListBySubscriptionCompleteMatchingPredicate(ctx, id, options, StorageApplianceOperationPredicate{})
}

// StorageAppliancesListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) StorageAppliancesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options StorageAppliancesListBySubscriptionOperationOptions, predicate StorageApplianceOperationPredicate) (result StorageAppliancesListBySubscriptionCompleteResult, err error) {
	items := make([]StorageAppliance, 0)

	resp, err := c.StorageAppliancesListBySubscription(ctx, id, options)
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

	result = StorageAppliancesListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
