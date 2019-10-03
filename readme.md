# beego Blog Test

> beego + vue-element-admin 实现的简单博客，基本的页面展示跟CURD
> https://gvb.hezhizheng.com/

## 安装
```
git clone 
cd beego_vb/conf
cp app.conf.example app.conf
```

### 后端
- api文件
  - controllers
    - api

- sql文件(数据库结构)
  - sql

#### 运行/开发
```
cd beego_vb

// 安装依赖
go mod download

// 运行
bee run / go run main.go
```

### 前端
- 前端vue（后台）默认访问地址 http://127.0.0.1:8080/vue-admin
  - frontend
    - govb-vue-admin

#### 运行/开发/打包
```
cd beego_vb/frontend/govb-vue-admin

// 安装依赖
npm run install

// 开发
npn run dev

// 打包，打包之后的文件放在 beego_vb/static/vue-admin，如修改过打包输出目录，需重新定义静态文件地址
npn run build
```

## 发布
参考 [Golang在Linux下的部署](https://hezhizheng.com/blog/deployment-of-golang-under-linux)