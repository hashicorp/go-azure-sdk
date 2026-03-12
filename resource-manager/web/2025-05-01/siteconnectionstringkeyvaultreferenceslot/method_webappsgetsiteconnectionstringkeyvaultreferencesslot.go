package siteconnectionstringkeyvaultreferenceslot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsGetSiteConnectionStringKeyVaultReferencesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApiKVReference
}

type WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApiKVReference
}

type WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsGetSiteConnectionStringKeyVaultReferencesSlot ...
func (c SiteConnectionStringKeyVaultReferenceSlotClient) WebAppsGetSiteConnectionStringKeyVaultReferencesSlot(ctx context.Context, id SlotId) (result WebAppsGetSiteConnectionStringKeyVaultReferencesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCustomPager{},
		Path:       fmt.Sprintf("%s/config/configReferences/connectionStrings", id.ID()),
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
		Values *[]ApiKVReference `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsGetSiteConnectionStringKeyVaultReferencesSlotComplete retrieves all the results into a single object
func (c SiteConnectionStringKeyVaultReferenceSlotClient) WebAppsGetSiteConnectionStringKeyVaultReferencesSlotComplete(ctx context.Context, id SlotId) (WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCompleteResult, error) {
	return c.WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCompleteMatchingPredicate(ctx, id, ApiKVReferenceOperationPredicate{})
}

// WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteConnectionStringKeyVaultReferenceSlotClient) WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ApiKVReferenceOperationPredicate) (result WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCompleteResult, err error) {
	items := make([]ApiKVReference, 0)

	resp, err := c.WebAppsGetSiteConnectionStringKeyVaultReferencesSlot(ctx, id)
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

	result = WebAppsGetSiteConnectionStringKeyVaultReferencesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
