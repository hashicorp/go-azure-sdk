package meprofilepatent

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

type ListMeProfilePatentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemPatentCollectionResponse
}

type ListMeProfilePatentsCompleteResult struct {
	Items []models.ItemPatentCollectionResponse
}

// ListMeProfilePatents ...
func (c MeProfilePatentClient) ListMeProfilePatents(ctx context.Context) (result ListMeProfilePatentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/patents",
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
		Values *[]models.ItemPatentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeProfilePatentsComplete retrieves all the results into a single object
func (c MeProfilePatentClient) ListMeProfilePatentsComplete(ctx context.Context) (ListMeProfilePatentsCompleteResult, error) {
	return c.ListMeProfilePatentsCompleteMatchingPredicate(ctx, models.ItemPatentCollectionResponseOperationPredicate{})
}

// ListMeProfilePatentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfilePatentClient) ListMeProfilePatentsCompleteMatchingPredicate(ctx context.Context, predicate models.ItemPatentCollectionResponseOperationPredicate) (result ListMeProfilePatentsCompleteResult, err error) {
	items := make([]models.ItemPatentCollectionResponse, 0)

	resp, err := c.ListMeProfilePatents(ctx)
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

	result = ListMeProfilePatentsCompleteResult{
		Items: items,
	}
	return
}
