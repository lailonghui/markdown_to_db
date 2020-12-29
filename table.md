# 营运车安全监管平台数据表

#### 业务介绍

营运车安全监管平台是管理机构，政府机构，运输企业等进行实时监管公路客运、旅游客运、危化品车、网约出租车等重点车辆的管控平台。主要的功能有基础信息管理、动态监督管理、车辆违法信息管理、人车分离管理、黑名单、安全教育培训管理、驾驶员培训资格等，有利于提高对违法企业、违法驾驶员和违法车辆的监管效率。

# 1. 车辆信息管理

### 1.1 vehicle_info 车辆信息表

```sh
#车辆信息表主要存放车辆的基础信息
```

| Name                          | Type        | Description                                                  | **Required** | default                                                      |
| ----------------------------- | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------------ |
| id                            | bigint      | 按指定方法生成，生成方法见下面说明                           | required     | 主键                                                         |
| vehicle_id                    | text        | 车辆外部编码，由golang程序生成的xid，暴露到外部使用          | required     | 联合主键                                                     |
| enterprise_id                 | text        | 所在企业id                                                   | optional     | **enterprise_info表**的enterprise_id                         |
| department_id                 | text        | 所在部门id                                                   | optional     | **department 部门信息表**                                    |
| industry_category             | integer     | 行业类别                                                     | optional     | **行业类别**字典                                             |
| business_scope                | integer     | 经营范围                                                     | optional     | **经营范围**字典                                             |
| vehicle_type                  | integer     | 车辆类型                                                     | optional     | **车辆类型**字典                                             |
| operating_type                | integer     | 营运类型                                                     | optional     | **营运类型**字典                                             |
| operating_state               | integer     | 营运状态                                                     | optional     | **营运状态**字典                                             |
| operating_route               | text        | 营运线路                                                     | optional     |                                                              |
| terminal_id                   | text        | 终端ID                                                       | optional     |                                                              |
| is_apply_install_terminal     | boolean     | 是否申请安装智能终端                                         | optional     |                                                              |
| license_plate_number          | text        | 车牌号                                                       | optional     |                                                              |
| license_plate_color           | integer     | 车牌颜色                                                     | optional     | **车牌颜色**字典                                             |
| license_plate_type            | integer     | 号牌种类                                                     | optional     | **号牌种类**字典                                             |
| vehicle_identification_number | text        | 车架号(后6位)                                                | optional     | 车辆识别代号vin,如D02133                                     |
| road_transport_license_number | text        | 道路运输证号                                                 | optional     |                                                              |
| heavy                         | numeric     | 吨位                                                         | optional     |                                                              |
| seats                         | integer     | 座位                                                         | optional     |                                                              |
| vehicle_manager               | text        | 机动车管理人                                                 | optional     |                                                              |
| vehicle_manager_phone         | text        | 机动车管理人联系电话                                         | optional     |                                                              |
| vehicle_manager_id_card       | text        | 机动车管理人身份证                                           |              |                                                              |
| owner                         | text        | 机动车所有人（六合一）                                       | optional     |                                                              |
| inspection_date               | timestamptz | 检验日期（六合一）                                           | optional     |                                                              |
| retirement_date               | timestamptz | 报废日期（六合一）                                           | optional     |                                                              |
| use_nature                    | text        | 使用性质（六合一）                                           | optional     |                                                              |
| vehicle_state                 | integer     | 机动车状态                                                   | optional     | **车辆状态**字典                                             |
| update_time_in                | timestamptz | 内网更新时间                                                 | optional     |                                                              |
| remark_in                     | text        | 车辆信息同步内网反馈信息                                     | optional     | 车辆信息同步到公安内网后内网的反馈内容，如车牌号填写错误会反馈车辆号牌错误 |
| is_complete                   | boolean     | 是否完成                                                     | optional     | 用于标志车辆资料是否处于确定状态。未确定状态的车辆信息在系统上除车辆管理外的功能中都查不到 |
| driving_licensee_pic          | text        | 行驶证照片,云储存系统返回的路径                              | optional     |                                                              |
| is_active                     | boolean     | 是否激活                                                     | optional     |                                                              |
| is_input                      | boolean     | 是否录入完成                                                 | optional     |                                                              |
| car_rental_price              | numeric     | 租车标准价格                                                 | optional     |                                                              |
| insurance_company             | integer     | 投保公司                                                     | optional     | **投保公司**字典                                             |
| insurance_date                | timestamptz | 投保日期                                                     | optional     |                                                              |
| vehicle_maintenances          | jsonb[]     | 维保数据数组，字段包括: 1.maintenance_ date 维保时间<br />2.maintenance_ kilometers 维保公里数 | optional     |                                                              |
| vehicle_displacement          | text        | 汽车排量                                                     | optional     |                                                              |
| vehicle_brand                 | integer     | 车辆品牌                                                     | optional     | **车辆品牌**字典                                             |
| quasi_driving_models          | integer     | 准驾车型                                                     | optional     | **准驾车型**字典                                             |
| is_upload_province            | boolean     | 是否上传省厅                                                 | optional     |                                                              |
| check_state                   | integer     | 校验状态                                                     | optional     | **车辆校验状态**字典                                         |
| is_import                     | boolean     | 是否导入                                                     | optional     | 是否通过外部导入的车辆信息                                   |
| is_engineering_vehicle        | boolean     | 是否工程运输车                                               | optional     |                                                              |
| is_catalog_library            | boolean     | 是否目录库                                                   | optional     |                                                              |
| remarks                       | text        | 备注                                                         | optional     |                                                              |
| is_deleted                    | boolean     | 是否删除                                                     | optional     |                                                              |
| record_at                     | timestamptz | 登记时间                                                     | optional     |                                                              |
| record_by                     | text        | 登记人                                                       | optional     | **system_user表**的user_id                                   |
| created_at                    | timestamptz | 创建时间                                                     | required     |                                                              |
| created_by                    | text        | 创建人                                                       | required     | **system_user表**的user_id                                   |
| updated_at                    | timestamptz | 修改时间                                                     | optional     |                                                              |
| updated_by                    | text        | 修改人                                                       | optional     | **system_user表**的user_id                                   |
| deleted_at                    | timestamptz | 删除时间                                                     | optional     |                                                              |
| deleted_by                    | text        | 删除人                                                       | optional     | **system_user表**的user_id                                   |

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

| Name                            | Type        | Description                              | **Required** | default                    |
| ------------------------------- | ----------- | ---------------------------------------- | ------------ | -------------------------- |
| id                              | bigint      | 按指定方法生成                           | required     | 主键                       |
| vehicle_id                      | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                            |
| agent                           | text        | 代理商                                   | optional     |                            |
| contact_number                  | text        | 合同编号                                 | optional     |                            |
| platform                        | integer     | 平台                                     | optional     | **平台标识**字典           |
| driving_license_owner           | text        | 行驶证所有人                             | optional     |                            |
| speed_mode_status_time          | timestamptz | 速度模式状态获得时间                     | optional     |                            |
| speed_mode_status               | text        | 速度模式状态                             | optional     |                            |
| scrap_reason                    | text        | 报废原因                                 | optional     |                            |
| scrap_time                      | timestamptz | 报废时间                                 | optional     |                            |
| scrap_time_check                | integer     | 报废时间审核                             | optional     |                            |
| administrative_region           | text        | 行政区域                                 | optional     |                            |
| license_plate_photo             | text        | 车牌照片,云储存系统返回的路径            | optional     |                            |
| other_photo                     | text        | 其他照片,云储存系统返回的路径            | optional     |                            |
| serial_number                   | text        | 编号                                     | optional     |                            |
| is_beidou                       | boolean     | 是否北斗部标平台                         | optional     |                            |
| is_in_operating_system          | boolean     | 是否在运证系统                           | optional     |                            |
| is_in_upload_platform           | boolean     | 是否上传货运平台                         | optional     |                            |
| is_supervise                    | boolean     | 是否监管                                 | optional     |                            |
| is_need_supervise               | boolean     | 是否需要监管                             | optional     |                            |
| is_function_ok                  | boolean     | 是否功能测试通过                         | optional     |                            |
| is_block                        | boolean     | 是否屏蔽                                 | optional     |                            |
| is_applay_terminal_installation | boolean     | 是否办理终端安装证明                     | optional     |                            |
| first_online_time               | timestamptz | 第一次上线时间                           | optional     |                            |
| last_binding_terminal_time      | timestamptz | 最后一次绑定终端时间                     | optional     |                            |
| service_expiration_time         | timestamptz | 服务到期时间                             | optional     |                            |
| contract_time                   | timestamptz | 合同时间                                 | optional     |                            |
| installation_time               | timestamptz | 安装时间                                 | optional     |                            |
| created_at                      | timestamptz | 创建时间                                 | required     |                            |
| created_by                      | text        | 创建人                                   | required     | **system_user表**的user_id |
| updated_at                      | timestamptz | 修改时间                                 | optional     |                            |
| updated_by                      | text        | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz | 删除时间                                 | optional     |                            |
| deleted_by                      | text        | 删除人                                   | optional     | **system_user表**的user_id |

### 1.3 muck_truck_ext 渣土车信息扩展表

```sh
#渣土车信息扩展表主要存放渣土车平台所需要的信息
```

| Name                                     | Type        | Description                                            | **Required** | default                    |
| ---------------------------------------- | ----------- | ------------------------------------------------------ | ------------ | -------------------------- |
| id                                       | bigint      | 按指定方法生成                                         | required     | 主键                       |
| vehicle_id                               | text        | **vehicle_info** 车辆信息表 的vehicle_id               | required     |                            |
| vehicle_picture                          | text        | 车辆图片                                               | optional     |                            |
| engine_number                            | text        | 发动机号                                               | optional     |                            |
| review_time                              | timestamptz | 审核时间                                               | optional     |                            |
| reviewer                                 | text        | 审核人                                                 | optional     | **system_user表**的user_id |
| review_notes                             | text        | 审核备注                                               | optional     |                            |
| review_status                            | text        | 审核状态                                               | optional     |                            |
| insurance_expiry_time                    | timestamptz | 保险到期时间                                           | optional     |                            |
| annual_inspection_expiration_time        | timestamptz | 年检到期时间                                           | optional     |                            |
| vehicle_description                      | text        | 车辆描述                                               | optional     |                            |
| driving_license_id_number                | text        | 行驶证登记的车主身份证号                               | optional     |                            |
| driving_license_contact_phone            | text        | 行驶证登记的车主联系电话                               | optional     |                            |
| is_due_security_alarm_processing         | boolean     | 是否安检到期报警处理                                   | optional     |                            |
| processor                                | text        | 处理人                                                 | optional     | **system_user表**的user_id |
| processing_time                          | timestamptz | 处理时间                                               | optional     |                            |
| processing_notes                         | text        | 处理备注                                               | optional     |                            |
| is_send_sms                              | boolean     | 是否发送短信                                           | optional     |                            |
| is_blacklist                             | boolean     | 是否黑名单                                             | optional     |                            |
| blacklist_deadline                       | timestamptz | 黑名单截止日期                                         | optional     |                            |
| compulsory_traffic_insurance_expiry_date | timestamptz | 交强险到期时间                                         | optional     |                            |
| owner_id_photo                           | text        | 车主身份证照片                                         | optional     |                            |
| vehicle_nature                           | integer     | 车辆性质  1.本企业车辆  2.企业挂靠车辆  3.车队挂靠车辆 | optional     |                            |
| actual_owner_name                        | text        | 实际车主姓名                                           | optional     |                            |
| actual_owner_id_number                   | text        | 实际车主身份证号                                       | optional     |                            |
| actual_owner_id_photo                    | text        | 实际车主身份证照片                                     | optional     |                            |
| actual_owner_contact_phone               | text        | 实际车主联系电话                                       | optional     |                            |
| illegal_number_endorsement               | text        | 违法编号签注                                           | optional     |                            |
| illegal_notice_number_endorsement        | text        | 违法通知书编号签注                                     | optional     |                            |
| illegal_compulsory_measures_number       | text        | 违法强制措施编号签注                                   | optional     |                            |
| incident_number_endorsement              | text        | 事故编号签注                                           | optional     |                            |
| is_muck_office_audit                     | boolean     | 是否渣土办审核                                         | optional     |                            |
| is_first_register                        | boolean     | 是否首次注册                                           | optional     |                            |
| secondary_maintenance_location           | text        | 二级维护地点                                           | optional     |                            |
| secondary_maintenance_expiry_date        | timestamptz | 二级维护检测到期时间                                   | optional     |                            |
| total_mass                               | numeric     | 总质量                                                 | optional     |                            |
| axes_number                              | integer     | 轴数                                                   | optional     |                            |
| driving_liscense_owner_id_photo          | text        | 行驶证登记的车主身份证照片                             | optional     |                            |
| initial_registration_date                | timestamptz | 初次登记日期                                           | optional     |                            |
| vehicle_operating_certificate_number     | text        | 车辆营运证号                                           | optional     |                            |
| vehicle_operating_certificate_photo      | text        | 车辆营运证照片                                         | optional     |                            |
| issue_date_of_quangong_number            | timestamptz | 泉工号发放日期                                         | optional     |                            |
| apply_inspection_time                    | timestamptz | 申请验车时间                                           | optional     |                            |
| signing_acceptance_application_time      | timestamptz | 签收验车申请时间                                       | optional     |                            |
| third_party_insurance_expiry_time        | timestamptz | 第三者保险到期时间                                     | optional     |                            |
| third_party_insurance_coverage           | text        | 第三者保额                                             | optional     |                            |
| third_party_insurance_policy_picture     | text        | 第三者保单图片                                         | optional     |                            |
| compulsory_insurance_policy_picture      | text        | 交强险保单图片                                         | optional     |                            |
| is_deleted                               | boolean     | 是否删除                                               | optional     |                            |
| sim_card_number                          | text        | SIM卡号                                                | optional     |                            |
| registration_date                        | timestamptz | 注册日期                                               | optional     |                            |
| area_id                                  | bigint      | 地区ID                                                 | optional     |                            |
| expiry_date                              | timestamptz | 费用到期时间                                           | optional     |                            |
| muck_truck_type                          | integer     | 渣土车类型（1.渣土车 2.混凝土车 3.砂石车）             | optional     | **工程运输车车辆类型**字典 |
| is_reserve_library                       | boolean     | 是否预备库                                             | optional     |                            |
| self_number                              | text        | 自编号                                                 | optional     |                            |
| mobile_card_location                     | text        | 移动办卡地                                             | optional     |                            |
| load_category                            | text        | 装载类别                                               | optional     |                            |
| created_at                               | timestamptz | 创建时间                                               | required     |                            |
| created_by                               | text        | 创建人                                                 | required     | **system_user表**的user_id |
| updated_at                               | timestamptz | 修改时间                                               | optional     |                            |
| updated_by                               | text        | 修改人                                                 | optional     | **system_user表**的user_id |
| deleted_at                               | timestamptz | 删除时间                                               | optional     |                            |
| deleted_by                               | text        | 删除人                                                 | optional     | **system_user表**的user_id |

### 1.4 muck_truck_preview_number 渣土车车辆预编号表

| Name                         | Type        | Description                                                  | **Required** | default                    |
| ---------------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                               | required     | 主键                       |
| muck_truck_preview_number_id | text        | 渣土车车辆预编号外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                   | text        | **vehicle_info** 车辆信息表 的vehicle_id                     | required     |                            |
| svn_number                   | text        | 预编号                                                       | optional     |                            |
| registration_time            | text        | 登记时间                                                     | optional     |                            |
| confirmor                    | text        | 确认人                                                       | optional     | **system_user表**的user_id |
| confirm_time                 | timestamptz | 确认时间                                                     | optional     |                            |
| confirm_status               | integer     | 确认状态                                                     | optional     |                            |
| front_license_plate          | text        | 前车牌                                                       | optional     |                            |
| rear_license_plate           | text        | 后车牌                                                       | optional     |                            |
| side_license_plate           | text        | 侧车牌                                                       | optional     |                            |
| original_number              | text        | 原编号                                                       | optional     |                            |
| is_review_automatically      | boolean     | 是否自动审核                                                 | optional     |                            |
| production_status            | integer     | 制作状态                                                     | optional     |                            |
| production_time              | timestamptz | 制作时间                                                     | optional     |                            |
| marking_time                 | timestamptz | 制作中时间                                                   | optional     |                            |
| production_times             | integer     | 制作次数                                                     | optional     |                            |
| submit_production_time       | timestamptz | 提交制牌厂时间                                               | optional     |                            |
| contact_person               | text        | 联系人                                                       | optional     |                            |
| contact_phone                | text        | 联系电话                                                     | optional     |                            |
| work_number_plate_color      | text        | 工号牌颜色（green.绿色 yellow.黄色）                         | optional     |                            |
| initial_registration_date    | timestamptz | 初次登记日期                                                 | optional     |                            |
| unlawful_violation_number    | integer     | 违法未处理数                                                 | optional     |                            |
| is_register_sale_order       | boolean     | 是否登记销售订单                                             | optional     |                            |
| remarks                      | text        | 备注                                                         | optional     |                            |
| is_deleted                   | boolean     | 是否删除                                                     | optional     |                            |
| created_at                   | timestamptz | 创建时间                                                     | required     |                            |
| created_by                   | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                   | timestamptz | 修改时间                                                     | optional     |                            |
| updated_by                   | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz | 删除时间                                                     | optional     |                            |
| deleted_by                   | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### 1.5 muck_truck_worker_identity_card_orders 渣土车工号牌制作订单表

| Name                        | Type        | Description                                                  | **Required** | default                    |
| --------------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                          | bigint      | 按指定方法生成                                               | required     | 主键                       |
| muck_truck_worker_orders_id | text        | 渣土车工号牌制作订单外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                  | text        | **vehicle_info** 车辆信息表 的vehicle_id                     | required     |                            |
| preview_number_id           | text        | **muck_truck_preview_number**  渣土车车辆预编号表的id        | required     |                            |
| created_at                  | timestamptz | 创建时间                                                     | required     |                            |
| created_by                  | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                  | timestamptz | 修改时间                                                     | optional     |                            |
| updated_by                  | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz | 删除时间                                                     | optional     |                            |
| deleted_by                  | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### 1.6 vehicle_supervision_photo 车辆监控图片表

| Name                       | Type        | Description                                                 | **Required** | default                              |
| -------------------------- | ----------- | ----------------------------------------------------------- | ------------ | ------------------------------------ |
| id                         | bigint      | 按指定方法生成                                              | required     | 主键                                 |
| supervision_photo_id       | text        | 车辆监控图片外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| vehicle_id                 | text        | 车辆ID                                                      | required     | **vehicle_info**表vehicle_id         |
| driver_id                  | text        | 驾驶员id                                                    | optional     | **driver_info 表**的driver_id        |
| enterprise_id              | text        | 所在企业id                                                  | optional     | **enterprise_info表**的enterprise_id |
| camera_id                  | integer     | 摄像头ID                                                    | optional     | **摄像头ID**字典                     |
| photo_condition            | text        | 拍照条件                                                    | optional     | **拍照条件**字典                     |
| update_time                | timestamptz | 终端上报时间                                                | optional     |                                      |
| monitoring_pic_name        | text        | 监控图片名称                                                | optional     |                                      |
| monitoring_pic_address     | text        | 监控图片地址                                                | optional     |                                      |
| monitoring_pic_upload_time | timestamptz | 监控图片上传时间                                            | optional     |                                      |
| imel                       | text        | 终端IMEI                                                    | optional     | 国际移动设备标识别码                 |
| sim_number                 | text        | SIM卡号                                                     | optional     |                                      |

### 1.7 owner_info 车主信息表

| Name          | Type        | Description                                             | **Required** | default                              |
| ------------- | ----------- | ------------------------------------------------------- | ------------ | ------------------------------------ |
| id            | bigint      | 按指定方法生成                                          | required     | 主键                                 |
| owner_id      | text        | 车主信息外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| department_id | text        | 所在部门id                                              | optional     | **department 部门信息表**            |
| name          | text        | 车主姓名                                                | optional     |                                      |
| address       | text        | 联系地址                                                | optional     |                                      |
| cellphone     | text        | 固定电话                                                | optional     |                                      |
| telephone     | text        | 手机号码                                                | optional     |                                      |
| email         | text        | 邮箱地址                                                | optional     |                                      |
| expiry_date   | timestamptz | 证件过期日期                                            | optional     |                                      |
| id_number     | text        | 身份证号                                                | optional     |                                      |
| remarks       | text        | 备注                                                    | optional     |                                      |
| sex           | integer     | 车主性别                                                | optional     | **性别**字典                         |
| agent         | text        | 代理商                                                  | optional     | **enterprise_info表**的enterprise_id |
| operator      | text        | 运营商                                                  | optional     | **enterprise_info表**的enterprise_id |
| is_deleted    | boolean     | 是否删除                                                | optional     |                                      |
| created_at    | timestamptz | 创建时间                                                | required     |                                      |
| created_by    | text        | 创建人                                                  | required     | **system_user表**的user_id           |
| updated_at    | timestamptz | 修改时间                                                | optional     |                                      |
| updated_by    | text        | 修改人                                                  | optional     | **system_user表**的user_id           |
| deleted_at    | timestamptz | 删除时间                                                | optional     |                                      |
| deleted_by    | text        | 删除人                                                  | optional     | **system_user表**的user_id           |

### 1.8 jj_vehicle 公安内网六合一平台同步车辆表(不修改字段)

| Name       | Type        | Description                              | **Required** | default |
| ---------- | ----------- | ---------------------------------------- | ------------ | ------- |
| id         | bigint      | 按指定方法生成                           | required     | 主键    |
| vehicle_id | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |         |
| HPHM       | text        | 车牌号码                                 | optional     |         |
| HPZL       | text        | 号牌种类                                 | optional     |         |
| SYR        | text        | 所有人                                   | optional     |         |
| SYXZ       | text        | 使用性质                                 | optional     |         |
| ZZL        | numeric     | 总质量                                   | optional     |         |
| CLLX       | text        | 车辆类型                                 | optional     |         |
| YXQZ       | timestamptz | 有效期至                                 | optional     |         |
| QZBFQZ     | timestamptz | 强制报废期止                             | optional     |         |
| CLSBDH     | text        | 车辆识别代号                             | optional     |         |
| COUNTY     | text        | 所在县                                   | optional     |         |
| LXDZ       | text        | 联系地址                                 | optional     |         |
| OTHER      | text        | 固话                                     | optional     |         |
| LXDH       | text        | 联系电话                                 | optional     |         |
| ZT         | integer     | 机动车状态                               | optional     |         |
| VEHMONTYPE | text        | 经营范围                                 | optional     |         |
| CCDJRQ     | timestamptz | 初次登记日期                             | optional     |         |
| DJRQ       | timestamptz | 登记日期                                 | optional     |         |
| FPRQ       | timestamptz | 发牌日期                                 | optional     |         |
| UPDATETIME | timestamptz | 更新时间                                 | optional     |         |
| XH         | text        | 型号                                     | optional     |         |
| IS_DETELED | numeric     | 是否删除                                 | optional     |         |

### 1.9 vehicle_enterprise_change_review 车辆单位变更审核表

| Name                                | Type        | Description                                     | **Required** | default                    |
| ----------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                  | bigint      | 按指定方法生成                                  | required     | 主键                       |
| vehicle_enterprise_change_review_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                          | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| target_enterprise_id                | text        | 目标企业ID                                      | optional     |                            |
| remarks                             | text        | 备注                                            | optional     |                            |
| upload_document_src                 | text        | 上传证明文件                                    | optional     |                            |
| review_status                       | integer     | 审核状态  0.未审核  1.审核通过  2.退回          | optional     |                            |
| change_type                         | integer     | 变更类型  1.本地区变更  2.跨区变更              | optional     |                            |
| original_area                       | text        | 车辆原所属地区                                  | optional     |                            |
| changed_area                        | text        | 变更地区                                        | optional     |                            |
| changed_self_number                 | text        | 变更后自编号                                    | optional     |                            |
| is_deleted                          | boolean     | 是否删除                                        | optional     |                            |
| created_at                          | timestamptz | 创建时间                                        | required     |                            |
| created_by                          | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                          | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                          | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                          | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                          | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 1.10 vehicle_enterprise_change_log 车辆单位变更日志表

| Name                             | Type        | Description                                     | **Required** | default                                                     |
| -------------------------------- | ----------- | ----------------------------------------------- | ------------ | ----------------------------------------------------------- |
| id                               | bigint      | 按指定方法生成                                  | required     | 主键                                                        |
| vehicle_enterprise_change_log_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                    |
| changed_type                     | integer     | 变更类型                                        | optional     |                                                             |
| changed_step                     | integer     | 变更步骤                                        | optional     |                                                             |
| operator                         | text        | 操作人                                          | optional     | **system_user表**的user_id                                  |
| change_review_id                 | text        | 变更审核表ID                                    | optional     | **vehicle_enterprise_change_review 车辆单位变更审核表**的id |
| review_status                    | integer     | 审核状态  0.未审核  1.通过  2.退回  3.撤销      | optional     |                                                             |
| created_at                       | timestamptz | 创建时间                                        | required     |                                                             |
| created_by                       | text        | 创建人                                          | required     | **system_user表**的user_id                                  |
| updated_at                       | timestamptz | 修改时间                                        | optional     |                                                             |
| updated_by                       | text        | 修改人                                          | optional     | **system_user表**的user_id                                  |
| deleted_at                       | timestamptz | 删除时间                                        | optional     |                                                             |
| deleted_by                       | text        | 删除人                                          | optional     | **system_user表**的user_id                                  |

### 1.11 vehicle_operation_history 车辆操作历史记录表

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| vehicle_operation_history_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                   | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| remarks                      | text        | 备注                                            | optional     |                            |
| operation_type               | integer     | 操作类型  1.添加  2.删除                        | optional     |                            |
| operator                     | text        | 操作人                                          | optional     | **system_user表**的user_id |
| review_status                | integer     | 审核状态  0.未审批  1.已审批                    | optional     |                            |
| area                         | text        | 地区                                            | optional     |                            |
| reviewer                     | text        | 审核人                                          | optional     | **system_user表**的user_id |
| created_at                   | timestamptz | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 1.12 vehicle_online_time 车辆在线时长表

| Name           | Type        | Description                                     | **Required** | default                    |
| -------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id             | bigint      | 按指定方法生成                                  | required     | 主键                       |
| online_time_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id     | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| online_time    | numeric     | 在线时长                                        | optional     |                            |
| offline_time   | numeric     | 离线时长                                        | optional     |                            |
| total_time     | numeric     | 总时长                                          | optional     |                            |
| is_online      | boolean     | 是否在线                                        | optional     |                            |
| created_at     | timestamptz | 创建时间                                        | required     |                            |
| created_by     | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at     | timestamptz | 修改时间                                        | optional     |                            |
| updated_by     | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at     | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by     | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 1.13 temp_identity_card_download_log 临时工号牌下载记录表

| Name                           | Type        | Description                                     | **Required** | default                    |
| ------------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                             | bigint      | 按指定方法生成                                  | required     | 主键                       |
| temp_identity_card_download_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                     | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| valid_from                     | timestamptz | 有效期起始                                      | optional     |                            |
| valid_until                    | timestamptz | 有效期截止                                      | optional     |                            |
| operator                       | text        | 操作人                                          | optional     | **system_user表**的user_id |
| is_deleted                     | boolean     | 是否删除                                        | optional     |                            |
| created_at                     | timestamptz | 创建时间                                        | required     |                            |
| created_by                     | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                     | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                     | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                     | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 1.14 muck_truck_test_situation 新型渣土车测试情况表

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| muck_truck_test_situation_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                   | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| airtight_state               | integer     | 密闭状态                                        | optional     |                            |
| load_state                   | integer     | 载重状态                                        | optional     |                            |
| lifting_status               | integer     | 举升状态                                        | optional     |                            |
| video                        | integer     | 视频                                            | optional     |                            |
| fingerprint                  | integer     | 指纹                                            | optional     |                            |
| car_lock_function            | integer     | 锁车功能                                        | optional     |                            |
| speed_limit_function         | integer     | 限速功能                                        | optional     |                            |
| ministry_standard_function   | integer     | 部标功能                                        | optional     |                            |
| weight_limit_function        | integer     | 限举功能                                        | optional     |                            |
| left_turn_light              | integer     | 左转灯                                          | optional     |                            |
| right_turn_right             | integer     | 右转灯                                          | optional     |                            |
| high_beam                    | integer     | 远光灯                                          | optional     |                            |
| low_beam                     | integer     | 近光灯                                          | optional     |                            |
| brake                        | integer     | 刹车                                            | optional     |                            |
| speed                        | integer     | 车速                                            | optional     |                            |
| is_detect_illegal_spoil      | boolean     | 是否检测非法弃土                                | optional     |                            |
| is_detect_illegal_start      | boolean     | 是否检测违规启动                                | optional     |                            |
| is_passed                    | boolean     | 是否通过                                        | optional     |                            |
| is_deleted                   | boolean     | 是否删除                                        | optional     |                            |
| created_at                   | timestamptz | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 1.15 vehicle_status_change_log 车辆状态变更记录表

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| vehicle_status_change_log_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                   | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| terminal_id                  | text        | 终端ID                                          | optional     |                            |
| start_time                   | timestamptz | 开始时间                                        | optional     |                            |
| end_time                     | timestamptz | 结束时间                                        | optional     |                            |
| vehicle_status_type          | integer     | 车辆状态类型(车厢状态,举升状态,载重状态)        | optional     | **车辆状态类型**字典       |
| value                        | text        | 值                                              | optional     |                            |
| is_completed                 | boolean     | 是否完成                                        | optional     |                            |
| created_at                   | timestamptz | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 1.16 vehicle_state_latest 新型车辆最新状态表

| Name                    | Type        | Description                                     | **Required** | default                              |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                                 |
| vehicle_state_latest_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| vehicle_id              | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                                      |
| operation_type          | text        | 操作类型（lock.锁车 speed.限速 lift.限举）      | optional     |                                      |
| operator                | text        | 操作人                                          | optional     | **system_user表**的user_id           |
| operator_institution    | text        | 操作人单位                                      | optional     | **enterprise_info表**的enterprise_id |
| status                  | text        | 状态                                            | optional     |                                      |
| speed_limit             | text        | 限速值                                          | optional     |                                      |
| created_at              | timestamptz | 创建时间                                        | required     |                                      |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at              | timestamptz | 修改时间                                        | optional     |                                      |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at              | timestamptz | 删除时间                                        | optional     |                                      |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id           |

### 1.17 vehicle_increment_record 车辆新增记录

| Name                 | Type        | Description                                     | **Required** | default                              |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                                 |
| increment_record_id  | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                                      |
| enterprise_id        | text        | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id |
| self_number          | text        | 自编号                                          | optional     |                                      |
| operation_type       | text        | 操作类型                                        | optional     |                                      |
| license_plate_number | text        | 车牌号                                          | optional     |                                      |
| enterprise_name      | text        | 单位名称                                        | optional     |                                      |
| area                 | text        | 地区                                            | optional     |                                      |
| original_number      | text        | 原编号                                          | optional     |                                      |
| operator             | text        | 运营商                                          | optional     | **enterprise_info表**的enterprise_id |
| created_at           | timestamptz | 创建时间                                        | required     |                                      |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at           | timestamptz | 修改时间                                        | optional     |                                      |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at           | timestamptz | 删除时间                                        | optional     |                                      |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id           |

### 1.18 outage_registration 停运登记表

| Name                         | Type        | Description                                     | **Required** | default                                                      |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                                                         |
| outage_registration_id       | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| vehicle_id                   | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                                                              |
| user_id                      | text        | 用户                                            | optional     | **system_user表**的user_id                                   |
| outage_start_time            | timestamptz | 停运起始时间                                    | optional     |                                                              |
| outage_end_time              | timestamptz | 停运截止时间                                    | optional     |                                                              |
| review_status                | text        | 审核状态                                        | optional     |                                                              |
| reviewer                     | text        | 审核人                                          | optional     | **system_user表**的user_id                                   |
| review_time                  | timestamptz | 审核时间                                        | optional     |                                                              |
| outage_filing_upload_file_id | text        | 停运报备上传文件表ID                            | optional     | **outage_filing_upload_file** 的outage_filing_upload_file_id |
| outage_start_coordinate      | point       | 停运起始经纬度                                  | optional     |                                                              |
| outage_end_coordinate        | point       | 停运结束经纬度                                  | optional     |                                                              |
| online_time                  | timestamptz | 上线时间                                        | optional     |                                                              |
| outage_start_position        | text        | 停运起始位置                                    | optional     |                                                              |
| is_management_review         | boolean     | 是否管理部门审核                                | optional     |                                                              |
| is_invalid                   | boolean     | 是否失效                                        | optional     |                                                              |
| is_latest                    | boolean     | 是否最新                                        | optional     |                                                              |
| is_deleted                   | boolean     | 是否删除                                        | optional     |                                                              |
| created_at                   | timestamptz | 创建时间                                        | required     |                                                              |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at                   | timestamptz | 修改时间                                        | optional     |                                                              |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at                   | timestamptz | 删除时间                                        | optional     |                                                              |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id                                   |

### 1.19 outage_filing_upload_file 停运报备上传文件表

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| outage_filing_upload_file_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| file_path                    | text        | 文件路径                                        | optional     |                            |
| file_type                    | text        | 文件类型(commitment.承诺书 other.其他)          | optional     |                            |
| is_deleted                   | boolean     | 是否删除                                        | optional     |                            |
| created_at                   | timestamptz | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 1.20 vehicle_reserve_history_record 车辆预备库历史记录表

| Name                              | Type        | Description                                     | **Required** | default                    |
| --------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                | bigint      | 按指定方法生成                                  | required     | 主键                       |
| vehicle_reserve_history_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                        | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| operation                         | integer     | 操作 1.加入预备库  2.移出预备库                 | optional     |                            |
| operation_user                    | text        | 操作用户                                        | optional     | **system_user表**的user_id |
| operation_time                    | timestamptz | 操作时间                                        | optional     |                            |
| operation_source                  | integer     | 操作来源 1.车辆  2.驾驶员                       | optional     |                            |
| created_at                        | timestamptz | 创建时间                                        | required     |                            |
| created_by                        | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                        | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                        | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                        | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 1.21 vehicle_security_check_record 车辆安全检查记录表

| Name                             | Type        | Description                                     | **Required** | default                              |
| -------------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                               | bigint      | 按指定方法生成                                  | required     | 主键                                 |
| vehicle_security_check_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| vehicle_id                       | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                                      |
| enterprise_id                    | text        | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id |
| brake                            | integer     | 刹车                                            | optional     |                                      |
| tire                             | integer     | 轮胎                                            | optional     |                                      |
| screw                            | integer     | 螺丝                                            | optional     |                                      |
| hydraulic_oil                    | integer     | 液压油                                          | optional     |                                      |
| engine_oil                       | integer     | 机油                                            | optional     |                                      |
| water                            | integer     | 水                                              | optional     |                                      |
| headlight                        | integer     | 大灯                                            | optional     |                                      |
| taillight                        | integer     | 尾灯                                            | optional     |                                      |
| turn_signal                      | integer     | 转向灯                                          | optional     |                                      |
| brake_light                      | integer     | 刹车灯                                          | optional     |                                      |
| last_check_time                  | timestamptz | 最后检查时间                                    | optional     |                                      |
| created_at                       | timestamptz | 创建时间                                        | required     |                                      |
| created_by                       | text        | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at                       | timestamptz | 修改时间                                        | optional     |                                      |
| updated_by                       | text        | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at                       | timestamptz | 删除时间                                        | optional     |                                      |
| deleted_by                       | text        | 删除人                                          | optional     | **system_user表**的user_id           |

### 1.22 catalog_new_vehicle_record 目录新型车辆记录表

| Name                          | Type        | Description                                     | **Required** | default                    |
| ----------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| catalog_new_vehicle_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                    | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| terminal_id                   | text        | 终端ID                                          | optional     |                            |
| reviewer                      | text        | 审核人                                          | optional     | **system_user表**的user_id |
| is_deleted                    | boolean     | 是否删除                                        | optional     |                            |
| created_at                    | timestamptz | 创建时间                                        | required     |                            |
| created_by                    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                    | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 1.23 vehicle_exit_catalog_review 车辆退出目录库审核表

| Name                           | Type        | Description                                                  | **Required** | default                              |
| ------------------------------ | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------------ |
| id                             | bigint      | 按指定方法生成                                               | required     | 主键                                 |
| vehicle_exit_catalog_review_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                             |
| vehicle_id                     | text        | **vehicle_info** 车辆信息表 的vehicle_id                     | required     |                                      |
| enterprise_id                  | text        | 所在企业id                                                   | optional     | **enterprise_info表**的enterprise_id |
| remarks                        | text        | 备注                                                         | optional     |                                      |
| review_status                  | integer     | 审核状态  0.未完成 1.完成                                    | optional     |                                      |
| exit_type                      | integer     | 类别  3.企业车辆退出目录库  6.零散车辆退出目录库             | optional     |                                      |
| exit_step                      | integer     | 退出步骤  Type=3:{1.AA市A县a企业提交材料,2.A县交警同意,3.A县渣土办同意}  Type=6:{1.A县交警提交材料,2.A县渣土办同意} | optional     |                                      |
| area_id                        | text        | 地区ID                                                       | optional     |                                      |
| original_enterprise_id         | text        | 原单位ID                                                     | optional     | **enterprise_info表**的enterprise_id |
| original_enterprise_name       | text        | 原单位名称                                                   | optional     |                                      |
| original_self_number           | text        | 原自编号                                                     | optional     |                                      |
| is_deleted                     | boolean     | 是否删除                                                     | optional     |                                      |
| created_at                     | timestamptz | 创建时间                                                     | required     |                                      |
| created_by                     | text        | 创建人                                                       | required     | **system_user表**的user_id           |
| updated_at                     | timestamptz | 修改时间                                                     | optional     |                                      |
| updated_by                     | text        | 修改人                                                       | optional     | **system_user表**的user_id           |
| deleted_at                     | timestamptz | 删除时间                                                     | optional     |                                      |
| deleted_by                     | text        | 删除人                                                       | optional     | **system_user表**的user_id           |

### 1.24 vehicle_exit_catalog_log 车辆退出目录库日志表

| Name                           | Type        | Description                                                  | **Required** | default                                                      |
| ------------------------------ | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------------ |
| id                             | bigint      | 按指定方法生成                                               | required     | 主键                                                         |
| vehicle_exit_catalog_log_id    | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                                                     |
| vehicle_exit_catalog_review_id | text        | 车辆退出目录库审核表id                                       |              | **vehicle_exit_catalog_review** 车辆退出目录库审核表的vehicle_exit_catalog_review_id |
| exit_type                      | integer     | 类别  3.企业车辆退出目录库  6.零散车辆退出目录库             | optional     |                                                              |
| exit_step                      | integer     | 退出步骤  Type=3:{1.AA市A县a企业提交材料,2.A县交警同意,3.A县渣土办同意}  Type=6:{1.A县交警提交材料,2.A县渣土办同意} | optional     |                                                              |
| operator                       | text        | 操作人                                                       | optional     |                                                              |
| review_status                  | integer     | 审核状态  0.申请  1.审批  2.退回  3.撤销                     | optional     |                                                              |
| review_user_group              | integer     | 审批用户组  1.运输企业  2.管理部门                           | optional     |                                                              |
| created_at                     | timestamptz | 创建时间                                                     | required     |                                                              |
| created_by                     | text        | 创建人                                                       | required     | **system_user表**的user_id                                   |
| updated_at                     | timestamptz | 修改时间                                                     | optional     |                                                              |
| updated_by                     | text        | 修改人                                                       | optional     | **system_user表**的user_id                                   |
| deleted_at                     | timestamptz | 删除时间                                                     | optional     |                                                              |
| deleted_by                     | text        | 删除人                                                       | optional     | **system_user表**的user_id                                   |

### 1.25  province_upload_vehicle 省厅新平台上传车辆表

| Name                                  | Type        | Description                                     | **Required** | default                              |
| ------------------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                                    | bigint      | 按指定方法生成                                  | required     | 主键                                 |
| province_upload_vehicle_id            | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| license_plate_number                  | text        | 车牌号码                                        | optional     |                                      |
| license_plate_color                   | integer     | 车牌颜色                                        | optional     | **车牌颜色**字典                     |
| enterprise_id                         | text        | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id |
| operator                              | text        | 运营商                                          | optional     | **enterprise_info表**的enterprise_id |
| vehicle_registration_place            | text        | 车籍地(行驶证上的车辆注册地 )                   | optional     |                                      |
| vehicle_type                          | integer     | 车辆类型                                        | optional     | **车辆类型**字典                     |
| heavy                                 | numeric     | 吨位                                            | optional     |                                      |
| seats                                 | integer     | 座位                                            | optional     |                                      |
| operating_flag                        | text        | 运营标识                                        | optional     |                                      |
| operating_state                       | integer     | 营运状态                                        | optional     | **营运状态**字典                     |
| remarks                               | text        | 备注                                            | optional     |                                      |
| online_status                         | text        | 在线状态                                        | optional     |                                      |
| last_report_time                      | timestamptz | 最后汇报时间                                    | optional     |                                      |
| transport_agency_synchronization_flag | text        | 运政同步标识                                    | optional     |                                      |
| is_transport_agency_synchronization   | boolean     | 是否运政同步                                    | optional     |                                      |
| transport_agency_synchronization_time | timestamptz | 运政同步时间                                    | optional     |                                      |
| device_model                          | text        | 设备型号                                        | optional     |                                      |
| terminal_model                        | text        | 终端型号                                        | optional     |                                      |
| created_at                            | timestamptz | 创建时间                                        | required     |                                      |
| created_by                            | text        | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at                            | timestamptz | 修改时间                                        | optional     |                                      |
| updated_by                            | text        | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at                            | timestamptz | 删除时间                                        | optional     |                                      |
| deleted_by                            | text        | 删除人                                          | optional     | **system_user表**的user_id           |

# 2. 驾驶员信息管理

### 2.1 driver_info 驾驶员信息表

| Name               | Type        | Description                                     | **Required** | default                                                      |
| ------------------ | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                 | bigint      | 按指定方法生成                                  | required     | 主键                                                         |
| driver_id          | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| enterprise_id      | text        | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id                         |
| department_id      | text        | 所在部门id                                      | optional     | **department 部门信息表**                                    |
| driver_identity_id | text        | 驾驶员身份验证信息ID                            | optional     | **driver_identity表**的driver_identity_id                    |
| driver_name        | text        | 驾驶员姓名                                      | optional     |                                                              |
| telephone          | text        | 手机号码                                        | optional     |                                                              |
| sex                | integer     | 性别                                            | optional     | **性别**字典                                                 |
| files_number       | text        | 档案编号(后6位)                                 | optional     |                                                              |
| contact_address    | text        | 联系地址                                        | optional     |                                                              |
| mailing_address    | text        | 邮寄地址                                        | optional     |                                                              |
| is_submit          | boolean     | 是否提交                                        | optional     | 用于标志驾驶员资料是否处于确定状态。未确定状态的驾驶员信息在系统上除驾驶员管理外的功能中都查不到。 |
| submit_content     | text        | 提交内容                                        | optional     |                                                              |
| submit_at          | timestamptz | 提交时间                                        | optional     |                                                              |
| submit_by          | text        | 提交人                                          | optional     | system_user表的user_id                                       |
| is_manual_input    | boolean     | 是否手动录入                                    | optional     | 驾驶员资料分为使用身份证读卡器读取身份证自动录入资料和手动填写资料 |
| is_input           | boolean     | 是否录入                                        | optional     |                                                              |
| input_at           | timestamptz | 录入时间                                        | optional     |                                                              |
| input_by           | text        | 录入人                                          | optional     | **system_user表**的user_id                                   |
| is_check_data      | boolean     | 是否校验数据                                    | optional     | 该字段代表是否用于校验驾驶员信息，未正式录入系统，但会同步到公安内容，用于查询驾驶员的违章数据。 |
| check_at           | timestamptz | 检验时间                                        | optional     |                                                              |
| check_by           | text        | 校验人                                          | optional     | **system_user表**的user_id                                   |
| remark_in          | text        | 驾驶员信息同步内网反馈信息                      | optional     | 驾驶员信息同步内网反馈信息。驾驶员信息同步到公安内网后内网的反馈内容，如档案编号填写错误会反馈档案编号后六位不正确 |
| update_time_in     | timestamptz | 内网更新时间                                    | optional     |                                                              |
| is_check_sms       | boolean     | 是否通过短信验证                                | optional     |                                                              |
| remarks            | text        | 备注                                            | optional     |                                                              |
| is_deleted         | boolean     | 是否删除                                        | optional     |                                                              |
| agent              | text        | 代理商                                          | optional     | **enterprise_info表**的enterprise_id                         |
| operator           | text        | 运营商                                          | optional     | **enterprise_info表**的enterprise_id                         |
| is_blacklist       | boolean     | 是否黑名单                                      | optional     | false                                                        |
| blacklist_deadline | timestamptz | 黑名单截止日期                                  | optional     |                                                              |
| created_at         | timestamptz | 创建时间                                        | required     |                                                              |
| created_by         | text        | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at         | timestamptz | 修改时间                                        | optional     |                                                              |
| updated_by         | text        | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at         | timestamptz | 删除时间                                        | optional     |                                                              |
| deleted_by         | text        | 删除人                                          | optional     | **system_user表**的user_id                                   |

### 2.2 driver_identity 驾驶员身份验证信息(各种证件信息，验证状态)

| Name                             | Type        | Description                                     | **Required** | default                    |
| -------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                               | bigint      | 按指定方法生成                                  | required     | 主键                       |
| identity_id                      | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| id_card_num                      | text        | 身份证号码                                      | optional     |                            |
| id_card_birthday                 | timestamptz | 身份证出生日期                                  | optional     |                            |
| id_card_sign_government          | text        | 身份证签发机关                                  | optional     |                            |
| id_card_nation                   | text        | 身份证民族                                      | optional     |                            |
| id_card_start_date               | timestamptz | 身份证有效起始日期                              | optional     |                            |
| id_card_end_date                 | timestamptz | 身份证有效截止日期                              | optional     |                            |
| id_card_front_pic                | text        | 身份证正面照，云存储地址                        | optional     |                            |
| id_card_back_pic                 | text        | 身份证背面照，云存储地址                        | optional     |                            |
| id_card_address                  | text        | 身份证住址                                      | optional     |                            |
| driver_holding_id_photo          | text        | 驾驶员手持身份证照片,云储存系统返回的路径       | optional     |                            |
| driver_photo                     | text        | 驾驶员的正面照,云储存系统返回的路径             | optional     |                            |
| driver_signature                 | text        | 驾驶员签名,云储存系统返回的路径                 | optional     |                            |
| occupational_number              | text        | 从业资格证号码                                  | optional     |                            |
| occupational_expire_date         | timestamptz | 从业资格证有效期至                              | optional     |                            |
| occupational_issuing_authority   | text        | 从业资格证发证机构                              | optional     |                            |
| labor_contract                   | text[]      | 劳动合同,云储存系统返回的完整劳动合同的图片路径 | optional     |                            |
| driver_license_pic               | text        | 驾驶员驾驶证,云储存系统返回的路径               | optional     |                            |
| driver_license_issuing_authority | text        | 驾驶证发证机关                                  | optional     |                            |
| annual_review_date               | timestamptz | 年审日期（六合一）                              | optional     |                            |
| renewal_date                     | timestamptz | 换证日期（六合一）                              | optional     |                            |
| accumulatived_points             | numeric     | 累计积分（六合一）                              | optional     |                            |
| sorting_date                     | timestamptz | 清分日期（六合一）                              | optional     |                            |
| quasi_driving_models             | integer     | 准驾车型（六合一）                              | optional     | **准驾车型**字典           |
| driver_license_province_id       | text        | 驾驶证发证所在地的省份ID                        | optional     | 省份表                     |
| driver_license_city_id           | text        | 驾驶证发证所在地的城市ID                        | optional     | 城市表                     |
| driver_license_district_id       | text        | 驾驶证发证所在地的区域ID                        | optional     | 区域表                     |
| driver_license_status            | integer     | 驾驶证状态                                      | optional     | **驾驶证状态**字典         |
| driver_license_issue_date        | timestamptz | 驾驶证初次领证日期                              | optional     |                            |
| working_time                     | timestamptz | 从业时间                                        | optional     |                            |
| is_review                        | boolean     | 是否审核                                        | optional     |                            |
| created_at                       | timestamptz | 创建时间                                        | required     |                            |
| created_by                       | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                       | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                       | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                       | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                       | text        | 删除人                                          | optional     | **system_user表**的user_id |
| is_deleted                       | boolean     | 是否删除                                        | optional     |                            |

### 2.3 driving_license_registration_inspection 驾驶证年检登记

| Name                     | Type        | Description                                     | **Required** | default                    |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                       |
| driver_id_info_report_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id               | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| ic_card_status           | integer     | IC状态(从业资格证IC卡插入,从业资格证IC卡拔出)   | optional     | **从业资格证IC卡**字典     |
| operation_time           | timestamptz | 操作时间                                        | optional     |                            |
| driver_name              | text        | 驾驶员姓名                                      | optional     |                            |
| license_number           | text        | 证件号码                                        | optional     |                            |
| imel                     | text        | 终端IMEI                                        | optional     | 国际移动设备标识别码       |
| ic_card_reading_result   | text        | IC卡读取结果                                    | optional     | **IC卡读卡**字典           |
| occupational_number      | text        | 从业资格证编码                                  | optional     |                            |
| driver_license_name      | text        | 发证机构名称                                    | optional     |                            |
| license_expire_date      | timestamptz | 证件有效期                                      | optional     |                            |
| registration_time        | timestamptz | 登记时间                                        | optional     |                            |
| created_at               | timestamptz | 创建时间                                        | required     |                            |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz | 修改时间                                        | optional     |                            |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 2.4 driver_fingerprint 驾驶员指纹表

| Name                  | Type        | Description                                     | **Required** | default                    |
| --------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                    | bigint      | 按指定方法生成                                  | required     | 主键                       |
| driver_fingerprint_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| driver_id             | text        | **driver_info驾驶员信息表**的driver_id          | required     |                            |
| driver_name           | text        | 驾驶员姓名                                      | optional     |                            |
| user_id               | text        | 用户ID                                          | optional     | **system_user表**的user_id |
| signature             | text        | 特征码                                          | optional     |                            |
| fingerprint_name      | text        | 指纹名称                                        | optional     |                            |
| sim_number            | text        | SIM卡号                                         | optional     |                            |
| instruction_id        | text        | 指令ID                                          | optional     |                            |
| operation_type        | integer     | 操作类型                                        | optional     |                            |
| content               | text        | 内容                                            | optional     |                            |
| operation_time        | timestamptz | 操作时间                                        | optional     |                            |
| upload_time           | timestamptz | 上传时间                                        | optional     |                            |
| timestamp             | text        | 时间戳                                          | optional     |                            |
| ternimal_id           | text        | 终端ID                                          | optional     |                            |
| is_success            | boolean     | 是否成功                                        | optional     |                            |
| is_deleted            | boolean     | 是否删除                                        | optional     |                            |
| created_at            | timestamptz | 创建时间                                        | required     |                            |
| created_by            | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz | 修改时间                                        | optional     |                            |
| updated_by            | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by            | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 2.5 driver_fingerprint_association 驾驶员指纹关联

| Name                              | Type        | Description                                                | **Required** | default                    |
| --------------------------------- | ----------- | ---------------------------------------------------------- | ------------ | -------------------------- |
| id                                | bigint      | 按指定方法生成                                             | required     | 主键                       |
| driver_fingerprint_association_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用            | required     | 联合主键                   |
| driver_id                         | text        | **driver_info驾驶员信息表**的driver_id                     | required     |                            |
| fingerprint_name                  | text        | 指纹名称                                                   | optional     |                            |
| driver_fingerprint_id             | text        | **driver_fingerprint 驾驶员指纹表**的driver_fingerprint_id | optional     |                            |
| picture_address                   | text        | 图片地址                                                   | optional     |                            |
| is_deleted                        | boolean     | 是否删除                                                   | optional     |                            |
| created_at                        | timestamptz | 创建时间                                                   | required     |                            |
| created_by                        | text        | 创建人                                                     | required     | **system_user表**的user_id |
| updated_at                        | timestamptz | 修改时间                                                   | optional     |                            |
| updated_by                        | text        | 修改人                                                     | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz | 删除时间                                                   | optional     |                            |
| deleted_by                        | text        | 删除人                                                     | optional     | **system_user表**的user_id |

### 2.6 driver_identity_info_report 驾驶员身份信息采集上报

| Name                           | Type        | Description                                     | **Required** | default                    |
| ------------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                             | bigint      | 按指定方法生成                                  | required     | 主键                       |
| driver_identity_info_report_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                     | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| ic_card_status                 | integer     | IC状态(从业资格证IC卡插入,从业资格证IC卡拔出)   | optional     | **从业资格证IC卡**字典     |
| operation_time                 | timestamptz | 操作时间                                        | optional     |                            |
| driver_name                    | text        | 驾驶员姓名                                      | optional     |                            |
| license_number                 | text        | 证件号码                                        | optional     |                            |
| imel                           | text        | 终端IMEI                                        | optional     | 国际移动设备标识别码       |
| ic_card_reading_result         | text        | IC卡读取结果                                    | optional     | **IC卡读卡**字典           |
| occupational_number            | text        | 从业资格证编码                                  | optional     |                            |
| driver_license_name            | text        | 发证机构名称                                    | optional     |                            |
| license_expire_date            | timestamptz | 证件有效期                                      | optional     |                            |
| registration_time              | timestamptz | 登记时间                                        | optional     |                            |
| created_at                     | timestamptz | 创建时间                                        | required     |                            |
| created_by                     | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                     | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                     | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                     | text        | 删除人                                          | optional     | **system_user表**的user_id |

# 3.人车分离管理

### 3.1 vehicle_driver_binding  车辆驾驶员绑定表

| Name                      | Type        | Description                                                  | **Required** | default                        |
| ------------------------- | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------ |
| id                        | bigint      | 按指定方法生成                                               | required     | 主键                           |
| vehicle_driver_binding_id | text        | 车辆驾驶员绑定外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                       |
| driver_id                 | text        | 驾驶员id                                                     | optional     | **driver_info表**的driver_id   |
| vehicle_id                | text        | 车辆id                                                       | optional     | **vehicle_info表**的vehicle_id |
| remarks                   | text        | 备注                                                         | optional     |                                |
| created_at                | timestamptz | 创建时间                                                     | required     |                                |
| created_by                | text        | 创建人                                                       | required     | **system_user表**的user_id     |
| updated_at                | timestamptz | 修改时间                                                     | optional     |                                |
| updated_by                | text        | 修改人                                                       | optional     | **system_user表**的user_id     |
| deleted_at                | timestamptz | 删除时间                                                     | optional     |                                |
| deleted_by                | text        | 删除人                                                       | optional     | **system_user表**的user_id     |
| is_deleted                | boolean     | 是否删除                                                     | optional     |                                |

### 3.2 driving_log_info 行车日志信息

| Name                | Type        | Description                                     | **Required** | default                        |
| ------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------ |
| id                  | bigint      | 按指定方法生成                                  | required     | 主键                           |
| driving_log_info_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                       |
| vehicle_id          | text        | 车辆id                                          | optional     | **vehicle_info表**的vehicle_id |
| driver_id           | text        | 驾驶员id                                        | optional     | **driver_info表**的driver_id   |
| driving_start_time  | timestamptz | 用车起始日期                                    | optional     |                                |
| driving_end_time    | timestamptz | 用车结束日期                                    | optional     |                                |
| cause               | text        | 事由                                            | optional     |                                |
| route               | text        | 路线                                            | optional     |                                |
| remarks             | text        | 备注                                            | optional     |                                |
| start_time          | text        | 开始时间                                        | optional     |                                |
| end_time            | text        | 结束时间                                        | optional     |                                |
| review_status       | integer     | 审核状态                                        | optional     |                                |
| review_agecy_level  | integer     | 审核机构级别                                    | optional     |                                |
| is_markup           | boolean     | 是否补录                                        | optional     |                                |
| created_at          | timestamptz | 创建时间                                        | required     |                                |
| created_by          | text        | 创建人                                          | required     | **system_user表**的user_id     |
| updated_at          | timestamptz | 修改时间                                        | optional     |                                |
| updated_by          | text        | 修改人                                          | optional     | **system_user表**的user_id     |
| deleted_at          | timestamptz | 删除时间                                        | optional     |                                |
| deleted_by          | text        | 删除人                                          | optional     | **system_user表**的user_id     |

### 3.3 case_approval_review_operation 案件审批审核操作

| Name                              | Type        | Description                                     | **Required** | default                    |
| --------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                | bigint      | 按指定方法生成                                  | required     | 主键                       |
| case_approval_review_operation_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| review_opinion                    | text        | 审核意见                                        | optional     |                            |
| review_result                     | text        | 审核结果                                        | optional     |                            |
| reviewer                          | text        | 审核人                                          | optional     |                            |
| review_time                       | timestamptz | 审核时间                                        | optional     |                            |
| created_at                        | timestamptz | 创建时间                                        | required     |                            |
| created_by                        | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                        | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                        | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                        | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 3.4 case_approval_review_call 案件审核审批电话告知

| Name                              | Type        | Description                                                  | **Required** | default                    |
| --------------------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                                | bigint      | 按指定方法生成                                               | required     | 主键                       |
| case_approval_review_call_id      | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| case_approval_review_operation_id | text        | **case_approval_review_operation 案件审批审核操作**的case_approval_review_operation_id | optional     |                            |
| telephone_number                  | text        | 电话号码                                                     | optional     |                            |
| dial_time                         | timestamptz | 拨打时间                                                     | optional     |                            |
| is_connected                      | boolean     | 是否接通                                                     | optional     |                            |
| inform_content                    | text        | 告知内容                                                     | optional     |                            |
| reviewer                          | text        | 审核人                                                       | optional     |                            |
| review_time                       | timestamptz | 审核时间                                                     | optional     |                            |
| created_at                        | timestamptz | 创建时间                                                     | required     |                            |
| created_by                        | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                        | timestamptz | 修改时间                                                     | optional     |                            |
| updated_by                        | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz | 删除时间                                                     | optional     |                            |
| deleted_by                        | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### 3.5 dispute_violation_record 违章争议记录表

| Name                          | Type        | Description                                                  | **Required** | default                                                |
| ----------------------------- | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------ |
| id                            | bigint      | 按指定方法生成                                               | required     | 主键                                                   |
| dispute_violation_id          | text        | 违章争议记录表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                               |
| violation_detail_id           | text        | 违章明细表id                                                 | optional     | **vehicle_violation _details** 表的violation_detail_id |
| written_application_materials | text        | 书面申请材料                                                 | optional     |                                                        |
| labor_contract                | text        | 劳动合同或租赁合同                                           | optional     |                                                        |
| driving_log                   | text        | 行车日志                                                     | optional     |                                                        |
| witness                       | text        | 证人证言                                                     | optional     |                                                        |
| statement                     | text        | 当事人陈述                                                   | optional     |                                                        |
| pic_evidence                  | text        | 图像证据材料                                                 | optional     |                                                        |
| driver_license                | text        | 行为人驾驶证                                                 | optional     |                                                        |
| driving_license               | text        | 机动车行驶证                                                 | optional     |                                                        |
| id_card                       | text        | 行为人身份证                                                 | optional     |                                                        |
| business_license              | text        | 机动车所有人营业执照                                         | optional     |                                                        |
| organization_code             | text        | 机动车所有人组织机构代码证                                   | optional     |                                                        |
| legal_person_id_number        | text        | 法定代表人身份证                                             | optional     |                                                        |
| agent_id_number               | text        | 委托代理人身份证                                             | optional     |                                                        |
| vehicle_manager_id_card       | text        | 机动车管理人身份证                                           | optional     |                                                        |
| other_evidence                | text[]      | 其他证据材料                                                 | optional     |                                                        |
| approve_state                 | integer     | 审批状态                                                     | optional     | **车辆违法审批状态**字典                               |
| update_time_in                | timestamptz | 内网更新时间                                                 | optional     |                                                        |
| contact_address               | text        | 联系地址                                                     | optional     |                                                        |
| created_at                    | timestamptz | 创建时间                                                     | required     |                                                        |
| created_by                    | text        | 创建人                                                       | required     | **system_user表**的user_id                             |
| updated_at                    | timestamptz | 修改时间                                                     | optional     |                                                        |
| updated_by                    | text        | 修改人                                                       | optional     | **system_user表**的user_id                             |
| deleted_at                    | timestamptz | 删除时间                                                     | optional     |                                                        |
| deleted_by                    | text        | 删除人                                                       | optional     | **system_user表**的user_id                             |
| is_deleted                    | boolean     | 是否删除                                                     | optional     |                                                        |

### 3.6 dispute_violation_record_log 违章争议审批日志表

| Name                     | Type        | Description                                                  | **Required** | default                                            |
| ------------------------ | ----------- | ------------------------------------------------------------ | ------------ | -------------------------------------------------- |
| id                       | bigint      | 按指定方法生成                                               | required     | 主键                                               |
| dispute_violation_log_id | text        | 违章争议审批日志外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                           |
| dispute_violation_id     | text        | 违章争议记录表id                                             | optional     | **dispute_violation_record**的dispute_violation_id |
| reviewer                 | text        | 审核人                                                       | optional     | **system_user表**的user_id                         |
| review_time              | timestamptz | 审核时间                                                     | optional     |                                                    |
| review_opinion           | text        | 审核意见                                                     | optional     |                                                    |
| review_result            | text        | 审核结果                                                     | optional     |                                                    |
| review_action_name       | text        | 审核动作名称                                                 | optional     |                                                    |
| approver                 | text        | 审批人                                                       | optional     | **system_user表**的user_id                         |
| update_time_in           | text        | 内网更新时间                                                 | optional     |                                                    |

# 4. 动态监管

### 4.1 dynamic_supervision 动态监管抽查主表

| Name                    | Type        | Description                                                  | **Required** | default                        |
| ----------------------- | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------ |
| id                      | bigint      | 按指定方法生成                                               | required     | 主键                           |
| supervision_id          | text        | 动态监管抽查主表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                       |
| spot_check_date         | timestamptz | 抽查日期                                                     | optional     |                                |
| spot_check_total_number | integer     | 抽查总数                                                     | optional     |                                |
| spot_check_number       | integer     | 抽查数量                                                     | optional     |                                |
| spot_check_ratio        | numeric     | 抽查比例                                                     | optional     |                                |
| check_user_id           | text        | 抽查人员                                                     | optional     | 引用**system_user表**的user_id |
| total_number_vehicle    | integer     | 总车辆数                                                     | optional     |                                |
| province_id             | text        | 抽查人员位置的省份ID                                         | optional     | 省份表province_id              |
| city_id                 | text        | 抽查人员位置的城市ID                                         | optional     | 城市表city_id                  |
| district_id             | text        | 抽查人员位置的区域ID                                         | optional     | 区域表district_id              |
| year                    | integer     | 年                                                           | optional     |                                |
| month                   | integer     | 月                                                           | optional     |                                |
| day                     | integer     | 日                                                           | optional     |                                |
| created_at              | timestamptz | 创建时间                                                     | required     |                                |
| created_by              | text        | 创建人                                                       | required     | **system_user表**的user_id     |
| updated_at              | timestamptz | 修改时间                                                     | optional     |                                |
| updated_by              | text        | 修改人                                                       | optional     | **system_user表**的user_id     |
| deleted_at              | timestamptz | 删除时间                                                     | optional     |                                |
| deleted_by              | text        | 删除人                                                       | optional     | **system_user表**的user_id     |
| is_deleted              | boolean     | 是否被删除                                                   | optional     |                                |

### 4.2 dynamic_supervision_detail 动态监管抽查明细表

| Name                           | Type        | Description                                                  | **Required** | default                                                 |
| ------------------------------ | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------- |
| id                             | bigint      | 按指定方法生成                                               | required     | 主键                                                    |
| supervision_detail_id          | text        | 动态监管抽查明细表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                |
| supervision_id                 | text        | 动态监管抽查主表的ID                                         | optional     | 动态监管抽查主表**dynamic_supervision**的supervision_id |
| driver_id                      | text        | 驾驶员id                                                     | optional     | **driver_info 表**的driver_id                           |
| enterprise_id                  | text        | 车辆所在企业id                                               | optional     | **enterprise_info表**的enterprise_id                    |
| vehicle_id                     | text        | 车辆id                                                       | optional     | **vehicle_info 表**的vehicle_id                         |
| monitoring_time                | timestamptz | 监控平台时间                                                 | optional     |                                                         |
| monitoring_location            | text        | 监控平台显示位置                                             | optional     |                                                         |
| is_online                      | boolean     | 是否在线（是/否）                                            | optional     |                                                         |
| is_tachograph_record_normal    | boolean     | 行车记录仪数据（是否异常）                                   | optional     |                                                         |
| tachograph_data_exception      | text        | 行车记录仪异常数据项                                         | optional     |                                                         |
| tachograph_speed               | text        | 行车记录仪速度                                               | optional     |                                                         |
| gps_speed                      | text        | 卫星定位速度                                                 | optional     |                                                         |
| is_speeding                    | boolean     | 是否超速（是/否）                                            | optional     |                                                         |
| taxi_state                     | integer     | 出租空/重车状态（空/重）                                     | optional     | **出租车空/重车状态**字典                               |
| is_fatigue_driving             | boolean     | 客运疲劳驾驶（是/否）                                        | optional     |                                                         |
| is_morning_outage              | boolean     | 客运凌晨2-5时停运（是/否）                                   | optional     |                                                         |
| curve                          | integer     | 曲线情况（曲线完整/回传异常/零速度）                         | optional     | **曲线情况**字典                                        |
| trail                          | integer     | 轨迹情况（正常/漂移/其他）                                   | optional     | **GPS轨迹情况**字典                                     |
| lens_position                  | integer     | 镜头位置（正/偏）                                            | optional     | **镜头位置**字典                                        |
| equipment                      | integer     | 设备情况（图像正常/无图像/摄像头 号损坏）                    | optional     | **设备情况**字典                                        |
| other_infraction               | text        | 其他违规                                                     | optional     |                                                         |
| disposal_measures              | text        | 处置措施                                                     | optional     |                                                         |
| disposal_results               | text        | 处置结果                                                     | optional     |                                                         |
| treatment_time                 | timestamptz | 处置时间                                                     | optional     |                                                         |
| assignee                       | text        | 受理人                                                       | optional     |                                                         |
| feedback_time                  | timestamptz | 反馈时间                                                     | optional     |                                                         |
| remarks                        | text        | 备注                                                         | optional     |                                                         |
| others                         | text        | 轨迹其他情况                                                 | optional     |                                                         |
| lens_on                        | text        | 摄像头损坏号                                                 | optional     |                                                         |
| monitor_end_time               | timestamptz | 监管费到期时间                                               | optional     |                                                         |
| is_locate                      | boolean     | 是否定位                                                     | optional     |                                                         |
| coordinate                     | point       | 空间数据类型point表示经纬度                                  | optional     |                                                         |
| latitude_longitude_description | text        | 经纬度描述                                                   | optional     |                                                         |
| is_send                        | boolean     | 是否发送                                                     | optional     |                                                         |
| business_scope                 | integer     | 经营范围                                                     | optional     | **经营范围**字典                                        |
| outage_alarm_time              | timestamptz | 凌晨2点到5点停运报警时间                                     | optional     |                                                         |
| speed_alarm_time               | timestamptz | 超速报警时间                                                 | optional     |                                                         |
| speeding_speed                 | text        | 超速速度                                                     | optional     |                                                         |
| fatigue_alarm_time             | timestamptz | 疲劳驾驶报警时间                                             | optional     |                                                         |
| disposal_measures1             | text        | 是否在线处置措施                                             | optional     |                                                         |
| disposal_measures2             | text        | 是否超速处置措施                                             | optional     |                                                         |
| disposal_measures3             | text        | 曲线情况处置措施                                             | optional     |                                                         |
| disposal_measures4             | text        | 客运疲劳驾驶处置措施                                         | optional     |                                                         |
| disposal_measures5             | text        | 客运凌晨停运处置措施                                         | optional     |                                                         |
| disposal_measures6             | text        | 行车记录仪数据处置措施                                       | optional     |                                                         |
| disposal_measures7             | text        | 轨迹情况处置措施                                             | optional     |                                                         |
| disposal_results1              | text        | 是否在线处置结果                                             | optional     |                                                         |
| disposal_results2              | text        | 是否超速处置结果                                             | optional     |                                                         |
| disposal_results3              | text        | 曲线情况处置结果                                             | optional     |                                                         |
| disposal_results4              | text        | 客运疲劳驾驶处置结果                                         | optional     |                                                         |
| disposal_results5              | text        | 客运疲劳驾驶处置结果                                         | optional     |                                                         |
| disposal_results6              | text        | 行车记录仪数据处置结果                                       | optional     |                                                         |
| disposal_results7              | text        | 轨迹情况处置结果                                             | optional     |                                                         |
| is_deleted                     | boolean     | 是否被删除                                                   | optional     |                                                         |
| created_at                     | timestamptz | 创建时间                                                     | required     |                                                         |
| created_by                     | text        | 创建人                                                       | required     | **system_user表**的user_id                              |
| updated_at                     | timestamptz | 修改时间                                                     | optional     |                                                         |
| updated_by                     | text        | 修改人                                                       | optional     | **system_user表**的user_id                              |
| deleted_at                     | timestamptz | 删除时间                                                     | optional     |                                                         |
| deleted_by                     | text        | 删除人                                                       | optional     | **system_user表**的user_id                              |

### 4.3  dynamic_spot_check_disposal 动态抽查处置表

| Name                             | Type        | Description                              | **Required** | default                                                      |
| -------------------------------- | ----------- | ---------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                               | bigint      | 按指定方法生成                           | required     | 主键                                                         |
| vehicle_id                       | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                                                              |
| enterprise_id                    | text        | 所在企业id                               | optional     | **enterprise_info表**的enterprise_id                         |
| telephone                        | text        | 手机号码                                 | optional     |                                                              |
| send_time                        | timestamptz | 发送时间                                 | optional     |                                                              |
| supervision_detail_id            | text        | 动态监管抽查明细表id                     | optional     | **dynamic_supervision_ detail** 动态监管抽查明细表的supervision_detail_id |
| image_abnormal_handing_measure   | text        | 图像异常处置措施                         | optional     |                                                              |
| feedback_time                    | timestamptz | 反馈时间                                 | optional     |                                                              |
| tachograph_data_disposal_measure | text        | 行车记录仪数据处置措施                   | optional     |                                                              |
| operation_user                   | text        | 操作用户                                 | optional     | **system_user表**的user_id                                   |
| is_sms_push                      | boolean     | 是否短信推送                             | optional     |                                                              |
| is_notify                        | boolean     | 是否通报                                 | optional     |                                                              |
| is_announce                      | boolean     | 是否语音通知                             | optional     |                                                              |
| is_app_push                      | boolean     | 是否APP推送                              | optional     |                                                              |
| notify_content                   | text        | 通报内容                                 | optional     |                                                              |
| announce_content                 | text        | 语音内容                                 | optional     |                                                              |
| app_push_content                 | text        | APP推送内容                              | optional     |                                                              |
| disposal_content                 | text        | 处置内容                                 | optional     |                                                              |
| disposal_method                  | integer     | 处置方式                                 | optional     | **处置方式**字典                                             |
| disposal_result                  | text        | 处置结果                                 | optional     |                                                              |
| is_deleted                       | boolean     | 是否删除                                 | optional     | false                                                        |
| created_at                       | timestamptz | 创建时间                                 | required     |                                                              |
| created_by                       | text        | 创建人                                   | required     | **system_user表**的user_id                                   |
| updated_at                       | timestamptz | 修改时间                                 | optional     |                                                              |
| updated_by                       | text        | 修改人                                   | optional     | **system_user表**的user_id                                   |
| deleted_at                       | timestamptz | 删除时间                                 | optional     |                                                              |
| deleted_by                       | text        | 删除人                                   | optional     | **system_user表**的user_id                                   |

### 4.4 muck_truck_online 渣土车在线

| Name                       | Type        | Description                                     | **Required** | default                    |
| -------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                  | required     | 主键                       |
| muck_truck_online_id       | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                 | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| is_online                  | boolean     | 是否在线                                        | optional     |                            |
| is_pay                     | boolean     | 是否缴费                                        | optional     |                            |
| is_gps_online              | boolean     | 是否GPS在线                                     | optional     |                            |
| is_catalog_library         | boolean     | 是否目录库                                      | optional     |                            |
| is_install                 | boolean     | 是否安装                                        | optional     |                            |
| is_vehicle_pass_inspection | boolean     | 是否验车通过                                    | optional     |                            |
| is_new_vehicle             | boolean     | 是否新型渣土车                                  | optional     |                            |
| location_time              | timestamptz | 定位时间                                        | optional     |                            |
| self_number                | text        | 自编号                                          | optional     |                            |
| area                       | text        | 地区                                            | optional     |                            |
| created_at                 | timestamptz | 创建时间                                        | required     |                            |
| created_by                 | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                 | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                 | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 4.5 vehicle_offline_disposal 离线车辆处置表

| Name                        | Type        | Description                                     | **Required** | default                                                      |
| --------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                          | bigint      | 按指定方法生成                                  | required     | 主键                                                         |
| vehicle_offline_disposal_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| enterprise_id               | text        | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id                         |
| telephone                   | text        | 手机号码                                        | optional     |                                                              |
| content                     | text        | 内容                                            | optional     |                                                              |
| send_time                   | timestamptz | 发送时间                                        | optional     |                                                              |
| user_id                     | text        | 用户ID                                          | optional     | **system_user表**的user_id                                   |
| is_sms_push                 | boolean     | 是否短信推送                                    | optional     |                                                              |
| is_report                   | boolean     | 是否通报                                        | optional     |                                                              |
| is_voice_notification       | boolean     | 是否语音通知                                    | optional     |                                                              |
| is_app_push                 | boolean     | 是否APP推送                                     | optional     |                                                              |
| notification_content        | text        | 通报内容                                        | optional     |                                                              |
| voice_content               | text        | 语音内容                                        | optional     |                                                              |
| app_push_content            | text        | APP推送内容                                     | optional     |                                                              |
| supervision_detail_id       | text        | 抽查表ID                                        | optional     | **dynamic_supervision_detail 动态监管抽查明细表**的supervision_detail_id |
| disposal_method             | text        | 处置方式                                        | optional     |                                                              |
| is_deleted                  | boolean     | 是否删除                                        | optional     |                                                              |
| created_at                  | timestamptz | 创建时间                                        | required     |                                                              |
| created_by                  | text        | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at                  | timestamptz | 修改时间                                        | optional     |                                                              |
| updated_by                  | text        | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at                  | timestamptz | 删除时间                                        | optional     |                                                              |
| deleted_by                  | text        | 删除人                                          | optional     | **system_user表**的user_id                                   |

### 4.6 vehicle_alarm_supervision  车辆报警监管

| Name                                 | Type        | Description                                     | **Required** | default                                                      |
| ------------------------------------ | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                                   | bigint      | 按指定方法生成                                  | required     | 主键                                                         |
| vehicle_alarm_supervision_id         | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| superivision_authority_id            | text        | 监管单位id                                      | optional     |                                                              |
| superivision_type                    | text        | 监管类型                                        | optional     | 地区机构监管，企业监管，执法机构监管，执法人员监管，地区抽查机构监管 |
| area_id                              | text        | 区域ID                                          | optional     |                                                              |
| vehicle_alarm_number                 | integer     | 报警车辆数                                      | optional     |                                                              |
| alarm_type                           | text        | 报警类型                                        | optional     |                                                              |
| alarm_times                          | integer     | 报警次数                                        | optional     |                                                              |
| disposal_times                       | integer     | 处置次数                                        | optional     |                                                              |
| alarm_disposal_rate                  | numeric     | 报警处置率                                      | optional     |                                                              |
| district_rectification               | text        | 县级提出的整改                                  | optional     |                                                              |
| city_rectification                   | text        | 市级提出的整改                                  | optional     |                                                              |
| province_rectification               | text        | 省级提出的整改                                  | optional     |                                                              |
| country_rectification                | text        | 部级提出的整改                                  | optional     |                                                              |
| registration_time                    | timestamptz | 登记时间                                        | optional     |                                                              |
| statistics_date                      | text        | 统计日期                                        | optional     |                                                              |
| should_supervision_enterprise_number | integer     | 应监管企业数                                    | optional     |                                                              |
| actual_supervision_enterprise_number | integer     | 实监管企业数                                    | optional     |                                                              |
| supervision_rate                     | numeric     | 监管率                                          | optional     |                                                              |
| business_scope                       | integer     | 经营范围                                        | optional     | **经营范围**字典                                             |
| should_check_vehicle_number          | integer     | 应检查车辆数                                    | optional     |                                                              |
| actual_check_vechicle_number         | integer     | 实检查车辆数                                    | optional     |                                                              |
| check_abnormal_vehicle_number        | integer     | 检查异常车辆数                                  | optional     |                                                              |
| abnormal_disposal_number             | integer     | 异常处置数                                      | optional     |                                                              |
| abnormal_disposal_rate               | numeric     | 异常处置率                                      | optional     |                                                              |
| created_at                           | timestamptz | 创建时间                                        | required     |                                                              |
| created_by                           | text        | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at                           | timestamptz | 修改时间                                        | optional     |                                                              |
| updated_by                           | text        | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at                           | timestamptz | 删除时间                                        | optional     |                                                              |
| deleted_by                           | text        | 删除人                                          | optional     | **system_user表**的user_id                                   |

### 4.7 region_management 进出区域报警--区域管理

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| region_management_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| region_id            | text        | 区域ID                                          | optional     |                            |
| region_name          | text        | 区域名称                                        | optional     |                            |
| region_type          | text        | 区域类型                                        | optional     |                            |
| region_nature        | text        | 区域性质                                        | optional     |                            |
| alarm_begin_time     | timestamptz | 报警开始时间                                    | optional     |                            |
| alarm_end_time       | timestamptz | 报警截止时间                                    | optional     |                            |
| enterprise_name      | text        | 企业名称                                        | optional     |                            |
| duration_            | integer     | 持续时间                                        | optional     |                            |
| max_speed            | integer     | 最高速度                                        | optional     |                            |
| circle_radius        | integer     | 圆形半径                                        | optional     |                            |
| region_coordinate    | point       | 区域经纬度                                      | optional     |                            |
| is_super_region      | boolean     | 是否超级区域                                    | optional     |                            |
| is_deleted           | boolean     | 是否删除                                        | optional     |                            |
| created_at           | timestamptz | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 4.8 region_issued 进出区域报警--区域下发

| Name             | Type        | Description                                     | **Required** | default                    |
| ---------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id               | bigint      | 按指定方法生成                                  | required     | 主键                       |
| region_issued_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| region_name      | text        | 区域名称                                        | optional     |                            |
| imel             | text        | 终端IMEI                                        | optional     |                            |
| created_at       | timestamptz | 创建时间                                        | required     |                            |
| created_by       | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at       | timestamptz | 修改时间                                        | optional     |                            |
| updated_by       | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at       | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by       | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 4.9 vehicle_alarm_times_record 同一车辆报警次数记录表

| Name                          | Type        | Description                                     | **Required** | default                    |
| ----------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| vehicle_alarm_times_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                    | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| alarm_type                    | text        | 报警类型                                        | optional     | **报警类型**字典           |
| disposal_measure              | text        | 处置措施                                        | optional     |                            |
| disposal_time                 | timestamptz | 处置时间                                        | optional     |                            |
| disposal_result               | text        | 处置结果                                        | optional     |                            |
| is_disposal                   | boolean     | 是否处置                                        | optional     |                            |
| disposal_method               | integer     | 处置方式                                        | optional     | **处置方式**字典           |
| duty_person                   | text        | 值班人                                          | optional     |                            |
| alarm_times                   | text        | 报警次数                                        | optional     |                            |
| remarks                       | text        | 备注                                            | optional     |                            |
| record_time                   | timestamptz | 记录时间                                        | optional     |                            |
| created_at                    | timestamptz | 创建时间                                        | required     |                            |
| created_by                    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                    | text        | 删除人                                          | optional     | **system_user表**的user_id |

# 5. 报警管理

### 5.1 vehicle_alarm_data  报警数据表

| Name                          | Type        | Description                                             | **Required** | default                            |
| ----------------------------- | ----------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id                            | bigint      | 按指定方法生成                                          | required     | 主键                               |
| vehicle_alarm_data_id         | text        | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id                    | text        | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| alarm_type                    | text        | 报警类型                                                | optional     | **报警类型**字典                   |
| alarm_start_time              | timestamptz | 报警开始时间                                            | optional     |                                    |
| alarm_end_time                | timestamptz | 报警结束时间                                            | optional     |                                    |
| alarm_end_position            | text        | 报警结束位置                                            | optional     |                                    |
| latest_alarm_time             | timestamptz | 最新报警时间                                            | optional     |                                    |
| latest_alarm_position         | integer     | 最新报警位置                                            | optional     |                                    |
| is_alarm_effective            | boolean     | 报警是否有效                                            | optional     |                                    |
| is_alarm_over                 | boolean     | 报警是否结束                                            | optional     |                                    |
| is_cancel_alarm               | boolean     | 是否取消报警                                            | optional     |                                    |
| alarm_source                  | text        | 报警来源:1.终端报警 2.平台报警                          | optional     | **报警来源**字典                   |
| processing_time               | timestamptz | 处理时间                                                | optional     |                                    |
| processing_method             | text        | 处理方式                                                | optional     | **处警处理方式**字典               |
| processing_description        | text        | 处理描述                                                | optional     |                                    |
| processor                     | text        | 处理人                                                  | optional     | **system_user表**的user_id         |
| processing_status             | text        | 处理状态                                                | optional     | **处警处理状态**字典               |
| tachograph_speed              | numeric     | 行驶记录仪速度                                          | optional     |                                    |
| GPS_speed                     | numeric     | GPS速度                                                 | optional     |                                    |
| maximum_speed                 | numeric     | 最高速度                                                | optional     |                                    |
| speed_limit_threshold         | numeric     | 限速阀值                                                | optional     |                                    |
| coordinate                    | point       | 空间数据类型point表示经度(longitude)和纬度(latitude)    | optional     |                                    |
| location_description          | text        | 位置描述                                                | optional     |                                    |
| duration                      | text        | 持续时间                                                | optional     |                                    |
| road_grade                    | text        | 道路等级                                                | optional     | **道路等级**字典                   |
| road_name                     | text        | 道路名称                                                | optional     |                                    |
| area_id                       | text        | 进区域ID                                                | optional     |                                    |
| alarm_deal_id                 | text        | 处警ID                                                  | optional     |                                    |
| pid                           | text        | 地区                                                    | optional     |                                    |
| record_time                   | timestamptz | 记录时间                                                | optional     |                                    |
| supervisor                    | text        | 监管人                                                  | optional     | **system_user表**的user_id         |
| is_supervise                  | boolean     | 管理部门是否监管                                        | optional     |                                    |
| supervision_time              | timestamptz | 管理部门监管时间                                        | optional     |                                    |
| supervision_note              | text        | 监管备注                                                | optional     |                                    |
| is_resolve                    | boolean     | 是否解析                                                | optional     |                                    |
| is_construction_site_handle   | boolean     | 工地是否处理                                            | optional     |                                    |
| construction_site_handle_time | timestamptz | 工地处理时间                                            | optional     |                                    |
| created_at                    | timestamptz | 创建时间                                                | required     |                                    |
| created_by                    | text        | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at                    | timestamptz | 修改时间                                                | optional     |                                    |
| updated_by                    | text        | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at                    | timestamptz | 删除时间                                                | optional     |                                    |
| deleted_by                    | text        | 删除人                                                  | optional     | **system_user表**的user_id         |

### 5.2 alarm_processing_record 报警处理记录表

| Name                         | Type        | Description                                                  | **Required** | default                    |
| ---------------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                               | required     | 主键                       |
| alarm_data_id                | text        | **vehicle_alarm_data**报警数据表的alarm_data_id              | required     |                            |
| alarm_supervision_picture_id | bigint      | **alarm_supervision_picture_ upload**报警监管图片上传表的alarm_supervision_picture_id | required     |                            |
| processing_content           | text        | 处理内容                                                     | optional     |                            |
| processing_time              | timestamptz | 处理时间                                                     | optional     |                            |
| processing_type              | integer     | 处理类型  1.超速报警  2.疲劳驾驶  3.工程报警  4.超三天断电报警  5.进出区域报警  6.进出区域报警  7.安检到期报警  11.进出工地报警 | optional     |                            |
| operation_user               | text        | 操作用户                                                     | optional     | **system_user表**的user_id |
| is_sms_push                  | boolean     | 是否短信推送                                                 | optional     |                            |
| is_notify                    | boolean     | 是否通报                                                     | optional     |                            |
| is_announce                  | boolean     | 是否语音通知                                                 | optional     |                            |
| is_app_push                  | boolean     | 是否APP推送                                                  | optional     |                            |
| notify_content               | text        | 通报内容                                                     | optional     |                            |
| announce_content             | text        | 语音内容                                                     | optional     |                            |
| app_push_content             | text        | APP推送内容                                                  | optional     |                            |
| disposal_method              | integer     | 处置方式                                                     | optional     | **处置方式**字典           |
| disposal_result              | text        | 处置结果                                                     | optional     |                            |
| is_deleted                   | boolean     | 是否删除                                                     | optional     | false                      |
| created_at                   | timestamptz | 创建时间                                                     | required     |                            |
| created_by                   | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                   | timestamptz | 修改时间                                                     | optional     |                            |
| updated_by                   | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz | 删除时间                                                     | optional     |                            |
| deleted_by                   | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### 5.3 district_alarm_content_push 	各县市区报警内容推送表

| Name          | Type        | Description                                     | **Required** | default                    |
| ------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_data_id | text        | **vehicle_alarm_data**报警数据表的alarm_data_id | required     |                            |
| alarm_type    | text        | 报警类型                                        | optional     | **报警类型**字典           |
| alarm_content | text        | 内容                                            | optional     |                            |
| province_id   | text        | 省份ID                                          | optional     | **省份表**province_id      |
| city_id       | text        | 城市ID                                          | optional     | **城市表**city_id          |
| district_id   | text        | 区ID                                            | optional     | **区域表**district_id      |
| is_deleted    | boolean     | 是否删除                                        | optional     | false                      |
| created_at    | timestamptz | 创建时间                                        | required     |                            |
| created_by    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at    | timestamptz | 修改时间                                        | optional     |                            |
| updated_by    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at    | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by    | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 5.4 alarm_supervision_picture_upload 报警监管图片上传表

| Name                         | Type        | Description                                                  | **Required** | default                              |
| ---------------------------- | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------------ |
| id                           | bigint      | 按指定方法生成                                               | required     | 主键                                 |
| alarm_supervision_picture_id | text        | 报警监管图片上传外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| vehicle_id                   | text        | 车辆ID                                                       | required     | **vehicle_info**表vehicle_id         |
| driver_id                    | text        | 驾驶员id                                                     | optional     | **driver_info 表**的driver_id        |
| enterprise_id                | text        | 所在企业id                                                   | optional     | **enterprise_info表**的enterprise_id |
| camera_id                    | integer     | 摄像头ID                                                     | optional     | **摄像头ID**字典                     |
| photo_condition              | text        | 拍照条件                                                     | optional     | **拍照条件**字典                     |
| update_time                  | timestamptz | 终端上报时间                                                 | optional     |                                      |
| monitoring_pic_name          | text        | 报警监控图片名称                                             | optional     |                                      |
| monitoring_pic_address       | text        | 报警监控图片地址                                             | optional     |                                      |
| monitoring_pic_upload_time   | timestamptz | 报警监控图片上传时间                                         | optional     |                                      |
| imel                         | text        | 终端IMEI                                                     | optional     | 国际移动设备标识别码                 |
| sim_number                   | text        | SIM卡号                                                      | optional     |                                      |
| created_at                   | timestamptz | 创建时间                                                     | required     |                                      |
| created_by                   | text        | 创建人                                                       | required     | **system_user表**的user_id           |
| updated_at                   | timestamptz | 修改时间                                                     | optional     |                                      |
| updated_by                   | text        | 修改人                                                       | optional     | **system_user表**的user_id           |
| deleted_at                   | timestamptz | 删除时间                                                     | optional     |                                      |
| deleted_by                   | text        | 删除人                                                       | optional     | **system_user表**的user_id           |

### 5.5 video_platform_alarm_type  视频平台报警类型字典

| Name                         | Type        | Description                                     | **Required** | default                                                   |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | --------------------------------------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                                                      |
| video_platform_alarm_type_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                  |
| vehicle_alarm_data_id        | integer     | 报警数据表id                                    | optional     | **vehicle_alarm_data  报警数据表**的vehicle_alarm_data_id |
| alarm_type                   | text        | 报警类型                                        | optional     |                                                           |
| alarm_source                 | text        | 报警来源                                        | optional     |                                                           |
| alarm_classify               | text        | 报警分类                                        | optional     | **Adas报警**字典                                          |
| alarm_code                   | text        | 报警代码                                        | optional     |                                                           |
| is_deleted                   | boolean     | 是否删除                                        | optional     |                                                           |
| created_at                   | timestamptz | 创建时间                                        | required     |                                                           |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id                                |
| updated_at                   | timestamptz | 修改时间                                        | optional     |                                                           |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id                                |
| deleted_at                   | timestamptz | 删除时间                                        | optional     |                                                           |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id                                |

### 5.6 offline_alarm_registration 不在线报警登记

| Name                          | Type        | Description                                     | **Required** | default                    |
| ----------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| offline_alarm_registration_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                    | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| offline_start_time            | timestamptz | 离线开始时间                                    | optional     |                            |
| offline_end_time              | timestamptz | 离线结束时间                                    | optional     |                            |
| registration_user             | text        | 登记用户                                        | optional     | **system_user表**的user_id |
| registration_time             | timestamptz | 登记时间                                        | optional     |                            |
| sms_content                   | text        | 短信发送内容                                    | optional     |                            |
| phone_reminder_content        | text        | 电话提醒内容                                    | optional     |                            |
| sms_send_time                 | timestamptz | 短信发送时间                                    | optional     |                            |
| phone_reminder_time           | timestamptz | 电话提醒时间                                    | optional     |                            |
| offline_reason                | text        | 离线原因                                        | optional     |                            |
| alarm_type                    | integer     | 报警类型                                        | optional     | **报警类型**字典           |
| is_registration               | boolean     | 是否登记                                        | optional     |                            |
| is_end_alarm                  | boolean     | 是否结束报警                                    | optional     |                            |
| is_send_sms                   | boolean     | 是否发送短信                                    | optional     |                            |
| is_need_maintain              | boolean     | 是否需要维护                                    | optional     |                            |
| created_at                    | timestamptz | 创建时间                                        | required     |                            |
| created_by                    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                    | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 5.7 voice_alarm_record 语音报警记录

| Name                  | Type        | Description                                     | **Required** | default                    |
| --------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                    | bigint      | 按指定方法生成                                  | required     | 主键                       |
| vioce_alarm_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id            | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| alarm_time            | timestamptz | 报警时间                                        | optional     |                            |
| alarm_type            | text        | 报警类型                                        | optional     |                            |
| remind_time           | timestamptz | 提醒时间                                        | optional     |                            |
| remind_content        | text        | 提醒内容                                        | optional     |                            |
| input_person          | text        | 录入人                                          | optional     | **system_user表**的user_id |
| input_time            | timestamptz | 录入时间                                        | optional     |                            |
| is_success            | boolean     | 是否成功                                        | optional     |                            |
| created_at            | timestamptz | 创建时间                                        | required     |                            |
| created_by            | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz | 修改时间                                        | optional     |                            |
| updated_by            | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by            | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 5.8 enterprise_alarm_send_police 企业报警发送交警联系人

| Name                            | Type        | Description                                     | **Required** | default                              |
| ------------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                              | bigint      | 按指定方法生成                                  | required     | 主键                                 |
| enterprise_alarm_send_police_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| enterprise_id                   | text        | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id |
| enterprise_name                 | text        | 企业名称                                        | optional     |                                      |
| enterprise_contact              | text        | 企业联系人                                      | optional     |                                      |
| enterprise_phone                | text        | 企业联系电话                                    | optional     |                                      |
| police                          | text        | 交警                                            | optional     |                                      |
| police_phone                    | text        | 交警联系电话                                    | optional     |                                      |
| phlice_department               | text        | 交警所属部门                                    | optional     |                                      |
| created_at                      | timestamptz | 创建时间                                        | required     |                                      |
| created_by                      | text        | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at                      | timestamptz | 修改时间                                        | optional     |                                      |
| updated_by                      | text        | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at                      | timestamptz | 删除时间                                        | optional     |                                      |
| deleted_by                      | text        | 删除人                                          | optional     | **system_user表**的user_id           |

# 6. 交通违法信息管理

### 6.1 vehicle_violation_details 车辆违章明细表

| Name                    | Type        | Description                                                 | **Required** | default                                                      |
| ----------------------- | ----------- | ----------------------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                      | bigint      | 按指定方法生成                                              | required     | 主键                                                         |
| violation_detail_id     | text        | 车辆违章明细外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| vehicle_id              | text        | 违章车辆id                                                  | optional     | **vehicle_info表**的vehicle_id                               |
| driver_id               | text        | 违章驾驶员id                                                | optional     | **driver_info表**的driver_id                                 |
| enterprise_id           | text        | 所在企业id                                                  | optional     | **enterprise_info表**的enterprise_id                         |
| illegal_code            | text        | 违法代码                                                    | optional     | **VIO_CODEWFDM** 违法描述字典表                              |
| illegal_time            | timestamptz | 违法时间                                                    | optional     |                                                              |
| illegal_handling_status | integer     | 违法处理状态                                                | optional     | **车辆违法处理状态**字典                                     |
| illegal_location        | text        | 违法地点                                                    | optional     |                                                              |
| standard_value          | text        | 标准值                                                      | optional     | 路段的限速阈值或核载的人数，根据违法的种类不同而不同。       |
| measured_value          | text        | 实测值                                                      | optional     | 车辆实际行驶的车速或实际载的人数，根据违法的种类不同而不同。 |
| discovery_agency        | text        | 发现机构                                                    | optional     |                                                              |
| illegal_photo           | text        | 违法照片                                                    | optional     |                                                              |
| is_notice_driver        | boolean     | 是否通知驾驶员                                              | optional     |                                                              |
| notice_time             | timestamptz | 通知时间                                                    | optional     |                                                              |
| decision_number         | text        | 决定书号                                                    | optional     |                                                              |
| payment_mark            | integer     | 缴款标记                                                    | optional     | **是否缴款**字典                                             |
| party_name              | text        | 当事人姓名                                                  | optional     |                                                              |
| information_source      | integer     | 信息来源：1，强制，2，非现场，0，简易                       | optional     | **信息来源**字典表                                           |
| vehicle_information     | text        | 驾驶人处理的交通违法记录对应的机动车信息                    | optional     |                                                              |
| update_time_in          | timestamptz | 内网更新时间                                                | optional     |                                                              |
| is_handle               | boolean     | 是否处理                                                    | optional     |                                                              |
| handle_by               | text        | 处理人                                                      | optional     | **system_user表**的user_id                                   |
| handle_at               | timestamptz | 处理时间                                                    | optional     |                                                              |
| is_send                 | boolean     | 是否发送短信                                                | optional     |                                                              |
| is_deleted              | boolean     | 是否删除                                                    | optional     |                                                              |
| created_at              | timestamptz | 创建时间                                                    | required     |                                                              |
| created_by              | text        | 创建人                                                      | required     | **system_user表**的user_id                                   |
| updated_at              | timestamptz | 修改时间                                                    | optional     |                                                              |
| updated_by              | text        | 修改人                                                      | optional     | **system_user表**的user_id                                   |
| deleted_at              | timestamptz | 删除时间                                                    | optional     |                                                              |
| deleted_by              | text        | 删除人                                                      | optional     | **system_user表**的user_id                                   |

### 6.2 vehicle_violation_scoring_record  车辆违规计分记录

| Name                      | Type        | Description                                                  | **Required** | default                                                      |
| ------------------------- | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------------ |
| id                        | bigint      | 按指定方法生成                                               | required     | 主键                                                         |
| violation_scoring_id      | text        | 车辆违规计分记录外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| vehicle_id                | text        | 扣分车辆id                                                   | optional     | **vehicle_info表**的vehicle_id                               |
| violation_scoring_item_id | text        | 扣分明细id                                                   | optional     | **vehicle_violation_scoring_ items**表的violation_scoring_item_id |
| demerit_points            | numeric     | 扣分分值                                                     | optional     |                                                              |
| remarks                   | text        | 备注                                                         | optional     |                                                              |
| is_deleted                | boolean     | 是否删除                                                     | optional     |                                                              |
| created_at                | timestamptz | 创建时间                                                     | required     |                                                              |
| created_by                | text        | 创建人                                                       | required     | **system_user表**的user_id                                   |
| updated_at                | timestamptz | 修改时间                                                     | optional     |                                                              |
| updated_by                | text        | 修改人                                                       | optional     | **system_user表**的user_id                                   |
| deleted_at                | timestamptz | 删除时间                                                     | optional     |                                                              |
| deleted_by                | text        | 删除人                                                       | optional     | **system_user表**的user_id                                   |

### 6.3 vehicle_violation_scoring_items 车辆违规计分项表

| Name                       | Type        | Description                                                  | **Required** | default                    |
| -------------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                               | required     | 主键                       |
| violation_scoring_item_id  | text        | 车辆违规计分项表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| deduction_item_description | text        | 扣分事项描述                                                 | optional     |                            |
| deduction_category         | integer     | 扣分事项类别                                                 | optional     | **车辆评分扣分类别**字典   |
| demerit_points             | numeric     | 扣分分值                                                     | optional     |                            |
| is_deleted                 | boolean     | 是否删除                                                     | optional     |                            |
| created_at                 | timestamptz | 创建时间                                                     | required     |                            |
| created_by                 | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                 | timestamptz | 修改时间                                                     | optional     |                            |
| updated_by                 | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz | 删除时间                                                     | optional     |                            |
| deleted_by                 | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### 6.4 vio_codewfdm 违法描述字典表(交警提供数据表，暂无修改)

```sh
#交警提供的数据表，交警只提供了部分字段的含义，没用到的字段交警未提供。
```

| Name     | Type    | Description  | **Required** | default |
| -------- | ------- | ------------ | ------------ | ------- |
| WFXW     | text    | 违法行为     | required     |         |
| WFMS     | text    | 违法描述     | optional     |         |
| WFJFS    | numeric | 违法计分数   | optional     |         |
| FKJE_MIN | numeric | 最小罚款金额 | optional     |         |
| FKJE_MAX | numeric | 最大罚款金额 | optional     |         |
| XH       | text    | 序号         | optional     |         |

### 6.5 regional_violation_register  区域处理机关交通违法违规登记表

| Name                           | Type        | Description                                                  | **Required** | default                                            |
| ------------------------------ | ----------- | ------------------------------------------------------------ | ------------ | -------------------------------------------------- |
| id                             | bigint      | 按指定方法生成                                               | required     | 主键                                               |
| regional_violation_register_id | text        | 区域处理机关交通违法违规登记表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                           |
| violation_detail_id            | text        | 违法记录表ID                                                 | optional     | **vehicle_violation_details**的violation_detail_id |
| vehicle_id                     | text        | 违章车辆id                                                   | optional     | **vehicle_info表**的vehicle_id                     |
| driver_id                      | text        | 驾驶员id                                                     | optional     | **driver_info 表**的driver_id                      |
| illegal_code                   | integer     | 违法代码                                                     | optional     | **VIO_CODEWFDM** 违法描述字典表                    |
| illegal_time                   | timestamptz | 违法时间                                                     | optional     |                                                    |
| illegal_type                   | integer     | 类型(1车辆2驾驶员)                                           | optional     | **违法类型**字典表                                 |
| register_time                  | timestamptz | 登记时间                                                     | optional     |                                                    |
| processing_agency              | text        | 处理机关                                                     | optional     |                                                    |
| operator                       | text        | 操作员                                                       | optional     | **system_user表**的user_id                         |
| is_register                    | boolean     | 类型(false未登记true已登记)                                  | optional     | fase                                               |
| created_at                     | timestamptz | 创建时间                                                     | required     |                                                    |
| created_by                     | text        | 创建人                                                       | required     | **system_user表**的user_id                         |
| updated_at                     | timestamptz | 修改时间                                                     | optional     |                                                    |
| updated_by                     | text        | 修改人                                                       | optional     | **system_user表**的user_id                         |
| deleted_at                     | timestamptz | 删除时间                                                     | optional     |                                                    |
| deleted_by                     | text        | 删除人                                                       | optional     | **system_user表**的user_id                         |
| is_deleted                     | boolean     | 是否删除                                                     | optional     |                                                    |

### 6.6 illegal_photo 违法照片表

| Name             | Type        | Description                                               | **Required** | default                    |
| ---------------- | ----------- | --------------------------------------------------------- | ------------ | -------------------------- |
| id               | bigint      | 按指定方法生成                                            | required     | 主键                       |
| illegal_photo_id | text        | 违法照片表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| picture_name     | text        | 违法照片名称                                              | optional     |                            |
| picture_address  | text        | 违法照片地址                                              | optional     |                            |
| is_synchronized  | boolean     | 是否同步                                                  | optional     | false                      |
| is_deleted       | boolean     | 是否删除                                                  | optional     | false                      |
| created_at       | timestamptz | 创建时间                                                  | required     |                            |
| created_by       | text        | 创建人                                                    | required     | **system_user表**的user_id |
| updated_at       | timestamptz | 修改时间                                                  | optional     |                            |
| updated_by       | text        | 修改人                                                    | optional     | **system_user表**的user_id |
| deleted_at       | timestamptz | 删除时间                                                  | optional     |                            |
| deleted_by       | text        | 删除人                                                    | optional     | **system_user表**的user_id |

### 6.7 app_enforcement APP现场执法表

| Name                 | Type        | Description                                                  | **Required** | default                    |
| -------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                               | required     | 主键                       |
| illegal_photo_id     | text        | APP现场执法表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id                     | optional     |                            |
| operation_user       | text        | 操作用户                                                     | optional     | **system_user表**的user_id |
| edit_text            | text        | 编辑文本                                                     | optional     |                            |
| coordinate           | point       | 空间数据类型point表示经度(longitude)和纬度(latitude)         | optional     |                            |
| location_description | text        | 位置描述                                                     | optional     |                            |
| enterprise_type      | text        | 企业类型                                                     | optional     | **企业类型**字典           |
| picket_status        | integer     | 纠察状态（1.反馈辖区管理 2.执法考评 3.其他 4.查处“两非”渣土车） | optional     |                            |
| created_at           | timestamptz | 创建时间                                                     | required     |                            |
| created_by           | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at           | timestamptz | 修改时间                                                     | optional     |                            |
| updated_by           | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz | 删除时间                                                     | optional     |                            |
| deleted_by           | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### 6.8 driver_peccancy_check  驾驶员违法核实记录表

| Name          | Type        | Description    | **Required** | default                              |
| ------------- | ----------- | -------------- | ------------ | ------------------------------------ |
| id            | bigint      | 按指定方法生成 | required     | 主键                                 |
| vehicle_id    | text        | 违章车辆id     | optional     | **vehicle_info表**的vehicle_id       |
| driver_id     | text        | 违章驾驶员id   | optional     | **driver_info表**的driver_id         |
| enterprise_id | text        | 所在企业id     | optional     | **enterprise_info表**的enterprise_id |
| created_at    | timestamptz | 创建时间       | required     |                                      |
| created_by    | text        | 创建人         | required     | **system_user表**的user_id           |
| updated_at    | timestamptz | 修改时间       | optional     |                                      |
| updated_by    | text        | 修改人         | optional     | **system_user表**的user_id           |
| deleted_at    | timestamptz | 删除时间       | optional     |                                      |
| deleted_by    | text        | 删除人         | optional     | **system_user表**的user_id           |

### 6.9 serious_traffic_violation 严重交通违法行为表

| Name                         | Type        | Description                                     | **Required** | default                         |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                            |
| serious_traffic_violation_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                        |
| vehicle_id                   | text        | 违章车辆id                                      | optional     | **vehicle_info表**的vehicle_id  |
| illegal_code                 | text        | 违法代码                                        | optional     | **VIO_CODEWFDM** 违法描述字典表 |
| illegal_time                 | timestamptz | 违法日期                                        | optional     |                                 |
| created_at                   | timestamptz | 创建时间                                        | required     |                                 |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id      |
| updated_at                   | timestamptz | 修改时间                                        | optional     |                                 |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id      |
| deleted_at                   | timestamptz | 删除时间                                        | optional     |                                 |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id      |

### 6.10 deduction_report 扣分报表

| Name                | Type        | Description                                     | **Required** | default                    |
| ------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                  | bigint      | 按指定方法生成                                  | required     | 主键                       |
| deduction_report_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| area_id             | text        | 地区                                            | optional     |                            |
| date                | text        | 年份.月份  如2015.3                             | optional     |                            |
| cause               | text        | 扣分事由                                        | optional     |                            |
| input_person        | text        | 录入人                                          | optional     | **system_user表**的user_id |
| input_time          | timestamptz | 录入时间                                        | optional     |                            |
| demerit_points      | numeric     | 扣分分值                                        | optional     |                            |
| created_at          | timestamptz | 创建时间                                        | required     |                            |
| created_by          | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at          | timestamptz | 修改时间                                        | optional     |                            |
| updated_by          | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at          | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by          | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 6.11 violation_registration 违法信息登记表

| Name                      | Type        | Description                                     | **Required** | default                        |
| ------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------ |
| id                        | bigint      | 按指定方法生成                                  | required     | 主键                           |
| violation_registration_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                       |
| vehicle_id                | text        | 违章车辆id                                      | optional     | **vehicle_info表**的vehicle_id |
| license_plate_number      | text        | 车牌号                                          | optional     |                                |
| id_card_num               | text        | 码身份证号                                      | optional     |                                |
| name                      | text        | 姓名                                            | optional     |                                |
| location                  | text        | 地点                                            | optional     |                                |
| cause                     | text        | 原因                                            | optional     |                                |
| illegal_time              | timestamptz | 违法时间                                        | optional     |                                |
| operator                  | text        | 操作人                                          | optional     | **system_user表**的user_id     |
| illegal_area              | text        | 违法地区                                        | optional     |                                |
| illegal_code              | text        | 违法代码                                        | optional     |                                |
| vehicle_area              | text        | 车辆所属地区                                    | optional     |                                |
| vehicle_enterprise        | text        | 车辆所属单位                                    | optional     |                                |
| province_id               | text        | 所在省                                          | optional     | **省份表**province_id          |
| city_id                   | text        | 所在市                                          | optional     | **城市表**city_id              |
| district_id               | text        | 所在县                                          | optional     | **区域表**district_id          |
| supervisor                | text        | 监管人                                          | optional     |                                |
| supervision_time          | text        | 监管时间                                        | optional     |                                |
| sepervision_remarks       | text        | 监管备注                                        | optional     |                                |
| is_supervised             | text        | 是否监管                                        | optional     |                                |
| is_accident               | text        | 是否事故                                        | optional     |                                |
| is_deleted                | boolean     | 是否删除                                        | optional     |                                |
| created_at                | timestamptz | 创建时间                                        | required     |                                |
| created_by                | text        | 创建人                                          | required     | **system_user表**的user_id     |
| updated_at                | timestamptz | 修改时间                                        | optional     |                                |
| updated_by                | text        | 修改人                                          | optional     | **system_user表**的user_id     |
| deleted_at                | timestamptz | 删除时间                                        | optional     |                                |
| deleted_by                | text        | 删除人                                          | optional     | **system_user表**的user_id     |

# 7. 新型渣土车推荐目录管理

### 7.1 new_muck_truck_recommend_catalog 新型渣土车推荐目录

| Name                                | Type        | Description                                     | **Required** | default                    |
| ----------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                  | bigint      | 按指定方法生成                                  | required     | 主键                       |
| new_muck_truck_recommend_catalog_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| brand_name                          | text        | 品牌简称                                        | optional     |                            |
| created_at                          | timestamptz | 创建时间                                        | required     |                            |
| created_by                          | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                          | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                          | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                          | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                          | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 7.2 vehicle_sale_record 汽车销售备案

| Name                   | Type        | Description                                     | **Required** | default                                                 |
| ---------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------- |
| id                     | bigint      | 按指定方法生成                                  | required     | 主键                                                    |
| vehicle_sale_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                |
| seller                 | text        | 销售商                                          | optional     | **enterprise_info表**的enterprise_id                    |
| vehicle_model          | text        | 汽车型号                                        | optional     |                                                         |
| vehicle_picture        | text[]      | 汽车相关图片的路径                              | optional     |                                                         |
| vehicle_parameter      | text        | 汽车参数                                        | optional     |                                                         |
| reference_price        | text        | 参考报价                                        | optional     |                                                         |
| other_material         | text        | 其他相关材料                                    | optional     |                                                         |
| remarks                | text        | 备注                                            | optional     |                                                         |
| registration_time      | timestamptz | 登记日期                                        | optional     |                                                         |
| registration_user      | text        | 登记用户                                        | optional     | **system_user表**的user_id                              |
| cause                  | text        | 原因                                            | optional     |                                                         |
| catalog_id             | text        | 汽车品牌(简称)                                  | optional     | new_muck_truck_recommend_catalog 新型渣土车推荐目录的id |
| axis_type              | integer     | 轴数类型                                        | optional     | **轴数类型**字典                                        |
| transport_volume       | text        | 运输方量                                        | optional     |                                                         |
| is_review              | boolean     | 是否审核                                        | optional     |                                                         |
| is_deleted             | boolean     | 是否删除                                        | optional     |                                                         |
| created_at             | timestamptz | 创建时间                                        | required     |                                                         |
| created_by             | text        | 创建人                                          | required     | **system_user表**的user_id                              |
| updated_at             | timestamptz | 修改时间                                        | optional     |                                                         |
| updated_by             | text        | 修改人                                          | optional     | **system_user表**的user_id                              |
| deleted_at             | timestamptz | 删除时间                                        | optional     |                                                         |
| deleted_by             | text        | 删除人                                          | optional     | **system_user表**的user_id                              |

### 7.3 muck_truck_sale_order 新型渣土车销售订单

| Name                     | Type        | Description                                     | **Required** | default                                  |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | ---------------------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                                     |
| muck_truck_sale_order_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                 |
| customer_name            | text        | 顾客企业id                                      | optional     |                                          |
| customer_id              | text        | 购车企业或个人姓名                              | optional     | **enterprise_info表**的enterprise_id     |
| id_card_num              | text        | 身份证号                                        | optional     |                                          |
| contact_phone            | text        | 联系电话                                        | optional     |                                          |
| notification             | text        | 告知书                                          | optional     |                                          |
| vehicle_sale_contract    | text[]      | 汽车销售合同(销售合同照片地址)                  | optional     |                                          |
| vehicle_deposit_contract | text[]      | 订金合同(订金合同照片地址)                      | optional     |                                          |
| commitment               | text        | 承诺书                                          | optional     |                                          |
| book_date                | timestamptz | 订车日期                                        | optional     |                                          |
| registration_time        | timestamptz | 登记时间                                        | optional     |                                          |
| belonging_enterprise     | text        | 所属单位                                        | optional     |                                          |
| id_card_picture          | text        | 身份证(身份证照片路径)                          | optional     |                                          |
| progress                 | numeric     | 进度                                            | optional     |                                          |
| vehicle_sale_record_id   | text        | 汽车型号                                        | optional     | **vehicle_sale_record 汽车销售备案**的id |
| book_vehicle_number      | integer     | 订购车辆数                                      | optional     |                                          |
| vehicle_length           | numeric     | 整车长度                                        | optional     |                                          |
| vehicle_width            | numeric     | 整车宽度                                        | optional     |                                          |
| vehicle_height           | numeric     | 整车高度                                        | optional     |                                          |
| tipper_length            | numeric     | 自卸车长度                                      | optional     |                                          |
| tipper_width             | numeric     | 自卸车宽度                                      | optional     |                                          |
| tipper_height            | numeric     | 自卸车高度                                      | optional     |                                          |
| tank_truck_marked_volume | numeric     | 罐车标注容积                                    | optional     |                                          |
| tank_truck_actual_volume | numeric     | 罐车实际搅动容积                                | optional     |                                          |
| carriage_sealing_device  | text        | 车厢密闭装置                                    | optional     |                                          |
| u_shaped_cargo_box       | text        | U型货箱                                         | optional     |                                          |
| is_other_vehicle_model   | boolean     | 是否其他汽车型号                                | optional     |                                          |
| is_completed             | boolean     | 是否完成                                        | optional     |                                          |
| is_deleted               | boolean     | 是否删除                                        | optional     |                                          |
| created_at               | timestamptz | 创建时间                                        | required     |                                          |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id               |
| updated_at               | timestamptz | 修改时间                                        | optional     |                                          |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id               |
| deleted_at               | timestamptz | 删除时间                                        | optional     |                                          |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id               |

### 7.4 muck_truck_sale_order_detail 新型渣土车销售订单明细

| Name                            | Type        | Description                                     | **Required** | default                                                      |
| ------------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                              | bigint      | 按指定方法生成                                  | required     | 主键                                                         |
| muck_truck_sale_order_detail_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| order_id                        | text        | 订单号                                          | optional     | **muck_truck_sale_order 新型渣土车销售订单**的muck_truck_sale_order_id |
| vehicle_identification_number   | text        | 车架号(后6位)                                   | optional     | 车辆识别代号vin,如D02133                                     |
| certificate                     | text        | 合格证                                          | optional     |                                                              |
| terminal_proof                  | text        | 终端证明                                        | optional     |                                                              |
| driver_license_pic              | text        | 行驶证(照片)                                    | optional     |                                                              |
| license_plate_number            | text        | 车牌号码                                        | optional     |                                                              |
| license_plate_color             | integer     | 车牌颜色                                        | optional     | **车牌颜色**字典                                             |
| license_plate_type              | integer     | 号牌种类                                        | optional     | **号牌种类**字典                                             |
| arrive_qz_date                  | timestamptz | 抵达泉州日期                                    | optional     |                                                              |
| first_registration_date         | timestamptz | 初次登记日期                                    | optional     |                                                              |
| registration_date               | timestamptz | 登记时间                                        | optional     |                                                              |
| step                            | integer     | 步骤(销售订单登记进度（2.到车登记 3.上牌登记）) | optional     |                                                              |
| vehicle_photo                   | text        | 车辆照片                                        | optional     |                                                              |
| seller_preview_number           | text        | 供应商销售预编号                                | optional     |                                                              |
| is_deleted                      | boolean     | 是否删除                                        | optional     |                                                              |
| created_at                      | timestamptz | 创建时间                                        | required     |                                                              |
| created_by                      | text        | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at                      | timestamptz | 修改时间                                        | optional     |                                                              |
| updated_by                      | text        | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at                      | timestamptz | 删除时间                                        | optional     |                                                              |
| deleted_by                      | text        | 删除人                                          | optional     | **system_user表**的user_id                                   |

### 7.5 seller_rating_record 销售商评分记录

| Name                    | Type        | Description                                     | **Required** | default                              |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                                 |
| seller_rating_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| demerit_obj             | text        | 扣分对象                                        | optional     | **enterprise_info表**的enterprise_id |
| demerit_points          | numeric     | 扣分分值                                        | optional     |                                      |
| demerit_reason          | text        | 扣分原因                                        | optional     |                                      |
| operator                | text        | 操作人                                          | optional     | **system_user表**的user_id           |
| operation_time          | timestamptz | 操作时间                                        | optional     |                                      |
| is_deleted              | boolean     | 是否删除                                        | optional     |                                      |
| created_at              | timestamptz | 创建时间                                        | required     |                                      |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at              | timestamptz | 修改时间                                        | optional     |                                      |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at              | timestamptz | 删除时间                                        | optional     |                                      |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id           |

### 7.6 seller_filing 销售商备案

| Name                                  | Type        | Description                                     | **Required** | default                                                 |
| ------------------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------- |
| id                                    | bigint      | 按指定方法生成                                  | required     | 主键                                                    |
| seller_filing_id                      | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                |
| seller                                | text        | 销售商                                          | optional     | **enterprise_info表**的enterprise_id                    |
| business_license                      | text        | 工商营业执照                                    | optional     |                                                         |
| organization_code                     | text        | 组织机构代码证                                  | optional     |                                                         |
| legal_representative_id_card          | text        | 法人代表身份证                                  | optional     |                                                         |
| entrusted_agent_id_card               | text        | 委托代理人身份证                                | optional     |                                                         |
| ministry_production_sale_permit       | text        | 工信部汽车生产销售许可证明                      | optional     |                                                         |
| industry_self_discipline_commitment   | text        | 行业自律承诺书                                  | optional     |                                                         |
| manufacturer_production_sale_permit   | text        | 汽车厂家汽车销售许可证明                        | optional     |                                                         |
| repair_service_station_name           | text        | 维修服务服务站名称                              | optional     |                                                         |
| repair_service_station_address        | text        | 维修服务服务站地址                              | optional     |                                                         |
| repair_service_station_material       | text        | 维修服务服务站相关材料                          | optional     |                                                         |
| service_staff_name                    | text        | 服务人员名称                                    | optional     |                                                         |
| service_staff_phone                   | text        | 服务人员电话                                    | optional     |                                                         |
| service_staff_id                      | text        | 服务人员身份证                                  | optional     |                                                         |
| service_staff_material                | text        | 服务人员相关材料                                | optional     |                                                         |
| other_material                        | text        | 其他相关材料                                    | optional     |                                                         |
| remarks                               | text        | 备注                                            | optional     |                                                         |
| registration_time                     | timestamptz | 登记日期                                        | optional     |                                                         |
| registration_user                     | text        | 登记用户                                        | optional     | **system_user表**的user_id                              |
| is_verify                             | boolean     | 是否核查                                        | optional     |                                                         |
| cause                                 | text        | 原因                                            | optional     |                                                         |
| catalog_id                            | text        | 汽车品牌(简称)                                  | optional     | new_muck_truck_recommend_catalog 新型渣土车推荐目录的id |
| platform_docking_technology_agreement | text        | 平台对接技术协议                                | optional     |                                                         |
| is_deleted                            | boolean     | 是否删除                                        | optional     |                                                         |
| created_at                            | timestamptz | 创建时间                                        | required     |                                                         |
| created_by                            | text        | 创建人                                          | required     | **system_user表**的user_id                              |
| updated_at                            | timestamptz | 修改时间                                        | optional     |                                                         |
| updated_by                            | text        | 修改人                                          | optional     | **system_user表**的user_id                              |
| deleted_at                            | timestamptz | 删除时间                                        | optional     |                                                         |
| deleted_by                            | text        | 删除人                                          | optional     | **system_user表**的user_id                              |

### 7.7 new_muck_truck_info  新型渣土车信息表

| Name                   | Type        | Description                                     | **Required** | default                        |
| ---------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------ |
| id                     | bigint      | 按指定方法生成                                  | required     | 主键                           |
| new_muck_truck_info_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                       |
| vehicle_id             | text        | 扣分车辆id                                      | optional     | **vehicle_info表**的vehicle_id |
| license_plate_number   | text        | 车牌号                                          | optional     |                                |
| registration_time      | timestamptz | 登记注册时间                                    | optional     |                                |
| emission_standard      | text        | 排放标准                                        | optional     |                                |
| is_u_shaped_cargo_box  | boolean     | 是否U型货箱                                     | optional     |                                |
| length                 | text        | 长                                              | optional     |                                |
| width                  | text        | 宽                                              | optional     |                                |
| height                 | text        | 高                                              | optional     |                                |
| sealing_device         | text        | 密封设备                                        | optional     |                                |
| top_cover_height       | text        | 顶盖高度                                        | optional     |                                |
| created_at             | timestamptz | 创建时间                                        | required     |                                |
| created_by             | text        | 创建人                                          | required     | **system_user表**的user_id     |
| updated_at             | timestamptz | 修改时间                                        | optional     |                                |
| updated_by             | text        | 修改人                                          | optional     | **system_user表**的user_id     |
| deleted_at             | timestamptz | 删除时间                                        | optional     |                                |
| deleted_by             | text        | 删除人                                          | optional     | **system_user表**的user_id     |

### 7.8 new_muck_truck_photo 新型渣土车拍照图片表

| Name                    | Type        | Description                                     | **Required** | default                        |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------ |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                           |
| new_muck_truck_photo_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                       |
| vehicle_id              | text        | 扣分车辆id                                      | optional     | **vehicle_info表**的vehicle_id |
| file_path               | text        | 路径                                            | optional     |                                |
| upload_user             | text        | 上传用户                                        | optional     | **system_user表**的user_id     |
| created_at              | timestamptz | 创建时间                                        | required     |                                |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id     |
| updated_at              | timestamptz | 修改时间                                        | optional     |                                |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id     |
| deleted_at              | timestamptz | 删除时间                                        | optional     |                                |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id     |

### 7.9 muck_truck_purchase_intention 渣土车购车意向

| Name                             | Type        | Description                                     | **Required** | default                              |
| -------------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                               | bigint      | 按指定方法生成                                  | required     | 主键                                 |
| muck_truck_purchase_intention_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| supplier                         | text        | 供应商                                          | optional     | **enterprise_info表**的enterprise_id |
| customer_name                    | text        | 购车用户姓名                                    | optional     |                                      |
| customer_phone                   | text        | 购车用户电话                                    | optional     |                                      |
| owner_enterprise                 | text        | 所有人企业                                      | optional     | **enterprise_info表**的enterprise_id |
| province_id                      | text        | 所有人所在省                                    | optional     | **省份表**province_id                |
| city_id                          | text        | 所有人所在市                                    | optional     | **城市表**city_id                    |
| district_id                      | text        | 所有人所在县                                    | optional     | **区域表**district_id                |
| capacigy_application             | text        | 运力申请                                        | optional     |                                      |
| brand_model                      | text        | 品牌型号                                        | optional     |                                      |
| vehicle_purchase                 | integer     | 购车数量                                        | optional     |                                      |
| code                             | text        | 编码                                            | optional     |                                      |
| registration_time                | timestamptz | 登记日期                                        | optional     |                                      |
| registration_user                | text        | 登记用户                                        | optional     | **system_user表**的user_id           |
| review                           | text        | 审核                                            | optional     |                                      |
| remarks                          | text        | 备注                                            | optional     |                                      |
| is_deleted                       | boolean     | 是否删除                                        | optional     |                                      |
| created_at                       | timestamptz | 创建时间                                        | required     |                                      |
| created_by                       | text        | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at                       | timestamptz | 修改时间                                        | optional     |                                      |
| updated_by                       | text        | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at                       | timestamptz | 删除时间                                        | optional     |                                      |
| deleted_by                       | text        | 删除人                                          | optional     | **system_user表**的user_id           |

# 8. 车辆抓拍识别系统

### 8.1 construction_info 工地信息表

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| construction_info_id         | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| project_name                 | text        | 项目名称                                        | optional     |                            |
| project_code                 | text        | 工程项目代码                                    | optional     |                            |
| contract_start_date          | timestamptz | 合同开工日期                                    | optional     |                            |
| contract_completion_date     | timestamptz | 合同竣工日期                                    | optional     |                            |
| build_agency                 | text        | 建设单位                                        | optional     |                            |
| project_manager              | text        | 项目负责人                                      | optional     |                            |
| area_id                      | text        | 地区                                            | optional     |                            |
| address                      | text        | 地址                                            | optional     |                            |
| map_coordinate               | point       | 地图经纬度                                      | optional     |                            |
| construction_agency          | text        | 施工单位                                        | optional     | 与建设单位的区别？         |
| supervision_agency           | text        | 监理单位                                        | optional     |                            |
| device_installation_agency   | text        | 设备安装单位                                    | optional     |                            |
| device_installation_location | text        | 设备安装地点                                    | optional     |                            |
| service_contract_number      | text        | 服务合同编号                                    | optional     |                            |
| contract_signing_time        | timestamptz | 合同签署时间                                    | optional     |                            |
| service_activation_date      | timestamptz | 服务启用日期                                    | optional     |                            |
| proof_valid_date             | timestamptz | 证明有效期                                      | optional     |                            |
| coordinate                   | point       | 经纬度                                          | optional     |                            |
| installation_time            | timestamptz | 安装时间                                        | optional     |                            |
| project_name_tag             | text        | 项目名称标签                                    | optional     |                            |
| import_export_quantity       | text        | 进出口数量                                      | optional     |                            |
| deactivation_reason          | text        | 停用原因                                        | optional     |                            |
| uninstalled_reason           | text        | 未安装原因                                      | optional     |                            |
| proof_valid_date_until       | timestamptz | 证明起始有效期                                  | optional     |                            |
| construction_type            | integer     | 类型（1.工地 2.矿山）                           | optional     |                            |
| contact_person               | text        | 联系人                                          | optional     |                            |
| contact_phone                | text        | 联系电话                                        | optional     |                            |
| finished                     | integer     | 是否完工（0.未完工 1.已完工 2.已停用）          | optional     |                            |
| is_city_directly             | text        | 是否市直属                                      | optional     |                            |
| is_installation              | text        | 是否安装（0.未申请 1.已安装 2.已申请未安装）    | optional     |                            |
| is_deleted                   | boolean     | 是否删除                                        | optional     |                            |
| registration_time            | timestamptz | 登记时间                                        | optional     |                            |
| created_at                   | timestamptz | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### 8.2 construction_upload_pic 工地上传图片表

| Name                       | Type        | Description                                     | **Required** | default                                                |
| -------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------ |
| id                         | bigint      | 按指定方法生成                                  | required     | 主键                                                   |
| construction_upload_pic_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                               |
| construction_info_id       | text        | 工地ID                                          | optional     | **construction_info 工地信息表**的construction_info_id |
| picture_url                | text        | 图片                                            | optional     |                                                        |
| is_deleted                 | boolean     | 是否删除                                        | optional     |                                                        |
| created_at                 | timestamptz | 创建时间                                        | required     |                                                        |
| created_by                 | text        | 创建人                                          | required     | **system_user表**的user_id                             |
| updated_at                 | timestamptz | 修改时间                                        | optional     |                                                        |
| updated_by                 | text        | 修改人                                          | optional     | **system_user表**的user_id                             |
| deleted_at                 | timestamptz | 删除时间                                        | optional     |                                                        |
| deleted_by                 | text        | 删除人                                          | optional     | **system_user表**的user_id                             |

### 8.3 construction_camera 工地摄像头表

| Name                   | Type        | Description                                     | **Required** | default                                                |
| ---------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------ |
| id                     | bigint      | 按指定方法生成                                  | required     | 主键                                                   |
| construction_camera_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                               |
| device_name            | text        | 设备名称                                        | optional     |                                                        |
| ip_address             | text        | IP地址                                          | optional     |                                                        |
| port                   | text        | 端口                                            | optional     |                                                        |
| device_id              | text        | 设备编号                                        | optional     |                                                        |
| construction_info_id   | text        | 关联工地                                        | optional     | **construction_info 工地信息表**的construction_info_id |
| sim_number             | text        | 卡号                                            | optional     |                                                        |
| registration_time      | timestamptz | 登记时间                                        | optional     |                                                        |
| is_deleted             | boolean     | 是否删除                                        | optional     |                                                        |
| created_at             | timestamptz | 创建时间                                        | required     |                                                        |
| created_by             | text        | 创建人                                          | required     | **system_user表**的user_id                             |
| updated_at             | timestamptz | 修改时间                                        | optional     |                                                        |
| updated_by             | text        | 修改人                                          | optional     | **system_user表**的user_id                             |
| deleted_at             | timestamptz | 删除时间                                        | optional     |                                                        |
| deleted_by             | text        | 删除人                                          | optional     | **system_user表**的user_id                             |

### 8.4 vehicle_passing_record 过车记录表

| Name                      | Type        | Description                                     | **Required** | default                                                      |
| ------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                        | bigint      | 按指定方法生成                                  | required     | 主键                                                         |
| vehicle_passing_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| vehicle_id                | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                                                              |
| license_plate_number      | text        | 车牌号                                          | optional     |                                                              |
| passing_time              | timestamptz | 过车时间                                        | optional     |                                                              |
| registration_time         | timestamptz | 登记时间                                        | optional     |                                                              |
| camera_id                 | text        | 摄像头ID                                        | optional     | **construction_camera 工地摄像头表**的construction_camera_id |
| pass_type                 | text        | 过车类型                                        | optional     |                                                              |
| picture_url               | text        | 图片地址                                        | optional     |                                                              |
| review_status             | text        | 审核状态                                        | optional     |                                                              |
| vehicle_type              | text        | 车辆类型                                        | optional     |                                                              |
| is_online                 | boolean     | 是否在线                                        | optional     |                                                              |
| created_at                | timestamptz | 创建时间                                        | required     |                                                              |
| created_by                | text        | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at                | timestamptz | 修改时间                                        | optional     |                                                              |
| updated_by                | text        | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at                | timestamptz | 删除时间                                        | optional     |                                                              |
| deleted_by                | text        | 删除人                                          | optional     | **system_user表**的user_id                                   |

### 8.5 snapshot_system_passing_alarm 抓拍系统过车报警表

| Name                             | Type        | Description                                                | **Required** | default                                                      |
| -------------------------------- | ----------- | ---------------------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                               | bigint      | 按指定方法生成                                             | required     | 主键                                                         |
| snapshot_system_passing_alarm_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用            | required     | 联合主键                                                     |
| passing_id                       | text        | 过车记录ID                                                 |              | **vehicle_passing_record 过车记录表**的vehicle_passing_record_id |
| vehicle_id                       | text        | **vehicle_info** 车辆信息表 的vehicle_id                   | required     |                                                              |
| license_plate_number             | text        | 车牌号                                                     | optional     |                                                              |
| registration_time                | timestamptz | 登记时间                                                   | optional     |                                                              |
| supervision_time                 | timestamptz | 监管时间                                                   | optional     |                                                              |
| remarks                          | text        | 备注                                                       | optional     |                                                              |
| camera_id                        | text        | 摄像头ID                                                   | optional     | **construction_camera 工地摄像头表**的construction_camera_id |
| alarm_type                       | integer     | 报警类型（1.未纳入管控平台 2.未报备工程 3.未纳入管控平台） | optional     |                                                              |
| review_status                    | text        | 审核状态                                                   | optional     |                                                              |
| vehicle_type                     | integer     | 车辆类型                                                   | optional     | **车辆类型**字典                                             |
| is_online                        | boolean     | 是否在线                                                   | optional     |                                                              |
| is_supervision                   | boolean     | 是否监管                                                   | optional     |                                                              |
| created_at                       | timestamptz | 创建时间                                                   | required     |                                                              |
| created_by                       | text        | 创建人                                                     | required     | **system_user表**的user_id                                   |
| updated_at                       | timestamptz | 修改时间                                                   | optional     |                                                              |
| updated_by                       | text        | 修改人                                                     | optional     | **system_user表**的user_id                                   |
| deleted_at                       | timestamptz | 删除时间                                                   | optional     |                                                              |
| deleted_by                       | text        | 删除人                                                     | optional     | **system_user表**的user_id                                   |

### 8.6 offline_registration_record 不在线登记记录表

| Name                        | Type        | Description                                     | **Required** | default                                                    |
| --------------------------- | ----------- | ----------------------------------------------- | ------------ | ---------------------------------------------------------- |
| id                          | bigint      | 按指定方法生成                                  | required     | 主键                                                       |
| offline_registration_record | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                   |
| outage_registration_id      | text        | 停运登记表的ID                                  | optional     | **outage_registration 停运登记表**的outage_registration_id |
| registration_content        | text        | 登记内容                                        | optional     |                                                            |
| registration_time           | timestamptz | 登记时间                                        | optional     |                                                            |
| registrant                  | text        | 登记人                                          | optional     |                                                            |
| is_deleted                  | boolean     | 是否删除                                        | optional     |                                                            |
| created_at                  | timestamptz | 创建时间                                        | required     |                                                            |
| created_by                  | text        | 创建人                                          | required     | **system_user表**的user_id                                 |
| updated_at                  | timestamptz | 修改时间                                        | optional     |                                                            |
| updated_by                  | text        | 修改人                                          | optional     | **system_user表**的user_id                                 |
| deleted_at                  | timestamptz | 删除时间                                        | optional     |                                                            |
| deleted_by                  | text        | 删除人                                          | optional     | **system_user表**的user_id                                 |