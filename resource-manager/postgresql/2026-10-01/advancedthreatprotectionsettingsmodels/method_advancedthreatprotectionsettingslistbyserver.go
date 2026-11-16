package advancedthreatprotectionsettingsmodels

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionSettingsListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AdvancedThreatProtectionSettingsModel
}

type AdvancedThreatProtectionSettingsListByServerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AdvancedThreatProtectionSettingsModel
}

type AdvancedThreatProtectionSettingsListByServerCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AdvancedThreatProtectionSettingsListByServerCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AdvancedThreatProtectionSettingsListByServer ...
func (c AdvancedThreatProtectionSettingsModelsClient) AdvancedThreatProtectionSettingsListByServer(ctx context.Context, id FlexibleServerId) (result AdvancedThreatProtectionSettingsListByServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AdvancedThreatProtectionSettingsListByServerCustomPager{},
		Path:       fmt.Sprintf("%s/advancedThreatProtectionSettings", id.ID()),
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
		Values *[]AdvancedThreatProtectionSettingsModel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AdvancedThreatProtectionSettingsListByServerComplete retrieves all the results into a single object
func (c AdvancedThreatProtectionSettingsModelsClient) AdvancedThreatProtectionSettingsListByServerComplete(ctx context.Context, id FlexibleServerId) (AdvancedThreatProtectionSettingsListByServerCompleteResult, error) {
	return c.AdvancedThreatProtectionSettingsListByServerCompleteMatchingPredicate(ctx, id, AdvancedThreatProtectionSettingsModelOperationPredicate{})
}

// AdvancedThreatProtectionSettingsListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdvancedThreatProtectionSettingsModelsClient) AdvancedThreatProtectionSettingsListByServerCompleteMatchingPredicate(ctx context.Context, id FlexibleServerId, predicate AdvancedThreatProtectionSettingsModelOperationPredicate) (result AdvancedThreatProtectionSettingsListByServerCompleteResult, err error) {
	items := make([]AdvancedThreatProtectionSettingsModel, 0)

	resp, err := c.AdvancedThreatProtectionSettingsListByServer(ctx, id)
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

	result = AdvancedThreatProtectionSettingsListByServerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
