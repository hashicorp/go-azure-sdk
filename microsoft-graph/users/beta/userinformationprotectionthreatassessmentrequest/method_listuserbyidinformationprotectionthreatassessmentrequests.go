package userinformationprotectionthreatassessmentrequest

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

type ListUserByIdInformationProtectionThreatAssessmentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ThreatAssessmentRequestCollectionResponse
}

type ListUserByIdInformationProtectionThreatAssessmentRequestsCompleteResult struct {
	Items []models.ThreatAssessmentRequestCollectionResponse
}

// ListUserByIdInformationProtectionThreatAssessmentRequests ...
func (c UserInformationProtectionThreatAssessmentRequestClient) ListUserByIdInformationProtectionThreatAssessmentRequests(ctx context.Context, id UserId) (result ListUserByIdInformationProtectionThreatAssessmentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/informationProtection/threatAssessmentRequests", id.ID()),
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
		Values *[]models.ThreatAssessmentRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdInformationProtectionThreatAssessmentRequestsComplete retrieves all the results into a single object
func (c UserInformationProtectionThreatAssessmentRequestClient) ListUserByIdInformationProtectionThreatAssessmentRequestsComplete(ctx context.Context, id UserId) (ListUserByIdInformationProtectionThreatAssessmentRequestsCompleteResult, error) {
	return c.ListUserByIdInformationProtectionThreatAssessmentRequestsCompleteMatchingPredicate(ctx, id, models.ThreatAssessmentRequestCollectionResponseOperationPredicate{})
}

// ListUserByIdInformationProtectionThreatAssessmentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInformationProtectionThreatAssessmentRequestClient) ListUserByIdInformationProtectionThreatAssessmentRequestsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ThreatAssessmentRequestCollectionResponseOperationPredicate) (result ListUserByIdInformationProtectionThreatAssessmentRequestsCompleteResult, err error) {
	items := make([]models.ThreatAssessmentRequestCollectionResponse, 0)

	resp, err := c.ListUserByIdInformationProtectionThreatAssessmentRequests(ctx, id)
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

	result = ListUserByIdInformationProtectionThreatAssessmentRequestsCompleteResult{
		Items: items,
	}
	return
}
