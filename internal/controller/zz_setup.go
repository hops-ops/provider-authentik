// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/unbounded-tech/provider-authentik/internal/controller/applications/application"
	entitlement "github.com/unbounded-tech/provider-authentik/internal/controller/applications/entitlement"
	blueprint "github.com/unbounded-tech/provider-authentik/internal/controller/blueprints/blueprint"
	serviceconnection "github.com/unbounded-tech/provider-authentik/internal/controller/docker/serviceconnection"
	license "github.com/unbounded-tech/provider-authentik/internal/controller/enterprise/license"
	rule "github.com/unbounded-tech/provider-authentik/internal/controller/events/rule"
	transport "github.com/unbounded-tech/provider-authentik/internal/controller/events/transport"
	flow "github.com/unbounded-tech/provider-authentik/internal/controller/flows/flow"
	flowstagebinding "github.com/unbounded-tech/provider-authentik/internal/controller/flows/flowstagebinding"
	serviceconnectionk8s "github.com/unbounded-tech/provider-authentik/internal/controller/k8s/serviceconnection"
	outpost "github.com/unbounded-tech/provider-authentik/internal/controller/outposts/outpost"
	binding "github.com/unbounded-tech/provider-authentik/internal/controller/policies/binding"
	dummy "github.com/unbounded-tech/provider-authentik/internal/controller/policies/dummy"
	eventmatcher "github.com/unbounded-tech/provider-authentik/internal/controller/policies/eventmatcher"
	expiry "github.com/unbounded-tech/provider-authentik/internal/controller/policies/expiry"
	expression "github.com/unbounded-tech/provider-authentik/internal/controller/policies/expression"
	geoip "github.com/unbounded-tech/provider-authentik/internal/controller/policies/geoip"
	password "github.com/unbounded-tech/provider-authentik/internal/controller/policies/password"
	reputation "github.com/unbounded-tech/provider-authentik/internal/controller/policies/reputation"
	uniquepassword "github.com/unbounded-tech/provider-authentik/internal/controller/policies/uniquepassword"
	notification "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/notification"
	providergoogleworkspace "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/providergoogleworkspace"
	providermicrosoftentra "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/providermicrosoftentra"
	providerrac "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/providerrac"
	providerradius "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/providerradius"
	providersaml "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/providersaml"
	providerscim "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/providerscim"
	providerscope "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/providerscope"
	sourcekerberos "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/sourcekerberos"
	sourceldap "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/sourceldap"
	sourceoauth "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/sourceoauth"
	sourceplex "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/sourceplex"
	sourcesaml "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/sourcesaml"
	sourcescim "github.com/unbounded-tech/provider-authentik/internal/controller/propertymapping/sourcescim"
	providerconfig "github.com/unbounded-tech/provider-authentik/internal/controller/providerconfig"
	googleworkspace "github.com/unbounded-tech/provider-authentik/internal/controller/providers/googleworkspace"
	ldap "github.com/unbounded-tech/provider-authentik/internal/controller/providers/ldap"
	microsoftentra "github.com/unbounded-tech/provider-authentik/internal/controller/providers/microsoftentra"
	oauth2 "github.com/unbounded-tech/provider-authentik/internal/controller/providers/oauth2"
	proxy "github.com/unbounded-tech/provider-authentik/internal/controller/providers/proxy"
	rac "github.com/unbounded-tech/provider-authentik/internal/controller/providers/rac"
	radius "github.com/unbounded-tech/provider-authentik/internal/controller/providers/radius"
	saml "github.com/unbounded-tech/provider-authentik/internal/controller/providers/saml"
	scim "github.com/unbounded-tech/provider-authentik/internal/controller/providers/scim"
	ssf "github.com/unbounded-tech/provider-authentik/internal/controller/providers/ssf"
	endpoint "github.com/unbounded-tech/provider-authentik/internal/controller/rac/endpoint"
	initialpermissions "github.com/unbounded-tech/provider-authentik/internal/controller/rbac/initialpermissions"
	permissionrole "github.com/unbounded-tech/provider-authentik/internal/controller/rbac/permissionrole"
	permissionuser "github.com/unbounded-tech/provider-authentik/internal/controller/rbac/permissionuser"
	role "github.com/unbounded-tech/provider-authentik/internal/controller/rbac/role"
	kerberos "github.com/unbounded-tech/provider-authentik/internal/controller/sources/kerberos"
	ldapsources "github.com/unbounded-tech/provider-authentik/internal/controller/sources/ldap"
	oauth "github.com/unbounded-tech/provider-authentik/internal/controller/sources/oauth"
	plex "github.com/unbounded-tech/provider-authentik/internal/controller/sources/plex"
	samlsources "github.com/unbounded-tech/provider-authentik/internal/controller/sources/saml"
	authenticatorduo "github.com/unbounded-tech/provider-authentik/internal/controller/stages/authenticatorduo"
	authenticatoremail "github.com/unbounded-tech/provider-authentik/internal/controller/stages/authenticatoremail"
	authenticatorendpointgdtc "github.com/unbounded-tech/provider-authentik/internal/controller/stages/authenticatorendpointgdtc"
	authenticatorsms "github.com/unbounded-tech/provider-authentik/internal/controller/stages/authenticatorsms"
	authenticatorstatic "github.com/unbounded-tech/provider-authentik/internal/controller/stages/authenticatorstatic"
	authenticatortotp "github.com/unbounded-tech/provider-authentik/internal/controller/stages/authenticatortotp"
	authenticatorvalidate "github.com/unbounded-tech/provider-authentik/internal/controller/stages/authenticatorvalidate"
	authenticatorwebauthn "github.com/unbounded-tech/provider-authentik/internal/controller/stages/authenticatorwebauthn"
	captcha "github.com/unbounded-tech/provider-authentik/internal/controller/stages/captcha"
	consent "github.com/unbounded-tech/provider-authentik/internal/controller/stages/consent"
	deny "github.com/unbounded-tech/provider-authentik/internal/controller/stages/deny"
	dummystages "github.com/unbounded-tech/provider-authentik/internal/controller/stages/dummy"
	email "github.com/unbounded-tech/provider-authentik/internal/controller/stages/email"
	identification "github.com/unbounded-tech/provider-authentik/internal/controller/stages/identification"
	invitation "github.com/unbounded-tech/provider-authentik/internal/controller/stages/invitation"
	mutualtls "github.com/unbounded-tech/provider-authentik/internal/controller/stages/mutualtls"
	passwordstages "github.com/unbounded-tech/provider-authentik/internal/controller/stages/password"
	prompt "github.com/unbounded-tech/provider-authentik/internal/controller/stages/prompt"
	promptfield "github.com/unbounded-tech/provider-authentik/internal/controller/stages/promptfield"
	redirect "github.com/unbounded-tech/provider-authentik/internal/controller/stages/redirect"
	source "github.com/unbounded-tech/provider-authentik/internal/controller/stages/source"
	userdelete "github.com/unbounded-tech/provider-authentik/internal/controller/stages/userdelete"
	userlogin "github.com/unbounded-tech/provider-authentik/internal/controller/stages/userlogin"
	userlogout "github.com/unbounded-tech/provider-authentik/internal/controller/stages/userlogout"
	userwrite "github.com/unbounded-tech/provider-authentik/internal/controller/stages/userwrite"
	brand "github.com/unbounded-tech/provider-authentik/internal/controller/system/brand"
	certificatekeypair "github.com/unbounded-tech/provider-authentik/internal/controller/system/certificatekeypair"
	systemsettings "github.com/unbounded-tech/provider-authentik/internal/controller/system/systemsettings"
	group "github.com/unbounded-tech/provider-authentik/internal/controller/users/group"
	token "github.com/unbounded-tech/provider-authentik/internal/controller/users/token"
	user "github.com/unbounded-tech/provider-authentik/internal/controller/users/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		entitlement.Setup,
		blueprint.Setup,
		serviceconnection.Setup,
		license.Setup,
		rule.Setup,
		transport.Setup,
		flow.Setup,
		flowstagebinding.Setup,
		serviceconnectionk8s.Setup,
		outpost.Setup,
		binding.Setup,
		dummy.Setup,
		eventmatcher.Setup,
		expiry.Setup,
		expression.Setup,
		geoip.Setup,
		password.Setup,
		reputation.Setup,
		uniquepassword.Setup,
		notification.Setup,
		providergoogleworkspace.Setup,
		providermicrosoftentra.Setup,
		providerrac.Setup,
		providerradius.Setup,
		providersaml.Setup,
		providerscim.Setup,
		providerscope.Setup,
		sourcekerberos.Setup,
		sourceldap.Setup,
		sourceoauth.Setup,
		sourceplex.Setup,
		sourcesaml.Setup,
		sourcescim.Setup,
		providerconfig.Setup,
		googleworkspace.Setup,
		ldap.Setup,
		microsoftentra.Setup,
		oauth2.Setup,
		proxy.Setup,
		rac.Setup,
		radius.Setup,
		saml.Setup,
		scim.Setup,
		ssf.Setup,
		endpoint.Setup,
		initialpermissions.Setup,
		permissionrole.Setup,
		permissionuser.Setup,
		role.Setup,
		kerberos.Setup,
		ldapsources.Setup,
		oauth.Setup,
		plex.Setup,
		samlsources.Setup,
		authenticatorduo.Setup,
		authenticatoremail.Setup,
		authenticatorendpointgdtc.Setup,
		authenticatorsms.Setup,
		authenticatorstatic.Setup,
		authenticatortotp.Setup,
		authenticatorvalidate.Setup,
		authenticatorwebauthn.Setup,
		captcha.Setup,
		consent.Setup,
		deny.Setup,
		dummystages.Setup,
		email.Setup,
		identification.Setup,
		invitation.Setup,
		mutualtls.Setup,
		passwordstages.Setup,
		prompt.Setup,
		promptfield.Setup,
		redirect.Setup,
		source.Setup,
		userdelete.Setup,
		userlogin.Setup,
		userlogout.Setup,
		userwrite.Setup,
		brand.Setup,
		certificatekeypair.Setup,
		systemsettings.Setup,
		group.Setup,
		token.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
