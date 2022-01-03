use actix_web::{web, App, HttpServer};
use yup_oauth2::ServiceAccountAuthenticator;
use google_youtube3::YouTube;
use hyper;
use hyper_rustls;

pub mod routes;

#[actix_web::main]
async fn main() -> std::io::Result<()> {

    let app_secret = yup_oauth2::read_service_account_key(&"service-account.json".to_string())
        .await
        .expect("service-account");
    let auth = ServiceAccountAuthenticator::builder(app_secret)
        .build()
        .await
        .expect("authenticator");
    let client = hyper::Client::builder().build(hyper_rustls::HttpsConnector::with_native_roots());
    let hub = YouTube::new(client, auth);

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(hub.clone()))
            .configure(routes::controllers::init)
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await    
}