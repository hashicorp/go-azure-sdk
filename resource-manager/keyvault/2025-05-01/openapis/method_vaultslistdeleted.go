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

type VaultsListDeletedOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeletedVault
}

type VaultsListDeletedCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeletedVault
}

type VaultsListDeletedCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VaultsListDeletedCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VaultsListDeleted ...
func (c OpenapisClient) VaultsListDeleted(ctx context.Context, id commonids.SubscriptionId) (result VaultsListDeletedOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &VaultsListDeletedCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.KeyVault/deletedVaults", id.ID()),
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
		Values *[]DeletedVault `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VaultsListDeletedComplete retrieves all the results into a single object
func (c OpenapisClient) VaultsListDeletedComplete(ctx context.Context, id commonids.SubscriptionId) (VaultsListDeletedCompleteResult, error) {
	return c.VaultsListDeletedCompleteMatchingPredicate(ctx, id, DeletedVaultOperationPredicate{})
}

// VaultsListDeletedCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) VaultsListDeletedCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate DeletedVaultOperationPredicate) (result VaultsListDeletedCompleteResult, err error) {
	items := make([]DeletedVault, 0)

	resp, err := c.VaultsListDeleted(ctx, id)
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

	result = VaultsListDeletedCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
