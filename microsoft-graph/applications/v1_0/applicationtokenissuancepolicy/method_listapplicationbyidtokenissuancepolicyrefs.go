package applicationtokenissuancepolicy

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

type ListApplicationByIdTokenIssuancePolicyRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListApplicationByIdTokenIssuancePolicyRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListApplicationByIdTokenIssuancePolicyRefs ...
func (c ApplicationTokenIssuancePolicyClient) ListApplicationByIdTokenIssuancePolicyRefs(ctx context.Context, id ApplicationId) (result ListApplicationByIdTokenIssuancePolicyRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tokenIssuancePolicies/$ref", id.ID()),
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

// ListApplicationByIdTokenIssuancePolicyRefsComplete retrieves all the results into a single object
func (c ApplicationTokenIssuancePolicyClient) ListApplicationByIdTokenIssuancePolicyRefsComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdTokenIssuancePolicyRefsCompleteResult, error) {
	return c.ListApplicationByIdTokenIssuancePolicyRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListApplicationByIdTokenIssuancePolicyRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationTokenIssuancePolicyClient) ListApplicationByIdTokenIssuancePolicyRefsCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.StringCollectionResponseOperationPredicate) (result ListApplicationByIdTokenIssuancePolicyRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListApplicationByIdTokenIssuancePolicyRefs(ctx, id)
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

	result = ListApplicationByIdTokenIssuancePolicyRefsCompleteResult{
		Items: items,
	}
	return
}
