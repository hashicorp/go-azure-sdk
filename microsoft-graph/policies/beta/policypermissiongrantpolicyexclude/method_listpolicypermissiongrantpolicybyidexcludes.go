package policypermissiongrantpolicyexclude

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

type ListPolicyPermissionGrantPolicyByIdExcludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PermissionGrantConditionSetCollectionResponse
}

type ListPolicyPermissionGrantPolicyByIdExcludesCompleteResult struct {
	Items []models.PermissionGrantConditionSetCollectionResponse
}

// ListPolicyPermissionGrantPolicyByIdExcludes ...
func (c PolicyPermissionGrantPolicyExcludeClient) ListPolicyPermissionGrantPolicyByIdExcludes(ctx context.Context, id PolicyPermissionGrantPolicyId) (result ListPolicyPermissionGrantPolicyByIdExcludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/excludes", id.ID()),
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

// ListPolicyPermissionGrantPolicyByIdExcludesComplete retrieves all the results into a single object
func (c PolicyPermissionGrantPolicyExcludeClient) ListPolicyPermissionGrantPolicyByIdExcludesComplete(ctx context.Context, id PolicyPermissionGrantPolicyId) (ListPolicyPermissionGrantPolicyByIdExcludesCompleteResult, error) {
	return c.ListPolicyPermissionGrantPolicyByIdExcludesCompleteMatchingPredicate(ctx, id, models.PermissionGrantConditionSetCollectionResponseOperationPredicate{})
}

// ListPolicyPermissionGrantPolicyByIdExcludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyPermissionGrantPolicyExcludeClient) ListPolicyPermissionGrantPolicyByIdExcludesCompleteMatchingPredicate(ctx context.Context, id PolicyPermissionGrantPolicyId, predicate models.PermissionGrantConditionSetCollectionResponseOperationPredicate) (result ListPolicyPermissionGrantPolicyByIdExcludesCompleteResult, err error) {
	items := make([]models.PermissionGrantConditionSetCollectionResponse, 0)

	resp, err := c.ListPolicyPermissionGrantPolicyByIdExcludes(ctx, id)
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

	result = ListPolicyPermissionGrantPolicyByIdExcludesCompleteResult{
		Items: items,
	}
	return
}
