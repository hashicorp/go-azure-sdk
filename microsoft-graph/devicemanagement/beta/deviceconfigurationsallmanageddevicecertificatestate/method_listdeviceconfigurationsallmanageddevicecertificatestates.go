package deviceconfigurationsallmanageddevicecertificatestate

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

type ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedAllDeviceCertificateState
}

type ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedAllDeviceCertificateState
}

type ListDeviceConfigurationsAllManagedDeviceCertificateStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceConfigurationsAllManagedDeviceCertificateStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceConfigurationsAllManagedDeviceCertificateStates ...
func (c DeviceConfigurationsAllManagedDeviceCertificateStateClient) ListDeviceConfigurationsAllManagedDeviceCertificateStates(ctx context.Context) (result ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceConfigurationsAllManagedDeviceCertificateStatesCustomPager{},
		Path:       "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates",
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
		Values *[]beta.ManagedAllDeviceCertificateState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationsAllManagedDeviceCertificateStatesComplete retrieves all the results into a single object
func (c DeviceConfigurationsAllManagedDeviceCertificateStateClient) ListDeviceConfigurationsAllManagedDeviceCertificateStatesComplete(ctx context.Context) (ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteResult, error) {
	return c.ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteMatchingPredicate(ctx, ManagedAllDeviceCertificateStateOperationPredicate{})
}

// ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationsAllManagedDeviceCertificateStateClient) ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteMatchingPredicate(ctx context.Context, predicate ManagedAllDeviceCertificateStateOperationPredicate) (result ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteResult, err error) {
	items := make([]beta.ManagedAllDeviceCertificateState, 0)

	resp, err := c.ListDeviceConfigurationsAllManagedDeviceCertificateStates(ctx)
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

	result = ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
