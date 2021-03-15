[![CI](https://github.com/mdb/terraguard/actions/workflows/ci.yml/badge.svg)](https://github.com/mdb/terraguard/actions/workflows/ci.yml)

# terraguard

`terraguard` helps automate [Terraform plan](https://www.terraform.io/docs/cli/commands/plan.html) reviews by checking a Terraform plan JSON for problematic resource changes.

## CLI Usage

`terraguard check` examines a Terraform plan JSON file for changes to guarded resources.

```text
terraguard check --help
Check if a Terraform plan seeks to modify the specified resources

Usage:
  terraguard check [flags]

Flags:
  -g, --guard strings   A comma-separated list of guarded resource addresses
  -h, --help            help for check
  -p, --plan string     The path to a Terraform plan output JSON file
```

Basic example:

```bash
terraguard \
  check \
    --guard="*foo*" \
    --plan="test_fixtures/basic_plan.json"
Error: test_fixtures/basic_plan.json indicates changes to guarded resources:

module.foo.null_resource.aliased
module.foo.null_resource.foo
null_resource.foo
```

With multiple guarded resources:

```bash
terraguard \
  check \
    --guard="*foo*" \
    --guard="*biz*" \
    --plan="test_fixtures/basic_plan.json"
Error: test_fixtures/basic_plan.json indicates changes to guarded resources:

module.foo.null_resource.aliased
module.foo.null_resource.foo
null_resource.bar
null_resource.foo
```

## Disclaimer

Tools like [Open Policy Agent](https://www.openpolicyagent.org/) and [its Terraform capabilities](https://www.openpolicyagent.org/docs/latest/terraform/) arguably offer more robust, extendable, and fully featured means of enforcing Terraform policies. `terraguard` is a simple alternative, though is far less mature.
