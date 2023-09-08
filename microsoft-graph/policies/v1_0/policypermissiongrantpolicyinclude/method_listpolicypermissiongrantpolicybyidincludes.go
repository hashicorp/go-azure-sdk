package policypermissiongrantpolicyinclude

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

type ListPolicyPermissionGrantPolicyByIdIncludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PermissionGrantConditionSetCollectionResponse
}

type ListPolicyPermissionGrantPolicyByIdIncludesCompleteResult struct {
	Items []models.PermissionGrantConditionSetCollectionResponse
}

// ListPolicyPermissionGrantPolicyByIdIncludes ...
func (c PolicyPermissionGrantPolicyIncludeClient) ListPolicyPermissionGrantPolicyByIdIncludes(ctx context.Context, id PolicyPermissionGrantPolicyId) (result ListPolicyPermissionGrantPolicyByIdIncludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/includes", id.ID()),
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
		Values *[]models.PermissionGrantConditionSetCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyPermissionGrantPolicyByIdIncludesComplete retrieves all the results into a single object
func (c PolicyPermissionGrantPolicyIncludeClient) ListPolicyPermissionGrantPolicyByIdIncludesComplete(ctx context.Context, id PolicyPermissionGrantPolicyId) (ListPolicyPermissionGrantPolicyByIdIncludesCompleteResult, error) {
	return c.ListPolicyPermissionGrantPolicyByIdIncludesCompleteMatchingPredicate(ctx, id, models.PermissionGrantConditionSetCollectionResponseOperationPredicate{})
}

// ListPolicyPermissionGrantPolicyByIdIncludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyPermissionGrantPolicyIncludeClient) ListPolicyPermissionGrantPolicyByIdIncludesCompleteMatchingPredicate(ctx context.Context, id PolicyPermissionGrantPolicyId, predicate models.PermissionGrantConditionSetCollectionResponseOperationPredicate) (result ListPolicyPermissionGrantPolicyByIdIncludesCompleteResult, err error) {
	items := make([]models.PermissionGrantConditionSetCollectionResponse, 0)

	resp, err := c.ListPolicyPermissionGrantPolicyByIdIncludes(ctx, id)
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

	result = ListPolicyPermissionGrantPolicyByIdIncludesCompleteResult{
		Items: items,
	}
	return
}
