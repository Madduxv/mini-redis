package protocol

/* RESP:
   *x-> An array with x elements
   $x-> A bulk string with x elements
*/

/* example Requests:
SET: {
  *3
  $3
  SET
  $4
  mykey
  $13
  Hello, World!
}

GET: {
  *2
  $3
  GET
  $4
  mykey
}

DEL: {
  *2
  $3
  DEL
  $4
  mykey
}

*/
