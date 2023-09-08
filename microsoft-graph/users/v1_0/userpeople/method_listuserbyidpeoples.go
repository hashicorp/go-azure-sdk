package userpeople

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

type ListUserByIdPeoplesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonCollectionResponse
}

type ListUserByIdPeoplesCompleteResult struct {
	Items []models.PersonCollectionResponse
}

// ListUserByIdPeoples ...
func (c UserPeopleClient) ListUserByIdPeoples(ctx context.Context, id UserId) (result ListUserByIdPeoplesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/people", id.ID()),
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
		Values *[]models.PersonCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdPeoplesComplete retrieves all the results into a single object
func (c UserPeopleClient) ListUserByIdPeoplesComplete(ctx context.Context, id UserId) (ListUserByIdPeoplesCompleteResult, error) {
	return c.ListUserByIdPeoplesCompleteMatchingPredicate(ctx, id, models.PersonCollectionResponseOperationPredicate{})
}

// ListUserByIdPeoplesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPeopleClient) ListUserByIdPeoplesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PersonCollectionResponseOperationPredicate) (result ListUserByIdPeoplesCompleteResult, err error) {
	items := make([]models.PersonCollectionResponse, 0)

	resp, err := c.ListUserByIdPeoples(ctx, id)
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

	result = ListUserByIdPeoplesCompleteResult{
		Items: items,
	}
	return
}
