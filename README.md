[![CI](https://github.com/mdb/terraguard/actions/workflows/ci.yml/badge.svg)](https://github.com/mdb/terraguard/actions/workflows/ci.yml)

# terraguard

`terraguard` helps automate [Terraform plan](https://www.terraform.io/docs/cli/commands/plan.html) reviews by checking a Terraform plan for problematic resource changes.

## Disclaimer

Tools like [Open Policy Agent](https://www.openpolicyagent.org/) and [its Terraform policy capabilities](https://www.openpolicyagent.org/docs/latest/terraform/) arguably offer more robust, extendable, and fully featured means of enforcing Terraform policies. `terraguard` is a simple alternative, though is far less mature.
