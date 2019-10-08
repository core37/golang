package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"os"
	// "os/exec"
	// "bufio"
	// "io"
)

var (
	h bool
	f bool
	l int
	s, e int
	d string
	sstdin string
	sstdout string
	sstderr string
	filename string
)

func init() {
	flag.BoolVarP(&h, "h","h", false, "this help")

	flag.IntVarP(&s, "s","s", -1, "start page")
	flag.IntVarP(&e, "e", "e", -1, "end page")

	flag.BoolVarP(&f, "f","f", false, "page-end-mark used")
	flag.IntVarP(&l, "l","l", 72, "line of page")

	flag.StringVarP(&d, "d","d", "", "send to destination")
	// // 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	// flag.StringVarP(&sstdin, "<","<", "", "send `signal` to a master process: stop, quit, reopen, reload")
	// flag.StringVarP(&sstdout, ">",">", "/usr/local/nginx/", "set `prefix` path")
	// flag.StringVarP(&sstderr, "2>","2>", "conf/nginx.conf", "set configuration `file`")

	// 改变默认的 Usage
	flag.Usage = usage
}

func main() {
	flag.Parse()
	// for i:=0; i<flag.NArg(); i++{
	// 	if flag.Args()[i] == "2<"{
	// 		f, _ := os.OpenFile(flag.Args()[i + 1], os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND,0755)
	// 		os.Stdin = f
	// 	}
	// 	if flag.Args()[i] == "3>"{
	// 		fmt.Fprintf(os.Stdout, flag.Args()[i + 1])
	// 		// f, _ := os.OpenFile(flag.Args()[i + 1], os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND,0755)
	// 		// os.Stdout = f
	// 	}
	// 	if flag.Args()[i] == "2>"{
	// 		f, _ := os.OpenFile(flag.Args()[i + 1], os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND,0755)
	// 		os.Stderr = f
	// 	}
	// }
	if s <= 0 || e <= 0 || s > e{
		fmt.Fprintf(os.Stderr, "you need to set a correct page range")
		return
	}
	if l < 0{
		fmt.Fprintf(os.Stderr, "you need to set a correct page line number")
		return
	}
	if f == true && l!=72{
		fmt.Fprintf(os.Stderr, "can't appear -f and -l at the s")
	}
	if h {
		flag.Usage()
		return
	}

	if flag.NArg != 0 {
		_, err := os.Stat(flag.Agrs()[0])
		if (err != nil){
			fmt.Fprintf(os.Stderr, "File not exist")
		}
		filename = flag.Agrs()[0]
	}


	var count,page int
	count = 0
	page = 1


	if flag.NArg != 0{
		fin := os.Stdin
	}else{
		var err error
		fin, err = os.Open(filename)
	}
	cmd := &exec.Cmd{}
	buffer := bufio.NewReader(fin)


	if len(d) == 0{
		fout := os.Stdout
	}else{
		cmd = exec.Command("cat")

		var err error
		cmd.Stdout, err = os.OpenFile(d, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err !=nil{
			fmt.Fprintf(os.Stderr, "Can't open output file")
		}

		fout, err = cmd.StdinPipe()

		if err != nil{
			fmt.Fprintf(os.Stderr, "Can't open pipe")
		}

		cmd.Start()
		defer fout.Close()
	}


	if f{
		page = 1
		for{
			pagestr, err := bufFin.ReadString('\f')
			if err!=nil{
				break
			}

			if page >= s && page <= e{
				_, err := fout.Write([]byte(pagestr))
			}
			page++
		}
	}else{
		line, _ := bufFin.ReadString('\n')
		count ++
		if count == l{
				page++
			count = 0
		}
		if page >= s && page <= e{
			_, err := fout.Write([]byte(line))
		}
	}
}



















	}




}

func usage() {
	flag.PrintDefaults()
}