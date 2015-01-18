<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>DOM render</title>
  <script>
   var data = {{.}};
  </script>
  <script>
    function onLoad() {
      var x = window.chrome.loadTimes();
      if (x.firstPaintAfterLoadTime == 0) {
        window.setTimeout(onLoad, 1000);
        return;
      }
      console.log('Seconds to first paint: ' + (x.firstPaintTime - x.startLoadTime));
      console.log('Seconds to first paint after load: ' + (x.firstPaintAfterLoadTime - x.startLoadTime));
    }
  </script>
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
    <tbody id="body"></tbody>
  </table>
  <script>
    var body = document.querySelector('#body');
    for (var i = 0; i < data.length; i += 1) {
      var cat = data[i];
      var row = document.createElement('tr');

      var cell = document.createElement('td');
      cell.textContent = cat.name;
      row.appendChild(cell);

      cell = document.createElement('td');
      cell.textContent = cat.color;
      row.appendChild(cell);

      cell = document.createElement('td');
      cell.textContent = cat.sex;
      row.appendChild(cell);

      cell = document.createElement('td');
      cell.textContent = cat.legs;
      row.appendChild(cell);

      body.appendChild(row);
    }
  </script>
  <script>onLoad()</script>
  <!-- Example based on https://html.spec.whatwg.org/multipage/scripting.html#the-template-element -->
</body>
</html>
