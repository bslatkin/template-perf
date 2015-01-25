<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Server render</title>
  <script>
    if (typeof window.chrome === 'object') {
      window.onLoad = function() {
        var x = window.chrome.loadTimes();
        if (x.firstPaintAfterLoadTime == 0) {
          window.setTimeout(onLoad, 1000);
          return;
        }
        console.log('Seconds to first paint: ' + (x.firstPaintTime - x.startLoadTime));
        console.log('Seconds to last paint: ' + (x.firstPaintAfterLoadTime - x.startLoadTime));
      };
    } else {
      // For measuring Firefox
      //
      // Must set the dom.send_after_paint_to_content preference to true to
      // make this event work. See:
      //
      // https://developer.mozilla.org/en-US/docs/Web/Events/MozAfterPaint
      var firstPaint = 0;
      function log(e) {
        firstPaint = performance.now();
        window.removeEventListener("MozAfterPaint", log, false);
      }
      window.addEventListener("MozAfterPaint", log, false);

      window.onLoad = function() {
        if (window.performance.timing.loadEventEnd == 0) {
          window.setTimeout(onLoad, 1000);
          return
        }
        var lastPaintDelaySeconds = (window.performance.timing.loadEventEnd -
            window.performance.timing.navigationStart) / 1000.0;
        var firstPaintDelaySeconds = firstPaint != 0 ? firstPaint / 1000.0 : lastPaintDelaySeconds;
        console.log('Seconds to first paint: ' + firstPaintDelaySeconds);
        console.log('Seconds to last paint: ' + lastPaintDelaySeconds);
      };
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
  <script>onLoad()</script>
  <!-- Example based on https://html.spec.whatwg.org/multipage/scripting.html#the-template-element -->
</body>
</html>
