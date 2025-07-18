# Release Process for Wavy WhatsApp CLI

This document outlines the process for creating new releases of Wavy WhatsApp CLI.

## Version Tagging Strategy

We follow [Semantic Versioning](https://semver.org/) (SemVer) for version numbers:

- **MAJOR** version for incompatible API changes (e.g., v1.0.0 to v2.0.0)
- **MINOR** version for new functionality in a backward-compatible manner (e.g., v1.0.0 to v1.1.0)
- **PATCH** version for backward-compatible bug fixes (e.g., v1.0.0 to v1.0.1)

## Creating a New Release

1. Ensure the codebase is ready for release:

   - All features for this version are complete
   - Tests are passing
   - Documentation is up-to-date

2. Create and push a new tag:

   ```bash
   # Replace X.Y.Z with the appropriate version number
   git tag -a vX.Y.Z -m "Release version X.Y.Z"
   git push origin vX.Y.Z
   ```

3. The GitHub Actions workflow will automatically:

   - Build the binaries for multiple platforms (Linux, macOS, Windows)
   - Create a GitHub Release with those binaries
   - Generate release notes based on commits since the last release

4. Once the workflow completes, visit the [Releases page](https://github.com/YOUR_USERNAME/wavy-whatsapp-cli/releases) to verify the new release.

## Release Artifacts

Each release includes the following binaries:

- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

## Pre-releases

For pre-releases, use suffixes like `-alpha`, `-beta`, or `-rc` in your tag:

```bash
git tag -a vX.Y.Z-alpha -m "Alpha release X.Y.Z"
git push origin vX.Y.Z-alpha
```

## Release Checklist

Before tagging a release, ensure:

- [ ] All features planned for this release are complete
- [ ] Tests are passing
- [ ] Documentation is updated
- [ ] CHANGELOG.md is updated (if applicable)
- [ ] README.md has the latest information
