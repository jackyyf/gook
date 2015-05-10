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
          <li class="active"><a href="/book/list">Books</a></li>
          <li><a href="/orderin/list">Buy Orders</a></li>
          <li><a href="/orderout/list">Sell Orders</a></li>
          <li><a href="/bill/list">Billing</a></li>
        </ul>
        <ul class="nav navbar-nav navbar-right">
          <li><a href="/user/me">Welcome, {{.user.Name}}</a></li>
          <li><a href="/user/logout">Logout</a></li>
          {{if .user.IsAdmin}}
          <li><a href="/admin/">Admin</a></li>
          {{end}}
        </ul>
      </div>
    </div>
  </nav>
  <div class="container">
    <div class="col-md-6 col-md-offset-3" style="margin-top: 32px;">
      <div class="login-panel panel panel-default">
        <div class="panel-heading">
          <h3 class="panel-title" style="text-align: center;">Book Info</h3>
        </div>
        <div class="panel-body">
          <form role="form" method="post" id="changeme-form" action="/book/edit/{{.book.ID}}">
            <fieldset>
              {{if .errmsg}}
              <div class="form-group">
                <div class="alert alert-danger alert-dismissible col-md-12">
                  {{.errmsg}}
                </div>
              </div>
              {{end}}
              <div class="form-group">
                <label for="book-title" class="col-md-2 control-label">Title</label>
                <div class="col-md-10">
                  <input class="form-control" placeholder="Title" name="title" type="text" value="{{.book.Name}}" id="book-title">
                </div>
              </div>
              <div class="form-group">
                <label for="book-isbn" class="col-md-2 control-label">ISBN</label>
                <div class="col-md-10">
                  <input class="form-control" placeholder="ISBN" name="isbn" type="text" value="{{.book.ISBN}}" id="book-isbn">
                </div>
              </div>
              <div class="form-group">
                <label for="book-publisher" class="col-md-2 control-label">Publisher</label>
                <div class="col-md-10">
                  <input class="form-control" placeholder="Publisher" name="publisher" type="text" value="{{.book.Publisher}}" id="book-publisher">
                </div>
              </div>
              <div class="form-group">
                <label for="book-author" class="col-md-2 control-label">Author</label>
                <div class="col-md-10">
                  <input class="form-control" placeholder="Author" name="author" type="text" value="{{.book.Author}}" id="book-author">
                </div>
              </div>
              <div class="form-group">
                <label for="book-price" class="col-md-2 control-label">Price</label>
                <div class="col-md-10">
                  <input class="form-control" placeholder="Price" name="price" type="text" value='{{printf "%.2f" .book.Price}}' id="book-price">
                </div>
              </div>
              <div class="form-group">
                <label for="book-amount" class="col-md-2 control-label">Amount</label>
                <div class="col-md-10">
                  <input class="form-control" placeholder="Amount" type="text" value='{{.book.Amount}}' id="book-amount" disabled>
                </div>
              </div>
              <input type="submit" class="btn btn-success col-md-4" value="Change Book">
              <div class="col-md-4" style="text-align: center"><a href="/book/buy/{{.book.ID}}" class="btn btn-info form-control"> Buy Book </a></div>
              <div class="col-md-4" style="text-align: center"><a href="/book/sell/{{.book.ID}}" class="btn btn-info form-control"> Sell Book </a></div>
          </fieldset>
        </form>
      </div>
    </div>
  </div>
</div>
</body>
</html>
