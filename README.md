# github.com/octokit/go-sdk

An "alpha" version of a generated Go SDK from [GitHub's OpenAPI spec](https://github.com/github/rest-api-description), built on [Kiota](https://github.com/microsoft/kiota).

## How do I use it?

See example client instantiations and requests in [example_test.go](pkg/example_test.go) or in the [cmd/ directory](cmd/).

⚠️ **Note**: This SDK is not yet stable. Breaking changes may occur at any time.

### Authentication

This SDK supports [Personal Access Tokens (classic)](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#personal-access-tokens-classic), [fine-grained Personal Access Tokens](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#fine-grained-personal-access-tokens), and [GitHub Apps](https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/about-authentication-with-a-github-app) authentication.

In order to use either type of Personal Access token, you can use the `WithTokenAuthentication("YOUR_TOKEN_HERE")` functional option when constructing a client, like so:

```go
client, err := pkg.NewApiClient(
	pkg.WithTokenAuthentication(os.Getenv("GITHUB_TOKEN")),
)
if err != nil {
	log.Fatalf("error creating client: %v", err)
}
```

In order to authenticate as a GitHub App, you can use the `WithGitHubAppAuthentication` functional option:

```go
client, err := pkg.NewApiClient(
	pkg.WithGitHubAppAuthentication("/path/to/your/pem/file.pem", "your-client-ID", yourInstallationIDInt),
)
if err != nil {
	log.Fatalf("error creating client: %v", err)
}
```

To see more detailed examples, view [the cmd/ directory in this repo](cmd/).

⚠️ **Note**: There are [three types](https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/about-authentication-with-a-github-app) of GitHub App authentication:
1. [As the App itself (meta endpoints)](https://docs.github.com/en/rest/apps/apps?apiVersion=2022-11-28)
1. [As an App installation](https://docs.github.com/en/rest/authentication/endpoints-available-for-github-app-installation-access-tokens?apiVersion=2022-11-28)
1. On behalf of a user

Authenticating on behalf of a user is not supported in an SDK, as it requires a UI authentication flow with redirects. This SDK supports authenticating as the App itself and as an App installation.

Note that the SDK **does not yet** support authenticating as the App itself and as an App installation using the same client transparently to the user. Authenticating as the App itself requires [creating a JSON Web Token (JWT)](https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-json-web-token-jwt-for-a-github-app) and using that as token authentication. For helpers to create and sign a JWT in Go, you may use the [golang-jwt/jwt](https://github.com/golang-jwt/jwt) library.

Authenticating as an App installation can be done using the `WithGitHubAppAuthentication` functional option. Future work is planned to make the App meta endpoints vs. App installation endpoints auth schemes transparent to the user and only require one client setup.

## Why a generated SDK?

We want to...
1.  provide 100% coverage of the API in our SDK
2.  use this as a building block for future SDK tooling.

## Why Go?

We don't currently support a Go SDK and we wanted a narrow scope for the initial effort without worrying about cutting over users of an existing SDK.

## How can I report on my experience or issues with the SDK?

Please use this project's [issues](https://github.com/octokit/go-sdk/issues)!

## Releasing this project

Periodically (based on the frequency of [this workflow](https://github.com/octokit/source-generator/blob/main/.github/workflows/build-go.yml)), the source-generator repository will ingest the latest version of [GitHub's OpenAPI spec](https://github.com/github/rest-api-description) and generate a new version of this SDK. If there is a diff, a PR (similar to [this one](https://github.com/octokit/go-sdk/pull/22)) will be generated.

When reviewing the PR, analyze the diff and determine whether the changes are breaking (for which a major version number must be incremented), feature additions (for which a minor version number must be incremented), or bug fixes or docs changes (for which a patch number must be incremented). For more details about how to select an appropriate semantic version, see [semver.org](https://semver.org/). In many/most cases, due to the scale of GitHub's specification and the rate of change on it, the diff will be large and the changes will be technically breaking. This will mean incrementing a major version number. If a major version is being incremented, it must be changed in the [go.mod](./go.mod) file as described in [these docs](https://go.dev/doc/modules/release-workflow#breaking).

When changes are analyzed, change the PR title appropriately (see [this PR](https://github.com/octokit/go-sdk/pull/40) for an example) and merge it. Then go to [repository releases](https://github.com/octokit/go-sdk/releases), tag the release with the chosen version, title it with the chosen version, use the "Generate release notes" button to see what PRs will be included in the release, and manually edit the release notes grouping the changes under the headings `Features`, `Fixes`, `Maintenance`, and `Documentation` when appropriate. After clicking "Publish Release", the new version will be available for use!
