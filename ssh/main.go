package main

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"strings"
)

/**
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' main.go
upx main.exe
*/
type Tunnel struct {
	Remote string
	Local  string
}

type SSHTunnel struct {
	sshClient  *ssh.Client
	Addr       string
	User       string
	Pass       string
	Tunnels    []Tunnel
	BufferSise int64
	Timeout    time.Duration
}

func (st *SSHTunnel) Close() {
	if nil != st.sshClient {
		st.sshClient.Close()
	}
}

func (st *SSHTunnel) GetSSHClient() (*ssh.Client, error) {
	if st.sshClient != nil {
		return st.sshClient, nil
	}
	var auth []ssh.AuthMethod
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(st.Pass))

	sc := &ssh.ClientConfig{
		User: st.User,
		Auth: auth,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: st.Timeout * time.Millisecond,
	}
	var err error
	st.sshClient, err = ssh.Dial("tcp", st.Addr, sc)
	if err != nil {
		return nil, err
	}
	//log.Printf("连接到服务器成功: %s", st.Addr)
	return st.sshClient, err
}

func (st *SSHTunnel) ClientClose() {
	if st.sshClient != nil {
		st.sshClient.Close()
		st.sshClient = nil
	}
}

func (st *SSHTunnel) connect(t Tunnel) {
	ll, err := net.Listen("tcp", t.Local)
	if err != nil {
		//log.Printf(`开启本地端口监听失败: %s, %s`, t.Local, err)
		return
	}
	defer func() {
		ll.Close()
		//log.Printf(`断开隧道连接：%s <=> %s`, t.Local, t.Remote)
	}()
	//log.Printf(`开启隧道：%s <=> %s`, t.Local, t.Remote)
	sno := int64(0)
	for {
		lc, err := ll.Accept()
		if err != nil {
			//log.Printf(`接受来自本地的连接失败: %s`, err)
			return
		}
		//log.Printf(`接收到本地连接 => %s`, t.Local)
		sc, err := st.GetSSHClient()
		if err != nil {
			//log.Printf(`连接到服务器失败: %s`, err)
			lc.Close()
			continue
		}
		rc, err := sc.Dial("tcp", t.Remote)
		if err != nil {
			//log.Printf(`连接到远程主机失败: %s`, err)
			st.ClientClose()
			lc.Close()
			continue
		}
		//log.Printf(`连接到远程主机 => %s `, t.Remote)
		sno = sno + 1
		cid := fmt.Sprintf("%s <=> %s: %d", t.Local, t.Remote, sno)
		st.transfer(cid, lc, rc)
	}
}

func main() {
	if getmac("8c:16:45:13:0c:d4") || getmac("52:54:00:3c:50:50") {
		println("开始使用吧~~")
		A := ""
		B := ""
		C := ""
		AA, _ := b64.StdEncoding.DecodeString(A)
		BB, _ := b64.StdEncoding.DecodeString(B)
		CC, _ := b64.StdEncoding.DecodeString(C)
		sts := []SSHTunnel{{
			Addr: BytesToString(AA),
			User: BytesToString(BB),
			Pass: BytesToString(CC),
			Tunnels: []Tunnel{{
				Remote: "127.0.0.1:37799",
				Local:  "0.0.0.0:37799",
			}, {
				Remote: "127.0.0.1:7072",
				Local:  "0.0.0.0:7071",
			},
			},
		},
		}
		var wg sync.WaitGroup
		for _, st := range sts {
			st.check()
			wg.Add(1)
			go func() {
				start(st)
				wg.Done()
			}()
			//log.Printf(`启动隧道配置：%s`, st.Addr)
		}
		wg.Wait()
	} else {
		os.Exit(1)
	}
}

func (st *SSHTunnel) setPass() {
	//fmt.Printf("请输入登陆密码[%s@%s]:", st.User, st.Addr)
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	st.Pass = string(bytePassword)
	//fmt.Println()
}

func (st *SSHTunnel) check() {
	if len(st.Pass) == 0 {
		st.setPass()
	}

	if st.BufferSise == 0 {
		st.BufferSise = 5 * 1024
	}

	if st.Timeout == 0 {
		st.Timeout = 3000
	}

	st.initSSHClient()
}

func (st *SSHTunnel) initSSHClient() {
	_, err := st.GetSSHClient()
	if nil != err {
		error := err.Error()
		//log.Printf(`连接主机失败: %s`, error)
		if strings.Contains(error, "unable to authenticate") {
			st.Pass = ""
			st.setPass()
			st.initSSHClient()
			return
		}
		if strings.Contains(error, "i/o timeout") {
			//log.Printf(`连接到服务器超时: %s`, st.Addr)
			os.Exit(-1)
		}
	}
}

func start(st SSHTunnel) {
	defer st.Close()
	var wg sync.WaitGroup
	for _, t := range st.Tunnels {
		wg.Add(1)
		go func(tunnel Tunnel) {
			st.connect(tunnel)
			wg.Done()
		}(t)
	}
	wg.Wait()
}

func (st SSHTunnel) transfer(cid string, lc net.Conn, rc net.Conn) {
	copyBufPool := sync.Pool{
		New: func() interface{} {
			b := make([]byte, st.BufferSise)
			return &b
		},
	}
	go func() {
		defer lc.Close()
		defer rc.Close()
		//log.Printf(`连接下行通道：%s`, cid)
		bufp := copyBufPool.Get().(*[]byte)
		defer copyBufPool.Put(bufp)
		io.CopyBuffer(lc, rc, *bufp)
		//log.Printf(`断开下行通道：%s`, cid)
	}()
	go func() {
		defer rc.Close()
		defer lc.Close()
		//log.Printf(`连接上行通道：%s`, cid)
		bufp := copyBufPool.Get().(*[]byte)
		defer copyBufPool.Put(bufp)
		io.CopyBuffer(rc, lc, *bufp)
		//log.Printf(`断开上行通道：%s`, cid)
	}()
}
func getmac(m string) bool {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}
	var t int = 0
	for _, inter := range interfaces {
		//var mac []byte
		//mac = inter.HardwareAddr //获取本机MAC地址
		mac := inter.HardwareAddr.String()
		//fmt.Println("MAC = ",mac)
		if mac == m {
			t = t + 1
		}
	}
	if t >= 1 {
		return true
	} else {
		return false
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}
