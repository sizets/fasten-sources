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

// https://fhir-myrecord.cerner.com/r4/c63e0ebe-63f4-43e1-a630-b2adca03ff23/metadata
func GetSourceHospitalDist6OfHarperCounty(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/c63e0ebe-63f4-43e1-a630-b2adca03ff23/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/c63e0ebe-63f4-43e1-a630-b2adca03ff23/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/c63e0ebe-63f4-43e1-a630-b2adca03ff23"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/c63e0ebe-63f4-43e1-a630-b2adca03ff23"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeHospitalDist6OfHarperCounty]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Hospital Dist #6 of Harper County"
	sourceDef.SourceType = pkg.SourceTypeHospitalDist6OfHarperCounty
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}