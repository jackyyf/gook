<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge" />
  <title>GOOK</title>
  <link rel="stylesheet" href="/static/css/bootstrap.min.css" />
  <link rel="stylesheet" href="/static/css/font-awesome.min.css" />
  <link rel="stylesheet" href="/static/css/style.css" />
    <link rel="stylesheet" href="/static/css/timeline.css" />
  <script language="javascript" type="text/javascript" src="/static/js/jquery.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/jquery.placeholder.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/bootstrap.min.js"></script>
</head>
<body>
  <nav class="navbar navbar-default navbar-fixed-top">
    <div class="container">
      <div class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li class="active"><a href="/">Home</a></li>
          <li><a href="/book/list">Books</a></li>
          <li><a href="/orderin/list">Buy Orders</a></li>
          <li><a href="/orderout/list">Sell Orders</a></li>
          <li><a href="/bill/list">Billing</a></li>
        </ul>
        <ul class="nav navbar-nav navbar-right">
          <li><a href="/user/me">Welcome, {{.user.Name}}</a></li>
          <li><a href="/user/logout">Logout</a></li>
          {{if .user.IsAdmin}}
          <li><a href="/admin/list">Admin</a></li>
          {{end}}
        </ul>
      </div>
    </div>
  </nav>
  <div class="container" style="font-size: 24px;">
    <h1 style="margin-left: 1%; font-size: 96px;">GOOK</h1>
    <p style="margin-left: 2%;">Book management system in golang.</p>
  </div>
</body>
</html>
