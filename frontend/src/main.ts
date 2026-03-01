import type { InitialData } from './types';

const canvas = document.getElementById('engine') as HTMLCanvasElement;
// Do not get a 2D context here: the WASM engine uses this canvas for WebGL. A canvas can only have one context.
// We only get a 2D context in fallbackRender() when WASM fails to load.

function resize() {
  const canvasTop = canvas.getBoundingClientRect().top;
  canvas.height = Math.max(0, window.innerHeight - canvasTop);
  canvas.width = window.innerWidth;
}

window.addEventListener('resize', resize);
resize();

const data: InitialData = (window as any).__INITIAL_DATA__ ?? { stories: [] };

async function initWasm() {
  try {
    const wasm = await import('./wasm/news_engine');
    // Default export is the init: load and instantiate the WASM (run_web/start runs on init)
    await wasm.default();
    console.log('Wasm engine initialized');
  } catch (e) {
    console.warn('Wasm not available, using fallback renderer:', e);
    fallbackRender();
  }
}

function fallbackRender() {
  const ctx = canvas.getContext('2d');
  if (!ctx) return;

  ctx.fillStyle = '#1a1a2e';
  ctx.fillRect(0, 0, canvas.width, canvas.height);

  ctx.fillStyle = '#eee';
  ctx.font = '24px system-ui, sans-serif';
  ctx.textAlign = 'center';
  ctx.fillText('News Graph', canvas.width / 2, 60);

  ctx.font = '14px system-ui, sans-serif';
  ctx.fillStyle = '#888';
  ctx.fillText(
    `${data.stories.length} stories loaded`,
    canvas.width / 2,
    90
  );

  ctx.fillStyle = '#666';
  ctx.fillText(
    'Wasm engine not loaded - showing placeholder',
    canvas.width / 2,
    canvas.height / 2
  );
}

initWasm();
