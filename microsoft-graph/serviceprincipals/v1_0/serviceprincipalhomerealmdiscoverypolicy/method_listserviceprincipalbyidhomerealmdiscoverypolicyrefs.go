package serviceprincipalhomerealmdiscoverypolicy

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

type ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefs ...
func (c ServicePrincipalHomeRealmDiscoveryPolicyClient) ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefs(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/homeRealmDiscoveryPolicies/$ref", id.ID()),
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

// ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsComplete retrieves all the results into a single object
func (c ServicePrincipalHomeRealmDiscoveryPolicyClient) ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsCompleteResult, error) {
	return c.ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalHomeRealmDiscoveryPolicyClient) ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.StringCollectionResponseOperationPredicate) (result ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefs(ctx, id)
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

	result = ListServicePrincipalByIdHomeRealmDiscoveryPolicyRefsCompleteResult{
		Items: items,
	}
	return
}
