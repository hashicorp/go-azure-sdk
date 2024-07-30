package domainjoinconnector

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

type ListDomainJoinConnectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementDomainJoinConnector
}

type ListDomainJoinConnectorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementDomainJoinConnector
}

type ListDomainJoinConnectorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDomainJoinConnectorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDomainJoinConnectors ...
func (c DomainJoinConnectorClient) ListDomainJoinConnectors(ctx context.Context) (result ListDomainJoinConnectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDomainJoinConnectorsCustomPager{},
		Path:       "/deviceManagement/domainJoinConnectors",
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
		Values *[]beta.DeviceManagementDomainJoinConnector `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDomainJoinConnectorsComplete retrieves all the results into a single object
func (c DomainJoinConnectorClient) ListDomainJoinConnectorsComplete(ctx context.Context) (ListDomainJoinConnectorsCompleteResult, error) {
	return c.ListDomainJoinConnectorsCompleteMatchingPredicate(ctx, DeviceManagementDomainJoinConnectorOperationPredicate{})
}

// ListDomainJoinConnectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DomainJoinConnectorClient) ListDomainJoinConnectorsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementDomainJoinConnectorOperationPredicate) (result ListDomainJoinConnectorsCompleteResult, err error) {
	items := make([]beta.DeviceManagementDomainJoinConnector, 0)

	resp, err := c.ListDomainJoinConnectors(ctx)
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

	result = ListDomainJoinConnectorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
