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
get-jwt azure --client-id 00000000-1111-2222-3333-444444444444 --tenant-id 55555555-6666-7777-8888-999999999999 --scope api://00000000-1111-2222-3333-444444444444/Read
```

## Configuration

| Flag          | Environment variable         |
| ------------- | ---------------------------- |
| `--client-id` | `GET_JWT_AZURE_AD_CLIENT_ID` |
| `--tenant-id` | `GET_JWT_AZURE_AD_TENANT_ID` |
| `--scope`     | `GET_JWT_AZURE_AD_SCOPE`     |
| `--copy`      | `GET_JWT_COPY_TO_CLIPBOARD`  |
