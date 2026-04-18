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

Create a full frontend build (js, html, css, wasm + glue)
```bash 
npm run build
```

You can build wasm separately too
```bash 
npm run build:wasm
```

You can build wasm natively and run it (assuming your cd is frontend/engine)
```bash
cargo build --release
cargo run
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
