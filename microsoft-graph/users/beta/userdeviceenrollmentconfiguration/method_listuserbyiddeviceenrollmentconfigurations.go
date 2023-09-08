package userdeviceenrollmentconfiguration

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

type ListUserByIdDeviceEnrollmentConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DeviceEnrollmentConfigurationCollectionResponse
}

type ListUserByIdDeviceEnrollmentConfigurationsCompleteResult struct {
	Items []models.DeviceEnrollmentConfigurationCollectionResponse
}

// ListUserByIdDeviceEnrollmentConfigurations ...
func (c UserDeviceEnrollmentConfigurationClient) ListUserByIdDeviceEnrollmentConfigurations(ctx context.Context, id UserId) (result ListUserByIdDeviceEnrollmentConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/deviceEnrollmentConfigurations", id.ID()),
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
		Values *[]models.DeviceEnrollmentConfigurationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdDeviceEnrollmentConfigurationsComplete retrieves all the results into a single object
func (c UserDeviceEnrollmentConfigurationClient) ListUserByIdDeviceEnrollmentConfigurationsComplete(ctx context.Context, id UserId) (ListUserByIdDeviceEnrollmentConfigurationsCompleteResult, error) {
	return c.ListUserByIdDeviceEnrollmentConfigurationsCompleteMatchingPredicate(ctx, id, models.DeviceEnrollmentConfigurationCollectionResponseOperationPredicate{})
}

// ListUserByIdDeviceEnrollmentConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserDeviceEnrollmentConfigurationClient) ListUserByIdDeviceEnrollmentConfigurationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.DeviceEnrollmentConfigurationCollectionResponseOperationPredicate) (result ListUserByIdDeviceEnrollmentConfigurationsCompleteResult, err error) {
	items := make([]models.DeviceEnrollmentConfigurationCollectionResponse, 0)

	resp, err := c.ListUserByIdDeviceEnrollmentConfigurations(ctx, id)
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

	result = ListUserByIdDeviceEnrollmentConfigurationsCompleteResult{
		Items: items,
	}
	return
}
