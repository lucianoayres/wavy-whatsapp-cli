# WhatsApp CLI Client Refactoring Plan

## 1. Restructure Project Layout

- [ ] Create `/internal` directory for non-public packages
- [ ] Create `/pkg` directory for public packages
- [ ] Reorganize command handlers into `/cmd/wavy/commands`
- [ ] Move common utilities to `/internal/common`

## 2. Create Domain Layer

- [ ] Define core entities in `/internal/domain/models`
  - [ ] Contact model
  - [ ] Message model
  - [ ] Group model
- [ ] Define interfaces in `/internal/domain/ports`
  - [ ] WhatsAppClient interface
  - [ ] MessageRepository interface
  - [ ] ConfigRepository interface

## 3. Create Application Layer

- [ ] Implement use cases in `/internal/application/usecases`
  - [ ] SetupUseCase
  - [ ] SendMessageUseCase
  - [ ] CheckContactUseCase
  - [ ] ListGroupsUseCase
- [ ] Create service interfaces in `/internal/application/services`
  - [ ] MessageService
  - [ ] GroupService
  - [ ] AuthenticationService

## 4. Create Infrastructure Layer

- [ ] Implement WhatsApp client adapter in `/internal/infrastructure/whatsapp`
- [ ] Create file system adapter in `/internal/infrastructure/filesystem`
- [ ] Implement configuration manager in `/internal/infrastructure/config`
- [ ] Add database adapter in `/internal/infrastructure/database`

## 5. Refactor CLI Interface

- [ ] Create command factory in `/cmd/wavy/factory`
- [ ] Refactor command handlers to use dependency injection
- [ ] Extract presentation logic from command handlers
- [ ] Implement proper error handling middleware

## 6. Add Testing Infrastructure

- [ ] Create mock implementations of interfaces for testing
- [ ] Add unit tests for domain models
- [ ] Add unit tests for use cases
- [ ] Add integration tests for infrastructure
- [ ] Add end-to-end tests for commands

## 7. Improve Error Handling

- [ ] Create custom error types in `/internal/domain/errors`
- [ ] Implement consistent error handling strategy
- [ ] Add proper error logging
- [ ] Improve user-facing error messages

## 8. Add Configuration Management

- [ ] Implement proper configuration file handling
- [ ] Add environment variable support
- [ ] Create default configuration values
- [ ] Add configuration validation

## 9. Documentation

- [ ] Document interfaces and their implementations
- [ ] Add package documentation
- [ ] Update README.md with new project structure
- [ ] Add examples for extending the application

## 10. CI/CD Integration

- [ ] Set up GitHub Actions for automated testing
- [ ] Add linting and code quality checks
- [ ] Implement automated releases
