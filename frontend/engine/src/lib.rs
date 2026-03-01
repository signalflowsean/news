use wasm_bindgen::prelude::*;
use web_sys::console;

#[wasm_bindgen]
pub fn init() {
    console::log_1(&format!("news-engine: initialized").into());
    // Future: Initialize wgpu device, surface, pipelines
}

#[wasm_bindgen]
pub fn render() {
    console::log_1(&format!("news-engine: render called").into());
    // Future: Run graph layout algorithm, render nodes/edges to canvas via WebGPU
}

#[wasm_bindgen]
pub fn set_data(_json: &str) {
    console::log_1(&format!("news-engine: data received").into());
    // Future: Parse story/event data, build graph structure
}
