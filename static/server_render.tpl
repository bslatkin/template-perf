<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Server render</title>
</head>
<body>
  <table>
   <thead>
    <tr>
     <th>Name <th>Colour <th>Sex <th>Legs
   <tbody>
    {{range .}}
     <tr><td>{{.Name}}<td>{{.Color}}<td>{{.Sex}}<td>{{.Legs}}
    {{end}}
  </table>
  <!-- Example based on https://html.spec.whatwg.org/multipage/scripting.html#the-template-element -->
</body>
</html>
