// DO NOT EDIT ** This file was generated with the bake tool ** DO NOT EDIT //

package static

var Files = map[string]string{
	"dom_render.tpl": `<!DOCTYPE html>
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
`,

	"fragment_divs_render.tpl": `<!DOCTYPE html>
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
  <div hidden>
    <template id="row">
      <div>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
      </div>
    </template>
  </div>
  <script>
    var template = document.querySelector('#row');
    var frag = document.createDocumentFragment();
    for (var i = 0; i < data.length; i += 1) {
      var cat = data[i];
      var clone = template.content.cloneNode(true);
      var cells = clone.querySelectorAll('span');
      cells[0].textContent = cat.name;
      cells[1].textContent = cat.color;
      cells[2].textContent = cat.sex;
      cells[3].textContent = cat.legs;
      frag.appendChild(clone);
    }
    document.body.appendChild(frag);
  </script>
  <script>onLoad()</script>
  <!-- Example based on https://html.spec.whatwg.org/multipage/scripting.html#the-template-element -->
</body>
</html>
`,

	"raf_render.tpl": `<!DOCTYPE html>
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
      console.log(x);
      console.log('Seconds to first paint: ' + (x.firstPaintTime - x.startLoadTime));
      console.log('Seconds to first paint after load: ' + (x.firstPaintAfterLoadTime - x.startLoadTime));
    }
  </script>
</head>
<body>
  <div hidden>
    <template id="row">
      <div>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
      </div>
    </template>
  </div>
  <script>
    var template = document.querySelector('#row');
    var i = 0;
    function step() {
      for (var j = 0; j < 10 && i < data.length; i++, j++) {
        var cat = data[i];
        var clone = template.content.cloneNode(true);
        var cells = clone.querySelectorAll('span');
        cells[0].textContent = cat.name;
        cells[1].textContent = cat.color;
        cells[2].textContent = cat.sex;
        cells[3].textContent = cat.legs;
        document.body.appendChild(clone);
      }
      if (i < data.length) {
        window.requestAnimationFrame(step);
      } else {
        // Trigger timing
        onLoad()
      }
    }
    window.requestAnimationFrame(step);
  </script>
  <!-- Example based on https://html.spec.whatwg.org/multipage/scripting.html#the-template-element -->
</body>
</html>
`,

	"server_divs_render.tpl": `<!DOCTYPE html>
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
`,

	"server_render.tpl": `<!DOCTYPE html>
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
      var paintTimes = [];
      var loadDone = false;
      function log(e) {
        paintTimes.push(performance.now());
        if (loadDone) {
          window.removeEventListener("MozAfterPaint", log, false);
        }
      }
      window.addEventListener("MozAfterPaint", log, false);

      window.onLoad = function() {
        if (window.performance.timing.loadEventEnd != 0) {
          loadDone = true;
        }
        if (paintTimes.length == 0 || window.performance.timing.loadEventEnd == 0) {
          window.requestAnimationFrame(onLoad);
          return;
        }
        var firstPaintDelaySeconds = paintTimes[0] / 1000.0;
        var lastPaintDelaySeconds = paintTimes[paintTimes.length - 1] / 1000.0;
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
`,

	"template_tag_render.tpl": `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Template tag render</title>
  <script>
    var data = {{.}};
  </script>
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
      var paintTimes = [];
      var loadDone = false;
      function log(e) {
        paintTimes.push(performance.now());
        if (loadDone) {
          window.removeEventListener("MozAfterPaint", log, false);
        }
      }
      window.addEventListener("MozAfterPaint", log, false);

      window.onLoad = function() {
        if (window.performance.timing.loadEventEnd != 0) {
          loadDone = true;
        }
        if (!loadDone || paintTimes.length == 0) {
          window.requestAnimationFrame(onLoad);
          return;
        }
        var firstPaintDelaySeconds = paintTimes[0] / 1000.0;
        var lastPaintDelaySeconds = paintTimes[paintTimes.length - 1] / 1000.0;
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
`,

}
