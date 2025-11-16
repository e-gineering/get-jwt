# get-jwt

[![License](https://img.shields.io/github/license/egineering-llc/get-jwt)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/egineering-llc/get-jwt)](go.mod)
[![Latest Release](https://img.shields.io/github/v/release/egineering-llc/get-jwt)](https://github.com/egineering-llc/get-jwt/releases)

A command-line helper utility to quickly obtain JSON Web Tokens (JWT) from identity providers. Currently supports Azure AD authentication with interactive browser-based login.

## Features

- **Interactive Authentication**: Opens your browser for secure OAuth authentication
- **Multiple Output Options**: Print to terminal or copy directly to clipboard
- **Environment Variable Support**: Configure via flags or environment variables
- **Cross-Platform**: Pre-built binaries for Linux, macOS, and Windows
- **Shell Completion**: Tab completion support for bash, zsh, PowerShell, and fish
- **Homebrew Support**: Easy installation on macOS and Linux

## Table of Contents

- [Installing](#installing)
  - [Homebrew](#homebrew)
  - [Pre-built Binaries](#pre-built-binaries)
  - [Go Install](#go-install)
  - [Command Completion](#command-completion)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Azure AD Setup](#azure-ad-setup)
- [Usage Examples](#usage-examples)
- [Contributing](#contributing)
  - [Building from Source](#building-from-source)
  - [Releasing](#releasing)
  - [Roadmap](#roadmap)
- [License](#license)

## Installing

### Homebrew

If you have Homebrew installed for macOS or Linux:

```bash
brew tap egineering-llc/get-jwt https://github.com/egineering-llc/get-jwt
brew install get-jwt
```

### Pre-built Binaries

Download pre-built binaries for Windows, macOS, and Linux from [GitHub Releases](https://github.com/egineering-llc/get-jwt/releases).

**Example for Linux:**

```bash
# Download the binary
curl -LO https://github.com/egineering-llc/get-jwt/releases/download/v0.1.2/get-jwt_0.1.2_linux_amd64

# Move to a directory in your PATH
sudo mv get-jwt_0.1.2_linux_amd64 /usr/local/bin/get-jwt

# Make it executable
chmod +x /usr/local/bin/get-jwt
```

**Example for macOS:**

```bash
# Download the binary (for Apple Silicon)
curl -LO https://github.com/egineering-llc/get-jwt/releases/download/v0.1.2/get-jwt_0.1.2_darwin_arm64

# Move and rename
sudo mv get-jwt_0.1.2_darwin_arm64 /usr/local/bin/get-jwt

# Make it executable
chmod +x /usr/local/bin/get-jwt
```

### Go Install

If you have Go 1.20 or later installed:

```bash
go install github.com/egineering-llc/get-jwt@latest
```

### Command Completion

Enable tab completion for your shell by adding the appropriate line to your shell configuration file:

**Bash** (`~/.bashrc` or `~/.bash_profile`):
```bash
source <(get-jwt completion bash)
```

**Zsh** (`~/.zshrc`):
```zsh
source <(get-jwt completion zsh)
```

**Fish** (`~/.config/fish/config.fish`):
```fish
get-jwt completion fish | source
```

**PowerShell** (add to your PowerShell profile):
```powershell
get-jwt completion powershell | Out-String | Invoke-Expression
```

After updating your configuration, restart your terminal or source the configuration file.

## Quick Start

The basic command requires three pieces of information from your Azure AD App Registration:

```bash
get-jwt azure \
  --client-id <your-client-id> \
  --tenant-id <your-tenant-id> \
  --scope <your-scope>
```

**Example:**

```bash
get-jwt azure \
  --client-id c3ba59ce-1840-4824-b0b5-539d951c3b9c \
  --tenant-id 76dd4f83-97f4-429d-8f93-b230bcf24989 \
  --scope api://c3ba59ce-1840-4824-b0b5-539d951c3b9c/Read
```

This will:
1. Open your default browser for Azure AD authentication
2. Print the JWT to your terminal after successful login

## Configuration

Configuration can be provided via command-line flags or environment variables. Flags take precedence over environment variables.

| Flag          | Environment Variable          | Required | Default | Description                           |
|---------------|-------------------------------|----------|---------|---------------------------------------|
| `--client-id` | `GET_JWT_AZURE_AD_CLIENT_ID`  | Yes      | -       | Azure AD Application (client) ID      |
| `--tenant-id` | `GET_JWT_AZURE_AD_TENANT_ID`  | Yes      | -       | Azure AD Directory (tenant) ID        |
| `--scope`     | `GET_JWT_AZURE_AD_SCOPE`      | Yes      | -       | OAuth scope to request                |
| `--copy`      | `GET_JWT_COPY_TO_CLIPBOARD`   | No       | `false` | Copy JWT to clipboard instead of printing |

### Using Environment Variables

Set environment variables to avoid typing flags repeatedly:

```bash
export GET_JWT_AZURE_AD_CLIENT_ID="c3ba59ce-1840-4824-b0b5-539d951c3b9c"
export GET_JWT_AZURE_AD_TENANT_ID="76dd4f83-97f4-429d-8f93-b230bcf24989"
export GET_JWT_AZURE_AD_SCOPE="api://c3ba59ce-1840-4824-b0b5-539d951c3b9c/Read"

# Now you can simply run:
get-jwt azure
```

## Azure AD Setup

Before using this tool with an Azure AD App Registration for the first time, you need to configure the app to allow public client flows.

### Prerequisites

1. **Add Mobile and Desktop Applications Platform**
   - Navigate to your App Registration in the [Azure Portal](https://portal.azure.com)
   - Go to **Authentication** in the left sidebar
   - Click **Add a platform**
   - Select **Mobile and desktop applications**
   - Under **Custom redirect URIs**, enter: `http://localhost`
   - Click **Configure**

2. **Enable Public Client Flows**
   - Stay on the **Authentication** page
   - Scroll down to **Advanced settings**
   - Toggle **Allow public client flows** to **Yes**
   - Click **Save**

### Finding Your IDs

- **Client ID**: Found on the app registration's **Overview** page as "Application (client) ID"
- **Tenant ID**: Found on the app registration's **Overview** page as "Directory (tenant) ID"
- **Scope**: Typically in the format `api://<client-id>/<permission-name>` or use default scopes like `https://graph.microsoft.com/.default`

## Usage Examples

### Copy JWT to Clipboard

```bash
get-jwt azure \
  --client-id c3ba59ce-1840-4824-b0b5-539d951c3b9c \
  --tenant-id 76dd4f83-97f4-429d-8f93-b230bcf24989 \
  --scope api://c3ba59ce-1840-4824-b0b5-539d951c3b9c/Read \
  --copy
```

### Using Microsoft Graph API Scope

```bash
get-jwt azure \
  --client-id c3ba59ce-1840-4824-b0b5-539d951c3b9c \
  --tenant-id 76dd4f83-97f4-429d-8f93-b230bcf24989 \
  --scope https://graph.microsoft.com/.default
```

### Store in a Variable (for scripting)

```bash
JWT_TOKEN=$(get-jwt azure \
  --client-id c3ba59ce-1840-4824-b0b5-539d951c3b9c \
  --tenant-id 76dd4f83-97f4-429d-8f93-b230bcf24989 \
  --scope api://c3ba59ce-1840-4824-b0b5-539d951c3b9c/Read)

# Use the token in an API request
curl -H "Authorization: Bearer $JWT_TOKEN" https://api.example.com/resource
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Building from Source

**Clone the repository:**

```bash
git clone https://github.com/egineering-llc/get-jwt.git
cd get-jwt
```

**Build the binary:**

```bash
go build .
```

This creates a `get-jwt` binary in the current directory.

**Run without building:**

```bash
go run main.go azure --client-id <client-id> --tenant-id <tenant-id> --scope <scope>
```

**Run tests:**

```bash
go test ./...
```

### Releasing

Releases are automated using [GoReleaser](https://goreleaser.com/) via GitHub Actions.

To create a new release:

1. Tag the commit with a version number:
   ```bash
   git tag v0.2.1
   ```

2. Push the tag to GitHub:
   ```bash
   git push origin v0.2.1
   ```

3. GitHub Actions will automatically build and publish the release

### Roadmap

- [ ] Add a `version` subcommand to display the current version
- [ ] Handle common MSAL errors more gracefully with user-friendly messages
- [ ] Add a `--quiet` flag to output only the JWT for easier use in scripts
- [ ] Integrate [Viper](https://github.com/spf13/viper) for better configuration management
- [ ] Use `cmd.MarkFlagRequired()` for required flags
- [ ] Implement MSAL token caching for improved performance
- [ ] Add support for additional identity providers (Okta, Auth0, etc.)
- [ ] Add support for certificate-based authentication

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
