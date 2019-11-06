# wait_for defines the ability to wait for a given condition to happen in a
# juju status output. The output is JSON, so everything that the API server
# knows about should be valid.
# The query argument is a jq query.
#
# ```
# wait_for <model name> <query>
# ```
wait_for() {
    local name query

    name=${1}
    query=${2}

    attempt=0
    # shellcheck disable=SC2046,SC2143
    until [ "$(juju status --format=json 2> /dev/null | jq -S "${query}" | grep "${name}")" ]; do
        echo "[+] (attempt ${attempt}) polling status"
        juju status --relations 2>&1 | sed 's/^/    | /g'
        sleep 5
        let attempt=attempt+1
    done
}

idle_condition() {
    local name app_index unit_index

    name=${1}
    app_index=${2:-0}
    unit_index=${3:-0}

    echo ".applications | select(.[\"$name\"] | .units | .[\"$name/$unit_index\"] | .[\"juju-status\"] | .current == \"idle\") | keys[$app_index]"
}

# workload_status gets the workload-status object for the unit - use
# .current or .message to select the actual field you need.
workload_status() {
    local app unit

    app=$1
    unit=$2

    echo ".applications[\"$app\"].units[\"$app/$unit\"][\"workload-status\"]"
}

# agent_status gets the juju-status object for the unit - use
# .current or .message to select the actual field you need.
agent_status() {
    local app unit

    app=$1
    unit=$2

    echo ".applications[\"$app\"].units[\"$app/$unit\"][\"juju-status\"]"
}

# wait_for_machine_agent_status blocks until the machine agent for the specified
# machine instance ID reports the requested status.
#
# ```
# wait_for_machine_agent_status <instance-id> <status>
#
# example:
# wait_for_machine_agent_status "i-1234" "started"
# ```
wait_for_machine_agent_status() {
    local inst_id status

    inst_id=${1}
    status=${2}

    attempt=0
    # shellcheck disable=SC2046,SC2143
    until [ $(juju show-machine --format json | jq -r ".[\"machines\"] | .[\"${inst_id}\"] | .[\"juju-status\"] | .[\"current\"]" | grep "${status}") ]; do
        echo "[+] (attempt ${attempt}) polling status"
        juju machines | grep "$inst_id" 2>&1 | sed 's/^/    | /g'
        sleep 5
        let attempt=attempt+1
    done
}

# wait_for_machine_netif_count blocks until the number of detected network
# interfaces for the requested machine instance ID becomes equal to the desired
# value.
#
# ```
# wait_for_machine_netif_count <instance-id> <count>
#
# example:
# wait_for_machine_netif_count "i-1234" "42"
# ```
wait_for_machine_netif_count() {
    local inst_id count

    inst_id=${1}
    count=${2}

    attempt=0
    # shellcheck disable=SC2046,SC2143
    until [ $(juju show-machine --format json | jq -r ".[\"machines\"] | .[\"${inst_id}\"] | .[\"network-interfaces\"] | length" | grep "${count}") ]; do
        # shellcheck disable=SC2046,SC2143
        echo "[+] (attempt ${attempt}) network interface count for instance ${inst_id} = "$(juju show-machine --format json | jq -r ".[\"machines\"] | .[\"${inst_id}\"] | .[\"network-interfaces\"] | length")
        sleep 5
        let attempt=attempt+1
    done
}