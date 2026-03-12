package appsettingkeyvaultreferenceslot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsGetAppSettingsKeyVaultReferencesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApiKVReference
}

type WebAppsGetAppSettingsKeyVaultReferencesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApiKVReference
}

type WebAppsGetAppSettingsKeyVaultReferencesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsGetAppSettingsKeyVaultReferencesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsGetAppSettingsKeyVaultReferencesSlot ...
func (c AppSettingKeyVaultReferenceSlotClient) WebAppsGetAppSettingsKeyVaultReferencesSlot(ctx context.Context, id SlotId) (result WebAppsGetAppSettingsKeyVaultReferencesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsGetAppSettingsKeyVaultReferencesSlotCustomPager{},
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

// WebAppsGetAppSettingsKeyVaultReferencesSlotComplete retrieves all the results into a single object
func (c AppSettingKeyVaultReferenceSlotClient) WebAppsGetAppSettingsKeyVaultReferencesSlotComplete(ctx context.Context, id SlotId) (WebAppsGetAppSettingsKeyVaultReferencesSlotCompleteResult, error) {
	return c.WebAppsGetAppSettingsKeyVaultReferencesSlotCompleteMatchingPredicate(ctx, id, ApiKVReferenceOperationPredicate{})
}

// WebAppsGetAppSettingsKeyVaultReferencesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppSettingKeyVaultReferenceSlotClient) WebAppsGetAppSettingsKeyVaultReferencesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ApiKVReferenceOperationPredicate) (result WebAppsGetAppSettingsKeyVaultReferencesSlotCompleteResult, err error) {
	items := make([]ApiKVReference, 0)

	resp, err := c.WebAppsGetAppSettingsKeyVaultReferencesSlot(ctx, id)
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

	result = WebAppsGetAppSettingsKeyVaultReferencesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
