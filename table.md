# 系统管理
### enterprise 企业 
| Name | Type | Description | Remark | 
| ---- | ---- | ---- | ---- | 
| id | int8 | ID |  | 
| enterprise_id | text | 企业ID |  | 
| superior_enterprise_id | text | 上级企业ID |  | 
| enterprise_code | text | 企业码 |  | 
| enterprise_name | text | 企业名称 |  | 
| enterprise_level | int4 | 企业级别 |  | 
| display_number | int4 | 显示顺序 |  | 
| contact_persons | _jsonb | 联系人 |  | 
| enterprise_address | text | 企业地址 |  | 
| business_scope | int4 | 经营范围 |  | 
| fax_number | text | 传真号码 |  | 
| province_id | int8 | 省份ID |  | 
| city_id | int8 | 城市ID |  | 
| district_id | int8 | 区域ID |  | 
| institution_category | int8 | 机构类别 |  | 
| operating_license_photo | text | 经营许可证图片 |  | 
| business_license_photo | text | 营业执照图片 |  | 
| business_license_issuance_date | timestamptz | 营业执照发证日期 |  | 
| business_license_expiry_date | timestamptz | 营业执照到期日期 |  | 
| enterprise_nature | int4 | 企业性质 |  | 
| legal_representative | text | 企业法人代表 |  | 
| legal_representative_phone | text | 企业法人代表联系电话 |  | 
| legal_representative_id_card_photo | text | 企业法人代表身份证-图片 |  | 
| legal_representative_id_card | text | 企业法人-身份证号码 |  | 
| entrusted_agent | text | 委托代理人 |  | 
| entrusted_agent_phone | text | 委托代理人-电话号码 |  | 
| entrusted_agent_id_card_photo | text | 委托代理人身份证图片 |  | 
| entrusted_agent_id_card | text | 委托代理人-身份证号码 |  | 
| organization_code | text | 组织机构代码(企业的营运证) |  | 
| organization_code_certificate_photo | text | 组织机构代码证照片 |  | 
| update_time_in | timestamptz | 内网更新时间 |  | 
| business_photo | text | 业务办理扫描件照片 |  | 
| is_black | bool | 是否黑名单 |  | 
| check_status | int4 | 审核状态 |  | 
| is_install | bool | 是否安装 |  | 
| is_input | bool | 是否录入完成 |  | 
| is_upload_province | bool | 是否上次省厅 |  | 
| score | int4 | 记分 |  | 
| brigade_id | text | 所属大队ID |  | 
| police_station_id | text | 所属派出所ID |  | 
| brigade_review_opinion | text | 大队审核意见 |  | 
| brigade_review_time | timestamptz | 大队审核时间 |  | 
| brigade_review_by | text | 大队审核人 |  | 
| association_review_opinion | text | 协会审核意见 |  | 
| association_review_time | timestamptz | 协会审核时间 |  | 
| association_review_by | timestamptz | 协会审核时间 |  | 
| create_at | timestamptz | 创建时间 |  | 
| create_by | text | 创建人 |  | 
| update_at | timestamptz | 修改时间 |  | 
| update_by | text | 修改人 |  | 
| delete_at | timestamptz | 删除时间 |  | 
| delete_by | text | 删除人 |  | 
| record_at | timestamptz | 登记时间 |  | 
| record_by | text | 登记人 |  | 
| remarks | text | 备注 |  | 
| is_deleted | bool | 是否删除 |  | 
### system_user 系统用户 
| Name | Type | Description | Remark | 
| ---- | ---- | ---- | ---- | 
| id | int8 | ID |  | 
| user_id | text | 用户ID |  | 
| password | text | 密码 |  | 
| uername | text | 用户名 |  | 
| enterprise_id | text | 企业ID |  | 
| department_id | text | 部门ID |  | 
| grade | int4 | 级别 |  | 
| audit_level | int4 | 审核等级 |  | 
| user_type | int4 | 用户类型 |  | 
| is_valid | bool | 是否有效 |  | 
| user_state | int4 | 状态 |  | 
| ip_address | text | ip地址 |  | 
| is_bind_ip | bool | 是否绑定IP |  | 
| email | text | 邮箱 |  | 
| telephone | text | 电话号码 |  | 
| mobile | text | 手机号码 |  | 
| ukey | text | 加密串码 |  | 
| mkey | text | 手机串号 |  | 
| app_version | text | 客户端版本号 |  | 
| created_at | timestamptz | 创建时间 |  | 
| create_by | text | 创建人 |  | 
| update_at | timestamptz | 修改时间 |  | 
| update_by | text | 修改人 |  | 
| delete_at | timestamptz | 删除时间 |  | 
| delete_by | text | 删除人 |  | 
| is_delete | bool | 是否删除 |  | 
| remarks | text | 备注 |  | 
### department 部门 
| Name | Type | Description | Remark | 
| ---- | ---- | ---- | ---- | 
| id | int8 | ID |  | 
| department_id | text | 部门ID |  | 
| enterprise_id | text | 企业ID |  | 
| superior_department_id | text | 上级部门ID |  | 
| department_name | text | 部门名称 |  | 
| department_code | text | 部门编号 |  | 
| department_category | int4 | 部门类型 |  | 
| internal_number | int4 | 排序 |  | 
| create_at | timestamptz | 创建时间 |  | 
| create_by | text | 创建人 |  | 
| update_at | timestamptz | 修改时间 |  | 
| update_by | text | 修改人 |  | 
| delete_at | timestamptz | 删除时间 |  | 
| delete_by | text | 删除人 |  | 
| is_delete | bool | 是否删除 |  | 
| remarks | text | 备注 |  | 
# 位置管理
### vehicle_location_last 车辆最新位置表 
| Name | Type | Description | Remark | 
| ---- | ---- | ---- | ---- | 
| id | int8 | ID |  | 
| vehicle_id | text | 车辆ID |  | 
| driver_id | text | 驾驶员ID |  | 
| enterprise_id | text | 企业ID |  | 
| supervision_photo_id | text | 监控图片ID |  | 
| acceleration | text | 加速度 |  | 
| alarm_content | text | 报警内容 |  | 
| alititude | text | 海拔 |  | 
| direction | text | 方向 |  | 
| tachograph_speed | numeric | 行驶记录仪速度 |  | 
| is_locate | bool | 是否定位 |  | 
| locate_time | timestamptz | 定位时间 |  | 
| coordinate | point | 坐标 |  | 
| location_description | text | 位置描述 |  | 
| mileage | text | 里程 |  | 
| gps_speed | numeric | GPS速度 |  | 
| star_count | text | 星数 |  | 
| star_status | text | 星况 |  | 
| road_name | text | 道路名称 |  | 
| speed_limit_threshold | numeric | 限速阀值 |  | 
| correct_coordinate | point | 纠偏后坐标 |  | 
| vehicle_status | text | 车辆状态 |  | 
| imei | text | 终端IMEI |  | 
| sim_number | text | SIM卡号 |  | 
| district_id | text | 区域ID |  | 
### vehicle_location_his 车辆位置历史表 
| Name | Type | Description | Remark | 
| ---- | ---- | ---- | ---- | 
| id | int8 | ID |  | 
| vehicle_id | text | 车辆ID |  | 
| driver_id | text | 驾驶员ID |  | 
| enterprise_id | text | 企业ID |  | 
| supervision_photo_id | text | 监控图片ID |  | 
| acceleration | text | 加速度 |  | 
| alarm_content | text | 报警内容 |  | 
| alititude | text | 海拔 |  | 
| direction | text | 方向 |  | 
| tachograph_speed | numeric | 行驶记录仪速度 |  | 
| is_locate | bool | 是否定位 |  | 
| locate_time | timestamptz | 定位时间 |  | 
| coordinate | point | 坐标 |  | 
| location_description | text | 位置描述 |  | 
| mileage | text | 里程 |  | 
| gps_speed | numeric | GPS速度 |  | 
| star_count | text | 星数 |  | 
| star_status | text | 星况 |  | 
| road_name | text | 道路名称 |  | 
| speed_limit_threshold | numeric | 限速阀值 |  | 
| correct_coordinate | point | 纠偏后坐标 |  | 
| vehicle_status | text | 车辆状态 |  | 
| imei | text | 终端IMEI |  | 
| sim_number | text | SIM卡号 |  | 
| district_id | text | 区域ID |  | 
