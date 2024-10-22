# Grapic

🌟 开源项目介绍 🌟
<br />
Grapic 近期计划开源工具，是提供一个私密且便捷的平台，用于分享代码和文本。并且提供一个直观、易用的界面，让用户能够轻松地分享信息。

🛠️ 技术栈
Grapic 后端采用以下技术构建：

* Go语言：使用Go语言编写，确保了后端服务的性能和并发处理能力。
* Gin框架：轻量级的Web框架，为开发提供快速且灵活的路由处理。
* Xorm：一个简单而强大的ORM库，用于处理数据库操作，简化开发流程。
* SQLite：轻量级的数据库，无需复杂的配置，适用于快速开发和部署。

前端则采用现代的JavaScript技术栈：
* Vue 3：使用最新的Vue.js版本，提供更好的性能和更易维护的代码结构。
* Vuetify：基于Material Design的UI库，为应用提供美观、响应式的界面。
* Pinia：作为Vuex的替代品，提供更简洁的状态管理解决方案。
* Vue Router：官方的路由管理器，为单页应用提供强大的路由功能。
* Axios：一个基于Promise的HTTP客户端，用于浏览器和node.js，简化了与后端的通信。


🚀 功能特点
* 端到端加密：如果启用密码验证，则只有拥有正确密码的用户才能访问。
* 自销毁链接：分享的链接在一定时间后或一定次数的访问后自动销毁。
* 简洁的界面：易于使用的界面设计，让用户快速上手，无需复杂的学习曲线。
* 高度可定制：用户可以根据自己的需求定制分享的选项，如过期时间、访问次数等。

📝 使用场景
* 开发者：分享代码片段或配置文件，无需担心敏感信息泄露。
* 团队协作：在团队成员之间安全地共享文档和笔记。
* 隐私保护：为需要保密的信息提供一个安全的分享渠道。

💻 关于项目
* 项目依然在开发中，虽说已经可以开始使用，但是功能方面依然不够完善，并且没有做足够的测试，当前阶段也正在逐步完善中。


# 配置结构
* 以下是与二进制程序同级的目录: default的结构
```text
.
├── config
│   └── config.json
├── dist
│   ├── assets
│   │   ├── xxxxxx.js
│   │   ├── xxxxxx.css
│   ├── index.html
│   └── vite.svg
├── logs
├── sqlite
│   └── default.db
└── uploads
```

* config/config.json: 配置文件，用于配置程序的运行参数，如端口号、数据库连接信息等。
* dist: 前端静态资源，包括js、css文件。
* sqlite/default.db: 数据库文件，用于存储数据。

其中比较重要的是config.json文件，它包含了程序的运行参数，包括端口号、地址、启用SSL等，具体介绍如下：

```json
{
    "address": "", // URL地址，例如: baidu.com
    "http": {
        "port": "" // HTTP端口号, 默认9099
    },
    "https": {
        "enabled": false,   // 是否启用SSL，默认false
        "port": "",         // HTTPS端口号，默认443
        "cert_file": "",    // SSL证书文件路径
        "key_file": ""      // SSL密钥文件路径
    }
}
```

如果需要更改配置，可以手动修改后重启二进制程序，或者删除config.json文件，并且重新运行程序，则会要求重新交互式配置。

# 运行程序

脚本简单写了一下，大概长这样，对于linux环境，可以直接运行，另外，如果配置要重新生成的话，可以删除config.json文件，然后运行程序，会要求重新配置，或者手动编辑config.json文件。

```bash
{
    cd /opt/
    git clone git@github.com:hyhacct/grapic.git

    # 编译
    cd /opt/grapic/go-admin
    go build -o grapic cmd/main.go

    cd /opt/grapic/vite-project
    yarn
    yarn build

    mkdir -p /opt/grapic/data/default
    mv /opt/grapic/go-admin/grapic /opt/grapic/data/
    mv /opt/grapic/vite-project/dist /opt/grapic/data/default/

    # 启动
    cd /opt/grapic/data/
    chmod +x grapic
    ls
}
```