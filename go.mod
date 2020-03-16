module wetalk

go 1.14

require (
	github.com/Unknwon/goconfig v0.0.0-20191126170842-860a72fb44fd
	github.com/astaxie/beego v1.12.1
	github.com/beego/compress v0.0.0-20131031122501-fbacc486da67
	github.com/beego/i18n v0.0.0-20161101132742-e9308947f407
	github.com/beego/social-auth v0.0.0-20161029090309-a8ef647f2fd8
	github.com/beego/wetalk v0.0.0-20160915090959-b0d9d4d558e3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/howeyc/fsnotify v0.9.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/slene/blackfriday v0.0.0-20140117134005-fd3fc8f180b2
)

//replace github.com/beego/wetalk => /Users/ln/mysrc/go/wetalk
replace github.com/beego/wetalk => /home/zcw/mysrc/wetalk
