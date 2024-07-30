package virtualendpointexternalpartnersetting

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

type ListVirtualEndpointExternalPartnerSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCExternalPartnerSetting
}

type ListVirtualEndpointExternalPartnerSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCExternalPartnerSetting
}

type ListVirtualEndpointExternalPartnerSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointExternalPartnerSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointExternalPartnerSettings ...
func (c VirtualEndpointExternalPartnerSettingClient) ListVirtualEndpointExternalPartnerSettings(ctx context.Context) (result ListVirtualEndpointExternalPartnerSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointExternalPartnerSettingsCustomPager{},
		Path:       "/deviceManagement/virtualEndpoint/externalPartnerSettings",
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
		Values *[]beta.CloudPCExternalPartnerSetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointExternalPartnerSettingsComplete retrieves all the results into a single object
func (c VirtualEndpointExternalPartnerSettingClient) ListVirtualEndpointExternalPartnerSettingsComplete(ctx context.Context) (ListVirtualEndpointExternalPartnerSettingsCompleteResult, error) {
	return c.ListVirtualEndpointExternalPartnerSettingsCompleteMatchingPredicate(ctx, CloudPCExternalPartnerSettingOperationPredicate{})
}

// ListVirtualEndpointExternalPartnerSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointExternalPartnerSettingClient) ListVirtualEndpointExternalPartnerSettingsCompleteMatchingPredicate(ctx context.Context, predicate CloudPCExternalPartnerSettingOperationPredicate) (result ListVirtualEndpointExternalPartnerSettingsCompleteResult, err error) {
	items := make([]beta.CloudPCExternalPartnerSetting, 0)

	resp, err := c.ListVirtualEndpointExternalPartnerSettings(ctx)
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

	result = ListVirtualEndpointExternalPartnerSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
