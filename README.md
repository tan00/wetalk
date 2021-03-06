# WeTalk

An open source project for Gopher community.

### Usage

```
cd /path/to/wetalk
go mod init wetalk
```
add 'replace github.com/beego/wetalk => /path/to/local/wetalk'  in go.mod

```
go mod tidy
```


Copy `conf/global/app.ini` to `conf/app.ini` and edit it. All configure has comment in it.

The files in `conf/` can overwrite `conf/global/` in runtime.


**Run WeTalk**

```
bee run watchall
```

### Dependencies
* Beego [https://github.com/astaxie/beego](https://github.com/astaxie/beego) (develop branch)
* Social-Auth [https://github.com/beego/social-auth](https://github.com/beego/social-auth)
* Compress [https://github.com/beego/compress](https://github.com/beego/compress)
* i18n [https://github.com/beego/i18n](https://github.com/beego/i18n)
* Mysql [https://github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
* goconfig [https://github.com/Unknwon/goconfig](https://github.com/Unknwon/goconfig)
* fsnotify [https://github.com/howeyc/fsnotify](https://github.com/howeyc/fsnotify)
* resize [https://github.com/nfnt/resize](https://github.com/nfnt/resize)
* blackfriday [https://github.com/slene/blackfriday](https://github.com/slene/blackfriday)



### Static Files

WeTalk use `Google Closure Compile` and `Yui Compressor` compress js and css files.

So you could need Java Runtime. Or close this feature in code by yourself.

### WeTalk in world

[Go China Community](http://bbs.go-china.org/)

### Contact

Maintain by [slene](https://github.com/slene)

## License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
