package newrelics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivateSaaSParameterRequest struct {
	PublisherId string `json:"publisherId"`
	SaasGuid    string `json:"saasGuid"`
}
