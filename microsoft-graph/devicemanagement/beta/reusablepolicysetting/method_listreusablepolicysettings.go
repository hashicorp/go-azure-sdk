package reusablepolicysetting

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

type ListReusablePolicySettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementReusablePolicySetting
}

type ListReusablePolicySettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementReusablePolicySetting
}

type ListReusablePolicySettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListReusablePolicySettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListReusablePolicySettings ...
func (c ReusablePolicySettingClient) ListReusablePolicySettings(ctx context.Context) (result ListReusablePolicySettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListReusablePolicySettingsCustomPager{},
		Path:       "/deviceManagement/reusablePolicySettings",
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
		Values *[]beta.DeviceManagementReusablePolicySetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListReusablePolicySettingsComplete retrieves all the results into a single object
func (c ReusablePolicySettingClient) ListReusablePolicySettingsComplete(ctx context.Context) (ListReusablePolicySettingsCompleteResult, error) {
	return c.ListReusablePolicySettingsCompleteMatchingPredicate(ctx, DeviceManagementReusablePolicySettingOperationPredicate{})
}

// ListReusablePolicySettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReusablePolicySettingClient) ListReusablePolicySettingsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementReusablePolicySettingOperationPredicate) (result ListReusablePolicySettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementReusablePolicySetting, 0)

	resp, err := c.ListReusablePolicySettings(ctx)
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

	result = ListReusablePolicySettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
