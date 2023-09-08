package identityuserflowattribute

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

type ListIdentityUserFlowAttributesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.IdentityUserFlowAttributeCollectionResponse
}

type ListIdentityUserFlowAttributesCompleteResult struct {
	Items []models.IdentityUserFlowAttributeCollectionResponse
}

// ListIdentityUserFlowAttributes ...
func (c IdentityUserFlowAttributeClient) ListIdentityUserFlowAttributes(ctx context.Context) (result ListIdentityUserFlowAttributesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/identity/userFlowAttributes",
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
		Values *[]models.IdentityUserFlowAttributeCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityUserFlowAttributesComplete retrieves all the results into a single object
func (c IdentityUserFlowAttributeClient) ListIdentityUserFlowAttributesComplete(ctx context.Context) (ListIdentityUserFlowAttributesCompleteResult, error) {
	return c.ListIdentityUserFlowAttributesCompleteMatchingPredicate(ctx, models.IdentityUserFlowAttributeCollectionResponseOperationPredicate{})
}

// ListIdentityUserFlowAttributesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityUserFlowAttributeClient) ListIdentityUserFlowAttributesCompleteMatchingPredicate(ctx context.Context, predicate models.IdentityUserFlowAttributeCollectionResponseOperationPredicate) (result ListIdentityUserFlowAttributesCompleteResult, err error) {
	items := make([]models.IdentityUserFlowAttributeCollectionResponse, 0)

	resp, err := c.ListIdentityUserFlowAttributes(ctx)
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

	result = ListIdentityUserFlowAttributesCompleteResult{
		Items: items,
	}
	return
}
