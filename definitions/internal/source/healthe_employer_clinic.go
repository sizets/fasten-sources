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

// https://fhir-myrecord.cerner.com/r4/0635029c-ad2b-4c76-90de-10b216b6fefa/metadata
func GetSourceHealtheEmployerClinic(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/0635029c-ad2b-4c76-90de-10b216b6fefa/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/0635029c-ad2b-4c76-90de-10b216b6fefa/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/0635029c-ad2b-4c76-90de-10b216b6fefa"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/0635029c-ad2b-4c76-90de-10b216b6fefa"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeHealtheEmployerClinic]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Healthe Employer Clinic"
	sourceDef.SourceType = pkg.SourceTypeHealtheEmployerClinic
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}