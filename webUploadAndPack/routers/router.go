package routers

import (
	"github.com/astaxie/beego"
	"golang-test/webUploadAndPack/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/file", &controllers.FileController{})
	beego.Router("/download", &controllers.DownloadController{})
	beego.Router("/erro", &controllers.ErroController{})
}
