package mobilethreatdefenseconnector

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

type ListMobileThreatDefenseConnectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MobileThreatDefenseConnector
}

type ListMobileThreatDefenseConnectorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MobileThreatDefenseConnector
}

type ListMobileThreatDefenseConnectorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileThreatDefenseConnectorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileThreatDefenseConnectors ...
func (c MobileThreatDefenseConnectorClient) ListMobileThreatDefenseConnectors(ctx context.Context) (result ListMobileThreatDefenseConnectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMobileThreatDefenseConnectorsCustomPager{},
		Path:       "/deviceManagement/mobileThreatDefenseConnectors",
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
		Values *[]beta.MobileThreatDefenseConnector `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileThreatDefenseConnectorsComplete retrieves all the results into a single object
func (c MobileThreatDefenseConnectorClient) ListMobileThreatDefenseConnectorsComplete(ctx context.Context) (ListMobileThreatDefenseConnectorsCompleteResult, error) {
	return c.ListMobileThreatDefenseConnectorsCompleteMatchingPredicate(ctx, MobileThreatDefenseConnectorOperationPredicate{})
}

// ListMobileThreatDefenseConnectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileThreatDefenseConnectorClient) ListMobileThreatDefenseConnectorsCompleteMatchingPredicate(ctx context.Context, predicate MobileThreatDefenseConnectorOperationPredicate) (result ListMobileThreatDefenseConnectorsCompleteResult, err error) {
	items := make([]beta.MobileThreatDefenseConnector, 0)

	resp, err := c.ListMobileThreatDefenseConnectors(ctx)
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

	result = ListMobileThreatDefenseConnectorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
