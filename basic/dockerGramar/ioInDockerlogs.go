package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
  @Author:      He Bao Jing
  @Date:        5/31/2023 9:52 AM
  @Description:
*/

func ioCopyDemo() {
	url := "https://cn.bing.com/images/search?view=detailV2&ccid=5fsFGVxE&id=FC129B2ACEB7317091C0F1F7ADF787FF88F8ADBB&thid=OIP.5fsFGVxEMzwz8Ccd6Km24gHaFv&mediaurl=https%3a%2f%2fts1.cn.mm.bing.net%2fth%2fid%2fR-C.e5fb05195c44333c33f0271de8a9b6e2%3frik%3du634iP%252bH96338Q%26riu%3dhttp%253a%252f%252fcircuitcellar.com%252fwp-content%252fuploads%252f2018%252f12%252fGo-IO-PR-Copy.jpg%26ehk%3dE3pFE5%252fba0LVDV83xndhNPhkmEN9Kn6KIUrNdQYLJiw%253d%26risl%3d%26pid%3dImgRaw%26r%3d0&exph=621&expw=800&q=go+io.copy+%e7%94%a8%e6%b3%95&simid=608028612102464589&FORM=IRPRST&ck=DD6C0E1E0652F62B946A438F53E1A0BF&selectedIndex=2"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("fail to get the url, ", err)
		return
	}
	defer res.Body.Close()
	out, err := os.Create("D:\\tmp\\test.jpg")
	wt := bufio.NewWriter(out)
	defer out.Close()

	n, err := io.Copy(wt, res.Body)
	fmt.Println("n is ", n)

	wt.Flush()
}

func main() {
	ioCopyDemo()
}
