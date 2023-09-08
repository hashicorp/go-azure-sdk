package meonlinemeeting

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

type ListMeOnlineMeetingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnlineMeetingCollectionResponse
}

type ListMeOnlineMeetingsCompleteResult struct {
	Items []models.OnlineMeetingCollectionResponse
}

// ListMeOnlineMeetings ...
func (c MeOnlineMeetingClient) ListMeOnlineMeetings(ctx context.Context) (result ListMeOnlineMeetingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/onlineMeetings",
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
		Values *[]models.OnlineMeetingCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOnlineMeetingsComplete retrieves all the results into a single object
func (c MeOnlineMeetingClient) ListMeOnlineMeetingsComplete(ctx context.Context) (ListMeOnlineMeetingsCompleteResult, error) {
	return c.ListMeOnlineMeetingsCompleteMatchingPredicate(ctx, models.OnlineMeetingCollectionResponseOperationPredicate{})
}

// ListMeOnlineMeetingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnlineMeetingClient) ListMeOnlineMeetingsCompleteMatchingPredicate(ctx context.Context, predicate models.OnlineMeetingCollectionResponseOperationPredicate) (result ListMeOnlineMeetingsCompleteResult, err error) {
	items := make([]models.OnlineMeetingCollectionResponse, 0)

	resp, err := c.ListMeOnlineMeetings(ctx)
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

	result = ListMeOnlineMeetingsCompleteResult{
		Items: items,
	}
	return
}
