package policyactivitybasedtimeoutpolicyappliesto

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

type ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListPolicyActivityBasedTimeoutPolicyByIdAppliesTos ...
func (c PolicyActivityBasedTimeoutPolicyAppliesToClient) ListPolicyActivityBasedTimeoutPolicyByIdAppliesTos(ctx context.Context, id PolicyActivityBasedTimeoutPolicyId) (result ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appliesTo", id.ID()),
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

// ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosComplete retrieves all the results into a single object
func (c PolicyActivityBasedTimeoutPolicyAppliesToClient) ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosComplete(ctx context.Context, id PolicyActivityBasedTimeoutPolicyId) (ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosCompleteResult, error) {
	return c.ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyActivityBasedTimeoutPolicyAppliesToClient) ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosCompleteMatchingPredicate(ctx context.Context, id PolicyActivityBasedTimeoutPolicyId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListPolicyActivityBasedTimeoutPolicyByIdAppliesTos(ctx, id)
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

	result = ListPolicyActivityBasedTimeoutPolicyByIdAppliesTosCompleteResult{
		Items: items,
	}
	return
}
