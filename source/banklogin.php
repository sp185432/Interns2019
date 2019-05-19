<?php 
session_start();
        
if(isset($_SESSION['staff_login'])) 
    header('location:bank_homepage.php');   
?>
<!DOCTYPE html>
<html>
    <head>
        <noscript><meta http-equiv="refresh" content="0;url=no-js.php"></noscript>  
        <meta charset="UTF-8">
        <title>Bank Login</title>
        
        <link rel="stylesheet" href="newcss.css">
    </head>
<?php
include 'header.php'; ?>

<div class='content'>
<div class="user_login">
    <form action='' method='POST'>
        <table align="center">
            <tr><td><span class="caption">Bank Login</span></td></tr>
            <tr><td colspan="2"><hr></td></tr>
            <tr><td>Username:</td></tr>
            <tr><td><input type="text" name="uname"></td></tr>
            <tr><td>Password:</td></tr>
            <tr><td><input type="password" name="pwd"></td></tr>
            <tr><td class="button1"><input type="submit" name="submitBtn" value="Log In" class="button"></td></tr>
        </table>
    </form>
            </div>
        </div>
          
<?php include 'footer.php';
?>
<?php
if(isset($_REQUEST['submitBtn'])){
    include '_inc/dbconn.php';
    $username=$_REQUEST['uname'];
    $password=$_REQUEST['pwd'];
  
    $sql="SELECT * FROM staff WHERE email='$username' AND pwd='$password'";
    $result=mysqli_query($conn,$sql) or die(mysqli_error($conn));
    $rws=  mysqli_fetch_array($result,MYSQLI_BOTH);
    
    
    
    if($rws["email"]==$username && $rws["pwd"]==$password){
        session_start();
        $_SESSION['staff_login']=1;
        $_SESSION['staff_id']=$rws["id"];
        $_SESSION['bank']=$rws["name"];
    header('location:bank_homepage.php'); 
    }
    else
        header('location:banklogin.php');
        
}
