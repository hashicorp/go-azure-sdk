package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByApplicationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServiceResource
}

type ListByApplicationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ServiceResource
}

// ListByApplications ...
func (c ServiceClient) ListByApplications(ctx context.Context, id ApplicationId) (result ListByApplicationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/services", id.ID()),
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
		Values *[]ServiceResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByApplicationsComplete retrieves all the results into a single object
func (c ServiceClient) ListByApplicationsComplete(ctx context.Context, id ApplicationId) (ListByApplicationsCompleteResult, error) {
	return c.ListByApplicationsCompleteMatchingPredicate(ctx, id, ServiceResourceOperationPredicate{})
}

// ListByApplicationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServiceClient) ListByApplicationsCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate ServiceResourceOperationPredicate) (result ListByApplicationsCompleteResult, err error) {
	items := make([]ServiceResource, 0)

	resp, err := c.ListByApplications(ctx, id)
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

	result = ListByApplicationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
