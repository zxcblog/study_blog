# openapi 注释语法
## 公共文档信息
```protobuf
syntax = "proto3";

option go_package = "github.com/zxcblog/study-blog/pb;pb";

import "protoc-gen-openapiv2/options/annotations.proto";

// 在.proto文件的包声明之后，添加以下内容来定义Swagger文档的基础信息
option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  info: {
    title: "study-blog";
    version: "1.0.0";
    description: "学习笔记";
    contact: {
      name: "zxc";
      url: "https://github.com/zxcblog/study-blog";
      email: "zxc_7310@163.com";
    };
    license: {
      name: "Apache 2.0";
      url: "https://www.apache.org/licenses/LICENSE-2.0.html";
    };
  };
  host: "localhost:9090"
  base_path: "/gateway"
  external_docs: {
    url: "http://localhost:9090/gateway"
    description: "关于日记账V2项目的详细介绍"
  }
  schemes: HTTP
  schemes: HTTPS
  consumes: "application/json"
  produces: "application/json"
  // 认证相关配置
  security_definitions: {
    security: {
      key: "ApiKeyAuth"
      value: {
        type: TYPE_API_KEY
        in: IN_HEADER
        name: "x-Auth"
        extensions: {
          key: "x-amazon-apigateway-authorizer"
          value: {
            struct_value: {
              fields: {
                key: "type"
                value: {string_value: "token"}
              }
              fields: {
                key: "authorizerResultTtlInSeconds"
                value: {number_value: 60}
              }
            }
          }
        }
      }
    }
  }
  security: {
    security_requirement: {
      key: "ApiKeyAuth"
      value: {}
    }
  }
  // 设置统一返回信息
  responses: {
    key: "403"
    value: {description: "用户权限不够时返回错误信息"}
  }
  responses: {
    key: "404"
    value: {
      description: "找不到资源时返回错误信息"
      schema: {
        json_schema: {type: STRING}
      }
    }
  }
  responses: {
    key: "418"
    value: {
      description: "测试其他文件中的请求体"
      schema: {
        json_schema: {ref: ".base.v1.Error"}
      }
    }
  }

  // 定义自定义字段信息， swagger中不能查看
  extensions: {
    key: "x-grpc-gateway-foo"
    value: {string_value: "bar"}
  }
  extensions: {
    key: "x-grpc-gateway-baz-list"
    value: {
      list_value: {
        values: {string_value: "one"}
        values: {bool_value: true}
      }
    }
  }
};
```
- info: 定义Swagger文档的基础信息
- host: 定义请求的host
- base_path: 定义请求的基本前缀
- external_docs: 定义项目的外部文档连接
- schemes: 定义请求协议
- consumes: 定义请求格式， json, xml等等
- produces: 定义返回格式
- security_definitions: 定义认证信息
- security: 定义使用的认证信息
- responses: 定义统一返回信息
- extensions: 定义自定义字段信息

## 定义服务，接口，字段注释
```protobuf
syntax = "proto3";

package user.v1;

option go_package = "github.com/zxcblog/study-blog/user;user";

import "google/api/annotations.proto";
import "validate/validate.proto";
import "user/token.proto";
import "protoc-gen-openapiv2/options/annotations.proto";

// User 用户管理
service User {
  // 服务描述
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_tag) = {
    name: "User"
    description: "用户管理"
    external_docs: { //添加服务文档信息
      url: "http://localhost:9090/user.doc"
      description: "查看用户管理功能文档"
    }
  };

  rpc QQRegister(RegisterReq)returns(UserAuthRes){
    option (google.api.http) = {
      post: "/v1/user/register/qq",
      body: "*",
    };

    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "QQ注册"
      description: "使用QQ号进行注册，已废弃"
      operation_id: "QQRegister"
      deprecated: true // 添加废弃标记
      external_docs: { // 添加单个接口文档信息
        url: "http://localhost:9090/user.doc#qq"
        description: "QQ注册文档信息"
      }

      // 不需要认证
      security:{}

      // 设置请求头信息
      parameters: {
        headers: {
          name: "x-Use-Env",
          description: "使用的环境标识"
          type: STRING
        }
      }

      responses: {
        key: "500"
        value: {
          description: "服务器出现问题了"
          schema: {
            json_schema: {type: STRING}
          }
        }
      }
    };

  }

  rpc Register(RegisterReq)returns(UserAuthRes){
    option (google.api.http) = {
      post: "/v1/user/register",
      body: "*",
    };
  }

  // Login 用户登录
  rpc Login(LoginReq)returns(UserAuthRes){
    option (google.api.http) = {
      post: "/v1/user/login",
      body: "*",
    };
  }
}

// 注册请求
message RegisterReq {
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
    json_schema: {
      title: "RegisterReq";
      description: "注册请求";
      required: ["username","password","confirm_password"],
    }
    external_docs: {
      url: "http://localhost:9090/user.doc#RegisterReq";
      description: "注册请求文档";
    }
    // 请求示例
    example: "{\"username\": \"zhouXiaoChuan\", \"password\": \"123456\", \"confirm_password\":\"123456\", \"email\":\"zxc_7310@163.com\"}"
  };

  string password         = 3[(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) ={
    description: "密码",
    min_length: 6,
    max_length: 20,
    pattern: "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]+$",
  },(validate.rules).string = {min_len: 1}];// 密码


  string username         = 1[(validate.rules).string = {min_len: 1}]; // 用户名
  string email = 2[(validate.rules).string = {email: true}]; // 邮箱
  string confirm_password = 4; // 确认密码
  string img_cache        = 7; // 图形验证码
}

// 登录类型
enum LoginType {
  Detault_Login = 0; // 邮箱密码登录
  Mobile_Login = 1; // 手机验证码登录
}

// 登录请求
message LoginReq {
  LoginType type  = 1; // 登录类型
  string account  = 2 [(validate.rules).string = {min_len: 1}]; // 账号
  string password = 3; // 认证密码
  string captcha  = 4; // 验证码
  string captcha_id = 5; // 验证码id
}

// 登录注册返回信息
message UserAuthRes {
  user.v1.Auth token_info = 1; // token认证信息
}
```

- option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_tag): 定义服务描述信息
  - external_docs: 定义服务文档信息
- option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation): 定义接口描述信息
  - security: 定义该接口的认证信息， 如果swagger文档中定义，此处的优先级高，会使用此处的认证替换swagger文档中的认证
- option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema): 定义请求体公共描述信息
  - json_schema 定义请求体信息
  - external_docs: 定义请求体文档信息
  - example: 定义请求示例
- (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field): 定义单个字段描述信息
  - min_length, max_length: 定义字段长度限制
  - min_items, max_items: 定义数组长度限制
  - default: 定义默认值




















