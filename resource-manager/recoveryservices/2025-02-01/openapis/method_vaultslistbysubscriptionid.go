package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultsListBySubscriptionIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Vault
}

type VaultsListBySubscriptionIdCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Vault
}

type VaultsListBySubscriptionIdCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VaultsListBySubscriptionIdCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VaultsListBySubscriptionId ...
func (c OpenapisClient) VaultsListBySubscriptionId(ctx context.Context, id commonids.SubscriptionId) (result VaultsListBySubscriptionIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &VaultsListBySubscriptionIdCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.RecoveryServices/vaults", id.ID()),
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
		Values *[]Vault `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VaultsListBySubscriptionIdComplete retrieves all the results into a single object
func (c OpenapisClient) VaultsListBySubscriptionIdComplete(ctx context.Context, id commonids.SubscriptionId) (VaultsListBySubscriptionIdCompleteResult, error) {
	return c.VaultsListBySubscriptionIdCompleteMatchingPredicate(ctx, id, VaultOperationPredicate{})
}

// VaultsListBySubscriptionIdCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) VaultsListBySubscriptionIdCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate VaultOperationPredicate) (result VaultsListBySubscriptionIdCompleteResult, err error) {
	items := make([]Vault, 0)

	resp, err := c.VaultsListBySubscriptionId(ctx, id)
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

	result = VaultsListBySubscriptionIdCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
