package virtualmachineextensionimages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTypesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachineExtensionImage
}

// ListTypes ...
func (c VirtualMachineExtensionImagesClient) ListTypes(ctx context.Context, id PublisherId) (result ListTypesOperationResponse, err error) {
	req, err := c.preparerForListTypes(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineextensionimages.VirtualMachineExtensionImagesClient", "ListTypes", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineextensionimages.VirtualMachineExtensionImagesClient", "ListTypes", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListTypes(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineextensionimages.VirtualMachineExtensionImagesClient", "ListTypes", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListTypes prepares the ListTypes request.
func (c VirtualMachineExtensionImagesClient) preparerForListTypes(ctx context.Context, id PublisherId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/artifactTypes/vmExtension/types", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListTypes handles the response to the ListTypes request. The method always
// closes the http.Response Body.
func (c VirtualMachineExtensionImagesClient) responderForListTypes(resp *http.Response) (result ListTypesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
