# get-jwt

- [get-jwt](#get-jwt)
  - [Installing](#installing)
    - [Homebrew](#homebrew)
    - [Go](#go)
    - [Binary](#binary)
  - [Running](#running)
  - [Configuration](#configuration)
  - [Azure prerequisites](#azure-prerequisites)

A helper utility to make it easier to get a JSON Web Token (JWT) from a given service. Currently only Azure AD is supported.

If this is the first time that get-jwt is being run against an App Registration in Azure AD, there's a couple one-time steps to do (see [Azure prerequisites](#azure-prerequisites)).

## Installing

### Homebrew

### Go

### Binary

## Running

Starting with the base command of `get-jwt azure`, you'll need to provide three flags:
  - `--client-id`: The Application (client) ID
  - `--tenant-id`: The Directory (tenant) ID
  - `--scope`: The scope you are requesting access to

```
get-jwt azure --client-id c3ba59ce-1840-4824-b0b5-539d951c3b9c --tenant-id 76dd4f83-97f4-429d-8f93-b230bcf24989 --scope api://c3ba59ce-1840-4824-b0b5-539d951c3b9c/Read
```

## Configuration

| Flag          | Environment variable         |
| ------------- | ---------------------------- |
| `--client-id` | `GET_JWT_AZURE_AD_CLIENT_ID` |
| `--tenant-id` | `GET_JWT_AZURE_AD_TENANT_ID` |
| `--scope`     | `GET_JWT_AZURE_AD_SCOPE`     |
| `--copy`      | `GET_JWT_COPY_TO_CLIPBOARD`  |

## Azure prerequisites

There is some initial setup we need to do in App Registrations if this is the first time this tool is being run against it:

1. Add a **Mobile and desktop applications** platform, with `http://localhost` as a Redirect URI
  - Navigate to your App Registration in the Azure portal
  - Open the Authentication tab in the left sidebar
  - Click **Add a platform**
  - Choose **Mobile and desktop applications**
  - Under **Custom redirect URIs**, enter `http://localhost` to allow Azure to redirect back to the random port opened by get-jwt locally.
2. Enable **Allow public client flows**
  - Remain on the same Authentication tab
  - Under **Advanced settings** at the bottom, toggle the slider for **Allow public client flows** to `Yes`
