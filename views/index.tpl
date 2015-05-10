<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge" />
  <title>{% block title %}{% trans 'Home' %}{% endblock %} | {{ config.SITE_NAME }}</title>
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
          <li><a href="/">Home</a></li>
          <li><a href="/books">Books</a></li>
          <li><a href="/orders_in">Buy Orders</a></li>
          <li><a href="/orders_out">Sell Orders</a></li>
          <li><a href="/billing">Billing</a></li>
        </ul>
        <ul class="nav navbar-nav navbar-right">
          <li><a href="/user/me">Welcome, {{.user}}</a></li>
          <li><a href="/user/logout">Logout'</a></li>
          {{with .user.IsAdmin}}
          <li><a href="/admin/">Admin</a></li>
          {{end}}
        </ul>
      </div>
    </div>
  </nav>
  <div class="container">
    <h1>GOOK</h1>
    <p>Book management in golang.</p>
  </div>
</body>
</html>
