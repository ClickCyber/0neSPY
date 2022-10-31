<?php
session_start();

include 'functions.php';

if (!isset($_SESSION['login'])){
  die(header("location: login.php"));
}


// menu admin


?>
