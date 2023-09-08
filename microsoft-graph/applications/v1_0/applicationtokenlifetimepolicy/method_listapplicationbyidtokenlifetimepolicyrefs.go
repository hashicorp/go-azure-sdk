package applicationtokenlifetimepolicy

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

type ListApplicationByIdTokenLifetimePolicyRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListApplicationByIdTokenLifetimePolicyRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListApplicationByIdTokenLifetimePolicyRefs ...
func (c ApplicationTokenLifetimePolicyClient) ListApplicationByIdTokenLifetimePolicyRefs(ctx context.Context, id ApplicationId) (result ListApplicationByIdTokenLifetimePolicyRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tokenLifetimePolicies/$ref", id.ID()),
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

// ListApplicationByIdTokenLifetimePolicyRefsComplete retrieves all the results into a single object
func (c ApplicationTokenLifetimePolicyClient) ListApplicationByIdTokenLifetimePolicyRefsComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdTokenLifetimePolicyRefsCompleteResult, error) {
	return c.ListApplicationByIdTokenLifetimePolicyRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListApplicationByIdTokenLifetimePolicyRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationTokenLifetimePolicyClient) ListApplicationByIdTokenLifetimePolicyRefsCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.StringCollectionResponseOperationPredicate) (result ListApplicationByIdTokenLifetimePolicyRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListApplicationByIdTokenLifetimePolicyRefs(ctx, id)
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

	result = ListApplicationByIdTokenLifetimePolicyRefsCompleteResult{
		Items: items,
	}
	return
}
