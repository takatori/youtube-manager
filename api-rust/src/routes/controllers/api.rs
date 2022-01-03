use actix_web::{get, web, HttpResponse, Responder};
use google_youtube3::YouTube;
use serde::Deserialize;

#[get("/api/popular")]
pub async fn fetch_most_popular_videos(hub: web::Data<YouTube>) -> impl Responder {
    let response = hub
        .videos()
        .list(&vec!["id,snippet".into()])
        .max_results(3)
        .chart("mostPopular")
        .doit()
        .await;

    let body = response
        .map(|res| serde_json::to_string(&res.1).unwrap())
        .unwrap();

    HttpResponse::Ok()
        .content_type("application/json")
        .body(body)
}

#[get("/api/related/{id}")]
pub async fn fetch_related_videos(
    path: web::Path<String>,
    hub: web::Data<YouTube>,
) -> impl Responder {
    let id = path.into_inner();

    let response = hub
        .search()
        .list(&vec!["id,snippet".into()])
        .related_to_video_id(&id)
        .max_results(3)
        .doit()
        .await;

    match response {
        Ok(_) => println!("Success"),
        Err(ref e) => println!("{:?}", e),
    }

    let body = response
        .map(|res| serde_json::to_string(&res.1).unwrap())
        .unwrap();

    HttpResponse::Ok()
        .content_type("application/json")
        .body(body)
}

#[derive(Deserialize)]
pub struct SearchQuery {
    pub q: String,
}

#[get("/api/search")]
pub async fn search_video(query: web::Query<SearchQuery>, hub: web::Data<YouTube>) -> impl Responder {

    println!("{:?}", &query.q);

    let response = hub
        .search()
        .list(&vec!["id,snippet".into()])
        .q(&query.q)
        .max_results(3)
        .doit()
        .await;

    let body = response
        .map(|res| serde_json::to_string(&res.1).unwrap())
        .unwrap();

    HttpResponse::Ok()
        .content_type("application/json")
        .body(body)
}

#[get("/api/video/{id}")]
pub async fn get_video(path: web::Path<(u32,)>, hub: web::Data<YouTube>) -> impl Responder {
    HttpResponse::Ok()
}
