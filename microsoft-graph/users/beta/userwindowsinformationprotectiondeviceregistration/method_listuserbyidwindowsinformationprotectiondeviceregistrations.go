package userwindowsinformationprotectiondeviceregistration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdWindowsInformationProtectionDeviceRegistrationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.WindowsInformationProtectionDeviceRegistrationCollectionResponse
}

type ListUserByIdWindowsInformationProtectionDeviceRegistrationsCompleteResult struct {
	Items []models.WindowsInformationProtectionDeviceRegistrationCollectionResponse
}

// ListUserByIdWindowsInformationProtectionDeviceRegistrations ...
func (c UserWindowsInformationProtectionDeviceRegistrationClient) ListUserByIdWindowsInformationProtectionDeviceRegistrations(ctx context.Context, id UserId) (result ListUserByIdWindowsInformationProtectionDeviceRegistrationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.WindowsInformationProtectionDeviceRegistrationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdWindowsInformationProtectionDeviceRegistrationsComplete retrieves all the results into a single object
func (c UserWindowsInformationProtectionDeviceRegistrationClient) ListUserByIdWindowsInformationProtectionDeviceRegistrationsComplete(ctx context.Context, id UserId) (ListUserByIdWindowsInformationProtectionDeviceRegistrationsCompleteResult, error) {
	return c.ListUserByIdWindowsInformationProtectionDeviceRegistrationsCompleteMatchingPredicate(ctx, id, models.WindowsInformationProtectionDeviceRegistrationCollectionResponseOperationPredicate{})
}

// ListUserByIdWindowsInformationProtectionDeviceRegistrationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserWindowsInformationProtectionDeviceRegistrationClient) ListUserByIdWindowsInformationProtectionDeviceRegistrationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.WindowsInformationProtectionDeviceRegistrationCollectionResponseOperationPredicate) (result ListUserByIdWindowsInformationProtectionDeviceRegistrationsCompleteResult, err error) {
	items := make([]models.WindowsInformationProtectionDeviceRegistrationCollectionResponse, 0)

	resp, err := c.ListUserByIdWindowsInformationProtectionDeviceRegistrations(ctx, id)
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

	result = ListUserByIdWindowsInformationProtectionDeviceRegistrationsCompleteResult{
		Items: items,
	}
	return
}
