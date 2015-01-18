<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Server divs render</title>
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
  {{range .}}
    <div>
      <span>{{.Name}}</span>
      <span>{{.Color}}</span>
      <span>{{.Sex}}</span>
      <span>{{.Legs}}</span>
    </div>
  {{end}}
  <script>onLoad()</script>
  <!-- Example based on https://html.spec.whatwg.org/multipage/scripting.html#the-template-element -->
</body>
</html>
