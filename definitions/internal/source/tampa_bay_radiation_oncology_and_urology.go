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

// https://fhir-myrecord.cerner.com/r4/_equjXi-I6k5_iN9ulQyh3xNMVDgQ7-w/metadata
func GetSourceTampaBayRadiationOncologyAndUrology(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/_equjXi-I6k5_iN9ulQyh3xNMVDgQ7-w/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/_equjXi-I6k5_iN9ulQyh3xNMVDgQ7-w/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/_equjXi-I6k5_iN9ulQyh3xNMVDgQ7-w"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/_equjXi-I6k5_iN9ulQyh3xNMVDgQ7-w"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTampaBayRadiationOncologyAndUrology]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Tampa Bay Radiation Oncology & Urology"
	sourceDef.SourceType = pkg.SourceTypeTampaBayRadiationOncologyAndUrology
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}