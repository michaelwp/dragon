<h2 align="center">
  <a target="_blank">
    <img src="https://github.com/JayantGoel001/JayantGoel001/blob/master/GIF/Earth.gif" width="24px" style="max-width:100%;">
  </a>
  𝐇𝐞𝐥𝐥𝐨, &lt;𝚌𝚘𝚍𝚎𝚛𝚜/&gt;!
  <a target="_blank">
    <img src="https://github.com/JayantGoel001/JayantGoel001/blob/master/GIF/Hi.gif" width="40px" />
  </a>
</h2>

# Dragon 🐉
<div align="center">

&nbsp; ![Open Source Love](https://img.shields.io/badge/Open%20Source-%F0%9F%92%9B-cyan.svg?style=flat) &nbsp; ![](https://visitor-badge.glitch.me/badge?page_id=michaelwp.dragon&style=flat-square&color=0088cc) &nbsp; <img src="https://img.shields.io/static/v1?label=%F0%9F%8C%9F&message=If%20Useful&style=style=flat&color=BC4E99" alt="Star Badge"/>
	
</div>
<img height="50" alt="Thanks for visiting me" width="100%" src="https://raw.githubusercontent.com/BrunnerLivio/brunnerlivio/master/images/marquee.svg" />

<br>

<a target="_blank">
  <img align="right" height="250" width="400" alt="GIF" src="https://github.com/JayantGoel001/JayantGoel001/blob/master/GIF/image.gif">
</a>

## Golang router package 🍁

* &nbsp; ⚙️ [Install](#install)
* &nbsp; 🚀 [Examples](#examples)

</br>


## Install

&nbsp; 👾 Run this command to install.
<h3>
	
```bash
go get github.com/michaelwp/dragon
```
	
<h3>

## Examples

<h3>
	
```go
package main

import (
	"github.com/michaelwp/dragon"
	"log"
)

type response struct {
	Message string `json:"message"`
}

func main() {
	r := dragon.NewRouter()
	
	api := r.Group("/api/v1")
	api.GET("/home", func(d *dragon.Dragon) error {
		var resp response
		resp.Message = "Hello World"
		return d.ResponseJSON(200, resp)
	})

	log.Fatal(r.Run(":8090"))
}
```
<h3>

## License

&nbsp; ⚡ BSD licensed. See the LICENSE file for details.

---
	
<div align="center">

### Show some ❤️ by starring the repository!

</div>
	
#

![footer](https://github.com/JayantGoel001/JayantGoel001/blob/master/PNG/footer.png)	
