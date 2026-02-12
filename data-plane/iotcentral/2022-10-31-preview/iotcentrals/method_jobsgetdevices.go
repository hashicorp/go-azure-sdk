package iotcentrals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobsGetDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]JobDeviceStatus
}

type JobsGetDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []JobDeviceStatus
}

type JobsGetDevicesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *JobsGetDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// JobsGetDevices ...
func (c IotcentralsClient) JobsGetDevices(ctx context.Context, id JobId) (result JobsGetDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &JobsGetDevicesCustomPager{},
		Path:       fmt.Sprintf("%s/devices", id.Path()),
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
		Values *[]JobDeviceStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// JobsGetDevicesComplete retrieves all the results into a single object
func (c IotcentralsClient) JobsGetDevicesComplete(ctx context.Context, id JobId) (JobsGetDevicesCompleteResult, error) {
	return c.JobsGetDevicesCompleteMatchingPredicate(ctx, id, JobDeviceStatusOperationPredicate{})
}

// JobsGetDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) JobsGetDevicesCompleteMatchingPredicate(ctx context.Context, id JobId, predicate JobDeviceStatusOperationPredicate) (result JobsGetDevicesCompleteResult, err error) {
	items := make([]JobDeviceStatus, 0)

	resp, err := c.JobsGetDevices(ctx, id)
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

	result = JobsGetDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
