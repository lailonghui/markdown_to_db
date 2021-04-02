## 车辆信息(公共)

### veh_info 车辆信息表

**根据以下数据表提取出公用部分**

重点车辆：
[VD_VEHICLE 车辆信息](#major_VD_VEHICLE)

营运车辆：
[INFO_VEHICLE 车辆信息 ](#oper_INFO_VEHICLE)

工程车:

[GPSVEHICLE 基础车辆信息表 ](#egn_GPSVEHICLE)

[VEHICLEMANAGE 目录库车辆信息表](#egn_VEHICLEMANAGE)

| Name                                     | Type           | Description                                                  | **Required** | default                                                      |
| ---------------------------------------- | -------------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------------ |
| id                                       | int8           | 按指定方法生成，生成方法见下面说明                           | required     | 主键                                                         |
| veh_id                                   | text           | 车辆外部编码，由golang程序生成的xid，暴露到外部使用          | required     | 联合主键                                                     |
| owner_type                               | int4           | 车主类型（1.企业，2，个体）                                  | optional     |                                                              |
| enterprise_id                            | text           | 所在企业id                                                   | optional     | **enterprise_info表**的enterprise_id                         |
| owner_id                                 | text           | 车主id(个体车主信息表)                                       | optional     |                                                              |
| industry_category                        | int4           | 行业类别                                                     | optional     | **行业类别**字典                                             |
| business_scope                           | int4           | 经营范围                                                     | optional     | **经营范围**字典                                             |
| vehicle_type                             | int4           | 车辆类型                                                     | optional     | **车辆类型**字典                                             |
| operating_type                           | int4           | 营运类型                                                     | optional     | **营运类型**字典                                             |
| operating_state                          | int4           | 营运状态                                                     | optional     | **营运状态**字典                                             |
| operating_route                          | text           | 营运线路                                                     | optional     |                                                              |
| terminal_id                              | text           | 终端ID                                                       | optional     |                                                              |
| is_apply_install_terminal                | boolean        | 是否申请安装智能终端                                         | optional     |                                                              |
| license_plate_number                     | text           | 车牌号                                                       | optional     |                                                              |
| license_plate_color                      | int4           | 车牌颜色                                                     | optional     | **车牌颜色**字典                                             |
| license_plate_type                       | int4           | 号牌种类                                                     | optional     | **号牌种类**字典                                             |
| vehicle_identification_number            | text           | 车架号(后6位)                                                | optional     | 车辆识别代号vin,如D02133                                     |
| road_transport_license_number            | text           | 道路运输证号                                                 | optional     |                                                              |
| heavy                                    | numeric        | 吨位                                                         | optional     |                                                              |
| seats                                    | int4           | 座位                                                         | optional     |                                                              |
| vehicle_manager                          | text           | 机动车管理人                                                 | optional     |                                                              |
| vehicle_manager_phone                    | text           | 机动车管理人联系电话                                         | optional     |                                                              |
| vehicle_manager_id_card                  | text           | 机动车管理人身份证                                           | optional     |                                                              |
| owner                                    | text           | 机动车所有人（六合一）                                       | optional     |                                                              |
| inspection_date                          | timestamptz(6) | 检验日期（六合一）                                           | optional     |                                                              |
| retirement_date                          | timestamptz(6) | 报废日期（六合一）                                           | optional     |                                                              |
| use_nature                               | text           | 使用性质（六合一）                                           | optional     |                                                              |
| vehicle_state                            | int4           | 机动车状态                                                   | optional     | **车辆状态**字典                                             |
| update_time_in                           | timestamptz(6) | 内网更新时间                                                 | optional     |                                                              |
| remark_in                                | text           | 车辆信息同步内网反馈信息                                     | optional     | 车辆信息同步到公安内网后内网的反馈内容，如车牌号填写错误会反馈车辆号牌错误 |
| is_complete                              | boolean        | 是否完成                                                     | optional     | 用于标志车辆资料是否处于确定状态。未确定状态的车辆信息在系统上除车辆管理外的功能中都查不到 |
| driving_license_owner                    | text           | 行驶证所有人                                                 | optional     |                                                              |
| driving_licensee_pic                     | text           | 行驶证照片,云储存系统返回的路径                              | optional     |                                                              |
| license_plate_photo                      | text           | 车牌照片,云储存系统返回的路径                                | optional     |                                                              |
| other_photo                              | text           | 其他照片,云储存系统返回的路径                                | optional     |                                                              |
| serial_number                            | text           | 客户自定义编号                                               | optional     |                                                              |
| is_beidou                                | boolean        | 是否北斗部标平台                                             | optional     |                                                              |
| is_in_upload_platform                    | boolean        | 是否上传货运平台                                             | optional     |                                                              |
| service_expiration_time                  | timestamptz(6) | 服务到期时间                                                 | optional     |                                                              |
| is_active                                | boolean        | 缴费是否激活                                                 | optional     |                                                              |
| insurance_company                        | int4           | 投保公司                                                     | optional     | **投保公司**字典                                             |
| insurance_date                           | timestamptz(6) | 投保日期                                                     | optional     |                                                              |
| insurance_expiry_time                    | timestamptz(6) | 保险到期时间                                                 | optional     |                                                              |
| compulsory_traffic_insurance_expiry_date | timestamptz(6) | 交强险到期时间                                               | optional     |                                                              |
| annual_inspection_expiration_time        | timestamptz(6) | 年检到期时间                                                 | optional     |                                                              |
| vehicle_maintenances                     | jsonb          | 维保数据，字段包括: 1.maintenance_ date 维保时间<br />2.maintenance_ kilometers 维保公里数 | optional     |                                                              |
| vehicle_displacement                     | text           | 汽车排量                                                     | optional     |                                                              |
| vehicle_brand                            | int4           | 车辆品牌                                                     | optional     | **车辆品牌**字典                                             |
| quasi_driving_models                     | int4           | 准驾车型（驾驶证允许的驾驶车型）                             | optional     | **准驾车型**字典                                             |
| is_upload_province                       | boolean        | 是否上传省厅                                                 | optional     |                                                              |
| data_source                              | int4           | 数据来源（1，2，3）                                          | optional     | 是否通过外部导入的车辆信息                                   |
| engine_number                            | text           | 发动机号                                                     | optional     |                                                              |
| remarks                                  | text           | 备注                                                         | optional     |                                                              |
| function                                 | jsonb          | 功能测试(kv,是否通过，功能测试状态)                          | optional     |                                                              |
| first_online_time                        | timestamptz(6) | 第一次上线时间                                               | optional     |                                                              |
| lastest_binding_terminal_time            | timestamptz(6) | 最新绑定终端时间                                             | optional     |                                                              |
| contract_time                            | timestamptz(6) | 合同时间                                                     | optional     |                                                              |
| installation_time                        | timestamptz(6) | 安装时间                                                     | optional     |                                                              |
| contact_number                           | text           | 合同编号                                                     | optional     |                                                              |
| driving_liscense_owner_id_photo          | text           | 行驶证登记的车主身份证照片                                   | optional     |                                                              |
| initial_registration_date                | timestamptz(6) | 行驶证初次登记日期                                           | optional     |                                                              |
| vehicle_operating_certificate_number     | text           | 车辆营运证号                                                 | optional     |                                                              |
| vehicle_operating_certificate_photo      | text           | 车辆营运证照片                                               | optional     |                                                              |
| third_party_insurance_expiry_time        | timestamptz(6) | 第三者保险到期时间                                           | optional     |                                                              |
| third_party_insurance_coverage           | text           | 第三者保额                                                   | optional     |                                                              |
| third_party_insurance_policy_picture     | text           | 第三者保单图片                                               | optional     |                                                              |
| compulsory_insurance_policy_picture      | text           | 交强险保单图片                                               | optional     |                                                              |
| expiry_date                              | timestamptz(6) | 费用到期时间                                                 | optional     |                                                              |
| is_deleted                               | boolean        | 是否删除                                                     | optional     |                                                              |
| created_at                               | timestamptz(6) | 创建时间                                                     | required     |                                                              |
| created_by                               | text           | 创建人                                                       | required     | **system_user表**的user_id                                   |
| updated_at                               | timestamptz(6) | 修改时间                                                     | optional     |                                                              |
| updated_by                               | text           | 修改人                                                       | optional     | **system_user表**的user_id                                   |
| deleted_at                               | timestamptz(6) | 删除时间                                                     | optional     |                                                              |
| deleted_by                               | text           | 删除人                                                       | optional     | **system_user表**的user_id                                   |

## 车辆信息(重点车)

### veh_major_ext 重点车辆拓展表

**根据以下数据表整合而成**

重点车:

[VD_VEHICLE 车辆信息 ](#major_VD_VEHICLE)

| Name         | Type           | Description                       | **Required** | default                    |
| ------------ | -------------- | --------------------------------- | ------------ | -------------------------- |
| id           | int8           | 按指定方法生成                    | required     | 主键                       |
| vehicle_id   | text           | 车辆信息表 的vehicle_id           | required     |                            |
| is_input     | int4           | 录入完整性状态（1.完整，2不完整） | optional     |                            |
| is_blacklist | boolean        | 是否黑名单                        | optional     |                            |
| created_at   | timestamptz(6) | 创建时间                          | required     |                            |
| created_by   | text           | 创建人                            | required     | **system_user表**的user_id |
| updated_at   | timestamptz(6) | 修改时间                          | optional     |                            |
| updated_by   | text           | 修改人                            | optional     | **system_user表**的user_id |
| deleted_at   | timestamptz(6) | 删除时间                          | optional     |                            |
| deleted_by   | text           | 删除人                            | optional     | **system_user表**的user_id |

### veh_security_check 车辆安全检查表

**根据以下数据表整合而成**

重点车:

[INFO_VEH_SAFECHECK 车辆安全检查表](#major_INFO_VEH_SAFECHECK)  

| Name                  | Type           | Description                                     | **Required** | default                    |
| --------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                    | bigint         | 按指定方法生成                                  | required     | 主键                       |
| veh_security_check_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id            | text           | 车辆信息表 的vehicle_id                         | required     |                            |
| enterprise_id         | text           | 所在企业id                                      | optional     |                            |
| brake                 | int2           | 刹车                                            | optional     |                            |
| tire                  | int2           | 轮胎                                            | optional     |                            |
| screw                 | int2           | 螺丝                                            | optional     |                            |
| hydraulic_oil         | int2           | 液压油                                          | optional     |                            |
| engine_oil            | int2           | 机油                                            | optional     |                            |
| water                 | int2           | 水                                              | optional     |                            |
| headlight             | int2           | 大灯                                            | optional     |                            |
| taillight             | int2           | 尾灯                                            | optional     |                            |
| turn_signal           | int2           | 转向灯                                          | optional     |                            |
| brake_light           | int2           | 刹车灯                                          | optional     |                            |
| last_check_time       | timestamptz(6) | 最后检查时间                                    | optional     |                            |
| created_at            | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by            | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by            | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by            | text           | 删除人                                          | optional     | **system_user表**的user_id |

## 车辆信息(营运车)

### veh_oper_ext  营运车辆扩展表

**根据以下数据表整合而成**

营运车:

[INFO_VEHICLE 车辆信息  ](#oper_INFO_VEHICLE)

| Name                            | Type           | Description                                | **Required** | default                    |
| ------------------------------- | -------------- | ------------------------------------------ | ------------ | -------------------------- |
| id                              | int8           | 按指定方法生成                             | required     | 主键                       |
| vehicle_id                      | text           | 车辆信息表 的vehicle_id                    | required     |                            |
| scrap_reason                    | text           | 营运证报废原因                             | optional     |                            |
| scrap_time                      | timestamptz(6) | 营运证报废时间                             | optional     |                            |
| scrap_time_check                | int4           | 营运证报废时间审核                         | optional     |                            |
| is_in_operating_system          | boolean        | 是否在运证系统                             | optional     |                            |
| supervise                       | jsonb          | 监管信息(是否需要监管，被谁监管，监管状态) | optional     |                            |
| is_applay_terminal_installation | boolean        | 是否办理终端安装证明                       | optional     |                            |
| is_input                        | int4           | 录入完整性状态（1.完整，2不完整）          | optional     |                            |
| is_expired                      | boolean        | 是否过期车辆                               | optional     |                            |
| created_at                      | timestamptz(6) | 创建时间                                   | required     |                            |
| created_by                      | text           | 创建人                                     | required     | **system_user表**的user_id |
| updated_at                      | timestamptz(6) | 修改时间                                   | optional     |                            |
| updated_by                      | text           | 修改人                                     | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz(6) | 删除时间                                   | optional     |                            |
| deleted_by                      | text           | 删除人                                     | optional     | **system_user表**的user_id |

### owner_info 个体车主信息表

**根据以下数据表整合而成**

营运车:

[INFO_VEHICLE_OWNER 车主信息](#oper_INFO_VEHICLE_OWNER)

| Name                          | Type           | Description                                             | **Required** | default                    |
| ----------------------------- | -------------- | ------------------------------------------------------- | ------------ | -------------------------- |
| id                            | int8           | 按指定方法生成                                          | required     | 主键                       |
| owner_id                      | text           | 车主信息外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| department_id            | text           | 所在部门id                                              | optional     | **department 部门信息表**  |
| name                          | text           | 车主姓名                                                | optional     |                            |
| address                       | text           | 联系地址                                                | optional     |                            |
| cellphone                     | text           | 固定电话                                                | optional     |                            |
| telephone                     | text           | 手机号码                                                | optional     |                            |
| email                         | text           | 邮箱地址                                                | optional     |                            |
| expiry_date                   | timestamptz(6) | 证件过期日期                                            | optional     |                            |
| id_number                     | text           | 身份证号                                                | optional     |                            |
| remarks                       | text           | 备注                                                    | optional     |                            |
| sex                           | int4           | 车主性别                                                | optional     | **性别**字典               |
| driving_license_id_number     | text           | 行驶证登记的车主身份证号                                | optional     |                            |
| driving_license_contact_phone | text           | 行驶证登记的车主联系电话                                | optional     |                            |
| is_deleted                    | boolean        | 是否删除                                                | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                                | required     |                            |
| created_by                    | text           | 创建人                                                  | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                                | optional     |                            |
| updated_by                    | text           | 修改人                                                  | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                                | optional     |                            |
| deleted_by                    | text           | 删除人                                                  | optional     | **system_user表**的user_id |

### oper_outage_registration 车辆停运报备记录

**根据以下数据表整合而成**

营运车:

[MON_VEH_STOP_RECORD 车辆停运报备记录](#oper_MON_VEH_STOP_RECORD)

| Name                        | Type           | Description                                     | **Required** | default                    |
| --------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                          | int8           | 按指定方法生成                                  | required     | 主键                       |
| oper_outage_registration_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| bill_num                    | text           | 报备单据号                                      | optional     |                            |
| vehicle_id                  | text           | 车辆信息表 的vehicle_id                         | optional     |                            |
| license_plate_number        | text           | 车牌号                                          | optional     |                            |
| outage_start_time           | timestamptz(6) | 报备起始时间                                    | optional     |                            |
| outage_end_time             | timestamptz(6) | 报备结束时间                                    | optional     |                            |
| outage_person               | text           | 报备人                                          | optional     | **system_user表**的user_id |
| effect_flag                 | int4           | 生效标识                                        | optional     | **车辆报备生效状态**字典   |
| effect_start_time           | timestamptz(6) | 生效起始时间                                    | optional     |                            |
| effect_end_time             | timestamptz(6) | 生效结束时间                                    | optional     |                            |
| outage_reason               | text           | 停运原因                                        | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                  | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                  | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                  | text           | 删除人                                          | optional     | **system_user表**的user_id |

### latest_veh_photo 车辆最新拍照图片表

**根据以下数据表整合而成**

营运车:

[MON_LATEST_PICTURE 车辆最新拍照图片表](#oper_MON_LATEST_PICTURE)  

| Name                | Type           | Description                                     | **Required** | default                    |
| ------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                  | int8           | 按指定方法生成                                  | required     | 主键                       |
| latest_veh_photo_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| latest_pic_url      | text           | 最新照片                                        | optional     |                            |
| pic_id              | text           | 照片id                                          | optional     |                            |
| created_at          | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by          | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at          | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by          | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at          | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by          | text           | 删除人                                          | optional     | **system_user表**的user_id |

### oper_veh_trans_sys 运政系统车辆表

**根据以下数据表整合而成**

营运车:

[YZ_VEHICLEINFO 运政系统车辆表](#oper_YZ_VEHICLEINFO)

| Name                          | Type           | Description      | **Required** | default                    |
| ----------------------------- | -------------- | ---------------- | ------------ | -------------------------- |
| id                            | int8           | 按指定方法生成   | required     | 主键                       |
| license_plate_number          | text           | 车牌号码         | required     |                            |
| license_plate_color           | text           | 车牌颜色         | required     |                            |
| seats                         | text           | 核定座位         | required     |                            |
| household_name                | text           | 业户名           | optional     |                            |
| transport_license_number      | text           | 运输证号         | optional     |                            |
| business_scope                | text           | 经营范围         | optional     |                            |
| license_number                | text           | 许可证号         | optional     |                            |
| annual_review_date            | text           | 年审日期         | optional     |                            |
| issue_date                    | text           | 发证日期         | optional     |                            |
| annual_review_valid           | text           | 年审有效期       | optional     |                            |
| household_addr                | text           | 业户地址         | optional     |                            |
| engine_number                 | text           | 发动机号         | optional     |                            |
| vehicle_identification_number | text           | 车架号码         | optional     |                            |
| admin_division                | text           | 行政区域         | optional     |                            |
| brand_model                   | text           | 厂牌型号         | optional     |                            |
| contact_phone                 | text           | 联系电话         | optional     |                            |
| driving_models                | text           | 车型             | optional     |                            |
| heavy                         | text           | 核定吨位         | optional     |                            |
| annual_review_valid_until     | text           | 行驶证年审有效期 | optional     |                            |
| total_mass                    | text           | 总质量           | optional     |                            |
| specific_admin_division       | text           | 具体行政区域     | optional     |                            |
| vehicle_state                 | text           | 车辆状态         | optional     |                            |
| operating_type                | text           | 营运类型         | optional     |                            |
| is_major                      | text           | 是否重点车辆     | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间         | required     |                            |
| created_by                    | text           | 创建人           | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间         | optional     |                            |
| updated_by                    | text           | 修改人           | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间         | optional     |                            |
| deleted_by                    | text           | 删除人           | optional     | **system_user表**的user_id |

### veh_province_online_num 省厅在线数

**根据以下数据表整合而成**

营运车:

[INFO_VEHICLE_ONLINE  省厅在线数](#oper_INFO_VEHICLE_ONLINE)

| Name                       | Type           | Description                                     | **Required** | default                    |
| -------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | int8           | 按指定方法生成                                  | required     | 主键                       |
| veh_province_online_num_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| total_veh_num              | int4           | 总车辆数                                        | optional     |                            |
| online_veh_num             | int4           | 在线车辆数                                      | optional     |                            |
| local_online_veh_num       | int4           | 本地车辆在线数                                  | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text           | 删除人                                          | optional     | **system_user表**的user_id |

### veh_province_upload 省厅平台车辆信息表

**根据以下数据表整合而成**

营运车:

[INFO_VEHICLE_ONLINE_DETAIL 省厅车辆在线数明细](#oper_INFO_VEHICLE_ONLINE_DETAIL)

[INFO_VEHOL_DET_LAST_NEW 省厅新平台上传车辆明细（最新）](#oper_INFO_VEHOL_DET_LAST_NEW)

[INFO_VEHOL_DET_NEW 省厅新平台上传车辆总表](#oper_INFO_VEHOL_DET_NEW)

[INFO_VEH_OL_DETAIL_LAST 省厅车辆在线明细](#oper_INFO_VEH_OL_DETAIL_LAST)

[ST_VEHINFO 省厅车辆表](#new_ST_VEHINFO)

| Name                                  | Type           | Description                                     | **Required** | default                              |
| ------------------------------------- | -------------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                                    | int8           | 按指定方法生成                                  | required     | 主键                                 |
| veh_province_upload_id                | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| license_plate_number                  | text           | 车牌号码                                        | optional     |                                      |
| license_plate_color                   | int4           | 车牌颜色                                        | optional     | **车牌颜色**字典                     |
| enterprise_id                         | text           | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id |
| operator                              | text           | 运营商                                          | optional     | **enterprise_info表**的enterprise_id |
| vehicle_registration_place            | text           | 车籍地(行驶证上的车辆注册地 )                   | optional     |                                      |
| vehicle_type                          | int4           | 车辆类型                                        | optional     | **车辆类型**字典                     |
| heavy                                 | numeric        | 吨位                                            | optional     |                                      |
| seats                                 | int4           | 座位                                            | optional     |                                      |
| operating_flag                        | text           | 运营标识                                        | optional     |                                      |
| operating_state                       | int4           | 营运状态                                        | optional     | **营运状态**字典                     |
| remarks                               | text           | 备注                                            | optional     |                                      |
| online_status                         | text           | 在线状态                                        | optional     |                                      |
| last_report_time                      | timestamptz(6) | 最后汇报时间                                    | optional     |                                      |
| transport_agency_synchronization_flag | text           | 运政同步标识                                    | optional     |                                      |
| is_transport_agency_synchronization   | boolean        | 是否运政同步                                    | optional     |                                      |
| transport_agency_synchronization_time | timestamptz(6) | 运政同步时间                                    | optional     |                                      |
| device_model                          | text           | 设备型号                                        | optional     |                                      |
| terminal_model                        | text           | 终端型号                                        | optional     |                                      |
| created_at                            | timestamptz(6) | 创建时间                                        | required     |                                      |
| created_by                            | text           | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at                            | timestamptz(6) | 修改时间                                        | optional     |                                      |
| updated_by                            | text           | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at                            | timestamptz(6) | 删除时间                                        | optional     |                                      |
| deleted_by                            | text           | 删除人                                          | optional     | **system_user表**的user_id           |

### veh_night_operation_whitelist 车辆夜间运营白名单表

**根据以下数据表整合而成**

营运车:

[MON_VEH_NIGHT_RECORD 车辆夜间运营白名单表](#oper_MON_VEH_NIGHT_RECORD) 

| Name                             | Type           | Description                                     | **Required** | default                    |
| -------------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                               | int8           | 按指定方法生成                                  | required     | 主键                       |
| veh_night_operation_whitelist_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                       | text           | 车辆信息表 的vehicle_id                         | optional     |                            |
| ent_id                           | text           | 所属企业                                        | optional     |                            |
| trans_route                      | text           | 运输路线                                        | optional     |                            |
| operator                         | text           | 运营商                                          | optional     |                            |
| veh_type                         | text           | 车辆类型                                        | optional     |                            |
| operating_type                   | int4           | 营运类型                                        | optional     | **营运类型**字典           |
| is_effect                        | int4           | 是否生效(1生效2未审核3审核不通过)               | optional     | **夜间白名单是否生效**字典 |
| audit_fail_reason                | text           | 审核不通过原因                                  | optional     |                            |
| ent_time                         | timestamptz(6) | 结束时间                                        | optional     |                            |
| remarks                          | text           | 备注                                            | optional     |                            |
| created_at                       | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                       | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                       | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                       | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                       | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                       | text           | 删除人                                          | optional     | **system_user表**的user_id |

### vehicle_supervision_photo 车辆监控图片表

**根据以下数据表整合而成**

营运车:

REPORT_PICTURE 车辆图片信息

| Name                       | Type        | Description                                                 | **Required** | default          |
| -------------------------- | ----------- | ----------------------------------------------------------- | ------------ | ---------------- |
| id                         | bigint      | 按指定方法生成                                              | required     | 主键             |
| supervision_photo_id       | text        | 车辆监控图片外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键         |
| license_plate_number       | text        | 车牌号                                                      | optional     |                  |
| pic_name        | text        | 图片名称                                                    | optional     |                  |
| pic_time | timestamptz(6) | 照片时间                                                    | optional |                  |
| pic_upload_time | timestamptz(6) | 上传时间                                                    | optional     |                  |
| camera_id                  | integer     | 摄像头ID                                                    | optional     | **摄像头ID**字典 |
| photo_condition            | text        | 拍照条件                                                    | optional     | **拍照条件**字典 |
| pic_address     | text        | 监控图片地址                                                | optional     |                  |
| created_at                       | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                       | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                       | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                       | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                       | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                       | text           | 删除人                                          | optional     | **system_user表**的user_id |


## 车辆信息(工程车)

### veh_egn_ext 工程运输车扩展表

**根据以下数据表整合而成**

工程车:

[GPSVEHICLE 基础车辆信息表 ](#egn_GPSVEHICLE)

[VEHICLEMANAGE 目录库车辆信息表](#egn_VEHICLEMANAGE)

| Name                                | Type           | Description                                                  | **Required** | default                    |
| ----------------------------------- | -------------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                                  | int8           | 按指定方法生成                                               | required     | 主键                       |
| vehicle_id                          | text           | **veh_info** 车辆信息表 的vehicle_id                         | required     |                            |
| vehicle_picture                     | text           | 自有车辆上传图片                                             | optional     |                            |
| review                              | jsonb          | 审核信息(review_time审核时间，reviewer 审核人，review_depart审核部门,review_notes审核备注,review_status审核状态) | optional     |                            |
| vehicle_description                 | text           | 工程车辆用户描述                                             | optional     |                            |
| is_blacklist                        | boolean        | 是否黑名单                                                   | optional     |                            |
| blacklist_deadline                  | timestamptz(6) | 黑名单截止日期                                               | optional     |                            |
| owner_id_photo                      | text           | 车主身份证照片                                               | optional     |                            |
| vehicle_nature                      | int4           | 车辆性质  1.本企业车辆  2.企业挂靠车辆  3.车队挂靠车辆       | optional     |                            |
| actual_owner_name                   | text           | 实际车主姓名                                                 | optional     |                            |
| actual_owner_id_number              | text           | 实际车主身份证号                                             | optional     |                            |
| actual_owner_id_photo               | text           | 实际车主身份证照片                                           | optional     |                            |
| actual_owner_contact_phone          | text           | 实际车主联系电话                                             | optional     |                            |
| illegal_number_endorsement          | text           | 违法编号签注                                                 | optional     |                            |
| illegal_notice_number_endorsement   | text           | 违法通知书编号签注                                           | optional     |                            |
| illegal_compulsory_measures_number  | text           | 违法强制措施编号签注                                         | optional     |                            |
| incident_number_endorsement         | text           | 事故编号签注                                                 | optional     |                            |
| is_engin_vehoffice_audit            | boolean        | 是否渣土办审核                                               | optional     |                            |
| is_first_register                   | boolean        | 是否首次注册                                                 | optional     |                            |
| total_mass                          | numeric        | 总质量                                                       | optional     |                            |
| axes_number                         | int4           | 轴数                                                         | optional     |                            |
| issue_date_of_quangong_number       | timestamptz(6) | 泉工号发放日期                                               | optional     |                            |
| apply_inspection_time               | timestamptz(6) | 申请验车时间                                                 | optional     |                            |
| signing_acceptance_application_time | timestamptz(6) | 签收验车申请时间                                             | optional     |                            |
| area_id                             | int8           | 管辖地区ID（工号牌所属地区）                                 | optional     | 350600                     |
| veh_egn_type                        | int4           | 工程运输车类型（ABCFDE）                                     | optional     | **工程运输车车辆类型**字典 |
| is_reserve_library                  | boolean        | 是否预备库                                                   | optional     |                            |
| worker_id_card                      | text           | 工号牌（自编号：泉工A014015）                                | optional     |                            |
| is_input                            | int4           | 录入完整性状态（1.完整，2不完整）                            | optional     |                            |
| is_deleted                          | boolean        | 是否删除                                                     | optional     |                            |
| created_at                          | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by                          | text           | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                          | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by                          | text           | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                          | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by                          | text           | 删除人                                                       | optional     | **system_user表**的user_id |

### veh_new_record 车辆新增记录

**根据以下数据表整合而成**

工程车:

[NEWSVNUMADDLOG 车辆新增记录](#egn_NEWSVNUMADDLOG)

| Name                    | Type           | Description                                     | **Required** | default                              |
| ----------------------- | -------------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                      | int8           | 按指定方法生成                                  | required     | 主键                                 |
| veh_new_record_id       | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | optional     | 联合主键                             |
| vehicle_id              | text           | **veh_info** 车辆信息表 的vehicle_id            | optional     |                                      |
| license_plate_number    | text           | 车牌号                                          | optional     |                                      |
| enterprise_id           | text           | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id |
| worker_id_card          | text           | 工号牌（自编号：泉工A014015）                   | optional     | 联合主键                             |
| operation_type          | text           | 操作类型                                        | optional     |                                      |
| enterprise_name         | text           | 单位名称                                        | optional     |                                      |
| area                    | text           | 地区                                            | optional     |                                      |
| original_worker_id_card | text           | 原工号牌                                        | optional     |                                      |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                                      |
| created_by              | text           | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                                      |
| updated_by              | text           | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                                      |
| deleted_by              | text           | 删除人                                          | optional     | **system_user表**的user_id           |

### veh_overdue_his 车辆年检逾期历史表(新表)

**把逾期登记表改为逾期历史表**

工程车:

[VEHYCKMARK 车辆年检逾期登记表](#egn_VEHYCKMARK)

| Name               | Type           | Description                                     | **Required** | default                    |
| ------------------ | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                 | int8           | 按指定方法生成                                  | required     | 主键                       |
| veh_overdue_his_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | optional     | 联合主键                   |
| vehicle_id         | text           | **veh_info** 车辆信息表 的vehicle_id            | optional     |                            |
| start_time         | timestamptz(6) | 逾期开始时间                                    | optional     |                            |
| end_time           | timestamptz(6) | 逾期结束时间                                    | optional     |                            |
| reason             | text           | 逾期原因                                        | optional     |                            |
| remarks            | text           | 备注                                            | optional     |                            |
| created_at         | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by         | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at         | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by         | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at         | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by         | text           | 删除人                                          | optional     | **system_user表**的user_id |



### veh_egn_pre_num 工程运输车预编号表

**根据以下数据表整合而成**

工程车:

[VEHICLE_PREVIEW_SVNUM 车辆预编号表 ](#egn_VEHICLE_PREVIEW_SVNUM)

| Name                      | Type           | Description                                     | **Required** | default                    |
| ------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | int8           | 按指定方法生成                                  | required     | 主键                       |
| veh_egn_pre_num_id        | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                | text           | **vehicle_info** 车辆信息表 的vehicle_id        | optional     |                            |
| svn_number                | text           | 预编号                                          | optional     |                            |
| registration_time         | text           | 登记时间                                        | optional     |                            |
| reviewer                  | text           | 审核人                                          | optional     | **system_user表**的user_id |
| review_time               | timestamptz(6) | 审核时间                                        | optional     |                            |
| review_status             | int4           | 审核状态                                        | optional     |                            |
| front_license_plate       | text           | 前车牌(自编号车牌：泉工A014015)                 | optional     |                            |
| rear_license_plate        | text           | 后车牌(自编号车牌：泉工A014015)                 | optional     |                            |
| side_license_plate        | text           | 侧车牌(自编号车牌：泉工A014015)                 | optional     |                            |
| original_number           | text           | 原编号                                          | optional     |                            |
| production_status         | int4           | 制作状态                                        | optional     |                            |
| production_time           | timestamptz(6) | 制作时间                                        | optional     |                            |
| marking_time              | timestamptz(6) | 制作中时间                                      | optional     |                            |
| production_times          | int4           | 制作次数                                        | optional     |                            |
| submit_production_time    | timestamptz(6) | 提交制牌厂时间                                  | optional     |                            |
| contact_person            | text           | 联系人                                          | optional     |                            |
| contact_phone             | text           | 联系电话                                        | optional     |                            |
| work_number_plate_color   | text           | 工号牌颜色（green.绿色 yellow.黄色）            | optional     |                            |
| initial_registration_date | timestamptz(6) | 初次登记日期                                    | optional     |                            |
| unlawful_violation_number | int4           | 违法未处理数                                    | optional     |                            |
| is_register_sale_order    | boolean        | 是否登记销售订单                                | optional     |                            |
| remarks                   | text           | 备注                                            | optional     |                            |
| is_deleted                | boolean        | 是否删除                                        | optional     |                            |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                | text           | 删除人                                          | optional     | **system_user表**的user_id |

### work_num_production_order 工号牌制作订单

**根据以下数据表整合而成**

工程车:

[VEHICLE_SVNUM_MAKE_BILL 工号牌制作订单 ](#egn_VEHICLE_SVNUM_MAKE_BILL)

| Name                    | Type           | Description                              | **Required** | default                    |
| ----------------------- | -------------- | ---------------------------------------- | ------------ | -------------------------- |
| id                      | int8           | 按指定方法生成                           | required     | 主键                       |
| veh_egn_pre_num_id      | text           | 预编号表ID                               | optional     |                            |
| vehicle_id              | text           | **vehicle_info** 车辆信息表 的vehicle_id | optional     |                            |
| svn_number              | text           | 预编号                                   | optional     |                            |
| registration_time       | text           | 登记时间                                 | optional     |                            |
| front_license_plate     | text           | 前车牌                                   | optional     |                            |
| rear_license_plate      | text           | 后车牌                                   | optional     |                            |
| side_license_plate      | text           | 侧车牌                                   | optional     |                            |
| production_status       | int4           | 制作状态                                 | optional     |                            |
| production_time         | timestamptz(6) | 制作时间                                 | optional     |                            |
| marking_time            | timestamptz(6) | 制作中时间                               | optional     |                            |
| production_times        | int4           | 制作次数                                 | optional     |                            |
| contact_person          | text           | 联系人                                   | optional     |                            |
| contact_phone           | text           | 联系电话                                 | optional     |                            |
| work_number_plate_color | text           | 工号牌颜色（green.绿色 yellow.黄色）     | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by              | text           | 创建人                                   | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by              | text           | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by              | text           | 删除人                                   | optional     | **system_user表**的user_id |

### outage_registration 停运登记表 

**根据以下数据表整合而成**

工程车:

[OUTAGECHECK 停运登记表](#egn_OUTAGECHECK)

| Name                    | Type           | Description                                     | **Required** | default                                        |
| ----------------------- | -------------- | ----------------------------------------------- | ------------ | ---------------------------------------------- |
| id                      | int8           | 按指定方法生成                                  | required     | 主键                                           |
| outage_registration_id  | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                       |
| vehicle_id              | text           | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                                                |
| user_id                 | text           | 用户                                            | optional     | **system_user表**的user_id                     |
| outage_start_time       | timestamptz(6) | 停运起始时间                                    | optional     |                                                |
| outage_end_time         | timestamptz(6) | 停运截止时间                                    | optional     |                                                |
| review_status           | text           | 审核状态                                        | optional     |                                                |
| reviewer                | text           | 审核人                                          | optional     | **system_user表**的user_id                     |
| review_time             | timestamptz(6) | 审核时间                                        | optional     |                                                |
| outage_upload_file_id   | text           | 停运报备上传文件表ID                            | optional     | **outage_upload_file** 的outage_upload_file_id |
| outage_start_coordinate | point          | 停运起始经纬度                                  | optional     |                                                |
| outage_end_coordinate   | point          | 停运结束经纬度                                  | optional     |                                                |
| online_time             | timestamptz(6) | 上线时间                                        | optional     |                                                |
| outage_start_position   | text           | 停运起始位置                                    | optional     |                                                |
| is_management_review    | boolean        | 是否管理部门审核                                | optional     |                                                |
| is_invalid              | boolean        | 是否失效                                        | optional     |                                                |
| is_latest               | boolean        | 是否最新                                        | optional     |                                                |
| is_deleted              | boolean        | 是否删除                                        | optional     |                                                |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                                                |
| created_by              | text           | 创建人                                          | required     | **system_user表**的user_id                     |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                                                |
| updated_by              | text           | 修改人                                          | optional     | **system_user表**的user_id                     |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                                                |
| deleted_by              | text           | 删除人                                          | optional     | **system_user表**的user_id                     |



### outage_upload_file 停运报备上传文件表

**根据以下数据表整合而成**

工程车:

[OUTAGEFILE 停运报备上传文件表](#egn_OUTAGEFILE)

| Name                  | Type           | Description                                     | **Required** | default                    |
| --------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                    | int8           | 按指定方法生成                                  | required     | 主键                       |
| outage_upload_file_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| file_path             | text           | 文件路径                                        | optional     |                            |
| file_type             | text           | 文件类型(commitment.承诺书 other.其他)          | optional     |                            |
| is_deleted            | boolean        | 是否删除                                        | optional     |                            |
| created_at            | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by            | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by            | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by            | text           | 删除人                                          | optional     | **system_user表**的user_id |



### veh_egn_exit_catalog_review 工程车辆退出目录库审核表

**根据以下数据表整合而成**

工程车:

[VEHDELDETAIL 车辆退出目录库审核表](#egn_VEHDELDETAIL)

| Name                             | Type           | Description                                                  | **Required** | default                              |
| -------------------------------- | -------------- | ------------------------------------------------------------ | ------------ | ------------------------------------ |
| id                               | int8           | 按指定方法生成                                               | required     | 主键                                 |
| engin_veh_exit_catalog_review_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                             |
| vehicle_id                       | text           | **vehicle_info** 车辆信息表 的vehicle_id                     | required     |                                      |
| enterprise_id                    | text           | 所在企业id                                                   | optional     | **enterprise_info表**的enterprise_id |
| remarks                          | text           | 备注                                                         | optional     |                                      |
| review_status                    | int4           | 审核状态  0.未完成 1.完成                                    | optional     |                                      |
| exit_type                        | int4           | 类别  3.企业车辆退出目录库  6.零散车辆退出目录库             | optional     |                                      |
| exit_step                        | int4           | 退出步骤  Type=3:{1.AA市A县a企业提交材料,2.A县交警同意,3.A县渣土办同意}  Type=6:{1.A县交警提交材料,2.A县渣土办同意} | optional     |                                      |
| area_id                          | text           | 地区ID                                                       | optional     |                                      |
| original_enterprise_id           | text           | 原单位ID                                                     | optional     | **enterprise_info表**的enterprise_id |
| original_enterprise_name         | text           | 原单位名称                                                   | optional     |                                      |
| original_self_number             | text           | 原自编号                                                     | optional     |                                      |
| is_deleted                       | boolean        | 是否删除                                                     | optional     |                                      |
| created_at                       | timestamptz(6) | 创建时间                                                     | required     |                                      |
| created_by                       | text           | 创建人                                                       | required     | **system_user表**的user_id           |
| updated_at                       | timestamptz(6) | 修改时间                                                     | optional     |                                      |
| updated_by                       | text           | 修改人                                                       | optional     | **system_user表**的user_id           |
| deleted_at                       | timestamptz(6) | 删除时间                                                     | optional     |                                      |
| deleted_by                       | text           | 删除人                                                       | optional     | **system_user表**的user_id           |

### joint_penalty_release 联罚联放表

**根据以下数据表整合而成**

工程车:

[ENFORCEMENTINFO 联罚联放表 ](#egn_ENFORCEMENTINFO)

| Name                     | Type           | Description                                                 | **Required** | default                              |
| ------------------------ | -------------- | ----------------------------------------------------------- | ------------ | ------------------------------------ |
| id                       | int8           | 按指定方法生成                                              | required     | 主键                                 |
| joint_penalty_release_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用             | required     | 联合主键                             |
| vehicle_id               | text           | **vehicle_info** 车辆信息表 的vehicle_id                    | required     |                                      |
| enterprise_id            | text           | 所在企业id                                                  | optional     | **enterprise_info表**的enterprise_id |
| enterprise_type          | int4           | 企业类型 1：企业 2：个人                                    | optional     |                                      |
| self_number              | text           | 自编号                                                      | optional     |                                      |
| veh_owner                | int4           | 车辆权属  1：自有 2：挂靠                                   | optional     |                                      |
| license_type             | int4           | 牌照类型  1：本地 2：外地                                   | optional     |                                      |
| veh_type                 | int4           | 车辆类型  1：环保渣土车 2：非环保渣土车 3：混凝土车 4：其他 | optional     |                                      |
| first_catch_unit         | int4           | 首抓单位 1：交警 2：交通 3：城管                            | optional     |                                      |
| seizure_date             | timestamptz(6) | 查扣日期                                                    | optional     |                                      |
| seizure_days_single      | int4           | 单部门查扣天数                                              | optional     |                                      |
| penalty_single           | numeric        | 单部门罚款金额                                              | optional     |                                      |
| seizure_days_multiple    | int4           | 联罚查扣天数                                                | optional     |                                      |
| penalty_multiple         | numeric        | 联罚罚款金额                                                | optional     |                                      |
| closed_days              | int4           | 企业停业天数                                                | optional     |                                      |
| penalty                  | numeric        | 企业罚款金额                                                | optional     |                                      |
| rectification_situation  | text           | 企业整改情况                                                | optional     |                                      |
| region_code              | text           | 区域编码                                                    | optional     |                                      |
| user_id                  | text           | 用户ID（录入数据的用户）                                    | optional     |                                      |
| user_role                | int4           | 用户角色（3：城管，4：交警，5：交通）                       | optional     |                                      |
| seizure_reason           | text           | 查扣原因                                                    | optional     |                                      |
| other_reason             | text           | 其他原因                                                    | optional     |                                      |
| is_joint                 | boolean        | 是否联罚联放                                                | optional     |                                      |
| transfer_department      | text           | 移交部门                                                    | optional     |                                      |
| joint_exit_id            | text           | 联罚联放附件ID                                              | optional     |                                      |
| final_state              | int4           | 最终状态 1：完成，2：待移交，3：待联罚                      | optional     |                                      |
| created_at               | timestamptz(6) | 创建时间                                                    | required     |                                      |
| created_by               | text           | 创建人                                                      | required     | **system_user表**的user_id           |
| updated_at               | timestamptz(6) | 修改时间                                                    | optional     |                                      |
| updated_by               | text           | 修改人                                                      | optional     | **system_user表**的user_id           |
| deleted_at               | timestamptz(6) | 删除时间                                                    | optional     |                                      |
| deleted_by               | text           | 删除人                                                      | optional     | **system_user表**的user_id           |

### joint_penalty_release_sub 联罚联放附属表

**根据以下数据表整合而成**

工程车:

[ENFORCEMENTINFO_PUSH 联罚联放附属表 ](#egn_ENFORCEMENTINFO)

| Name                     | Type           | Description                           | **Required** | default                    |
| ------------------------ | -------------- | ------------------------------------- | ------------ | -------------------------- |
| id                       | int8           | 按指定方法生成                        | required     | 主键                       |
| joint_penalty_release_id | text           | 联罚联放主表id                        | required     |                            |
| transfer_role            | int4           | 移交角色（3：城管，4：交警，5：交通） | optional     |                            |
| region_code              | text           | 区域编码                              | optional     |                            |
| is_deal                  | boolean        | 是否处理                              | optional     |                            |
| processor                | text           | 处理人                                | optional     |                            |
| seizure_reason           | text           | 查扣原因                              | optional     |                            |
| other_reason             | text           | 其他原因                              | optional     |                            |
| seizure_date             | timestamptz(6) | 查扣日期                              | optional     |                            |
| seizure_days_multiple    | int4           | 联罚查扣天数                          | optional     |                            |
| penalty_multiple         | numeric        | 联罚罚款金额                          | optional     |                            |
| closed_days_ent          | int4           | 企业停业天数                          | optional     |                            |
| penalty_ent              | numeric        | 企业罚款金额                          | optional     |                            |
| rectification_situation  | text           | 企业整改情况                          | optional     |                            |
| transfer_person          | text           | 移交人                                | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                              | required     |                            |
| created_by               | text           | 创建人                                | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                              | optional     |                            |
| updated_by               | text           | 修改人                                | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                              | optional     |                            |
| deleted_by               | text           | 删除人                                | optional     | **system_user表**的user_id |

### tem_work_num_download 临时工号牌下载记录表

**根据以下数据表整合而成**

工程车:

VEH_TMPSVNUM_DOWNLOAD_LOG 临时工号牌下载记录表

| Name        | Type           | Description                              | **Required** | default                    |
| ----------- | -------------- | ---------------------------------------- | ------------ | -------------------------- |
| id          | int8           | 按指定方法生成                           | required     | 主键                       |
| vehicle_id  | text           | **vehicle_info** 车辆信息表 的vehicle_id | optional     |                            |
| valid_from  | timestamptz(6) | 有效期起始                               | optional     |                            |
| valid_until | timestamptz(6) | 有效期截止                               | optional     |                            |
| is_deleted  | boolean        | 是否删除                                 | optional     |                            |
| created_at  | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by  | text           | 创建人                                   | required     | **system_user表**的user_id |
| updated_at  | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by  | text           | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at  | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by  | text           | 删除人                                   | optional     | **system_user表**的user_id |

### veh_egn_change_review 工程运输车辆单位变更审核表

**根据以下数据表整合而成**

工程车:

[VEHTEAMCHANGEDETAIL 车辆单位变更审核表 ](#egn_VEHTEAMCHANGEDETAIL)

| Name                     | Type           | Description                                                  | **Required** | default                    |
| ------------------------ | -------------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                       | int8           | 按指定方法生成                                               | required     | 主键                       |
| veh_egn_change_review_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| vehicle_id               | text           | **vehicle_info** 车辆信息表 的vehicle_id                     | required     |                            |
| target_enterprise_id     | text           | 目标企业ID                                                   | optional     |                            |
| remarks                  | text           | 备注                                                         | optional     |                            |
| upload_document_src      | text           | 上传证明文件                                                 | optional     |                            |
| review_status            | int4           | 审核状态  0.未审核  1.审核通过  2.退回                       | optional     |                            |
| change_type              | int4           | 变更类型  1.本地区变更  2.跨区变更                           | optional     |                            |
| review_step              | int4           | 审核步骤  变更类型1：1.交警提交申请  2.渣土办同意  变更类型2：1.A县交警提交申请  2.A县渣土办同意  3.B县交警同意  4.县渣土办同意 |              |                            |
| original_area            | text           | 车辆原所属地区                                               | optional     |                            |
| changed_area             | text           | 变更地区                                                     | optional     |                            |
| changed_self_number      | text           | 变更后自编号                                                 | optional     |                            |
| is_deleted               | boolean        | 是否删除                                                     | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by               | text           | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by               | text           | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by               | text           | 删除人                                                       | optional     | **system_user表**的user_id |

### vehicle_inspection_pic 验车图片表

**根据以下数据表整合而成**

工程车:

INFO_SMARTCAR_CKCAR_CKIMG 验车图片表

| Name                      | Type           | Description                                     | **Required** | default                    |
| ------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | int8           | 按指定方法生成                                  | required     | 主键                       |
| vehicle_inspection_pic_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| img_addr                  | text           | 图片地址                                        | optional     |                            |
| img_type                  | text           | 图片类型                                        | optional     |                            |
| is_deleted                | boolean        | 是否删除                                        | optional     |                            |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                | text           | 删除人                                          | optional     | **system_user表**的user_id |

### vehicle_inspection_record 验车记录表

**根据以下数据表整合而成**

工程车:

INFO_SMARTCAR_CKCAR 验车记录表

| Name                         | Type           | Description                                     | **Required** | default                    |
| ---------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | int8           | 按指定方法生成                                  | required     | 主键                       |
| vehicle_inspection_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                   | text           | 车辆信息表的vehicle_id                          | optional     |                            |
| terminal_id                  | text           | 终端ID                                          | optional     |                            |
| vehicle_inspection_pic_id    | text           | 验车图片表的id                                  | optional     |                            |
| axed_num                     | integer        | 轴数                                            | optional     |                            |
| cover_thickness              | numeric        | 车厢密闭顶盖厚度                                | optional     |                            |
| carriage_length              | numeric        | 车厢长度                                        | optional     |                            |
| carriage_width               | numeric        | 车厢宽度                                        | optional     |                            |
| carriage_height              | numeric        | 车厢高度                                        | optional     |                            |
| tank_length                  | numeric        | 罐体长度                                        | optional     |                            |
| tank_volume                  | numeric        | 罐体容积                                        | optional     |                            |
| veh_length                   | numeric        | 车辆长度                                        | optional     |                            |
| operating_paltform           | text           | 操作平台                                        | optional     |                            |
| police_num                   | text           | 警号                                            | optional     |                            |
| remarks                      | text           | 备注                                            | optional     |                            |
| veh_type                     | text           | 车辆类型                                        | optional     |                            |
| is_pass_check                | boolean        | 是否通过校验                                    | optional     |                            |
| is_smart_terminal            | boolean        | 是否智能终端                                    | optional     |                            |
| is_airtight_device           | boolean        | 是否具有密闭装置                                | optional     |                            |
| is_generate_pre_num          | boolean        | 是否生成预编号                                  | optional     |                            |
| is_country_five              | boolean        | 是否国V及以上排放标准                           | optional     |                            |
| is_install_gps               | boolean        | 是否安装GPS                                     | optional     |                            |
| is_u_shaped_cargo            | boolean        | 是否带全密闭顶盖U型货箱                         | optional     |                            |
| is_pass                      | boolean        | 是否通过                                        | optional     |                            |
| is_deleted                   | boolean        | 是否删除                                        | optional     |                            |
| created_at                   | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                   | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                   | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                   | text           | 删除人                                          | optional     | **system_user表**的user_id |

### veh_operation_his 车辆操作历史记录表

**根据以下数据表整合而成**

工程车:

VEHICLEOPERATE 车辆操作历史记录表

| Name                 | Type           | Description                                     | **Required** | default                    |
| -------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | int8           | 按指定方法生成                                  | required     | 主键                       |
| veh_operation_his_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | optional     | 联合主键                   |
| vehicle_id           | text           | **veh_info** 车辆信息表 的vehicle_id            | optional     |                            |
| remarks              | text           | 备注                                            | optional     |                            |
| operation_type       | integer        | 操作类型  1.添加  2.删除                        | optional     |                            |
| operation_person     | text           | 操作人                                          | optional     |                            |
| audit_status         | integer        | 审核状态  0.未审批  1.已审批                    | optional     |                            |
| audit_person         | text           | 审核人                                          | optional     |                            |
| area                 | text           | 地区                                            | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text           | 删除人                                          | optional     | **system_user表**的user_id |

## 驾驶员信息(公共)

### driver_info 驾驶员信息表

**根据以下数据表提取出公用部分**

重点车辆：
[VD_DRIVER 驾驶员信息](#major_VD_DRIVER)

营运车辆：
[INFO_EMPLOYEES 驾驶员信息 ](#oper_INFO_EMPLOYEES)

工程车:

[DRIVERSINFO 驾驶员信息表 ](#egn_DRIVERSINFO)

| Name               | Type           | Description                                     | **Required** | default                                                      |
| ------------------ | -------------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                 | int8           | 按指定方法生成                                  | required     | 主键                                                         |
| driver_id          | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| enterprise_id      | text           | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id                         |
| department_id      | text           | 所在部门id                                      | optional     | **department 部门信息表**                                    |
| driver_identity_id | text           | 驾驶员身份验证信息ID                            | optional     | **driver_identity表**的driver_identity_id                    |
| driver_name        | text           | 驾驶员姓名                                      | optional     |                                                              |
| telephone          | text           | 手机号码                                        | optional     |                                                              |
| sex                | int4           | 性别                                            | optional     | **性别**字典                                                 |
| files_number       | text           | 档案编号(后6位)                                 | optional     |                                                              |
| contact_address    | text           | 联系地址                                        | optional     |                                                              |
| mailing_address    | text           | 邮寄地址                                        | optional     |                                                              |
| is_active          | boolean        | 是否激活（默认激活）                            | optional     | 用于标志驾驶员资料是否处于确定状态。未确定状态的驾驶员信息在系统上除驾驶员管理外的功能中都查不到。 |
| is_manual_input    | boolean        | 是否手动录入                                    | optional     | 驾驶员资料分为使用身份证读卡器读取身份证自动录入资料和手动填写资料 |
| is_input           | int4           | 资料完整性(1完整，2不完整)                      | optional     |                                                              |
| input_at           | timestamptz(6) | 录入时间                                        | optional     |                                                              |
| input_by           | text           | 录入人                                          | optional     | **system_user表**的user_id                                   |
| is_check_data      | boolean        | 是否上传公安内网获取6合1数据                    | optional     | 该字段代表是否用于校验驾驶员信息，未正式录入系统，但会同步到公安内网，用于查询驾驶员的违章数据。 |
| check_at           | timestamptz(6) | 检验时间                                        | optional     |                                                              |
| check_by           | text           | 校验人                                          | optional     | **system_user表**的user_id                                   |
| remark_in          | text           | 驾驶员信息同步内网反馈信息                      | optional     | 驾驶员信息同步内网反馈信息。驾驶员信息同步到公安内网后内网的反馈内容，如档案编号填写错误会反馈档案编号后六位不正确 |
| update_time_in     | timestamptz(6) | 内网更新时间                                    | optional     |                                                              |
| is_check_sms       | boolean        | 是否通过短信验证                                | optional     |                                                              |
| remarks            | text           | 备注                                            | optional     |                                                              |
| is_deleted         | boolean        | 是否删除                                        | optional     |                                                              |
| is_blacklist       | boolean        | 是否黑名单(重点车)                              | optional     | false                                                        |
| blacklist_deadline | timestamptz(6) | 黑名单截止日期                                  | optional     |                                                              |
| is_clear_out       | boolean        | 是否清退                                        | optional     |                                                              |
| clear_out_time     | timestamptz(6) | 清退时间                                        | optional     |                                                              |
| created_at         | timestamptz(6) | 创建时间                                        | required     |                                                              |
| created_by         | text           | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at         | timestamptz(6) | 修改时间                                        | optional     |                                                              |
| updated_by         | text           | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at         | timestamptz(6) | 删除时间                                        | optional     |                                                              |
| deleted_by         | text           | 删除人                                          | optional     | **system_user表**的user_id                                   |

### driver_identity 驾驶员身份信息拓展表(新表)

**新表,把驾驶员身份信息提取出来**

重点车辆：
[VD_DRIVER 驾驶员信息](#major_VD_DRIVER)

营运车辆：
[INFO_EMPLOYEES 驾驶员信息 ](#oper_INFO_EMPLOYEES)

工程车:

[DRIVERSINFO 驾驶员信息表 ](#egn_DRIVERSINFO)

| Name                             | Type           | Description                                     | **Required** | default                    |
| -------------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                               | int8           | 按指定方法生成                                  | required     | 主键                       |
| driver_identity_id               | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| driver_id                        | text           | 驾驶员ID                                        | optional     |                            |
| id_card_num                      | text           | 身份证号码                                      | optional     |                            |
| id_card_birthday                 | timestamptz(6) | 身份证出生日期                                  | optional     |                            |
| id_card_sign_government          | text           | 身份证签发机关                                  | optional     |                            |
| id_card_nation                   | text           | 身份证民族                                      | optional     |                            |
| id_card_start_date               | timestamptz(6) | 身份证有效起始日期                              | optional     |                            |
| id_card_end_date                 | timestamptz(6) | 身份证有效截止日期                              | optional     |                            |
| id_card_front_pic                | text           | 身份证正面照，云存储地址                        | optional     |                            |
| id_card_back_pic                 | text           | 身份证背面照，云存储地址                        | optional     |                            |
| id_card_address                  | text           | 身份证住址                                      | optional     |                            |
| driver_holding_id_photo          | text           | 驾驶员手持身份证照片,云储存系统返回的路径       | optional     |                            |
| driver_photo                     | text           | 驾驶员的正面照,云储存系统返回的路径             | optional     |                            |
| driver_signature                 | text           | 驾驶员签名,云储存系统返回的路径                 | optional     |                            |
| occupational_number              | text           | 从业资格证号码                                  | optional     |                            |
| occupational_expire_date         | timestamptz(6) | 从业资格证有效期至                              | optional     |                            |
| occupational_issuing_authority   | text           | 从业资格证发证机构                              | optional     |                            |
| labor_contract                   | text[]         | 劳动合同,云储存系统返回的完整劳动合同的图片路径 | optional     |                            |
| driver_license_pic               | text           | 驾驶员驾驶证,云储存系统返回的路径               | optional     |                            |
| driver_license_issuing_authority | text           | 驾驶证发证机关                                  | optional     |                            |
| annual_review_date               | timestamptz(6) | 年审日期（六合一）                              | optional     |                            |
| renewal_date                     | timestamptz(6) | 换证日期（六合一）                              | optional     |                            |
| accumulatived_points             | numeric        | 累计积分（六合一）                              | optional     |                            |
| sorting_date                     | timestamptz(6) | 清分日期（六合一）                              | optional     |                            |
| quasi_driving_models             | int4           | 准驾车型（六合一）                              | optional     | **准驾车型**字典           |
| driver_license_province_id       | text           | 驾驶证发证所在地的省份ID                        | optional     | 省份表                     |
| driver_license_city_id           | text           | 驾驶证发证所在地的城市ID                        | optional     | 城市表                     |
| driver_license_district_id       | text           | 驾驶证发证所在地的区域ID                        | optional     | 区域表                     |
| driver_license_status            | int4           | 驾驶证状态                                      | optional     | **驾驶证状态**字典         |
| driver_license_issue_date        | timestamptz(6) | 驾驶证初次领证日期                              | optional     |                            |
| working_time                     | timestamptz(6) | 从业时间                                        | optional     |                            |
| created_at                       | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                       | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                       | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                       | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                       | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                       | text           | 删除人                                          | optional     | **system_user表**的user_id |
| is_deleted                       | boolean        | 是否删除                                        | optional     |                            |

## 驾驶员信息(营运车)

### driver_identity_info_report 驾驶员身份信息采集上报

**根据以下数据表整合而成**

营运车:

[MON_DRVCARD_OP 驾驶员身份信息采集上报](#oper_MON_DRVCARD_OP)

| Name                     | Type        | Description                                     | **Required** | default                    |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                       |
| driver_identity_info_report_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id               | text        | 车辆信息表 的vehicle_id                         | optional     |                            |
| ic_card_status           | integer     | IC状态(从业资格证IC卡插入,从业资格证IC卡拔出)   | optional     | **从业资格证IC卡**字典     |
| operation_time           | timestamptz(6) | 操作时间                                        | optional     |                            |
| driver_name              | text        | 驾驶员姓名                                      | optional     |                            |
| license_number           | text        | 证件号码                                        | optional     |                            |
| imel                     | text        | 终端IMEI                                        | optional     | 国际移动设备标识别码       |
| ic_card_reading_result   | text        | IC卡读取结果                                    | optional     | **IC卡读卡**字典           |
| occupational_number      | text        | 从业资格证编码                                  | optional     |                            |
| driver_license_name      | text        | 发证机构名称                                    | optional     |                            |
| license_expire_date      | timestamptz(6) | 证件有效期                                      | optional     |                            |
| registration_time        | timestamptz(6) | 登记时间                                        | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id |

## 驾驶员信息(工程车)

### driver_license_overdue_his 驾驶证年检逾期历史表(新表)

**新表,把驾驶证年检登记表改为年检逾期历史表**

工程车:

[DRIVERYCKMARK  驾驶证年检登记](#egn_DRIVERYCKMARK)

| Name                          | Type           | Description                                     | **Required** | default                    |
| ----------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | int8           | 按指定方法生成                                  | required     | 主键                       |
| driver_license_overdue_his_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | optional     | 联合主键                   |
| driver_id                     | text           | 驾驶员ID                                        | optional     |                            |
| start_time                    | timestamptz(6) | 逾期开始时间                                    | optional     |                            |
| end_time                      | timestamptz(6) | 逾期结束时间                                    | optional     |                            |
| reason                        | text           | 逾期原因                                        | optional     |                            |
| remarks                       | text           | 备注                                            | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                    | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                    | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                    | text           | 删除人                                          | optional     | **system_user表**的user_id |



## 驾驶员信息(重点车)

### vehicle_driver_binding  车辆驾驶员绑定表

**根据以下数据表整合而成**

重点车:

| Name                      | Type           | Description                                                  | **Required** | default                        |
| ------------------------- | -------------- | ------------------------------------------------------------ | ------------ | ------------------------------ |
| id                        | int8           | 按指定方法生成                                               | required     | 主键                           |
| vehicle_driver_binding_id | text           | 车辆驾驶员绑定外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                       |
| driver_id                 | text           | 驾驶员id                                                     | optional     | **driver_info表**的driver_id   |
| vehicle_id                | text           | 车辆id                                                       | optional     | **vehicle_info表**的vehicle_id |
| remarks                   | text           | 备注                                                         | optional     |                                |
| created_at                | timestamptz(6) | 创建时间                                                     | required     |                                |
| created_by                | text           | 创建人                                                       | required     | **system_user表**的user_id     |
| updated_at                | timestamptz(6) | 修改时间                                                     | optional     |                                |
| updated_by                | text           | 修改人                                                       | optional     | **system_user表**的user_id     |
| deleted_at                | timestamptz(6) | 删除时间                                                     | optional     |                                |
| deleted_by                | text           | 删除人                                                       | optional     | **system_user表**的user_id     |
| is_deleted                | boolean        | 是否删除                                                     | optional     |                                |

## 渣土车推荐目录

### vehicle_brand_record 车辆品牌备案 

**根据以下数据表整合而成**

重点车:

[NEWZTC_SELLERRECORD 销售商备案 ](#major_NEWZTC_SELLERRECORD)

| Name                                  | Type           | Description                                     | **Required** | default                                                    |
| ------------------------------------- | -------------- | ----------------------------------------------- | ------------ | ---------------------------------------------------------- |
| id                                    | int8           | 按指定方法生成                                  | required     | 主键                                                       |
| vehicle_brand_record_id               | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                   |
| seller                                | text           | 销售商                                          | optional     | **enterprise_info表**的enterprise_id                       |
| business_license                      | text           | 工商营业执照                                    | optional     |                                                            |
| organization_code                     | text           | 组织机构代码证                                  | optional     |                                                            |
| legal_representative_id_card          | text           | 法人代表身份证                                  | optional     |                                                            |
| entrusted_agent_id_card               | text           | 委托代理人身份证                                | optional     |                                                            |
| ministry_production_sale_permit       | text           | 工信部汽车生产销售许可证明                      | optional     |                                                            |
| industry_self_discipline_commitment   | text           | 行业自律承诺书                                  | optional     |                                                            |
| manufacturer_production_sale_permit   | text           | 汽车厂家汽车销售许可证明                        | optional     |                                                            |
| repair_service_station_name           | text           | 维修服务服务站名称                              | optional     |                                                            |
| repair_service_station_address        | text           | 维修服务服务站地址                              | optional     |                                                            |
| repair_service_station_material       | text           | 维修服务服务站相关材料                          | optional     |                                                            |
| service_staff_name                    | text           | 服务人员名称                                    | optional     |                                                            |
| service_staff_phone                   | text           | 服务人员电话                                    | optional     |                                                            |
| service_staff_id                      | text           | 服务人员身份证                                  | optional     |                                                            |
| service_staff_material                | text           | 服务人员相关材料                                | optional     |                                                            |
| other_material                        | text           | 其他相关材料                                    | optional     |                                                            |
| remarks                               | text           | 备注                                            | optional     |                                                            |
| registration_time                     | timestamptz(6) | 登记日期                                        | optional     |                                                            |
| registration_user                     | text           | 登记用户                                        | optional     | **system_user表**的user_id                                 |
| is_verify                             | boolean        | 是否核查                                        | optional     |                                                            |
| cause                                 | text           | 原因                                            | optional     |                                                            |
| catalog_id                            | text           | 汽车品牌(简称)                                  | optional     | new_engin_veh_recommend_catalog 新型工程运输车推荐目录的id |
| platform_docking_technology_agreement | text           | 平台对接技术协议                                | optional     |                                                            |
| is_deleted                            | boolean        | 是否删除                                        | optional     |                                                            |
| created_at                            | timestamptz(6) | 创建时间                                        | required     |                                                            |
| created_by                            | text           | 创建人                                          | required     | **system_user表**的user_id                                 |
| updated_at                            | timestamptz(6) | 修改时间                                        | optional     |                                                            |
| updated_by                            | text           | 修改人                                          | optional     | **system_user表**的user_id                                 |
| deleted_at                            | timestamptz(6) | 删除时间                                        | optional     |                                                            |
| deleted_by                            | text           | 删除人                                          | optional     | **system_user表**的user_id                                 |

### vehicle_model_record 车辆型号备案

**根据以下数据表整合而成**

重点车:

[NEWZTC_CARRECORD 汽车销售备案 ](#major_NEWZTC_CARRECORD)

| Name                    | Type           | Description                                     | **Required** | default                                                    |
| ----------------------- | -------------- | ----------------------------------------------- | ------------ | ---------------------------------------------------------- |
| id                      | int8           | 按指定方法生成                                  | required     | 主键                                                       |
| vehicle_model_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                   |
| seller                  | text           | 销售商                                          | optional     | **enterprise_info表**的enterprise_id                       |
| vehicle_model           | text           | 汽车型号                                        | optional     |                                                            |
| vehicle_picture         | text[]         | 汽车相关图片的路径                              | optional     |                                                            |
| vehicle_parameter       | text           | 汽车参数                                        | optional     |                                                            |
| reference_price         | text           | 参考报价                                        | optional     |                                                            |
| other_material          | text           | 其他相关材料                                    | optional     |                                                            |
| remarks                 | text           | 备注                                            | optional     |                                                            |
| registration_time       | timestamptz(6) | 登记日期                                        | optional     |                                                            |
| registration_user       | text           | 登记用户                                        | optional     | **system_user表**的user_id                                 |
| cause                   | text           | 原因                                            | optional     |                                                            |
| catalog_id              | text           | 汽车品牌(简称)                                  | optional     | new_engin_veh_recommend_catalog 新型工程运输车推荐目录的id |
| axis_type               | int4           | 轴数类型                                        | optional     | **轴数类型**字典                                           |
| transport_volume        | text           | 运输方量                                        | optional     |                                                            |
| is_review               | boolean        | 是否审核                                        | optional     |                                                            |
| is_deleted              | boolean        | 是否删除                                        | optional     |                                                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                                                            |
| created_by              | text           | 创建人                                          | required     | **system_user表**的user_id                                 |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                                                            |
| updated_by              | text           | 修改人                                          | optional     | **system_user表**的user_id                                 |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                                                            |
| deleted_by              | text           | 删除人                                          | optional     | **system_user表**的user_id                                 |

### egn_recommend_catalog 工程车品牌字典

**根据以下数据表整合而成**

重点车:

[NEWZTC_BRAND_CATALOG 新型渣土车推荐目录 ](#major_NEWZTC_BRAND_CATALOG)

| Name                     | Type           | Description                                     | **Required** | default                    |
| ------------------------ | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | int8           | 按指定方法生成                                  | required     | 主键                       |
| egn_recommend_catalog_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| brand_name               | text           | 品牌简称                                        | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text           | 删除人                                          | optional     | **system_user表**的user_id |

### veh_egn_sale_order 工程车销售订单

**根据以下数据表整合而成**

重点车:

[NEWZTC_SALE_BILL 新型渣土车销售订单 ](#major_NEWZTC_SALE_BILL)

工程车：

[NEWZTC_SALE_BILL_DETAIL 销售订单表 ](#egn_NEWZTC_SALE_BILL_DETAIL )

| Name                     | Type           | Description                                     | **Required** | default                                  |
| ------------------------ | -------------- | ----------------------------------------------- | ------------ | ---------------------------------------- |
| id                       | int8           | 按指定方法生成                                  | required     | 主键                                     |
| veh_egn_sale_order_id    | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                 |
| customer_name            | text           | 顾客企业id                                      | optional     |                                          |
| customer_id              | text           | 购车企业或个人姓名                              | optional     | **enterprise_info表**的enterprise_id     |
| id_card_num              | text           | 身份证号                                        | optional     |                                          |
| contact_phone            | text           | 联系电话                                        | optional     |                                          |
| notification             | text           | 告知书                                          | optional     |                                          |
| vehicle_sale_contract    | text[]         | 汽车销售合同(销售合同照片地址)                  | optional     |                                          |
| vehicle_deposit_contract | text[]         | 订金合同(订金合同照片地址)                      | optional     |                                          |
| commitment               | text           | 承诺书                                          | optional     |                                          |
| book_date                | timestamptz(6) | 订车日期                                        | optional     |                                          |
| registration_time        | timestamptz(6) | 登记时间                                        | optional     |                                          |
| belonging_enterprise     | text           | 所属单位                                        | optional     |                                          |
| id_card_picture          | text           | 身份证(身份证照片路径)                          | optional     |                                          |
| progress                 | numeric        | 进度                                            | optional     |                                          |
| vehicle_sale_record_id   | text           | 汽车型号                                        | optional     | **vehicle_sale_record 汽车销售备案**的id |
| book_vehicle_number      | int4           | 订购车辆数                                      | optional     |                                          |
| vehicle_length           | numeric        | 整车长度                                        | optional     |                                          |
| vehicle_width            | numeric        | 整车宽度                                        | optional     |                                          |
| vehicle_height           | numeric        | 整车高度                                        | optional     |                                          |
| tipper_length            | numeric        | 自卸车长度                                      | optional     |                                          |
| tipper_width             | numeric        | 自卸车宽度                                      | optional     |                                          |
| tipper_height            | numeric        | 自卸车高度                                      | optional     |                                          |
| tank_truck_marked_volume | numeric        | 罐车标注容积                                    | optional     |                                          |
| tank_truck_actual_volume | numeric        | 罐车实际搅动容积                                | optional     |                                          |
| carriage_sealing_device  | text           | 车厢密闭装置                                    | optional     |                                          |
| u_shaped_cargo_box       | text           | U型货箱                                         | optional     |                                          |
| is_other_vehicle_model   | boolean        | 是否其他汽车型号                                | optional     |                                          |
| is_completed             | boolean        | 是否完成                                        | optional     |                                          |
| is_deleted               | boolean        | 是否删除                                        | optional     |                                          |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                                          |
| created_by               | text           | 创建人                                          | required     | **system_user表**的user_id               |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                                          |
| updated_by               | text           | 修改人                                          | optional     | **system_user表**的user_id               |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                                          |
| deleted_by               | text           | 删除人                                          | optional     | **system_user表**的user_id               |

### veh_egn_sale_order_detail 工程车销售订单明细

**根据以下数据表整合而成**

重点车:

[NEWZTC_SALE_BILL_DETAIL 新型渣土车销售订单明细 ](#major_NEWZTC_SALE_BILL_DETAIL)

| Name                          | Type           | Description                                     | **Required** | default                                                      |
| ----------------------------- | -------------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                            | int8           | 按指定方法生成                                  | required     | 主键                                                         |
| veh_egn_sale_order_detail_id  | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| order_id                      | text           | 订单号                                          | optional     | **engin_veh_sale_order 新型工程运输车销售订单**的engin_veh_sale_order_id |
| vehicle_identification_number | text           | 车架号(后6位)                                   | optional     | 车辆识别代号vin,如D02133                                     |
| certificate                   | text           | 合格证                                          | optional     |                                                              |
| terminal_proof                | text           | 终端证明                                        | optional     |                                                              |
| driver_license_pic            | text           | 行驶证(照片)                                    | optional     |                                                              |
| license_plate_number          | text           | 车牌号码                                        | optional     |                                                              |
| license_plate_color           | int4           | 车牌颜色                                        | optional     | **车牌颜色**字典                                             |
| license_plate_type            | int4           | 号牌种类                                        | optional     | **号牌种类**字典                                             |
| arrive_qz_date                | timestamptz(6) | 抵达泉州日期                                    | optional     |                                                              |
| first_registration_date       | timestamptz(6) | 初次登记日期                                    | optional     |                                                              |
| registration_date             | timestamptz(6) | 登记时间                                        | optional     |                                                              |
| step                          | int4           | 步骤(销售订单登记进度（2.到车登记 3.上牌登记）) | optional     |                                                              |
| vehicle_photo                 | text           | 车辆照片                                        | optional     |                                                              |
| seller_preview_number         | text           | 供应商销售预编号                                | optional     |                                                              |
| is_deleted                    | boolean        | 是否删除                                        | optional     |                                                              |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                                                              |
| created_by                    | text           | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                                                              |
| updated_by                    | text           | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                                                              |
| deleted_by                    | text           | 删除人                                          | optional     | **system_user表**的user_id                                   |

### seller_rating_record 销售商评分记录

**根据以下数据表整合而成**

重点车:

[NEWZTC_SALE_KF_LOG 销售商评分记录](#major_NEWZTC_SALE_KF_LOG)

| Name                    | Type           | Description                                     | **Required** | default                              |
| ----------------------- | -------------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                      | int8           | 按指定方法生成                                  | required     | 主键                                 |
| seller_rating_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| demerit_obj             | text           | 扣分对象                                        | optional     | **enterprise_info表**的enterprise_id |
| demerit_points          | numeric        | 扣分分值                                        | optional     |                                      |
| demerit_reason          | text           | 扣分原因                                        | optional     |                                      |
| operator                | text           | 操作人                                          | optional     | **system_user表**的user_id           |
| operation_time          | timestamptz(6) | 操作时间                                        | optional     |                                      |
| is_deleted              | boolean        | 是否删除                                        | optional     |                                      |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                                      |
| created_by              | text           | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                                      |
| updated_by              | text           | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                                      |
| deleted_by              | text           | 删除人                                          | optional     | **system_user表**的user_id           |

### ent_purchase_intention 企业购车意向

**根据以下数据表整合而成**

重点车:

[ZTC_GCYX 渣土车购车意向](#major_ZTC_GCYX)

| Name                      | Type           | Description                                     | **Required** | default                              |
| ------------------------- | -------------- | ----------------------------------------------- | ------------ | ------------------------------------ |
| id                        | int8           | 按指定方法生成                                  | required     | 主键                                 |
| ent_purchase_intention_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| supplier                  | text           | 供应商                                          | optional     | **enterprise_info表**的enterprise_id |
| customer_name             | text           | 购车用户姓名                                    | optional     |                                      |
| customer_phone            | text           | 购车用户电话                                    | optional     |                                      |
| owner_enterprise          | text           | 所有人企业                                      | optional     | **enterprise_info表**的enterprise_id |
| province_id               | text           | 所有人所在省                                    | optional     | **省份表**province_id                |
| city_id                   | text           | 所有人所在市                                    | optional     | **城市表**city_id                    |
| district_id               | text           | 所有人所在县                                    | optional     | **区域表**district_id                |
| capacigy_application      | text           | 运力申请                                        | optional     |                                      |
| brand_model               | text           | 品牌型号                                        | optional     |                                      |
| vehicle_purchase          | int4           | 购车数量                                        | optional     |                                      |
| code                      | text           | 编码                                            | optional     |                                      |
| registration_time         | timestamptz(6) | 登记日期                                        | optional     |                                      |
| registration_user         | text           | 登记用户                                        | optional     | **system_user表**的user_id           |
| review                    | text           | 审核                                            | optional     |                                      |
| remarks                   | text           | 备注                                            | optional     |                                      |
| is_deleted                | boolean        | 是否删除                                        | optional     |                                      |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                                      |
| created_by                | text           | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                                      |
| updated_by                | text           | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                                      |
| deleted_by                | text           | 删除人                                          | optional     | **system_user表**的user_id           |

### sale_order_whitelist 销售订单白名单

**根据以下数据表整合而成**

工程车:

SALEBILLWHITELIST 销售订单白名单

| Name                      | Type           | Description                                     | **Required** | default                    |
| ------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | int8           | 按指定方法生成                                  | required     | 主键                       |
| vehicle_passing_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                | text           | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| regis_time                | timestamptz(6) | 登记时间                                        | optional     |                            |
| is_deleted                | boolean        | 是否删除                                        | optional     |                            |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                | text           | 删除人                                          | optional     | **system_user表**的user_id |

## 工地进出报警(抓拍系统和电子围栏)

### construction_info 工地信息表

**根据以下数据表整合而成**

工程车:

 CONSTRUCTIONINFO工地信息表

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| construction_info_id         | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| project_name                 | text        | 项目名称                                        | optional     |                            |
| project_code                 | text        | 工程项目代码                                    | optional     |                            |
| contract_start_date          | timestamptz(6) | 合同开工日期                                    | optional     |                            |
| contract_completion_date     | timestamptz(6) | 合同竣工日期                                    | optional     |                            |
| build_agency                 | text        | 建设单位                                        | optional     |                            |
| project_manager              | text        | 项目负责人                                      | optional     |                            |
| area_id                      | text        | 地区                                            | optional     |                            |
| address                      | text        | 地址                                            | optional     |                            |
| map_coordinate               | point       | 地图经纬度                                      | optional     |                            |
| construction_agency          | text        | 施工单位                                        | optional     |                            |
| supervision_agency           | text        | 监理单位                                        | optional     |                            |
| device_installation_agency   | text        | 设备安装单位                                    | optional     |                            |
| device_installation_location | text        | 设备安装地点                                    | optional     |                            |
| service_contract_number      | text        | 服务合同编号                                    | optional     |                            |
| contract_signing_time        | timestamptz(6) | 合同签署时间                                    | optional     |                            |
| service_activation_date      | timestamptz(6) | 服务启用日期                                    | optional     |                            |
| proof_valid_date             | timestamptz(6) | 证明有效期                                      | optional     |                            |
| coordinate                   | point       | 经纬度                                          | optional     |                            |
| installation_time            | timestamptz(6) | 安装时间                                        | optional     |                            |
| project_name_tag             | text        | 项目名称标签                                    | optional     |                            |
| import_export_quantity       | text        | 进出口数量                                      | optional     |                            |
| deactivation_reason          | text        | 停用原因                                        | optional     |                            |
| uninstalled_reason           | text        | 未安装原因                                      | optional     |                            |
| proof_valid_date_until       | timestamptz(6) | 证明起始有效期                                  | optional     |                            |
| type                         | integer     | 类型（1.工地 2.矿山）                           | optional     |                            |
| contact_person               | text        | 联系人                                          | optional     |                            |
| contact_phone                | text        | 联系电话                                        | optional     |                            |
| finished                     | integer     | 是否完工（0.未完工 1.已完工 2.已停用）          | optional     |                            |
| is_city_directly             | text        | 是否市直属                                      | optional     |                            |
| is_installation              | text        | 是否安装（0.未申请 1.已安装 2.已申请未安装）    | optional     |                            |
| is_deleted                   | boolean     | 是否删除                                        | optional     |                            |
| registration_time            | timestamptz(6) | 登记时间                                        | optional     |                            |
| created_at                   | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### vehicle_passing_record 抓拍系统过车记录表

**根据以下数据表整合而成**

工程车:

[ORIG_FILEVEHPASS 过车记录表](#egn_ORIG_FILEVEHPASS)

| Name                      | Type           | Description                                     | **Required** | default                    |
| ------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | int8           | 按指定方法生成                                  | required     | 主键                       |
| vehicle_passing_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                | text           | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| license_plate_number      | text           | 车牌号                                          | optional     |                            |
| passing_time              | timestamptz(6) | 过车时间                                        | optional     |                            |
| registration_time         | timestamptz(6) | 登记时间                                        | optional     |                            |
| camera_id                 | text           | 摄像头ID                                        | optional     |                            |
| pass_type                 | text           | 过车类型                                        | optional     |                            |
| picture_url               | text           | 图片地址                                        | optional     |                            |
| review_status             | text           | 审核状态                                        | optional     |                            |
| vehicle_type              | text           | 车辆类型                                        | optional     |                            |
| is_online                 | boolean        | 是否在线                                        | optional     |                            |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                | text           | 删除人                                          | optional     | **system_user表**的user_id |

### construction_passing_alarm 工地进出报警表(新)

**抓拍系统报警和工地过车报警都放在工地进出报警表里**

工程车:

ORIG_FILEVEHPASSALARMMSG 抓拍系统过车报警表

INFO_ELECTRIC_CRAWL_ALARM_MSG 电子围栏报警

| Name                          | Type           | Description                                                | **Required** | default                                                      |
| ----------------------------- | -------------- | ---------------------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                            | int8           | 按指定方法生成                                             | required     | 主键                                                         |
| construction_passing_alarm_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用            | required     | 联合主键                                                     |
| passing_id                    | text           | 过车记录ID(抓拍系统过车记录表或电子围栏进出记录表的id)     | optional     |                                                              |
| passing_type                  | int4           | 过车报警类型(1.抓拍系统报警 2.电子围栏报警)                | optional     |                                                              |
| vehicle_id                    | text           | 车辆信息表 的vehicle_id                                    | optional     |                                                              |
| license_plate_number          | text           | 车牌号                                                     | optional     |                                                              |
| registration_time             | timestamptz(6) | 登记时间                                                   | optional     |                                                              |
| supervision_time              | timestamptz(6) | 监管时间                                                   | optional     |                                                              |
| remarks                       | text           | 备注                                                       | optional     |                                                              |
| camera_id                     | text           | 摄像头ID                                                   | optional     | **construction_camera 工地摄像头表**的construction_camera_id |
| alarm_type                    | int4           | 报警类型（1.未纳入管控平台 2.未报备工程 3.未纳入管控平台） | optional     |                                                              |
| review_status                 | text           | 审核状态                                                   | optional     |                                                              |
| vehicle_type                  | int4           | 车辆类型                                                   | optional     | **车辆类型**字典                                             |
| is_online                     | boolean        | 是否在线                                                   | optional     |                                                              |
| is_supervision                | boolean        | 是否监管                                                   | optional     |                                                              |
| created_at                    | timestamptz(6) | 创建时间                                                   | required     |                                                              |
| created_by                    | text           | 创建人                                                     | required     | **system_user表**的user_id                                   |
| updated_at                    | timestamptz(6) | 修改时间                                                   | optional     |                                                              |
| updated_by                    | text           | 修改人                                                     | optional     | **system_user表**的user_id                                   |
| deleted_at                    | timestamptz(6) | 删除时间                                                   | optional     |                                                              |
| deleted_by                    | text           | 删除人                                                     | optional     | **system_user表**的user_id                                   |

### electric_fence 电子围栏 

**根据以下数据表整合而成**

工程车:

[INFO_ELECTRIC_CRAWL 电子围栏表](#egn_INFO_ELECTRIC_CRAWL)
[INFO_ELECTRIC_CRAWL_LNGLAT 电子围栏经纬度表](#egn_INFO_ELECTRIC_CRAWL_LNGLAT)

| Name              | Type           | Description                             | Required | Remark   |
| ----------------- | -------------- | --------------------------------------- | -------- | -------- |
| id                | int8           | ID                                      | required | 主键     |
| electric_fence_id | text           | 电子围栏ID                              | required | 联合主键 |
| area_type         | int4           | 区域类型(1矩形 2圆形 3多边形)           | required |          |
| fence_type        | int4           | 围栏类型(1消纳场 2工地 3工程 4重点区域) | optional |          |
| fence_name        | text           | 围栏名称                                | optional |          |
| address           | text           | 地点                                    | optional |          |
| district_id       | text           | 所属区域                                | optional |          |
| gis_data          | polygon        | 围栏数据                                | optional |          |
| created_at        | timestamptz(6) | 创建时间                                | required |          |
| created_by        | text           | 创建人                                  | optional |          |
| updated_at        | timestamptz(6) | 更新时间                                | optional |          |
| updated_by        | text           | 更新人                                  | optional |          |
| deleted_at        | timestamptz(6) | 删除时间                                | optional |          |
| deleted_by        | text           | 删除人                                  | optional |          |
| is_deleted        | bool           | 是否删除                                | required |          |

### electric_fence_enterance_record 电子围栏进出记录表  

**根据以下数据表整合而成**

工程车:

[INFO_VEHPASSCRAWL 电子围栏进出记录表](#egn_INFO_VEHPASSCRAWL)

| Name              | Type           | Description          | Required | Remark   |
| ----------------- | -------------- | -------------------- | -------- | -------- |
| id                | int8           | ID                   | required | 主键     |
| record_id         | text           | 电子围栏进出记录表ID | required | 联合主键 |
| vehicle_id        | text           | 车辆ID               | required |          |
| electric_fence_id | text           | 电子围栏ID           | required |          |
| in_time           | timestamptz(6) | 进入时间             | optional |          |
| out_time          | timestamptz(6) | 离开时间             | optional |          |
| is_out            | bool           | 是否离开             | optional |          |
| position          | point          | 位置点               | optional |          |
| is_online         | bool           | 是否在线             | optional |          |
| created_at        | timestamptz(6) | 创建时间             | required |          |
| created_by        | text           | 创建人               | optional |          |
| updated_at        | timestamptz(6) | 更新时间             | optional |          |
| updated_by        | text           | 更新人               | optional |          |
| deleted_at        | timestamptz(6) | 删除时间             | optional |          |
| deleted_by        | text           | 删除人               | optional |          |
| is_deleted        | bool           | 是否删除             | required |          |



### construction_alarm_push 工地过车报警推送关联表

**根据以下数据表整合而成**

工程车:

[FILEVEHPASSDEAL 工地过车报警推送表](#egn_FILEVEHPASSDEAL)

| Name                       | Type           | Description                                     | **Required** | default                    |
| -------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | int8           | 按指定方法生成                                  | required     | 主键                       |
| construction_alarm_push_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| alarm_id                   | text           | 报警ID                                          | optional     |                            |
| role                       | text           | 角色                                            | optional     |                            |
| is_supervision             | boolean        | 是否监管                                        | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text           | 删除人                                          | optional     | **system_user表**的user_id |

### snapshot_whitelist 抓拍系统白名单表

**根据以下数据表整合而成**

工程车:

[FileVehPassNotPlatform 抓拍系统白名单表](#egn_FileVehPassNotPlatform)

| Name                  | Type           | Description                                     | **Required** | default                    |
| --------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                    | int8           | 按指定方法生成                                  | required     | 主键                       |
| snapshot_whitelist_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number  | text           | 车牌号                                          | optional     |                            |
| vehicle_type          | int4           | 车辆类型                                        | optional     | **车辆类型**字典           |
| user_id               | text           | 操作人ID                                        | optional     |                            |
| vehicle_id            | text           | 车辆ID（非平台车辆时为空）                      | optional     |                            |
| check_type            | int4           | 库的类型（1.需要报警 2.无需报警）               | optional     |                            |
| is_deleted            | boolean        | 是否删除                                        | optional     |                            |
| created_at            | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by            | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by            | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by            | text           | 删除人                                          | optional     | **system_user表**的user_id |

## 工程车报警

### alarm_overspeed  超速报警表

**根据以下数据表整合而成**

工程车:

[ECDSPEEDALARMMSG 超速报警表](#egn_ECDSPEEDALARMMSG)

| Name               | Type           | Description                                             | **Required** | default                            |
| ------------------ | -------------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id                 | int8           | 按指定方法生成                                          | required     | 主键                               |
| alarm_overspeed_id | text           | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id         | text           | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| line_id            | text           | 路线ID                                                  | optional     |                                    |
| alarm_speed        | numeric        | 报警车速                                                | optional     |                                    |
| alarm_times        | int4           | 报警次数                                                | optional     |                                    |
| drive_time         | timestamptz(6) | 行驶时间                                                | optional     |                                    |
| alarm_start_time   | timestamptz(6) | 报警开始时间                                            | optional     |                                    |
| alarm_end_time     | timestamptz(6) | 报警结束时间                                            | optional     |                                    |
| is_cancled         | boolean        | 是否取消报警                                            | optional     |                                    |
| is_supervised      | boolean        | 是否监管                                                | optional     |                                    |
| supervisor         | text           | 监管人                                                  | optional     |                                    |
| supervise_time     | timestamptz(6) | 监管时间                                                | optional     |                                    |
| supervise_remark   | text           | 监管备注                                                | optional     |                                    |
| coordinate         | point          | 经纬度                                                  | optional     |                                    |
| is_resolve         | boolean        | 是否解析                                                | optional     |                                    |
| is_send_msg        | boolean        | 是否发送短信                                            | optional     |                                    |
| alarm_type         | int4           | 报警类型=0                                              | optional     |                                    |
| area               | text           | 地区                                                    | optional     |                                    |
| created_at         | timestamptz(6) | 创建时间                                                | required     |                                    |
| created_by         | text           | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at         | timestamptz(6) | 修改时间                                                | optional     |                                    |
| updated_by         | text           | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at         | timestamptz(6) | 删除时间                                                | optional     |                                    |
| deleted_by         | text           | 删除人                                                  | optional     | **system_user表**的user_id         |

### alarm_fatigue 疲劳驾驶报警表

**根据以下数据表整合而成**

工程车:

[ECDALARMMSG 疲劳驾驶报警表](#egn_ECDALARMMSG)

| Name             | Type           | Description                                             | **Required** | default                            |
| ---------------- | -------------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id               | int8           | 按指定方法生成                                          | required     | 主键                               |
| alarm_fatigue_id | text           | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id       | text           | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| line_id          | text           | 路线ID                                                  | optional     |                                    |
| alarm_speed      | numeric        | 报警车速                                                | optional     |                                    |
| tachograph_speed | numeric        | 行驶记录仪速度                                          | optional     |                                    |
| drive_time       | timestamptz(6) | 行驶时间                                                | optional     |                                    |
| alarm_start_time | timestamptz(6) | 报警开始时间                                            | optional     |                                    |
| alarm_end_time   | timestamptz(6) | 报警结束时间                                            | optional     |                                    |
| is_cancled       | boolean        | 是否取消报警                                            | optional     |                                    |
| is_supervised    | boolean        | 是否监管                                                | optional     |                                    |
| supervisor       | text           | 监管人                                                  | optional     |                                    |
| supervise_time   | timestamptz(6) | 监管时间                                                | optional     |                                    |
| supervise_remark | text           | 监管备注                                                | optional     |                                    |
| coordinate       | point          | 经纬度                                                  | optional     |                                    |
| is_resolve       | boolean        | 是否解析                                                | optional     |                                    |
| is_send_msg      | boolean        | 是否发送短信                                            | optional     |                                    |
| alarm_type       | int4           | 报警类型=4                                              | optional     |                                    |
| area             | text           | 地区                                                    | optional     |                                    |
| created_at       | timestamptz(6) | 创建时间                                                | required     |                                    |
| created_by       | text           | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at       | timestamptz(6) | 修改时间                                                | optional     |                                    |
| updated_by       | text           | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at       | timestamptz(6) | 删除时间                                                | optional     |                                    |
| deleted_by       | text           | 删除人                                                  | optional     | **system_user表**的user_id         |

### alarm_power_outage 超三天断电报警表

**根据以下数据表整合而成**

工程车:

[ONLINEECDALARMMSG 超三天断电报警表](#egn_ONLINEECDALARMMSG)

| Name                | Type           | Description                                             | **Required** | default                            |
| ------------------- | -------------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id                  | int8           | 按指定方法生成                                          | required     | 主键                               |
| alarm_fatigue_id    | text           | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id          | text           | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| line_id             | text           | 路线ID                                                  | optional     |                                    |
| alarm_speed         | numeric        | 报警车速                                                | optional     |                                    |
| tachograph_speed    | numeric        | 行驶记录仪速度                                          | optional     |                                    |
| alarm_time          | timestamptz(6) | 报警日期                                                | optional     |                                    |
| alarm_start_time    | timestamptz(6) | 离线开始时间                                            | optional     |                                    |
| alarm_end_time      | timestamptz(6) | 离线结束时间                                            | optional     |                                    |
| is_power_outage_end | boolean        | 离线是否结束                                            |              |                                    |
| is_cancled          | boolean        | 是否取消报警                                            | optional     |                                    |
| is_supervised       | boolean        | 是否监管                                                | optional     |                                    |
| supervisor          | text           | 监管人                                                  | optional     |                                    |
| supervise_time      | timestamptz(6) | 监管时间                                                | optional     |                                    |
| supervise_remark    | text           | 监管备注                                                | optional     |                                    |
| coordinate          | point          | 经纬度                                                  | optional     |                                    |
| is_resolve          | boolean        | 是否解析                                                | optional     |                                    |
| is_send_msg         | boolean        | 是否发送短信                                            | optional     |                                    |
| is_register_outage  | boolean        | 是否登记停运                                            | optional     |                                    |
| area                | text           | 地区                                                    | optional     |                                    |
| created_at          | timestamptz(6) | 创建时间                                                | required     |                                    |
| created_by          | text           | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at          | timestamptz(6) | 修改时间                                                | optional     |                                    |
| updated_by          | text           | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at          | timestamptz(6) | 删除时间                                                | optional     |                                    |
| deleted_by          | text           | 删除人                                                  | optional     | **system_user表**的user_id         |

## 新型渣土车管理

### vehicle_state_latest 新型车辆最新状态表

**根据以下数据表整合而成**

工程车:

[INFO_SMARTCAR_OPERATION_STATE 新型车辆最新状态表](#egn_INFO_SMARTCAR_OPERATION_STATE)

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
| created_at              | timestamptz(6) | 创建时间                                        | required     |                                      |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id           |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                                      |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id           |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                                      |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id           |

### driver_fingerprint 驾驶员指纹表

**根据以下数据表整合而成**

工程车:

[FINGERPRINTS 指纹信息表 ](#egn_FINGERPRINTS)

[ZTC_FINGERPRINT_INFORMATION 指纹表](#egn_ZTC_FINGERPRINT_INFORMATION)

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
| operation_time        | timestamptz(6) | 操作时间                                        | optional     |                            |
| upload_time           | timestamptz(6) | 上传时间                                        | optional     |                            |
| timestamp             | text        | 时间戳                                          | optional     |                            |
| ternimal_id           | text        | 终端ID                                          | optional     |                            |
| is_success            | boolean     | 是否成功                                        | optional     |                            |
| is_deleted            | boolean     | 是否删除                                        | optional     |                            |
| created_at            | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by            | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by            | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by            | text        | 删除人                                          | optional     | **system_user表**的user_id |

### driver_fingerprint_association  驾驶员指纹关联

**根据以下数据表整合而成**

工程车:

[ZTC_FINGERPRINT_DRIVER 驾驶员指纹关联](#egn_ZTC_FINGERPRINT_DRIVER)

| Name                              | Type        | Description                                                | **Required** | default                    |
| --------------------------------- | ----------- | ---------------------------------------------------------- | ------------ | -------------------------- |
| id                                | bigint      | 按指定方法生成                                             | required     | 主键                       |
| driver_fingerprint_association_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用            | required     | 联合主键                   |
| driver_id                         | text        | **driver_info驾驶员信息表**的driver_id                     | required     |                            |
| fingerprint_name                  | text        | 指纹名称                                                   | optional     |                            |
| driver_fingerprint_id             | text        | **driver_fingerprint 驾驶员指纹表**的driver_fingerprint_id | optional     |                            |
| picture_address                   | text        | 图片地址                                                   | optional     |                            |
| is_deleted                        | boolean     | 是否删除                                                   | optional     |                            |
| created_at                        | timestamptz(6) | 创建时间                                                   | required     |                            |
| created_by                        | text        | 创建人                                                     | required     | **system_user表**的user_id |
| updated_at                        | timestamptz(6) | 修改时间                                                   | optional     |                            |
| updated_by                        | text        | 修改人                                                     | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz(6) | 删除时间                                                   | optional     |                            |
| deleted_by                        | text        | 删除人                                                     | optional     | **system_user表**的user_id |

###  issue_instruction 下发指令表

**根据以下数据表整合而成**

工程车:

[INFO_ORDERLIST  下发指令表(用户下发指令存在该表，后台程序取出指令下发给终端)](#egn_INFO_ORDERLIST)

| Name                  | Type        | Description                                     | **Required** | default                            |
| --------------------- | ----------- | ----------------------------------------------- | ------------ | ---------------------------------- |
| id                    | bigint      | 按指定方法生成                                  | required     | 主键                               |
| issue_instructions_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| license_plate_number  | text        | 车牌号                                          | optional     |                                    |
| vehicle_id            | text        | 车辆ID                                          | optional     | 引用**vehicle_info**表的vehicle_id |
| sim_num               | text        | SIM卡号                                         | optional     |                                    |
| instruction_type      | text        | 指令类型                                        | optional     |                                    |
| instruction_content   | text        | 指令内容                                        | optional     |                                    |
| result                | int4        | 返回结果（1.成功 3.未响应）                     | optional     |                                    |
| user_id               | text        | 操作用户ID                                      | optional     |                                    |
| group_id              | text        | 用户组                                          | optional     |                                    |
| user_region           | text        | 操作人单位                                      | optional     |                                    |
| created_at            | timestamptz(6) | 创建时间                                        | required     |                                    |
| created_by            | text        | 创建人                                          | required     | **system_user表**的user_id         |
| updated_at            | timestamptz(6) | 修改时间                                        | optional     |                                    |
| updated_by            | text        | 修改人                                          | optional     | **system_user表**的user_id         |
| deleted_at            | timestamptz(6) | 删除时间                                        | optional     |                                    |
| deleted_by            | text        | 删除人                                          | optional     | **system_user表**的user_id         |

### alarm_lift_violation 违规举升报警表

**根据以下数据表整合而成**

工程车:

[INFO_SMARTCAR_ALARM_FFQYJS 非法举升报警表](#egn_INFO_SMARTCAR_ALARM_FFQYJS)

| Name                     | Type           | Description                                             | **Required** | default                            |
| ------------------------ | -------------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id                       | int8           | 按指定方法生成                                          | required     | 主键                               |
| alarm_lift_violation_id  | text           | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id               | text           | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| alarm_start_time         | timestamptz(6) | 报警开始时间                                            | optional     |                                    |
| alarm_end_time           | timestamptz(6) | 报警结束时间                                            | optional     |                                    |
| remark                   | text           | 备注                                                    | optional     |                                    |
| is_supervised            | boolean        | 是否监管                                                | optional     |                                    |
| supervisor               | text           | 监管人                                                  | optional     |                                    |
| supervise_time           | timestamptz(6) | 监管时间                                                | optional     |                                    |
| is_completed             | boolean        | 是否完成                                                | optional     |                                    |
| is_supervised_gov        | boolean        | 管理部门是否监管                                        | optional     |                                    |
| supervise_time_gov       | timestamptz(6) | 管理部门监管时间                                        | optional     |                                    |
| is_construction_handle   | boolean        | 工地是否处理                                            | optional     |                                    |
| construction_handle_time | timestamptz(6) | 工地处理时间                                            | optional     |                                    |
| coordinate               | point          | 经纬度                                                  | optional     |                                    |
| created_at               | timestamptz(6) | 创建时间                                                | required     |                                    |
| created_by               | text           | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at               | timestamptz(6) | 修改时间                                                | optional     |                                    |
| updated_by               | text           | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at               | timestamptz(6) | 删除时间                                                | optional     |                                    |
| deleted_by               | text           | 删除人                                                  | optional     | **system_user表**的user_id         |

### alarm_processing_record 报警处理记录表

**根据以下数据表整合而成**

工程车:

[ALARMDEAL 报警处理记录](#egn_ALARMDEAL)

| Name                         | Type           | Description                                                  | **Required** | default                    |
| ---------------------------- | -------------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                           | int8           | 按指定方法生成                                               | required     | 主键                       |
| alarm_supervision_picture_id | int8           | **alarm_supervision_picture_ upload**报警监管图片上传表的alarm_supervision_picture_id | required     |                            |
| processing_content           | text           | 处理内容                                                     | optional     |                            |
| processing_time              | timestamptz(6) | 处理时间                                                     | optional     |                            |
| processing_type              | int4           | 处理类型  1.超速报警  2.疲劳驾驶  3.工程报警  4.超三天断电报警  5.进出区域报警  6.进出区域报警  7.安检到期报警  11.进出工地报警 | optional     |                            |
| alarm_data_id                | text           | 报警ID                                                       | required     |                            |
| operation_user               | text           | 操作用户                                                     | optional     | **system_user表**的user_id |
| is_deleted                   | boolean        | 是否删除                                                     | optional     | false                      |
| created_at                   | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by                   | text           | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by                   | text           | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by                   | text           | 删除人                                                       | optional     | **system_user表**的user_id |

###  alarm_supervision_picture_upload 报警监管图片上传表

**根据以下数据表整合而成**

工程车:

[ALARMDEALPHOTO 报警监管图片上传](#egn_ALARMDEALPHOTO)

| Name                         | Type        | Description                                                  | **Required** | default                    |
| ---------------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                               | required     | 主键                       |
| alarm_supervision_picture_id | text        | 报警监管图片上传外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| pic_address                  | text        | 图片地址                                                     | optional     |                            |
| created_at                   | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by                   | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by                   | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by                   | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### alarm_start_violation 违规启动报警表

**根据以下数据表整合而成**

工程车:

[INFO_SMARTCAR_ALARM_WGQD 新型车辆违规启动报警](#egn_INFO_SMARTCAR_ALARM_WGQD)

| Name                     | Type           | Description                                             | **Required** | default                            |
| ------------------------ | -------------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id                       | int8           | 按指定方法生成                                          | required     | 主键                               |
| alarm_start_violation_id | text           | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id               | text           | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| alarm_start_time         | timestamptz(6) | 报警开始时间                                            | optional     |                                    |
| alarm_end_time           | timestamptz(6) | 报警结束时间                                            | optional     |                                    |
| remark                   | text           | 备注                                                    | optional     |                                    |
| is_supervised            | boolean        | 是否监管                                                | optional     |                                    |
| supervisor               | text           | 监管人                                                  | optional     |                                    |
| supervise_time           | timestamptz(6) | 监管时间                                                | optional     |                                    |
| is_completed             | boolean        | 是否完成                                                | optional     |                                    |
| is_supervised_gov        | boolean        | 管理部门是否监管                                        | optional     |                                    |
| supervise_time_gov       | timestamptz(6) | 管理部门监管时间                                        | optional     |                                    |
| is_construction_handle   | boolean        | 工地是否处理                                            | optional     |                                    |
| construction_handle_time | timestamptz(6) | 工地处理时间                                            | optional     |                                    |
| coordinate               | point          | 经纬度                                                  | optional     |                                    |
| created_at               | timestamptz(6) | 创建时间                                                | required     |                                    |
| created_by               | text           | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at               | timestamptz(6) | 修改时间                                                | optional     |                                    |
| updated_by               | text           | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at               | timestamptz(6) | 删除时间                                                | optional     |                                    |
| deleted_by               | text           | 删除人                                                  | optional     | **system_user表**的user_id         |

### alarm_spoil_violation 非法弃土报警表

**根据以下数据表整合而成**

工程车:

[INFO_SMARTCAR_ALARM_FFQT 非法弃土报警表](#egn_INFO_SMARTCAR_ALARM_FFQT)

| Name                     | Type           | Description                                             | **Required** | default                            |
| ------------------------ | -------------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id                       | int8           | 按指定方法生成                                          | required     | 主键                               |
| alarm_spoil_violation_id | text           | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id               | text           | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| alarm_start_time         | timestamptz(6) | 报警开始时间                                            | optional     |                                    |
| alarm_end_time           | timestamptz(6) | 报警结束时间                                            | optional     |                                    |
| remark                   | text           | 备注                                                    | optional     |                                    |
| is_supervised            | boolean        | 是否监管                                                | optional     |                                    |
| supervisor               | text           | 监管人                                                  | optional     |                                    |
| supervise_time           | timestamptz(6) | 监管时间                                                | optional     |                                    |
| is_completed             | boolean        | 是否完成                                                | optional     |                                    |
| is_supervised_gov        | boolean        | 管理部门是否监管                                        | optional     |                                    |
| supervise_time_gov       | timestamptz(6) | 管理部门监管时间                                        | optional     |                                    |
| is_construction_handle   | boolean        | 工地是否处理                                            | optional     |                                    |
| construction_handle_time | timestamptz(6) | 工地处理时间                                            | optional     |                                    |
| coordinate               | point          | 经纬度                                                  | optional     |                                    |
| alarm_type               | text           | 报警类型                                                | optional     |                                    |
| created_at               | timestamptz(6) | 创建时间                                                | required     |                                    |
| created_by               | text           | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at               | timestamptz(6) | 修改时间                                                | optional     |                                    |
| updated_by               | text           | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at               | timestamptz(6) | 删除时间                                                | optional     |                                    |
| deleted_by               | text           | 删除人                                                  | optional     | **system_user表**的user_id         |

### vehicle_status_change 车辆状态变更表(新表)

**根据以下数据表整合而成**

工程车:

[INFO_SMARTCAR_JSZT 举升状态变更记录](#egn_INFO_SMARTCAR_JSZT)

[INFO_SMARTCAR_KZZT 车辆载重状态变更记录](#egn_INFO_SMARTCAR_KZZT)

[INFO_SMARTCAR_CXZT 车辆车厢状态变更记录表](#egn_INFO_SMARTCAR_CXZT)

| Name                     | Type        | Description                                     | **Required** | default                    |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                       |
| vehicle_status_change_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id               | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| terminal_id              | text        | 终端ID                                          | optional     |                            |
| start_time               | timestamptz(6) | 开始时间                                        | optional     |                            |
| end_time                 | timestamptz(6) | 结束时间                                        | optional     |                            |
| vehicle_status_type      | integer     | 车辆状态类型(车厢状态,举升状态,载重状态)        | optional     | **车辆状态类型**字典       |
| value                    | text        | 值                                              | optional     |                            |
| is_completed             | boolean     | 是否完成                                        | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_unsealed 未密闭行驶报警表

**根据以下数据表整合而成**

工程车:

[INFO_SMARTCAR_ALARM_DGWMB 顶盖未密闭行驶报警表](#egn_INFO_SMARTCAR_ALARM_DGWMB)

| Name                     | Type           | Description                                             | **Required** | default                            |
| ------------------------ | -------------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id                       | int8           | 按指定方法生成                                          | required     | 主键                               |
| alarm_unsealed_id        | text           | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id               | text           | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| alarm_start_time         | timestamptz(6) | 报警开始时间                                            | optional     |                                    |
| alarm_end_time           | timestamptz(6) | 报警结束时间                                            | optional     |                                    |
| remark                   | text           | 备注                                                    | optional     |                                    |
| is_supervised            | boolean        | 是否监管                                                | optional     |                                    |
| supervisor               | text           | 监管人                                                  | optional     |                                    |
| supervise_time           | timestamptz(6) | 监管时间                                                | optional     |                                    |
| is_completed             | boolean        | 是否完成                                                | optional     |                                    |
| is_supervised_gov        | boolean        | 管理部门是否监管                                        | optional     |                                    |
| supervise_time_gov       | timestamptz(6) | 管理部门监管时间                                        | optional     |                                    |
| is_construction_handle   | boolean        | 工地是否处理                                            | optional     |                                    |
| construction_handle_time | timestamptz(6) | 工地处理时间                                            | optional     |                                    |
| coordinate               | point          | 经纬度                                                  | optional     |                                    |
| created_at               | timestamptz(6) | 创建时间                                                | required     |                                    |
| created_by               | text           | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at               | timestamptz(6) | 修改时间                                                | optional     |                                    |
| updated_by               | text           | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at               | timestamptz(6) | 删除时间                                                | optional     |                                    |
| deleted_by               | text           | 删除人                                                  | optional     | **system_user表**的user_id         |

### alarm_load_violation 非法载重报警表

**根据以下数据表整合而成**

工程车:

[INFO_SMARTCAR_ALARM_FFQYZZ 非法区域载重报警表](#egn_INFO_SMARTCAR_ALARM_FFQYZZ)

| Name                     | Type           | Description                                             | **Required** | default                            |
| ------------------------ | -------------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id                       | int8           | 按指定方法生成                                          | required     | 主键                               |
| alarm_unsealed_id        | text           | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id               | text           | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| alarm_start_time         | timestamptz(6) | 报警开始时间                                            | optional     |                                    |
| alarm_end_time           | timestamptz(6) | 报警结束时间                                            | optional     |                                    |
| remark                   | text           | 备注                                                    | optional     |                                    |
| is_supervised            | boolean        | 是否监管                                                | optional     |                                    |
| supervisor               | text           | 监管人                                                  | optional     |                                    |
| supervise_time           | timestamptz(6) | 监管时间                                                | optional     |                                    |
| is_completed             | boolean        | 是否完成                                                | optional     |                                    |
| is_supervised_gov        | boolean        | 管理部门是否监管                                        | optional     |                                    |
| supervise_time_gov       | timestamptz(6) | 管理部门监管时间                                        | optional     |                                    |
| is_construction_handle   | boolean        | 工地是否处理                                            | optional     |                                    |
| construction_handle_time | timestamptz(6) | 工地处理时间                                            | optional     |                                    |
| coordinate               | point          | 经纬度                                                  | optional     |                                    |
| created_at               | timestamptz(6) | 创建时间                                                | required     |                                    |
| created_by               | text           | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at               | timestamptz(6) | 修改时间                                                | optional     |                                    |
| updated_by               | text           | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at               | timestamptz(6) | 删除时间                                                | optional     |                                    |
| deleted_by               | text           | 删除人                                                  | optional     | **system_user表**的user_id         |

## 工程车违规记分(黑名单)

### violation_file 违规记分文件表

**根据以下数据表整合而成**

工程车:

[KFFILE 违规记分文件表](#egn_KFFILE)

| Name              | Type           | Description                                     | **Required** | default                    |
| ----------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                | bigint         | 按指定方法生成                                  | required     | 主键                       |
| violation_file_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| file_path         | text           | 文件地址                                        | optional     |                            |
| created_at        | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by        | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at        | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by        | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at        | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by        | text           | 删除人                                          | optional     | **system_user表**的user_id |


### ent_violation_scoring_season 企业违规记分季度排名

**根据以下数据表整合而成**

工程车:

[COMPANYSCORERANK 企业违规记分季度排名](#egn_COMPANYSCORERANK)

| Name                            | Type           | Description                                     | **Required** | default                    |
| ------------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                              | bigint         | 按指定方法生成                                  | required     | 主键                       |
| ent_violation_scoring_season_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| ent_id                          | text           | 企业id                                          | optional     |                            |
| rank                            | int4           | 排名                                            | optional     |                            |
| violation_points                | int4           | 违规记分                                        | optional     |                            |
| year                            | int4           | 年份                                            | optional     |                            |
| season                          | int4           | 季度                                            | optional     |                            |
| created_at                      | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                      | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                      | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                      | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                      | text           | 删除人                                          | optional     | **system_user表**的user_id |

### ent_violation_scoring_items 企业扣分事项表

**根据以下数据表整合而成**

工程车:

COMPANYSCORESET 企业扣分项

| Name                           | Type           | Description                                     | **Required** | default                    |
| ------------------------------ | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                             | bigint         | 按指定方法生成                                  | required     | 主键                       |
| ent_violation_scoring_items_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| demerit_content                | text           | 扣分内容                                        | optional     |                            |
| demerit_type                   | text           | 扣分类型                                        | optional     |                            |
| demerit_points                 | numeric        | 扣分分值                                        | optional     |                            |
| is_deleted                     | boolean        | 是否删除                                        | optional     |                            |
| created_at                     | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                     | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                     | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                     | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                     | text           | 删除人                                          | optional     | **system_user表**的user_id |

### ent_violation_scoring 企业违规记分

**根据以下数据表整合而成**

工程车:

COMPANYSCORELOG 企业扣分日志表

| Name                     | Type           | Description                                     | **Required** | default                    |
| ------------------------ | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | int8           | 按指定方法生成                                  | required     | 主键                       |
| ent_violation_scoring_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| ent_id                   | text           | 企业id                                          | optional     |                            |
| current_points           | numeric        | 当前分值                                        | optional     |                            |
| demerit_points           | numeric        | 扣分分值                                        | optional     |                            |
| demerit_content          | text           | 扣分内容                                        | optional     |                            |
| file_id                  | text           | 文件ID                                          | optional     |                            |
| remarks                  | text           | 备注                                            | optional     |                            |
| is_deleted               | boolean        | 是否删除                                        | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text           | 删除人                                          | optional     | **system_user表**的user_id |

### ent_blacklist_early_warning 企业黑名单预警表  

**根据以下数据表整合而成**

工程车:

CompanyAlarm 企业黑名单预警表

| Name          | Type           | Description                          | Required | Remark   |
| ------------- | -------------- | ------------------------------------ | -------- | -------- |
| id            | int8           | ID                                   | required | 主键     |
| alarm_id      | text           | 企业黑名单预警表ID                   | required | 联合主键 |
| enterprise_id | text           | 运输企业ID                           | required |          |
| type          | int4           | 类型                                 | required |          |
| year          | int4           | 年份                                 | required |          |
| is_black      | bool           | 企业的车辆是否10%以上的车进入黑名单  | required |          |
| is_rank       | bool           | 企业是否连续两个季度违规记分排名前十 | required |          |
| created_at    | timestamptz(6) | 创建时间                             | required |          |
| created_by    | text           | 创建人                               | optional |          |
| updated_at    | timestamptz(6) | 更新时间                             | optional |          |
| updated_by    | text           | 更新人                               | optional |          |
| deleted_at    | timestamptz(6) | 删除时间                             | optional |          |
| deleted_by    | text           | 删除人                               | optional |          |
| is_deleted    | bool           | 是否删除                             | optional |          |

### veh_blacklist_early_warning 车辆黑名单预警表

**根据以下数据表整合而成**

工程车:

VehicleAlarm 车辆黑名单预警表

| Name                 | Type           | Description      | Required | Remark   |
| -------------------- | -------------- | ---------------- | -------- | -------- |
| id                   | int8           | ID               | required | 主键     |
| alarm_id             | text           | 车辆黑名单警告ID | required | 联合主键 |
| vehicle_id           | text           | 车辆ID           | required |          |
| license_plate_number | text           | 车牌号           | optional |          |
| year                 | int4           | 年份             | optional |          |
| score_num            | int4           | 扣分数量         | optional |          |
| score_count          | int4           | 扣分次数         | optional |          |
| created_at           | timestamptz(6) | 创建时间         | required |          |
| created_by           | text           | 创建人           | optional |          |
| updated_at           | timestamptz(6) | 更新时间         | optional |          |
| updated_by           | text           | 更新人           | optional |          |
| deleted_at           | timestamptz(6) | 删除时间         | optional |          |
| deleted_by           | text           | 删除人           | optional |          |
| is_deleted           | bool           | 是否删除         | required |          |

### veh_violation_scoring 车辆违规记分表

**根据以下数据表整合而成**

工程车:

[VEHICLESCORELOG 车辆违规记分记录](#egn_VEHICLESCORELOG)

| Name                     | Type           | Description                                     | **Required** | default                            |
| ------------------------ | -------------- | ----------------------------------------------- | ------------ | ---------------------------------- |
| id                       | int8           | 按指定方法生成                                  | required     | 主键                               |
| veh_violation_scoring_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id               | text           | 车辆ID                                          | optional     | 引用**vehicle_info**表的vehicle_id |
| demerit_points           | numeric        | 扣分分值                                        | optional     |                                    |
| remarks                  | text           | 备注                                            | optional     |                                    |
| demerit_content          | text           | 扣分内容                                        | optional     |                                    |
| violation_date           | timestamptz(6) | 违法日期                                        | optional     |                                    |
| operation_person         | text           | 操作人                                          | optional     |                                    |
| operation_date           | timestamptz(6) | 操作时间                                        | optional     |                                    |
| file_id                  | text           | 文件ID                                          | optional     |                                    |
| is_deleted               | boolean        | 是否删除                                        | optional     |                                    |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                                    |
| created_by               | text           | 创建人                                          | required     | **system_user表**的user_id         |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                                    |
| updated_by               | text           | 修改人                                          | optional     | **system_user表**的user_id         |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                                    |
| deleted_by               | text           | 删除人                                          | optional     | **system_user表**的user_id         |

### veh_violation_scoring_items 车辆扣分事项表

**根据以下数据表整合而成**

工程车:

[VEHICLESCORESET 车辆扣分项](#egn_VEHICLESCORESET)

| Name                           | Type           | Description                                     | **Required** | default                    |
| ------------------------------ | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                             | bigint         | 按指定方法生成                                  | required     | 主键                       |
| ent_violation_scoring_items_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| demerit_content                | text           | 扣分内容                                        | optional     |                            |
| demerit_type                   | text           | 扣分类型                                        | optional     |                            |
| demerit_points                 | numeric        | 扣分分值                                        | optional     |                            |
| operation_person               | text           | 操作人                                          | optional     |                            |
| operation_date                 | timestamptz(6) | 操作时间                                        | optional     |                            |
| is_deleted                     | boolean        | 是否删除                                        | optional     |                            |
| created_at                     | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                     | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                     | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                     | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                     | text           | 删除人                                          | optional     | **system_user表**的user_id |

### veh_accident 车辆交通事故表

**根据以下数据表整合而成**

工程车:

[VehicleAccidenLog 车辆交通事故表](#egn_VehicleAccidenLog)

| Name                 | Type           | Description                                     | **Required** | default                            |
| -------------------- | -------------- | ----------------------------------------------- | ------------ | ---------------------------------- |
| id                   | bigint         | 按指定方法生成                                  | required     | 主键                               |
| veh_accident_id      | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id           | text           | 车辆ID                                          | optional     | 引用**vehicle_info**表的vehicle_id |
| score                | numeric        | 记分分值                                        | optional     |                                    |
| remarks              | text           | 备注                                            | optional     |                                    |
| demerit_id           | text           | 扣分项ID                                        | optional     |                                    |
| file_id              | text           | 文件ID                                          | optional     |                                    |
| violation_time       | timestamptz(6) | 违法时间                                        | optional     |                                    |
| driver_name          | text           | 驾驶员姓名                                      | optional     |                                    |
| driving_liscense_num | text           | 驾驶证号                                        | optional     |                                    |
| driver_id            | text           | 驾驶员ID                                        | optional     |                                    |
| accident_hurt        | int4           | 事故致伤人数                                    | optional     |                                    |
| accident_dead        | int4            | 事故致死人数                                    | optional     |                                    |
| violation_plate_num  | text           | 违法车牌                                        | optional     |                                    |
| ent_id               | text           | 所属企业ID                                      | optional     |                                    |
| ent_name             | text           | 所属企业名称                                    | optional     |                                    |
| veh_location_code    | text           | 车辆属地编码                                    | optional     |                                    |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                                    |
| created_by           | text           | 创建人                                          | required     | **system_user表**的user_id         |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                                    |
| updated_by           | text           | 修改人                                          | optional     | **system_user表**的user_id         |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                                    |
| deleted_by           | text           | 删除人                                          | optional     | **system_user表**的user_id         |



### veh_accident_file 车辆交通事故附件表

**根据以下数据表整合而成**

工程车:

[KFFile 车辆交通事故表](#egn_KFFile)

| Name                 | Type           | Description                                     | **Required** | default                    |
| -------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint         | 按指定方法生成                                  | required     | 主键                       |
| veh_accident_file_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| file_path            | text           | 文件地址                                        | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text           | 删除人                                          | optional     | **system_user表**的user_id |

### driver_blacklist_apply  驾驶员黑名单申请

**根据以下数据表整合而成**

工程车:

DriverBlackListMain 驾驶员黑名单申请主表

| Name           | Type           | Description                                                  | Required | Remark   |
| -------------- | -------------- | ------------------------------------------------------------ | -------- | -------- |
| id             | int8           | ID                                                           | required | 主键     |
| apply_id       | text           | 驾驶员黑名单申请ID                                           | required | 联合主键 |
| driver_id      | text           | 驾驶员ID                                                     | required |          |
| apply_reason   | text           | 申请原因                                                     | optional |          |
| apply_type     | int4           | 申请类型（1：加入黑名单，2：撤销黑名单）                     | optional |          |
| blacklist_type | int4           | 黑名单类型（1.半挂牵引车违规运载砂石 2.逃避交通违法处罚 3.拒不参加安全教育培训 4.其他） | optional |          |
| state          | int4           | 状态（1：待交警大队审批2：待交警支队审批，3：退回，4：完成)  | optional |          |
| created_at     | timestamptz(6) | 创建时间                                                     | required |          |
| created_by     | text           | 创建人                                                       | required |          |
| updated_at     | timestamptz(6) | 更新时间                                                     | optional |          |
| updated_by     | text           | 更新人                                                       | optional |          |
| deleted_at     | timestamptz(6) | 删除时间                                                     | optional |          |
| deleted_by     | text           | 删除人                                                       | optional |          |
| is_deleted     | bool           | 是否删除                                                     | required |          |

### driver_blacklist 驾驶员黑名单

**根据以下数据表整合而成**

工程车:

DRIVER_BLACK_LIST 驾驶员黑名单

| Name                | Type           | Description    | Required | Remark   |
| ------------------- | -------------- | -------------- | -------- | -------- |
| id                  | int8           | ID             | required | 主键     |
| driver_blacklist_id | text           | 驾驶员黑名单ID | required | 联合主键 |
| driver_id           | text           | 驾驶员ID       | optional |          |
| driver_num          | text           | 驾驶证号码     | optional |          |
| apply_reason        | text           | 申请原因       | optional |          |
| blacklist_type      | int4           | 黑名单类型     | optional |          |
| driver_name         | text           | 驾驶员姓名     | optional |          |
| driver_tel          | text           | 驾驶员联系电话 | optional |          |
| vehicle_license_num | text           | 车牌号         | optional |          |
| driver_sex          | int4           | 性别           | optional |          |
| created_at          | timestamptz(6) | 创建时间       | required |          |
| created_by          | text           | 创建人         | optional |          |
| updated_at          | timestamptz(6) | 更新时间       | optional |          |
| updated_by          | text           | 更新人         | optional |          |
| deleted_at          | timestamptz(6) | 删除时间       | optional |          |
| deleted_by          | text           | 删除人         | optional |          |
| is_deleted          | bool           | 是否删除       | required |          |

### violation_registration 违法信息登记表

**根据以下数据表整合而成**

工程车:

Peccancy 违法信息登记表

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
| illegal_time              | timestamptz(6) | 违法时间                                        | optional     |                                |
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
| created_at                | timestamptz(6) | 创建时间                                        | required     |                                |
| created_by                | text        | 创建人                                          | required     | **system_user表**的user_id     |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                                |
| updated_by                | text        | 修改人                                          | optional     | **system_user表**的user_id     |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                                |
| deleted_by                | text        | 删除人                                          | optional     | **system_user表**的user_id     |

### serious_traffic_violation 严重交通违法行为表

**根据以下数据表整合而成**

工程车:

AllYZWFData 严重交通违法行为表

| Name                         | Type        | Description                                     | **Required** | default                         |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                            |
| serious_traffic_violation_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                        |
| vehicle_id                   | text        | 违章车辆id                                      | optional     | **vehicle_info表**的vehicle_id  |
| illegal_code                 | text        | 违法代码                                        | optional     | **VIO_CODEWFDM** 违法描述字典表 |
| illegal_time                 | timestamptz(6) | 违法日期                                        | optional     |                                 |
| created_at                   | timestamptz(6) | 创建时间                                        | required     |                                 |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id      |
| updated_at                   | timestamptz(6) | 修改时间                                        | optional     |                                 |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id      |
| deleted_at                   | timestamptz(6) | 删除时间                                        | optional     |                                 |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id      |

## 黑名单(重点车)

### blacklist_oper 黑名单操作记录表

**根据以下数据表整合而成**

重点车:

[INFO_BLACKLISTLOG 黑名单操作记录表](#major_INFO_BLACKLISTLOG)

| Name              | Type           | Description                                     | **Required** | default                    |
| ----------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                | bigint         | 按指定方法生成                                  | required     | 主键                       |
| blacklist_oper_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| blacklist_type    | int4           | 黑名单类别                                      | optional     |                            |
| oper_type         | int4           | 操作类别                                        | optional     |                            |
| reason            | text           | 原由                                            | optional     |                            |
| oper_time         | timestamptz(6) | 操作时间                                        | optional     |                            |
| oper_person       | text           | 操作人                                          | optional     |                            |
| created_at        | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by        | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at        | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by        | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at        | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by        | text           | 删除人                                          | optional     | **system_user表**的user_id |

## 营运车报警

### oper_alarm_data 营运车报警数据表

**根据以下数据表整合而成**

营运车:

[REPORT_ALARM_DATA 报警数据](#oper_REPORT_ALARM_DATA)

| Name                   | Type           | Description                                          | **Required** | default                            |
| ---------------------- | -------------- | ---------------------------------------------------- | ------------ | ---------------------------------- |
| id                     | int8           | 按指定方法生成                                       | required     | 主键                               |
| vehicle_alarm_data_id  | text           | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                           |
| vehicle_id             | text           | 车辆ID                                               | optional     | 引用**vehicle_info**表的vehicle_id |
| alarm_type             | text           | 报警类型(营运车共60种报警)                           | optional     | **报警类型**字典                   |
| alarm_start_time       | timestamptz(6) | 报警开始时间                                         | optional     |                                    |
| alarm_end_time         | timestamptz(6) | 报警结束时间                                         | optional     |                                    |
| alarm_end_position     | text           | 报警结束位置                                         | optional     |                                    |
| latest_alarm_time      | timestamptz(6) | 最新报警时间                                         | optional     |                                    |
| latest_alarm_position  | int4           | 最新报警位置                                         | optional     |                                    |
| is_alarm_effective     | boolean        | 报警是否有效                                         | optional     |                                    |
| is_alarm_over          | boolean        | 报警是否结束                                         | optional     |                                    |
| is_cancel_alarm        | boolean        | 是否取消报警                                         | optional     |                                    |
| alarm_source           | text           | 报警来源:1.终端报警 2.平台报警                       | optional     | **报警来源**字典                   |
| processing_time        | timestamptz(6) | 处理时间                                             | optional     |                                    |
| processing_method      | text           | 处理方式                                             | optional     | **处警处理方式**字典               |
| processing_description | text           | 处理描述                                             | optional     |                                    |
| processor              | text           | 处理人                                               | optional     | **system_user表**的user_id         |
| processing_status      | text           | 处理状态                                             | optional     | **处警处理状态**字典               |
| tachograph_speed       | numeric        | 行驶记录仪速度                                       | optional     |                                    |
| GPS_speed              | numeric        | GPS速度                                              | optional     |                                    |
| maximum_speed          | numeric        | 最高速度                                             | optional     |                                    |
| speed_limit_threshold  | numeric        | 限速阀值                                             | optional     |                                    |
| coordinate             | point          | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                                    |
| location_description   | text           | 位置描述                                             | optional     |                                    |
| duration               | text           | 持续时间                                             | optional     |                                    |
| road_grade             | text           | 道路等级                                             | optional     | **道路等级**字典                   |
| road_name              | text           | 道路名称                                             | optional     |                                    |
| area_id                | text           | 进区域ID                                             | optional     |                                    |
| alarm_deal_id          | text           | 处警ID                                               | optional     |                                    |
| pid                    | text           | 地区                                                 | optional     |                                    |
| created_at             | timestamptz(6) | 创建时间                                             | required     |                                    |
| created_by             | text           | 创建人                                               | required     | **system_user表**的user_id         |
| updated_at             | timestamptz(6) | 修改时间                                             | optional     |                                    |
| updated_by             | text           | 修改人                                               | optional     | **system_user表**的user_id         |
| deleted_at             | timestamptz(6) | 删除时间                                             | optional     |                                    |
| deleted_by             | text           | 删除人                                               | optional     | **system_user表**的user_id         |

### oper_alarm_overspeed 营运车超速报警表

**根据以下数据表整合而成**

营运车:

[REPORT_ALARM_DATA_OS_NEW 超速报警明细](#oper_REPORT_ALARM_DATA_OS_NEW)

[REPORT_ALARM_DATA_OVERSPEED 超速报警表](#oper_REPORT_ALARM_DATA_OVERSPEED)

| Name                    | Type           | Description                                     | **Required** | default                            |
| ----------------------- | -------------- | ----------------------------------------------- | ------------ | ---------------------------------- |
| id                      | int8           | 按指定方法生成                                  | required     | 主键                               |
| oper_alarm_overspeed_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| vehicle_id              | text           | 车辆ID                                          | optional     | 引用**vehicle_info**表的vehicle_id |
| alarm_start_time        | timestamptz(6) | 报警开始时间                                    | optional     |                                    |
| alarm_end_time          | timestamptz(6) | 报警结束时间                                    | optional     |                                    |
| start_coordinate        | point          | 开始经纬度                                      | optional     |                                    |
| end_coordinate          | point          | 结束经纬度                                      | optional     |                                    |
| max_speed               | numeric        | 最高速度                                        | optional     |                                    |
| start_pos               | text           | 开始位置                                        | optional     |                                    |
| end_pos                 | text           | 结束位置                                        | optional     |                                    |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                                    |
| created_by              | text           | 创建人                                          | required     | **system_user表**的user_id         |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                                    |
| updated_by              | text           | 修改人                                          | optional     | **system_user表**的user_id         |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                                    |
| deleted_by              | text           | 删除人                                          | optional     | **system_user表**的user_id         |

### oper_alarm_fatigue 营运车疲劳驾驶报警表

**根据以下数据表整合而成**

营运车:

[REPORT_ALARM_DATA_FATIGUE 疲劳驾驶报警表](#oper_REPORT_ALARM_DATA_FATIGUE) 

| Name                     | Type           | Description                                             | **Required** | default                            |
| ------------------------ | -------------- | ------------------------------------------------------- | ------------ | ---------------------------------- |
| id                       | int8           | 按指定方法生成                                          | required     | 主键                               |
| oper_alarm_fatigue_id    | text           | 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                           |
| oper_alarm_processing_id | text           | 报警处理ID                                              | optional     |                                    |
| alarm_type               | int4           | 报警类型                                                | optional     |                                    |
| alarm_start_time         | timestamptz(6) | 报警开始时间                                            | optional     |                                    |
| alarm_end_time           | timestamptz(6) | 报警结束时间                                            | optional     |                                    |
| vehicle_id               | text           | 车辆ID                                                  | optional     | 引用**vehicle_info**表的vehicle_id |
| coordinate               | point          | 经纬度                                                  | optional     |                                    |
| tachograph_speed         | numeric        | 行驶记录仪速度                                          | optional     |                                    |
| gps_speed                | numeric        | GPS速度                                                 | optional     |                                    |
| lastest_alarm_time       | timestamptz(6) | 最新报警时间                                            | optional     |                                    |
| alarm_last_time          | timestamptz(6) | 报警持续时间                                            | optional     |                                    |
| alarm_source             | int4           | 报警来源 1.终端报警 2.平台报警                          | optional     |                                    |
| is_valid                 | boolean        | 是否有效                                                | optional     |                                    |
| is_make_up               | boolean        | 是否补传                                                | optional     |                                    |
| is_alarm_end             | boolean        | 报警是否结束                                            | optional     |                                    |
| created_at               | timestamptz(6) | 创建时间                                                | required     |                                    |
| created_by               | text           | 创建人                                                  | required     | **system_user表**的user_id         |
| updated_at               | timestamptz(6) | 修改时间                                                | optional     |                                    |
| updated_by               | text           | 修改人                                                  | optional     | **system_user表**的user_id         |
| deleted_at               | timestamptz(6) | 删除时间                                                | optional     |                                    |
| deleted_by               | text           | 删除人                                                  | optional     | **system_user表**的user_id         |

### oper_alarm_processing 营运车报警处理表

**根据以下数据表整合而成**

营运车:

[MON_ALARM_DEAL_ZF 报警处置表](#oper_MON_ALARM_DEAL_ZF) 

| Name           | Type           | Description    | **Required** | default                    |
| -------------- | -------------- | -------------- | ------------ | -------------------------- |
| id             | int8           | 按指定方法生成 | required     | 主键                       |
| alarm_id       | text           | 报警ID         | optional     |                            |
| handler        | text           | 处理人         | optional     |                            |
| handle_time    | timestamptz(6) | 处理时间       | optional     |                            |
| handle_content | text           | 处理内容       | optional     |                            |
| created_at     | timestamptz(6) | 创建时间       | required     |                            |
| created_by     | text           | 创建人         | required     | **system_user表**的user_id |
| updated_at     | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by     | text           | 修改人         | optional     | **system_user表**的user_id |
| deleted_at     | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by     | text           | 删除人         | optional     | **system_user表**的user_id |

### alarm_vioce_record 语音报警记录

**根据以下数据表整合而成**

营运车:

[SPEECH_REMIND_LOG 语音报警记录](#oper_SPEECH_REMIND_LOG)

| Name                  | Type        | Description                                     | **Required** | default                    |
| --------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                    | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_vioce_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id            | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| alarm_time            | timestamptz(6) | 报警时间                                        | optional     |                            |
| alarm_type            | text        | 报警类型                                        | optional     |                            |
| remind_time           | timestamptz(6) | 提醒时间                                        | optional     |                            |
| remind_content        | text        | 提醒内容                                        | optional     |                            |
| input_person          | text        | 录入人                                          | optional     | **system_user表**的user_id |
| input_time            | timestamptz(6) | 录入时间                                        | optional     |                            |
| is_success            | boolean     | 是否成功                                        | optional     |                            |
| created_at            | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by            | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by            | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by            | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_veh_times_record 同一车辆报警次数记录表(新表)

**根据以下数据表整合而成**

营运车:

[REPORT_ALARM_DATA_COUNT 4次以上同一报警同一车辆统计表](#oper_REPORT_ALARM_DATA_COUNT)

[REPORT_ALARM_DATA_FOUR 同部车同类型报警4次以上表](#oper_REPORT_ALARM_DATA_FOUR) 

| Name                      | Type           | Description                                     | **Required** | default                    |
| ------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | bigint         | 按指定方法生成                                  | required     | 主键                       |
| alarm_veh_times_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                | text           | 车辆信息表 的vehicle_id                         | required     |                            |
| alarm_type                | text           | 报警类型                                        | optional     | **报警类型**字典           |
| disposal_measure          | text           | 处置措施                                        | optional     |                            |
| disposal_time             | timestamptz(6) | 处置时间                                        | optional     |                            |
| disposal_result           | text           | 处置结果                                        | optional     |                            |
| is_disposal               | boolean        | 是否处置                                        | optional     |                            |
| disposal_method           | int4           | 处置方式                                        | optional     | **处置方式**字典           |
| duty_person               | text           | 值班人                                          | optional     |                            |
| alarm_times               | text           | 报警次数                                        | optional     |                            |
| remarks                   | text           | 备注                                            | optional     |                            |
| record_time               | timestamptz(6) | 记录时间                                        | optional     |                            |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                | text           | 删除人                                          | optional     | **system_user表**的user_id |



## 动态监管

### dynamic_supervision_detail 动态监管抽查明细表

**根据以下数据表整合而成**

重点车:

[INFO_VEHICLE_GPS 动态监管抽查明细表](#major_INFO_VEHICLE_GPS) 

营运车:

[INFO_DYNAMIC_CHECK_RG_DETAIL 企业动态监管抽查明细](#oper_INFO_DYNAMIC_CHECK_RG_DETAIL)

[INFO_VEHICLE_DYNAMIC 道路运输车辆动态监控记录表](#oper_INFO_VEHICLE_DYNAMIC)

[INFO_VEHICLE_GPS 道路运输车辆动态监控记录表（县/企业）](oper_INFO_VEHICLE_GPS)

[INFO_VEHICLE_HISTORY 道路运输车辆历史监控记录表](#oper_INFO_VEHICLE_HISTORY)

| Name                           | Type           | Description                                     | **Required** | default                    |
| ------------------------------ | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                             | int8           | 按指定方法生成                                  | required     | 主键                       |
| supervision_detail_id          | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| supervision_id                 | text           | 动态监管抽查主表的ID                            | optional     |                            |
| driver_id                      | text           | 驾驶员id                                        | optional     |                            |
| enterprise_id                  | text           | 车辆所在企业id                                  | optional     |                            |
| vehicle_id                     | text           | 车辆id                                          | optional     |                            |
| monitoring_time                | timestamptz(6) | 监控平台时间                                    | optional     |                            |
| monitoring_location            | text           | 监控平台显示位置                                | optional     |                            |
| is_online                      | boolean        | 是否在线（是/否）                               | optional     |                            |
| is_tachograph_record_normal    | boolean        | 行车记录仪数据（是否异常）                      | optional     |                            |
| tachograph_data_exception      | text           | 行车记录仪异常数据项                            | optional     |                            |
| tachograph_speed               | text           | 行车记录仪速度                                  | optional     |                            |
| gps_speed                      | text           | 卫星定位速度                                    | optional     |                            |
| is_speeding                    | boolean        | 是否超速（是/否）                               | optional     |                            |
| taxi_state                     | int4           | 出租空/重车状态（空/重）                        | optional     | **出租车空/重车状态**字典  |
| is_fatigue_driving             | boolean        | 客运疲劳驾驶（是/否）                           | optional     |                            |
| is_morning_outage              | boolean        | 客运凌晨2-5时停运（是/否）                      | optional     |                            |
| curve                          | int4           | 曲线情况（曲线完整/回传异常/零速度）            | optional     | **曲线情况**字典           |
| trail                          | int4           | 轨迹情况（正常/漂移/其他）                      | optional     | **GPS轨迹情况**字典        |
| lens_position                  | int4           | 镜头位置（正/偏）                               | optional     | **镜头位置**字典           |
| equipment                      | int4           | 设备情况（图像正常/无图像/摄像头 号损坏）       | optional     | **设备情况**字典           |
| other_infraction               | text           | 其他违规                                        | optional     |                            |
| disposal_measures              | text           | 处置措施                                        | optional     |                            |
| disposal_results               | text           | 处置结果                                        | optional     |                            |
| treatment_time                 | timestamptz(6) | 处置时间                                        | optional     |                            |
| assignee                       | text           | 受理人                                          | optional     |                            |
| feedback_time                  | timestamptz(6) | 反馈时间                                        | optional     |                            |
| remarks                        | text           | 备注                                            | optional     |                            |
| others                         | text           | 轨迹其他情况                                    | optional     |                            |
| lens_on                        | text           | 摄像头损坏号                                    | optional     |                            |
| monitor_end_time               | timestamptz(6) | 监管费到期时间                                  | optional     |                            |
| is_locate                      | boolean        | 是否定位                                        | optional     |                            |
| coordinate                     | point          | 空间数据类型point表示经纬度                     | optional     |                            |
| latitude_longitude_description | text           | 经纬度描述                                      | optional     |                            |
| is_send                        | boolean        | 是否发送                                        | optional     |                            |
| business_scope                 | int4           | 经营范围                                        | optional     | **经营范围**字典           |
| outage_alarm_time              | timestamptz(6) | 凌晨2点到5点停运报警时间                        | optional     |                            |
| speed_alarm_time               | timestamptz(6) | 超速报警时间                                    | optional     |                            |
| speeding_speed                 | text           | 超速速度                                        | optional     |                            |
| fatigue_alarm_time             | timestamptz(6) | 疲劳驾驶报警时间                                | optional     |                            |
| disposal_measures1             | text           | 是否在线处置措施                                | optional     |                            |
| disposal_measures2             | text           | 是否超速处置措施                                | optional     |                            |
| disposal_measures3             | text           | 曲线情况处置措施                                | optional     |                            |
| disposal_measures4             | text           | 客运疲劳驾驶处置措施                            | optional     |                            |
| disposal_measures5             | text           | 客运凌晨停运处置措施                            | optional     |                            |
| disposal_measures6             | text           | 行车记录仪数据处置措施                          | optional     |                            |
| disposal_measures7             | text           | 轨迹情况处置措施                                | optional     |                            |
| disposal_results1              | text           | 是否在线处置结果                                | optional     |                            |
| disposal_results2              | text           | 是否超速处置结果                                | optional     |                            |
| disposal_results3              | text           | 曲线情况处置结果                                | optional     |                            |
| disposal_results4              | text           | 客运疲劳驾驶处置结果                            | optional     |                            |
| disposal_results5              | text           | 客运疲劳驾驶处置结果                            | optional     |                            |
| disposal_results6              | text           | 行车记录仪数据处置结果                          | optional     |                            |
| disposal_results7              | text           | 轨迹情况处置结果                                | optional     |                            |
| dynamic_type                   | int4           | 动态抽查类型                                    | require      |                            |
| is_deleted                     | boolean        | 是否被删除                                      | optional     |                            |
| created_at                     | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                     | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                     | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                     | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                     | text           | 删除人                                          | optional     | **system_user表**的user_id |

### ent_dynamic_supervision_count 企业动态监管统计表

**根据以下数据表整合而成**

营运车:

[INFO_DYNAMIC_CHECK_RG_TJ  企业车辆动态监控台账排查统计表](#oper_INFO_DYNAMIC_CHECK_RG_TJ) 

| Name                   | Type           | Description                                     | **Required** | default                    |
| ---------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                     | int8           | 按指定方法生成                                  | required     | 主键                       |
| ent_dynamic_count_id   | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| ent_id                 | text           | 企业id                                          | optional     |                            |
| spot_checks_num        | int4           | 应抽查数                                        | optional     |                            |
| actual_spot_checks_num | int4           | 实抽查数                                        | optional     |                            |
| exception_num          | int4           | 异常数                                          | optional     |                            |
| disposal_num           | int4           | 处置数                                          | optional     |                            |
| area                   | text           | 地区                                            | optional     |                            |
| created_at             | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by             | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at             | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by             | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at             | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by             | text           | 删除人                                          | optional     | **system_user表**的user_id |

### dynamic_supervision 动态监管抽查主表

**根据以下数据表整合而成**

重点车:

[INFO_VEHICLE_GPS_MAIN 动态监管抽查主表](#major_INFO_VEHICLE_GPS_MAIN)<br />[INFO_VEHICLE_GPS_MAIN_BASE 动态监管抽查主表(临时)](#major_INFO_VEHICLE_GPS_MAIN_BASE)

| Name                    | Type           | Description                                                  | **Required** | default                        |
| ----------------------- | -------------- | ------------------------------------------------------------ | ------------ | ------------------------------ |
| id                      | int8           | 按指定方法生成                                               | required     | 主键                           |
| supervision_id          | text           | 动态监管抽查主表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                       |
| spot_check_date         | timestamptz(6) | 抽查日期                                                     | optional     |                                |
| spot_check_total_number | int4           | 抽查总数                                                     | optional     |                                |
| spot_check_number       | int4           | 抽查数量                                                     | optional     |                                |
| spot_check_ratio        | numeric        | 抽查比例                                                     | optional     |                                |
| check_user_id           | text           | 抽查人员                                                     | optional     | 引用**system_user表**的user_id |
| total_number_vehicle    | int4           | 总车辆数                                                     | optional     |                                |
| province_id             | text           | 抽查人员位置的省份ID                                         | optional     | 省份表province_id              |
| city_id                 | text           | 抽查人员位置的城市ID                                         | optional     | 城市表city_id                  |
| district_id             | text           | 抽查人员位置的区域ID                                         | optional     | 区域表district_id              |
| year                    | int4           | 年                                                           | optional     |                                |
| month                   | int4           | 月                                                           | optional     |                                |
| day                     | int4           | 日                                                           | optional     |                                |
| created_at              | timestamptz(6) | 创建时间                                                     | required     |                                |
| created_by              | text           | 创建人                                                       | required     | **system_user表**的user_id     |
| updated_at              | timestamptz(6) | 修改时间                                                     | optional     |                                |
| updated_by              | text           | 修改人                                                       | optional     | **system_user表**的user_id     |
| deleted_at              | timestamptz(6) | 删除时间                                                     | optional     |                                |
| deleted_by              | text           | 删除人                                                       | optional     | **system_user表**的user_id     |
| is_deleted              | boolean        | 是否被删除                                                   | optional     |                                |

### alarm_outage_date 客运凌晨两点至五点是否停运详细表

**根据以下数据表整合而成**

重点车:

[MON_DATE_ALARM 客运凌晨两点至五点是否停运详细表](#major_MON_DATE_ALARM) 

| Name                 | Type           | Description                                     | **Required** | default                    |
| -------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | int8           | 按指定方法生成                                  | required     | 主键                       |
| alarm_outage_date_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id               | text           | 人车分离车辆id                                  | optional     |                            |
| veh_cislc_id         | text           | cislcGPS车辆id                                  | optional     |                            |
| alarm_time           | timestamptz(6) | 报警时间                                        | optional     |                            |
| is_alarm             | boolean        | 是否报警                                        | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text           | 删除人                                          | optional     | **system_user表**的user_id |

### veh_offline_disposal 离线车辆处置表

**根据以下数据表整合而成**

重点车:

[MON_VEH_HANDLEDEAL 离线车辆处置表](#major_MON_VEH_HANDLEDEAL)

| Name                    | Type           | Description                                     | **Required** | default                                                      |
| ----------------------- | -------------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                      | int8           | 按指定方法生成                                  | required     | 主键                                                         |
| veh_offline_disposal_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| enterprise_id           | text           | 所在企业id                                      | optional     | **enterprise_info表**的enterprise_id                         |
| telephone               | text           | 手机号码                                        | optional     |                                                              |
| content                 | text           | 内容                                            | optional     |                                                              |
| send_time               | timestamptz(6) | 发送时间                                        | optional     |                                                              |
| user_id                 | text           | 用户ID                                          | optional     | **system_user表**的user_id                                   |
| is_sms_push             | boolean        | 是否短信推送                                    | optional     |                                                              |
| is_report               | boolean        | 是否通报                                        | optional     |                                                              |
| is_voice_notification   | boolean        | 是否语音通知                                    | optional     |                                                              |
| is_app_push             | boolean        | 是否APP推送                                     | optional     |                                                              |
| notification_content    | text           | 通报内容                                        | optional     |                                                              |
| voice_content           | text           | 语音内容                                        | optional     |                                                              |
| app_push_content        | text           | APP推送内容                                     | optional     |                                                              |
| supervision_detail_id   | text           | 抽查表ID                                        | optional     | **dynamic_supervision_detail 动态监管抽查明细表**的supervision_detail_id |
| disposal_method         | text           | 处置方式                                        | optional     |                                                              |
| is_deleted              | boolean        | 是否删除                                        | optional     |                                                              |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                                                              |
| created_by              | text           | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                                                              |
| updated_by              | text           | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                                                              |
| deleted_by              | text           | 删除人                                          | optional     | **system_user表**的user_id                                   |

## 统计查询(营运车)

### veh_daily_driving_summary 车辆每日行驶汇总情况

**根据以下数据表整合而成**

营运车:

[REPORT_LOCATION_MILE 车辆每日行驶汇总情况](#oper_REPORT_LOCATION_MILE) 

| Name                         | Type           | Description                                     | **Required** | default                    |
| ---------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | int8           | 按指定方法生成                                  | required     | 主键                       |
| veh_daily_driving_summary_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                   | text           | 车辆id                                          | optional     |                            |
| start_time                   | timestamptz(6) | 开始时间                                        | optional     |                            |
| deadline_time                | timestamptz(6) | 截止时间                                        | optional     |                            |
| coordinate                   | point          | 经纬度                                          | optional     |                            |
| speed                        | numeric        | 速度                                            | optional     |                            |
| direction                    | text           | 方向                                            | optional     |                            |
| is_located                   | boolean        | 是否定位                                        | optional     |                            |
| start_mileage                | numeric        | 开始里程                                        | optional     |                            |
| deadline_mileage             | numeric        | 截止里程                                        | optional     |                            |
| statistics_mileage           | numeric        | 统计里程                                        | optional     |                            |
| created_at                   | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                   | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                   | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                   | text           | 删除人                                          | optional     | **system_user表**的user_id |

### veh_online_rate 车辆在线率表

**根据以下数据表整合而成**

营运车:

[INFO_VEHPOSITION_QUE_RATE 车辆在线率表](#oper_INFO_VEHPOSITION_QUE_RATE)

| Name                    | Type           | Description                                            | **Required** | default                    |
| ----------------------- | -------------- | ------------------------------------------------------ | ------------ | -------------------------- |
| id                      | bigint         | 按指定方法生成                                         | required     | 主键                       |
| veh_online_time_id      | text           | 外部编码，由golang程序生成的xid，暴露到外部使用        | required     | 联合主键                   |
| vehicle_id              | text           | 车辆信息表 的vehicle_id                                | required     |                            |
| first_anchor_time       | timestamptz(6) | 当天的第一个定位点时间                                 | required     |                            |
| total_location_points   | int4           | 总定位点数                                             | required     |                            |
| data_qualified_points   | int4           | 数据合格点数                                           | required     |                            |
| data_unqualified_points | int4           | 数据不合格点数                                         | required     |                            |
| data_qualified_rate     | numeric        | 数据合格率                                             | required     |                            |
| data_unqualified_rate   | numeric        | 数据不合格率                                           | required     |                            |
| no_location_points      | int4           | 不定位点点数                                           | optional     |                            |
| today_total_online_time | int4           | 当天在线总时间（单位：秒）                             | optional     |                            |
| max_continuous_time     | int4           | 最大的连续在线时间（单位：秒）                         | optional     |                            |
| total_online_time       | int4           | 在线总时间                                             | optional     |                            |
| data_source             | text           | 数据来源类型                                           | optional     |                            |
| speed_points            | numeric        | 行车记录仪速度与GPS速度差值在3km/h以内（剔除不定位点） | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                               | required     |                            |
| created_by              | text           | 创建人                                                 | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                               | optional     |                            |
| updated_by              | text           | 修改人                                                 | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                               | optional     |                            |
| deleted_by              | text           | 删除人                                                 | optional     | **system_user表**的user_id |

### veh_online_time 车辆在线时长表

**根据以下数据表整合而成**

营运车:

[INFO_VEH_ONLINE 车辆在线时长表](#oper_INFO_VEH_ONLINE)

| Name               | Type           | Description                                     | **Required** | default                    |
| ------------------ | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                 | bigint         | 按指定方法生成                                  | required     | 主键                       |
| veh_online_time_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id         | text           | 车辆信息表 的vehicle_id                         | required     |                            |
| online_time        | numeric        | 在线时长                                        | optional     |                            |
| offline_time       | numeric        | 离线时长                                        | optional     |                            |
| total_time         | numeric        | 总时长                                          | optional     |                            |
| is_online          | boolean        | 是否在线                                        | optional     |                            |
| created_at         | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by         | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at         | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by         | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at         | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by         | text           | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_offline_registration 不在线报警登记

**根据以下数据表整合而成**

营运车:

[MON_OFFLINE_ALARM 不在线报警登记](#oper_MON_OFFLINE_ALARM)

| Name                          | Type           | Description                                     | **Required** | default                    |
| ----------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint         | 按指定方法生成                                  | required     | 主键                       |
| alarm_offline_registration_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                    | text           | 车辆信息表 的vehicle_id                         | required     |                            |
| offline_start_time            | timestamptz(6) | 离线开始时间                                    | optional     |                            |
| offline_end_time              | timestamptz(6) | 离线结束时间                                    | optional     |                            |
| registration_user             | text           | 登记用户                                        | optional     | **system_user表**的user_id |
| registration_time             | timestamptz(6) | 登记时间                                        | optional     |                            |
| sms_content                   | text           | 短信发送内容                                    | optional     |                            |
| phone_reminder_content        | text           | 电话提醒内容                                    | optional     |                            |
| sms_send_time                 | timestamptz(6) | 短信发送时间                                    | optional     |                            |
| phone_reminder_time           | timestamptz(6) | 电话提醒时间                                    | optional     |                            |
| offline_reason                | text           | 离线原因                                        | optional     |                            |
| alarm_type                    | integer        | 报警类型                                        | optional     | **报警类型**字典           |
| is_registration               | boolean        | 是否登记                                        | optional     |                            |
| is_end_alarm                  | boolean        | 是否结束报警                                    | optional     |                            |
| is_send_sms                   | boolean        | 是否发送短信                                    | optional     |                            |
| is_need_maintain              | boolean        | 是否需要维护                                    | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                    | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                    | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                    | text           | 删除人                                          | optional     | **system_user表**的user_id |

### client_online_status 客户端在线情况

**根据以下数据表整合而成**

营运车:

[REPORT_CLIENT_ONLINE 客户端在线情况](#oper_REPORT_CLIENT_ONLINE)

| Name                    | Type           | Description                                     | **Required** | default                    |
| ----------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint         | 按指定方法生成                                  | required     | 主键                       |
| client_online_status_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| user_id                 | text           | 用户id                                          | optional     |                            |
| start_time              | timestamptz(6) | 开始时间                                        | optional     |                            |
| end_time                | timestamptz(6) | 结束时间                                        | optional     |                            |
| heartbeat               | int4           | 心跳次数                                        | optional     |                            |
| is_updated              | boolean        | 是否更新                                        | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text           | 删除人                                          | optional     | **system_user表**的user_id |

### sms_remind_record 短信提醒记录

**根据以下数据表整合而成**

营运车:

[SMS_REMIND_LOG 短信提醒记录](#oper_SMS_REMIND_LOG)  

| Name                 | Type           | Description                                     | **Required** | default                    |
| -------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint         | 按指定方法生成                                  | required     | 主键                       |
| sms_remind_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id           | text           | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| alarm_time           | timestamptz(6) | 报警时间                                        | optional     |                            |
| alarm_type           | text           | 报警类型                                        | optional     |                            |
| sms_receiving_num    | text           | 短信接收号码                                    | optional     |                            |
| sms_recipient        | text           | 短信接收人                                      | optional     |                            |
| sms_content          | text           | 短信内容                                        | optional     |                            |
| input_person         | text           | 录入人                                          | optional     | **system_user表**的user_id |
| input_time           | timestamptz(6) | 录入时间                                        | optional     |                            |
| is_send              | boolean        | 是否发送                                        | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text           | 删除人                                          | optional     | **system_user表**的user_id |

### operator_data_statistice 运营商数据统计

**根据以下数据表整合而成**

营运车:

[INFO_JTJ_STATISTICAL_YYS 运营商数据统计](#oper_INFO_JTJ_STATISTICAL_YYS)

| Name                        | Type           | Description                                     | **Required** | default                    |
| --------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                          | bigint         | 按指定方法生成                                  | required     | 主键                       |
| operator_data_statistice_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| speeding_time               | int4           | 超速次数                                        | optional     |                            |
| full_mileage                | numeric        | 完整里程                                        | optional     |                            |
| total_mileage               | numeric        | 总里程                                          | optional     |                            |
| fatigue_total_time          | numeric        | 疲劳驾驶合计时间                                | optional     |                            |
| network_num                 | int4           | 入网数                                          | optional     |                            |
| online_num                  | int4           | 在线数                                          | optional     |                            |
| operator                    | text           | 运营商                                          | optional     |                            |
| night_operation_num         | int4           | 夜间营运次数                                    | optional     |                            |
| veh_num                     | int4           | 车辆数                                          | optional     |                            |
| operating_type              | int4           | 营运类型                                        | optional     | **营运类型**字典           |
| created_at                  | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                  | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                  | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                  | text           | 删除人                                          | optional     | **system_user表**的user_id |



## 缴费管理

### payment_receipt_info 缴费单据信息

**根据以下数据表整合而成**

营运车:

[FEE_BILL 缴费单据信息](#oper_FEE_BILL)

| Name                       | Type           | Description                                     | **Required** | default                    |
| -------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint         | 按指定方法生成                                  | required     | 主键                       |
| payment_receipt_info_id    | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| payment_ent                | text           | 缴费企业                                        | optional     |                            |
| payment_veh_num            | int4           | 缴费车辆数                                      | optional     |                            |
| payment_date               | timestamptz(6) | 缴费日期                                        | optional     |                            |
| total_payment_amount       | numeric        | 缴费总金额                                      | optional     |                            |
| payment_status             | int4           | 缴费状态                                        | optional     | **缴费状态**字典           |
| revocation_operator        | text           | 撤销操作人                                      | optional     |                            |
| revocation_time            | timestamptz(6) | 撤销时间                                        | optional     |                            |
| revocation_reason          | int4           | 撤销原因                                        | optional     | **缴费撤销原因**字典       |
| sim_total_fee              | numeric        | SIM卡费总额                                     | optional     |                            |
| sim_payment_method         | int4           | SIM卡费缴费方式                                 | optional     | **缴费方式**字典           |
| sim_payment_months         | int4           | SIM卡费缴费月数                                 | optional     |                            |
| service_total_fee          | numeric        | 服务费总额                                      | optional     |                            |
| service_payment_method     | int4           | 服务费缴费方式                                  | optional     | **缴费方式**字典           |
| service_payment_months     | int4           | 服务费缴费月数                                  | optional     |                            |
| service_payment_years      | int4           | 服务费缴费年数                                  | optional     |                            |
| supervision_total_fee      | numeric        | 监管费总额                                      | optional     |                            |
| supervision_payment_method | int4           | 监管费缴费方式                                  | optional     | **缴费方式**字典           |
| supervision_payment_months | int4           | 监管费缴费月数                                  | optional     |                            |
| payment_operator           | text           | 缴费操作人                                      | optional     |                            |
| audit_operator             | text           | 审核操作人                                      | optional     |                            |
| audit_time                 | timestamptz(6) | 审核时间                                        | optional     |                            |
| remarks                    | text           | 备注                                            | optional     |                            |
| fee_type                   | int4           | 费用类型                                        | optional     | **费用类型**字典           |
| revocation_fee_status      | int4           | 撤销缴费的状态                                  | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text           | 删除人                                          | optional     | **system_user表**的user_id |

### payment_receipt_detail 缴费单据明细

**根据以下数据表整合而成**

营运车:

[FEE_BILL_DETAIL 缴费单据明细](#oper_FEE_BILL_DETAIL)

| Name                           | Type           | Description        | **Required** | default                    |
| ------------------------------ | -------------- | ------------------ | ------------ | -------------------------- |
| id                             | bigint         | 按指定方法生成     | required     | 主键                       |
| payment_receipt_info_id        | text           | 缴费单据id         | required     | 联合主键                   |
| veh_id                         | text           | 车辆id             | optional     |                            |
| total_payment_amount           | numeric        | 缴费总金额         | optional     |                            |
| sim_month_fee                  | numeric        | SIM卡月费          | optional     |                            |
| sim_payment_start_date         | timestamptz(6) | SIM卡缴费起始日期  | optional     |                            |
| sim_payment_deadline           | timestamptz(6) | SIM卡缴费截止日期  | optional     |                            |
| sim_fee_months                 | int4           | SIM卡费月数        | optional     |                            |
| sim_total_fee                  | numeric        | SIM卡费用总额      | optional     |                            |
| service_month_fee              | numeric        | 服务月费           | optional     |                            |
| service_payment_start_date     | timestamptz(6) | 服务费缴费起始日期 | optional     |                            |
| service_payment_deadline       | timestamptz(6) | 服务费缴费截止日期 | optional     |                            |
| service_fee_years              | int4           | 服务费年数         | optional     |                            |
| service_fee_months             | int4           | 服务费月数         | optional     |                            |
| service_total_fee              | numeric        | 服务费用总额       | optional     |                            |
| supervision_month_fee          | numeric        | 监管月费           | optional     |                            |
| supervision_payment_start_date | timestamptz(6) | 监管费缴费起始日期 | optional     |                            |
| supervision_payment_deadline   | timestamptz(6) | 监管费缴费截止日期 | optional     |                            |
| supervision_fee_months         | int4           | 监管费月数         | optional     |                            |
| supervision_total_fee          | numeric        | 监管费用总额       | optional     |                            |
| created_at                     | timestamptz(6) | 创建时间           | required     |                            |
| created_by                     | text           | 创建人             | required     | **system_user表**的user_id |
| updated_at                     | timestamptz(6) | 修改时间           | optional     |                            |
| updated_by                     | text           | 修改人             | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz(6) | 删除时间           | optional     |                            |
| deleted_by                     | text           | 删除人             | optional     | **system_user表**的user_id |

### arrears_sms_record 欠费短信发送记录表

**根据以下数据表整合而成**

营运车:

[MON_FEEREMIND_SMS 欠费短信发送记录表](#oper_MON_FEEREMIND_SMS) 

| Name                  | Type           | Description                                     | **Required** | default                    |
| --------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                    | bigint         | 按指定方法生成                                  | required     | 主键                       |
| arrears_sms_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| sms_num               | text           | 短信号码                                        | optional     |                            |
| sms_content           | text           | 短信内容                                        | optional     |                            |
| license_plate_number  | text           | 车牌号                                          | optional     |                            |
| remind_type           | int4           | 提醒类型                                        | optional     | **欠费短信提醒类型**字典   |
| created_at            | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by            | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by            | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by            | text           | 删除人                                          | optional     | **system_user表**的user_id |

### driver_card_base_fee 司机卡基础费用

**根据以下数据表整合而成**

营运车:

[FEE_DRIVERBASEFEE 司机卡基础费用](#oper_FEE_DRIVERBASEFEE)

[FEE_DRVCARDBASEFEE 司机卡基础费用](#oper_FEE_DRVCARDBASEFEE)

| Name                          | Type           | Description                                     | **Required** | default                    |
| ----------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint         | 按指定方法生成                                  | required     | 主键                       |
| driver_card_base_fee_id       | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| driver_id                     | text           | 驾驶员id                                        | optional     |                            |
| card_fee                      | numeric        | 制卡费                                          | optional     |                            |
| need_pay_card_fee             | numeric        | 需缴制卡费                                      | optional     |                            |
| service_fee                   | numeric        | 服务费                                          | optional     |                            |
| need_pay_service_fee          | numeric        | 需缴服务费                                      | optional     |                            |
| service_fee_due_time          | timestamptz(6) | 服务费到期时间                                  | optional     |                            |
| service_fee_next_payment_date | timestamptz(6) | 服务费下次起始缴费日期                          | optional     |                            |
| is_manually                   | boolean        | 是否手工设置                                    | optional     |                            |
| registration_time             | timestamptz(6) | 注册时间                                        | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                    | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                    | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                    | text           | 删除人                                          | optional     | **system_user表**的user_id |

### driver_card_payment 司机卡缴费单

**根据以下数据表整合而成**

营运车:

[FEE_CARD_BILL 司机卡缴费主表](#oper_FEE_CARD_BILL)

[FEE_DRVCARBILL 司机卡缴费单](#oper_FEE_DRVCARBILL)

| Name                    | Type           | Description                                     | **Required** | default                    |
| ----------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint         | 按指定方法生成                                  | required     | 主键                       |
| driver_card_payment_id  | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| payment_ent             | text           | 缴费企业                                        | optional     |                            |
| payment_driver_card_num | int4           | 缴费司机卡数                                    | optional     |                            |
| payment_date            | timestamptz(6) | 缴费日期                                        | optional     |                            |
| total_payment_amount    | numeric        | 缴费总金额                                      | optional     |                            |
| payment_status          | int4           | 缴费状态                                        | optional     | **缴费状态**字典           |
| revocation_operator     | text           | 撤销操作人                                      | optional     |                            |
| revocation_time         | timestamptz(6) | 撤销时间                                        | optional     |                            |
| revocation_reason       | int4           | 撤销原因                                        | optional     | **缴费撤销原因**字典       |
| card_fee                | numeric        | 制卡费总额                                      | optional     |                            |
| card_fee_payment_method | int4           | 制卡费缴费方式                                  | optional     | **缴费方式**字典           |
| card_fee_month          | int4           | 制卡费缴费月数                                  | optional     |                            |
| service_total_fee       | numeric        | 服务费总额                                      | optional     |                            |
| service_payment_method  | int4           | 服务费缴费方式                                  | optional     | **缴费方式**字典           |
| service_payment_months  | int4           | 服务费缴费月数                                  | optional     |                            |
| service_payment_years   | int4           | 服务费缴费年数                                  | optional     |                            |
| payment_operator        | text           | 缴费操作员                                      | optional     |                            |
| audit_operator          | text           | 审核操作员                                      | optional     |                            |
| audit_time              | timestamptz(6) | 审核时间                                        | optional     |                            |
| remarks                 | text           | 备注                                            | optional     |                            |
| fee_type                | int4           | 费用类型                                        | optional     | **费用类型**字典           |
| revocation_fee_status   | int4           | 撤销缴费的状态                                  | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text           | 删除人                                          | optional     | **system_user表**的user_id |



### driver_card_payment 司机卡缴费明细

**根据以下数据表整合而成**

营运车:

[FEE_CARD_BILL_DETAIL 司机卡缴费明细](#oper_FEE_CARD_BILL_DETAIL)

[FEE_DRVCARBILL_DETAIL 司机卡缴费明细](#oper_FEE_DRVCARBILL_DETAIL)

| Name                           | Type           | Description        | **Required** | default                    |
| ------------------------------ | -------------- | ------------------ | ------------ | -------------------------- |
| id                             | bigint         | 按指定方法生成     | required     | 主键                       |
| driver_card_payment_id         | text           | 司机卡缴费单id     | optional     |                            |
| driver_card_id                 | text           | 司机卡编号         | optional     |                            |
| card_fee                       | numeric        | 制卡费             | optional     |                            |
| service_month_fee              | numeric        | 服务月费           | optional     |                            |
| total_payment_amount           | numeric        | 缴费总金额         | optional     |                            |
| driver_card_payment_start_date | timestamptz(6) | 司机卡缴费起始日期 | optional     |                            |
| driver_card_payment_deadline   | timestamptz(6) | 司机卡缴费截止日期 | optional     |                            |
| driver_card_fee_months         | int4           | 司机卡费月数       | optional     |                            |
| driver_card_total_fee          | numeric        | 司机卡费用总额     | optional     |                            |
| service_payment_start_date     | timestamptz(6) | 服务费缴费起始日期 | optional     |                            |
| service_payment_deadline       | timestamptz(6) | 服务费缴费截止日期 | optional     |                            |
| service_fee_months             | int4           | 服务费月数         | optional     |                            |
| service_fee_years              | int4           | 服务费年数         | optional     |                            |
| service_total_fee              | numeric        | 服务费用总额       | optional     |                            |
| created_at                     | timestamptz(6) | 创建时间           | required     |                            |
| created_by                     | text           | 创建人             | required     | **system_user表**的user_id |
| updated_at                     | timestamptz(6) | 修改时间           | optional     |                            |
| updated_by                     | text           | 修改人             | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz(6) | 删除时间           | optional     |                            |
| deleted_by                     | text           | 删除人             | optional     | **system_user表**的user_id |

### driver_card_revocation_record 司机卡缴费撤销记录表

**根据以下数据表整合而成**

营运车:

[FEE_CARD_BILL_RET_LOG 司机卡缴费撤销记录表](#oper_FEE_CARD_BILL_RET_LOG)

| Name                   | Type           | Description    | **Required** | default                    |
| ---------------------- | -------------- | -------------- | ------------ | -------------------------- |
| id                     | bigint         | 按指定方法生成 | required     | 主键                       |
| driver_card_payment_id | text           | 司机卡缴费单id | optional     |                            |
| revocation_status      | int4           | 撤销状态       | optional     | **撤销费用状态**字典       |
| created_at             | timestamptz(6) | 创建时间       | required     |                            |
| created_by             | text           | 创建人         | required     | **system_user表**的user_id |
| updated_at             | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by             | text           | 修改人         | optional     | **system_user表**的user_id |
| deleted_at             | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by             | text           | 删除人         | optional     | **system_user表**的user_id |

### driver_card_fee_transfer 司机卡费用转移表

**根据以下数据表整合而成**

营运车:

[FEE_CARD_MOVEBILL 司机卡费用转移表](#oper_FEE_CARD_MOVEBILL) 

| Name                   | Type           | Description                  | **Required** | default                    |
| ---------------------- | -------------- | ---------------------------- | ------------ | -------------------------- |
| id                     | bigint         | 按指定方法生成               | required     | 主键                       |
| driver_card_payment_id | text           | 费用转移单号(司机卡缴费单id) | optional     |                            |
| ent_id | text | 所在企业 | optional | |
| fee_status | int4 | 单据状态 | optional | **缴费状态**字典 |
| src_driver_cart | text | 源司机卡                     | optional |                            |
| src_driver_regis_time | timestamptz(6) | 源驾驶员注册日期             | optional |                            |
| src_driver_old_service_fee_deadline | timestamptz(6) | 源驾驶员原服务费截止日期     | optional |                            |
| src_driver_new_service_fee_deadline | timestamptz(6) | 源驾驶员新服务费截止日期     | optional |                            |
| transfer_date | timestamptz(6) | 转移日期                     | optional |                            |
| dest_driver_cart | text | 目的司机卡                   | optional |                            |
| dest_driver_regis_time | timestamptz(6) | 目的驾驶员注册日期           | optional |                            |
| dest_driver_old_service_fee_deadline | timestamptz(6) | 目的驾驶员原服务费截止日期   | optional |                            |
| dest_driver_new_service_fee_deadline | timestamptz(6) | 目的驾驶员最终服务费截止日期 | optional |                            |
| revocation_operator     | text           | 撤销操作人                                      | optional     |                            |
| revocation_time         | timestamptz(6) | 撤销时间                                        | optional     |                            |
| revocation_reason       | int4           | 撤销原因                  |optional||
| producer | text | 制单人 | optional |                            |
| produce_date | timestamptz(6) | 制单日期 | optional | |
| audit_operator | text | 审核操作人                   | optional     |                            |
| audit_time                           | timestamptz(6) | 审核时间                     | optional |                            |
| remarks | text | 备注 | optional |                            |
| created_at             | timestamptz(6) | 创建时间                     | required     |                            |
| created_by             | text           | 创建人                       | required     | **system_user表**的user_id |
| updated_at             | timestamptz(6) | 修改时间                     | optional     |                            |
| updated_by             | text           | 修改人                       | optional     | **system_user表**的user_id |
| deleted_at             | timestamptz(6) | 删除时间                     | optional     |                            |
| deleted_by             | text           | 删除人                       | optional     | **system_user表**的user_id |

## 人车分离

### driving_log_info 行车日志信息

**根据以下数据表整合而成**

重点车:

[VD_DRIVING_LOG 行车日志信息](#major_VD_DRIVING_LOG)  

| Name                | Type        | Description                                     | **Required** | default                        |
| ------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------ |
| id                  | bigint      | 按指定方法生成                                  | required     | 主键                           |
| driving_log_info_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                       |
| vehicle_id          | text        | 车辆id                                          | optional     | **vehicle_info表**的vehicle_id |
| driver_id           | text        | 驾驶员id                                        | optional     | **driver_info表**的driver_id   |
| driving_start_time  | timestamptz(6) | 用车起始日期                                    | optional     |                                |
| driving_end_time    | timestamptz(6) | 用车结束日期                                    | optional     |                                |
| cause               | text        | 事由                                            | optional     |                                |
| route               | text        | 路线                                            | optional     |                                |
| remarks             | text        | 备注                                            | optional     |                                |
| start_time          | text        | 开始时间                                        | optional     |                                |
| end_time            | text        | 结束时间                                        | optional     |                                |
| review_status       | integer     | 审核状态                                        | optional     |                                |
| review_agecy_level  | integer     | 审核机构级别                                    | optional     |                                |
| is_markup           | boolean     | 是否补录                                        | optional     |                                |
| created_at          | timestamptz(6) | 创建时间                                        | required     |                                |
| created_by          | text        | 创建人                                          | required     | **system_user表**的user_id     |
| updated_at          | timestamptz(6) | 修改时间                                        | optional     |                                |
| updated_by          | text        | 修改人                                          | optional     | **system_user表**的user_id     |
| deleted_at          | timestamptz(6) | 删除时间                                        | optional     |                                |
| deleted_by          | text        | 删除人                                          | optional     | **system_user表**的user_id     |


### driver_peccancy_check 驾驶员违法核实记录表


**根据以下数据表整合而成**

重点车:

[VD_PECCANCY_CHECK 驾驶员违法核实记录](#major_VD_PECCANCY_CHECK)

| Name          | Type           | Description    | **Required** | default                              |
| ------------- | -------------- | -------------- | ------------ | ------------------------------------ |
| id            | int8           | 按指定方法生成 | required     | 主键                                 |
| vehicle_id    | text           | 违章车辆id     | optional     | **vehicle_info表**的vehicle_id       |
| driver_id     | text           | 违章驾驶员id   | optional     | **driver_info表**的driver_id         |
| enterprise_id | text           | 所在企业id     | optional     | **enterprise_info表**的enterprise_id |
| created_at    | timestamptz(6) | 创建时间       | required     |                                      |
| created_by    | text           | 创建人         | required     | **system_user表**的user_id           |
| updated_at    | timestamptz(6) | 修改时间       | optional     |                                      |
| updated_by    | text           | 修改人         | optional     | **system_user表**的user_id           |
| deleted_at    | timestamptz(6) | 删除时间       | optional     |                                      |
| deleted_by    | text           | 删除人         | optional     | **system_user表**的user_id           |

### dispute_violation_record 违章争议记录表

**根据以下数据表整合而成**

重点车:

[VD_PECCANCY_VRGUE 争议违章记录](#major_VD_PECCANCY_VRGUE)

| Name                          | Type           | Description                                                  | **Required** | default                                                |
| ----------------------------- | -------------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------ |
| id                            | int8           | 按指定方法生成                                               | required     | 主键                                                   |
| dispute_violation_id          | text           | 违章争议记录表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                               |
| violation_detail_id           | text           | 违章明细表id                                                 | optional     | **vehicle_violation _details** 表的violation_detail_id |
| written_application_materials | text           | 书面申请材料                                                 | optional     |                                                        |
| labor_contract                | text           | 劳动合同或租赁合同                                           | optional     |                                                        |
| driving_log                   | text           | 行车日志                                                     | optional     |                                                        |
| witness                       | text           | 证人证言                                                     | optional     |                                                        |
| statement                     | text           | 当事人陈述                                                   | optional     |                                                        |
| pic_evidence                  | text           | 图像证据材料                                                 | optional     |                                                        |
| driver_license                | text           | 行为人驾驶证                                                 | optional     |                                                        |
| driving_license               | text           | 机动车行驶证                                                 | optional     |                                                        |
| id_card                       | text           | 行为人身份证                                                 | optional     |                                                        |
| business_license              | text           | 机动车所有人营业执照                                         | optional     |                                                        |
| organization_code             | text           | 机动车所有人组织机构代码证                                   | optional     |                                                        |
| legal_person_id_number        | text           | 法定代表人身份证                                             | optional     |                                                        |
| agent_id_number               | text           | 委托代理人身份证                                             | optional     |                                                        |
| vehicle_manager_id_card       | text           | 机动车管理人身份证                                           | optional     |                                                        |
| other_evidence                | text[]         | 其他证据材料                                                 | optional     |                                                        |
| approve_state                 | int4           | 审批状态                                                     | optional     | **车辆违法审批状态**字典                               |
| update_time_in                | timestamptz(6) | 内网更新时间                                                 | optional     |                                                        |
| contact_address               | text           | 联系地址                                                     | optional     |                                                        |
| created_at                    | timestamptz(6) | 创建时间                                                     | required     |                                                        |
| created_by                    | text           | 创建人                                                       | required     | **system_user表**的user_id                             |
| updated_at                    | timestamptz(6) | 修改时间                                                     | optional     |                                                        |
| updated_by                    | text           | 修改人                                                       | optional     | **system_user表**的user_id                             |
| deleted_at                    | timestamptz(6) | 删除时间                                                     | optional     |                                                        |
| deleted_by                    | text           | 删除人                                                       | optional     | **system_user表**的user_id                             |
| is_deleted                    | boolean        | 是否删除                                                     | optional     |                                                        |

### dispute_violation_record_log 违章争议审批日志表

**根据以下数据表整合而成**

重点车:

[VD_PECCANCY_VRGUE_LOG 违章争议审批日志表](#major_VD_PECCANCY_VRGUE_LOG)

| Name                     | Type           | Description                                                  | **Required** | default                                            |
| ------------------------ | -------------- | ------------------------------------------------------------ | ------------ | -------------------------------------------------- |
| id                       | int8           | 按指定方法生成                                               | required     | 主键                                               |
| dispute_violation_log_id | text           | 违章争议审批日志外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                           |
| dispute_violation_id     | text           | 违章争议记录表id                                             | optional     | **dispute_violation_record**的dispute_violation_id |
| reviewer                 | text           | 审核人                                                       | optional     | **system_user表**的user_id                         |
| review_time              | timestamptz(6) | 审核时间                                                     | optional     |                                                    |
| review_opinion           | text           | 审核意见                                                     | optional     |                                                    |
| review_result            | text           | 审核结果                                                     | optional     |                                                    |
| review_action_name       | text           | 审核动作名称                                                 | optional     |                                                    |
| approver                 | text           | 审批人                                                       | optional     | **system_user表**的user_id                         |
| update_time_in           | text           | 内网更新时间                                                 | optional     |                                                    |

### case_approval_review_operation 案件审批审核操作

**根据以下数据表整合而成**

重点车:

[VD_PECCANCY_OPERATION_LOG 案件审批审核操作](#major_VD_PECCANCY_OPERATION_LOG)

| Name                              | Type           | Description                                     | **Required** | default                    |
| --------------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                | int8           | 按指定方法生成                                  | required     | 主键                       |
| case_approval_review_operation_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| review_opinion                    | text           | 审核意见                                        | optional     |                            |
| review_result                     | text           | 审核结果                                        | optional     |                            |
| reviewer                          | text           | 审核人                                          | optional     |                            |
| review_time                       | timestamptz(6) | 审核时间                                        | optional     |                            |
| created_at                        | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                        | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                        | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                        | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                        | text           | 删除人                                          | optional     | **system_user表**的user_id |

### case_approval_review_call 案件审核审批电话告知

**根据以下数据表整合而成**

重点车:

[VD_PECCANCY_TEL_LOG 案件审核审批电话告知](#major_VD_PECCANCY_TEL_LOG) 

| Name                              | Type           | Description                                                  | **Required** | default                    |
| --------------------------------- | -------------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                                | int8           | 按指定方法生成                                               | required     | 主键                       |
| case_approval_review_call_id      | text           | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| case_approval_review_operation_id | text           | **case_approval_review_operation 案件审批审核操作**的case_approval_review_operation_id | optional     |                            |
| telephone_number                  | text           | 电话号码                                                     | optional     |                            |
| dial_time                         | timestamptz(6) | 拨打时间                                                     | optional     |                            |
| is_connected                      | boolean        | 是否接通                                                     | optional     |                            |
| inform_content                    | text           | 告知内容                                                     | optional     |                            |
| reviewer                          | text           | 审核人                                                       | optional     |                            |
| review_time                       | timestamptz(6) | 审核时间                                                     | optional     |                            |
| created_at                        | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by                        | text           | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                        | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by                        | text           | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by                        | text           | 删除人                                                       | optional     | **system_user表**的user_id |

## 交通违法信息

### vehicle_violation_details 车辆违章明细表

**根据以下数据表整合而成**

重点车:

[VD_PECCANCY_DETAIL 车辆违章明细](#major_VD_PECCANCY_DETAIL)

| Name                    | Type           | Description                                                 | **Required** | default                                                      |
| ----------------------- | -------------- | ----------------------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                      | int8           | 按指定方法生成                                              | required     | 主键                                                         |
| violation_detail_id     | text           | 车辆违章明细外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| vehicle_id              | text           | 违章车辆id                                                  | optional     | **vehicle_info表**的vehicle_id                               |
| driver_id               | text           | 违章驾驶员id                                                | optional     | **driver_info表**的driver_id                                 |
| enterprise_id           | text           | 所在企业id                                                  | optional     | **enterprise_info表**的enterprise_id                         |
| illegal_code            | text           | 违法代码                                                    | optional     | **VIO_CODEWFDM** 违法描述字典表                              |
| illegal_time            | timestamptz(6) | 违法时间                                                    | optional     |                                                              |
| illegal_handling_status | int4           | 违法处理状态                                                | optional     | **车辆违法处理状态**字典                                     |
| illegal_location        | text           | 违法地点                                                    | optional     |                                                              |
| standard_value          | text           | 标准值                                                      | optional     | 路段的限速阈值或核载的人数，根据违法的种类不同而不同。       |
| measured_value          | text           | 实测值                                                      | optional     | 车辆实际行驶的车速或实际载的人数，根据违法的种类不同而不同。 |
| discovery_agency        | text           | 发现机构                                                    | optional     |                                                              |
| illegal_photo           | text           | 违法照片                                                    | optional     |                                                              |
| is_notice_driver        | boolean        | 是否通知驾驶员                                              | optional     |                                                              |
| notice_time             | timestamptz(6) | 通知时间                                                    | optional     |                                                              |
| decision_number         | text           | 决定书号                                                    | optional     |                                                              |
| payment_mark            | int4           | 缴款标记                                                    | optional     | **是否缴款**字典                                             |
| party_name              | text           | 当事人姓名                                                  | optional     |                                                              |
| information_source      | int4           | 信息来源：1，强制，2，非现场，0，简易                       | optional     | **信息来源**字典表                                           |
| vehicle_information     | text           | 驾驶人处理的交通违法记录对应的机动车信息                    | optional     |                                                              |
| update_time_in          | timestamptz(6) | 内网更新时间                                                | optional     |                                                              |
| is_handle               | boolean        | 是否处理                                                    | optional     |                                                              |
| handle_by               | text           | 处理人                                                      | optional     | **system_user表**的user_id                                   |
| handle_at               | timestamptz(6) | 处理时间                                                    | optional     |                                                              |
| is_send                 | boolean        | 是否发送短信                                                | optional     |                                                              |
| is_deleted              | boolean        | 是否删除                                                    | optional     |                                                              |
| created_at              | timestamptz(6) | 创建时间                                                    | required     |                                                              |
| created_by              | text           | 创建人                                                      | required     | **system_user表**的user_id                                   |
| updated_at              | timestamptz(6) | 修改时间                                                    | optional     |                                                              |
| updated_by              | text           | 修改人                                                      | optional     | **system_user表**的user_id                                   |
| deleted_at              | timestamptz(6) | 删除时间                                                    | optional     |                                                              |
| deleted_by              | text           | 删除人                                                      | optional     | **system_user表**的user_id                                   |

### illegal_photo 违法照片表

**根据以下数据表整合而成**

重点车:

[VD_PICTURE_IN 违法照片](#major_VD_PICTURE_IN) 

| Name             | Type        | Description                                               | **Required** | default                    |
| ---------------- | ----------- | --------------------------------------------------------- | ------------ | -------------------------- |
| id               | bigint      | 按指定方法生成                                            | required     | 主键                       |
| illegal_photo_id | text        | 违法照片表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| picture_name     | text        | 违法照片名称                                              | optional     |                            |
| picture_address  | text        | 违法照片地址                                              | optional     |                            |
| is_synchronized  | boolean     | 是否同步                                                  | optional     | false                      |
| is_deleted       | boolean     | 是否删除                                                  | optional     | false                      |
| created_at       | timestamptz(6) | 创建时间                                                  | required     |                            |
| created_by       | text        | 创建人                                                    | required     | **system_user表**的user_id |
| updated_at       | timestamptz(6) | 修改时间                                                  | optional     |                            |
| updated_by       | text        | 修改人                                                    | optional     | **system_user表**的user_id |
| deleted_at       | timestamptz(6) | 删除时间                                                  | optional     |                            |
| deleted_by       | text        | 删除人                                                    | optional     | **system_user表**的user_id |

### veh_violation_scoring_items 车辆违规计分项表

**根据以下数据表整合而成**

重点车:

[INFO_VEHICLESCORESET 车辆违规计分项表](#major_INFO_VEHICLESCORESET)

| Name                       | Type        | Description                                                  | **Required** | default                    |
| -------------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                               | required     | 主键                       |
| violation_scoring_item_id  | text        | 车辆违规计分项表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| deduction_item_description | text        | 扣分事项描述                                                 | optional     |                            |
| deduction_category         | integer     | 扣分事项类别                                                 | optional     | **车辆评分扣分类别**字典   |
| demerit_points             | numeric     | 扣分分值                                                     | optional     |                            |
| is_deleted                 | boolean     | 是否删除                                                     | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by                 | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by                 | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by                 | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### veh_violation_scoring_record  车辆违规计分记录

**根据以下数据表整合而成**

重点车:

[INFO_VEHICLESCORELOG 车辆违规计分记录](#major_INFO_VEHICLESCORELOG)

| Name                      | Type        | Description                                                  | **Required** | default                                                      |
| ------------------------- | ----------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------------ |
| id                        | bigint      | 按指定方法生成                                               | required     | 主键                                                         |
| violation_scoring_id      | text        | 车辆违规计分记录外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| vehicle_id                | text        | 扣分车辆id                                                   | optional     | **vehicle_info表**的vehicle_id                               |
| violation_scoring_item_id | text        | 扣分明细id                                                   | optional     | **vehicle_violation_scoring_ items**表的violation_scoring_item_id |
| demerit_points            | numeric     | 扣分分值                                                     | optional     |                                                              |
| remarks                   | text        | 备注                                                         | optional     |                                                              |
| is_deleted                | boolean     | 是否删除                                                     | optional     |                                                              |
| created_at                | timestamptz(6) | 创建时间                                                     | required     |                                                              |
| created_by                | text        | 创建人                                                       | required     | **system_user表**的user_id                                   |
| updated_at                | timestamptz(6) | 修改时间                                                     | optional     |                                                              |
| updated_by                | text        | 修改人                                                       | optional     | **system_user表**的user_id                                   |
| deleted_at                | timestamptz(6) | 删除时间                                                     | optional     |                                                              |
| deleted_by                | text        | 删除人                                                       | optional     | **system_user表**的user_id                                   |

### vio_codewfdm 违法描述字典表

**根据以下数据表整合而成**

重点车:

[VIO_CODEWFDM 违法描述字典表](#major_VIO_CODEWFDM)

| Name     | Type    | Description    | **Required** | default |
| -------- | ------- | -------------- | ------------ | ------- |
| id       | int8    | 按指定方法生成 | required     | 主键    |
| WFXW     | text    | 违法行为       | required     |         |
| WFMS     | text    | 违法描述       | optional     |         |
| WFJFS    | numeric | 违法计分数     | optional     |         |
| FKJE_MIN | numeric | 最小罚款金额   | optional     |         |
| FKJE_MAX | numeric | 最大罚款金额   | optional     |         |
| XH       | text    | 序号           | optional     |         |

## 工程车相关数据表

### msg_reminder 消息提醒表

**根据以下数据表整合而成**

工程车:

MSGREMIND 消息提醒表

| Name            | Type        | Description                                                  | **Required** | default                    |
| --------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id              | bigint      | 按指定方法生成                                               | required     | 主键                       |
| msg_reminder_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| ent_id          | text        | 接收者ID(运输企业)                                           | optional     |                            |
| mag_id          | text        | 接收者ID(管理部门)                                           | optional     |                            |
| user_type       | integer     | 用户类型  1.渣土办  2.市政局  3.执法局  4.公安局             | optional     |                            |
| title           | text        | 标题                                                         | optional     |                            |
| content         | text        | 内容                                                         | optional     |                            |
| msg_type        | integer     | 消息类型 ： 0.审批  1.修改  2.报警  3.公司资料  4.驾驶员资料  5.车辆资料 11.车辆审核  12.驾驶员审核  13.教育培训  14.通知通告  15.意见反馈  16.车辆单位变更  17.车辆退出目录库  18.欠费通知  19.车辆不在线登记(车队申请)  --20.车辆不在线登记(管理部门审核)  21.事故登记  22.事故处理 | optional     |                            |
| msg_status      | integer     | 消息状态  0.未读  1.已读                                     | optional     |                            |
| area_id         | text        | 接收地区ID                                                   | optional     |                            |
| created_at      | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by      | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at      | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by      | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at      | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by      | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### veh_quarter_inspection 车辆季度查验表

**根据以下数据表整合而成**

工程车:

VEHICLEVERIFICATIONQUARTER 车辆季度查验表

| Name                      | Type        | Description                                     | **Required** | default                    |
| ------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_quarter_inspection_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                    | text        | 车辆ID                                          | optional     |                            |
| ent_id                    | text        | 企业ID                                          | optional     |                            |
| license_plate_number      | text        | 车牌号                                          | optional     |                            |
| work_num                  | text        | 工号牌                                          | optional     |                            |
| area_id                   | text        | 区域ID                                          | optional     |                            |
| year                      | int4        | 年份                                            | optional     |                            |
| season                    | int4        | 季度                                            | optional     |                            |
| auditor                   | text        | 审核人                                          | optional     |                            |
| audit_date                | timestamptz(6) | 审核时间                                        | optional     |                            |
| audit_times               | int4        | 审核次数                                        | optional     |                            |
| audit_status              | int4        | 审核状态                                        | optional     |                            |
| remarks                   | text        | 备注                                            | optional     |                            |
| submit_time               | timestamptz(6) | 提交时间                                        | optional     |                            |
| regis_time                | timestamptz(6) | 登记时间                                        | optional     |                            |
| is_deleted                | boolean     | 是否删除                                        | optional     |                            |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_quarter_inspection_item 车辆季度查验项  

**根据以下数据表整合而成**

工程车:

VEHICLEVERIFICATIONQUARTERDETAIL 车辆季度查验项

| Name                           | Type        | Description                                                  | **Required** | default                    |
| ------------------------------ | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                             | bigint      | 按指定方法生成                                               | required     | 主键                       |
| veh_quarter_inspection_item_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| veh_quarter_inspection_id      | text        | 车辆季度查验表id                                             | optional     |                            |
| inspection_item                | int4        | 查验项（1.反光材料 2.轮胎 3.灯光 4.右侧盲区设备 5.右侧盲区设备是否正常使用 6.栏板高度 7.密闭装置） | optional     |                            |
| file_url                       | text        | 文件地址                                                     | optional     |                            |
| regis_time                     | timestamptz(6) | 登记时间                                                     | optional     |                            |
| audit_status                   | int4        | 审核状态                                                     | optional     |                            |
| is_deleted                     | boolean     | 是否删除                                                     | optional     |                            |
| created_at                     | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by                     | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                     | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by                     | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by                     | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### veh_quarter_inspection_operation 车辆季度查验操作记录

**根据以下数据表整合而成**

工程车:

VEHICLEVERIFICATIONQUARTERLOG 车辆季度查验操作记录

| Name                                | Type        | Description                                     | **Required** | default                    |
| ----------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                  | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_quarter_inspection_operation_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_quarter_inspection_id           | text        | 车辆季度查验表id                                | optional     |                            |
| operation_type                      | int4        | 操作类型（1.企业提交 2.审核通过 3.退回）        | optional     |                            |
| remarks                             | text        | 备注                                            | optional     |                            |
| operation_platform                  | text        | 操作平台（gkpt/sjapp）                          | optional     |                            |
| created_at                          | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                          | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                          | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                          | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                          | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                          | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_radar_brand 右侧盲区视频监控及红外线雷达装置品牌型号

**根据以下数据表整合而成**

工程车:

VEHICLEVERIFICATIONALARMDEVICE 右侧盲区视频监控及红外线雷达装置品牌型号

| Name               | Type        | Description                                     | **Required** | default                    |
| ------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                 | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_radar_brand_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id             | text        | 车辆ID                                          | optional     |                            |
| factory            | text        | 厂家                                            | optional     |                            |
| brand              | text        | 品牌                                            | optional     |                            |
| model_type         | text        | 型号                                            | optional     |                            |
| img_url            | text        | 图片地址                                        | optional     |                            |
| is_deleted         | boolean     | 是否删除                                        | optional     |                            |
| created_at         | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by         | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at         | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by         | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at         | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by         | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_violation_upload 车辆违规上传表

**根据以下数据表整合而成**

工程车:

VEHICLEVIOLATION 车辆违规上传表

| Name                    | Type        | Description                                                  | **Required** | default                    |
| ----------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                      | bigint      | 按指定方法生成                                               | required     | 主键                       |
| veh_violation_upload_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| veh_id                  | text        | 车辆ID                                                       | optional     |                            |
| violation_type          | int4        | 违规类型（1.私自加高栏板 2.私自改变密封装置 3.违反分类管理规定 4.其他） | optional     |                            |
| violation_file          | text        | 违法证据文件                                                 | optional     |                            |
| remarks                 | text        | 备注                                                         | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by              | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by              | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by              | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### offline_registration 不在线登记记录表

**根据以下数据表整合而成**

工程车:

OUTAGEMSG 不在线登记记录表

| Name                    | Type        | Description                                     | **Required** | default                    |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                       |
| offline_registration_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| outage_registration_id  | text        | 停运登记表的id                                  | optional     |                            |
| registration_content    | text        | 登记内容                                        | optional     |                            |
| is_deleted              | boolean     | 是否删除                                        | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_speeding_warning_threshold 车辆超速预警阈值

**根据以下数据表整合而成**

工程车:

SET_VEHICLE_SPEED 车辆超速预警阈值

| Name                              | Type        | Description                                     | **Required** | default                    |
| --------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_speeding_warning_threshold_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                            | text        | 车辆ID                                          | optional     |                            |
| speed                             | text        | 限速阈值                                        | optional     |                            |
| is_deleted                        | boolean     | 是否删除                                        | optional     |                            |
| created_at                        | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                        | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                        | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                        | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                        | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_inspection_apply 验车申请记录表

**根据以下数据表整合而成**

工程车:

ApplicationCkeckVheicleLog 验车申请记录表

| Name                    | Type        | Description                                     | **Required** | default                    |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_inspection_apply_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                  | text        | 车辆ID                                          | optional     |                            |
| apply_person            | text        | 申请人                                          | optional     |                            |
| apply_group             | text        | 申请人用户组                                    | optional     |                            |
| apply_time              | timestamptz(6) | 申请时间                                        | optional     |                            |
| signer                  | text        | 签收人                                          | optional     |                            |
| signer_group            | text        | 签收人用户组                                    | optional     |                            |
| sign_time               | timestamptz(6) | 签收时间                                        | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_inspection_signee 验车签收人联系电话表

**根据以下数据表整合而成**

工程车:

CHECK_VEHICLE_USER 验车签收人联系电话表

| Name                     | Type        | Description                                     | **Required** | default                    |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_inspection_signee_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| area_code                | text        | 区域编码                                        | optional     |                            |
| name                     | text        | 姓名                                            | optional     |                            |
| phone                    | text        | 联系电话                                        | optional     |                            |
| department               | text        | 部门                                            | optional     |                            |
| is_deleted               | boolean     | 是否删除                                        | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id |

### driver_operation_log 驾驶员操作日志表

**根据以下数据表整合而成**

工程车:

DRIVEROPERATELOG 驾驶员操作日志表

| Name                    | Type        | Description                                                  | **Required** | default                    |
| ----------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                      | bigint      | 按指定方法生成                                               | required     | 主键                       |
| driver_operation_log_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| driver_id               | text        | 驾驶员ID                                                     | optional     |                            |
| driver_name             | text        | 驾驶员姓名                                                   | optional     |                            |
| driver_license_num      | text        | 驾驶证号                                                     | optional     |                            |
| veh_id                  | text        | 车辆ID                                                       | optional     |                            |
| license_plate_number    | text        | 车牌号                                                       | optional     |                            |
| ent_id                  | text        | 所属企业ID                                                   | optional     |                            |
| ent_name                | text        | 企业名称                                                     | optional     |                            |
| oper_type               | int4        | 操作类型（1.新增  2.编辑  3.删除  4.清退  5.审核  6.绑定车辆  7.解绑车辆  8.退出公司 9恢复审核资格） | optional     |                            |
| oper_content            | text        | 操作内容                                                     | optional     |                            |
| oper_time               | timestamptz(6) | 操作时间                                                     | optional     |                            |
| operator                | text        | 操作人                                                       | optional     |                            |
| operator_group          | text        | 操作人用户组                                                 | optional     |                            |
| keyword                 | text        | 关键字                                                       | optional     |                            |
| withdraw_reason         | text        | 清退原因（字典表DriverClearType）                            | optional     |                            |
| withdraw_remarks        | text        | 清退备注                                                     | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by              | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by              | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by              | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### driver_clear_type 清退原因字典

**根据以下数据表整合而成**

工程车:

DRIVERCLEARTYPE 清退原因字典

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| driver_clear_type_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| withdraw_reason      | text        | 清退原因                                        | optional     |                            |
| type                 | int4        | 类型（默认1）                                   | optional     |                            |
| is_deleted           | boolean     | 是否删除                                        | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### ent_state_file 企业状态文件表

**根据以下数据表整合而成**

工程车:

STATEFILE 企业状态文件表

| Name              | Type        | Description                                     | **Required** | default                    |
| ----------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                | bigint      | 按指定方法生成                                  | required     | 主键                       |
| ent_state_file_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| serial_num        | int8        | 序号                                            | optional     |                            |
| file_addr         | text        | 文件地址                                        | optional     |                            |
| created_at        | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by        | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at        | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by        | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at        | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by        | text        | 删除人                                          | optional     | **system_user表**的user_id |

### supplier_sign 供应商签到表

**根据以下数据表整合而成**

工程车:

INFO_SELLER_SIGN 供应商签到表

| Name             | Type        | Description                                     | **Required** | default                    |
| ---------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id               | bigint      | 按指定方法生成                                  | required     | 主键                       |
| supplier_sign_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| supplier         | text        | 供应商                                          | optional     |                            |
| oper_person      | text        | 操作人                                          | optional     |                            |
| sign_date        | timestamptz(6) | 签到日期                                        | optional     |                            |
| sign_type        | text        | 签到类型                                        | optional     |                            |
| year             | int4        | 年                                              | optional     |                            |
| month            | int4        | 月                                              | optional     |                            |
| day              | int4        | 日                                              | optional     |                            |
| created_at       | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by       | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at       | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by       | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at       | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by       | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_egn_purchase_intention 渣土车购车意向

**根据以下数据表整合而成**

重点车:

ZTC_GCYX 渣土车购车意向

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| supplier_sign_id     | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| supplier             | text        | 供应商                                          | optional     |                            |
| purchase_name        | text        | 购车用户姓名                                    | optional     |                            |
| purchase_phone       | text        | 购车用户电话                                    | optional     |                            |
| owner_ent_name       | text        | 所有人企业名称                                  | optional     |                            |
| owner_province       | text        | 所有人所在省                                    | optional     |                            |
| owner_city           | text        | 所有人所在市                                    | optional     |                            |
| owner_district       | text        | 所有人所在县                                    | optional     |                            |
| capacity_application | int4        | 运力申请                                        | optional     | 数字型字典5                |
| brand                | text        | 品牌型号                                        | optional     |                            |
| purchase_num         | int4        | 购车数量                                        | optional     |                            |
| coding               | text        | 编码                                            | optional     |                            |
| audit                | int4        | 审核                                            | optional     | 数字型字典115              |
| remarks              | text        | 备注                                            | optional     |                            |
| is_deleted           | boolean     | 是否删除                                        | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_uninstall_snapshot_system 工程未安装抓拍系统预警

**根据以下数据表整合而成**

工程车:

ECDFILEMATCHCONSTRUCTIONALARM 工程未安装抓拍系统预警

| Name                               | Type        | Description                                     | **Required** | default                    |
| ---------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                 | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_uninstall_snapshot_system_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| egn_order_num                      | text        | 工程单号                                        | optional     |                            |
| alarm_time                         | timestamptz(6) | 报警时间                                        | optional     |                            |
| is_cancel                          | boolean     | 是否撤销                                        | optional     |                            |
| cancel_time                        | timestamptz(6) | 撤销时间                                        | optional     |                            |
| created_at                         | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                         | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                         | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                         | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                         | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                         | text        | 删除人                                          | optional     | **system_user表**的user_id |

### snapshot_veh_whitelist 抓拍识别系统车辆白名单  

**根据以下数据表整合而成**

工程车:

FILEVEHPASSWHITELIST  抓拍识别系统车辆白名单

| Name                      | Type        | Description                                     | **Required** | default                    |
| ------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | bigint      | 按指定方法生成                                  | required     | 主键                       |
| snapshot_veh_whitelist_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number      | text        | 车牌号                                          | optional     |                            |
| ent_name                  | text        | 企业名称                                        | optional     |                            |
| owner_name                | text        | 车主姓名                                        | optional     |                            |
| owner_phone               | text        | 车主联系电话                                    | optional     |                            |
| construction_info_id      | text        | 进出工地表的ID                                  | optional     |                            |
| veh_front_pic             | text        | 车辆正面照片                                    | optional     |                            |
| veh_rear_pic              | text        | 车辆尾部照片                                    | optional     |                            |
| veh_sideways_pic          | text        | 车辆侧身照片                                    | optional     |                            |
| driving_license_veh_pic   | text        | 行驶证车辆照片                                  | optional     |                            |
| driving_license_pic       | text        | 行驶证照片                                      | optional     |                            |
| new_user                  | text        | 新增用户                                        | optional     |                            |
| user_group                | text        | 用户组                                          | optional     |                            |
| audit_status              | text        | 审核状态                                        | optional     |                            |
| auditor                   | text        | 审核人                                          | optional     |                            |
| audit_time                | timestamptz(6) | 审核时间                                        | optional     |                            |
| return_reason             | text        | 退回原因                                        | optional     |                            |
| is_deleted                | boolean     | 是否删除                                        | optional     |                            |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                | text        | 删除人                                          | optional     | **system_user表**的user_id |

### mq_send MQ发送记录

**根据以下数据表整合而成**

工程车:

MQTT_SEND MQTT发送记录

| Name         | Type        | Description                                     | **Required** | default                    |
| ------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| mq_send_id   | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| title        | text        | 标题                                            | optional     |                            |
| send_content | text        | 发送内容                                        | optional     |                            |
| status       | text        | 状态                                            | optional     |                            |
| created_at   | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at   | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at   | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### construction_camera 工地摄像头表

**根据以下数据表整合而成**

工程车:

FILECAMERA 工地摄像头表

| Name                   | Type        | Description                                     | **Required** | default                    |
| ---------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                     | bigint      | 按指定方法生成                                  | required     | 主键                       |
| construction_camera_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| sim_num                | text        | sim卡号                                         | optional     |                            |
| device_name            | text        | 设备名称                                        | optional     |                            |
| device_num             | text        | 设备编号                                        | optional     |                            |
| construction_info_id   | text        | 关联工地id                                      | optional     |                            |
| ip_address             | text        | IP地址                                          | optional     |                            |
| port                   | text        | 端口                                            | optional     |                            |
| is_deleted             | boolean     | 是否删除                                        | optional     |                            |
| created_at             | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by             | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at             | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by             | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at             | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by             | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_enter_out_constrution_manually 手动登记车辆进出工地

**根据以下数据表整合而成**

工程车:

FILEVEHINOROUT 手动登记车辆进出工地

| Name                   | Type        | Description                                     | **Required** | default                    |
| ---------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                     | bigint      | 按指定方法生成                                  | required     | 主键                       |
| construction_camera_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| construction_info_id   | text        | 工地id                                          | optional     |                            |
| ent_id                 | text        | 企业名称                                        | optional     |                            |
| veh_id                 | text        | 车辆ID                                          | optional     |                            |
| license_plate_number   | text        | 车牌号                                          | optional     |                            |
| leave_time             | timestamptz(6) | 离开时间                                        | optional     |                            |
| enter_time             | timestamptz(6) | 进入时间                                        | optional     |                            |
| driver_id              | text        | 驾驶员id                                        | optional     |                            |
| driver_phone           | text        | 驾驶员手机号                                    | optional     |                            |
| regis_time             | int4        | 登记类型（默认1）                               | optional     |                            |
| goods                  | text        | 货物                                            | optional     |                            |
| audit_status           | text        | 审核状态                                        | optional     |                            |
| auditor                | text        | 审核人                                          | optional     |                            |
| audit_time             | timestamptz(6) | 审核日期                                        | optional     |                            |
| is_alarm_verify        | boolean     | 是否已验证报警与否                              | optional     |                            |
| created_at             | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by             | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at             | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by             | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at             | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by             | text        | 删除人                                          | optional     | **system_user表**的user_id |

### site_engineering_association 工地工程关联表

**根据以下数据表整合而成**

工程车:

CONSTRUCTIONFILE 工地工程关联表

| Name                            | Type        | Description                                     | **Required** | default                    |
| ------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                              | bigint      | 按指定方法生成                                  | required     | 主键                       |
| site_engineering_association_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| construction_info_id            | text        | 工地id                                          | optional     |                            |
| engineering_id                  | text        | 工程id                                          | optional     |                            |
| created_at                      | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                      | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                      | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                      | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                      | text        | 删除人                                          | optional     | **system_user表**的user_id |

### camera_failure 摄像头故障表  

**根据以下数据表整合而成**

工程车:

CAMERAFAULT 摄像头故障表

| Name              | Type        | Description                                     | **Required**     | default                    |
| ----------------- | ----------- | ----------------------------------------------- | ---------------- | -------------------------- |
| id                | bigint      | 按指定方法生成                                  | required         | 主键                       |
| camera_failure_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required         | 联合主键                   |
| camera_id         | text        | 摄像头ID                                        | optional         |                            |
| status            | text        | 状态                                            | optional         |                            |
| operator          | text        | 操作人                                          | optional         |                            |
| operation_time    | timestamptz(6) | 操作时间                                        | optional         |                            |
| remarks           | text        | 备注                                            | optional         |                            |
| alarm_start_time  | timestamptz(6) | 开始报警时间                                    | optional         |                            |
| alarm_end_time    | timestamptz(6) | 结束报警时间                                    | optional         |                            |
| is_end            | boolean     | 是否结束                                        | optional         |                            |
| created_at        | timestamptz(6) | 创建时间                                        | optionalrequired |                            |
| created_by        | text        | 创建人                                          | required         | **system_user表**的user_id |
| updated_at        | timestamptz(6) | 修改时间                                        | optional         |                            |
| updated_by        | text        | 修改人                                          | optional         | **system_user表**的user_id |
| deleted_at        | timestamptz(6) | 删除时间                                        | optional         |                            |
| deleted_by        | text        | 删除人                                          | optional         | **system_user表**的user_id |

### construction_payment 工地缴费主表

**根据以下数据表整合而成**

工程车:

CONSTRUCTIONPAYMENTMAIN 工地缴费主表

| Name                    | Type        | Description                                     | **Required** | default                    |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                       |
| construction_payment_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| construction_info_id    | text        | 工地id                                          | optional     |                            |
| fee_deadline            | timestamptz(6) | 费用截止日期                                    | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id |

### construction_install_application 工地安装申请表

**根据以下数据表整合而成**

工程车:

ConstructionCameraServer 工地安装申请表

| Name                                | Type        | Description                                     | **Required** | default                    |
| ----------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                  | bigint      | 按指定方法生成                                  | required     | 主键                       |
| construction_install_application_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| construction_info_id                | text        | 工地id                                          | optional     |                            |
| operator                            | text        | 操作人                                          | optional     |                            |
| apply_time                          | timestamptz(6) | 申请时间                                        | optional     |                            |
| remarks                             | text        | 备注                                            |              |                            |
| created_at                          | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                          | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                          | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                          | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                          | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                          | text        | 删除人                                          | optional     | **system_user表**的user_id |

### camera_failure_handle 摄像头故障处理表

**根据以下数据表整合而成**

工程车:

CameraFaultDeal 摄像头故障处理表

| Name                     | Type        | Description                                     | **Required** | default                    |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                       |
| camera_failure_handle_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| camera_failure_id        | text        | 摄像头故障id                                    | optional     |                            |
| step                     | int4        | 步骤                                            | optional     |                            |
| operator                 | text        | 操作人                                          | optional     |                            |
| operation_time           | timestamptz(6) | 操作时间                                        | optional     |                            |
| remarks                  | text        | 备注                                            | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id |

### ent_rank_monthly_score 企业按月评分排名

**根据以下数据表整合而成**

工程车:

Last3CompanyMonth 企业按月评分排名

| Name                     | Type        | Description                                     | **Required** | default                    |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                       |
| camera_failure_handle_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| ent_id                   | text        | 企业ID                                          | optional     |                            |
| deduction_points         | numeric     | 扣分数                                          | optional     |                            |
| year                     | int4        | 年份                                            | optional     |                            |
| month                    | int4        | 月份                                            | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id |

### ent_rank_last_three_month 连续3月评分排名最后企业

**根据以下数据表整合而成**

工程车:

Continuous3MonthLast3 连续3月评分排名最后企业 

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| ent_rank_last_three_month_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| ent_id                       | text        | 企业ID                                          | optional     |                            |
| year                         | int4        | 年份                                            | optional     |                            |
| month                        | int4        | 月份                                            | optional     |                            |
| created_at                   | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_speed_fatigue 单月超速、疲劳驾驶报警5次以上驾驶员

**根据以下数据表整合而成**

工程车:

SpeedFatigueAlarm 单月超速、疲劳驾驶报警5次以上驾驶员

| Name                   | Type        | Description                                         | **Required** | default                                                      |
| ---------------------- | ----------- | --------------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                     | bigint      | 按指定方法生成                                      | required     | 主键                                                         |
| alarm_speed_fatigue_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用     | required     | 联合主键                                                     |
| veh_id                 | text        | 车辆ID                                              | optional     |                                                              |
| driver_id              | text        | 驾驶员id                                            | optional     |                                                              |
| driver_name            | text        | 驾驶员姓名                                          | optional     |                                                              |
| alarm_times            | int4        | 报警次数                                            | optional     |                                                              |
| alarm_type             | int4        | 报警类型（1.是否超速5次以上 2.是否疲劳驾驶5次以上） | optional     |                                                              |
| year                   | int4        | 年份                                                | optional     |                                                              |
| month                  | int4        | 月份                                                | optional     |                                                              |
| created_at             | timestamptz(6) | 创建时间                                            | required     |                                                              |
| created_by             | text        | 创建人                                              | required     | **system_user表**的user_id                                   |
| updated_at             | timestamptz(6) | 修改时间                                            | optional     |                                                              |
| updated_by             | text        | 修改人                                              | optional     | **system_user表**的user_id                                   |
| deleted_at             | timestamptz(6) | 删除时间                                            | optional     |                                                              |
| deleted_by             | text        | 删除人                                              | optional     | **system_user表**的user_ident_rank_last_three_month 连续3月评分排名最后企业 |

### offline_three_days_not_reason 无故不在线达3天以上车辆  

**根据以下数据表整合而成**

工程车:

OffLine3DayNotReason 无故不在线达3天以上车辆 

| Name                             | Type        | Description                                     | **Required** | default                                                      |
| -------------------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                               | bigint      | 按指定方法生成                                  | required     | 主键                                                         |
| offline_three_days_not_reason_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| veh_id                           | text        | 车辆ID                                          | optional     |                                                              |
| driver_id                        | text        | 驾驶员id                                        | optional     |                                                              |
| driver_name                      | text        | 驾驶员姓名                                      | optional     |                                                              |
| alarm_times                      | int4        | 报警次数                                        | optional     |                                                              |
| year                             | int4        | 年份                                            | optional     |                                                              |
| month                            | int4        | 月份                                            | optional     |                                                              |
| created_at                       | timestamptz(6) | 创建时间                                        | required     |                                                              |
| created_by                       | text        | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at                       | timestamptz(6) | 修改时间                                        | optional     |                                                              |
| updated_by                       | text        | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at                       | timestamptz(6) | 删除时间                                        | optional     |                                                              |
| deleted_by                       | text        | 删除人                                          | optional     | **system_user表**的user_ident_rank_last_three_month 连续3月评分排名最后企业 |

### alarm_deduction_num 当月违规扣分达到20分或连续2次停车学习驾驶员

**根据以下数据表整合而成**

工程车:

DeductionNumberAlarm 当月违规扣分达到20分或连续2次停车学习驾驶员

| Name                   | Type        | Description                                     | **Required** | default                                                      |
| ---------------------- | ----------- | ----------------------------------------------- | ------------ | ------------------------------------------------------------ |
| id                     | bigint      | 按指定方法生成                                  | required     | 主键                                                         |
| alarm_deduction_num_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                                     |
| veh_id                 | text        | 车辆ID                                          | optional     |                                                              |
| driver_id              | text        | 驾驶员id                                        | optional     |                                                              |
| driver_name            | text        | 驾驶员姓名                                      | optional     |                                                              |
| year                   | int4        | 年份                                            | optional     |                                                              |
| month                  | int4        | 月份                                            | optional     |                                                              |
| created_at             | timestamptz(6) | 创建时间                                        | required     |                                                              |
| created_by             | text        | 创建人                                          | required     | **system_user表**的user_id                                   |
| updated_at             | timestamptz(6) | 修改时间                                        | optional     |                                                              |
| updated_by             | text        | 修改人                                          | optional     | **system_user表**的user_id                                   |
| deleted_at             | timestamptz(6) | 删除时间                                        | optional     |                                                              |
| deleted_by             | text        | 删除人                                          | optional     | **system_user表**的user_ident_rank_last_three_month 连续3月评分排名最后企业 |

### construction_payment_apply 工地缴费申请表

**根据以下数据表整合而成**

工程车:

ConstructionPaymentApplication 工地缴费申请表

| Name                          | Type        | Description                                            | **Required** | default                    |
| ----------------------------- | ----------- | ------------------------------------------------------ | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                         | required     | 主键                       |
| construction_payment_apply_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用        | required     | 联合主键                   |
| construction_info_id          | text        | 工地id                                                 | optional     |                            |
| total_payment                 | numeric     | 缴费总额                                               | optional     |                            |
| monthly_payment               | numeric     | 缴费金额（单月）                                       | optional     |                            |
| payment_time                  | int4        | 缴费时长（月）                                         | optional     |                            |
| status                        | int4        | 状态（0：待审核，1：通过2：退回）                      | optional     |                            |
| payment_start_time            | timestamptz(6) | 缴费开始日期                                           | optional     |                            |
| payment_end_time              | timestamptz(6) | 缴费结束日期                                           | optional     |                            |
| return_reason                 | text        | 退回原因                                               | optional     |                            |
| is_cancel                     | boolean     | 是否撤销                                               | optional     |                            |
| cancel_reason                 | text        | 撤销原因                                               | optional     |                            |
| withdraw_status               | int4        | 撤销状态（1：通过，2：退回，3：待验证，4：待财务审核） | optional     |                            |
| withdraw_return_reason        | text        | 申请撤销财务退回原因                                   | optional     |                            |
| withdraw_user_id              | text        | 申请撤销用户ID                                         | optional     |                            |
| withdraw_verify_status        | int4        | 撤销验证状态（0：未验证，1：通过，2：退回）            | optional     |                            |
| withdraw_operator             | text        | 撤销验证操作人                                         | optional     |                            |
| is_deleted                    | boolean     | 是否删除                                               | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                               | required     |                            |
| created_by                    | text        | 创建人                                                 | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                               | optional     |                            |
| updated_by                    | text        | 修改人                                                 | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                               | optional     |                            |
| deleted_by                    | text        | 删除人                                                 | optional     | **system_user表**的user_id |

### alarm_unreport_engineering 未报备工程报警明细

**根据以下数据表整合而成**

工程车:

UnReportedAlarmMsg 未报备工程报警明细

| Name                          | Type        | Description                                          | **Required** | default                    |
| ----------------------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                       | required     | 主键                       |
| alarm_unreport_engineering_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| license_plate_number          | text        | 车牌号                                               | optional     |                            |
| ent_id                        | text        | 所属企业ID                                           | optional     |                            |
| veh_id                        | text        | 车辆ID                                               | optional     |                            |
| line_id                       | text        | 路线ID                                               | optional     |                            |
| gps_speed                     | numeric     | GPS速度                                              | optional     |                            |
| tachograph_speed              | numeric     | 行驶记录仪速度                                       | optional     |                            |
| alarm_start_time              | timestamptz(6) | 报警开始时间                                         | optional     |                            |
| alarm_end_time                | timestamptz(6) | 报警结束时间                                         | optional     |                            |
| is_supervision                | boolean     | 是否监管                                             | optional     |                            |
| coordinate                    | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| pos_describe                  | text        | 位置描述                                             | optional     |                            |
| area_id                       | text        | 所属地区                                             | optional     |                            |
| is_send_sms                   | boolean     | 是否发送短信                                         | optional     |                            |
| work_num                      | text        | 工号                                                 | optional     |                            |
| alarm_type                    | text        | 报警类型                                             | optional     |                            |
| egn_type                      | text        | 工程类型（A/B）                                      | optional     |                            |
| is_save_driver                | boolean     | 是否保存驾驶员                                       | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by                    | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by                    | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by                    | text        | 删除人                                               | optional     | **system_user表**的user_id |

### veh_engineering_trans_record  车辆工程运输过程记录

**根据以下数据表整合而成**

工程车:

AlarmMain_CompleteTrans 车辆工程运输过程记录

| Name                            | Type        | Description                                     | **Required** | default                    |
| ------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                              | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_engineering_trans_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                          | text        | 车辆ID                                          | optional     |                            |
| egn_order_num                   | text        | 工程单号                                        | optional     |                            |
| start_report_point_id           | text        | 开始运输的报备点ID（路线表对应的起点ID）        | optional     |                            |
| end_report_point_id             | text        | 结束运输的报备点ID（路线表对应的终点ID）        | optional     |                            |
| start_time                      | timestamptz(6) | 开始时间                                        | optional     |                            |
| end_time                        | timestamptz(6) | 结束时间                                        | optional     |                            |
| created_at                      | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                      | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                      | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                      | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                      | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_deviation_route_statistics 偏移路线报警日统计表

**根据以下数据表整合而成**

工程车:

OffsetRouteAlarmDay 偏移路线报警日统计表

| Name                                | Type        | Description                                     | **Required** | default                    |
| ----------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                  | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_deviation_route_statistics_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number                | text        | 车牌号                                          | optional     |                            |
| ent_id                              | text        | 所属企业                                        | optional     |                            |
| veh_id                              | text        | 车辆ID                                          | optional     |                            |
| alarm_times                         | int4        | 报警次数                                        | optional     |                            |
| alarm_time                          | timestamptz(6) | 报警时间                                        | optional     |                            |
| is_supervision                      | boolean     | 是否监管                                        | optional     |                            |
| is_send_sms                         | boolean     | 是否发送短信                                    | optional     |                            |
| area_id                             | text        | 所属地区                                        | optional     |                            |
| work_num                            | text        | 工号                                            | optional     |                            |
| created_at                          | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                          | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                          | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                          | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                          | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                          | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_speed_limit_statistics 限速报警日统计

**根据以下数据表整合而成**

工程车:

LimitSpeedAlarmDay 限速报警日统计

| Name                            | Type        | Description                                     | **Required** | default                    |
| ------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                              | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_speed_limit_statistics_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number            | text        | 车牌号                                          | optional     |                            |
| ent_id                          | text        | 所属企业                                        | optional     |                            |
| veh_id                          | text        | 车辆ID                                          | optional     |                            |
| alarm_times                     | int4        | 报警次数                                        | optional     |                            |
| alarm_time                      | timestamptz(6) | 报警时间                                        | optional     |                            |
| is_supervision                  | boolean     | 是否监管                                        | optional     |                            |
| is_send_sms                     | boolean     | 是否发送短信                                    | optional     |                            |
| area_id                         | text        | 所属地区                                        | optional     |                            |
| work_num                        | text        | 工号                                            | optional     |                            |
| created_at                      | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                      | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                      | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                      | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                      | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_time_limit_statistics 限时报警日统计

**根据以下数据表整合而成**

工程车:

LimitTimeAlarmDay 限时报警日统计

| Name                           | Type        | Description                                     | **Required** | default                    |
| ------------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                             | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_time_limit_statistics_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number           | text        | 车牌号                                          | optional     |                            |
| ent_id                         | text        | 所属企业                                        | optional     |                            |
| veh_id                         | text        | 车辆ID                                          | optional     |                            |
| alarm_times                    | int4        | 报警次数                                        | optional     |                            |
| alarm_time                     | timestamptz(6) | 报警时间                                        | optional     |                            |
| is_supervision                 | boolean     | 是否监管                                        | optional     |                            |
| is_send_sms                    | boolean     | 是否发送短信                                    | optional     |                            |
| area_id                        | text        | 所属地区                                        | optional     |                            |
| work_num                       | text        | 工号                                            | optional     |                            |
| created_at                     | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                     | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                     | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                     | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                     | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_unreport_statistics 未报备工程报警日统计

**根据以下数据表整合而成**

工程车:

UnReportedAlarmDay 未报备工程报警日统计  

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_unreport_statistics_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number         | text        | 车牌号                                          | optional     |                            |
| ent_id                       | text        | 所属企业                                        | optional     |                            |
| veh_id                       | text        | 车辆ID                                          | optional     |                            |
| alarm_times                  | int4        | 报警次数                                        | optional     |                            |
| alarm_time                   | timestamptz(6) | 报警时间                                        | optional     |                            |
| is_supervision               | boolean     | 是否监管                                        | optional     |                            |
| is_send_sms                  | boolean     | 是否发送短信                                    | optional     |                            |
| area_id                      | text        | 所属地区                                        | optional     |                            |
| work_num                     | text        | 工号                                            | optional     |                            |
| alarm_start_time             | timestamptz(6) | 报警开始时间                                    | optional     |                            |
| alarm_end_time               | timestamptz(6) | 报警结束时间                                    | optional     |                            |
| alarm_duration               | numeric     | 报警时长（秒）                                  | optional     |                            |
| alarm_type                   | text        | 报警类型                                        | optional     |                            |
| created_at                   | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_driver_related 报警驾驶员关联表

**根据以下数据表整合而成**

工程车:

ECDSpeedAlarmMsgDriver 超速报警驾驶员关联表

ECDAlarmMsgDriver 疲劳驾驶报警驾驶员关联表

ONLINEECDALARMMSGDRIVER 超三天断电报警驾驶员关联

| Name                    | Type        | Description                                     | **Required** | default                    |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_driver_related_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| alarm_id                | text        | 报警ID                                          | optional     |                            |
| driver_id               | text        | 驾驶员ID                                        | optional     |                            |
| driver_name             | text        | 驾驶员姓名                                      | optional     |                            |
| telephone               | text        | 手机号码                                        | optional     |                            |
| driver_license          | text        | 驾驶证号                                        | optional     |                            |
| alarm_type              | text        | 报警类型                                        | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_ent_send_police 企业报警发送交警联系人

**根据以下数据表整合而成**

工程车:

PoliceSmsInfo 企业报警发送交警联系人

| Name                     | Type        | Description                                     | **Required** | default                    |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_ent_send_police_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| ent_id                   | text        | 企业ID                                          | optional     |                            |
| ent_contact              | text        | 企业联系人                                      | optional     |                            |
| ent_phone                | text        | 企业联系电话                                    | optional     |                            |
| police_id                | text        | 交警id                                          | optional     |                            |
| police_phone             | text        | 交警联系电话                                    | optional     |                            |
| police_department        | text        | 交警所属部门                                    | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id |

### sms_send_log 短信发送日志

**根据以下数据表整合而成**

工程车:

ECDSmsLog 短信发送日志

| Name            | Type        | Description                                                  | **Required** | default                    |
| --------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id              | bigint      | 按指定方法生成                                               | required     | 主键                       |
| sms_send_log_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| veh_id          | text        | 车辆id                                                       | optional     |                            |
| driver_id       | text        | 驾驶员id                                                     | optional     |                            |
| driver_phone    | text        | 驾驶员联系电话                                               | optional     |                            |
| sms_content     | text        | 短信内容                                                     | optional     |                            |
| operator_area   | text        | 操作用户所属地区                                             | optional     |                            |
| alarm_type      | int4        | 报警类型  1.超速报警  2.安检报警  3.疲劳驾驶  4.超三天断电  8.工程报警  11.进出工地报警 | optional     |                            |
| created_at      | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by      | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at      | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by      | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at      | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by      | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### area_veh_enter 前各县市车辆进入本辖区车辆

**根据以下数据表整合而成**

工程车:

AREAVEHICELEPID 当前各县市车辆进入本辖区车辆

| Name              | Type        | Description                                          | **Required** | default                    |
| ----------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                | bigint      | 按指定方法生成                                       | required     | 主键                       |
| area_veh_enter_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| veh_id            | text        | 车辆ID                                               | optional     |                            |
| area_id           | text        | 区域ID                                               | optional     |                            |
| enter_time        | timestamptz(6) | 进入时间                                             | optional     |                            |
| leave_time        | timestamptz(6) | 离开时间                                             | optional     |                            |
| interval          | int4        | 间隔时间                                             | optional     |                            |
| is_located        | boolean     | 是否定位                                             | optional     |                            |
| coordinate        | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| speed             | numeric     | 速度                                                 | optional     |                            |
| created_at        | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by        | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at        | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by        | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at        | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by        | text        | 删除人                                               | optional     | **system_user表**的user_id |

### alarm_veh_enter_overtime 各县市车辆进入本辖区域停留超过24小时的报警统计表

**根据以下数据表整合而成**

工程车:

ALARMAREAVEHICELE 各县市车辆进入本辖区域停留超过24小时的报警统计表

| Name                        | Type        | Description                                                  | **Required** | default                    |
| --------------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                          | bigint      | 按指定方法生成                                               | required     | 主键                       |
| alarm_veh_enter_overtime_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| veh_id                      | text        | 车辆ID                                                       | optional     |                            |
| area_id                     | text        | 区域ID                                                       | optional     |                            |
| enter_time                  | timestamptz(6) | 进入时间                                                     | optional     |                            |
| leave_time                  | timestamptz(6) | 离开时间                                                     | optional     |                            |
| interval                    | int4        | 间隔时间                                                     | optional     |                            |
| is_located                  | boolean     | 是否定位                                                     | optional     |                            |
| coordinate                  | point       | 空间数据类型point表示经度(longitude)和纬度(latitude)         | optional     |                            |
| speed                       | numeric     | 速度                                                         | optional     |                            |
| alarm_type                  | int4        | 报警类型 1.各县市车辆24小时内进出本辖区域超过3次的报警 2.各县市车辆进入本辖区域停留超过24小时的报警 | optional     |                            |
| is_supervise                | int4        | 是否监管 0.未监管 1.已监管                                   | optional     |                            |
| supervision_content         | text        | 监管内容                                                     | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by                  | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by                  | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by                  | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### alarm_deviation_route_detail 偏移路线报警明细表

**根据以下数据表整合而成**

工程车:

OffsetRouteAlarmMsg 偏移路线报警明细表

| Name                            | Type        | Description                                          | **Required** | default                    |
| ------------------------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                              | bigint      | 按指定方法生成                                       | required     | 主键                       |
| alarm_deviation_route_detail_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| license_plate_number            | text        | 车牌号                                               | optional     |                            |
| ent_id                          | text        | 所属企业ID                                           | optional     |                            |
| veh_id                          | text        | 车辆ID                                               | optional     |                            |
| line_id                         | text        | 路线ID                                               | optional     |                            |
| gps_speed                       | numeric     | GPS速度                                              | optional     |                            |
| tachograph_speed                | numeric     | 行驶记录仪速度                                       | optional     |                            |
| alarm_start_time                | timestamptz(6) | 报警开始时间                                         | optional     |                            |
| alarm_end_time                  | timestamptz(6) | 报警结束时间                                         | optional     |                            |
| is_supervision                  | boolean     | 是否监管                                             | optional     |                            |
| coordinate                      | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| pos_describe                    | text        | 位置描述                                             | optional     |                            |
| area_id                         | text        | 所属地区                                             | optional     |                            |
| is_send_sms                     | boolean     | 是否发送短信                                         | optional     |                            |
| work_num                        | text        | 工号                                                 | optional     |                            |
| alarm_type                      | text        | 报警类型                                             | optional     |                            |
| egn_type                        | text        | 工程类型（A/B）                                      | optional     |                            |
| is_save_driver                  | boolean     | 是否保存驾驶员                                       | optional     |                            |
| created_at                      | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by                      | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at                      | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by                      | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by                      | text        | 删除人                                               | optional     | **system_user表**的user_id |

### alarm_speed_limit_detail 限速报警明细

**根据以下数据表整合而成**

工程车:

LimitSpeedAlarmMsg 限速报警明细

| Name                        | Type        | Description                                          | **Required** | default                    |
| --------------------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                          | bigint      | 按指定方法生成                                       | required     | 主键                       |
| alarm_speed_limit_detail_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| license_plate_number        | text        | 车牌号                                               | optional     |                            |
| ent_id                      | text        | 所属企业ID                                           | optional     |                            |
| veh_id                      | text        | 车辆ID                                               | optional     |                            |
| line_id                     | text        | 路线ID                                               | optional     |                            |
| gps_speed                   | numeric     | GPS速度                                              | optional     |                            |
| tachograph_speed            | numeric     | 行驶记录仪速度                                       | optional     |                            |
| alarm_start_time            | timestamptz(6) | 报警开始时间                                         | optional     |                            |
| alarm_end_time              | timestamptz(6) | 报警结束时间                                         | optional     |                            |
| is_supervision              | boolean     | 是否监管                                             | optional     |                            |
| coordinate                  | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| pos_describe                | text        | 位置描述                                             | optional     |                            |
| area_id                     | text        | 所属地区                                             | optional     |                            |
| is_send_sms                 | boolean     | 是否发送短信                                         | optional     |                            |
| work_num                    | text        | 工号                                                 | optional     |                            |
| alarm_type                  | text        | 报警类型                                             | optional     |                            |
| egn_type                    | text        | 工程类型（A/B）                                      | optional     |                            |
| is_save_driver              | boolean     | 是否保存驾驶员                                       | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by                  | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by                  | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by                  | text        | 删除人                                               | optional     | **system_user表**的user_id |

### alarm_time_limit_detail 限时报警明细

**根据以下数据表整合而成**

工程车:

LimitTimeAlarmMsg 限时报警明细

| Name                       | Type        | Description                                          | **Required** | default                    |
| -------------------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                       | required     | 主键                       |
| alarm_time_limit_detail_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| license_plate_number       | text        | 车牌号                                               | optional     |                            |
| ent_id                     | text        | 所属企业ID                                           | optional     |                            |
| veh_id                     | text        | 车辆ID                                               | optional     |                            |
| line_id                    | text        | 路线ID                                               | optional     |                            |
| gps_speed                  | numeric     | GPS速度                                              | optional     |                            |
| tachograph_speed           | numeric     | 行驶记录仪速度                                       | optional     |                            |
| alarm_start_time           | timestamptz(6) | 报警开始时间                                         | optional     |                            |
| alarm_end_time             | timestamptz(6) | 报警结束时间                                         | optional     |                            |
| is_supervision             | boolean     | 是否监管                                             | optional     |                            |
| coordinate                 | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| pos_describe               | text        | 位置描述                                             | optional     |                            |
| area_id                    | text        | 所属地区                                             | optional     |                            |
| is_send_sms                | boolean     | 是否发送短信                                         | optional     |                            |
| work_num                   | text        | 工号                                                 | optional     |                            |
| alarm_type                 | text        | 报警类型                                             | optional     |                            |
| egn_type                   | text        | 工程类型（A/B）                                      | optional     |                            |
| is_save_driver             | boolean     | 是否保存驾驶员                                       | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by                 | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by                 | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by                 | text        | 删除人                                               | optional     | **system_user表**的user_id |

### veh_status 新型车辆状态

**根据以下数据表整合而成**

工程车:

ZTC_INFO_OPERATION_STATE 新型车辆状态

| Name                     | Type        | Description                                     | **Required** | default                    |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_status_id            | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                   | text        | 车辆ID                                          | optional     |                            |
| lock_status              | int4        | 锁车状态                                        | optional     |                            |
| limit_speed_status       | int4        | 限速状态                                        | optional     |                            |
| limit_lift_status        | int4        | 限举状态                                        | optional     |                            |
| control_unlock_status    | int4        | 管控解锁状态                                    | optional     |                            |
| certificate_status       | int4        | 核准证状态                                      | optional     |                            |
| fingerprint_status       | int4        | 指纹状态                                        | optional     |                            |
| query_fingerprint_status | int4        | 查询指纹状态                                    | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_online_status 新型车辆在线状态

**根据以下数据表整合而成**

工程车:

SmartCar_Online 新型车辆在线状态

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_online_status_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id               | text        | 车辆ID                                          | optional     |                            |
| license_plate_number | text        | 车牌号                                          | optional     |                            |
| last_locate_time     | timestamptz(6) | 最后定位时间                                    | optional     |                            |
| acc_status           | int4        | Acc状态（0.关 1.开）                            | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_egn_terminal 新型渣土车终端      

**根据以下数据表整合而成**

工程车:

INFO_SMARTCAR_TERMINAL 新型渣土车终端

| Name                   | Type        | Description                                     | **Required** | default                    |
| ---------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                     | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_egn_terminal_id    | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                 | text        | 车辆ID                                          | optional     |                            |
| terminal_id            | text        | 终端ID                                          | optional     |                            |
| is_test_pass           | boolean     | 是否测试通过                                    | optional     |                            |
| is_terminal_simulation | boolean     | 是否模拟终端                                    | optional     |                            |
| is_deleted             | boolean     | 是否删除                                        | optional     |                            |
| created_at             | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by             | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at             | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by             | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at             | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by             | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_send_screen_info 发送车载屏信息     

**根据以下数据表整合而成**

工程车:

INFO_SMARTCAR_SEND_MSG 发送车载屏信息

| Name                    | Type        | Description                                     | **Required** | default                    |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_send_screen_info_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| send_obj                | text        | 发送对象类型（1.辖区 2.运输主体 3.车辆）        | optional     |                            |
| send_content            | text        | 发送内容                                        | optional     |                            |
| user_group              | text        | 用户组                                          | optional     |                            |
| title                   | text        | 标题                                            | optional     |                            |
| is_deleted              | boolean     | 是否删除                                        | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_screen_info_send_obj 车载屏信息发送对象   

**根据以下数据表整合而成**

工程车:

INFO_SMARTCAR_SEND_MSG_TARGET 车载屏信息发送对象

| Name                        | Type        | Description                                     | **Required** | default                    |
| --------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                          | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_screen_info_send_obj_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_send_screen_info_id     | text        | 发送车载屏信息id                                | optional     |                            |
| send_obj_id                 | text        | 发送对象ID                                      | optional     |                            |
| certificate_id              | text        | 核准证ID                                        | optional     |                            |
| is_deleted                  | boolean     | 是否删除                                        | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                  | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                  | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                  | text        | 删除人                                          | optional     | **system_user表**的user_id |

### fingerprint_check_in_record 指纹打卡记录

**根据以下数据表整合而成**

工程车:

INFO_SMARTCAR_ZWDK 指纹打卡记录

| Name                           | Type        | Description                                          | **Required** | default                    |
| ------------------------------ | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                             | bigint      | 按指定方法生成                                       | required     | 主键                       |
| fingerprint_check_in_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| check_in_time                  | timestamptz(6) | 打卡时间                                             | optional     |                            |
| veh_id                         | text        | 车辆ID                                               | optional     |                            |
| license_plate_number           | text        | 车牌号                                               | optional     |                            |
| terminal_id                    | text        | 终端ID                                               | optional     |                            |
| user_id                        | text        | 用户ID                                               | optional     |                            |
| certificate_id                 | text        | 核准证ID                                             | optional     |                            |
| is_violation                   | boolean     | 是否违规                                             | optional     |                            |
| coordinate                     | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| created_at                     | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by                     | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at                     | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by                     | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by                     | text        | 删除人                                               | optional     | **system_user表**的user_id |

### veh_driver_fingerprint_related 车辆驾驶员指纹绑定

**根据以下数据表整合而成**

工程车:

ZTC_VEH_FINGERPRINT 车辆驾驶员指纹绑定

| Name                              | Type        | Description                                     | **Required** | default                    |
| --------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_driver_fingerprint_related_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| sim_num                           | text        | SIM卡号                                         | optional     |                            |
| veh_id                            | text        | 车辆ID                                          | optional     |                            |
| license_plate_number              | text        | 车牌号                                          | optional     |                            |
| driver_id                         | text        | 驾驶员ID                                        | optional     |                            |
| driver_name                       | text        | 姓名                                            | optional     |                            |
| fingerprint_id                    | text        | 指纹ID                                          | optional     |                            |
| is_deleted                        | boolean     | 是否删除                                        | optional     |                            |
| created_at                        | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                        | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                        | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                        | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                        | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_fingerprint_set 车辆指纹状态设置

**根据以下数据表整合而成**

工程车:

VehFingerFunctionManager 车辆指纹状态设置

| Name                   | Type        | Description                                     | **Required** | default                    |
| ---------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                     | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_fingerprint_set_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| sim_num                | text        | SIM卡号                                         | optional     |                            |
| is_oper                | boolean     | 是否启用                                        | optional     |                            |
| created_at             | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by             | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at             | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by             | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at             | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by             | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_destroy_terminal 疑似破坏终端报警

**根据以下数据表整合而成**

工程车:

INFO_SMARTCAR_ALARM_PHZD 疑似破坏终端报警

| Name                      | Type           | Description                                                  | **Required** | default                    |
| ------------------------- | -------------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                        | bigint         | 按指定方法生成                                               | required     | 主键                       |
| alarm_destroy_terminal_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| vehicle_id                | text           | 车辆ID                                                       | optional     |                            |
| alarm_start_time          | timestamptz(6) | 报警开始时间                                                 | optional     |                            |
| alarm_end_time            | timestamptz(6) | 报警结束时间                                                 | optional     |                            |
| is_handle                 | boolean        | 是否处理                                                     | optiona      |                            |
| processing_time           | timestamptz(6) | 处理时间                                                     | optional     |                            |
| processing_remarks        | text           | 处理备注                                                     | optional     |                            |
| processor                 | text           | 处理人                                                       | optional     |                            |
| is_ent_handle             | boolean        | 企业是否处理                                                 | optional     |                            |
| ent_handle_time           | timestamptz(6)    | 企业处理时间                                                 | optional     |                            |
| coordinate                | point          | 空间数据类型point表示经度(longitude)和纬度(latitude)         | optional     |                            |
| alarm_type                | int4           | 报警类型（1.开启SD&SIM卡槽 2.私自拔插电源线、天线 3.私自安装设备电源开关 4.遮挡、屏蔽天线） | optional     |                            |
| audit_status              | text           | 审核状态                                                     | optional     |                            |
| auditor                   | text           | 审核人                                                       | optional     |                            |
| pic_id                    | text           | 图片表ID                                                     | optional     |                            |
| alarm_position            | text           | 报警位置                                                     | optional     |                            |
| is_completed              | boolean        | 是否完成                                                     | optional     |                            |
| is_supervise              | boolean        | 是否监管                                                     | optional     |                            |
| supervision_time          | timestamptz(6) | 监管时间                                                     | optional     |                            |
| is_deleted                | boolean        | 是否删除                                                     | optional     |                            |
| created_at                | timestamptz(6)    | 创建时间                                                     | required     |                            |
| created_by                | text           | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6)    | 修改时间                                                     | optional     |                            |
| updated_by                | text           | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6)    | 删除时间                                                     | optional     |                            |
| deleted_by                | text           | 删除人                                                       | optional     | **system_user表**的user_id |

### destroy_terminal_pic 疑似破坏终端证据图片  

**根据以下数据表整合而成**

工程车:

SmartCarAlarmPHZD_PIC 疑似破坏终端证据图片

| Name                    | Type        | Description                                     | **Required** | default                    |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                       |
| destroy_terminal_pic_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| serial_num              | text        | 序号                                            | optional     |                            |
| img_src                 | text        | 图片地址                                        | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_egn_check_status 新型渣土车测试情况

**根据以下数据表整合而成**

工程车:

INFO_SMARTCAR_CHECK 新型渣土车测试情况

| Name                       | Type        | Description                                     | **Required** | default                    |
| -------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_egn_check_status_id    | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                 | text        | **vehicle_info** 车辆信息表 的vehicle_id        | required     |                            |
| airtight_state             | integer     | 密闭状态                                        | optional     |                            |
| load_state                 | integer     | 载重状态                                        | optional     |                            |
| lifting_status             | integer     | 举升状态                                        | optional     |                            |
| video                      | integer     | 视频                                            | optional     |                            |
| fingerprint                | integer     | 指纹                                            | optional     |                            |
| car_lock_function          | integer     | 锁车功能                                        | optional     |                            |
| speed_limit_function       | integer     | 限速功能                                        | optional     |                            |
| ministry_standard_function | integer     | 部标功能                                        | optional     |                            |
| weight_limit_function      | integer     | 限举功能                                        | optional     |                            |
| left_turn_light            | integer     | 左转灯                                          | optional     |                            |
| right_turn_right           | integer     | 右转灯                                          | optional     |                            |
| high_beam                  | integer     | 远光灯                                          | optional     |                            |
| low_beam                   | integer     | 近光灯                                          | optional     |                            |
| brake                      | integer     | 刹车                                            | optional     |                            |
| speed                      | integer     | 车速                                            | optional     |                            |
| is_detect_illegal_spoil    | boolean     | 是否检测非法弃土                                | optional     |                            |
| is_detect_illegal_start    | boolean     | 是否检测违规启动                                | optional     |                            |
| is_passed                  | boolean     | 是否通过                                        | optional     |                            |
| is_deleted                 | boolean     | 是否删除                                        | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text        | 删除人                                          | optional     | **system_user表**的user_id |

### blacklist_operation_record 黑名单操作记录表 

**根据以下数据表整合而成**

工程车:

BlackListLog 黑名单操作日志表

| Name                          | Type        | Description                                     | **Required** | default                    |
| ----------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| blacklist_operation_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| obj_type                      | int4        | 黑名单主体类别  1.运输主体  2.车辆  3.驾驶员    | optional     |                            |
| obj_id                        | text        | 主体id(运输主体,车辆，驾驶员)                   | optional     |                            |
| oper_type                     | int4        | 操作类别  1.加入黑名单  0.移出黑名单            | optional     |                            |
| remarks                       | text        | 备注                                            | optional     |                            |
| result                        | text        | 结果                                            | optional     |                            |
| blacklist_file_id             | text        | 黑名单文件表ID                                  | optional     |                            |
| driver_blacklist_apply        | text        | 驾驶员黑名单申请主表id                          | optional     |                            |
| step                          | int4        | 进度                                            | optional     |                            |
| blacklist_type                | int4        | 黑名单类型                                      | optional     |                            |
| driver_license_num            | text        | 驾驶证号                                        | optional     |                            |
| company_name                  | text        | 单位名称                                        | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                    | text        | 删除人                                          | optional     | **system_user表**的user_id |

### blacklist_file 黑名单文件表

**根据以下数据表整合而成**

工程车:

BlackListFiles 黑名单文件表

| Name              | Type        | Description                                     | **Required** | default                    |
| ----------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                | bigint      | 按指定方法生成                                  | required     | 主键                       |
| blacklist_file_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| serial_num        | text        | 序号                                            | optional     |                            |
| file_path         | text        | 文件路径                                        | optional     |                            |
| created_at        | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by        | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at        | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by        | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at        | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by        | text        | 删除人                                          | optional     | **system_user表**的user_id |

### area_police_force 各区域警力表

**根据以下数据表整合而成**

工程车:

Police 各区域警力表

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| area_police_force_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| area_id              | text        | 区域ID                                          | optional     |                            |
| police_num           | int4        | 警力数                                          | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_electric_fence_distribute 电子围栏报警分发表

**根据以下数据表整合而成**

工程车:

INFO_ELECTRIC_CRAWL_ALARM_DEAL 电子围栏报警分发表

| Name                               | Type        | Description                                     | **Required** | default                    |
| ---------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                 | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_electric_fence_distribute_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| alarm_id                           | text        | 报警ID                                          | optional     |                            |
| user_role                          | text        | 用户角色                                        | optional     |                            |
| is_supervise                       | boolean     | 是否监管                                        | optional     |                            |
| created_at                         | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                         | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                         | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                         | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                         | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                         | text        | 删除人                                          | optional     | **system_user表**的user_id |

### dailly_veh_overspeed_alarm 每日通报车辆超速报警

**根据以下数据表整合而成**

工程车:

DaySpeedAlarmVehDetail 每日通报车辆超速报警

| Name                 | Type        | Description                              | **Required** | default                    |
| -------------------- | ----------- | ---------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                           | required     | 主键                       |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                            |
| license_plate_number | text        | 车牌号                                   | optional     |                            |
| ent_id               | text        | 企业id                                   | optional     |                            |
| ent_name             | text        | 企业名称                                 | optional     |                            |
| area_id              | text        | 区域ID                                   | optional     |                            |
| work_num             | text        | 工号                                     | optional     |                            |
| statistics_date      | timestamptz(6) | 统计日期                                 | optional     |                            |
| alarm_times          | int4        | 报警次数                                 | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by           | text        | 创建人                                   | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by           | text        | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by           | text        | 删除人                                   | optional     | **system_user表**的user_id |

### dailly_veh_offline_alarm 每日通报不在线车辆明细

**根据以下数据表整合而成**

工程车:

DayOfflineVehDetail 每日通报不在线车辆明细

| Name                 | Type        | Description                              | **Required** | default                    |
| -------------------- | ----------- | ---------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                           | required     | 主键                       |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                            |
| license_plate_number | text        | 车牌号                                   | optional     |                            |
| ent_id               | text        | 企业id                                   | optional     |                            |
| ent_name             | text        | 企业名称                                 | optional     |                            |
| area_id              | text        | 区域ID                                   | optional     |                            |
| work_num             | text        | 工号                                     | optional     |                            |
| statistics_date      | timestamptz(6) | 统计日期                                 | optional     |                            |
| offline_time         | int4        | 离线时长（秒）                           | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by           | text        | 创建人                                   | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by           | text        | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by           | text        | 删除人                                   | optional     | **system_user表**的user_id |

### dailly_veh_fatigue_alarm 每日通报车辆疲劳驾驶报警

**根据以下数据表整合而成**

工程车:

DayFatigueVehDetail 每日通报车辆疲劳驾驶报警

| Name                 | Type        | Description                              | **Required** | default                    |
| -------------------- | ----------- | ---------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                           | required     | 主键                       |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                            |
| license_plate_number | text        | 车牌号                                   | optional     |                            |
| ent_id               | text        | 企业id                                   | optional     |                            |
| ent_name             | text        | 企业名称                                 | optional     |                            |
| area_id              | text        | 区域ID                                   | optional     |                            |
| work_num             | text        | 工号                                     | optional     |                            |
| statistics_date      | timestamptz(6) | 统计日期                                 | optional     |                            |
| alarm_times          | int4        | 报警次数                                 | optional     |                            |
| alarm_duration       | int4        | 报警时长                                 | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by           | text        | 创建人                                   | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by           | text        | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by           | text        | 删除人                                   | optional     | **system_user表**的user_id |

### dailly_veh__trajectory_integrity_alarm 每日通报车辆轨迹完整率明细

**根据以下数据表整合而成**

工程车:

DayTrajectoryVehDetail 每日通报车辆轨迹完整率明细

| Name                 | Type        | Description                              | **Required** | default                    |
| -------------------- | ----------- | ---------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                           | required     | 主键                       |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                            |
| license_plate_number | text        | 车牌号                                   | optional     |                            |
| ent_id               | text        | 企业id                                   | optional     |                            |
| ent_name             | text        | 企业名称                                 | optional     |                            |
| area_id              | text        | 区域ID                                   | optional     |                            |
| work_num             | text        | 工号                                     | optional     |                            |
| statistics_date      | timestamptz(6) | 统计日期                                 | optional     |                            |
| all_mileage          | numeric     | 全部里程                                 | optional     |                            |
| continuous_mileage   | numeric     | 连续里程                                 | optional     |                            |
| completeness_rate    | numeric     | 完整率                                   | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by           | text        | 创建人                                   | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by           | text        | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by           | text        | 删除人                                   | optional     | **system_user表**的user_id |

### dailly_veh_trans_alarm 每日通报车辆工程运输过程报警

**根据以下数据表整合而成**

工程车:

DayFileAlarmVehDetail 每日通报车辆工程运输过程报警

| Name                 | Type        | Description                              | **Required** | default                    |
| -------------------- | ----------- | ---------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                           | required     | 主键                       |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                            |
| license_plate_number | text        | 车牌号                                   | optional     |                            |
| ent_id               | text        | 企业id                                   | optional     |                            |
| ent_name             | text        | 企业名称                                 | optional     |                            |
| area_id              | text        | 区域ID                                   | optional     |                            |
| work_num             | text        | 工号                                     | optional     |                            |
| statistics_date      | timestamptz(6) | 统计日期                                 | optional     |                            |
| offset_route_times   | int4        | 偏移路线报警次数                         | optional     |                            |
| limit_speed_times    | int4        | 限速报警次数                             | optional     |                            |
| limit_time_times     | int4        | 限时报警次数                             | optional     |                            |
| total_times          | int4        | 合计报警次数                             | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by           | text        | 创建人                                   | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by           | text        | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by           | text        | 删除人                                   | optional     | **system_user表**的user_id |

### dailly_veh_unreport_alarm 每日通报车辆未报备工程报警

**根据以下数据表整合而成**

工程车:

DayUnReportVehDetail 每日通报车辆未报备工程报警

| Name                 | Type        | Description                              | **Required** | default                    |
| -------------------- | ----------- | ---------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                           | required     | 主键                       |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                            |
| license_plate_number | text        | 车牌号                                   | optional     |                            |
| ent_id               | text        | 企业id                                   | optional     |                            |
| ent_name             | text        | 企业名称                                 | optional     |                            |
| area_id              | text        | 区域ID                                   | optional     |                            |
| work_num             | text        | 工号                                     | optional     |                            |
| statistics_date      | timestamptz(6) | 统计日期                                 | optional     |                            |
| alarm_times          | int4        | 报警次数                                 | optional     |                            |
| trans_duration       | int4        | 运输时长                                 | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by           | text        | 创建人                                   | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by           | text        | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by           | text        | 删除人                                   | optional     | **system_user表**的user_id |

### dailly_ent_overspeed_alarm 每日通报企业超速报警

**根据以下数据表整合而成**

工程车:

DaySpeedAlarmCompanyDetail 每日通报企业超速报警

| Name            | Type        | Description    | **Required** | default                    |
| --------------- | ----------- | -------------- | ------------ | -------------------------- |
| id              | bigint      | 按指定方法生成 | required     | 主键                       |
| ent_id          | text        | 企业id         | optional     |                            |
| ent_name        | text        | 企业名称       | optional     |                            |
| area_id         | text        | 区域ID         | optional     |                            |
| statistics_date | timestamptz(6) | 统计日期       | optional     |                            |
| veh_num         | int4        | 车辆数         | optional     |                            |
| alarm_times     | int4        | 报警次数       | optional     |                            |
| created_at      | timestamptz(6) | 创建时间       | required     |                            |
| created_by      | text        | 创建人         | required     | **system_user表**的user_id |
| updated_at      | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by      | text        | 修改人         | optional     | **system_user表**的user_id |
| deleted_at      | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by      | text        | 删除人         | optional     | **system_user表**的user_id |

### dailly_ent_offline_alarm 每日通报企业不在线车辆数明细

**根据以下数据表整合而成**

工程车:

DayOfflineCompanyDetail 每日通报企业不在线车辆数明细

| Name            | Type        | Description    | **Required** | default                    |
| --------------- | ----------- | -------------- | ------------ | -------------------------- |
| id              | bigint      | 按指定方法生成 | required     | 主键                       |
| ent_id          | text        | 企业id         | optional     |                            |
| ent_name        | text        | 企业名称       | optional     |                            |
| area_id         | text        | 区域ID         | optional     |                            |
| statistics_date | timestamptz(6) | 统计日期       | optional     |                            |
| veh_num         | int4        | 车辆数         | optional     |                            |
| veh_offline_num | int4        | 离线车辆数     | optional     |                            |
| created_at      | timestamptz(6) | 创建时间       | required     |                            |
| created_by      | text        | 创建人         | required     | **system_user表**的user_id |
| updated_at      | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by      | text        | 修改人         | optional     | **system_user表**的user_id |
| deleted_at      | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by      | text        | 删除人         | optional     | **system_user表**的user_id |

### dailly_ent_fatigue_alarm 每日通报企业疲劳驾驶报警

**根据以下数据表整合而成**

工程车:

DayFatigueCompanyDetail 每日通报企业疲劳驾驶报警

| Name            | Type        | Description    | **Required** | default                    |
| --------------- | ----------- | -------------- | ------------ | -------------------------- |
| id              | bigint      | 按指定方法生成 | required     | 主键                       |
| ent_id          | text        | 企业id         | optional     |                            |
| ent_name        | text        | 企业名称       | optional     |                            |
| area_id         | text        | 区域ID         | optional     |                            |
| statistics_date | timestamptz(6) | 统计日期       | optional     |                            |
| veh_num         | int4        | 车辆数         | optional     |                            |
| alarm_times     | int4        | 报警次数       | optional     |                            |
| created_at      | timestamptz(6) | 创建时间       | required     |                            |
| created_by      | text        | 创建人         | required     | **system_user表**的user_id |
| updated_at      | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by      | text        | 修改人         | optional     | **system_user表**的user_id |
| deleted_at      | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by      | text        | 删除人         | optional     | **system_user表**的user_id |

### dailly_ent__trajectory_integrity_alarm 每日通报企业轨迹完整率

**根据以下数据表整合而成**

工程车:

DayTrajectoryCompanyDetail 每日通报企业轨迹完整率

| Name               | Type        | Description    | **Required** | default                    |
| ------------------ | ----------- | -------------- | ------------ | -------------------------- |
| id                 | bigint      | 按指定方法生成 | required     | 主键                       |
| ent_id             | text        | 企业id         | optional     |                            |
| ent_name           | text        | 企业名称       | optional     |                            |
| veh_num            | int4        | 车辆数         | optional     |                            |
| area_id            | text        | 区域ID         | optional     |                            |
| statistics_date    | timestamptz(6) | 统计日期       | optional     |                            |
| all_mileage        | numeric     | 全部里程       | optional     |                            |
| continuous_mileage | numeric     | 连续里程       | optional     |                            |
| created_at         | timestamptz(6) | 创建时间       | required     |                            |
| created_by         | text        | 创建人         | required     | **system_user表**的user_id |
| updated_at         | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by         | text        | 修改人         | optional     | **system_user表**的user_id |
| deleted_at         | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by         | text        | 删除人         | optional     | **system_user表**的user_id |

### dailly_ent_trans_alarm 每日通报企业工程运输过程报警

**根据以下数据表整合而成**

工程车:

DayFileAlarmCompanyDetail 每日通报企业工程运输过程报警

| Name               | Type        | Description      | **Required** | default                    |
| ------------------ | ----------- | ---------------- | ------------ | -------------------------- |
| id                 | bigint      | 按指定方法生成   | required     | 主键                       |
| ent_id             | text        | 企业id           | optional     |                            |
| ent_name           | text        | 企业名称         | optional     |                            |
| veh_num            | int4        | 车辆数           | optional     |                            |
| area_id            | text        | 区域ID           | optional     |                            |
| statistics_date    | timestamptz(6) | 统计日期         | optional     |                            |
| offset_route_times | int4        | 偏移路线报警次数 | optional     |                            |
| limit_speed_times  | int4        | 限速报警次数     | optional     |                            |
| limit_time_times   | int4        | 限时报警次数     | optional     |                            |
| total_times        | int4        | 合计报警次数     | optional     |                            |
| created_at         | timestamptz(6) | 创建时间         | required     |                            |
| created_by         | text        | 创建人           | required     | **system_user表**的user_id |
| updated_at         | timestamptz(6) | 修改时间         | optional     |                            |
| updated_by         | text        | 修改人           | optional     | **system_user表**的user_id |
| deleted_at         | timestamptz(6) | 删除时间         | optional     |                            |
| deleted_by         | text        | 删除人           | optional     | **system_user表**的user_id |

### dailly_ent_unreport_alarm 每日通报企业未报备工程报警

**根据以下数据表整合而成**

工程车:

DayUnReportCompanyDetail 每日通报企业未报备工程报警

| Name            | Type        | Description    | **Required** | default                    |
| --------------- | ----------- | -------------- | ------------ | -------------------------- |
| id              | bigint      | 按指定方法生成 | required     | 主键                       |
| ent_id          | text        | 企业id         | optional     |                            |
| ent_name        | text        | 企业名称       | optional     |                            |
| veh_num         | int4        | 车辆数         | optional     |                            |
| area_id         | text        | 区域ID         | optional     |                            |
| work_num        | text        | 工号           | optional     |                            |
| statistics_date | timestamptz(6) | 统计日期       | optional     |                            |
| trans_duration  | int4        | 运输时长       | optional     |                            |
| created_at      | timestamptz(6) | 创建时间       | required     |                            |
| created_by      | text        | 创建人         | required     | **system_user表**的user_id |
| updated_at      | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by      | text        | 修改人         | optional     | **system_user表**的user_id |
| deleted_at      | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by      | text        | 删除人         | optional     | **system_user表**的user_id |

### dailly_site_not_include_alarm 每日通报工地车辆未纳入管控平台报警

**根据以下数据表整合而成**

工程车:

DayNotPlatformVehDetail 每日通报工地车辆未纳入管控平台报警

| Name              | Type        | Description                              | **Required** | default                    |
| ----------------- | ----------- | ---------------------------------------- | ------------ | -------------------------- |
| id                | bigint      | 按指定方法生成                           | required     | 主键                       |
| vehicle_id        | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                            |
| construction_type | text        | 工地类型                                 | optional     |                            |
| area_id           | text        | 区域ID                                   | optional     |                            |
| statistics_date   | timestamptz(6) | 统计日期                                 | optional     |                            |
| alarm_times       | int4        | 报警次数                                 | optional     |                            |
| created_at        | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by        | text        | 创建人                                   | required     | **system_user表**的user_id |
| updated_at        | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by        | text        | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at        | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by        | text        | 删除人                                   | optional     | **system_user表**的user_id |

### dailly_site_unreport_alarm 每日通报工地车辆未报备工程报警

**根据以下数据表整合而成**

工程车:

DayNotRepFileVehDetail 每日通报工地车辆未报备工程报警

| Name                 | Type        | Description                              | **Required** | default                    |
| -------------------- | ----------- | ---------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                           | required     | 主键                       |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id | required     |                            |
| license_plate_number | text        | 车牌号                                   | optional     |                            |
| ent_id               | text        | 企业id                                   | optional     |                            |
| ent_name             | text        | 企业名称                                 | optional     |                            |
| area_id              | text        | 区域ID                                   | optional     |                            |
| work_num             | text        | 工号                                     | optional     |                            |
| statistics_date      | timestamptz(6) | 统计日期                                 | optional     |                            |
| alarm_times          | int4        | 报警次数                                 | optional     |                            |
| construction_id      | text        | 工地ID                                   | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                 | required     |                            |
| created_by           | text        | 创建人                                   | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                 | optional     |                            |
| updated_by           | text        | 修改人                                   | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                 | optional     |                            |
| deleted_by           | text        | 删除人                                   | optional     | **system_user表**的user_id |

### dailly_site_alarm 每日通报工地报警

**根据以下数据表整合而成**

工程车:

DayConstructionAlarmDetail 每日通报工地报警

| Name              | Type        | Description    | **Required** | default                    |
| ----------------- | ----------- | -------------- | ------------ | -------------------------- |
| id                | bigint      | 按指定方法生成 | required     | 主键                       |
| construction_id   | text        | 工地ID         | optional     |                            |
| construction_type | text        | 工地类型       | optional     |                            |
| construction_name | text        | 项目名称       | optional     |                            |
| area_id           | text        | 区域ID         | optional     |                            |
| alarm_type        | text        | 报警类型       | optional     |                            |
| statistics_date   | timestamptz(6) | 统计日期       | optional     |                            |
| alarm_times       | int4        | 报警次数       | optional     |                            |
| created_at        | timestamptz(6) | 创建时间       | required     |                            |
| created_by        | text        | 创建人         | required     | **system_user表**的user_id |
| updated_at        | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by        | text        | 修改人         | optional     | **system_user表**的user_id |
| deleted_at        | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by        | text        | 删除人         | optional     | **system_user表**的user_id |

### daily_veh_enter_out_site 日通报车辆进出工地次数

**根据以下数据表整合而成**

工程车:

DayConstructionVehPassDetail  日通报车辆进出工地次数

| Name                 | Type        | Description    | **Required** | default                    |
| -------------------- | ----------- | -------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成 | required     | 主键                       |
| license_plate_number | text        | 车牌号         | optional     |                            |
| construction_id      | text        | 工地ID         | optional     |                            |
| construction_name    | text        | 项目名称       | optional     |                            |
| area_id              | text        | 区域ID         | optional     |                            |
| statistics_date      | timestamptz(6) | 统计日期       | optional     |                            |
| alarm_times          | int4        | 过车次数       | optional     |                            |
| created_at           | timestamptz(6) | 创建时间       | required     |                            |
| created_by           | text        | 创建人         | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by           | text        | 修改人         | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by           | text        | 删除人         | optional     | **system_user表**的user_id |

### constrution_user_related 用户工地关联表

**根据以下数据表整合而成**

工程车:

ConstructionSubscribe 用户工地关联表

| Name                        | Type        | Description                                     | **Required** | default                    |
| --------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                          | bigint      | 按指定方法生成                                  | required     | 主键                       |
| constrution_user_related_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| user_id                     | text        | 用户ID                                          | optional     |                            |
| construction_id             | text        | 工地ID                                          | optional     |                            |
| is_deleted                  | boolean     | 是否删除                                        | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                  | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                  | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                  | text        | 删除人                                          | optional     | **system_user表**的user_id |

## 营运车相关数据表

### veh_online_status 车辆上线情况

**根据以下数据表整合而成**

营运车:

REPORT_LOCATION_ONLINE 车辆上线情况

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_online_status_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number | text        | 车牌号                                          | optional     |                            |
| start_time           | timestamptz(6) | 开始时间                                        | optional     |                            |
| deadline             | timestamptz(6) | 截止时间                                        | optional     |                            |
| acc_status           | text        | ACC状态                                         | optional     |                            |
| interval             | numeric     | 间隔分钟                                        | optional     |                            |
| is_deleted           | boolean     | 删除标识                                        | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_payment_standard 车辆缴费标准

**根据以下数据表整合而成**

营运车:

FEE_VEHBASEFEE 车辆缴费标准

| Name                           | Type           | Description            | **Required** | default                    |
| ------------------------------ | -------------- | ---------------------- | ------------ | -------------------------- |
| id                             | bigint         | 按指定方法生成         | required     | 主键                       |
| veh_payment_standard_id        | text           | 缴费单据id             | required     | 联合主键                   |
| veh_id                         | text           | 车辆id                 | optional     |                            |
| sim_month_fee                  | numeric        | SIM卡月费              | optional     |                            |
| service_month_fee              | numeric        | 服务月费               | optional     |                            |
| supervision_month_fee          | numeric        | 监管月费               | optional     |                            |
| is_manually                    | boolean        | 是否手工设置           | optional     |                            |
| registration_time              | timestamptz(6) | 注册时间               | optional     |                            |
| sim_payment_start_date         | timestamptz(6) | SIM费下次缴费起始日期  | optional     |                            |
| service_payment_start_date     | timestamptz(6) | 服务费下次缴费起始日期 | optional     |                            |
| supervision_payment_start_date | timestamptz(6) | 监管费下次缴费起始日期 | optional     |                            |
| sim_payment_expire             | timestamptz(6) | sim卡费到期时间        | optional     |                            |
| service_payment_expire         | timestamptz(6) | 服务费到期时间         | optional     |                            |
| supervision_payment_expire     | timestamptz(6) | 监管费到期时间         | optional     |                            |
| sim_fee_required               | numeric        | 需缴SIM费              | optional     |                            |
| service_fee_required           | numeric        | 需缴服务费             | optional     |                            |
| supervision_fee_required       | numeric        | 需缴监管费             | optional     |                            |
| total_payment_amount           | numeric        | 缴费总金额             | optional     |                            |
| created_at                     | timestamptz(6) | 创建时间               | required     |                            |
| created_by                     | text           | 创建人                 | required     | **system_user表**的user_id |
| updated_at                     | timestamptz(6) | 修改时间               | optional     |                            |
| updated_by                     | text           | 修改人                 | optional     | **system_user表**的user_id |
| deleted_at                     | timestamptz(6) | 删除时间               | optional     |                            |
| deleted_by                     | text           | 删除人                 | optional     | **system_user表**的user_id |

### check_post 查岗应答

**根据以下数据表整合而成**

营运车:

REPORT_PLATFORM_MSG 查岗应答

| Name                        | Type        | Description                                     | **Required** | default                    |
| --------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                          | bigint      | 按指定方法生成                                  | required     | 主键                       |
| check_post_id               | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| ent_id                      | text        | 企业的用户ID                                    | optional     |                            |
| msg_id                      | text        | 信息ID                                          | optional     |                            |
| sub_business_id             | text        | 子业务类型标识                                  | optional     |                            |
| subsequent_date_length      | text        | 后续数据长度                                    | optional     |                            |
| request_date_length         | text        | 请求的数据长度                                  | optional     |                            |
| request_check_post_content  | text        | 请求查岗内容                                    | optional     |                            |
| response_data_length        | text        | 应答的数据长度                                  | optional     |                            |
| response_check_post_content | text        | 应答查岗内容                                    | optional     |                            |
| is_response                 | text        | 是否应答                                        | optional     |                            |
| auto_send                   | int4        | 0,上级平台请求查岗;1，平台自我请求查岗          | optional     |                            |
| check_post_time             | timestamptz(6) | 查岗的时间                                      | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                  | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                  | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                  | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_call_notice_record 报警电话通知记录  

**根据以下数据表整合而成**

营运车:

PHONE_NOTICE_LOG 报警电话通知日志

| Name                        | Type        | Description                                     | **Required** | default                    |
| --------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                          | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_call_notice_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                      | text        | 车辆id                                          | optional     |                            |
| alarm_id                    | text        | 报警ID                                          | optional     |                            |
| entry_person                | text        | 录入人                                          | optional     |                            |
| entry_time                  | timestamptz(6) | 录入时间                                        | optional     |                            |
| notification_contact        | text        | 通知联系人                                      | optional     |                            |
| notification_phone          | text        | 通知电话                                        | optional     |                            |
| disposal_measure            | text        | 处置措施                                        | optional     |                            |
| disposal_result             | text        | 处置结果                                        | optional     |                            |
| is_send_sms                 | boolean     | 是否发送短信                                    | optional     |                            |
| second_contact_phone        | text        | 第二联系人电话                                  | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                  | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                  | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                  | text        | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_sms_notice_record 报警短信通知记录

**根据以下数据表整合而成**

营运车:

PHONE_NOTICE_SMS_LOG 报警电话通知日志短信通知记录

| Name                       | Type        | Description                                     | **Required** | default                    |
| -------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_sms_notice_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| alarm_id                   | text        | 报警ID                                          | optional     |                            |
| veh_id                     | text        | 车辆id                                          | optional     |                            |
| notification_id            | text        | 通知表ID                                        | optional     |                            |
| send_time                  | timestamptz(6) | 发送时间                                        | optional     |                            |
| sms_content                | text        | 短信内容                                        | optional     |                            |
| contact_num                | text        | 联系号码                                        | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text        | 删除人                                          | optional     | **system_user表**的user_id |

### daily_gps_can_speed_record 每日车辆GPS速度和CAN速度排查记录

**根据以下数据表整合而成**

营运车:

INFO_VEHPOSITION_QUA 每日车辆GPS速度和CAN速度排查记录

| Name                          | Type        | Description                                                  | **Required** | default                    |
| ----------------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                               | required     | 主键                       |
| daily_gps_can_speed_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用              | required     | 联合主键                   |
| veh_id                        | text        | 车辆id                                                       | optional     |                            |
| total_pos_point               | numeric     | 总定位点数(仅已定位的点)                                     | optional     |                            |
| pos_point_qualified_num       | numeric     | 定位点合格数                                                 | optional     |                            |
| pos_point_unqualified_num     | numeric     | 定位点不合格数                                               | optional     |                            |
| pos_point_qualified_rate      | numeric     | 定位点合格率                                                 | optional     |                            |
| pos_point_unqualified_rate    | numeric     | 定位点不合格率                                               | optional     |                            |
| non_pos_point                 | numeric     | 不定位点数                                                   | optional     |                            |
| online_total_time             | numeric     | 在线总时间（不包含不定位点）                                 | optional     |                            |
| max_continuous_online_time    | numeric     | 最大连续在线时间                                             | optional     |                            |
| online_total_time_all         | numeric     | 在线总时间（包含不定位点）                                   | optional     |                            |
| speed_normal_count            | numeric     | 行车记录仪速度与GPS速度差值在3km/h以内（不包含不定位点）     | optional     |                            |
| veh_data_type                 | numeric     | 车辆数据类型                                                 | optional     |                            |
| not_pos_point_unqualified_num | numeric     | 定位点合格数（不包含海拔异常）                               | optional     |                            |
| not_pos_point_qualified_rate  | numeric     | 定位点合格率（不包含海拔异常）                               | optional     |                            |
| driving_pos_point             | numeric     | 行驶的定位点                                                 | optional     |                            |
| pulse_speed_less_ten          | numeric     | 脉冲速度差10以下                                             | optional     |                            |
| acc_close_point               | numeric     | ACC关闭的点(已定位)                                          | optional     |                            |
| illegal_acc_close_point       | numeric     | ACC关闭的点，但不合法（已定位点，GPS速度或行驶记录仪速度大于5km/h） | optional     |                            |
| can_speed_differ_than_ten     | numeric     | CAN速度差10以上                                              | optional     |                            |
| can_speed_point_num           | numeric     | CAN速度的点数                                                | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by                    | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by                    | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by                    | text        | 删除人                                                       | optional     | **system_user表**的user_id |

### oper_msg_reminder 消息提醒信息表

**根据以下数据表整合而成**

营运车:

ZX_NOTICE 消息提醒信息表

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| oper_msg_reminder_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| title                | text        | 标题                                            | optional     |                            |
| content              | text        | 内容                                            | optional     |                            |
| msg_type             | text        | 消息类型                                        | optional     |                            |
| is_read              | boolean     | 是否已读                                        | optional     |                            |
| sender_id            | text        | 发送人ID                                        | optional     |                            |
| receiver_id          | text        | 接收信息人id                                    | optional     |                            |
| send_time            | timestamptz(6) | 发送时间                                        | optional     |                            |
| read_time            | timestamptz(6) | 读取时间                                        | optional     |                            |
| send_url             | text        | 发送连接URL                                     | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_offline_registration 不在线车辆登记记录

**根据以下数据表整合而成**

营运车:

MON_OFFLINE_ALARM_LOG 不在线车辆登记记录

| Name                        | Type           | Description                                     | **Required** | default                    |
| --------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                          | int8           | 按指定方法生成                                  | required     | 主键                       |
| veh_offline_registration_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| offline_start_time          | timestamptz(6) | 离线开始时间                                    | optional     |                            |
| offline_end_time            | timestamptz(6) | 离线结束时间                                    | optional     |                            |
| license_plate_number        | text           | 车牌号                                          | optional     |                            |
| phone_reminder_content      | text           | 电话提醒内容                                    | optional     |                            |
| phone_reminder_time         | timestamptz(6) | 电话提醒时间                                    | optional     |                            |
| offline_reason              | text           | 离线原因                                        | optional     |                            |
| is_need_maintain            | boolean        | 是否需要维护                                    | optional     |                            |
| registration_user           | text           | 登记用户                                        | optional     | **system_user表**的user_id |
| registration_time           | timestamptz(6) | 登记时间                                        | optional     |                            |
| alarm_type                  | text           | 报警ID                                          | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                  | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                  | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                  | text           | 删除人                                          | optional     | **system_user表**的user_id |

### daily_veh_breakdown_record 日常车辆故障登记

**根据以下数据表整合而成**

营运车:

MON_TERMINAL_DRIVING 日常车辆故障登记

| Name                          | Type        | Description                                     | **Required** | default                    |
| ----------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| daily_veh_breakdown_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| terminal_id                   | text        | 终端ID                                          | optional     |                            |
| acc_status                    | text        | ACC状态                                         | optional     |                            |
| brake_signal                  | text        | 制动信号(刹车)                                  | optional     |                            |
| left_turn_signal              | text        | 左转向灯信号                                    | optional     |                            |
| right_turn_signal             | text        | 右转向灯信号                                    | optional     |                            |
| low_beam_signal               | text        | 近光灯信号                                      | optional     |                            |
| high_beam_signal              | text        | 远光灯信号(大灯)                                | optional     |                            |
| horm_signal                   | text        | 喇叭信号                                        | optional     |                            |
| pos_status                    | text        | 定位状态                                        | optional     |                            |
| front_door                    | text        | 前门                                            | optional     |                            |
| gps_signal                    | text        | GPS信号                                         | optional     |                            |
| beidou_signal                 | text        | 北斗信号                                        | optional     |                            |
| camera                        | text        | 摄像头                                          | optional     |                            |
| emergency_alarm_button        | text        | 紧急报警按钮                                    | optional     |                            |
| speed                         | text        | 车速                                            | optional     |                            |
| alarm_main_power_failure      | text        | 主电源掉电报警                                  | optional     |                            |
| remarks                       | text        | 备注                                            | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                    | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_addiliate_department 车辆隶属部门表

营运车:

INFO_VEHICLE_AFFILIATE 车辆隶属部门表

| Name                        | Type        | Description                                     | **Required** | default                    |
| --------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                          | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_addiliate_department_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| dept_id                     | text        | 所在部门                                        | optional     |                            |
| veh_id                      | text        | 车辆id                                          | optional     |                            |
| user_id                     | text        | 用户ID                                          | optional     |                            |
| is_deleted                  | boolean     | 是否删除                                        | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                  | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                  | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                  | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_dynamic_supervision 道路运输车辆动态监控记录表 

营运车:

INFO_VEHICLE_DYNAMIC 道路运输车辆动态监控记录表

| Name                        | Type           | Description                                                  | **Required** | default                         |
| --------------------------- | -------------- | ------------------------------------------------------------ | ------------ | ------------------------------- |
| id                          | int8           | 按指定方法生成                                               | required     | 主键                            |
| veh_dynamic_supervision_id  | text           | 动态监管抽查明细表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                        |
| vehicle_id                  | text           | 车辆id                                                       | optional     | **vehicle_info 表**的vehicle_id |
| monitoring_location         | text           | 监控平台显示位置                                             | optional     |                                 |
| is_tachograph_record_normal | boolean        | 行车记录仪数据（是否异常）                                   | optional     |                                 |
| tachograph_data_exception   | text           | 行车记录仪异常数据项                                         | optional     |                                 |
| is_speeding                 | boolean        | 是否超速（是/否）                                            | optional     |                                 |
| taxi_state                  | int4           | 出租空/重车状态（空/重）                                     | optional     | **出租车空/重车状态**字典       |
| is_fatigue_driving          | boolean        | 客运疲劳驾驶（是/否）                                        | optional     |                                 |
| is_morning_outage           | boolean        | 客运凌晨2-5时停运（是/否）                                   | optional     |                                 |
| curve                       | int4           | 曲线情况（曲线完整/回传异常/零速度）                         | optional     | **曲线情况**字典                |
| trail                       | int4           | 轨迹情况（正常/漂移/其他）                                   | optional     | **GPS轨迹情况**字典             |
| lens_position               | int4           | 镜头位置（正/偏）                                            | optional     | **镜头位置**字典                |
| other_infraction            | text           | 其他违规                                                     | optional     |                                 |
| lens_on                     | text           | 摄像头损坏号                                                 | optional     |                                 |
| coordinate                  | point          | 空间数据类型point表示经纬度                                  | optional     |                                 |
| tachograph_speed            | text           | 行车记录仪速度                                               | optional     |                                 |
| gps_speed                   | text           | 卫星定位速度                                                 | optional     |                                 |
| monitoring_time             | timestamptz(6) | 监控平台时间                                                 | optional     |                                 |
| is_locate                   | boolean        | 是否定位                                                     | optional     |                                 |
| created_at                  | timestamptz(6) | 创建时间                                                     | required     |                                 |
| created_by                  | text           | 创建人                                                       | required     | **system_user表**的user_id      |
| updated_at                  | timestamptz(6) | 修改时间                                                     | optional     |                                 |
| updated_by                  | text           | 修改人                                                       | optional     | **system_user表**的user_id      |
| deleted_at                  | timestamptz(6) | 删除时间                                                     | optional     |                                 |
| deleted_by                  | text           | 删除人                                                       | optional     | **system_user表**的user_id      |

### his_veh_dynamic_supervision 道路运输车辆历史监控记录表

营运车:

INFO_VEHICLE_HISTORY 道路运输车辆历史监控记录表

| Name                           | Type           | Description                                                  | **Required** | default                         |
| ------------------------------ | -------------- | ------------------------------------------------------------ | ------------ | ------------------------------- |
| id                             | int8           | 按指定方法生成                                               | required     | 主键                            |
| his_veh_dynamic_supervision_id | text           | 动态监管抽查明细表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                        |
| vehicle_id                     | text           | 车辆id                                                       | optional     | **vehicle_info 表**的vehicle_id |
| monitoring_location            | text           | 监控平台显示位置                                             | optional     |                                 |
| is_tachograph_record_normal    | boolean        | 行车记录仪数据（是否异常）                                   | optional     |                                 |
| tachograph_data_exception      | text           | 行车记录仪异常数据项                                         | optional     |                                 |
| is_speeding                    | boolean        | 是否超速（是/否）                                            | optional     |                                 |
| taxi_state                     | int4           | 出租空/重车状态（空/重）                                     | optional     | **出租车空/重车状态**字典       |
| is_fatigue_driving             | boolean        | 客运疲劳驾驶（是/否）                                        | optional     |                                 |
| is_morning_outage              | boolean        | 客运凌晨2-5时停运（是/否）                                   | optional     |                                 |
| curve                          | int4           | 曲线情况（曲线完整/回传异常/零速度）                         | optional     | **曲线情况**字典                |
| trail                          | int4           | 轨迹情况（正常/漂移/其他）                                   | optional     | **GPS轨迹情况**字典             |
| lens_position                  | int4           | 镜头位置（正/偏）                                            | optional     | **镜头位置**字典                |
| other_infraction               | text           | 其他违规                                                     | optional     |                                 |
| lens_on                        | text           | 摄像头损坏号                                                 | optional     |                                 |
| coordinate                     | point          | 空间数据类型point表示经纬度                                  | optional     |                                 |
| tachograph_speed               | text           | 行车记录仪速度                                               | optional     |                                 |
| gps_speed                      | text           | 卫星定位速度                                                 | optional     |                                 |
| monitoring_time                | timestamptz(6) | 监控平台时间                                                 | optional     |                                 |
| is_locate                      | boolean        | 是否定位                                                     | optional     |                                 |
| disposal_measures              | text           | 处置措施                                                     | optional     |                                 |
| disposal_results               | text           | 处置结果                                                     | optional     |                                 |
| treatment_time                 | timestamptz(6) | 处置时间                                                     | optional     |                                 |
| assignee                       | text           | 受理人                                                       | optional     |                                 |
| feedback_time                  | timestamptz(6) | 反馈时间                                                     | optional     |                                 |
| remarks                        | text           | 备注                                                         | optional     |                                 |
| created_at                     | timestamptz(6) | 创建时间                                                     | required     |                                 |
| created_by                     | text           | 创建人                                                       | required     | **system_user表**的user_id      |
| updated_at                     | timestamptz(6) | 修改时间                                                     | optional     |                                 |
| updated_by                     | text           | 修改人                                                       | optional     | **system_user表**的user_id      |
| deleted_at                     | timestamptz(6) | 删除时间                                                     | optional     |                                 |
| deleted_by                     | text           | 删除人                                                       | optional     | **system_user表**的user_id      |

### traffic_illegal_info_statistics 交通违法动态信息处理和统计

营运车:

MON_WF_COUNT 交通违法动态信息处理和统计

| Name                               | Type        | Description                                     | **Required** | default                    |
| ---------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                 | bigint      | 按指定方法生成                                  | required     | 主键                       |
| traffic_illegal_info_statistics_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                         | text        | 车辆id                                          | optional     |                            |
| dept_id                            | text        | 所在部门                                        | optional     |                            |
| overspeed_below_ten                | int4        | 超速10以下                                      | optional     |                            |
| overspeed_ten_to_twenty            | int4        | 超速10以上20以下                                | optional     |                            |
| overspeed_twenty_to_fifty          | int4        | 超速20以上50以下                                | optional     |                            |
| overspeed_over_fifty               | int4        | 超速50以上                                      | optional     |                            |
| situation_analysis                 | text        | 情况分析                                        | optional     |                            |
| disposal_measures                  | text        | 处置措施                                        | optional     |                            |
| remarks                            | text        | 备注                                            | optional     |                            |
| created_at                         | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                         | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                         | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                         | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                         | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                         | text        | 删除人                                          | optional     | **system_user表**的user_id |

### province_outage_registration 省厅停运报备表

营运车:

ST_OFFLINEINFO 省厅停运报备表

| Name                            | Type        | Description                                     | **Required** | default                    |
| ------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                              | bigint      | 按指定方法生成                                  | required     | 主键                       |
| province_outage_registration_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                      | text        | 车辆id                                          | optional     |                            |
| start_time                      | timestamptz(6) | 开始时间                                        | optional     |                            |
| end_time                        | timestamptz(6) | 结束时间                                        | optional     |                            |
| correcting                      | text        | 校正                                            | optional     |                            |
| cause                           | text        | 原因                                            | optional     |                            |
| operator                        | text        | 运营商                                          | optional     |                            |
| veh_status                      | text        | 车辆状态                                        | optional     |                            |
| veh_type                        | text        | 车辆类型                                        | optional     |                            |
| created_at                      | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                      | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                      | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                      | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                      | text        | 删除人                                          | optional     | **system_user表**的user_id |

### charter_report 包车报备表

营运车:

INFO_CHARTVEH 包车报备表

| Name                          | Type        | Description                                     | **Required** | default                    |
| ----------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| charter_report_id             | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number          | text        | 车牌号                                          | optional     |                            |
| license_plate_color           | int4        | 车牌颜色                                        | optional     | **车牌颜色**字典           |
| seats                         | int4        | 座位                                            | optional     |                            |
| departure_place               | text        | 出发地                                          | optional     |                            |
| pathway_place                 | text        | 途径地                                          | optional     |                            |
| destination                   | text        | 终点地                                          | optional     |                            |
| mileage                       | text        | 里程                                            | optional     |                            |
| departure_time                | timestamptz(6) | 去程出发时间                                    | optional     |                            |
| arrival_time                  | timestamptz(6) | 去程到达时间                                    | optional     |                            |
| rest_place                    | text        | 去程休息地点                                    | optional     |                            |
| rest_two_to_five_place        | text        | 去程2点到5点休息地点                            | optional     |                            |
| return_departure_time         | timestamptz(6) | 返程发车时间                                    | optional     |                            |
| return_arrival_time           | timestamptz(6) | 返程到达时间                                    | optional     |                            |
| return_rest_place             | text        | 返程休息地点                                    | optional     |                            |
| return_rest_two_to_five_place | text        | 返程2点到5点休息地点                            | optional     |                            |
| charter_start_time            | timestamptz(6) | 包车开始时间                                    | optional     |                            |
| charter_end_time              | timestamptz(6) | 包车结束时间                                    | optional     |                            |
| guest_num                     | int4        | 载客人数                                        | optional     |                            |
| business_name                 | text        | 业户名称                                        | optional     |                            |
| apply_type                    | text        | 申请类别                                        | optional     |                            |
| driver_list                   | text[]      | 驾驶员列表                                      | optional     |                            |
| departure_area_code           | text        | 出发地行政区划代码                              | optional     |                            |
| pathway_area_code             | text        | 途径地行政区划代码（暂时为空，以后扩展）        | optional     |                            |
| destination_area_code         | text        | 终点地行政区划代码                              | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                    | text        | 删除人                                          | optional     | **system_user表**的user_id |

### security_checklist 安检表

营运车:

INFO_SAFECHECK 安检表

| Name                  | Type        | Description                                     | **Required** | default                    |
| --------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                    | bigint      | 按指定方法生成                                  | required     | 主键                       |
| security_checklist_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number  | text        | 车牌号                                          | optional     |                            |
| license_plate_color   | int4        | 车牌颜色                                        | optional     | **车牌颜色**字典           |
| seats                 | int4        | 座位                                            | optional     |                            |
| check_time            | timestamptz(6) | 检查时间                                        | optional     |                            |
| is_qualified          | boolean     | 是否合格                                        | optional     |                            |
| add_person_name       | text        | 新增人姓名                                      | optional     |                            |
| add_time              | timestamptz(6) | 新增时间                                        | optional     |                            |
| ent_id                | text        | 所属企业名称                                    | optional     |                            |
| created_at            | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by            | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by            | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by            | text        | 删除人                                          | optional     | **system_user表**的user_id |

### enter_xm_whitelist 入厦白名单

营运车:

INFO_XM_WHITELIST 入厦白名单

| Name                  | Type        | Description                                     | **Required** | default                    |
| --------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                    | bigint      | 按指定方法生成                                  | required     | 主键                       |
| enter_xm_whitelist_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number  | text        | 车牌号                                          | optional     |                            |
| veh_id                | text        | 车辆id                                          | optional     |                            |
| sim_num               | text        | sim卡号                                         | optional     |                            |
| type                  | text        | 类型                                            | optional     |                            |
| area_id               | text        | 入站ID                                          | optional     |                            |
| created_at            | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by            | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at            | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by            | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at            | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by            | text        | 删除人                                          | optional     | **system_user表**的user_id |

### driver_card_manager 司机卡信息

营运车:

INFO_DRIVER_CARD 司机卡管理

| Name                      | Type           | Description                                     | **Required** | default                    |
| ------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | bigint         | 按指定方法生成                                  | required     | 主键                       |
| driver_card_manager_id    | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| driver_id                 | text           | 驾驶员信息表的driver_id                         | optional     |                            |
| driver_name               | text           | 驾驶员姓名                                      | optional     |                            |
| dept_id                   | text           | 所在部门                                        | optional     |                            |
| identity_card             | text           | 驾驶员编号                                      | optional     |                            |
| driving_license           | text           | 驾驶证号                                        | optional     |                            |
| qualification_id          | text           | 从业资格证号码                                  | optional     |                            |
| issuing_authority         | text           | 从业资格证发证机构                              | optional     |                            |
| encrypt_identity_card     | text           | 加密驾驶员编号                                  | optional     |                            |
| encrypt_driver_name       | text           | 加密驾驶员姓名                                  | optional     |                            |
| encrypt_driving_license   | text           | 加密驾驶证号                                    | optional     |                            |
| encrypt_qualification_id  | text           | 加密从业资格证号码                              | optional     |                            |
| encrypt_issuing_authority | text           | 加密从业资格证发证机构                          | optional     |                            |
| is_deleted                | boolean        | 是否删除                                        | optional     |                            |
| is_audit                  | boolean        | 是否审核                                        | optional     |                            |
| encrypt                   | text           | 密钥                                            | optional     |                            |
| driver_card_type          | text           | 司机卡类型                                      | optional     |                            |
| occupational_expire_date  | timestamptz(6) | 从业资格证有效期                                | optional     |                            |
| created_at                | timestamptz(6)    | 创建时间                                        | required     |                            |
| created_by                | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6)    | 修改时间                                        | optional     |                            |
| updated_by                | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6)    | 删除时间                                        | optional     |                            |
| deleted_by                | text           | 删除人                                          | optional     | **system_user表**的user_id |

### voice_text_sending_record 语音文本发送记录

营运车:

MON_MESSAGE_SEND 语音文本发送日志

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| voice_text_sending_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id                   | text        | 车辆id                                          | optional     |                            |
| send_type                    | text        | 发送方式                                        | optional     |                            |
| msg_content                  | text        | 消息内容                                        | optional     |                            |
| send_ent                     | text        | 发送企业                                        | optional     |                            |
| send_result                  | text        | 发送结果                                        | optional     |                            |
| operator                     | text        | 操作人                                          | optional     |                            |
| created_at                   | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### offline_msg_remind_record 离线4小时短信提醒记录

营运车:

MON_OFFLINE4_ALARM 离线4小时短信提醒日志

| Name                         | Type           | Description                                     | **Required** | default                    |
| ---------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | int8           | 按指定方法生成                                  | required     | 主键                       |
| offline_msg_remind_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| offline_start_time           | timestamptz(6) | 离线开始时间                                    | optional     |                            |
| offline_end_time             | timestamptz(6) | 离线结束时间                                    | optional     |                            |
| license_plate_number         | text           | 车牌号                                          | optional     |                            |
| registration_user            | text           | 登记用户                                        | optional     | **system_user表**的user_id |
| registration_time            | timestamptz(6) | 登记时间                                        | optional     |                            |
| sms_content                  | text           | 短信发送内容                                    | optional     |                            |
| phone_reminder_content       | text           | 电话提醒内容                                    | optional     |                            |
| sms_send_time                | timestamptz(6) | 短信发送时间                                    | optional     |                            |
| phone_reminder_time          | timestamptz(6) | 电话提醒时间                                    | optional     |                            |
| offline_reason               | text           | 离线原因                                        | optional     |                            |
| alarm_type                   | int4           | 报警类型                                        | optional     | **报警类型**字典           |
| is_registration              | boolean        | 是否登记                                        | optional     |                            |
| is_end_alarm                 | boolean        | 是否结束报警                                    | optional     |                            |
| is_send_sms                  | boolean        | 是否发送短信                                    | optional     |                            |
| is_need_maintain             | boolean        | 是否需要维护                                    | optional     |                            |
| created_at                   | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                   | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                   | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                   | text           | 删除人                                          | optional     | **system_user表**的user_id |

### dynamic_alarm_dispose_record 动态监控报警记录处置

营运车:

INFO_DYNAMIC_ALARM_DEAL  动态监控报警记录处置

| Name                            | Type           | Description                                     | **Required** | default                    |
| ------------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                              | int8           | 按指定方法生成                                  | required     | 主键                       |
| dynamic_alarm_dispose_record_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                          | text           | 车辆id                                          | optional     |                            |
| ent_id                          | text           | 企业id                                          | optional     |                            |
| phone                           | text           | 手机号码                                        | optional     |                            |
| alarm_type                      | text           | 报警类型                                        | optional     |                            |
| alarm_id                        | text           | 报警表ID                                        | optional     |                            |
| processing_content              | text           | 处理内容                                        | optional     |                            |
| is_sms_push                     | boolean        | 是否短信推送                                    | optional     |                            |
| processing_time                 | timestamptz(6) | 处理时间                                        | optional     |                            |
| processing_type                 | int4           | 处理类型                                        | optional     |                            |
| operation_user                  | text           | 操作用户                                        | optional     | **system_user表**的user_id |
| is_notify                       | boolean        | 是否通报                                        | optional     |                            |
| is_announce                     | boolean        | 是否语音通知                                    | optional     |                            |
| is_app_push                     | boolean        | 是否APP推送                                     | optional     |                            |
| notify_content                  | text           | 通报内容                                        | optional     |                            |
| announce_content                | text           | 语音内容                                        | optional     |                            |
| app_push_content                | text           | APP推送内容                                     | optional     |                            |
| disposal_method                 | int4           | 处置方式                                        | optional     | **处置方式**字典           |
| disposal_result                 | text           | 处置结果                                        | optional     |                            |
| is_deleted                      | boolean        | 是否删除                                        | optional     | false                      |
| created_at                      | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                      | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                      | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                      | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                      | text           | 删除人                                          | optional     | **system_user表**的user_id |

### dynamic_veh_record_investigation 道路运输车辆动态监控记录排查单位

营运车:

INFO_DYNAMIC_CHECK_RG  道路运输车辆动态监控记录排查单位

| Name                                | Type        | Description                                     | **Required** | default                    |
| ----------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                  | bigint      | 按指定方法生成                                  | required     | 主键                       |
| dynamic_veh_record_investigation_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| ent_id                              | text        | 所属企业                                        | optional     |                            |
| is_completed                        | boolean     | 是否抽查完成                                    | optional     |                            |
| created_at                          | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                          | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                          | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                          | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                          | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                          | text        | 删除人                                          | optional     | **system_user表**的user_id |

### dynamic_ent_check_detail 企业动态监管抽查明细  

营运车:

INFO_DYNAMIC_CHECK_RG_DETAIL 企业动态监管抽查明细

| Name                           | Type           | Description                                                  | **Required** | default                              |
| ------------------------------ | -------------- | ------------------------------------------------------------ | ------------ | ------------------------------------ |
| id                             | int8           | 按指定方法生成                                               | required     | 主键                                 |
| dynamic_ent_check_detail_id    | text           | 动态监管抽查明细表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                             |
| driver_id                      | text           | 驾驶员id                                                     | optional     | **driver_info 表**的driver_id        |
| enterprise_id                  | text           | 车辆所在企业id                                               | optional     | **enterprise_info表**的enterprise_id |
| vehicle_id                     | text           | 车辆id                                                       | optional     | **vehicle_info 表**的vehicle_id      |
| monitoring_time                | timestamptz(6) | 监控平台时间                                                 | optional     |                                      |
| monitoring_location            | text           | 监控平台显示位置                                             | optional     |                                      |
| is_online                      | boolean        | 是否在线（是/否）                                            | optional     |                                      |
| is_tachograph_record_normal    | boolean        | 行车记录仪数据（是否异常）                                   | optional     |                                      |
| tachograph_data_exception      | text           | 行车记录仪异常数据项                                         | optional     |                                      |
| tachograph_speed               | text           | 行车记录仪速度                                               | optional     |                                      |
| gps_speed                      | text           | 卫星定位速度                                                 | optional     |                                      |
| is_speeding                    | boolean        | 是否超速（是/否）                                            | optional     |                                      |
| taxi_state                     | int4           | 出租空/重车状态（空/重）                                     | optional     | **出租车空/重车状态**字典            |
| is_fatigue_driving             | boolean        | 客运疲劳驾驶（是/否）                                        | optional     |                                      |
| is_morning_outage              | boolean        | 客运凌晨2-5时停运（是/否）                                   | optional     |                                      |
| curve                          | int4           | 曲线情况（曲线完整/回传异常/零速度）                         | optional     | **曲线情况**字典                     |
| trail                          | int4           | 轨迹情况（正常/漂移/其他）                                   | optional     | **GPS轨迹情况**字典                  |
| lens_position                  | int4           | 镜头位置（正/偏）                                            | optional     | **镜头位置**字典                     |
| equipment                      | int4           | 设备情况（图像正常/无图像/摄像头 号损坏）                    | optional     | **设备情况**字典                     |
| other_infraction               | text           | 其他违规                                                     | optional     |                                      |
| disposal_measures              | text           | 处置措施                                                     | optional     |                                      |
| disposal_results               | text           | 处置结果                                                     | optional     |                                      |
| treatment_time                 | timestamptz(6) | 处置时间                                                     | optional     |                                      |
| assignee                       | text           | 受理人                                                       | optional     |                                      |
| feedback_time                  | timestamptz(6) | 反馈时间                                                     | optional     |                                      |
| remarks                        | text           | 备注                                                         | optional     |                                      |
| others                         | text           | 轨迹其他情况                                                 | optional     |                                      |
| lens_on                        | text           | 摄像头损坏号                                                 | optional     |                                      |
| monitor_end_time               | timestamptz(6) | 监管费到期时间                                               | optional     |                                      |
| is_locate                      | boolean        | 是否定位                                                     | optional     |                                      |
| coordinate                     | point          | 空间数据类型point表示经纬度                                  | optional     |                                      |
| latitude_longitude_description | text           | 经纬度描述                                                   | optional     |                                      |
| is_send                        | boolean        | 是否发送                                                     | optional     |                                      |
| business_scope                 | int4           | 经营范围                                                     | optional     | **经营范围**字典                     |
| outage_alarm_time              | timestamptz(6) | 凌晨2点到5点停运报警时间                                     | optional     |                                      |
| speed_alarm_time               | timestamptz(6) | 超速报警时间                                                 | optional     |                                      |
| speeding_speed                 | text           | 超速速度                                                     | optional     |                                      |
| fatigue_alarm_time             | timestamptz(6) | 疲劳驾驶报警时间                                             | optional     |                                      |
| disposal_measures1             | text           | 是否在线处置措施                                             | optional     |                                      |
| disposal_measures2             | text           | 是否超速处置措施                                             | optional     |                                      |
| disposal_measures3             | text           | 曲线情况处置措施                                             | optional     |                                      |
| disposal_measures4             | text           | 客运疲劳驾驶处置措施                                         | optional     |                                      |
| disposal_measures5             | text           | 客运凌晨停运处置措施                                         | optional     |                                      |
| disposal_measures6             | text           | 行车记录仪数据处置措施                                       | optional     |                                      |
| disposal_measures7             | text           | 轨迹情况处置措施                                             | optional     |                                      |
| disposal_results1              | text           | 是否在线处置结果                                             | optional     |                                      |
| disposal_results2              | text           | 是否超速处置结果                                             | optional     |                                      |
| disposal_results3              | text           | 曲线情况处置结果                                             | optional     |                                      |
| disposal_results4              | text           | 客运疲劳驾驶处置结果                                         | optional     |                                      |
| disposal_results5              | text           | 客运疲劳驾驶处置结果                                         | optional     |                                      |
| disposal_results6              | text           | 行车记录仪数据处置结果                                       | optional     |                                      |
| disposal_results7              | text           | 轨迹情况处置结果                                             | optional     |                                      |
| is_deleted                     | boolean        | 是否被删除                                                   | optional     |                                      |
| created_at                     | timestamptz(6) | 创建时间                                                     | required     |                                      |
| created_by                     | text           | 创建人                                                       | required     | **system_user表**的user_id           |
| updated_at                     | timestamptz(6) | 修改时间                                                     | optional     |                                      |
| updated_by                     | text           | 修改人                                                       | optional     | **system_user表**的user_id           |
| deleted_at                     | timestamptz(6) | 删除时间                                                     | optional     |                                      |
| deleted_by                     | text           | 删除人                                                       | optional     | **system_user表**的user_id           |

### dynamic_check_dispose 动态抽查处置表

营运车:

INFO_DYNAMIC_CHECK_RG_DEAL 动态抽查处置表

| Name                     | Type           | Description                                     | **Required** | default                    |
| ------------------------ | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | int8           | 按指定方法生成                                  | required     | 主键                       |
| dynamic_check_dispose_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| enterprise_id            | text           | 所在企业id                                      | optional     |                            |
| vehicle_id               | text           | 车辆id                                          | optional     |                            |
| telephone                | text           | 手机号码                                        | optional     |                            |
| content                  | text           | 内容                                            | optional     |                            |
| send_time                | timestamptz(6) | 发送时间                                        | optional     |                            |
| user_id                  | text           | 用户ID                                          | optional     | **system_user表**的user_id |
| is_sms_push              | boolean        | 是否短信推送                                    | optional     |                            |
| is_report                | boolean        | 是否通报                                        | optional     |                            |
| is_voice_notification    | boolean        | 是否语音通知                                    | optional     |                            |
| is_app_push              | boolean        | 是否APP推送                                     | optional     |                            |
| notification_content     | text           | 通报内容                                        | optional     |                            |
| voice_content            | text           | 语音内容                                        | optional     |                            |
| app_push_content         | text           | APP推送内容                                     | optional     |                            |
| supervision_detail_id    | text           | 抽查表ID                                        | optional     |                            |
| disposal_method          | text           | 处置方式                                        | optional     |                            |
| is_deleted               | boolean        | 是否删除                                        | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text           | 删除人                                          | optional     | **system_user表**的user_id |

### dynamic_veh_record_statistics 道路运输车辆动态监控记录表统计

营运车:

INFO_DYNAMIC_CHECK_RG_TJ 道路运输车辆动态监控记录表统计

| Name                             | Type           | Description                                     | **Required** | default                    |
| -------------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                               | int8           | 按指定方法生成                                  | required     | 主键                       |
| dynamic_veh_record_statistics_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| ent_id                           | text           | 企业id                                          | optional     |                            |
| spot_checks_num                  | int4           | 应抽查数                                        | optional     |                            |
| actual_spot_checks_num           | int4           | 实抽查数                                        | optional     |                            |
| exception_num                    | int4           | 异常数                                          | optional     |                            |
| disposal_num                     | int4           | 处置数                                          | optional     |                            |
| area_id                          | text           | 地区id                                          | optional     |                            |
| created_at                       | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                       | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                       | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                       | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                       | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                       | text           | 删除人                                          | optional     | **system_user表**的user_id |

### dynamic_spot_check_statistics 道路运输车辆动态监控记录抽查统计

营运车:

INFO_DYNAMIC_CHECK_COUNTY_TJ 道路运输车辆动态监控记录抽查统计

| Name                             | Type           | Description                                     | **Required** | default                    |
| -------------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                               | int8           | 按指定方法生成                                  | required     | 主键                       |
| dynamic_spot_check_statistics_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| spot_checks_num                  | int4           | 应抽查数                                        | optional     |                            |
| actual_spot_checks_num           | int4           | 实抽查数                                        | optional     |                            |
| exception_num                    | int4           | 异常数                                          | optional     |                            |
| disposal_num                     | int4           | 处置数                                          | optional     |                            |
| area_id                          | text           | 地区id                                          | optional     |                            |
| is_completed                     | boolean        | 是否排查完成                                    | optional     |                            |
| created_at                       | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                       | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                       | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                       | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                       | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                       | text           | 删除人                                          | optional     | **system_user表**的user_id |

### dynamic_ent_spot_check_manually 道路运输车辆动态监控记录企业手动登记表

营运车:

INFO_DYNAMIC_CHECK_RG_MANUAL 道路运输车辆动态监控记录企业手动登记表

| Name                               | Type           | Description                                     | **Required** | default                    |
| ---------------------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                 | int8           | 按指定方法生成                                  | required     | 主键                       |
| dynamic_ent_spot_check_manually_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| driver_id                          | text           | 驾驶员id                                        | optional     |                            |
| enterprise_id                      | text           | 车辆所在企业id                                  | optional     |                            |
| vehicle_id                         | text           | 车辆id                                          | optional     |                            |
| monitoring_time                    | timestamptz(6) | 监控平台时间                                    | optional     |                            |
| monitoring_location                | text           | 监控平台显示位置                                | optional     |                            |
| is_online                          | boolean        | 是否在线（是/否）                               | optional     |                            |
| is_tachograph_record_normal        | boolean        | 行车记录仪数据（是否异常）                      | optional     |                            |
| tachograph_data_exception          | text           | 行车记录仪异常数据项                            | optional     |                            |
| tachograph_speed                   | text           | 行车记录仪速度                                  | optional     |                            |
| gps_speed                          | text           | 卫星定位速度                                    | optional     |                            |
| is_speeding                        | boolean        | 是否超速（是/否）                               | optional     |                            |
| taxi_state                         | int4           | 出租空/重车状态（空/重）                        | optional     | **出租车空/重车状态**字典  |
| is_fatigue_driving                 | boolean        | 客运疲劳驾驶（是/否）                           | optional     |                            |
| is_morning_outage                  | boolean        | 客运凌晨2-5时停运（是/否）                      | optional     |                            |
| curve                              | int4           | 曲线情况（曲线完整/回传异常/零速度）            | optional     | **曲线情况**字典           |
| trail                              | int4           | 轨迹情况（正常/漂移/其他）                      | optional     | **GPS轨迹情况**字典        |
| lens_position                      | int4           | 镜头位置（正/偏）                               | optional     | **镜头位置**字典           |
| equipment                          | int4           | 设备情况（图像正常/无图像/摄像头 号损坏）       | optional     | **设备情况**字典           |
| other_infraction                   | text           | 其他违规                                        | optional     |                            |
| disposal_measures                  | text           | 处置措施                                        | optional     |                            |
| disposal_results                   | text           | 处置结果                                        | optional     |                            |
| treatment_time                     | timestamptz(6) | 处置时间                                        | optional     |                            |
| assignee                           | text           | 受理人                                          | optional     |                            |
| feedback_time                      | timestamptz(6) | 反馈时间                                        | optional     |                            |
| remarks                            | text           | 备注                                            | optional     |                            |
| others                             | text           | 轨迹其他情况                                    | optional     |                            |
| lens_on                            | text           | 摄像头损坏号                                    | optional     |                            |
| monitor_end_time                   | timestamptz(6) | 监管费到期时间                                  | optional     |                            |
| is_locate                          | boolean        | 是否定位                                        | optional     |                            |
| coordinate                         | point          | 空间数据类型point表示经纬度                     | optional     |                            |
| latitude_longitude_description     | text           | 经纬度描述                                      | optional     |                            |
| is_send                            | boolean        | 是否发送                                        | optional     |                            |
| business_scope                     | int4           | 经营范围                                        | optional     | **经营范围**字典           |
| outage_alarm_time                  | timestamptz(6) | 凌晨2点到5点停运报警时间                        | optional     |                            |
| speed_alarm_time                   | timestamptz(6) | 超速报警时间                                    | optional     |                            |
| speeding_speed                     | text           | 超速速度                                        | optional     |                            |
| fatigue_alarm_time                 | timestamptz(6) | 疲劳驾驶报警时间                                | optional     |                            |
| disposal_measures1                 | text           | 是否在线处置措施                                | optional     |                            |
| disposal_measures2                 | text           | 是否超速处置措施                                | optional     |                            |
| disposal_measures3                 | text           | 曲线情况处置措施                                | optional     |                            |
| disposal_measures4                 | text           | 客运疲劳驾驶处置措施                            | optional     |                            |
| disposal_measures5                 | text           | 客运凌晨停运处置措施                            | optional     |                            |
| disposal_measures6                 | text           | 行车记录仪数据处置措施                          | optional     |                            |
| disposal_measures7                 | text           | 轨迹情况处置措施                                | optional     |                            |
| disposal_results1                  | text           | 是否在线处置结果                                | optional     |                            |
| disposal_results2                  | text           | 是否超速处置结果                                | optional     |                            |
| disposal_results3                  | text           | 曲线情况处置结果                                | optional     |                            |
| disposal_results4                  | text           | 客运疲劳驾驶处置结果                            | optional     |                            |
| disposal_results5                  | text           | 客运疲劳驾驶处置结果                            | optional     |                            |
| disposal_results6                  | text           | 行车记录仪数据处置结果                          | optional     |                            |
| disposal_results7                  | text           | 轨迹情况处置结果                                | optional     |                            |
| dynamic_type                       | int4           | 动态抽查类型                                    | require      |                            |
| is_deleted                         | boolean        | 是否被删除                                      | optional     |                            |
| created_at                         | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                         | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                         | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                         | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                         | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                         | text           | 删除人                                          | optional     | **system_user表**的user_id |

### alarm_disposal_manually 手动录入报警处置记录  

营运车:

MONITOR_ALARM_DISPOSAL 手动录入报警处置记录

| Name                       | Type        | Description                                     | **Required** | default                    |
| -------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                  | required     | 主键                       |
| alarm_disposal_manually_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                     | text        | 车辆类型                                        | optional     |                            |
| monitor_time               | timestamptz(6) | 监控平台时间                                    | optional     |                            |
| monitor_display_location   | text        | 监控平台显示位置                                | optional     |                            |
| alarm_type                 | text        | 报警类型                                        | optional     |                            |
| disposal_measure           | text        | 处置措施                                        | optional     |                            |
| disposal_result            | text        | 处置结果                                        | optional     |                            |
| alarm_release_time         | timestamptz(6) | 报警解除时间                                    | optional     |                            |
| remarks                    | text        | 备注                                            | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text        | 删除人                                          | optional     | **system_user表**的user_id |

### dynamic_veh_record_setting 道路运输车辆动态监控记录表设置

营运车:

INFO_DYNAMIC_CHECK_CONFIG 道路运输车辆动态监控记录表设置

| Name                          | Type        | Description                                     | **Required** | default                    |
| ----------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| dynamic_veh_record_setting_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| user_id                       | text        | 用户ID                                          | optional     |                            |
| check_type_one                | int4        | 抽查重型货车企业或个体                          | optional     |                            |
| check_type_two                | int4        | 抽查两客一危企业或个体                          | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                    | text        | 删除人                                          | optional     | **system_user表**的user_id |

### driver_card_send_instruction 司机卡发送指令

营运车:

INFO_DRIVER_SMS 司机卡发送指令

| Name                            | Type        | Description                                     | **Required** | default                    |
| ------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                              | bigint      | 按指定方法生成                                  | required     | 主键                       |
| driver_card_send_instruction_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| driver_card_id                  | text        | 司机卡ID                                        | optional     |                            |
| driver_name                     | text        | 驾驶员姓名                                      | optional     |                            |
| id_card_num                     | text        | 身份证号码                                      | optional     |                            |
| phone_num                       | text        | 手机号码                                        | optional     |                            |
| content                         | text        | 内容                                            | optional     |                            |
| user_id                         | text        | 用户ID                                          | optional     |                            |
| created_at                      | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                      | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                      | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                      | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                      | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                      | text        | 删除人                                          | optional     | **system_user表**的user_id |

### monitor_whitelist 监控白名单

营运车:

INFO_MONITOR_WHITE_LIST 监控白名单

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| monitor_whitelist_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number | text        | 车牌号                                          | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### upload_province_evidence 上传省厅证明

营运车:

INFO_UPLOADPROVINCE_PIC 上传省厅证明

| Name                        | Type        | Description                                     | **Required** | default                    |
| --------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                          | bigint      | 按指定方法生成                                  | required     | 主键                       |
| upload_province_evidence_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                      | text        | 车辆id                                          | optional     |                            |
| upload_province_pic         | text        | 上传省厅截图                                    | optional     |                            |
| state                       | text        | 状态                                            | optional     |                            |
| created_at                  | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                  | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                  | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                  | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                  | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                  | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_cert_bind 车辆证书绑定

营运车:

INFO_VEH_CERT 车辆证书绑定

| Name             | Type        | Description                                     | **Required** | default                    |
| ---------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id               | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_cert_bind_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id           | text        | 车辆id                                          | optional     |                            |
| cert_path        | text        | 证书存放路径                                    | optional     |                            |
| is_deleted       | boolean     | 是否删除                                        | optional     |                            |
| created_at       | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by       | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at       | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by       | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at       | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by       | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_online_status 车辆在线情况

营运车:

INFO_VEH_ONLINETIME 车辆在线情况

| Name                       | Type        | Description                                     | **Required** | default                    |
| -------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_online_status_id       | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| online_total_time          | int4        | 在线总时间(秒,连续在线大于20分钟的时间段才算入) | optional     |                            |
| max_continuous_online_time | int4        | 最长的一次连续在线时间(秒)                      | optional     |                            |
| first_online_time          | timestamptz(6) | 当天首次在线时间                                | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_online_status_that_day 车辆当天在线情况

营运车:

INFO_VEH_ONLINETIME_CURRENT 车辆当天在线情况

| Name                          | Type        | Description                                     | **Required** | default                    |
| ----------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_online_status_that_day_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| online_total_time             | int4        | 在线总时间(秒,连续在线大于20分钟的时间段才算入) | optional     |                            |
| max_continuous_online_time    | int4        | 最长的一次连续在线时间(秒)                      | optional     |                            |
| first_online_time             | timestamptz(6) | 当天首次在线时间                                | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                    | text        | 删除人                                          | optional     | **system_user表**的user_id |

### palteform_investigation_detail 平台查岗明细

营运车:

INFO_DOWN_PLATFORMQUERY 平台查岗明细

| Name                              | Type        | Description                                     | **Required** | default                    |
| --------------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                                | bigint      | 按指定方法生成                                  | required     | 主键                       |
| palteform_investigation_detail_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| obj_type                          | text        | 对象类型                                        | optional     |                            |
| ent_id                            | text        | 所查企业的ID                                    | optional     |                            |
| subsequent_date_length            | text        | 后续数据长度                                    | optional     |                            |
| msg_id                            | text        | 信息ID                                          | optional     |                            |
| date_length                       | text        | 信息长度                                        | optional     |                            |
| data_content                      | text        | 信息内容                                        | optional     |                            |
| build_json                        | text        | 平台组装的JSON                                  | optional     |                            |
| is_response                       | text        | 是否应答                                        | optional     |                            |
| response_time                     | timestamptz(6) | 应答时间                                        | optional     |                            |
| reply_content                     | text        | 回复内容（JSON）                                | optional     |                            |
| is_response_correct               | boolean     | 应答答案是否正确                                | optional     |                            |
| is_need_down                      | boolean     | 是否需要下发                                    | optional     |                            |
| created_at                        | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                        | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                        | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                        | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                        | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_egn_approval_info 核准证信息表

营运车:

ZTC_APPROVAL_INFORMATION 新型车辆核准证信息

工程车：

ZTC_APPROVAL_INFORMATION 核准证信息表

| Name                     | Type        | Description                                     | **Required** | default                    |
| ------------------------ | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                       | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_egn_approval_info_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                   | text        | 车辆id                                          | optional     |                            |
| sim_num                  | text    | 终端SIM卡号                                     | optional     |                            |
| msg_id                   | text    | 消息ID                                          | optional     |                            |
| oper_type                | text    | 操作类型                                        | optional     |                            |
| business_value           | text    | 业务值                                          | optional     |                            |
| start_time               | timestamptz(6) | 开始时间                                        | optional     |                            |
| end_time                 | timestamptz(6) | 结束时间                                        | optional     |                            |
| content                  | text    | 内容                                            | optional     |                            |
| oper_time                | timestamptz(6) | 操作时间                                        | optional     |                            |
| feedback_time            | timestamptz(6) | 反馈时间                                        | optional     |                            |
| is_success               | boolean     | 是否成功                                        | optional     |                            |
| is_deleted               | boolean     | 是否删除                                        | optional     |                            |
| created_at               | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by               | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at               | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by               | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at               | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by               | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_state_upload 新型车辆状态上报

营运车:

INFO_VEHSTATEUPLOAD 新型车辆状态上报

| Name                | Type        | Description                                          | **Required** | default                    |
| ------------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                  | bigint      | 按指定方法生成                                       | required     | 主键                       |
| veh_state_upload_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| veh_id              | text        | 车辆id                                               | optional     |                            |
| status              | text        | 状态                                                 | optional     |                            |
| coordinate          | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| altitude            | numeric     | 高程                                                 | optional     |                            |
| speed               | numeric     | 速度                                                 | optional     |                            |
| direction           | numeric     | 方向                                                 | optional     |                            |
| driver_id           | text        | 司机ID                                               | optional     |                            |
| approval_id         | text        | 核准证ID                                             | optional     |                            |
| veh_case_status     | text        | 车箱状态                                             | optional     |                            |
| lift_status         | text        | 举升状态                                             | optional     |                            |
| empty_weight_status | text        | 空重状态                                             | optional     |                            |
| violation_status    | text        | 违规情况                                             | optional     |                            |
| overload_simulation | text        | 重空载模拟量                                         | optional     |                            |
| created_at          | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by          | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at          | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by          | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at          | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by          | text        | 删除人                                               | optional     | **system_user表**的user_id |

### veh_event_upload 新型车辆事件上报

营运车:

INFO_EVENTUPLOAD 新型车辆事件上报

| Name                | Type        | Description                                          | **Required** | default                    |
| ------------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                  | bigint      | 按指定方法生成                                       | required     | 主键                       |
| veh_event_upload_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| event_num           | text        | 事件流水号                                           | optional     |                            |
| event_type          | text        | 事件类型                                             | optional     |                            |
| coordinate          | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| altitude            | numeric     | 高程                                                 | optional     |                            |
| speed               | numeric     | 速度                                                 | optional     |                            |
| direction           | numeric     | 方向                                                 | optional     |                            |
| driver_id           | text        | 司机ID                                               | optional     |                            |
| approval_id         | text        | 核准证ID                                             | optional     |                            |
| veh_case_status     | text        | 车箱状态                                             | optional     |                            |
| lift_status         | text        | 举升状态                                             | optional     |                            |
| empty_weight_status | text        | 空重状态                                             | optional     |                            |
| violation_status    | text        | 违规情况                                             | optional     |                            |
| overload_simulation | text        | 重空载模拟量                                         | optional     |                            |
| created_at          | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by          | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at          | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by          | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at          | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by          | text        | 删除人                                               | optional     | **system_user表**的user_id |

### veh_self_check 新型车辆自检信息

营运车:

INFO_SELFCHECK 新型车辆自检信息

| Name                 | Type        | Description                                          | **Required** | default                    |
| -------------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                       | required     | 主键                       |
| veh_self_check_id    | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| license_plate_number | text        | 车牌号                                               | optional     |                            |
| driver_id            | text        | 司机ID                                               | optional     |                            |
| approval_id          | text        | 核准证ID                                             | optional     |                            |
| coordinate           | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by           | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by           | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by           | text        | 删除人                                               | optional     | **system_user表**的user_id |

### veh_verify_info 新型车辆验证信息

营运车:

INFO_VERIFICATION  新型车辆验证信息

| Name                 | Type        | Description                                          | **Required** | default                    |
| -------------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                       | required     | 主键                       |
| veh_verify_info_id   | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| license_plate_number | text        | 车牌号                                               | optional     |                            |
| driver_id            | text        | 司机ID                                               | optional     |                            |
| approval_id          | text        | 核准证ID                                             | optional     |                            |
| coordinate           | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| password             | text        | 密码                                                 | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by           | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by           | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by           | text        | 删除人                                               | optional     | **system_user表**的user_id |

### veh_fee_transfer 车辆费用转移表

营运车:

FEE_MOVEBILL 车辆费用转移表

| Name                              | Type           | Description                  | **Required** | default                    |
| --------------------------------- | -------------- | ---------------------------- | ------------ | -------------------------- |
| id                                | bigint         | 按指定方法生成               | required     | 主键                       |
| veh_fee_transfer_id               | text           | 费用转移单号(司机卡缴费单id) | optional     |                            |
| ent_id                            | text           | 所在企业                     | optional     |                            |
| fee_status                        | int4           | 单据状态                     | optional     | **缴费状态**字典           |
| src_veh_cart                      | text           | 源车辆                       | optional     |                            |
| src_veh_regis_time                | timestamptz(6) | 源车辆注册日期               | optional     |                            |
| src_veh_old_sim_fee_deadline      | timestamptz(6) | 源车辆原SIM卡费截止日期      | optional     |                            |
| src_veh_new_sim_fee_deadline      | timestamptz(6) | 源车辆新SIM卡费截止日期      | optional     |                            |
| src_veh_old_monitor_fee_deadline  | timestamptz(6) | 源车辆原监管费截止日期       | optional     |                            |
| src_veh_new_monitor_fee_deadline  | timestamptz(6) | 源车辆新监管费截止日期       | optional     |                            |
| src_veh_old_service_fee_deadline  | timestamptz(6) | 源车辆原服务费截止日期       | optional     |                            |
| src_veh_new_service_fee_deadline  | timestamptz(6) | 源车辆新服务费截止日期       | optional     |                            |
| transfer_date                     | timestamptz(6) | 转移日期                     | optional     |                            |
| dest_veh                          | text           | 目的车辆                     | optional     |                            |
| dest_veh_regis_time               | timestamptz(6) | 目的车辆注册日期             | optional     |                            |
| dest_veh_old_sim_fee_deadline     | timestamptz(6) | 目的车辆原SIM卡费截止日期    | optional     |                            |
| dest_veh_new_sim_fee_deadline     | timestamptz(6) | 目的车辆最终SIM卡费截止日期  | optional     |                            |
| dest_veh_old_monitor_fee_deadline | timestamptz(6) | 目的车辆原监管费截止日期     | optional     |                            |
| dest_veh_new_monitor_fee_deadline | timestamptz(6) | 目的车辆最终监管费截止日期   | optional     |                            |
| dest_veh_old_service_fee_deadline | timestamptz(6) | 目的车辆原服务费截止日期     | optional     |                            |
| dest_veh_new_service_fee_deadline | timestamptz(6) | 目的车辆最终服务费截止日期   | optional     |                            |
| revocation_operator               | text           | 撤销操作人                   | optional     |                            |
| revocation_time                   | timestamptz(6) | 撤销时间                     | optional     |                            |
| revocation_reason                 | int4           | 撤销原因                     | optional     |                            |
| producer                          | text           | 制单人                       | optional     |                            |
| produce_date                      | timestamptz(6) | 制单日期                     | optional     |                            |
| audit_operator                    | text           | 审核操作人                   | optional     |                            |
| audit_time                        | timestamptz(6) | 审核时间                     | optional     |                            |
| remarks                           | text           | 备注                         | optional     |                            |
| created_at                        | timestamptz(6) | 创建时间                     | required     |                            |
| created_by                        | text           | 创建人                       | required     | **system_user表**的user_id |
| updated_at                        | timestamptz(6) | 修改时间                     | optional     |                            |
| updated_by                        | text           | 修改人                       | optional     | **system_user表**的user_id |
| deleted_at                        | timestamptz(6) | 删除时间                     | optional     |                            |
| deleted_by                        | text           | 删除人                       | optional     | **system_user表**的user_id |

### refund_doc_info 退款单据信息

营运车:

FEE_RET_BILL 退款单据信息

| Name                    | Type           | Description                                     | **Required** | default                    |
| ----------------------- | -------------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint         | 按指定方法生成                                  | required     | 主键                       |
| payment_receipt_info_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| payment_ent             | text           | 退款企业                                        | optional     |                            |
| payment_veh_num         | int4           | 退款车辆数                                      | optional     |                            |
| payment_date            | timestamptz(6) | 退款日期                                        | optional     |                            |
| total_payment_amount    | numeric        | 退款总金额                                      | optional     |                            |
| payment_step            | int4           | 退款进度                                        | optional     |                            |
| payment_status          | int4           | 退款状态                                        | optional     | **缴费状态**字典           |
| revocation_operator     | text           | 撤销操作人                                      | optional     |                            |
| revocation_time         | timestamptz(6) | 撤销时间                                        | optional     |                            |
| revocation_reason       | int4           | 撤销原因                                        | optional     | **缴费撤销原因**字典       |
| service_total_fee       | numeric        | 服务费总额                                      | optional     |                            |
| service_payment_months  | int4           | 服务费退款月数                                  | optional     |                            |
| payment_operator        | text           | 退款操作人                                      | optional     |                            |
| audit_operator          | text           | 审核操作人                                      | optional     |                            |
| audit_time              | timestamptz(6) | 审核时间                                        | optional     |                            |
| remarks                 | text           | 备注                                            | optional     |                            |
| fee_type                | int4           | 费用类型                                        | optional     | **费用类型**字典           |
| revocation_fee_status   | int4           | 撤销缴费的状态                                  | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text           | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text           | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text           | 删除人                                          | optional     | **system_user表**的user_id |

### refund_doc_veh_detail 退款单据车辆明细

营运车:

FEE_RET_BILL_DETAIL 退款单据车辆明细

| Name                       | Type           | Description        | **Required** | default                    |
| -------------------------- | -------------- | ------------------ | ------------ | -------------------------- |
| id                         | bigint         | 按指定方法生成     | required     | 主键                       |
| refund_doc_veh_detail_id   | text           | 退款单据id         | required     | 联合主键                   |
| payment_receipt_info_id    | text           | 退款单号           | optional     |                            |
| veh_id                     | text           | 车辆id             | optional     |                            |
| total_payment_amount       | numeric        | 退款总金额         | optional     |                            |
| service_month_fee          | numeric        | 服务月费           | optional     |                            |
| service_payment_start_date | timestamptz(6) | 服务费缴费起始日期 | optional     |                            |
| service_payment_deadline   | timestamptz(6) | 服务费缴费截止日期 | optional     |                            |
| service_fee_years          | int4           | 服务费年数         | optional     |                            |
| service_fee_months         | int4           | 服务费月数         | optional     |                            |
| service_total_fee          | numeric        | 服务费用总额       | optional     |                            |
| expire_date                | timestamptz(6) | 报废日期           | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间           | required     |                            |
| created_by                 | text           | 创建人             | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间           | optional     |                            |
| updated_by                 | text           | 修改人             | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间           | optional     |                            |
| deleted_by                 | text           | 删除人             | optional     | **system_user表**的user_id |

### refund_doc_oper_record 退款单据操作记录

营运车:

FEE_RET_BILL_LOG 退款单据操作记录

| Name                      | Type        | Description                                     | **Required** | default                    |
| ------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | bigint      | 按指定方法生成                                  | required     | 主键                       |
| refund_doc_oper_record_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| payment_receipt_info_id   | text        | 退款单号                                        | optional     |                            |
| step                      | int4        | 进度                                            | optional     |                            |
| is_pass                   | boolean     | 是否验证通过                                    | optional     |                            |
| remarks                   | text        | 备注                                            | optional     |                            |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                | text        | 删除人                                          | optional     |                            |

### alarm_detection_line 交通局检测线报警

营运车:

REPORT_REGION_ALARM_DATA 交通局检测线报警

| Name                    | Type           | Description                                          | **Required** | default                            |
| ----------------------- | -------------- | ---------------------------------------------------- | ------------ | ---------------------------------- |
| id                      | int8           | 按指定方法生成                                       | required     | 主键                               |
| alarm_detection_line_id | text           | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                           |
| vehicle_id              | text           | 车辆ID                                               | optional     | 引用**vehicle_info**表的vehicle_id |
| alarm_type              | text           | 报警类型(营运车共60种报警)                           | optional     | **报警类型**字典                   |
| detection_start_time    | timestamptz(6) | 检测开始时间                                         | optional     |                                    |
| detection_end_time      | timestamptz(6) | 检测离开时间                                         | optional     |                                    |
| processing_description  | text           | 处理描述                                             | optional     |                                    |
| processor               | text           | 处理人                                               | optional     | **system_user表**的user_id         |
| processing_status       | text           | 处理状态                                             | optional     | **处警处理状态**字典               |
| processing_time         | timestamptz(6) | 处理时间                                             | optional     |                                    |
| processing_method       | text           | 处理方式                                             | optional     | **处警处理方式**字典               |
| tachograph_speed        | numeric        | 行驶记录仪速度                                       | optional     |                                    |
| coordinate              | point          | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                                    |
| location_description    | text           | 位置描述                                             | optional     |                                    |
| road_grade              | text           | 道路等级                                             | optional     | **道路等级**字典                   |
| gps_speed               | numeric        | GPS速度                                              | optional     |                                    |
| maximum_speed           | numeric        | 最高速度                                             | optional     |                                    |
| duration                | text           | 持续时间                                             | optional     |                                    |
| alarm_end_position      | text           | 报警结束位置                                         | optional     |                                    |
| latest_alarm_time       | timestamptz(6) | 最新报警时间                                         | optional     |                                    |
| latest_alarm_position   | int4           | 最新报警位置                                         | optional     |                                    |
| is_alarm_effective      | boolean        | 报警是否有效                                         | optional     |                                    |
| is_alarm_over           | boolean        | 报警是否结束                                         | optional     |                                    |
| is_cancel_alarm         | boolean        | 是否取消报警                                         | optional     |                                    |
| alarm_source            | text           | 报警来源:1.终端报警 2.平台报警                       | optional     | **报警来源**字典                   |
| speed_limit_threshold   | numeric        | 限速阀值                                             | optional     |                                    |
| road_name               | text           | 道路名称                                             | optional     |                                    |
| area_id                 | text           | 进区域ID                                             | optional     |                                    |
| alarm_deal_id           | text           | 处警ID                                               | optional     |                                    |
| pid                     | text           | 地区                                                 | optional     |                                    |
| detection_line          | text           | 检测线                                               | optional     |                                    |
| is_retransmit           | boolean        | 是否补传                                             | optional     |                                    |
| created_at              | timestamptz(6) | 创建时间                                             | required     |                                    |
| created_by              | text           | 创建人                                               | required     | **system_user表**的user_id         |
| updated_at              | timestamptz(6) | 修改时间                                             | optional     |                                    |
| updated_by              | text           | 修改人                                               | optional     | **system_user表**的user_id         |
| deleted_at              | timestamptz(6) | 删除时间                                             | optional     |                                    |
| deleted_by              | text           | 删除人                                               | optional     | **system_user表**的user_id


## 重点车相关数据表

### major_picture 图片表

**根据以下数据表整合而成**

重点车:

VD_PICTURE 图片表

| Name             | Type        | Description                                     | **Required** | default                    |
| ---------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id               | bigint      | 按指定方法生成                                  | required     | 主键                       |
| major_picture_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| pic_num          | int4        | 图片序号                                        | optional     |                            |
| pic_date         | text        | 图片                                            | optional     |                            |
| is_synchronized  | boolean     | 是否同步                                        | optional     |                            |
| is_deleted       | boolean     | 是否删除                                        | optional     |                            |
| created_at       | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by       | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at       | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by       | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at       | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by       | text        | 删除人                                          | optional     | **system_user表**的user_id |

### driver_card_apply 司机卡申请

**根据以下数据表整合而成**

重点车:

MON_DRIVERCARD_ALI 司机卡申请

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| driver_card_apply_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| driver_id            | text        | 驾驶员ID                                        | optional     |                            |
| ent_id               | text        | 驾驶员申请企业id                                | optional     |                            |
| auditor              | text        | 审核人                                          | optional     |                            |
| remarks              | text        | 备注                                            | optional     |                            |
| is_claimed           | boolean     | 是否已申领                                      | optional     |                            |
| is_audit             | boolean     | 是否审核                                        | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### pub_security_veh 公安内网六合一平台同步车辆表

**根据以下数据表整合而成**

重点车:

JJ_VEHICLE 公安内网六合一平台同步车辆表

| Name                          | Type           | Description    | **Required** | default                    |
| ----------------------------- | -------------- | -------------- | ------------ | -------------------------- |
| id                            | int8           | 按指定方法生成 | required     | 主键                       |
| license_plate_number          | text           | 车牌号码       | required     |                            |
| license_plate_type            | text           | 号牌种类       | required     |                            |
| owner                         | text           | 所有人         | optional     |                            |
| use_nature                    | text           | 使用性质       | optional     |                            |
| total_mass                    | numeric        | 总质量         | optional     |                            |
| vehicle_type                  | text           | 车辆类型       | optional     |                            |
| valid_until                   | timestamptz(6) | 有效期至       | optional     |                            |
| mandatory_retire_date         | timestamptz(6) | 强制报废期止   | optional     |                            |
| vehicle_identification_number | text           | 车辆识别代号   | optional     |                            |
| district                      | text           | 所在县         | optional     |                            |
| contact_addr                  | text           | 联系地址       | optional     |                            |
| phone                         | text           | 固话           | optional     |                            |
| cellphone                     | text           | 联系电话       | optional     |                            |
| vehicle_state                 | int4           | 机动车状态     | optional     |                            |
| business_scope                | text           | 经营范围       | optional     |                            |
| initial_registration_date     | timestamptz(6) | 初次登记日期   | optional     |                            |
| registration_date             | timestamptz(6) | 登记日期       | optional     |                            |
| distribute_date               | timestamptz(6) | 发牌日期       | optional     |                            |
| model                         | text           | 型号           | optional     |                            |
| is_deleted                    | numeric        | 是否删除       | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间       | required     |                            |
| created_by                    | text           | 创建人         | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间       | optional     |                            |
| updated_by                    | text           | 修改人         | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间       | optional     |                            |
| deleted_by                    | text           | 删除人         | optional     | **system_user表**的user_id |

### veh_network_info 网约车车辆信息

**根据以下数据表整合而成**

重点车:

JJ_WYCDA 网约车车辆信息

| Name                         | Type        | Description                                     | **Required** | default                    |
| ---------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                           | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_network_info_id          | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| access_network_platform      | text        | 接入网约平台                                    | optional     |                            |
| license_plate_number         | text        | 车牌号                                          | optional     |                            |
| license_plate_color          | integer     | 车牌颜色                                        | optional     | **车牌颜色**字典           |
| vehicle_type                 | integer     | 车辆类型                                        | optional     | **车辆类型**字典           |
| business_scope               | integer     | 经营范围                                        | optional     |                            |
| transport_license_number     | text        | 运输证号                                        | optional     |                            |
| issue_date                   | timestamptz(6) | 发证日期                                        | optional     |                            |
| transport_license_valid_date | timestamptz(6) | 运输证有效日期                                  | optional     |                            |
| annual_inspection_date       | timestamptz(6) | 年审日期                                        | optional     |                            |
| annual_inspection_valid_date | timestamptz(6) | 年审有效期                                      | optional     |                            |
| created_at                   | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                   | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                   | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                   | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                   | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                   | text        | 删除人                                          | optional     | **system_user表**的user_id |

### driver_photo_comparison 驾驶员拍照比对

**根据以下数据表整合而成**

重点车:

MON_FINGER_PHOTO_CONTRAST  驾驶员拍照比对

| Name                       | Type        | Description                                          | **Required** | default                    |
| -------------------------- | ----------- | ---------------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                       | required     | 主键                       |
| driver_photo_comparison_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用      | required     | 联合主键                   |
| comparison_time            | timestamptz(6) | 比对时间                                             | optional     |                            |
| license_plate_number       | text        | 车牌号                                               | optional     |                            |
| sim_num                    | text        | SIM卡号                                              | optional     |                            |
| veh_id                     | text        | 车辆ID                                               | optional     |                            |
| terminal_id                | text        | 终端ID                                               | optional     |                            |
| driver_id                  | text        | 驾驶员ID                                             | optional     |                            |
| approval_cert_id           | text        | 核准证ID                                             | optional     |                            |
| coordinate                 | point       | 空间数据类型point表示经度(longitude)和纬度(latitude) | optional     |                            |
| pic_id                     | text        | 图片ID                                               | optional     |                            |
| pic_upload_time            | timestamptz(6) | 图片上传时间                                         | optional     |                            |
| is_legal                   | boolean     | 是否合法                                             | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                             | required     |                            |
| created_by                 | text        | 创建人                                               | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                             | optional     |                            |
| updated_by                 | text        | 修改人                                               | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                             | optional     |                            |
| deleted_by                 | text        | 删除人                                               | optional     | **system_user表**的user_id |

### driving_log_statistics 行车日志统计

**根据以下数据表整合而成**

重点车:

VD_DRIVING_LOG_REG 行车日志统计

| Name                      | Type        | Description                                     | **Required** | default                    |
| ------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                        | bigint      | 按指定方法生成                                  | required     | 主键                       |
| driving_log_statistics_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                    | text        | 车辆ID                                          | optional     |                            |
| driving_log_id            | text        | 行车日志ID                                      | optional     |                            |
| check_time                | timestamptz(6) | 校验时间                                        | optional     |                            |
| is_register               | boolean     | 是否登记(true登记false未登记)                   | optional     |                            |
| created_at                | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                | text        | 删除人                                          | optional     | **system_user表**的user_id |

### driving_log_audit  行车日志审核

**根据以下数据表整合而成**

重点车:

VD_DRIVING_LOG_DETAIL 行车日志审核日志

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| driving_log_audit_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| driving_log_id       | text        | 行车日志ID                                      | optional     |                            |
| audit_status         | text        | 审核状态                                        | optional     |                            |
| auditor              | text        | 审核人                                          | optional     |                            |
| audit_time           | timestamptz(6) | 审核时间                                        | optional     |                            |
| remarks              | text        | 备注                                            | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### violation_report 违法情况上报

**根据以下数据表整合而成**

重点车:

INFO_WFEVIDENCE 违法情况上报

| Name                | Type        | Description                                     | **Required** | default                    |
| ------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                  | bigint      | 按指定方法生成                                  | required     | 主键                       |
| violation_report_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id              | text        | 车辆ID                                          | optional     |                            |
| driver_name         | text        | 驾驶员姓名                                      | optional     |                            |
| driver_id_num       | text        | 驾驶员身份证号                                  | optional     |                            |
| file_id             | text        | 文件表ID                                        | optional     |                            |
| violation_desc      | text        | 违法描述                                        | optional     |                            |
| org_type            | text        | 机构类型                                        | optional     |                            |
| cause_type          | text        | 原由字典                                        | optional     |                            |
| created_at          | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by          | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at          | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by          | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at          | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by          | text        | 删除人                                          | optional     | **system_user表**的user_id |

### violation_report_pic 违法情况上报图片表

**根据以下数据表整合而成**

重点车:

INFO_WFEVIDENCE_FILE 违法情况上报图片表

| Name                    | Type        | Description                                     | **Required** | default                    |
| ----------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                      | bigint      | 按指定方法生成                                  | required     | 主键                       |
| violation_report_pic_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| serial_num              | text        | 序号                                            | optional     |                            |
| img_addr                | text        | 图片地址                                        | optional     |                            |
| created_at              | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by              | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at              | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by              | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at              | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by              | text        | 删除人                                          | optional     | **system_user表**的user_id |

### multiply_type_violation 十大类违法类型表

**根据以下数据表整合而成**

重点车:

VIO_CODEWFDM_VAR 十大类违法类型表

| Name                       | Type        | Description                                     | **Required** | default                    |
| -------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                  | required     | 主键                       |
| multiply_type_violation_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| violation_code             | text        | 违法代码                                        | optional     |                            |
| violation_category         | text        | 违法类别                                        | optional     |                            |
| violation_type             | text        | 违法类型                                        | optional     |                            |
| multiply_violation_type    | text        | 违法十大类类别                                  | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text        | 删除人                                          | optional     | **system_user表**的user_id |

### illegal_handling_agency 违法处理机构

**根据以下数据表整合而成**

重点车:

INFO_CLJG_AREA 违法处理机构

| Name                       | Type        | Description                                     | **Required** | default                    |
| -------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                  | required     | 主键                       |
| illegal_handling_agency_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| area_id                    | text        | 区域ID                                          | optional     |                            |
| operation_agency           | text        | 处理机关                                        | optional     |                            |
| operator                   | text        | 操作员                                          | optional     |                            |
| superior_area              | text        | 上级区域                                        | optional     |                            |
| category                   | text        | 类别                                            | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text        | 删除人                                          | optional     | **system_user表**的user_id |

### regional_violation_register  区域处理机关交通违法违规登记表

**根据以下数据表整合而成**

重点车:

WF_DEAL_REGEIT 区域处理机关交通违法违规登记表

| Name                           | Type        | Description                                                  | **Required** | default                                            |
| ------------------------------ | ----------- | ------------------------------------------------------------ | ------------ | -------------------------------------------------- |
| id                             | bigint      | 按指定方法生成                                               | required     | 主键                                               |
| regional_violation_register_id | text        | 区域处理机关交通违法违规登记表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                                           |
| violation_detail_id            | text        | 违法记录表ID                                                 | optional     | **vehicle_violation_details**的violation_detail_id |
| vehicle_id                     | text        | 违章车辆id                                                   | optional     | **vehicle_info表**的vehicle_id                     |
| driver_id                      | text        | 驾驶员id                                                     | optional     | **driver_info 表**的driver_id                      |
| illegal_code                   | integer     | 违法代码                                                     | optional     | **VIO_CODEWFDM** 违法描述字典表                    |
| illegal_time                   | timestamptz(6) | 违法时间                                                     | optional     |                                                    |
| illegal_type                   | integer     | 类型(1车辆2驾驶员)                                           | optional     | **违法类型**字典表                                 |
| register_time                  | timestamptz(6) | 登记时间                                                     | optional     |                                                    |
| processing_agency              | text        | 处理机关                                                     | optional     |                                                    |
| operator                       | text        | 操作员                                                       | optional     | **system_user表**的user_id                         |
| is_register                    | boolean     | 类型(false未登记true已登记)                                  | optional     | fase                                               |
| created_at                     | timestamptz(6) | 创建时间                                                     | required     |                                                    |
| created_by                     | text        | 创建人                                                       | required     | **system_user表**的user_id                         |
| updated_at                     | timestamptz(6) | 修改时间                                                     | optional     |                                                    |
| updated_by                     | text        | 修改人                                                       | optional     | **system_user表**的user_id                         |
| deleted_at                     | timestamptz(6) | 删除时间                                                     | optional     |                                                    |
| deleted_by                     | text        | 删除人                                                       | optional     | **system_user表**的user_id                         |
| is_deleted                     | boolean     | 是否删除                                                     | optional     |                                                    |

### veh_network_violation_code 网约车违法信息代码        

**根据以下数据表整合而成**

重点车:

WYC_WF_CODE  网约车违法信息代码

| Name                          | Type        | Description                                     | **Required** | default                    |
| ----------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                            | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_network_violation_code_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| category                      | text        | 类别                                            | optional     |                            |
| violation_code                | text        | 违法代码                                        | optional     |                            |
| violation_description         | text        | 违法描述                                        | optional     |                            |
| created_at                    | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                    | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                    | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                    | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                    | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                    | text        | 删除人                                          | optional     | **system_user表**的user_id |

### veh_spot_check_disposal 车辆抽查处置表

**根据以下数据表整合而成**

重点车:

INFO_VEH_HANDLE 车辆抽查处置表

| Name                       | Type        | Description                                     | **Required** | default                    |
| -------------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                         | bigint      | 按指定方法生成                                  | required     | 主键                       |
| veh_spot_check_disposal_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| veh_id                     | text        | 车辆ID                                          | optional     |                            |
| ent_id                     | text        | 企业                                            | optional     |                            |
| phone_num                  | text        | 手机号码                                        | optional     |                            |
| content                    | text        | 内容                                            | optional     |                            |
| send_time                  | timestamptz(6) | 发送时间                                        | optional     |                            |
| user_id                    | text        | 用户ID                                          | optional     |                            |
| notify_content             | text        | 通报内容                                        | optional     |                            |
| announce_content           | text        | 语音内容                                        | optional     |                            |
| app_push_content           | text        | APP推送内容                                     | optional     |                            |
| spot_check_id              | timestamptz(6) | 抽查表ID                                        | optional     |                            |
| disposal_method            | int4        | 处置方式                                        | optional     |                            |
| regis_time                 | timestamptz(6) | 登记时间                                        | optional     |                            |
| is_app_push                | boolean     | 是否APP推送                                     | optional     |                            |
| is_announce                | boolean     | 是否语音通知                                    | optional     |                            |
| is_notify                  | boolean     | 是否通报                                        | optional     |                            |
| is_sms_push                | boolean     | 是否短信推送                                    | optional     |                            |
| is_deleted                 | boolean     | 是否删除                                        | optional     |                            |
| created_at                 | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by                 | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at                 | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by                 | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at                 | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by                 | text        | 删除人                                          | optional     | **system_user表**的user_id |

### deduction_detail 扣车明细表

**根据以下数据表整合而成**

重点车:

INFO_VEH_DETAIN 扣车明细表

| Name                 | Type        | Description                                     | **Required** | default                    |
| -------------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                  | required     | 主键                       |
| deduction_detail_id  | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| license_plate_number | text        | 车牌号码                                        | optional     |                            |
| license_plate_color  | integer     | 车牌颜色                                        | optional     |                            |
| deduction_pic        | text        | 扣车图片                                        | optional     |                            |
| deduction_time       | timestamptz(6) | 扣车时间                                        | optional     |                            |
| release_pic          | text        | 放车图片                                        | optional     |                            |
| release_time         | timestamptz(6) | 放车时间                                        | optional     |                            |
| status               | text        | 状态                                            | optional     |                            |
| bayonet_pic          | text        | 卡口图片                                        | optional     |                            |
| regis_institution    | text        | 登记所在机构                                    | optional     |                            |
| is_catalog_library   | boolean     | 是否目录库                                      | optional     |                            |
| remarks              | text        | 备注                                            | optional     |                            |
| is_violation_judge   | boolean     | 违法判断(是否通过)                              | optional     |                            |
| is_satellite_judge   | boolean     | 卫星定位判断(是否通过)                          | optional     |                            |
| is_bayonet_judge     | boolean     | 卡口判断(是否通过)                              | optional     |                            |
| is_deleted           | boolean     | 是否删除                                        | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by           | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by           | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by           | text        | 删除人                                          | optional     | **system_user表**的user_id |

### deduction_pic 扣车图片

**根据以下数据表整合而成**

重点车:

DETAIN_PICTURE 扣车图片

| Name             | Type        | Description                                     | **Required** | default                    |
| ---------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id               | bigint      | 按指定方法生成                                  | required     | 主键                       |
| deduction_pic_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| serial_num       | text        | 序号                                            | optional     |                            |
| img_url          | text        | 图片地址                                        | optional     |                            |
| is_deleted       | boolean     | 是否删除                                        | optional     |                            |
| created_at       | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by       | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at       | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by       | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at       | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by       | text        | 删除人                                          | optional     | **system_user表**的user_id |

### release_pic 放行图片

**根据以下数据表整合而成**

重点车:

NEWZTC_PICTURE 放行图片

| Name           | Type        | Description                                     | **Required** | default                    |
| -------------- | ----------- | ----------------------------------------------- | ------------ | -------------------------- |
| id             | bigint      | 按指定方法生成                                  | required     | 主键                       |
| release_pic_id | text        | 外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| serial_num     | text        | 序号                                            | optional     |                            |
| img_url        | text        | 图片地址                                        | optional     |                            |
| is_deleted     | boolean     | 是否删除                                        | optional     |                            |
| created_at     | timestamptz(6) | 创建时间                                        | required     |                            |
| created_by     | text        | 创建人                                          | required     | **system_user表**的user_id |
| updated_at     | timestamptz(6) | 修改时间                                        | optional     |                            |
| updated_by     | text        | 修改人                                          | optional     | **system_user表**的user_id |
| deleted_at     | timestamptz(6) | 删除时间                                        | optional     |                            |
| deleted_by     | text        | 删除人                                          | optional     | **system_user表**的user_id |

### app_enforcement APP现场执法表

**根据以下数据表整合而成**

重点车:

INFO_MONITOR APP现场执法表

| Name                 | Type        | Description                                                  | **Required** | default                    |
| -------------------- | ----------- | ------------------------------------------------------------ | ------------ | -------------------------- |
| id                   | bigint      | 按指定方法生成                                               | required     | 主键                       |
| app_enforcement_id   | text        | APP现场执法表外部编码，由golang程序生成的xid，暴露到外部使用 | required     | 联合主键                   |
| vehicle_id           | text        | **vehicle_info** 车辆信息表 的vehicle_id                     | optional     |                            |
| operation_user       | text        | 操作用户                                                     | optional     | **system_user表**的user_id |
| edit_text            | text        | 编辑文本                                                     | optional     |                            |
| coordinate           | point       | 空间数据类型point表示经度(longitude)和纬度(latitude)         | optional     |                            |
| location_description | text        | 位置描述                                                     | optional     |                            |
| enterprise_type      | text        | 企业类型                                                     | optional     | **企业类型**字典           |
| picket_status        | integer     | 纠察状态（1.反馈辖区管理 2.执法考评 3.其他 4.查处“两非”渣土车） | optional     |                            |
| created_at           | timestamptz(6) | 创建时间                                                     | required     |                            |
| created_by           | text        | 创建人                                                       | required     | **system_user表**的user_id |
| updated_at           | timestamptz(6) | 修改时间                                                     | optional     |                            |
| updated_by           | text        | 修改人                                                       | optional     | **system_user表**的user_id |
| deleted_at           | timestamptz(6) | 删除时间                                                     | optional     |                            |
| deleted_by           | text        | 删除人                                                       | optional     | **system_user表**的user_id |