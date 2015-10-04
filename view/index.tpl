<!DOCTYPE html>

{% macro napt_print(row) %}
  <div class="napt_item">
    Allocate: arg.vc:{{row.BoundPort}} ⇒ Target:{{row.TargetIP}}:{{row.TargetPort}} 
    <a href="/remove/{{row.Id}}">release</a>
  </div>
{% endmacro %}

<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>mitei || index </title>
</head>
<body>
<h2>NAPT Table</h2>
  {% for row in naptlist %}
   {{ napt_print(row) }}
  {% endfor %}

<h3>Allocate</h3>
<form action='/create' method='post'>
TargetIP: <input type='text' name='TargetIP'></input><br />
TargetPort: <input type='text' name='TargetPort'></input><br />
BoundPort: <input type='text' name='BoundPort'></input><br />
<input type='submit'></input>
</form>
</body>
</html>

