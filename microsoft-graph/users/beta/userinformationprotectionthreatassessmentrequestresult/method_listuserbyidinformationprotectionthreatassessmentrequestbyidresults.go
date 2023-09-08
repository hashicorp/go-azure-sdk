package userinformationprotectionthreatassessmentrequestresult

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

type ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ThreatAssessmentResultCollectionResponse
}

type ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteResult struct {
	Items []models.ThreatAssessmentResultCollectionResponse
}

// ListUserByIdInformationProtectionThreatAssessmentRequestByIdResults ...
func (c UserInformationProtectionThreatAssessmentRequestResultClient) ListUserByIdInformationProtectionThreatAssessmentRequestByIdResults(ctx context.Context, id UserInformationProtectionThreatAssessmentRequestId) (result ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/results", id.ID()),
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
		Values *[]models.ThreatAssessmentResultCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsComplete retrieves all the results into a single object
func (c UserInformationProtectionThreatAssessmentRequestResultClient) ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsComplete(ctx context.Context, id UserInformationProtectionThreatAssessmentRequestId) (ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteResult, error) {
	return c.ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteMatchingPredicate(ctx, id, models.ThreatAssessmentResultCollectionResponseOperationPredicate{})
}

// ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInformationProtectionThreatAssessmentRequestResultClient) ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteMatchingPredicate(ctx context.Context, id UserInformationProtectionThreatAssessmentRequestId, predicate models.ThreatAssessmentResultCollectionResponseOperationPredicate) (result ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteResult, err error) {
	items := make([]models.ThreatAssessmentResultCollectionResponse, 0)

	resp, err := c.ListUserByIdInformationProtectionThreatAssessmentRequestByIdResults(ctx, id)
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

	result = ListUserByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteResult{
		Items: items,
	}
	return
}
