package routers

import (
	"golang-test/webUploadAndPack/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/upload", &controllers.UploadController{})
    beego.Router("/file", &controllers.FileController{})
    beego.Router("/download", &controllers.DownloadController{})
    beego.Router("/erro", &controllers.ErroController{})
}
