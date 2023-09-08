package medeviceenrollmentconfigurationassignment

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

type ListMeDeviceEnrollmentConfigurationByIdAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EnrollmentConfigurationAssignmentCollectionResponse
}

type ListMeDeviceEnrollmentConfigurationByIdAssignmentsCompleteResult struct {
	Items []models.EnrollmentConfigurationAssignmentCollectionResponse
}

// ListMeDeviceEnrollmentConfigurationByIdAssignments ...
func (c MeDeviceEnrollmentConfigurationAssignmentClient) ListMeDeviceEnrollmentConfigurationByIdAssignments(ctx context.Context, id MeDeviceEnrollmentConfigurationId) (result ListMeDeviceEnrollmentConfigurationByIdAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]models.EnrollmentConfigurationAssignmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeDeviceEnrollmentConfigurationByIdAssignmentsComplete retrieves all the results into a single object
func (c MeDeviceEnrollmentConfigurationAssignmentClient) ListMeDeviceEnrollmentConfigurationByIdAssignmentsComplete(ctx context.Context, id MeDeviceEnrollmentConfigurationId) (ListMeDeviceEnrollmentConfigurationByIdAssignmentsCompleteResult, error) {
	return c.ListMeDeviceEnrollmentConfigurationByIdAssignmentsCompleteMatchingPredicate(ctx, id, models.EnrollmentConfigurationAssignmentCollectionResponseOperationPredicate{})
}

// ListMeDeviceEnrollmentConfigurationByIdAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeDeviceEnrollmentConfigurationAssignmentClient) ListMeDeviceEnrollmentConfigurationByIdAssignmentsCompleteMatchingPredicate(ctx context.Context, id MeDeviceEnrollmentConfigurationId, predicate models.EnrollmentConfigurationAssignmentCollectionResponseOperationPredicate) (result ListMeDeviceEnrollmentConfigurationByIdAssignmentsCompleteResult, err error) {
	items := make([]models.EnrollmentConfigurationAssignmentCollectionResponse, 0)

	resp, err := c.ListMeDeviceEnrollmentConfigurationByIdAssignments(ctx, id)
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

	result = ListMeDeviceEnrollmentConfigurationByIdAssignmentsCompleteResult{
		Items: items,
	}
	return
}
