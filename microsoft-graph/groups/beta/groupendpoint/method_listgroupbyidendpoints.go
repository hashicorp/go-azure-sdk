package groupendpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdEndpointsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EndpointCollectionResponse
}

type ListGroupByIdEndpointsCompleteResult struct {
	Items []models.EndpointCollectionResponse
}

// ListGroupByIdEndpoints ...
func (c GroupEndpointClient) ListGroupByIdEndpoints(ctx context.Context, id GroupId) (result ListGroupByIdEndpointsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/endpoints", id.ID()),
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
		Values *[]models.EndpointCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdEndpointsComplete retrieves all the results into a single object
func (c GroupEndpointClient) ListGroupByIdEndpointsComplete(ctx context.Context, id GroupId) (ListGroupByIdEndpointsCompleteResult, error) {
	return c.ListGroupByIdEndpointsCompleteMatchingPredicate(ctx, id, models.EndpointCollectionResponseOperationPredicate{})
}

// ListGroupByIdEndpointsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupEndpointClient) ListGroupByIdEndpointsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.EndpointCollectionResponseOperationPredicate) (result ListGroupByIdEndpointsCompleteResult, err error) {
	items := make([]models.EndpointCollectionResponse, 0)

	resp, err := c.ListGroupByIdEndpoints(ctx, id)
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

	result = ListGroupByIdEndpointsCompleteResult{
		Items: items,
	}
	return
}
