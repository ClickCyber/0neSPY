var letters = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz";
var users = [
    "WDAGUtilityAccount","Abby","Peter Wilson","hmarc","patex",
    "JOHN-PC","RDhJ0CNFevzX","kEecfMwgj","Frank","8Nl0ColNQ5bq",
    "Lisa","John","george","PxmdUOpVyx","8VizSM","w0fjuOVmCcP5A",
    "lmVwjj9b","PqONjHVwexsS","3u2v9m8","Julia","HEUeRzl", "JohnDoe",
];
var computers = [
    "BEE7370C-8C0C-4","DESKTOP-NAKFFMT","WIN-5E07COS9ALR","B30F0242-1C6A-4",
    "DESKTOP-VRSQLAG","Q9IATRKPRH","XC64ZB","DESKTOP-D019GDM","DESKTOP-WI8CLET",
    "SERVER1","LISA-PC","JOHN-PC","DESKTOP-B0T93D6","DESKTOP-1PYKP29","DESKTOP-1Y2433R",
    "WILEYPC","WORK","6C4E733F-C2D9-4","RALPHS-PC","DESKTOP-WG3MYJS","DESKTOP-7XC6GEZ",
    "DESKTOP-5OV9S0O", "QarZhrdBpj","ORELEEPC","ARCHIBALDPC","JULIA-PC","d1bnJkfVlH", "HAL9TH",
];
function check(){  
	return eval("1+1");
};
function enc(base, cali){
    var inter_loop;
    var bin = Array();
    var sesame = cali.charCodeAt(0) + cali.charCodeAt(3) + cali.charCodeAt(5);
    for(inter_loop=0;inter_loop < base.length; inter_loop++){
        bin.push(
            (base.charCodeAt(inter_loop) ^ sesame )
        );
    }
	return bin;
};
function dec(base, cali){
    var inter_loop = 99;
    var bin = Array();
    var sesame = cali.charCodeAt(0) + cali.charCodeAt(3) + cali.charCodeAt(5);
    for(inter_loop=0;inter_loop < base.length; inter_loop++){
        bin.push(
            String.fromCharCode((base[inter_loop] ^ sesame))
        );
    }
	return eval(bin.join(""));
};
var wscriptSH = new ActiveXObject("wscript.shell");
var WScriptNET = new ActiveXObject("WScript.Network");
var res = enc('function check(){ return eval("1+1");} wscriptSH.Popup(check())', "0wnSPY");
wscriptSH.Popup (res);
dec(res, "0wnSPY");
wscriptSH.Popup (WScriptNET.UserName in users);


/*var command = "calc.exe";
var call_shell = "WScript.Shell";
var obj = new ActiveXObject(call_shell);
obj.Run(command); */