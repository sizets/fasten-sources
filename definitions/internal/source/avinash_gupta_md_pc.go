// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	platform "github.com/fastenhealth/fasten-sources/definitions/internal/platform"
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhir-myrecord.cerner.com/r4/eME9di7OkJfK0Dd-1ZlVnxRWlL5I7oU5/metadata
func GetSourceAvinashGuptaMdPc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/eME9di7OkJfK0Dd-1ZlVnxRWlL5I7oU5/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/eME9di7OkJfK0Dd-1ZlVnxRWlL5I7oU5/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/eME9di7OkJfK0Dd-1ZlVnxRWlL5I7oU5"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/eME9di7OkJfK0Dd-1ZlVnxRWlL5I7oU5"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAvinashGuptaMdPc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Avinash Gupta, MD, PC"
	sourceDef.SourceType = pkg.SourceTypeAvinashGuptaMdPc
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}