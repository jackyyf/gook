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
  <script language="javascript" type="text/javascript" src="/static/js/jquery.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/jquery.placeholder.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/bootstrap-datepicker.min.js"></script>
</head>
<body>
  <nav class="navbar navbar-default navbar-fixed-top">
    <div class="container">
      <div class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li><a href="/">Home</a></li>
          <li><a href="/book/list">Books</a></li>
          <li><a href="/orderin/list">Buy Orders</a></li>
          <li><a href="/orderout/list">Sell Orders</a></li>
          <li><a href="/bill/list">Billing</a></li>
        </ul>
        <ul class="nav navbar-nav navbar-right">
          <li class="active"><a href="/user/me">Welcome, {{.user.Name}}</a></li>
          <li><a href="/user/logout">Logout</a></li>
          {{if .user.IsAdmin}}
          <li><a href="/admin/list">Admin</a></li>
          {{end}}
        </ul>
      </div>
    </div>
  </nav>
  <div class="container">
    <div class="col-md-4 col-md-offset-4" style="margin-top: 32px;">
      <div class="login-panel panel panel-default">
        <div class="panel-heading">
          <h3 class="panel-title" style="text-align: center;">Modify your info</h3>
        </div>
        <div class="panel-body">
          <form role="form" method="post" id="changeme-form">
            <fieldset>
              {{if .errmsg}}
              <div class="form-group">
                <div class="alert alert-danger alert-dismissible col-md-12">
                  {{.errmsg}}
                </div>
              </div>
              {{end}}
              <div class="form-group">
                <input class="form-control" placeholder="ID" type="text" value="{{.user.ID}}" disabled>
              </div>
              <div class="form-group">
                <input class="form-control" placeholder="Username" type="user" value="{{.user.Name}}" disabled>
              </div>
              <div class="form-group">
                <input class="form-control" placeholder="Current Password" name="password" type="password" value="">
              </div>
              <div class="form-group">
                <input class="form-control" placeholder="Realname" name="realname" type="user" value="{{.user.RealName}}">
              </div>
              <div class="form-group">
                  <select class="form-control" name="gender">
                    <option value="0" {{if .user.Gender}}{{else}}selected{{end}}>Female</option>
                    <option value="1" {{if .user.Gender}}selected{{end}}>Male</option>
                  </select>
              </div>
              <div class="form-group" style="color: black">
                <input class="form-control" placeholder="Realname" name="born" type="user" value='{{dateformat .user.Born "2006-01-02"}}' id="born">
                <script type="text/javascript">
                  $('#born').datepicker({
                    format: 'yyyy-mm-dd'
                  })
                </script>
              </div>
              <div class="form-group">
                <input class="form-control" placeholder="New Password" name="npass" type="password" value="">
              </div>
              <div class="form-group">
                <input class="form-control" placeholder="New Password again" name="nrpass" type="password" value="">
              </div>
              <input type="submit" class="btn btn-success btn-block" value="Submit">
          </fieldset>
        </form>
      </div>
    </div>
  </div>
</div>
</body>
</html>
