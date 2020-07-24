//产生JDBC加密文件
package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

//go build -ldflags="-H windowsgui"

var (
	url      *walk.LineEdit
	user     *walk.LineEdit
	password *walk.LineEdit
	path     *walk.LineEdit
)

const (
	DEFAULT_FILE_NAME = "jdbc.c_properties"
	DEFAULT_PATH      = "D:\\"
	DEFAULT_URL       = "jdbc:oracle:thin:@//IP:PORT/SID"
	HELP_INFO         = "填写信息后将生成的文件jdbc.c_properties放入WEB-INF\\classes\\configs中"
)

var mymw = &browseMw{}

//主窗口
var mw = MainWindow{
	AssignTo: &mymw.MainWindow,
	Title:    "JDBC加密工具",
	MinSize:  Size{Width: 220, Height: 50},
	Layout:   VBox{},
	Children: widget,
}

var widget = []Widget{
	Composite{
		Layout: Grid{Columns: 2}, //columns表示显示的列数
		Children: []Widget{
			urlLabel,
			urlle,

			userLabel,
			userle,

			passwordLabel,
			passwordle,

			pathLabel,
			pathle,

			browseButton,
			genButton,

			helpLabel,
		},
	},
}

//URL标签
var urlLabel = Label{Text: "JDBC URL:"}

//用户名的标签
var userLabel = Label{Text: "用户名:"}

//密码的标签
var passwordLabel = Label{Text: "密码:"}

//路径标签
var pathLabel = Label{Text: "保存路径:"}

var helpLabel = Label{
	Background: SolidColorBrush{Color: walk.RGB(255, 191, 0)},
	Text:       HELP_INFO,
}

//URL输入框
//Alignment-AlignNear: 左对齐
//MaxLength：最大长度
var urlle = LineEdit{AssignTo: &url, Text: DEFAULT_URL}

//用户名输入框
var userle = LineEdit{Alignment: AlignHNearVNear, AssignTo: &user}

//密码输入框
//PasswordMode:密码模式，使输入的字符显示为**
//var passwordle = LineEdit{MaxSize: textSize,Alignment:AlignNear,PasswordMode: true, MaxLength: 100, AssignTo: &password}
var passwordle = LineEdit{Alignment: AlignHNearVNear, PasswordMode: true, AssignTo: &password}

//保存路径输入框
var pathle = LineEdit{Alignment: AlignHNearVNear, AssignTo: &path, Text: DEFAULT_PATH}

//浏览按钮
var browseButton = PushButton{
	Text:      "浏览",
	OnClicked: mymw.getpath,
}

//生成按钮
var genButton = PushButton{
	Text:      "生成",
	OnClicked: genFile,
}

func main() {
	_, err := mw.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
