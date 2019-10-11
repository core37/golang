package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"os"
	"os/exec"
	"bufio"
    "io"
    // "strconv"
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


func main() {
    flag.BoolVarP(&h, "h","h", false, "this help")
	flag.IntVarP(&s, "s","s", -1, "start page")
	flag.IntVarP(&e, "e", "e", -1, "end page")
	flag.BoolVarP(&f, "f","f", false, "page-end-mark used")
	flag.IntVarP(&l, "l","l", 72, "line of page")
	flag.StringVarP(&d, "d","d", "", "send to destination")
	flag.Parse()
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

	if flag.NArg() != 0 {
		_, err := os.Stat(flag.Args()[0])
		if (err != nil){
			fmt.Fprintf(os.Stderr, "File not exist")
		}
		filename = flag.Args()[0]
	}


	var count,page int
	count = 0
	page = 1
    var fin *os.File
    fin = os.Stdin
	if len(filename) == 0{
		fin = os.Stdin
	}else{
        var err error
		fin, err = os.Open(filename)
        if  err!=nil{
            fmt.Fprintf(os.Stderr, " can't open file\n")
        }
	}
	cmd := &exec.Cmd{}
	bufFin := bufio.NewReader(fin)

    var fout io.WriteCloser
    fout = os.Stdout
	if len(d) == 0{
		fout = os.Stdout
	}else{
		cmd = exec.Command("cat")

		var err error
		cmd.Stdout, err = os.OpenFile(d, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err !=nil{
			fmt.Fprintf(os.Stderr, "Can't open output file")
		}

		fout, err = cmd.StdinPipe()

		if err != nil{
			fmt.Fprintf(os.Stderr, "Can't open pipe\n")
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
                if err !=nil{
			        fmt.Fprintf(os.Stderr, "Write err")
	        	}
			}
			page++
		}
	}else{

        page = 1
        count = 0
        for{
            line, erri := bufFin.ReadString('\n')
            if erri != nil{
                break
            }
            // fmt.Fprintf(os.Stderr, "!%d\n", count)
		    if count == l{
		    	page++
		    	count = 0
		    }
		    count ++
		    if page >= s && page <= e{
		    	_, err := fout.Write([]byte(line))
                if err !=nil{
		        	fmt.Fprintf(os.Stderr, "Write err")
		        }
		    }
            if page > e{
                break
            }
        }
	}
    fmt.Fprintf(os.Stderr, "ERROR TESTING\n")
}

func usage() {
	flag.PrintDefaults()
}