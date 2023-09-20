# Versioned migrations with Atlasgo.io and GORM in Golang

## Resources

- https://atlasgo.io/guides/orms/gorm
- https://gorm.io
- https://github.com/ariga/atlas-provider-gorm
- https://amal.sh/posts/gorm-atlas/

## Prerequisites
1. Setup Atlas & atlas-provider-gorm as documented in the [atlasgo guide](https://atlasgo.io/guides/orms/gorm)
2. This PoC assumes the "Standalone mode" described in the guides.

## How to do versioned mihgrations
1. Start by modifying the gorm models in `models/models.go`
2. Run `atlas migrate diff --env gorm` to generate a new migration file.
3. The new migration file is located in `migrations/` and referenced in the atlas.sum file.
4. To apply the migration, the filename needs to be added to the `DatabaseVersions` array in `main.go` (Example: `{MigrationFileName: "migrations/20230920174735.sql", Version: 20230920174735}`).
5. Run `go run .` to apply the migrations.
6. Perform steps 1 to 4 multiple times befor running `go run .` to do multiple migrations at once.

## How would this be used in production?
After updating the models, the dev would run the `atlas migrate diff --env gorm` command to generate the migration file and follow through to step 4.
This could also be implemented into the Azure Pipline. The generated migration files are checked into git and deployed. By running the updated program, the migrations are applied automatically.

## Special Files&Folders
1. `migrations/` contains the generated migration files.
2. `models/` has all database models, so they are detectable by the `atlas-provider-gorm`.
3. `atlas.db` is used by atlas to hold it's state.
4. `data.db` is the example database for the migrations.
5. `README.md` is this file...