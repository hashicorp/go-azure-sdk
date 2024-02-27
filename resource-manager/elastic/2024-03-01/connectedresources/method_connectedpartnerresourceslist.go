package connectedresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedPartnerResourcesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ConnectedPartnerResourcesListFormat
}

type ConnectedPartnerResourcesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ConnectedPartnerResourcesListFormat
}

// ConnectedPartnerResourcesList ...
func (c ConnectedResourcesClient) ConnectedPartnerResourcesList(ctx context.Context, id MonitorId) (result ConnectedPartnerResourcesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/listConnectedPartnerResources", id.ID()),
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
		Values *[]ConnectedPartnerResourcesListFormat `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConnectedPartnerResourcesListComplete retrieves all the results into a single object
func (c ConnectedResourcesClient) ConnectedPartnerResourcesListComplete(ctx context.Context, id MonitorId) (ConnectedPartnerResourcesListCompleteResult, error) {
	return c.ConnectedPartnerResourcesListCompleteMatchingPredicate(ctx, id, ConnectedPartnerResourcesListFormatOperationPredicate{})
}

// ConnectedPartnerResourcesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConnectedResourcesClient) ConnectedPartnerResourcesListCompleteMatchingPredicate(ctx context.Context, id MonitorId, predicate ConnectedPartnerResourcesListFormatOperationPredicate) (result ConnectedPartnerResourcesListCompleteResult, err error) {
	items := make([]ConnectedPartnerResourcesListFormat, 0)

	resp, err := c.ConnectedPartnerResourcesList(ctx, id)
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

	result = ConnectedPartnerResourcesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
