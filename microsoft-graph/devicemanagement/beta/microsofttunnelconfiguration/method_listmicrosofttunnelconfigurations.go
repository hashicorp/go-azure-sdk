package microsofttunnelconfiguration

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

type ListMicrosoftTunnelConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MicrosoftTunnelConfiguration
}

type ListMicrosoftTunnelConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MicrosoftTunnelConfiguration
}

type ListMicrosoftTunnelConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMicrosoftTunnelConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMicrosoftTunnelConfigurations ...
func (c MicrosoftTunnelConfigurationClient) ListMicrosoftTunnelConfigurations(ctx context.Context) (result ListMicrosoftTunnelConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMicrosoftTunnelConfigurationsCustomPager{},
		Path:       "/deviceManagement/microsoftTunnelConfigurations",
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
		Values *[]beta.MicrosoftTunnelConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMicrosoftTunnelConfigurationsComplete retrieves all the results into a single object
func (c MicrosoftTunnelConfigurationClient) ListMicrosoftTunnelConfigurationsComplete(ctx context.Context) (ListMicrosoftTunnelConfigurationsCompleteResult, error) {
	return c.ListMicrosoftTunnelConfigurationsCompleteMatchingPredicate(ctx, MicrosoftTunnelConfigurationOperationPredicate{})
}

// ListMicrosoftTunnelConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MicrosoftTunnelConfigurationClient) ListMicrosoftTunnelConfigurationsCompleteMatchingPredicate(ctx context.Context, predicate MicrosoftTunnelConfigurationOperationPredicate) (result ListMicrosoftTunnelConfigurationsCompleteResult, err error) {
	items := make([]beta.MicrosoftTunnelConfiguration, 0)

	resp, err := c.ListMicrosoftTunnelConfigurations(ctx)
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

	result = ListMicrosoftTunnelConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
