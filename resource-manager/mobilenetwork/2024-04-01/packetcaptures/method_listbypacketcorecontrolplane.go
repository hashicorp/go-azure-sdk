package packetcaptures

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByPacketCoreControlPlaneOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PacketCapture
}

type ListByPacketCoreControlPlaneCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PacketCapture
}

type ListByPacketCoreControlPlaneCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByPacketCoreControlPlaneCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByPacketCoreControlPlane ...
func (c PacketCapturesClient) ListByPacketCoreControlPlane(ctx context.Context, id PacketCoreControlPlaneId) (result ListByPacketCoreControlPlaneOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByPacketCoreControlPlaneCustomPager{},
		Path:       fmt.Sprintf("%s/packetCaptures", id.ID()),
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
		Values *[]PacketCapture `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByPacketCoreControlPlaneComplete retrieves all the results into a single object
func (c PacketCapturesClient) ListByPacketCoreControlPlaneComplete(ctx context.Context, id PacketCoreControlPlaneId) (ListByPacketCoreControlPlaneCompleteResult, error) {
	return c.ListByPacketCoreControlPlaneCompleteMatchingPredicate(ctx, id, PacketCaptureOperationPredicate{})
}

// ListByPacketCoreControlPlaneCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PacketCapturesClient) ListByPacketCoreControlPlaneCompleteMatchingPredicate(ctx context.Context, id PacketCoreControlPlaneId, predicate PacketCaptureOperationPredicate) (result ListByPacketCoreControlPlaneCompleteResult, err error) {
	items := make([]PacketCapture, 0)

	resp, err := c.ListByPacketCoreControlPlane(ctx, id)
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

	result = ListByPacketCoreControlPlaneCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
