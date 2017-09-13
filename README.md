domain
======
domain parser

## Usage
``` go
u, _ := domain.Parse("http://foo.bar.imroc.com.cn:8080/hello?q=arg")
println(u.Subdomain)    // foo.bar
println(u.Domain)       // imroc
println(u.PublicSuffix) // com.cn
```
