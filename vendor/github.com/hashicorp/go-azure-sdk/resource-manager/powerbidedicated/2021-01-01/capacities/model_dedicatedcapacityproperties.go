package capacities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DedicatedCapacityProperties struct {
	Administration    *DedicatedCapacityAdministrators `json:"administration"`
	FriendlyName      *string                          `json:"friendlyName,omitempty"`
	Mode              *Mode                            `json:"mode,omitempty"`
	ProvisioningState *CapacityProvisioningState       `json:"provisioningState,omitempty"`
	State             *State                           `json:"state,omitempty"`
	TenantId          *string                          `json:"tenantId,omitempty"`
}
