package siteconnectionstringkeyvaultreference

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

type WebAppsGetSiteConnectionStringKeyVaultReferencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApiKVReference
}

type WebAppsGetSiteConnectionStringKeyVaultReferencesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApiKVReference
}

type WebAppsGetSiteConnectionStringKeyVaultReferencesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsGetSiteConnectionStringKeyVaultReferencesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsGetSiteConnectionStringKeyVaultReferences ...
func (c SiteConnectionStringKeyVaultReferenceClient) WebAppsGetSiteConnectionStringKeyVaultReferences(ctx context.Context, id commonids.AppServiceId) (result WebAppsGetSiteConnectionStringKeyVaultReferencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsGetSiteConnectionStringKeyVaultReferencesCustomPager{},
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

// WebAppsGetSiteConnectionStringKeyVaultReferencesComplete retrieves all the results into a single object
func (c SiteConnectionStringKeyVaultReferenceClient) WebAppsGetSiteConnectionStringKeyVaultReferencesComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsGetSiteConnectionStringKeyVaultReferencesCompleteResult, error) {
	return c.WebAppsGetSiteConnectionStringKeyVaultReferencesCompleteMatchingPredicate(ctx, id, ApiKVReferenceOperationPredicate{})
}

// WebAppsGetSiteConnectionStringKeyVaultReferencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteConnectionStringKeyVaultReferenceClient) WebAppsGetSiteConnectionStringKeyVaultReferencesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate ApiKVReferenceOperationPredicate) (result WebAppsGetSiteConnectionStringKeyVaultReferencesCompleteResult, err error) {
	items := make([]ApiKVReference, 0)

	resp, err := c.WebAppsGetSiteConnectionStringKeyVaultReferences(ctx, id)
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

	result = WebAppsGetSiteConnectionStringKeyVaultReferencesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
