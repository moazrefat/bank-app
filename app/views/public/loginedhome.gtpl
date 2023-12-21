<!doctype html>
<html lang="ja">
 
<head>
  <center>
      <!-- Required meta tags -->
      <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  
      <!-- Bootstrap CSS -->
      <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
      <link rel="stylesheet" href="./assets/css/style.css" type="text/css"> 
  
      <title>Home Page</title>
      <div id="header_title">
      <p class="display-1 text-center">BankApp</p>
      </div>

      <ul id="nav">
      <li><a href="/home">Home</a></li>
      <li><a href="/profile">Profile</a></li>
      <li><a href="/logout">Logout</a></li>
      <ul>
  </center>
</head>
<br>
<body>
  <table>
    <tr>
      <th>User Information</th>
    </tr>
    <tr>
      <td>Name:{{.Name}}</td>
    </tr>
    <tr>
      <td>Age:{{.Age}}</td>
    </tr>
    <tr>
      <td>Email:{{.Email}}</td>
    </tr>
    <tr>
      <td>Address:{{.Address}}</td>
    </tr>
  </table>

</body>
</html>

