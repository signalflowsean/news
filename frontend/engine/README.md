# Engine (Rust/Wasm)

Graph layout and rendering engine. Will use wgpu/WebGPU for GPU-accelerated visualization.

## Prerequisites

- Rust (stable)
- wasm-pack

Install wasm-pack:
```bash
cargo install wasm-pack
```

## Build

```bash
wasm-pack build --target web --out-dir ../src/wasm
```

Or from the parent frontend directory:
```bash
npm run build:wasm
```

## Output

Generates in `frontend/src/wasm/`:
- `news_engine.js` - JS bindings
- `news_engine_bg.wasm` - Wasm binary
- `news_engine.d.ts` - TypeScript types

## Future Plans

- Implement force-directed graph layout algorithm
- Add wgpu/WebGPU rendering pipeline
- GPU-accelerated node/edge rendering
- Interactive pan/zoom/selection
