domain
======
domain parser

## Usage
``` go
u, _ := domain.Parse("http://foo.bar.imroc.com.cn:8080/hello?q=arg")
fmt.Println(u.Subdomain)    // foo.bar
fmt.Println(u.Domain)       // imroc
fmt.Println(u.PublicSuffix) // com.cn
```
