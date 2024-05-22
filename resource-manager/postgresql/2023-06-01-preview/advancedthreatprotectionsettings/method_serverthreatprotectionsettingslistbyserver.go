package advancedthreatprotectionsettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerThreatProtectionSettingsListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServerThreatProtectionSettingsModel
}

type ServerThreatProtectionSettingsListByServerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ServerThreatProtectionSettingsModel
}

// ServerThreatProtectionSettingsListByServer ...
func (c AdvancedThreatProtectionSettingsClient) ServerThreatProtectionSettingsListByServer(ctx context.Context, id FlexibleServerId) (result ServerThreatProtectionSettingsListByServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]ServerThreatProtectionSettingsModel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ServerThreatProtectionSettingsListByServerComplete retrieves all the results into a single object
func (c AdvancedThreatProtectionSettingsClient) ServerThreatProtectionSettingsListByServerComplete(ctx context.Context, id FlexibleServerId) (ServerThreatProtectionSettingsListByServerCompleteResult, error) {
	return c.ServerThreatProtectionSettingsListByServerCompleteMatchingPredicate(ctx, id, ServerThreatProtectionSettingsModelOperationPredicate{})
}

// ServerThreatProtectionSettingsListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdvancedThreatProtectionSettingsClient) ServerThreatProtectionSettingsListByServerCompleteMatchingPredicate(ctx context.Context, id FlexibleServerId, predicate ServerThreatProtectionSettingsModelOperationPredicate) (result ServerThreatProtectionSettingsListByServerCompleteResult, err error) {
	items := make([]ServerThreatProtectionSettingsModel, 0)

	resp, err := c.ServerThreatProtectionSettingsListByServer(ctx, id)
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

	result = ServerThreatProtectionSettingsListByServerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
