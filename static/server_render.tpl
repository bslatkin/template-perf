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
        <th>Name</th>
        <th>Colour</th>
        <th>Sex</th>
        <th>Legs</th>
      </tr>
    </thead>
    <tbody>
      {{range .}}
        <tr>
          <td>{{.Name}}</td>
          <td>{{.Color}}</td>
          <td>{{.Sex}}</td>
          <td>{{.Legs}}</td>
        </tr>
      {{end}}
    </tbody>
  </table>
  <!-- Example based on https://html.spec.whatwg.org/multipage/scripting.html#the-template-element -->
</body>
</html>
