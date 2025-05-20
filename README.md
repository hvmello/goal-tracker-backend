# Goal Tracker 

## Overview

This is a Goal Tracker made in Go. This project is meant to be used as a way to monitore all the goals, checking if it is done, doing or to-do.
The front-end is located in another repository called goal-tracker-front. 

## Project Structure

```plaintext
├── cmd/                 # Main applications for this project
│   └── appname/         # Application entry point
│       └── main.go      # Main package
├── pkg/                 # Library code that's OK to use by external applications
├── internal/            # Private application and library code
├── api/                 # API definitions and protocol files (e.g., protobuf)
├── configs/             # Configuration files and templates
├── scripts/             # Helper scripts for build, install, analysis, etc.
├── tests/               # Additional external test apps and test data
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── README.md            # Project documentation



### Branch naming pattern with semantic versioning

| Type            | Prefix + version                | Example                               |
| --------------- | ------------------------------ | ------------------------------------|
| Feature         | `feature/1.0.0/`               | `feature/1.0.0/create-goals-api`     |
| Bug fix         | `fix/1.0.1/`                   | `fix/1.0.1/database-connection`      |
| Infrastructure  | `chore/1.1.0/`                 | `chore/1.1.0/docker-setup`            |
| Refactoring     | `refactor/1.2.0/`              | `refactor/1.2.0/handler-cleanup`     |
| Urgent hotfix   | `hotfix/1.0.2/`                | `hotfix/1.0.2/fix-prod-crash`        |

---

### Tips for using this pattern:

- Version `1.0.0` follows [Semantic Versioning (SemVer)](https://semver.org/), meaning:  
  - **MAJOR** (1): incompatible API changes  
  - **MINOR** (0): backward-compatible feature additions  
  - **PATCH** (0): backward-compatible bug fixes

- Update the version in the branch name as development progresses:  
  - Increment **MINOR** for new features without breaking changes  
  - Increment **PATCH** for bug fixes or small improvements  
  - Increment **MAJOR** for breaking changes

- Including versions in branch names helps with automation, tagging, and release management.

---

Want help setting up automated Git tags and release workflows based on this versioning?
