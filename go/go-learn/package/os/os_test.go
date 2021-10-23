package os

import (
	"fmt"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	// 文件 ，进程，系统
	fmt.Println(os.Getpid())
	fmt.Println(os.Getenv("MN"))
	fmt.Println(os.Getpagesize())
	fmt.Println(os.Getwd())
	fmt.Println(os.Getppid())
	fmt.Println(os.ExpandEnv("hello $MN"))

	fmt.Println(os.Getgroups()) //获取调用者属于的组  [4 24 27 30 46 108 124 1000]
	fmt.Println(os.Getgid())    //获取调用者当前所在的组　1000
	//fmt.Println(os.Chown("tmp.txt", 1000, 46))  //更改文件所在的组

	//func Getpagesize() int　　　//获取底层系统内存页的数量
	//func Getpid() int　　　　//获取进程id
	//func Getppid() int             //获取调用者进程父id
	//
	//func Hostname() (name string, err error)    //获取主机名
	//func IsExist(err error) bool    　　　　　 //返回一个布尔值，它指明 err 错误是否报告了一个文件或者目录已经存在。它被ErrExist和其它系统调用满足。
	//
	//func IsNotExist(err error) bool　　　　　//返回一个布尔值，它指明 err 错误是否报告了一个文件或者目录不存在。它被ErrNotExist 和其它系统调用满足。
	//func IsPathSeparator(c uint8) bool

	fmt.Println(os.IsPathSeparator('/'))
	fmt.Println(os.Hostname())

	//func IsPermission(err error) bool   //判定 err 错误是否是权限错误。它被ErrPermission 和其它系统调用满足。
	//func Mkdir(name string, perm FileMode) error　//创建一个新目录，该目录具有FileMode权限， 当创建一个已经存在的目录时会报错

	//func MkdirAll(path string, perm FileMode) error　//创建一个新目录，该目录是利用路径（包括绝对路径和相对路径）进行创建的，如果需要创建对应的父目录，也一起进行创建，如果已经有了该目录，则不进行新的创建， 当创建一个已经存在的目录时，不会报错.
	//func Remove(name string) error           //删除文件或者目录
	//func RemoveAll(path string) error　　//删除目录以及其子目录和文件，如果path不存在的话，返回nil

	//
	//unc Rename(oldpath, newpath string) error　　//重命名文件，如果oldpath不存在，则报错no such file or directory
	//func SameFile(fi1, fi2 FileInfo) bool　　　　　　//查看f1和f2这两个是否是同一个文件，如果再Unix系统，这意味着底层结构的device和inode完全一致，在其他系统上可能是基于文件绝对路径的．SameFile只适用于本文件包stat返回的状态，其他情况下都返回false
	//func Setenv(key, value string) error           //设定环境变量，经常与Getenv连用，用来设定环境变量的值

	//func Create(name string) (file *File, err error)　　//创建一个文件，文件mode是0666(读写权限)，如果文件已经存在，则重新创建一个，原文件被覆盖，创建的新文件具有读写权限，能够备用与i/o操作．其相当于OpenFile的快捷操作，等同于OpenFile(name string, O_CREATE,0666)
	//func NewFile(fd uintptr, name string) *File　　　//根据文件描述符和名字创建一个新的文件
	//func Open(name string) (file *File, err error)　　　//打开一个文件，返回文件描述符，该文件描述符只有只读权限．他相当于OpenFile(name string,O_RDWR,0)
	//func OpenFile(name string, flag int, perm FileMode) (file *File, err error)　//指定文件权限和打开方式打开name文件或者create文件，其中flag标志如下:

	//O_RDONLY：只读模式(read-only)
	//O_WRONLY：只写模式(write-only)
	//O_RDWR：读写模式(read-write)
	//O_APPEND：追加模式(append)
	//O_CREATE：文件不存在就创建(create a new file if none exists.)
	//O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
	//O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
	//O_TRUNC：打开并清空文件
	//至于操作权限perm，除非创建文件时才需要指定，不需要创建新文件时可以将其设定为０.虽然go语言给perm权限设定了很多的常量，但是习惯上也可以直接使用数字，如0666(具体含义和Unix系统的一致).
	//func Pipe() (r *File, w *File, err error)        //返回一对连接的文件，从r中读取写入w中的数据，即首先向w中写入数据，此时从r中变能够读取到写入w中的数据，Pipe()函数返回文件和该过程中产生的错误．

	//func (f *File) Chdir() error　
	//func (f *File) Chmod(mode FileMode) error　　　//更改文件权限，其等价与os.Chmod(name string,mode FileMode)error
	//func (f *File) Chown(uid, gid int) error                     //更改文件所有者，与os.Chown等价
	//func (f *File) Close() error　　　　　　　　　　//关闭文件，使其不能够再进行i/o操作，其经常和defer一起使用，用在创建或者打开某个文件之后，这样在程序退出前变能够自己关闭响应的已经打开的文件．
	//func (f *File) Fd() uintptr　　　//返回系统文件描述符，也叫做文件句柄．
	//func (f *File) Name() string　　//返回文件名字，与file.Stat().Name()等价
	//func (f *File) Read(b []byte) (n int, err error)　//将len(b)的数据从f中读取到b中，如果无错，则返回n和nil,否则返回读取的字节数n以及响应的错误
	//func (f *File) ReadAt(b []byte, off int64) (n int, err error)　//和Read类似，不过ReadAt指定开始读取的位置offset
	//func (f *File) Readdir(n int) (fi []FileInfo, err error)
	//func (f *File) Stat() (fi FileInfo, err error)　　//返回文件描述相关信息，包括大小，名字等．等价于os.Stat(filename string)
	//func (f *File) Sync() (err error)                        //同步操作，将当前存在内存中的文件内容写入硬盘．
	//func (f *File) Truncate(size int64) error        //类似  os.Truncate(name, size),，将文件进行截断
	//func (f *File) Write(b []byte) (n int, err error)　　//将b中的数据写入f文件
	//func (f *File) WriteAt(b []byte, off int64) (n int, err error)　//将b中数据写入f文件中，写入时从offset位置开始进行写入操作
	//func (f *File) WriteString(s string) (ret int, err error)　　　//将字符串s写入文件中
	//func (f *File) Seek(offset int64, whence int) (ret int64, err error)　　　　//Seek设置下一次读或写操作的偏移量 offset ，根据 whence 来解析：0意味着相对于文件的原始位置，1意味着相对于当前偏移量，2意味着相对于文件结尾。它返回新的偏移量和错误（如果存在）。

	//func (m FileMode) IsDir() bool              //判断m是否是目录，也就是检查文件是否有设置的ModeDir位
	//func (m FileMode) IsRegular() bool　　//判断m是否是普通文件，也就是说检查m中是否有设置mode type
	//func (m FileMode) Perm() FileMode　　//返回m的权限位
	//func (m FileMode) String() string　　　　//返回m的字符串表示

	//func FindProcess(pid int) (p *Process, err error)　
	//func (p *Process) Signal(sig Signal) error　　　　//发送一个信号给进程p, 在windows中没有实现发送中断interrupt
	//func (p *Process) Wait() (*ProcessState, error)
	//p.Signal(os.Kill)
	//func (p *ProcessState) Exited() bool　　　　　　// 判断程序是否 已经退出
	//func (p *ProcessState) Pid() int                                //返回退出进程的pid
	//func (p *ProcessState) String() string                     //获取进程状态的字符串表示
	//func (p *ProcessState) Success() bool                    //判断程序是否 成功退出，而Exited则仅仅只是判断其是否退出
	//func (p *ProcessState) Sys() interface{}                //返回有关进程的系统独立的退出信息。并将它转换为恰当的底层类型（比如Unix上的syscall.WaitStatus），主要用来获取进程退出相关资源。
	//func (p *ProcessState) SysUsage() interface{}         //SysUsage返回关于退出进程的系统独立的资源使用信息。主要用来获取进程使用系统资源
	//func (p *ProcessState) SystemTime() time.Duration      //获取退出进程的内核态cpu使用时间
	//func (p *ProcessState) UserTime() time.Duration     //返回退出进程和子进程的用户态CPU使用时间

	fd, err := os.Stat("tmp.txt")
	if err != nil {
		fmt.Println(err)
	}
	fm := fd.Mode()
	fmt.Println(fm.IsDir())     //false
	fmt.Println(fm.IsRegular()) //true
	fmt.Println(fm.Perm())      //-rwxrwxrwx
	fmt.Println(fm.String())    //-rwxrwxrwx

	r, w, _ := os.Pipe()
	w.WriteString("hello,world!")
	var s = make([]byte, 20)
	n, _ := r.Read(s)
	fmt.Println(string(s[:n]))
	//os.Exit(1)
}
