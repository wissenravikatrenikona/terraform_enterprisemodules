package registrationdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistrationDefinition struct {
	Id         *string                           `json:"id,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Plan       *Plan                             `json:"plan"`
	Properties *RegistrationDefinitionProperties `json:"properties"`
	Type       *string                           `json:"type,omitempty"`
}
