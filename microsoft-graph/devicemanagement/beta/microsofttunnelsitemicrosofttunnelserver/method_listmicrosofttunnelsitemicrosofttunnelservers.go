package microsofttunnelsitemicrosofttunnelserver

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

type ListMicrosoftTunnelSiteMicrosoftTunnelServersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MicrosoftTunnelServer
}

type ListMicrosoftTunnelSiteMicrosoftTunnelServersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MicrosoftTunnelServer
}

type ListMicrosoftTunnelSiteMicrosoftTunnelServersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMicrosoftTunnelSiteMicrosoftTunnelServersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMicrosoftTunnelSiteMicrosoftTunnelServers ...
func (c MicrosoftTunnelSiteMicrosoftTunnelServerClient) ListMicrosoftTunnelSiteMicrosoftTunnelServers(ctx context.Context, id DeviceManagementMicrosoftTunnelSiteId) (result ListMicrosoftTunnelSiteMicrosoftTunnelServersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMicrosoftTunnelSiteMicrosoftTunnelServersCustomPager{},
		Path:       fmt.Sprintf("%s/microsoftTunnelServers", id.ID()),
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
		Values *[]beta.MicrosoftTunnelServer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMicrosoftTunnelSiteMicrosoftTunnelServersComplete retrieves all the results into a single object
func (c MicrosoftTunnelSiteMicrosoftTunnelServerClient) ListMicrosoftTunnelSiteMicrosoftTunnelServersComplete(ctx context.Context, id DeviceManagementMicrosoftTunnelSiteId) (ListMicrosoftTunnelSiteMicrosoftTunnelServersCompleteResult, error) {
	return c.ListMicrosoftTunnelSiteMicrosoftTunnelServersCompleteMatchingPredicate(ctx, id, MicrosoftTunnelServerOperationPredicate{})
}

// ListMicrosoftTunnelSiteMicrosoftTunnelServersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MicrosoftTunnelSiteMicrosoftTunnelServerClient) ListMicrosoftTunnelSiteMicrosoftTunnelServersCompleteMatchingPredicate(ctx context.Context, id DeviceManagementMicrosoftTunnelSiteId, predicate MicrosoftTunnelServerOperationPredicate) (result ListMicrosoftTunnelSiteMicrosoftTunnelServersCompleteResult, err error) {
	items := make([]beta.MicrosoftTunnelServer, 0)

	resp, err := c.ListMicrosoftTunnelSiteMicrosoftTunnelServers(ctx, id)
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

	result = ListMicrosoftTunnelSiteMicrosoftTunnelServersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
