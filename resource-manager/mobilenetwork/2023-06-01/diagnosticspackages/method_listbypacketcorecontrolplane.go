package diagnosticspackages

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
	Model        *[]DiagnosticsPackage
}

type ListByPacketCoreControlPlaneCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DiagnosticsPackage
}

// ListByPacketCoreControlPlane ...
func (c DiagnosticsPackagesClient) ListByPacketCoreControlPlane(ctx context.Context, id PacketCoreControlPlaneId) (result ListByPacketCoreControlPlaneOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/diagnosticsPackages", id.ID()),
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
		Values *[]DiagnosticsPackage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByPacketCoreControlPlaneComplete retrieves all the results into a single object
func (c DiagnosticsPackagesClient) ListByPacketCoreControlPlaneComplete(ctx context.Context, id PacketCoreControlPlaneId) (ListByPacketCoreControlPlaneCompleteResult, error) {
	return c.ListByPacketCoreControlPlaneCompleteMatchingPredicate(ctx, id, DiagnosticsPackageOperationPredicate{})
}

// ListByPacketCoreControlPlaneCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsPackagesClient) ListByPacketCoreControlPlaneCompleteMatchingPredicate(ctx context.Context, id PacketCoreControlPlaneId, predicate DiagnosticsPackageOperationPredicate) (result ListByPacketCoreControlPlaneCompleteResult, err error) {
	items := make([]DiagnosticsPackage, 0)

	resp, err := c.ListByPacketCoreControlPlane(ctx, id)
	if err != nil {
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
