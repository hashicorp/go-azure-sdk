package microsofttunnelhealththreshold

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

type ListMicrosoftTunnelHealthThresholdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MicrosoftTunnelHealthThreshold
}

type ListMicrosoftTunnelHealthThresholdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MicrosoftTunnelHealthThreshold
}

type ListMicrosoftTunnelHealthThresholdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMicrosoftTunnelHealthThresholdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMicrosoftTunnelHealthThresholds ...
func (c MicrosoftTunnelHealthThresholdClient) ListMicrosoftTunnelHealthThresholds(ctx context.Context) (result ListMicrosoftTunnelHealthThresholdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMicrosoftTunnelHealthThresholdsCustomPager{},
		Path:       "/deviceManagement/microsoftTunnelHealthThresholds",
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
		Values *[]beta.MicrosoftTunnelHealthThreshold `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMicrosoftTunnelHealthThresholdsComplete retrieves all the results into a single object
func (c MicrosoftTunnelHealthThresholdClient) ListMicrosoftTunnelHealthThresholdsComplete(ctx context.Context) (ListMicrosoftTunnelHealthThresholdsCompleteResult, error) {
	return c.ListMicrosoftTunnelHealthThresholdsCompleteMatchingPredicate(ctx, MicrosoftTunnelHealthThresholdOperationPredicate{})
}

// ListMicrosoftTunnelHealthThresholdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MicrosoftTunnelHealthThresholdClient) ListMicrosoftTunnelHealthThresholdsCompleteMatchingPredicate(ctx context.Context, predicate MicrosoftTunnelHealthThresholdOperationPredicate) (result ListMicrosoftTunnelHealthThresholdsCompleteResult, err error) {
	items := make([]beta.MicrosoftTunnelHealthThreshold, 0)

	resp, err := c.ListMicrosoftTunnelHealthThresholds(ctx)
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

	result = ListMicrosoftTunnelHealthThresholdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
