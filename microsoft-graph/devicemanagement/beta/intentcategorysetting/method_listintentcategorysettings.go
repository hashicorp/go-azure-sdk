package intentcategorysetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListIntentCategorySettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingInstance
}

type ListIntentCategorySettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingInstance
}

type ListIntentCategorySettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntentCategorySettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntentCategorySettings ...
func (c IntentCategorySettingClient) ListIntentCategorySettings(ctx context.Context, id DeviceManagementIntentIdCategoryId) (result ListIntentCategorySettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIntentCategorySettingsCustomPager{},
		Path:       fmt.Sprintf("%s/settings", id.ID()),
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
		Values *[]beta.DeviceManagementSettingInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIntentCategorySettingsComplete retrieves all the results into a single object
func (c IntentCategorySettingClient) ListIntentCategorySettingsComplete(ctx context.Context, id DeviceManagementIntentIdCategoryId) (ListIntentCategorySettingsCompleteResult, error) {
	return c.ListIntentCategorySettingsCompleteMatchingPredicate(ctx, id, DeviceManagementSettingInstanceOperationPredicate{})
}

// ListIntentCategorySettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntentCategorySettingClient) ListIntentCategorySettingsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementIntentIdCategoryId, predicate DeviceManagementSettingInstanceOperationPredicate) (result ListIntentCategorySettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingInstance, 0)

	resp, err := c.ListIntentCategorySettings(ctx, id)
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

	result = ListIntentCategorySettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
