/*
all server config code
*/

package cfg

import (
	"fmt"
	"muyusvr/convert"
	"muyusvr/muyudb"
	"muyusvr/muyulog"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/widuu/goini"
)

const (
	linuxcfg string = "/etc/muyu/muyu.ini"
)

const (
	defaultcfg string = `[server]
port=8080
addr=0.0.0.0
logdir=muyu.log
logsize=10
logsave=30
[db]
host=127.0.0.1
port=3306
user=muyu
pwd=123456
dbname=muyu`
)

type Server struct {
	Port    int
	Addr    string
	Logdir  string
	Logsize int
	Logsave int
}

type DbServer struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DBnane string
}

type ServerConf struct {
	ServerIni *Server
	DBServer  *DbServer
}

func createSvrcfgInstanse(p string) *ServerConf {
	if p == "" {
		return &ServerConf{
			ServerIni: &Server{
				Port:    8080,
				Addr:    "0.0.0.0",
				Logdir:  "muyu.log",
				Logsize: 10,
				Logsave: 30,
			},
			DBServer: &DbServer{
				Host:   "127.0.0.1",
				Port:   3306,
				User:   "muyu",
				Pwd:    "123456",
				DBnane: "muyu",
			},
		}
	} else {
		conf := goini.SetConfig(p)
		return &ServerConf{
			ServerIni: &Server{
				Port:    convert.StrToInt(conf.GetValue("server", "port")),
				Addr:    conf.GetValue("server", "addr"),
				Logdir:  conf.GetValue("server", "logdir"),
				Logsize: convert.StrToInt(conf.GetValue("server", "logsize")),
				Logsave: convert.StrToInt(conf.GetValue("server", "logsave")),
			},
			DBServer: &DbServer{
				Host:   conf.GetValue("db", "host"),
				Port:   convert.StrToInt(conf.GetValue("db", "port")),
				User:   conf.GetValue("db", "user"),
				Pwd:    conf.GetValue("db", "pwd"),
				DBnane: conf.GetValue("db", "dbname"),
			},
		}
	}
}

//InitedSvr 初始化服务
func InitedSvr() (error, *ServerConf) {
	systype := runtime.GOOS
	if systype == "linux" {
		return chkcfgfile(defaultcfg)
	} else if systype == "windows" {
		u, e := user.Current()
		if e != nil {
			return e, nil
		}
		return chkcfgfile(fmt.Sprintf("%s\\muyu\\muyu.ini", u.HomeDir))
		//return chkcfgfile("muyu.ini")

	} else {
		return fmt.Errorf("unknow system"), nil
	}

}

func chkcfgfile(fpath string) (error, *ServerConf) {
	s := new(ServerConf)
	var e error
	_, e = os.Stat(fpath)
	if e != nil {
		if !os.IsExist(e) {

			fd := filepath.Dir(fpath)
			_, e = os.Stat(fd)
			if os.IsNotExist(e) {
				os.MkdirAll(fd, os.ModePerm)
			}
			f, fe := os.OpenFile(fpath, os.O_CREATE, os.ModePerm)
			defer f.Close()
			if fe != nil {
				return fe, nil
			}
			f.WriteString(defaultcfg)
			f.Sync()
			e = fe
			s = createSvrcfgInstanse("")
		}
	} else {
		//读取配置文件
		s = createSvrcfgInstanse(fpath)

	}
	return e, s
}

func InitedDB(dbc *DbServer) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbc.User, dbc.Pwd, dbc.Host, dbc.Port, dbc.DBnane)
	return muyudb.InitedMuYuDB(dsn)
}

func InitedLog(s *Server) {
	_, e := os.Stat(s.Logdir)
	if e != nil {
		if !os.IsExist(e) {
			fd := filepath.Dir(s.Logdir)
			if fd != "." {
				_, e = os.Stat(fd)
				if os.IsNotExist(e) {
					os.MkdirAll(fd, os.ModePerm)
				}

			}
		}
	}
	muyulog.InitLogger(s.Logdir, s.Logsize, s.Logsave)
}
