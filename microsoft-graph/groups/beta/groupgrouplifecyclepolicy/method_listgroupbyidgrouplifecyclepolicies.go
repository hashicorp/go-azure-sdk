package groupgrouplifecyclepolicy

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

type ListGroupByIdGroupLifecyclePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.GroupLifecyclePolicyCollectionResponse
}

type ListGroupByIdGroupLifecyclePoliciesCompleteResult struct {
	Items []models.GroupLifecyclePolicyCollectionResponse
}

// ListGroupByIdGroupLifecyclePolicies ...
func (c GroupGroupLifecyclePolicyClient) ListGroupByIdGroupLifecyclePolicies(ctx context.Context, id GroupId) (result ListGroupByIdGroupLifecyclePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/groupLifecyclePolicies", id.ID()),
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
		Values *[]models.GroupLifecyclePolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdGroupLifecyclePoliciesComplete retrieves all the results into a single object
func (c GroupGroupLifecyclePolicyClient) ListGroupByIdGroupLifecyclePoliciesComplete(ctx context.Context, id GroupId) (ListGroupByIdGroupLifecyclePoliciesCompleteResult, error) {
	return c.ListGroupByIdGroupLifecyclePoliciesCompleteMatchingPredicate(ctx, id, models.GroupLifecyclePolicyCollectionResponseOperationPredicate{})
}

// ListGroupByIdGroupLifecyclePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupGroupLifecyclePolicyClient) ListGroupByIdGroupLifecyclePoliciesCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.GroupLifecyclePolicyCollectionResponseOperationPredicate) (result ListGroupByIdGroupLifecyclePoliciesCompleteResult, err error) {
	items := make([]models.GroupLifecyclePolicyCollectionResponse, 0)

	resp, err := c.ListGroupByIdGroupLifecyclePolicies(ctx, id)
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

	result = ListGroupByIdGroupLifecyclePoliciesCompleteResult{
		Items: items,
	}
	return
}
