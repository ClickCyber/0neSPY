<?php


$sql_hostname = "127.0.0.1";
$sql_username = "root";
$sql_password = "pass";
$sql_database = "0neSPY";

try {
  $sql = new PDO("mysql:host=$sql_hostname;dbname=$sql_database", $sql_username, $sql_password);
  $sql->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
  //echo "work!";
} catch(PDOException $e) {
  die("Connection failed Cheak Logs"); // write log $e->getMessage()
}
?>