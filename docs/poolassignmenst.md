# Poolassignments

Leostream poolassignments are used to link pools to users and groups. A user is given a policy directly or through group membership at login and the policy defines settings and contains poolassignments. When a user logs he or she receives a list of pools that they are assigned to. 


## Poolassignment command
```bash
Manage pool assignments, the poolassignment command can be used to list, get, delete or create and update pool assignments.

Usage:
  leostream-admin-cli poolassignment [flags]
  leostream-admin-cli poolassignment [command]

Available Commands:
  create      Create Leostream pool assignment
  delete      Delete Leostream pool assignment
  get         Get Leostream pool assignment
  list        List Leostream pool assignments
  update      Update Leostream pool assignment

Flags:
  -h, --help                       help for poolassignment
      --policy_id string           PolicyID to get pool assignments for
      --poolassignment_id string   PoolAssignmentID to get pool assignment for

Use "leostream-admin-cli poolassignment [command] --help" for more information about a command.
```

### Poolassignment list command

Displays a list of poolassignments.
Required flags: --policy_id

```bash
% leostream-admin-cli poolassignment list --help
List Leostream pool assignments by ID. For example:
                leostream-admin-cli poolassignment list --policy_id 1
                leostream-admin-cli poolassignment list --policy_id 1 --poolassignment_id 1

Usage:
  leostream-admin-cli poolassignment list [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --policy_id string           PolicyID to get pool assignments for
      --poolassignment_id string   PoolAssignmentID to get pool assignment for
```

### Poolassignment get command

Display pool assignment details
Required flags: --policy_id

```bash
% leostream-admin-cli poolassignment get --help
Get Leostream pool assignment by ID. For example:
                leostream-admin-cli poolassignment get --policy_id 1 --poolassignment_id 1

Usage:
  leostream-admin-cli poolassignment get [flags]

Flags:
  -h, --help   help for get

Global Flags:
      --policy_id string           PolicyID to get pool assignments for
      --poolassignment_id string   PoolAssignmentID to get pool assignment for
```

### Poolassignment create command

Create a poolassignment.
Required flags: --policy_id, --pool_id

```bash
% leostream-admin-cli poolassignment create --help
Create Leostream pool assignment by ID. For example:
                leostream-admin-cli poolassignment create --policy_id 1 --pool_id 1 --protocolplan_id 1 --powerplan_id 1 --releaseplan_id 1 --display_mode 9 --offer_filter 0 --offer_quantity 1 --start_if_stopped 0 --isMemberOf "AD Group" --isMemberOf "AD Group"

Usage:
  leostream-admin-cli poolassignment create [flags]

Flags:
      --display_mode string       Way desktop and pool are displayed to the user (default "9")
  -h, --help                      help for create
      --isMemberOf stringArray    SAML attribute or AD group to add to the pool assignment(can be used multiple times)
      --offer_filter string       Filter how desktops are offered to the user (default "0")
      --offer_quantity string     Number of desktops are offered to the user (default "1")
      --pool_id string            Pool to associate with  the pool assignment
      --powerplan_id string       Power plan to associate with  the pool assignment
      --protocolplan_id string    Protocol plan to associate with the pool assignment
      --releaseplan_id string     AD group to associate with the pool assignment
      --start_if_stopped string   Start stopped or suspended desktops (default "0")

Global Flags:
      --policy_id string           PolicyID to get pool assignments for
      --poolassignment_id string   PoolAssignmentID to get pool assignment for
```

### Poolassignment update command

Update a poolassignment.
Required flags: --policy_id, --poolassignment_id

```bash
% leostream-admin-cli poolassignment update --help
Update Leostream pool assignment by ID. For example:
                leostream-admin-cli poolassignment update --policy_id 1 --poolassignment_id 1 --pool_id 1 --protocolplan_id 1 --powerplan_id 1 --releaseplan_id 1 --display_mode 9 --offer_filter 0 --offer_quantity 1 --start_if_stopped 0

Usage:
  leostream-admin-cli poolassignment update [flags]

Flags:
      --display_mode string       Way desktop and pool are displayed to the user (default "9")
  -h, --help                      help for update
      --isMemberOf stringArray    SAML attribute or AD group to add to the pool assignment(can be used multiple times)
      --offer_filter string       Filter how desktops are offered to the user (default "0")
      --offer_quantity string     Number of desktops are offered to the user (default "1")
      --pool_id string            Pool to associate with  the pool assignment
      --powerplan_id string       Power plan to associate with  the pool assignment
      --protocolplan_id string    Protocol plan to associate with the pool assignment
      --releaseplan_id string     AD group to associate with the pool assignment
      --start_if_stopped string   Start stopped or suspended desktops (default "0")

Global Flags:
      --policy_id string           PolicyID to get pool assignments for
      --poolassignment_id string   PoolAssignmentID to get pool assignment for
➜  go-leostream-admin-cli git:(main) ✗ 
```


### Poolassignment delete command

Delete a poolassignment.
Required flags: --policy_id, --poolassignment_id

```bash
% leostream-admin-cli poolassignment delete --help
Delete Leostream pool assignment by ID. For example:
                leostream-admin-cli poolassignment delete --policy_id 1 --poolassignment_id 1

Usage:
  leostream-admin-cli poolassignment delete [flags]

Flags:
  -h, --help   help for delete

Global Flags:
      --policy_id string           PolicyID to get pool assignments for
      --poolassignment_id string   PoolAssignmentID to get pool assignment for
```
