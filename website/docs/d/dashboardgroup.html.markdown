---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_dashboard_group"
sidebar_current: "docs-logicmonitor-datasource-dashboard-group"
description: |-
  Use this datasource to get the ID of an available dashboard group.
---

# logicmonitor_dashboard_group

Use this datasource to get the ID of an available dashboard group.

## Example Usage

```hcl
# Look up a LogicMonitor dashboard group id
data "logicmonitor_dashboard_group" "TwentyFourLife" {
  filters {
    property = "fullPath"
    operator = ":"
    value = "Teams/Lakers/Guard"
  }
}
```

## Argument Reference

The following arguments are supported:

* `size` - (Optional) The number of results to display. Max is 1000. Default is 50
* `offset` - (Optional) The number of results to offset the displayed results by. Default is 0
* `filters` - (Optional) Filters the response according to the operator and value specified. Note that you can use * to match on more than one character. More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/dashboard-groups/

## Nested filters blocks

Nested `filters` blocks have the following structure: `property{operator}value`
* `property` - (Required if using filters) The name of filtered property. Currently the properties supported are `hostname` and `description`
* `operator` - (Required if using filters) The type of operator. Currently the operators supported are `:` `~` `!:` `!~`
* `value` - (Required if using filters) The value of the filtered property.
