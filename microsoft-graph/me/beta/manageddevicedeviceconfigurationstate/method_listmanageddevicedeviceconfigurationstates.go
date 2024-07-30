package manageddevicedeviceconfigurationstate

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

type ListManagedDeviceDeviceConfigurationStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceConfigurationState
}

type ListManagedDeviceDeviceConfigurationStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceConfigurationState
}

type ListManagedDeviceDeviceConfigurationStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceDeviceConfigurationStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceDeviceConfigurationStates ...
func (c ManagedDeviceDeviceConfigurationStateClient) ListManagedDeviceDeviceConfigurationStates(ctx context.Context, id MeManagedDeviceId) (result ListManagedDeviceDeviceConfigurationStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListManagedDeviceDeviceConfigurationStatesCustomPager{},
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

// ListManagedDeviceDeviceConfigurationStatesComplete retrieves all the results into a single object
func (c ManagedDeviceDeviceConfigurationStateClient) ListManagedDeviceDeviceConfigurationStatesComplete(ctx context.Context, id MeManagedDeviceId) (ListManagedDeviceDeviceConfigurationStatesCompleteResult, error) {
	return c.ListManagedDeviceDeviceConfigurationStatesCompleteMatchingPredicate(ctx, id, DeviceConfigurationStateOperationPredicate{})
}

// ListManagedDeviceDeviceConfigurationStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceDeviceConfigurationStateClient) ListManagedDeviceDeviceConfigurationStatesCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate DeviceConfigurationStateOperationPredicate) (result ListManagedDeviceDeviceConfigurationStatesCompleteResult, err error) {
	items := make([]beta.DeviceConfigurationState, 0)

	resp, err := c.ListManagedDeviceDeviceConfigurationStates(ctx, id)
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

	result = ListManagedDeviceDeviceConfigurationStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
