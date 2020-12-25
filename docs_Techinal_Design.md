## User permission (Version 2.0)

### History

| Version | Date time  | Description          | Author |
|---------|------------|----------------------|--------|
| 1.0     | 10/10/2019 | Version đầu tiên     | SonNH  |
| 2.0     | 24/06/2020 | Thêm phân quyền data | SonNH  |

### Mục lục

  * [Giới thiệu](#gi-i-thi-u)
  * [Thiết kế](#thi-t-k-)
    + [Tổng quan về hệ thống](#t-ng-quan-v--h--th-ng)
    + [Các iFrame](#c-c-iframe)
    + [Các API](#c-c-api)
    + [Database](#database)

### Giới thiệu

* Tài liệu này mô tả toàn bộ thiết kế của hệ thống phân quyền người dùng. Đây là hệ thống được tích hợp vào các hệ thống business khác, giúp phân quyền người dùng trên các app một cách hiệu quả và bảo mật

### Thiết kế

#### Tổng quan về hệ thống

![](User_Permission_System_Architecture.png)

* Hệ thống bao gồm các thành phần chính như sau:

  * **1. Website**: tương tác trực tiếp với người dùng. Các chức năng chính của website là:

    * Tạo 1 app integrate trực tiếp với hệ thống User Permission

    * Thêm các quyền, nhóm quyền, quyền truy cập dữ liệu cho 1 app

    * Thêm user, phân quyền cho user

  * **2. Iframe**: các màn hình được thiết kế theo dạng iframe, cho phép các trang web khác có thể nhúng các màn hình vào chính trang web của mình. Các màn hình được phép nhúng vào:

    * Quản lý người dùng

    * Quản lý quyền

    * Quản lý nhóm quyền

    * Quản lý quyền truy cập dữ liệu

    * Quản lý nhóm quyền truy cập dữ liệu

  * **3. Các API**: giúp các app khác truy cập trực tiếp vào User Permission qua REST API để thực hiện việc phân quyền truy cập trên app, ai có quyền, ai không có quyền vào chức năng nào trên app.

  * **4. Lưu trữ**: Hệ thống dùng database Postgres làm CSDL chính để lưu trữ và xử lý dữ liệu

* Kết nối với các hệ thống khác:

  * SSO: để đăng nhập vào website

  * HRIS API: để lấy thông tin nhân viên như MSNV, Tên, Chức Vụ từ HRM

#### Các iFrame

* 1. Quản lý người dùng

  ```url
  https://stg-te-acl.scommerce.asia/users/<app_id>
  ```

* 2. Quản lý quyền

  ```url
  https://stg-te-acl.scommerce.asia/permission/<app_id>
  ```

* 3. Quản lý nhóm quyền

  ```url
  https://stg-te-acl.scommerce.asia/rolepermission/<app_id>
  ```

* 4. Quản lý quyền truy cập dữ liệu

  ```url
  https://stg-te-acl.scommerce.asia/data/<app_id>
  ```

* 5. Quản lý nhóm quyền truy cập dữ liệu

  ```url
  https://stg-te-acl.scommerce.asia/functiondata/<app_id>
  ```

* _Lưu ý_: <app_id> được cung cấp bởi hệ thống User Permission

#### Các API

* 1. Lấy danh sách user có trong hệ thống

  * URL: `https://stg-te-acl.scommerce.asia/api/users/<app_id>?<option_query>`

  * Phương thức: `GET`

  * <app_id>: mã định danh của app, được cung cấp bởi hệ thống User Permission

  * <option_query>: tham số filter, nếu không dùng thì mặc định sẽ lấy tất cả nhân viên có trong hệ thống của app đó:

    * permissions: danh sách quyền muốn filter, mỗi quyền cách nhau bởi dấu phẩy (,)

    * datas: danh sách quyền truy cập dữ liệu muốn filter, mỗi quyền cách nhau bởi dấu phẩy (,)

  * Ví dụ:

    * Request

    ```bash
    curl https://stg-te-acl.scommerce.asia/api/users/b9edab79-af29-4bdc-9fba-122f085d2eba?permissions=BAO_CAO_TRUY_THU_THEO_DIA_DIEM,BAO_CAO_TRUY_THU_THEO_QUAN_LY&datas=TAT_CA,THEO_NGUOI_BAO_CAO
    ```

    * Response

    ```json
    {
        "status": 200,
        "msg": "Ok",
        "data": [
            {
                "employee_id": 27198,
                "name": "Ngô Huy Long",
                "id": "6de4ae6e-1959-4993-bcb9-04c646d9caba",
                "department": "Phòng Kiểm Soát Nội Bộ",
                "job_title": "Trưởng Nhóm Đền Bù"
            },
            {
                "employee_id": 1869742,
                "name": "Tạ Ngọc Bảo Lâm",
                "id": "b91ebad4-10d3-4c56-a4be-36f5ebd40dc0",
                "department": "Tech Enterprise",
                "job_title": "Chuyên Viên Phát Triển Phần Mềm Cấp Cao"
            }
        ]
    }
    ```

* 2. Lấy danh sách quyền của 1 user

  * URL: `https://stg-te-acl.scommerce.asia/api/userroles/<app_id>/<user_id>`

  * Phương thức: `GET`

  * <app_id>: mã định danh của app, được cung cấp bởi hệ thống User Permission

  * <user_id>: mã của user cần để lấy danh sách quyền

  * Ví dụ: 

    * Request

    ```bash
    curl https://stg-te-acl.scommerce.asia/api/users/api/userroles/b9edab79-af29-4bdc-9abf-122f085d2eba/199739
    ```

    * Response

    ```json
    {
        "status": 200,
        "msg": "Ok",
        "data": [
            {
                "user_name": "Huỳnh Minh Trọng",
                "permission_id": "4aa863b8-b3a9-418b-81c6-6af5a1a305c3",
                "data_permissions": {
                    "XUAT_DU_LIEU": [
                        {
                            "code": "KHONG",
                            "name": "Không"
                        }
                    ]
                },
                "permission_constant_name": "DANH_SACH_DON_AO",
                "role_name": "Area Manager",
                "permission_title": "Danh sách đơn ảo",
                "data_name": "Không",
                "user_id": "199739",
                "email": null,
                "info": null
            },
            {
                "user_name": "Huỳnh Minh Trọng",
                "permission_id": "0bec1e57-1f03-42d3-abed-b883e4b0c302",
                "data_permissions": null,
                "permission_constant_name": "TRUY_THU_DON_AO_NHAN_VIEN",
                "role_name": "Area Manager",
                "permission_title": "Thông tin truy thu đơn ảo NV",
                "data_name": null,
                "user_id": "199739",
                "email": null,
                "info": null
            }
        ]
    }
    ```

#### Database

![](User_Permission_Database.png)

* app: thông tin app được cấp quyền tích hợp với hệ thống User Permission

* users: thông tin user

* permission: thông tin quyền của app

* roles: thông tin nhóm quyền

* role_permissions: thông tin kết hợp để set quyền vào nhóm quyền

* user_roles: thiết lập nhân viên vào trong nhóm quyền

* data_permission: thông tin quyền truy cập dữ liệu

* data_roles: thộng tin nhóm quyền truy cập dữ liệu

* function_data: gắn nhóm quyền truy cập dữ liệu vào quyền cụ thể

* data_user_permission: thiết lập quyền truy cập dữ liệu cho nhân viên