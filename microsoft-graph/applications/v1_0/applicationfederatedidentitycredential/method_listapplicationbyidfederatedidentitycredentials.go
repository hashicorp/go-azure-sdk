package applicationfederatedidentitycredential

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

type ListApplicationByIdFederatedIdentityCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.FederatedIdentityCredentialCollectionResponse
}

type ListApplicationByIdFederatedIdentityCredentialsCompleteResult struct {
	Items []models.FederatedIdentityCredentialCollectionResponse
}

// ListApplicationByIdFederatedIdentityCredentials ...
func (c ApplicationFederatedIdentityCredentialClient) ListApplicationByIdFederatedIdentityCredentials(ctx context.Context, id ApplicationId) (result ListApplicationByIdFederatedIdentityCredentialsOperationResponse, err error) {
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

// ListApplicationByIdFederatedIdentityCredentialsComplete retrieves all the results into a single object
func (c ApplicationFederatedIdentityCredentialClient) ListApplicationByIdFederatedIdentityCredentialsComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdFederatedIdentityCredentialsCompleteResult, error) {
	return c.ListApplicationByIdFederatedIdentityCredentialsCompleteMatchingPredicate(ctx, id, models.FederatedIdentityCredentialCollectionResponseOperationPredicate{})
}

// ListApplicationByIdFederatedIdentityCredentialsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationFederatedIdentityCredentialClient) ListApplicationByIdFederatedIdentityCredentialsCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.FederatedIdentityCredentialCollectionResponseOperationPredicate) (result ListApplicationByIdFederatedIdentityCredentialsCompleteResult, err error) {
	items := make([]models.FederatedIdentityCredentialCollectionResponse, 0)

	resp, err := c.ListApplicationByIdFederatedIdentityCredentials(ctx, id)
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

	result = ListApplicationByIdFederatedIdentityCredentialsCompleteResult{
		Items: items,
	}
	return
}
