package groupowner

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdOwnerRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListGroupByIdOwnerRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListGroupByIdOwnerRefs ...
func (c GroupOwnerClient) ListGroupByIdOwnerRefs(ctx context.Context, id GroupId) (result ListGroupByIdOwnerRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/owners/$ref", id.ID()),
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
		Values *[]models.StringCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdOwnerRefsComplete retrieves all the results into a single object
func (c GroupOwnerClient) ListGroupByIdOwnerRefsComplete(ctx context.Context, id GroupId) (ListGroupByIdOwnerRefsCompleteResult, error) {
	return c.ListGroupByIdOwnerRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListGroupByIdOwnerRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupOwnerClient) ListGroupByIdOwnerRefsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.StringCollectionResponseOperationPredicate) (result ListGroupByIdOwnerRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListGroupByIdOwnerRefs(ctx, id)
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

	result = ListGroupByIdOwnerRefsCompleteResult{
		Items: items,
	}
	return
}
