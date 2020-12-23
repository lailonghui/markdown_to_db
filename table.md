

# 营运车安全监管平台数据表

#### 业务介绍

营运车安全监管平台是管理机构，政府机构，运输企业等进行实时监管公路客运、旅游客运、危化品车、网约出租车等重点车辆的管控平台。主要的功能有基础信息管理、动态监督管理、车辆违法信息管理、人车分离管理、黑名单、安全教育培训管理、驾驶员培训资格等，有利于提高对违法企业、违法驾驶员和违法车辆的监管效率。

# 1. 车辆信息管理

### 1.1 vehicle_info 车辆信息表

```sh
#车辆信息表主要存放车辆的基础信息
```

| Name          | Type   | Description                                         | **Required** | default                              |
| ------------- | ------ | --------------------------------------------------- | ------------ | ------------------------------------ |
| id            | bigint | 按指定方法生成，生成方法见下面说明                  | required     | 主键                                 |
| vehicle_id    | text   | 车辆外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| enterprise_id | text   | 所在企业id                                          | optional     | **enterprise_info表**的enterprise_id |
| department_id | text   | 所在部门id                                          | optional     | **department 部门信息表**            |

#### bigint主键ID生成方法

```json
# 参考链接，采用 Instagram 方案改进
https://rob.conery.io/2014/05/28/a-better-id-generator-for-postgresql/
# 生成都id将方便把数据同时插入多个数据库，以便PG分片和分布式部署时使用

create schema shard_1;
create sequence shard_1.global_id_sequence;

CREATE OR REPLACE FUNCTION shard_1.id_generator(OUT result bigint) AS $$
DECLARE
    our_epoch bigint := 1314220021721;
    seq_id bigint;
    now_millis bigint;
    -- the id of this DB shard, must be set for each
    -- schema shard you have - you could pass this as a parameter too
    shard_id int := 1;
BEGIN
    SELECT nextval('shard_1.global_id_sequence') % 1024 INTO seq_id;

    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
    result := (now_millis - our_epoch) << 23;
    result := result | (shard_id << 10);
    result := result | (seq_id);
END;
$$ LANGUAGE PLPGSQL;

select shard_1.id_generator();

# 运行该命令，会得到一个不错的，干净的bigint，可用于任何表的键

# 您可以通过以下方式声明用户表以使用此函数自动生成密钥
create table shard_1.users(
  id bigint not null default id_generator(),
  email varchar(255) not null unique,
  first varchar(50),
  last varchar(50)
)
```

### 1.2 operating_vehicle_ext  营运车信息扩展表

```sh
#营运车信息扩展表主要存放营运车平台所需要的信息
```

### 1.3 muck_truck_ext 渣土车信息扩展表

```sh
#渣土车信息扩展表主要存放渣土车平台所需要的信息
```

### 1.4 muck_truck_preview_number 渣土车车辆预编号表