package groupmemberof

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

type ListGroupByIdMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListGroupByIdMemberOfCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListGroupByIdMemberOf ...
func (c GroupMemberOfClient) ListGroupByIdMemberOf(ctx context.Context, id GroupId) (result ListGroupByIdMemberOfOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/memberOf", id.ID()),
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdMemberOfComplete retrieves all the results into a single object
func (c GroupMemberOfClient) ListGroupByIdMemberOfComplete(ctx context.Context, id GroupId) (ListGroupByIdMemberOfCompleteResult, error) {
	return c.ListGroupByIdMemberOfCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListGroupByIdMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupMemberOfClient) ListGroupByIdMemberOfCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListGroupByIdMemberOfCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListGroupByIdMemberOf(ctx, id)
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

	result = ListGroupByIdMemberOfCompleteResult{
		Items: items,
	}
	return
}
