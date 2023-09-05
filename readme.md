# Coding Test Simple Ecommerce

## Tech Stack
- Go Version 1.20
- Fiber
- Dependency Injection using Google Wire
- PostgreSQL

## Prerequisites
- Visual Studio Code (VS Code)
- Go for VS Code extension
- EditorConfig for VS Code extension
- Git version control
- Docker

## Local Setup
- Copy ``.env.example`` file to ``.env``, and adjust the contents (if you are using docker-compose please set env `DB_HOST` to value `db`)
- Adjust contents for ``docker-compose.yml`` and ``config/environment.yaml`` files (if necessary)

## Running
- Just run ``docker-compose up --build`` on your terminal.

## Debugging
- Go to **Run and Debug** menu on VS Code or press ``Ctrl+Shift+D``
- Click **Start Debugging** button or press ``F5``

## Postman
Click [Postman Online Documenter](https://documenter.getpostman.com/view/5040642/2s9YBxYbPm) **Or** you can import postman file from ``scripts`` folder

## Directory Structure
```shell
.vscode/            # VS Code files
cmd/                # Command or application entrypoints
config/             # Configuration files
deployment/         # Dockerfile
internal/
    controller/     # Handle request and response
    entity/         # Entity, value object, aggregates, etc.
    repository/     # Data access layer
    server/         # Bootstrap, initialization, wiring, etc.
    service/        # Business rules or logic layer
pkg/                # Library code that's ok to use by external apps
scripts/            # Scripts to perform various build, install, analysis, etc.
.editorconfig       # Standardized editors configuration
.gitignore          # Ignored files and dirs in Git
```
