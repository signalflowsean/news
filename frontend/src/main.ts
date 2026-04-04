import type { InitialData } from './types';
import init from './wasm/news_engine';

const canvas = document.getElementById('engine') as HTMLCanvasElement;
// Do not get a 2D context here: the WASM engine uses this canvas for WebGL. A canvas can only have one context.
// We only get a 2D context in fallbackRender() when WASM fails to load.

const data: InitialData = (window as any).__INITIAL_DATA__ ?? { stories: [] };

/** Winit's web EventLoop cannot block like on native; it schedules the loop and throws this Error to unwind the Rust stack. Our #[wasm_bindgen(start)] runs during init(), so the throw surfaces here even when startup succeeded. */
const WINIT_WEB_CONTROL_FLOW_PREFIX = 'Using exceptions for control flow';

async function initWasm() {
  try {
    await init();
    console.log('Wasm engine initialized');
  } catch (e) {
    if (e instanceof Error && e.message.startsWith(WINIT_WEB_CONTROL_FLOW_PREFIX)) {
      console.log('Wasm engine initialized');
      return;
    }
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
