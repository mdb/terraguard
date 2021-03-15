[![CI](https://github.com/mdb/terraguard/actions/workflows/ci.yml/badge.svg)](https://github.com/mdb/terraguard/actions/workflows/ci.yml)

# terraguard

`terraguard` helps automate [Terraform plan](https://www.terraform.io/docs/cli/commands/plan.html) reviews by checking a Terraform plan JSON for problematic resource changes.

## CLI Usage

```bash
terraguard \
  check \
    --resources="*foo*" \
    --plan="test_fixtures/basic_plan.json"
Error: test_fixtures/basic_plan.json indicates changes to guarded resources:

module.foo.null_resource.aliased
module.foo.null_resource.foo
null_resource.foo
```

## Disclaimer

Tools like [Open Policy Agent](https://www.openpolicyagent.org/) and [its Terraform capabilities](https://www.openpolicyagent.org/docs/latest/terraform/) arguably offer more robust, extendable, and fully featured means of enforcing Terraform policies. `terraguard` is a simple alternative, though is far less mature.
