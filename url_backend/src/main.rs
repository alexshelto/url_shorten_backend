use actix_web::{get, web, App, HttpServer, Responder};
use r2d2_sqlite::SqliteConnectionManager;

#[get("/hello/{name}")]
async fn greet(name: web::Path<String>) -> impl Responder {
    format!("Hello {name}!")
}


#[get("/{link}")]
async fn view_link(link: web::Path<String>) -> impl Responder {
    format!("link: {link}")
}



#[actix_web::main] 
async fn main() -> std::io::Result<()> {
    // Work on SQL connection 
    let manager = SqliteConnectionManager::file("../db/data.db");



    HttpServer::new(|| {
        App::new()
            .service(greet)
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}

