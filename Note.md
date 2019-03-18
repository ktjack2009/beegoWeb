# Beego项目配置
## 系统默认配置
    beego中带有很多可配置的参数，可以通过设置app.conf中对应的值，不区分大小写
### 基础配置
    * BConfig
### App配置
    * AppName：应用名称
    * RunMode：运行模式，可选值为prod、dev或test
    * RouterCaseSensitive：是否区分路由大小写，默认为true，区分大小写
    * ServerName：服务器默认在请求的时候输出server为beego
    * RecoverPanic：是否异常恢复，默认值为true，即当应用出现异常的情况，通过recover恢复回来，而不会导致应用异常退出
    * CopyRequestBody：是否允许在 HTTP 请求时，返回原始请求体数据字节，默认为 false （GET or HEAD or 上传文件请求除外）
    * EnableGzip：是否开启 gzip 支持，默认为 false 不支持 gzip，一旦开启了 gzip，那么在模板输出的内容会进行 gzip 或者 zlib 压缩，根据用户的 Accept-Encoding 来判断。
    * MaxMemory：文件上传默认内存缓存大小，默认值是 1 << 26(64M)。
    * EnableErrorsShow：是否显示系统错误信息，默认为 true。
    * EnableErrorsRender：是否将错误信息进行渲染，默认值为 true，即出错会提示友好的出错页面，对于 API 类型的应用可能需要将该选项设置为 false 以阻止在 dev 模式下不必要的模板渲染信息返回。
### Web配置
    * AutoRender：是否模板自动渲染，默认值为 true，对于 API 类型的应用，应用需要把该选项设置为 false，不需要渲染模板。
    * EnableDocs：是否开启文档内置功能，默认是 false
    * FlashName：Flash 数据设置时 Cookie 的名称，默认是 BEEGO_FLASH
    * FlashSeperator：Flash 数据的分隔符，默认是 BEEGOFLASH
    * DirectoryIndex：是否开启静态目录的列表显示，默认不显示目录，返回 403 错误。
    * StaticDir：静态文件目录设置，默认是static
    * StaticExtensionsToGzip：允许哪些后缀名的静态文件进行 gzip 压缩，默认支持 .css 和 .js
    * TemplateLeft：模板左标签，默认值是{{。
    * TemplateRight：模板右标签，默认值是}}。
    * ViewsPath：模板路径，默认值是 views。
    * EnableXSRF：是否开启 XSRF，默认为 false，不开启。
    * XSRFKEY：XSRF 的 key 信息，默认值是 beegoxsrf。 EnableXSRF＝true 才有效
    * XSRFExpire：XSRF 过期时间，默认值是 0，不过期。
### 监听配置
    * Graceful：是否开启热升级，默认是 false，关闭热升级。
    * ServerTimeOut：设置 HTTP 的超时时间，默认是 0，不超时。
    * ListenTCP4：监听本地网络地址类型，默认是TCP6，可以通过设置为true设置为TCP4。
    * EnableHTTP：是否启用 HTTP 监听，默认是 true。
    * HTTPAddr：应用监听地址，默认为空，监听所有的网卡 IP。
    * HTTPPort：应用监听端口，默认为 8080。
    * EnableHTTPS：是否启用 HTTPS，默认是 false 关闭。当需要启用时，先设置 EnableHTTPS = true，并设置 HTTPSCertFile 和 HTTPSKeyFile
    * HTTPSAddr：应用监听地址，默认为空，监听所有的网卡 IP。
    * HTTPSPort：应用监听端口，默认为 10443
    * HTTPSCertFile：开启 HTTPS 后，ssl 证书路径，默认为空。
    * HTTPSKeyFile：开启 HTTPS 之后，SSL 证书 keyfile 的路径。
    * EnableAdmin：是否开启进程内监控模块，默认 false 关闭。
    * AdminAddr：监控程序监听的地址，默认值是 localhost 。
    * AdminPort：监控程序监听的地址，默认值是 8088 。
    * EnableFcgi：是否启用 fastcgi ， 默认是 false。
    * EnableStdIo：通过fastcgi 标准I/O，启用 fastcgi 后才生效，默认 false。
    * SessionOn：session 是否开启，默认是 false。
    * SessionProvider：session 的引擎，默认是 memory
    * SessionName：存在客户端的 cookie 名称，默认值是 beegosessionID
    * SessionGCMaxLifetime：session 过期时间，默认值是 3600 秒。
    * SessionProviderConfig：配置信息，根据不同的引擎设置不同的配置信息
    * SessionCookieLifeTime：session 默认存在客户端的 cookie 的时间，默认值是 3600 秒
    * SessionAutoSetCookie：是否开启SetCookie, 默认值 true 开启。
    * SessionDomain：session cookie 存储域名, 默认空。
### Log配置
    * AccessLogs：是否输出日志到 Log，默认在 prod 模式下不会输出日志，默认为 false 不输出日志。此参数不支持配置文件配置。
    * FileLineNum：是否在日志里面显示文件名和输出日志行号，默认 true。此参数不支持配置文件配置。
    * Outputs：日志输出配置，参考 logs 模块，console file 等配置，此参数不支持配置文件配置。
## 控制器函数
beego.Controller实现了beego.ControllerInterface接口，实现了如下函数

* Init(ct *context.Context, childName string, app interface{})

        这个函数主要初始化了 Context、相应的 Controller 名称，模板名，初始化模板参数的容器 Data，app 即为当前执行的 Controller 的 reflecttype，这个 app 可以用来执行子类的方法。

* Prepare()

        这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。    

* Get()
    
        如果用户请求的 HTTP Method 是 GET，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Get 请求。

* Post()
* Delete()
* Put()
* Head()
* Patch()
* Options()
* Finish()
    
        这个函数是在执行完相应的 HTTP Method 方法之后执行的，默认是空，用户可以在子 struct 中重写这个函数，执行例如数据库关闭，清理数据之类的工作。

* Render() error

        这个函数主要用来实现渲染模板，如果 beego.AutoRender 为 true 的情况下才会执行。
    
比较流行的架构：实现一个自己的基类，实现一些初始化的方法，然后其它所有的逻辑继承自该基类
```
type NestPreparer interface {
        NestPrepare()
}

type baseController struct {
        beego.Controller
        i18n.Locale
        user    models.User
        isLogin bool
}

func (this *baseController) Prepare() {
        this.Data["PageStartTime"] = time.Now()
        this.Data["AppDescription"] = utils.AppDescription
        this.Data["AppKeywords"] = utils.AppKeywords
        this.Data["AppName"] = utils.AppName
        this.Data["AppVer"] = utils.AppVer
        this.Data["AppUrl"] = utils.AppUrl
        this.Data["AppLogo"] = utils.AppLogo
        this.Data["AvatarURL"] = utils.AvatarURL
        this.Data["IsProMode"] = utils.IsProMode

        if app, ok := this.AppController.(NestPreparer); ok {
                app.NestPrepare()
        }
}
```

提前终止运行：在prepare阶段进行判断，如果用户认证不通过，就输出一段信息，然后直接终止进程，之后的post、get之类的不再执行
```
type RController struct {
    beego.Controller
}

func (this *RController) Prepare() {
    this.Data["json"] = map[string]interface{}{"name": "astaxie"}
    this.ServeJSON()
    this.StopRun()
}
```