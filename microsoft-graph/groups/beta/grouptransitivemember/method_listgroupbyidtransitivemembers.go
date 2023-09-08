package grouptransitivemember

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

type ListGroupByIdTransitiveMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListGroupByIdTransitiveMembersCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListGroupByIdTransitiveMembers ...
func (c GroupTransitiveMemberClient) ListGroupByIdTransitiveMembers(ctx context.Context, id GroupId) (result ListGroupByIdTransitiveMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/transitiveMembers", id.ID()),
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

// ListGroupByIdTransitiveMembersComplete retrieves all the results into a single object
func (c GroupTransitiveMemberClient) ListGroupByIdTransitiveMembersComplete(ctx context.Context, id GroupId) (ListGroupByIdTransitiveMembersCompleteResult, error) {
	return c.ListGroupByIdTransitiveMembersCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListGroupByIdTransitiveMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTransitiveMemberClient) ListGroupByIdTransitiveMembersCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListGroupByIdTransitiveMembersCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListGroupByIdTransitiveMembers(ctx, id)
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

	result = ListGroupByIdTransitiveMembersCompleteResult{
		Items: items,
	}
	return
}
