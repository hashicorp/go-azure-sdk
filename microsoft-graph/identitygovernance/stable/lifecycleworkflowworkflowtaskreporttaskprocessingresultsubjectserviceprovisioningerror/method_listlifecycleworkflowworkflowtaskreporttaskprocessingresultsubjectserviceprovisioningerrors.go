package lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrors ...
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient) ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrors(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId) (result ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCustomPager{},
		Path:       fmt.Sprintf("%s/subject/serviceProvisioningErrors", id.ID()),
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
		Values *[]stable.ServiceProvisioningError `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient) ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId) (ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, ServiceProvisioningErrorOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient) ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId, predicate ServiceProvisioningErrorOperationPredicate) (result ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrors(ctx, id)
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

	result = ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
