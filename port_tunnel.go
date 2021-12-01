package main
import (
	"fmt"
	"io"
	"log"
	"net"
)
var (
	dst string = "127.0.0.1:1234"
	local string = "0.0.0.0:222"
)
func main() {
	listen()
}
func listen()  {
	for {
		ln,err := net.Listen("tcp",local)
		if err != nil {
			fmt.Println("tcp_listen:",err)
			return
		}
		defer ln.Close()
		for{
			tcp_Conn,err:=ln.Accept()
			if err!=nil{
				fmt.Println("Accept:",err)
				return
			}
			go tcp_handle(tcp_Conn)
		}
	}
}
func tcp_handle(tcpConn net.Conn){
	remote_tcp,err:=net.Dial("tcp",dst)
	if err!=nil{
		fmt.Println(err)
		return
	}
	log.Println(dst,"=>",local)
	go io.Copy(remote_tcp,tcpConn)
	log.Println(local,"=>",dst)
	go io.Copy(tcpConn,remote_tcp)
}
package main
import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)
var (
	local = flag.String("l", "0.0.0.0:222", "0.0.0.0:222 [本地网卡监听端口]")
	dst = flag.String("r", "8.8.8.8:1234", "8.8.8.8:1234 [NC、CobaltStrike、Metasploit等服务监听端口]")
)
func main() {
	flag.Parse()
	if len(os.Args)<=2 {
		flag.Usage()
	}else {
		listen()
	}

}


func listen()  {
	log.Println("Start listen",*local)
	for {
		ln,err := net.Listen("tcp",*local)
		if err != nil {
			fmt.Println("tcp_listen:",err)
			return
		}
		defer ln.Close()
		for{
			tcp_Conn,err:=ln.Accept()
			if err!=nil{
				fmt.Println("Accept:",err)
				return
			}
			go tcp_handle(tcp_Conn)
		}
	}
}
func tcp_handle(tcpConn net.Conn){
	remote_tcp,err:=net.Dial("tcp",*dst)
	if err!=nil{
		fmt.Println(err)
		return
	}
	log.Println(dst,"=>",*local)
	go io.Copy(remote_tcp,tcpConn)
	log.Println(local,"=>",*dst)
	go io.Copy(tcpConn,remote_tcp)

}
