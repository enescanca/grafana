---
description: Feature and improvement highlights for Grafana v10.1
keywords:
  - grafana
  - new
  - documentation
  - '10.1'
  - release notes
title: What's new in Grafana v10.1
weight: -38
---

# What’s new in Grafana v10.1

Welcome to Grafana 10.1! Read on to learn about changes to search and navigation, dashboards and visualizations, and security and authentication.

For even more detail about all the changes in this release, refer to the [changelog](https://github.com/grafana/grafana/blob/master/CHANGELOG.md). For the specific steps we recommend when you upgrade to v10.1, check out our [Upgrade Guide]({{< relref "../upgrade-guide/upgrade-v10.1/index.md" >}}).

<!-- Template below
## Feature
<!-- Name of contributor -->
<!-- [Generally available | Available in private/public preview | Experimental] in Grafana [Open Source, Enterprise, Cloud Free, Cloud Pro, Cloud Advanced]
Description. Include an overview of the feature and problem it solves, and where to learn more (like a link to the docs).
{{% admonition type="note" %}}
You must use relative references when linking to docs within the Grafana repo. Please do not use absolute URLs. For more information about relrefs, refer to [Links and references](/docs/writers-toolkit/writing-guide/references/).
{{% /admonition %}}
-->

## Dashboards and visualizations

### Disconnect values in time series, trend, and state timeline visualizations

_Generally available in all editions of Grafana._

<!-- Nathan Marrs -->

You can now choose whether to set a threshold above which values in the data should be disconnected. This can be useful in cases where you have sensors that report a value at a set interval, but you want to disconnect the values when the sensor does not respond. This feature complements the existing [connect null values functionality]({{< relref "../panels-visualizations/visualizations/time-series/#connect-null-values" >}}).

To learn more, refer to our [disconnect values documentation]({{< relref "../panels-visualizations/visualizations/time-series/#disconnect-values" >}}).

{{< figure src="/media/docs/grafana/screenshot-grafana-10-1-disconnect-values-examples.png" max-width="750px" caption="Disconnect values in time series, trend, and state timeline visualizations" >}}

## Authentication and authorization

### Configure refresh token handling separately for OAuth providers

<!-- Mihaly Gyongyosi -->

_Available in public preview in all editions of Grafana._

With Grafana v9.3, we introduced a feature toggle called `accessTokenExpirationCheck`. It improves the security of Grafana by checking the expiration of the access token and automatically refreshing the expired access token when the user is logged in using one of the OAuth providers.

With the current release, we introduce a new configuration option for each OAuth provider called `use_refresh_token` that allows you to configure whether the particular OAuth integration should use refresh tokens to automatically refresh access tokens when they expire. In addition, to further improve security and provide secure defaults, `use_refresh_token` is enabled by default for providers that support either refreshing tokens automatically or client-controlled fetching of refresh tokens. It's enabled by default for the following OAuth providers: `AzureAD`, `GitLab`, `Google`.

For more information on how to set up refresh token handling, please refer to [the documentation of the particular OAuth provider.]({{< relref "../setup-grafana/configure-security/configure-authentication/" >}}).

{{% admonition type="note" %}}
The `use_refresh_token` configuration must be used in conjunction with the `accessTokenExpirationCheck` feature toggle. If you disable the `accessTokenExpirationCheck` feature toggle, Grafana will not check the expiration of the access token and will not automatically refresh the expired access token, even if the `use_refresh_token` configuration is set to `true`.

The `accessTokenExpirationCheck` feature toggle will be removed in Grafana v10.2.
{{% /admonition %}}

### OAuth role mapping enforcement

<!-- AuthNZ -->

This change impacts `GitHub` OAuth, `Gitlab` OAuth, `Okta` OAuth and `Generic` OAuth.

Currently if no organization role mapping is found for a user when connecting via OAuth, Grafana doesn’t update the user’s organization role.

With Grafana 10.1, on every login, if the role_attribute_path property does not return a role, then the user is assigned the role specified by the auto_assign_org_role option or the default role for the organization, by default, Viewer.

To avoid overriding manually set roles, enable the skip_org_role_sync option in the Grafana configuration for your OAuth provider before the user logs in for the first time.

### GitLab OIDC support

<!-- AuthNZ -->

Grafana 10.1 now supports GitLab OIDC via the `GitLab` OAuth provider in addition to the existing `GitLab` OAuth2 provider. This allows you to use GitLab OIDC to authenticate users in Grafana.

This allows Grafana to reduce the access scope to only the required scopes for authentication and authorization instead
of full read API access.

To learn how to migrate your GitLab OAuth2 setup to OIDC, refer to our [GitLab authentication documentation]({{< relref "../setup-grafana/configure-security/configure-authentication/gitlab/" >}}).

### Google OIDC and teamsync support

<!-- AuthNZ -->

Grafana 10.1 now supports Google OIDC via the `Google` OAuth provider in addition to the existing `Google` OAuth2 provider. This allows you to use Google OIDC to authenticate users in Grafana.

This allows Grafana to reduce the access scope to only the required scopes for authentication and authorization.

This release also adds support for Google OIDC in Team Sync. You can now easily add users to teams by utilizing their Google groups.

To learn how to migrate your Google OAuth2 setup to OIDC and how to setup teamsync, refer to our [Google authentication documentation]({{< relref "../setup-grafana/configure-security/configure-authentication/google/" >}}).
