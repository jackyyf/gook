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
          <li class="active"><a href="/book/list">Books</a></li>
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
  <div class="container">
    <div class="col-lg-12">
      <div class="panel panel-default">
        <div class="panel-heading" style="height: 50px;">
          Imported books
          <span style="float: right">
            <a href="/book/new" class="btn btn-success btn-sm">Import a new Book</a>
          </span>
        </div>
        <div class="panel-body">
          <div class="row">
            <form method="get" action="/book/search">
              <div class="col-sm-2">
                <input class="form-control" placeholder="ISBN" name="isbn" type="text" value="{{.isbn}}">
              </div>
              <div class="col-sm-3">
                <input class="form-control" placeholder="Book Title" name="title" type="text" value="{{.title}}">
              </div>
              <div class="col-sm-3">
                <input class="form-control" placeholder="Author" name="author" type="text" value="{{.author}}">
              </div>
              <div class="col-sm-3">
                <input class="form-control" placeholder="Publisher" name="publisher" type="text" value="{{.publisher}}">
              </div>
              <div class="col-sm-1" style="padding-right: 5px;">
                <input class="btn btn-info" type="submit" value="Search">
              </div>
            </form>
          </div>
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
                    <th style="width: 150px;">ISBN</th>
                    <th style="width: 250px;">Name</th>
                    <th style="width: 200px;">Author</th>
                    <th style="width: 200px;">Publisher</th>
                    <th style="width: 80px;">Price</th>
                    <th style="width: 80px;">Amount</th>
                  </tr>
                </thead>
                <tbody>
                  {{range $idx, $book := .books}}
                  <tr style="cursor: pointer" onclick="window.location.href='/book/info/{{$book.ID}}'">
                    <td>{{$book.ISBN}}</td>
                    <td>{{$book.Name}}</td>
                    <td>{{$book.Author}}</td>
                    <td>{{$book.Publisher}}</td>
                    <td>{{printf "%.2f" $book.Price}}</td>
                    <td>{{$book.Amount}}</td>
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
