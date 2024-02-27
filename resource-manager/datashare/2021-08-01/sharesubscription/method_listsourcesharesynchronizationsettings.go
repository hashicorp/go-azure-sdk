package sharesubscription

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSourceShareSynchronizationSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SourceShareSynchronizationSetting
}

type ListSourceShareSynchronizationSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SourceShareSynchronizationSetting
}

// ListSourceShareSynchronizationSettings ...
func (c ShareSubscriptionClient) ListSourceShareSynchronizationSettings(ctx context.Context, id ShareSubscriptionId) (result ListSourceShareSynchronizationSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/listSourceShareSynchronizationSettings", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]SourceShareSynchronizationSetting, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := unmarshalSourceShareSynchronizationSettingImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for SourceShareSynchronizationSetting (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListSourceShareSynchronizationSettingsComplete retrieves all the results into a single object
func (c ShareSubscriptionClient) ListSourceShareSynchronizationSettingsComplete(ctx context.Context, id ShareSubscriptionId) (ListSourceShareSynchronizationSettingsCompleteResult, error) {
	return c.ListSourceShareSynchronizationSettingsCompleteMatchingPredicate(ctx, id, SourceShareSynchronizationSettingOperationPredicate{})
}

// ListSourceShareSynchronizationSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ShareSubscriptionClient) ListSourceShareSynchronizationSettingsCompleteMatchingPredicate(ctx context.Context, id ShareSubscriptionId, predicate SourceShareSynchronizationSettingOperationPredicate) (result ListSourceShareSynchronizationSettingsCompleteResult, err error) {
	items := make([]SourceShareSynchronizationSetting, 0)

	resp, err := c.ListSourceShareSynchronizationSettings(ctx, id)
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

	result = ListSourceShareSynchronizationSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
