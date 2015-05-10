<!DOCTYPE html>
<html>
  <head>
  <meta charset="UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge" />
  <title>GOOK</title>
  <link rel="stylesheet" href="/static/css/bootstrap-datetimepicker.min.css" />
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
  <script language="javascript" type="text/javascript" src="/static/js/moment.js"></script>
  <script language="javascript" type="text/javascript" src="/static/js/bootstrap-datetimepicker.min.js"></script>
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
          <li class="active"><a href="/bill/list">Billing</a></li>
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
          Billings
          <div class="col-md-6" style="display: inline-block; float: right;">
            <form method="get">
              <div class="col-md-9">
                <input class="form-control" name="after" type="text" id="time-after" value="{{if .after}}{{dateformat .after "2006-01-02 15:04:05"}}{{end}}" placeholder="After this time...">
              </div>
              <input class="btn btn-md btn-info" type="submit" value="Filter">
              <script type="text/javascript">
                $('#time-after').datetimepicker({
                  //sideBySide: true,
                  format: 'YYYY-MM-DD HH:mm:ss'
                })
              </script>
            </form>
          </div>
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
                    <th style="width: 150px;">Serial</th>
                    <th style="width: 250px;">Amount</th>
                    <th style="width: 200px;">Created</th>
                  </tr>
                </thead>
                <tbody>
                  {{range $idx, $bill := .bills}}
                  <tr style="cursor: pointer" class="bill-row">
                    <td>{{$bill.ID}}</td>
                    <td class="bill-amount">{{printf "%.2f" $bill.Amount}}</td>
                    <td>{{dateformat $bill.Created "2006-01-02 15:04:05 MST"}}</td>
                  </tr>
                  {{end}}
                  <tr style="cursor: pointer; font-weight: bold; font-size: 18px;">
                    <td>Total: </td>
                    <td colspan="2" id="bill-total"></td>
                  </tr>
                </tbody>
              </table>
              <script type="text/javascript">
                var sum = 0;
                $('.bill-row > .bill-amount').each(function() {
                  var $this = $(this)
                  var amount = parseFloat($this.text())
                  sum += amount;
                  $(this).parent().addClass(amount > 0 ? 'bill-in' : 'bill-out')
                })
                $('#bill-total').text(sum.toFixed(2)).parent().addClass(sum >= 0 ? 'bill-in' : 'bill-out')
              </script>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</body>
</html>
