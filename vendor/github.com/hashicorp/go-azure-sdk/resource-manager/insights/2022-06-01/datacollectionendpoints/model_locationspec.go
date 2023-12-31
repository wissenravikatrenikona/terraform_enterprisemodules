package datacollectionendpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationSpec struct {
	Location           *string                              `json:"location,omitempty"`
	ProvisioningStatus *KnownLocationSpecProvisioningStatus `json:"provisioningStatus,omitempty"`
}
