package botconnection

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

type ListByBotServiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ConnectionSetting
}

type ListByBotServiceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ConnectionSetting
}

// ListByBotService ...
func (c BotConnectionClient) ListByBotService(ctx context.Context, id commonids.BotServiceId) (result ListByBotServiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/connections", id.ID()),
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
		Values *[]ConnectionSetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBotServiceComplete retrieves all the results into a single object
func (c BotConnectionClient) ListByBotServiceComplete(ctx context.Context, id commonids.BotServiceId) (ListByBotServiceCompleteResult, error) {
	return c.ListByBotServiceCompleteMatchingPredicate(ctx, id, ConnectionSettingOperationPredicate{})
}

// ListByBotServiceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BotConnectionClient) ListByBotServiceCompleteMatchingPredicate(ctx context.Context, id commonids.BotServiceId, predicate ConnectionSettingOperationPredicate) (result ListByBotServiceCompleteResult, err error) {
	items := make([]ConnectionSetting, 0)

	resp, err := c.ListByBotService(ctx, id)
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

	result = ListByBotServiceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
