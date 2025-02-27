# pulumi/opentofu fork

This repository exists to expose internals of `opentofu/opentofu` for use in external projects.

## Usage

When linking to this repository, copy the `hcl/v2` replace into the consuming project `go.mod`, and
reference a particular commit of `pulumi/opentofu`:


``` go
replace (
	github.com/hashicorp/hcl/v2 => github.com/opentofu/hcl/v2 v2.0.0-20240814143621-8048794c5c52
	github.com/pulumi/opentofu => github.com/pulumi/opentofu v0.0.0-20250227012722-449d65590d90
)
```

You can then import functionality that is internal in `opentofu/opentofu` like this:

```go
import "github.com/pulumi/opentofu/configs"
```

WARNING: using internals implies stability risk as those APIs are liable to change and are not
guaranteed stable.

## Maintaining

### Branching

The latest code should always be merged into the `pulumi-main` branch. That is, to update, sync
`main` with `opentofu/opentofu`, then merge `main` into `pulumi-main`. Finally, update any sources
to make sure the project can be built. See Tools for helpers.

IMPORTANT: do not use git rebase! Use a merge workflow instead.

This permits consuming `go.mod` projects to reference particular commits from `pulumi-main` without
builds breaking.

### Tools

There is a script that assists with renaming go imports and exposing internals. To run:

``` shell
go run scripts/pulumi-imports/main.go
```

### Compatibility

There may be issues linking `github.com/pulumi/opentofu` and `github.com/opentofu/opentofu` into the
same `go.mod`. Replace directives do not seem to be sufficient to resolve.

There are issues with using `github.com/hashicorp/hcl/v2` alongside this library as
`github.com/pulumi/opentofu` insists on a particular replace.

The library is not currently compatible with `pulumi/pulumi-terraform-bridge`. The bridge currently
uses it own vendored version of opentofu codes and is likely to conflict.

Similarly to the problems in the bridge, there are issues with using another copy of
protobuf-generated Go code for `tfplugin5` and `tfplugin6` Terraform protocols. Protocol buffers
insist that only one copy of the generated code is included in-process, and linking
`github.com/pulumi/opentofu` brings one copy so you cannot bring another copy. At some point this
limitation may be lifted by re-packaging the protobuf-generated code in this repo to make it appear
distinct to protobuf tooling.
