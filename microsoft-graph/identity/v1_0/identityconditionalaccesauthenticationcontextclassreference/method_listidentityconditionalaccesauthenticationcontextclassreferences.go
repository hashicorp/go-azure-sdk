package identityconditionalaccesauthenticationcontextclassreference

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

type ListIdentityConditionalAccesAuthenticationContextClassReferencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AuthenticationContextClassReferenceCollectionResponse
}

type ListIdentityConditionalAccesAuthenticationContextClassReferencesCompleteResult struct {
	Items []models.AuthenticationContextClassReferenceCollectionResponse
}

// ListIdentityConditionalAccesAuthenticationContextClassReferences ...
func (c IdentityConditionalAccesAuthenticationContextClassReferenceClient) ListIdentityConditionalAccesAuthenticationContextClassReferences(ctx context.Context) (result ListIdentityConditionalAccesAuthenticationContextClassReferencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/identity/conditionalAccess/authenticationContextClassReferences",
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
		Values *[]models.AuthenticationContextClassReferenceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityConditionalAccesAuthenticationContextClassReferencesComplete retrieves all the results into a single object
func (c IdentityConditionalAccesAuthenticationContextClassReferenceClient) ListIdentityConditionalAccesAuthenticationContextClassReferencesComplete(ctx context.Context) (ListIdentityConditionalAccesAuthenticationContextClassReferencesCompleteResult, error) {
	return c.ListIdentityConditionalAccesAuthenticationContextClassReferencesCompleteMatchingPredicate(ctx, models.AuthenticationContextClassReferenceCollectionResponseOperationPredicate{})
}

// ListIdentityConditionalAccesAuthenticationContextClassReferencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityConditionalAccesAuthenticationContextClassReferenceClient) ListIdentityConditionalAccesAuthenticationContextClassReferencesCompleteMatchingPredicate(ctx context.Context, predicate models.AuthenticationContextClassReferenceCollectionResponseOperationPredicate) (result ListIdentityConditionalAccesAuthenticationContextClassReferencesCompleteResult, err error) {
	items := make([]models.AuthenticationContextClassReferenceCollectionResponse, 0)

	resp, err := c.ListIdentityConditionalAccesAuthenticationContextClassReferences(ctx)
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

	result = ListIdentityConditionalAccesAuthenticationContextClassReferencesCompleteResult{
		Items: items,
	}
	return
}
