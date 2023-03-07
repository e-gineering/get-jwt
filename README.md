# get-jwt

- [get-jwt](#get-jwt)
  - [Installing](#installing)
    - [Homebrew](#homebrew)
    - [Go](#go)
    - [Binary](#binary)
  - [Running](#running)
  - [Configuration](#configuration)

A 

Currently only Azure AD is supported.

## Installing

### Homebrew

### Go

### Binary

## Running

Three required flags are:
  - `--client-id`: 
  - `--tenant-id`:
  - `--scope`: The scope you would like to have access to

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
