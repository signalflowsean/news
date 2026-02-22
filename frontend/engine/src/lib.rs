use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern "C" {
    #[wasm_bindgen(js_namespace = console)]
    fn log(s: &str);
}

macro_rules! console_log {
    ($($t:tt)*) => (log(&format_args!($($t)*).to_string()))
}

/// Graph layout and rendering engine.
/// Future: Will use wgpu/WebGPU for GPU-accelerated rendering.
pub struct GraphRenderer {
    // Future fields:
    // device: wgpu::Device,
    // queue: wgpu::Queue,
    // surface: wgpu::Surface,
}

#[wasm_bindgen]
pub fn init() {
    console_log!("news-engine: initialized");
    // Future: Initialize wgpu device, surface, pipelines
}

#[wasm_bindgen]
pub fn render() {
    console_log!("news-engine: render called");
    // Future: Run graph layout algorithm, render nodes/edges to canvas via WebGPU
}

#[wasm_bindgen]
pub fn set_data(_json: &str) {
    console_log!("news-engine: data received");
    // Future: Parse story/event data, build graph structure
}
