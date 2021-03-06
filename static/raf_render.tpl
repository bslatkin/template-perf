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
