package microsofttunnelserverlogcollectionrespons

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

type ListMicrosoftTunnelServerLogCollectionResponsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MicrosoftTunnelServerLogCollectionResponse
}

type ListMicrosoftTunnelServerLogCollectionResponsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MicrosoftTunnelServerLogCollectionResponse
}

type ListMicrosoftTunnelServerLogCollectionResponsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMicrosoftTunnelServerLogCollectionResponsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMicrosoftTunnelServerLogCollectionRespons ...
func (c MicrosoftTunnelServerLogCollectionResponsClient) ListMicrosoftTunnelServerLogCollectionRespons(ctx context.Context) (result ListMicrosoftTunnelServerLogCollectionResponsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMicrosoftTunnelServerLogCollectionResponsCustomPager{},
		Path:       "/deviceManagement/microsoftTunnelServerLogCollectionResponses",
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
		Values *[]beta.MicrosoftTunnelServerLogCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMicrosoftTunnelServerLogCollectionResponsComplete retrieves all the results into a single object
func (c MicrosoftTunnelServerLogCollectionResponsClient) ListMicrosoftTunnelServerLogCollectionResponsComplete(ctx context.Context) (ListMicrosoftTunnelServerLogCollectionResponsCompleteResult, error) {
	return c.ListMicrosoftTunnelServerLogCollectionResponsCompleteMatchingPredicate(ctx, MicrosoftTunnelServerLogCollectionResponseOperationPredicate{})
}

// ListMicrosoftTunnelServerLogCollectionResponsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MicrosoftTunnelServerLogCollectionResponsClient) ListMicrosoftTunnelServerLogCollectionResponsCompleteMatchingPredicate(ctx context.Context, predicate MicrosoftTunnelServerLogCollectionResponseOperationPredicate) (result ListMicrosoftTunnelServerLogCollectionResponsCompleteResult, err error) {
	items := make([]beta.MicrosoftTunnelServerLogCollectionResponse, 0)

	resp, err := c.ListMicrosoftTunnelServerLogCollectionRespons(ctx)
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

	result = ListMicrosoftTunnelServerLogCollectionResponsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
