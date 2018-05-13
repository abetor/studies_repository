<?php 
// print "a: [" . TRUE . "]<br>";
// echo "b: [" . FALSE . "]<br>";

// $a = 2; $b = 3;
// if ($a > $b)  echo "$a is greater than $b<br>";
// if ($a < $b)  echo "$a is less than $b<br>";
// if ($a >= $b) echo "$a is greater than or equal to $b<br>";
// if ($a <= $b) echo "$a is less than or equal to $b<br>";

// if ($bank_balance < 100)
// {
//     $money         = 1000;
//     $bank_balance += $money;
// }
// elseif ($bank_balance > 200)
// {
//     $savings      += 100;
//     $bank_balance -= 100;
// }
// else
// {
//     $savings      += 50;
//     $bank_balance -= 50;
// }

// switch ($page):
//     case "Home":
//         echo "You selected Home";
//         break;

//     etc...

//     case "Links":
//         echo "You selected Links";
//         break;
// endswitch;

// $fp = fopen("text.txt", 'wb');

// for ($j = 0 ; $j < 100 ; ++$j)
// {
//   $written = fwrite($fp, "data");

//   if ($written == FALSE) break;
// }

// fclose($fp);

$j = 10;

while ($j > -10)
{
  $j--;

  if ($j == 0) continue;

  echo (10 / $j) . "<br>";
}

?>