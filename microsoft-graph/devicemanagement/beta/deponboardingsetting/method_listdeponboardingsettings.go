package deponboardingsetting

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

type ListDepOnboardingSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DepOnboardingSetting
}

type ListDepOnboardingSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DepOnboardingSetting
}

type ListDepOnboardingSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDepOnboardingSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDepOnboardingSettings ...
func (c DepOnboardingSettingClient) ListDepOnboardingSettings(ctx context.Context) (result ListDepOnboardingSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDepOnboardingSettingsCustomPager{},
		Path:       "/deviceManagement/depOnboardingSettings",
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
		Values *[]beta.DepOnboardingSetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDepOnboardingSettingsComplete retrieves all the results into a single object
func (c DepOnboardingSettingClient) ListDepOnboardingSettingsComplete(ctx context.Context) (ListDepOnboardingSettingsCompleteResult, error) {
	return c.ListDepOnboardingSettingsCompleteMatchingPredicate(ctx, DepOnboardingSettingOperationPredicate{})
}

// ListDepOnboardingSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DepOnboardingSettingClient) ListDepOnboardingSettingsCompleteMatchingPredicate(ctx context.Context, predicate DepOnboardingSettingOperationPredicate) (result ListDepOnboardingSettingsCompleteResult, err error) {
	items := make([]beta.DepOnboardingSetting, 0)

	resp, err := c.ListDepOnboardingSettings(ctx)
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

	result = ListDepOnboardingSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
