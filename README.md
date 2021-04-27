# germaniumlib
carbon-now-sh alternative code beautifyer Golang library

# Example

![image](https://github.com/raifpy/germaniumlib/blob/main/image.png)
```go
package main

import g "github.com/raifpy/germaniumlib"

func main(){
    f,_ := os.Create("image.png")
    //b := &bytes.Buffer{}
    err := g.Run(g.Options{
		Code:            "package main\nfunc main(){}",
		Writer:          f,
		Language:        "go",
		BackgroundColor: "#aaaaff",
    })
    f.Close()
    if err != nil{panic(err)}
    
    
    
}

```
