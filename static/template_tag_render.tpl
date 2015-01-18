<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Template tag render</title>
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
    </thead>
    <tbody>
      <template id="row">
        <tr>
          <td></td>
          <td></td>
          <td></td>
          <td></td>
        </tr>
      </template>
    </tbody>
  </table>
  <script>
    var template = document.querySelector('#row');
    for (var i = 0; i < data.length; i += 1) {
      var cat = data[i];
      var clone = template.content.cloneNode(true);
      var cells = clone.querySelectorAll('td');
      cells[0].textContent = cat.name;
      cells[1].textContent = cat.color;
      cells[2].textContent = cat.sex;
      cells[3].textContent = cat.legs;
      template.parentNode.appendChild(clone);
    }
  </script>
  <script>onLoad()</script>
  <!-- Example based on https://html.spec.whatwg.org/multipage/scripting.html#the-template-element -->
</body>
</html>
