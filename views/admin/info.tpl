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
    <div class="col-md-6 col-md-offset-3" style="margin-top: 32px;">
      <div class="login-panel panel panel-default">
        <div class="panel-heading">
          <h3 class="panel-title" style="text-align: center;">User Info</h3>
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
                <input class="form-control" placeholder="ID" type="text" value="{{.nuser.ID}}" disabled>
              </div>
              <div class="form-group">
                <input class="form-control" placeholder="Username" type="user" value="{{.nuser.Name}}" disabled>
              </div>
              <div class="form-group">
                <input class="form-control" placeholder="New Password" name="password" type="password" value="">
              </div>
              <div class="form-group">
                <input class="form-control" placeholder="Realname" name="realname" type="user" value="{{.nuser.RealName}}">
              </div>
              <div class="form-group">
                  <select class="form-control" name="gender">
                    <option value="0" {{if .nuser.Gender}}{{else}}selected{{end}}>Female</option>
                    <option value="1" {{if .nuser.Gender}}selected{{end}}>Male</option>
                  </select>
              </div>
              <div class="form-group" style="color: black">
                <input class="form-control" placeholder="Born" name="born" type="user" value='{{dateformat .nuser.Born "2006-01-02"}}' id="born">
                <script type="text/javascript">
                  $('#born').datepicker({
                    format: 'yyyy-mm-dd'
                  })
                </script>
              </div>
              <div class="form-group">
                  <label for="admin-select">Admin?</label>
                  <select class="form-control" name="admin" id="admin-select">
                    <option value="0" {{if .nuser.IsAdmin}}{{else}}selected{{end}}>No</option>
                    <option value="1" {{if .nuser.IsAdmin}}selected{{end}}>Yes</option>
                  </select>
              </div>
              <input type="submit" class="btn btn-success col-md-6" value="Submit">
              <div class="col-md-6" style="text-align: center"><a href="/admin/remove/{{.nuser.ID}}" class="btn btn-danger form-control"> Detele </a></div>
          </fieldset>
        </form>
      </div>
    </div>
  </div>
</div>
</body>
</html>
