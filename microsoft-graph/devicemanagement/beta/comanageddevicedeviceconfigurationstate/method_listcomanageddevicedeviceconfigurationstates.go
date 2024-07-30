package comanageddevicedeviceconfigurationstate

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

type ListComanagedDeviceDeviceConfigurationStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceConfigurationState
}

type ListComanagedDeviceDeviceConfigurationStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceConfigurationState
}

type ListComanagedDeviceDeviceConfigurationStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceDeviceConfigurationStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceDeviceConfigurationStates ...
func (c ComanagedDeviceDeviceConfigurationStateClient) ListComanagedDeviceDeviceConfigurationStates(ctx context.Context, id DeviceManagementComanagedDeviceId) (result ListComanagedDeviceDeviceConfigurationStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListComanagedDeviceDeviceConfigurationStatesCustomPager{},
		Path:       fmt.Sprintf("%s/deviceConfigurationStates", id.ID()),
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
		Values *[]beta.DeviceConfigurationState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceDeviceConfigurationStatesComplete retrieves all the results into a single object
func (c ComanagedDeviceDeviceConfigurationStateClient) ListComanagedDeviceDeviceConfigurationStatesComplete(ctx context.Context, id DeviceManagementComanagedDeviceId) (ListComanagedDeviceDeviceConfigurationStatesCompleteResult, error) {
	return c.ListComanagedDeviceDeviceConfigurationStatesCompleteMatchingPredicate(ctx, id, DeviceConfigurationStateOperationPredicate{})
}

// ListComanagedDeviceDeviceConfigurationStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceDeviceConfigurationStateClient) ListComanagedDeviceDeviceConfigurationStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementComanagedDeviceId, predicate DeviceConfigurationStateOperationPredicate) (result ListComanagedDeviceDeviceConfigurationStatesCompleteResult, err error) {
	items := make([]beta.DeviceConfigurationState, 0)

	resp, err := c.ListComanagedDeviceDeviceConfigurationStates(ctx, id)
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

	result = ListComanagedDeviceDeviceConfigurationStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
