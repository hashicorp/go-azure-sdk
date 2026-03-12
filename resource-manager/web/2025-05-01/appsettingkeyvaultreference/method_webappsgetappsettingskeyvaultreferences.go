package appsettingkeyvaultreference

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

type WebAppsGetAppSettingsKeyVaultReferencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApiKVReference
}

type WebAppsGetAppSettingsKeyVaultReferencesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApiKVReference
}

type WebAppsGetAppSettingsKeyVaultReferencesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsGetAppSettingsKeyVaultReferencesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsGetAppSettingsKeyVaultReferences ...
func (c AppSettingKeyVaultReferenceClient) WebAppsGetAppSettingsKeyVaultReferences(ctx context.Context, id commonids.AppServiceId) (result WebAppsGetAppSettingsKeyVaultReferencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsGetAppSettingsKeyVaultReferencesCustomPager{},
		Path:       fmt.Sprintf("%s/config/configReferences/appSettings", id.ID()),
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

// WebAppsGetAppSettingsKeyVaultReferencesComplete retrieves all the results into a single object
func (c AppSettingKeyVaultReferenceClient) WebAppsGetAppSettingsKeyVaultReferencesComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsGetAppSettingsKeyVaultReferencesCompleteResult, error) {
	return c.WebAppsGetAppSettingsKeyVaultReferencesCompleteMatchingPredicate(ctx, id, ApiKVReferenceOperationPredicate{})
}

// WebAppsGetAppSettingsKeyVaultReferencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppSettingKeyVaultReferenceClient) WebAppsGetAppSettingsKeyVaultReferencesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate ApiKVReferenceOperationPredicate) (result WebAppsGetAppSettingsKeyVaultReferencesCompleteResult, err error) {
	items := make([]ApiKVReference, 0)

	resp, err := c.WebAppsGetAppSettingsKeyVaultReferences(ctx, id)
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

	result = WebAppsGetAppSettingsKeyVaultReferencesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
