use news_engine::run;

// This is the entry point for the engine when built as a standalone binary
fn main() {
    if let Err(e) = run() {
        eprintln!("Error: {e:?}");
        std::process::exit(1);
    }
}