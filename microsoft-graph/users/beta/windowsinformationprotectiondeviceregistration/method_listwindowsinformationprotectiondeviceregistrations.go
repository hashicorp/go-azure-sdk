package windowsinformationprotectiondeviceregistration

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

type ListWindowsInformationProtectionDeviceRegistrationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsInformationProtectionDeviceRegistration
}

type ListWindowsInformationProtectionDeviceRegistrationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsInformationProtectionDeviceRegistration
}

type ListWindowsInformationProtectionDeviceRegistrationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsInformationProtectionDeviceRegistrationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsInformationProtectionDeviceRegistrations ...
func (c WindowsInformationProtectionDeviceRegistrationClient) ListWindowsInformationProtectionDeviceRegistrations(ctx context.Context, id UserId) (result ListWindowsInformationProtectionDeviceRegistrationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsInformationProtectionDeviceRegistrationsCustomPager{},
		Path:       fmt.Sprintf("%s/windowsInformationProtectionDeviceRegistrations", id.ID()),
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
		Values *[]beta.WindowsInformationProtectionDeviceRegistration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsInformationProtectionDeviceRegistrationsComplete retrieves all the results into a single object
func (c WindowsInformationProtectionDeviceRegistrationClient) ListWindowsInformationProtectionDeviceRegistrationsComplete(ctx context.Context, id UserId) (ListWindowsInformationProtectionDeviceRegistrationsCompleteResult, error) {
	return c.ListWindowsInformationProtectionDeviceRegistrationsCompleteMatchingPredicate(ctx, id, WindowsInformationProtectionDeviceRegistrationOperationPredicate{})
}

// ListWindowsInformationProtectionDeviceRegistrationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsInformationProtectionDeviceRegistrationClient) ListWindowsInformationProtectionDeviceRegistrationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate WindowsInformationProtectionDeviceRegistrationOperationPredicate) (result ListWindowsInformationProtectionDeviceRegistrationsCompleteResult, err error) {
	items := make([]beta.WindowsInformationProtectionDeviceRegistration, 0)

	resp, err := c.ListWindowsInformationProtectionDeviceRegistrations(ctx, id)
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

	result = ListWindowsInformationProtectionDeviceRegistrationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
