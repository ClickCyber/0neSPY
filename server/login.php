
<?php
include 'functions.php';
if (isset($_SESSION['login'])){
  die(header("location: ./index.php"));
}else if(isset($_POST['password']) and isset($_POST['email'])){
  if (Login($_POST['password'], $_POST['email'])){
    $_SESSION['login'] = true;
    $_SESSION['email'] = $_POST['email'];
    die(header("location: ./index.php"));
  }else
    $msg = "Invalid email/passowrd";
}
?>
<!doctype html>
<html lang="en">
  <head>
     <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="0neSPY - Login">
    <meta name="author" content="ClickCyber">
    <meta name="generator" content="beta">
    <title>0neSPY - Login</title>
    <link href="vendor/css/bootstrap.min.css" rel="stylesheet">
    <link href="vendor/css/signin.css" rel="stylesheet">
  </head>
   <body class="text-center">
    <main class="form-signin w-100 m-auto">
      <form method="POST">
        <img class="mb-4" src="/docs/5.2/assets/brand/bootstrap-logo.svg" alt="logo" width="72" height="57">
        <?php
        if (isset($msg))
          echo '<h1 class="h3 mb-3 fw-normal">'.$msg.'</h1>';
        else
          echo '<h1 class="h3 mb-3 fw-normal">Please sign in</h1>';
        ?>
        <div class="form-floating">
          <input type="email" class="form-control" name="email" id="floatingInput" placeholder="SpyWare@0neSPY.xyz">
          <label for="floatingInput">Email address</label>
        </div>
        <div class="form-floating">
          <input type="password" class="form-control" name="password" id="floatingPassword" placeholder="Password">
          <label for="floatingPassword">Password</label>
        </div>

        <div class="checkbox mb-3">
          <label>
            <input type="checkbox" value="remember-me"> Remember me
          </label>
        </div>
        <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
        <p class="mt-5 mb-3 text-muted">&copy; 2022-<?php echo date("Y");?></p>
      </form>
    <script src="vendor/js/bootstrap.bundle.min.js"></script>
    <script src="vendor/js/bootstrap.min.js"></script>
    <script src="vendor/js/popper.min.js"></script>
  </body>
</html>