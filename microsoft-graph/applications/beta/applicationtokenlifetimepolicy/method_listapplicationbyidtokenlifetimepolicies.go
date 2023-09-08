package applicationtokenlifetimepolicy

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

type ListApplicationByIdTokenLifetimePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TokenLifetimePolicyCollectionResponse
}

type ListApplicationByIdTokenLifetimePoliciesCompleteResult struct {
	Items []models.TokenLifetimePolicyCollectionResponse
}

// ListApplicationByIdTokenLifetimePolicies ...
func (c ApplicationTokenLifetimePolicyClient) ListApplicationByIdTokenLifetimePolicies(ctx context.Context, id ApplicationId) (result ListApplicationByIdTokenLifetimePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tokenLifetimePolicies", id.ID()),
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
		Values *[]models.TokenLifetimePolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApplicationByIdTokenLifetimePoliciesComplete retrieves all the results into a single object
func (c ApplicationTokenLifetimePolicyClient) ListApplicationByIdTokenLifetimePoliciesComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdTokenLifetimePoliciesCompleteResult, error) {
	return c.ListApplicationByIdTokenLifetimePoliciesCompleteMatchingPredicate(ctx, id, models.TokenLifetimePolicyCollectionResponseOperationPredicate{})
}

// ListApplicationByIdTokenLifetimePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationTokenLifetimePolicyClient) ListApplicationByIdTokenLifetimePoliciesCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.TokenLifetimePolicyCollectionResponseOperationPredicate) (result ListApplicationByIdTokenLifetimePoliciesCompleteResult, err error) {
	items := make([]models.TokenLifetimePolicyCollectionResponse, 0)

	resp, err := c.ListApplicationByIdTokenLifetimePolicies(ctx, id)
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

	result = ListApplicationByIdTokenLifetimePoliciesCompleteResult{
		Items: items,
	}
	return
}
