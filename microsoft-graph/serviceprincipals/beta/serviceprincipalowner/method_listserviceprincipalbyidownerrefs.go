package serviceprincipalowner

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

type ListServicePrincipalByIdOwnerRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListServicePrincipalByIdOwnerRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListServicePrincipalByIdOwnerRefs ...
func (c ServicePrincipalOwnerClient) ListServicePrincipalByIdOwnerRefs(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdOwnerRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/owners/$ref", id.ID()),
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

// ListServicePrincipalByIdOwnerRefsComplete retrieves all the results into a single object
func (c ServicePrincipalOwnerClient) ListServicePrincipalByIdOwnerRefsComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdOwnerRefsCompleteResult, error) {
	return c.ListServicePrincipalByIdOwnerRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdOwnerRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalOwnerClient) ListServicePrincipalByIdOwnerRefsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.StringCollectionResponseOperationPredicate) (result ListServicePrincipalByIdOwnerRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdOwnerRefs(ctx, id)
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

	result = ListServicePrincipalByIdOwnerRefsCompleteResult{
		Items: items,
	}
	return
}
