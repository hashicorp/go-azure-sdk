package policyhomerealmdiscoverypolicyappliesto

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

type ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTos ...
func (c PolicyHomeRealmDiscoveryPolicyAppliesToClient) ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTos(ctx context.Context, id PolicyHomeRealmDiscoveryPolicyId) (result ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosOperationResponse, err error) {
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

// ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosComplete retrieves all the results into a single object
func (c PolicyHomeRealmDiscoveryPolicyAppliesToClient) ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosComplete(ctx context.Context, id PolicyHomeRealmDiscoveryPolicyId) (ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosCompleteResult, error) {
	return c.ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyHomeRealmDiscoveryPolicyAppliesToClient) ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosCompleteMatchingPredicate(ctx context.Context, id PolicyHomeRealmDiscoveryPolicyId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTos(ctx, id)
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

	result = ListPolicyHomeRealmDiscoveryPolicyByIdAppliesTosCompleteResult{
		Items: items,
	}
	return
}
