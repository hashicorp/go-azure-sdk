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

type VolumesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Volume
}

type VolumesListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Volume
}

type VolumesListBySubscriptionOperationOptions struct {
	Top *int64
}

func DefaultVolumesListBySubscriptionOperationOptions() VolumesListBySubscriptionOperationOptions {
	return VolumesListBySubscriptionOperationOptions{}
}

func (o VolumesListBySubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o VolumesListBySubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o VolumesListBySubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type VolumesListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VolumesListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VolumesListBySubscription ...
func (c NetworkcloudsClient) VolumesListBySubscription(ctx context.Context, id commonids.SubscriptionId, options VolumesListBySubscriptionOperationOptions) (result VolumesListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &VolumesListBySubscriptionCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/volumes", id.ID()),
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
		Values *[]Volume `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VolumesListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) VolumesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options VolumesListBySubscriptionOperationOptions) (VolumesListBySubscriptionCompleteResult, error) {
	return c.VolumesListBySubscriptionCompleteMatchingPredicate(ctx, id, options, VolumeOperationPredicate{})
}

// VolumesListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) VolumesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options VolumesListBySubscriptionOperationOptions, predicate VolumeOperationPredicate) (result VolumesListBySubscriptionCompleteResult, err error) {
	items := make([]Volume, 0)

	resp, err := c.VolumesListBySubscription(ctx, id, options)
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

	result = VolumesListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
