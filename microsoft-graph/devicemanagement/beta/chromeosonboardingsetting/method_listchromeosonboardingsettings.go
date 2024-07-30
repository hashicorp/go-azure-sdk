package chromeosonboardingsetting

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

type ListChromeOSOnboardingSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ChromeOSOnboardingSettings
}

type ListChromeOSOnboardingSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ChromeOSOnboardingSettings
}

type ListChromeOSOnboardingSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChromeOSOnboardingSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChromeOSOnboardingSettings ...
func (c ChromeOSOnboardingSettingClient) ListChromeOSOnboardingSettings(ctx context.Context) (result ListChromeOSOnboardingSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListChromeOSOnboardingSettingsCustomPager{},
		Path:       "/deviceManagement/chromeOSOnboardingSettings",
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
		Values *[]beta.ChromeOSOnboardingSettings `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListChromeOSOnboardingSettingsComplete retrieves all the results into a single object
func (c ChromeOSOnboardingSettingClient) ListChromeOSOnboardingSettingsComplete(ctx context.Context) (ListChromeOSOnboardingSettingsCompleteResult, error) {
	return c.ListChromeOSOnboardingSettingsCompleteMatchingPredicate(ctx, ChromeOSOnboardingSettingsOperationPredicate{})
}

// ListChromeOSOnboardingSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChromeOSOnboardingSettingClient) ListChromeOSOnboardingSettingsCompleteMatchingPredicate(ctx context.Context, predicate ChromeOSOnboardingSettingsOperationPredicate) (result ListChromeOSOnboardingSettingsCompleteResult, err error) {
	items := make([]beta.ChromeOSOnboardingSettings, 0)

	resp, err := c.ListChromeOSOnboardingSettings(ctx)
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

	result = ListChromeOSOnboardingSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
