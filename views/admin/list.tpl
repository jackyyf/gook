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
  <link rel="stylesheet" href="/static/css/bootstrap-datepicker3.min.css" />
  <link rel="stylesheet" href="/static/css/dataTables.bootstrap.css" />
  <link rel="stylesheet" href="/static/css/dataTables.responsive.css" />
  <script language="javascript" type="text/javascript" src="/static/js/jquery.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/jquery.placeholder.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/bootstrap-datepicker.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/jquery.dataTables.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/dataTables.bootstrap.min.js"></script>
</head>
<body>
  <nav class="navbar navbar-default navbar-fixed-top">
    <div class="container">
      <div class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li><a href="/">Home</a></li>
          <li ><a href="/book/list">Books</a></li>
          <li><a href="/orderin/list">Buy Orders</a></li>
          <li><a href="/orderout/list">Sell Orders</a></li>
          <li><a href="/bill/list">Billing</a></li>
        </ul>
        <ul class="nav navbar-nav navbar-right">
          <li><a href="/user/me">Welcome, {{.user.Name}}</a></li>
          <li><a href="/user/logout">Logout</a></li>
          {{if .user.IsAdmin}}
          <li class="active"><a href="/admin/list">Admin</a></li>
          {{end}}
        </ul>
      </div>
    </div>
  </nav>
  <div class="container">
    <div class="col-lg-12">
      <div class="panel panel-default">
        <div class="panel-heading" style="height: 50px;">
          User List
          <span style="float: right">
            <a href="/admin/new" class="btn btn-success btn-sm">Create user</a>
          </span>
        </div>
        <div class="panel-body">
          {{if .errmsg}}
          <div class="row" style="margin-top: 10px;">
            <div class="alert alert-danger alert-dismissible">
              {{.errmsg}}
            </div>
          </div>
          {{end}}
          <div class="row" style="margin-top: 10px;">
            <div class="col-sm-12">
              <table class="table table-striped table-bordered table-hover no-footer">
                <thead>
                  <tr rol="row">
                    <th style="width: 100px;">ID</th>
                    <th style="width: 150px;">Username</th>
                    <th style="width: 250px;">Real Name</th>
                    <th style="width: 100px;">Gender</th>
                    <th style="width: 200px;">Born</th>
                    <th style="width: 80px;">Admin?</th>
                  </tr>
                </thead>
                <tbody>
                  {{range $idx, $user := .users}}
                  <tr style="cursor: pointer" onclick="window.location.href='/admin/info/{{$user.ID}}'">
					<td>{{$user.ID}}</td>
                    <td>{{$user.Name}}</td>
                    <td>{{$user.RealName}}</td>
                    <td>{{if compare $user.Gender 0}}Female{{else}}Male{{end}}</td>
                    <td>{{dateformat $user.Born "2006-01-02"}}</td>
                    <td>{{$user.IsAdmin}}</td>
                  </tr>
                  {{end}}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</body>
</html>
