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

// https://fhir4.eclinicalworks.com/fhir/r4/JHJJCD/metadata
func GetSourceRinaHealthMedicalPracticePC(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Audience = "https://fhir4.eclinicalworks.com/fhir/r4/JHJJCD"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.eclinicalworks.com/fhir/r4/JHJJCD"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeRinaHealthMedicalPracticePC]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "RINA HEALTH MEDICAL PRACTICE P.C."
	sourceDef.SourceType = pkg.SourceTypeRinaHealthMedicalPracticePC
	sourceDef.Category = []string{"207RI0200X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1124590559"}}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}