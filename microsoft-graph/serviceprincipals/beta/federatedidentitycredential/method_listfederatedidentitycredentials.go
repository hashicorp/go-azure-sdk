package federatedidentitycredential

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListFederatedIdentityCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.FederatedIdentityCredential
}

type ListFederatedIdentityCredentialsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.FederatedIdentityCredential
}

type ListFederatedIdentityCredentialsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListFederatedIdentityCredentialsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListFederatedIdentityCredentials ...
func (c FederatedIdentityCredentialClient) ListFederatedIdentityCredentials(ctx context.Context, id ServicePrincipalId) (result ListFederatedIdentityCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListFederatedIdentityCredentialsCustomPager{},
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
		Values *[]beta.FederatedIdentityCredential `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListFederatedIdentityCredentialsComplete retrieves all the results into a single object
func (c FederatedIdentityCredentialClient) ListFederatedIdentityCredentialsComplete(ctx context.Context, id ServicePrincipalId) (ListFederatedIdentityCredentialsCompleteResult, error) {
	return c.ListFederatedIdentityCredentialsCompleteMatchingPredicate(ctx, id, FederatedIdentityCredentialOperationPredicate{})
}

// ListFederatedIdentityCredentialsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FederatedIdentityCredentialClient) ListFederatedIdentityCredentialsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate FederatedIdentityCredentialOperationPredicate) (result ListFederatedIdentityCredentialsCompleteResult, err error) {
	items := make([]beta.FederatedIdentityCredential, 0)

	resp, err := c.ListFederatedIdentityCredentials(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListFederatedIdentityCredentialsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
