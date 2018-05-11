<?php
$username = "Fred Smith";
echo $username;
echo "<br>";
$current_user = $username;
echo $current_user;

define("ROOT_LOCATION", "usr/local/www");
$directory = ROOT_LOCATION;

function longdate($timestamp){
    return date("l F jS Y", $timestamp);
}

function Counter(){
    static $count = 0;
    echo $count;
    $count++;
}