use axum::{
    http::{header, StatusCode},
    response::IntoResponse,
    routing::post,
    Json, Router,
};
use qrcode::QrCode;
use image::Luma;
use serde::{Deserialize, Serialize};
use std::io::Cursor;
use tokio::net::TcpListener;

/// Defines the structure of the incoming JSON request body.
#[derive(Serialize, Deserialize, Debug)]
struct TicketRequest {
    order_id: String,
    event_id: String,
    user_id: String,
    seat: String,
}

/// The main entry point for the application.
#[tokio::main]
async fn main() {
    // Build the application's router with a single route.
    let app = Router::new().route("/generate", post(generate_ticket));

    // Start the TCP listener.
    println!("Ticket Generator listening on :7000");
    let listener = TcpListener::bind("0.0.0.0:7000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

/// This handler function receives ticket data, generates a QR code,
/// and returns it as a PNG image.
async fn generate_ticket(Json(payload): Json<TicketRequest>) -> impl IntoResponse {
    // Serialize the ticket data into a JSON string to embed in the QR code.
    let ticket_data_string = match serde_json::to_string(&payload) {
        Ok(json) => json,
        Err(_) => return Err((StatusCode::INTERNAL_SERVER_ERROR, "Failed to serialize data".to_string())),
    };

    // Generate the QR code from the JSON string.
    let code = match QrCode::new(ticket_data_string.as_bytes()) {
        Ok(c) => c,
        Err(_) => return Err((StatusCode::INTERNAL_SERVER_ERROR, "Failed to create QR code".to_string())),
    };

    // Render the QR code into a grayscale image.
    let image = code.render::<Luma<u8>>().build();

    // Encode the image into a PNG format in memory.
    let mut buffer = Cursor::new(Vec::new());
    match image.write_to(&mut buffer, image::ImageFormat::Png) {
        Ok(_) => (),
        Err(_) => return Err((StatusCode::INTERNAL_SERVER_ERROR, "Failed to encode PNG".to_string())),
    };
    
    // Prepare headers for the image response.
    let mut headers = header::HeaderMap::new();
    headers.insert(header::CONTENT_TYPE, "image/png".parse().unwrap());
    
    // Return the headers and the raw image bytes.
    Ok((headers, buffer.into_inner()))
}