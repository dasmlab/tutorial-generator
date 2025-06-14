// player/engine.js
const canvas = document.getElementById('canvas');
const ctx = canvas.getContext('2d');
let timeline = [];

fetch('/assets/preview_timeline.json')
  .then(res => res.json())
  .then(data => {
    timeline = data;
    startAnimation();
  });

function drawBox(params) {
  ctx.fillStyle = '#3df';
  ctx.fillRect(params.x, params.y, 120, 60);
  ctx.fillStyle = 'black';
  ctx.fillText(params.text || '', params.x + 10, params.y + 35);
}

function drawArrow(params) {
  // (stub logic: lines only for now)
  const from = findBox(params.from);
  const to = findBox(params.to);
  if (from && to) {
    ctx.beginPath();
    ctx.moveTo(from.x + 60, from.y + 30);
    ctx.lineTo(to.x + 60, to.y + 30);
    ctx.strokeStyle = '#f5f';
    ctx.stroke();
  }
}

// Helper to render Subtitles from vtt
function renderSubtitles(now) {
  timeline.forEach(step => {
    if (step.voice && now >= step.at && now <= step.at + 3000) {
      ctx.fillStyle = "#ffffff";
      ctx.font = "18px sans-serif";
      ctx.textAlign = "center";
      ctx.fillText(step.voice, canvas.width / 2, canvas.height - 30);
    }
  });
}


function startAnimation() {
  const start = Date.now();
  const boxes = {}; // in-memory reference by ID

  function render() {
    const now = Date.now() - start;
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    timeline.forEach(step => {
      if (now >= step.at) {
        switch (step.type) {
          case 'Box':
            boxes[step.params.id] = step.params;
            drawBox(step.params);
            break;
          case 'Arrow':
            drawArrow(step.params);
            break;
          case 'Callout':
            ctx.fillText(step.params.text, step.params.x, step.params.y);
            break;
          case 'FadeIn':
            ctx.globalAlpha = (now - step.at) / step.duration;
            break;
        }
      }
    });

    renderSubtitles(now);
    requestAnimationFrame(render);
  }

  function findBox(id) {
    return boxes[id];
  }


  render();
}

