package iosupdatestatus

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

type ListIosUpdateStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IosUpdateDeviceStatus
}

type ListIosUpdateStatusCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IosUpdateDeviceStatus
}

type ListIosUpdateStatusCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIosUpdateStatusCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIosUpdateStatus ...
func (c IosUpdateStatusClient) ListIosUpdateStatus(ctx context.Context) (result ListIosUpdateStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIosUpdateStatusCustomPager{},
		Path:       "/deviceManagement/iosUpdateStatuses",
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
		Values *[]beta.IosUpdateDeviceStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIosUpdateStatusComplete retrieves all the results into a single object
func (c IosUpdateStatusClient) ListIosUpdateStatusComplete(ctx context.Context) (ListIosUpdateStatusCompleteResult, error) {
	return c.ListIosUpdateStatusCompleteMatchingPredicate(ctx, IosUpdateDeviceStatusOperationPredicate{})
}

// ListIosUpdateStatusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IosUpdateStatusClient) ListIosUpdateStatusCompleteMatchingPredicate(ctx context.Context, predicate IosUpdateDeviceStatusOperationPredicate) (result ListIosUpdateStatusCompleteResult, err error) {
	items := make([]beta.IosUpdateDeviceStatus, 0)

	resp, err := c.ListIosUpdateStatus(ctx)
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

	result = ListIosUpdateStatusCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
