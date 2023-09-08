package serviceprincipalfederatedidentitycredential

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

type ListServicePrincipalByIdFederatedIdentityCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.FederatedIdentityCredentialCollectionResponse
}

type ListServicePrincipalByIdFederatedIdentityCredentialsCompleteResult struct {
	Items []models.FederatedIdentityCredentialCollectionResponse
}

// ListServicePrincipalByIdFederatedIdentityCredentials ...
func (c ServicePrincipalFederatedIdentityCredentialClient) ListServicePrincipalByIdFederatedIdentityCredentials(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdFederatedIdentityCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/federatedIdentityCredentials", id.ID()),
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
		Values *[]models.FederatedIdentityCredentialCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalByIdFederatedIdentityCredentialsComplete retrieves all the results into a single object
func (c ServicePrincipalFederatedIdentityCredentialClient) ListServicePrincipalByIdFederatedIdentityCredentialsComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdFederatedIdentityCredentialsCompleteResult, error) {
	return c.ListServicePrincipalByIdFederatedIdentityCredentialsCompleteMatchingPredicate(ctx, id, models.FederatedIdentityCredentialCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdFederatedIdentityCredentialsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalFederatedIdentityCredentialClient) ListServicePrincipalByIdFederatedIdentityCredentialsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.FederatedIdentityCredentialCollectionResponseOperationPredicate) (result ListServicePrincipalByIdFederatedIdentityCredentialsCompleteResult, err error) {
	items := make([]models.FederatedIdentityCredentialCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdFederatedIdentityCredentials(ctx, id)
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

	result = ListServicePrincipalByIdFederatedIdentityCredentialsCompleteResult{
		Items: items,
	}
	return
}
