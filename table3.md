# 工作流引擎work
## 1.数据表设计

### 1.1 proc_def 流程定义表

```sh
# proc_def(流程定义表)用于保存流程的配置
```

| Name         | Type        | Description                                     | **Required** | default  |
| ------------ | ----------- | ----------------------------------------------- | ------------ | -------- |
| id           | bigint      | 按指定方法生成                                  | required     | 主键     |
| proc_def_id  | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键 |
| name         | text        | 流程定义的名称，如:"请假流程"                   | optional     |          |
| version      | integer     | 流程定义的版本                                  | optional     |          |
| resource     | text        | 保存流程定义的具体配置,它是一个json格式的字符串 | optional     |          |
| user_id      | text        | 用户id                                          | optional     |          |
| user_name    | text        | 用户名称                                        | optional     |          |
| user_company | text        | 用户所在公司                                    | optional     |          |
| created_at   | timestamptz | 创建时间                                        | required     |          |
| created_by   | text        | 创建人                                          | required     |          |
| updated_at   | timestamptz | 修改时间                                        | optional     |          |
| updated_by   | text        | 修改人                                          | optional     |          |
| deleted_at   | timestamptz | 删除时间                                        | optional     |          |
| deleted_by   | text        | 删除人                                          | optional     |          |

### 1.2 proc_inst 流程实例表

```sh
#proc_inst(流程实例表)用于保存流程实例，当用户启动一个流程时，就会在这个表存入一个流程实例
```

| Name            | Type        | Description                                     | **Required** | default  |
| --------------- | ----------- | ----------------------------------------------- | ------------ | -------- |
| id              | bigint      | 按指定方法生成                                  | required     | 主键     |
| proc_inst_id    | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键 |
| proc_def_id     | text        | pro_def 流程定义表的id                          | optional     |          |
| proc_def_name   | text        | 流程定义的名称，如:"请假流程"                   | optional     |          |
| title           | text        | 标题，如："张三的请假流程"                      | optional     |          |
| department      | text        | 用户部门                                        | optional     |          |
| company         | text        | 用户公司                                        | optional     |          |
| node_id         | text        | 当前所处于节点的名称                            | optional     |          |
| candidate       | text        | 当前审批人或者审批用户组                        | optional     |          |
| task_id         | text        | 当前任务id                                      | optional     |          |
| start_time      | timestamptz | 开始时间                                        | optional     |          |
| end_time        | timestamptz | 结束时间                                        | optional     |          |
| duration        | bigint      | 持续时间                                        | optional     |          |
| start_user_name | text        | 发起人姓名                                      | optional     |          |
| is_finished     | boolean     | 是否结束                                        | optional     |          |
| created_at      | timestamptz | 创建时间                                        | required     |          |
| created_by      | text        | 创建人                                          | required     |          |
| updated_at      | timestamptz | 修改时间                                        | optional     |          |
| updated_by      | text        | 修改人                                          | optional     |          |
| deleted_at      | timestamptz | 删除时间                                        | optional     |          |
| deleted_by      | text        | 删除人                                          | optional     |          |

### 1.3 execution 执行流表

```sh
#execution(执行流表)用于保存执行流，当用户启动一个流程时，就会生成一条执行流，之后的流程就会按照执行流的顺序流转
#example: 开始->审批人1->审批人2->审批人3->结束
```

| Name          | Type        | Description                                     | **Required** | default  |
| ------------- | ----------- | ----------------------------------------------- | ------------ | -------- |
| id            | bigint      | 按指定方法生成                                  | required     | 主键     |
| execution_id  | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键 |
| rev           | integer     | 不明？？？                                      | optional     |          |
| proc_inst_id  | text        | 流程实例表id                                    | optional     |          |
| proc_def_id   | text        | pro_def 流程定义表的id                          | optional     |          |
| proc_def_name | text        | 流程定义的名称，如:"请假流程"                   | optional     |          |
| node_infos    | text[]      | 执行流经过的所有节点                            | optional     |          |
| is_active     | boolean     | 是否激活                                        | optional     |          |
| start_time    | timestamptz | 开始时间                                        | optional     |          |
| created_at    | timestamptz | 创建时间                                        | required     |          |
| created_by    | text        | 创建人                                          | required     |          |
| updated_at    | timestamptz | 修改时间                                        | optional     |          |
| updated_by    | text        | 修改人                                          | optional     |          |
| deleted_at    | timestamptz | 删除时间                                        | optional     |          |
| deleted_by    | text        | 删除人                                          | optional     |          |

### 1.4 task 任务表

```sh
# task(任务表)用于保存任务
```

| Name            | Type        | Description                                                  | **Required** | default  |
| --------------- | ----------- | ------------------------------------------------------------ | ------------ | -------- |
| id              | bigint      | 按指定方法生成                                               | required     | 主键     |
| task_id         | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键 |
| node_id         | text        | 当前执行流所在的节点                                         | optional     |          |
| step            | integer     | 表示任务对应的执行流位置，比如：有一个执行流：开始->审批人1->审批人2->审批人3->结束，<br />那么 step=0,则处于【开始】位置，step=1则处于【审批人1】位置 | optional     |          |
| proc_inst_id    | text        | 流程实例表id                                                 | optional     |          |
| assignee        | text        | 任务的处理人                                                 | optional     |          |
| member_count    | integer     | 表示当前任务需要多少人审批之后才能结束，默认是 1             | optional     | 1        |
| un_complete_num | integer     | 表示还有多少人没有审批，等于0代表会签已经全部审批结束，默认值为1 | optional     | 1        |
| agree_num       | integer     | 审批通过数                                                   | optional     |          |
| act_type        | text        | 表示任务类型 "or"表示或签，即一个人通过或者驳回就结束，"and"表示会签，要所有人通过就流转到下一步，如果有一个人驳回那么就跳转到上一步 | optional     |          |
| is_finished     | boolean     | 是否结束                                                     | optional     |          |
| created_at      | timestamptz | 创建时间                                                     | required     |          |
| created_by      | text        | 创建人                                                       | required     |          |
| updated_at      | timestamptz | 修改时间                                                     | optional     |          |
| updated_by      | text        | 修改人                                                       | optional     |          |
| deleted_at      | timestamptz | 删除时间                                                     | optional     |          |
| deleted_by      | text        | 删除人                                                       | optional     |          |

### 1.5 identity_link 关系表

```sh
# identidy_link(关系表)用于保存任务task的候选用户组或者候选用户以及用户所参与的流程信息
```

| Name             | Type        | Description                                                  | **Required** | default  |
| ---------------- | ----------- | ------------------------------------------------------------ | ------------ | -------- |
| id               | bigint      | 按指定方法生成                                               | required     | 主键     |
| identity_link_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键 |
| proc_inst_id     | text        | 流程实例表id                                                 | optional     |          |
| task_id          | text        | 对应任务task表的id                                           | optional     |          |
| proc_group       | text        | 表示当前审批的用户组                                         | optional     |          |
| type             | text        | 表示关系类型，有："candidate"和"participant"两种             | optional     |          |
| user_name        | text        | 用户名                                                       | optional     |          |
| step             | integer     | 表示任务对应的执行流位置，比如：有一个执行流：开始->审批人1->审批人2->审批人3->结束，<br />那么 step=0,则处于【开始】位置，step=1则处于【审批人1】位置 | optional     |          |
| company          | text        | 公司                                                         | optional     |          |
| comment          | text        | 评论                                                         | optional     |          |
| created_at       | timestamptz | 创建时间                                                     | required     |          |
| created_by       | text        | 创建人                                                       | required     |          |
| updated_at       | timestamptz | 修改时间                                                     | optional     |          |
| updated_by       | text        | 修改人                                                       | optional     |          |
| deleted_at       | timestamptz | 删除时间                                                     | optional     |          |
| deleted_by       | text        | 删除人                                                       | optional     |          |

### 1.6 his_proc_def 流程定义历史表

```sh
#历史数据表:字段同正常的表相同，每隔20秒，将已经结束的流程数据会自动迁移过来
```

| Name            | Type        | Description                                     | **Required** | default  |
| --------------- | ----------- | ----------------------------------------------- | ------------ | -------- |
| id              | bigint      | 按指定方法生成                                  | required     | 主键     |
| his_proc_def_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键 |
| name            | text        | 流程定义的名称，如:"请假流程"                   | optional     |          |
| version         | integer     | 流程定义的版本                                  | optional     |          |
| resource        | text        | 保存流程定义的具体配置,它是一个json格式的字符串 | optional     |          |
| user_id         | text        | 用户id                                          | optional     |          |
| user_name       | text        | 用户名称                                        | optional     |          |
| user_company    | text        | 用户所在公司                                    | optional     |          |
| created_at      | timestamptz | 创建时间                                        | required     |          |
| created_by      | text        | 创建人                                          | required     |          |
| updated_at      | timestamptz | 修改时间                                        | optional     |          |
| updated_by      | text        | 修改人                                          | optional     |          |
| deleted_at      | timestamptz | 删除时间                                        | optional     |          |
| deleted_by      | text        | 删除人                                          | optional     |          |

### 1.7 his_proc_inst 流程实例历史表

```sh
#历史数据表:字段同正常的表相同，每隔20秒，将已经结束的流程数据会自动迁移过来
```

| Name             | Type        | Description                                     | **Required** | default  |
| ---------------- | ----------- | ----------------------------------------------- | ------------ | -------- |
| id               | bigint      | 按指定方法生成                                  | required     | 主键     |
| his_proc_inst_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键 |
| proc_def_id      | text        | pro_def 流程定义表的id                          | optional     |          |
| proc_def_name    | text        | 流程定义的名称，如:"请假流程"                   | optional     |          |
| title            | text        | 标题，如："张三的请假流程"                      | optional     |          |
| department       | text        | 用户部门                                        | optional     |          |
| company          | text        | 用户公司                                        | optional     |          |
| node_id          | text        | 当前所处于节点的名称                            | optional     |          |
| candidate        | text        | 当前审批人或者审批用户组                        | optional     |          |
| task_id          | text        | 当前任务id                                      | optional     |          |
| start_time       | timestamptz | 开始时间                                        | optional     |          |
| end_time         | timestamptz | 结束时间                                        | optional     |          |
| duration         | integer     | 持续时间                                        | optional     |          |
| start_user_name  | text        | 发起人姓名                                      | optional     |          |
| is_finished      | boolean     | 是否结束                                        | optional     |          |
| created_at       | timestamptz | 创建时间                                        | required     |          |
| created_by       | text        | 创建人                                          | required     |          |
| updated_at       | timestamptz | 修改时间                                        | optional     |          |
| updated_by       | text        | 修改人                                          | optional     |          |
| deleted_at       | timestamptz | 删除时间                                        | optional     |          |
| deleted_by       | text        | 删除人                                          | optional     |          |

### 1.8 his_execution 执行流历史表

```sh
#历史数据表:字段同正常的表相同，每隔20秒，将已经结束的流程数据会自动迁移过来
```

| Name             | Type        | Description                                     | **Required** | default  |
| ---------------- | ----------- | ----------------------------------------------- | ------------ | -------- |
| id               | bigint      | 按指定方法生成                                  | required     | 主键     |
| his_execution_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键 |
| rev              | integer     | 不明？？？                                      | optional     |          |
| proc_inst_id     | text        | 流程实例表id                                    | optional     |          |
| proc_def_id      | text        | pro_def 流程定义表的id                          | optional     |          |
| proc_def_name    | text        | 流程定义的名称，如:"请假流程"                   | optional     |          |
| node_infos       | text[]      | 执行流经过的所有节点                            | optional     |          |
| is_active        | boolean     | 是否激活                                        | optional     |          |
| start_time       | timestamptz | 开始时间                                        | optional     |          |
| created_at       | timestamptz | 创建时间                                        | required     |          |
| created_by       | text        | 创建人                                          | required     |          |
| updated_at       | timestamptz | 修改时间                                        | optional     |          |
| updated_by       | text        | 修改人                                          | optional     |          |
| deleted_at       | timestamptz | 删除时间                                        | optional     |          |
| deleted_by       | text        | 删除人                                          | optional     |          |

### 1.9 his_task 任务历史表

```sh
#历史数据表:字段同正常的表相同，每隔20秒，将已经结束的流程数据会自动迁移过来
```

| Name            | Type        | Description                                                  | **Required** | default  |
| --------------- | ----------- | ------------------------------------------------------------ | ------------ | -------- |
| id              | bigint      | 按指定方法生成                                               | required     | 主键     |
| his_task_id     | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键 |
| node_id         | text        | 当前执行流所在的节点                                         | optional     |          |
| step            | integer     | 表示任务对应的执行流位置，比如：有一个执行流：开始->审批人1->审批人2->审批人3->结束，<br />那么 step=0,则处于【开始】位置，step=1则处于【审批人1】位置 | optional     |          |
| proc_inst_id    | text        | 流程实例表id                                                 | optional     |          |
| assignee        | text        | 任务的处理人                                                 | optional     |          |
| member_count    | integer     | 表示当前任务需要多少人审批之后才能结束，默认是 1             | optional     | 1        |
| un_complete_num | integer     | 表示还有多少人没有审批，等于0代表会签已经全部审批结束，默认值为1 | optional     | 1        |
| agree_num       | integer     | 审批通过数                                                   | optional     |          |
| act_type        | text        | 表示任务类型 "or"表示或签，即一个人通过或者驳回就结束，"and"表示会签，要所有人通过就流转到下一步，如果有一个人驳回那么就跳转到上一步 | optional     |          |
| is_finished     | boolean     | 是否结束                                                     | optional     |          |
| created_at      | timestamptz | 创建时间                                                     | required     |          |
| created_by      | text        | 创建人                                                       | required     |          |
| updated_at      | timestamptz | 修改时间                                                     | optional     |          |
| updated_by      | text        | 修改人                                                       | optional     |          |
| deleted_at      | timestamptz | 删除时间                                                     | optional     |          |
| deleted_by      | text        | 删除人                                                       | optional     |          |

### 1.10 his_identity_link 关系历史表

```sh
#历史数据表:字段同正常的表相同，每隔20秒，将已经结束的流程数据会自动迁移过来
```

| Name                 | Type        | Description                                                  | **Required** | default  |
| -------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------- |
| id                   | bigint      | 按指定方法生成                                               | required     | 主键     |
| his_identity_link_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键 |
| proc_inst_id         | text        | 流程实例表id                                                 | optional     |          |
| task_id              | text        | 对应任务task表的id                                           | optional     |          |
| proc_group           | text        | 表示当前审批的用户组                                         | optional     |          |
| type                 | text        | 表示关系类型，有："candidate"和"participant"两种             | optional     |          |
| user_name            | text        | 用户名                                                       | optional     |          |
| step                 | integer     | 表示任务对应的执行流位置，比如：有一个执行流：开始->审批人1->审批人2->审批人3->结束，<br />那么 step=0,则处于【开始】位置，step=1则处于【审批人1】位置 | optional     |          |
| company              | text        | 公司                                                         | optional     |          |
| comment              | text        | 评论                                                         | optional     |          |
| created_at           | timestamptz | 创建时间                                                     | required     |          |
| created_by           | text        | 创建人                                                       | required     |          |
| updated_at           | timestamptz | 修改时间                                                     | optional     |          |
| updated_by           | text        | 修改人                                                       | optional     |          |
| deleted_at           | timestamptz | 删除时间                                                     | optional     |          |
| deleted_by           | text        | 删除人                                                       | optional     |          |